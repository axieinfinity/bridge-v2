package listener

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/math"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/signer/core"
	"math/big"
	"time"

	"github.com/axieinfinity/bridge-contracts/generated_contracts/ethereum/gateway"
	governance "github.com/axieinfinity/bridge-contracts/generated_contracts/ethereum/governance"
	gateway2 "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/gateway"
	governance2 "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/governance"
	trustedOrg2 "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/trusted_org"
	bridgeCore "github.com/axieinfinity/bridge-core"
	bridgeCoreModels "github.com/axieinfinity/bridge-core/models"
	bridgeCoreStores "github.com/axieinfinity/bridge-core/stores"
	"github.com/axieinfinity/bridge-core/utils"
	"github.com/axieinfinity/bridge-v2/models"
	"github.com/axieinfinity/bridge-v2/stores"
	"github.com/axieinfinity/bridge-v2/task"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"
)

const oneHour = 3600

type RoninListener struct {
	*EthereumListener
	bridgeStore stores.BridgeStore
}

func NewRoninListener(ctx context.Context, cfg *bridgeCore.LsConfig, helpers utils.Utils, store bridgeCoreStores.MainStore) (*RoninListener, error) {
	listener, err := NewEthereumListener(ctx, cfg, helpers, store)
	if err != nil {
		panic(err)
	}
	l := &RoninListener{EthereumListener: listener}
	l.bridgeStore = stores.NewBridgeStore(store.GetDB())
	if err != nil {
		return nil, err
	}
	return l, nil
}

func (l *RoninListener) NewJobFromDB(job *bridgeCoreModels.Job) (bridgeCore.JobHandler, error) {
	return newJobFromDB(l, job)
}

// StoreMainchainWithdrawCallback stores the receipt to own database for future check from ProvideReceiptSignatureCallback
func (l *RoninListener) StoreMainchainWithdrawCallback(fromChainId *big.Int, tx bridgeCore.Transaction, data []byte) error {
	log.Info("[RoninListener] StoreMainchainWithdrawCallback", "tx", tx.GetHash().Hex())
	ronEvent := new(gateway2.GatewayMainchainWithdrew)
	ronGatewayAbi, err := gateway2.GatewayMetaData.GetAbi()
	if err != nil {
		return err
	}
	if err = ronGatewayAbi.UnpackIntoInterface(ronEvent, "MainchainWithdrew", data); err != nil {
		return err
	}
	receipt := ronEvent.Receipt
	// store ronEvent to database at withdrawal
	return l.bridgeStore.GetWithdrawalStore().Save(&models.Withdrawal{
		WithdrawalId:         receipt.Id.Int64(),
		ExternalAddress:      receipt.Mainchain.Addr.Hex(),
		ExternalTokenAddress: receipt.Mainchain.TokenAddr.Hex(),
		ExternalChainId:      receipt.Mainchain.ChainId.Int64(),
		RoninAddress:         receipt.Ronin.Addr.Hex(),
		RoninTokenAddress:    receipt.Ronin.TokenAddr.Hex(),
		TokenErc:             receipt.Info.Erc,
		TokenId:              receipt.Info.Id.Int64(),
		TokenQuantity:        receipt.Info.Quantity.String(),
		Transaction:          tx.GetHash().Hex(),
	})
}

func (l *RoninListener) IsUpTodate() bool {
	latestBlock, err := l.GetLatestBlock()
	if err != nil {
		log.Error("[RoninListener][IsUpTodate] error while get latest block", "err", err, "listener", l.GetName())
		return false
	}
	// true if timestamp is within 1 hour
	distance := uint64(time.Now().Unix()) - latestBlock.GetTimestamp()
	if distance > uint64(oneHour) {
		log.Info("Node is not up-to-date, keep waiting...", "distance (s)", distance, "listener", l.GetName())
		return false
	}
	return true
}

func (l *RoninListener) ProvideReceiptSignatureCallback(fromChainId *big.Int, tx bridgeCore.Transaction, data []byte) error {
	// check database if receipt exist then do nothing
	// Unpack event from data
	ronEvent := new(gateway2.GatewayMainchainWithdrew)
	ronGatewayAbi, err := gateway2.GatewayMetaData.GetAbi()
	if err != nil {
		return err
	}
	if err = ronGatewayAbi.UnpackIntoInterface(ronEvent, "MainchainWithdrew", data); err != nil {
		return err
	}
	receipt := ronEvent.Receipt

	// try getting withdrawal data from database by receipt.id
	withdrawal, _ := l.bridgeStore.GetWithdrawalStore().GetWithdrawalById(receipt.Id.Int64())
	if withdrawal != nil && withdrawal.ID > 0 {
		return nil
	}
	// try checking on smart contract
	// create caller
	caller, err := gateway2.NewGatewayCaller(common.HexToAddress(l.config.Contracts[task.GATEWAY_CONTRACT]), l.client)
	if err != nil {
		return err
	}
	result, err := caller.MainchainWithdrew(nil, receipt.Id)
	if err != nil {
		return err
	}
	log.Info("[RoninListener][ProvideReceiptSignatureCallback] result of calling MainchainWithdrew function", "result", result, "receiptId", receipt.Id.Int64(), "tx", tx.GetHash().Hex())
	if !result {
		// otherwise, create a task for submitting signature
		// get chainID
		chainId, err := l.GetChainID()
		if err != nil {
			return err
		}
		// create task and store to database
		withdrawalTask := &models.Task{
			ChainId:         hexutil.EncodeBig(chainId),
			FromChainId:     hexutil.EncodeBig(fromChainId),
			FromTransaction: tx.GetHash().Hex(),
			Type:            task.WITHDRAWAL_TASK,
			Data:            common.Bytes2Hex(data),
			Retries:         0,
			Status:          task.STATUS_PENDING,
			LastError:       "",
			CreatedAt:       time.Now().Unix(),
		}
		return l.bridgeStore.GetTaskStore().Save(withdrawalTask)
	}
	return nil
}

func (l *RoninListener) DepositRequestedCallback(fromChainId *big.Int, tx bridgeCore.Transaction, data []byte) error {
	log.Info("[RoninListener] DepositRequestedCallback", "tx", tx.GetHash().Hex())
	// check whether deposit is done or not
	// Unpack event data
	ethEvent := new(gateway.GatewayDepositRequested)
	ethGatewayAbi, err := gateway.GatewayMetaData.GetAbi()
	if err != nil {
		return err
	}

	if err = ethGatewayAbi.UnpackIntoInterface(ethEvent, "DepositRequested", data); err != nil {
		return err
	}
	// create caller
	caller, err := gateway2.NewGatewayCaller(common.HexToAddress(l.config.Contracts[task.GATEWAY_CONTRACT]), l.client)
	if err != nil {
		return err
	}
	// get chainID
	chainId, err := l.GetChainID()
	if err != nil {
		return err
	}

	// check if deposit has been executed or not
	result, err := caller.DepositVote(nil, ethEvent.Receipt.Mainchain.ChainId, ethEvent.Receipt.Id)
	if err != nil {
		return err
	}
	log.Info("[RoninListener][DepositRequestedCallback] result of calling DepositVote function", "status", result.Status, "finalHash", common.Bytes2Hex(result.FinalHash[:]), "receiptId", ethEvent.Receipt.Id, "tx", tx.GetHash().Hex())
	if result.Status == task.VoteStatusExecuted {
		return nil
	}

	// check if current validator has been voted for this deposit or not
	voted, err := caller.DepositVoted(nil, ethEvent.Receipt.Mainchain.ChainId, ethEvent.Receipt.Id, l.GetValidatorSign().GetAddress())
	if err != nil {
		return err
	}
	log.Info("[RoninListener][DepositRequestedCallback] result of calling DepositVoted function", "voted", voted, "receiptId", ethEvent.Receipt.Id, "tx", tx.GetHash().Hex())
	if voted {
		return nil
	}

	// create task and store to database
	depositTask := &models.Task{
		ChainId:         hexutil.EncodeBig(chainId),
		FromChainId:     hexutil.EncodeBig(fromChainId),
		FromTransaction: tx.GetHash().Hex(),
		Type:            task.DEPOSIT_TASK,
		Data:            common.Bytes2Hex(data),
		Retries:         0,
		Status:          task.STATUS_PENDING,
		LastError:       "",
		CreatedAt:       time.Now().Unix(),
	}
	return l.bridgeStore.GetTaskStore().Save(depositTask)
}

func (l *RoninListener) BridgeOperatorsUpdatedCallback(fromChainId *big.Int, tx bridgeCore.Transaction, data []byte) error {
	log.Info("[RoninListener] BridgeOperatorsUpdatedCallback", "tx", tx.GetHash().Hex())
	// Unpack event data
	ronEvent := new(gateway2.GatewayBridgeOperatorsUpdated)
	ronGatewayAbi, err := gateway2.GatewayMetaData.GetAbi()
	if err != nil {
		return err
	}

	if err = ronGatewayAbi.UnpackIntoInterface(ronEvent, "BridgeOperatorsUpdated", data); err != nil {
		return err
	}

	// create caller
	transactor, err := governance2.NewGatewayTransactor(common.HexToAddress(l.config.Contracts[task.GATEWAY_CONTRACT]), l.client)
	if err != nil {
		return err
	}

	bridgeOperatorsBallotTypes := core.TypedData{
		Types: core.Types{
			"BridgeOperatorsBallot": []core.Type{
				{
					Name: "period", Type: "uint256",
				},
				{
					Name: "operators", Type: "address[]",
				},
			},
		},
		PrimaryType: "RoninGatewayV2",
		Domain: core.TypedDataDomain{
			Name:              "RoninGatewayV2",
			Version:           "2",
			ChainId:           math.NewHexOrDecimal256(fromChainId.Int64()),
			VerifyingContract: l.config.Contracts[task.GATEWAY_CONTRACT],
		},
		Message: core.TypedDataMessage{
			"period":    ronEvent.Period,
			"operators": ronEvent.BridgeOperators,
		},
	}
	signatures, err := l.utilsWrapper.SignTypedData(bridgeOperatorsBallotTypes, l.GetValidatorSign())
	if err != nil {
		return err
	}

	_, err = l.utilsWrapper.SendContractTransaction(l.EthereumListener.GetValidatorSign(), fromChainId, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return transactor.VoteBridgeOperatorsBySignatures(opts, ronEvent.Period, ronEvent.BridgeOperators, string(signatures))
	})

	return err
}

func (l *RoninListener) BridgeOperatorsApprovedCallback(fromChainId *big.Int, tx bridgeCore.Transaction, data []byte) error {
	ronEvent := new(gateway2.GatewayBridgeOperatorsApproved)
	ronGatewayAbi, err := gateway2.GatewayMetaData.GetAbi()
	if err != nil {
		return err
	}

	if err = ronGatewayAbi.UnpackIntoInterface(ronEvent, "BridgeOperatorsUpdated", data); err != nil {
		return err
	}

	// create caller
	trustedCaller2, err := trustedOrg2.NewGatewayCaller(common.HexToAddress(l.config.Contracts[task.GATEWAY_CONTRACT]), l.client)
	if err != nil {
		return err
	}

	trustedOrgs, err := trustedCaller2.GetAllTrustedOrganizations(nil)
	if err != nil {
		return err
	}

	var voters []common.Address
	for _, node := range trustedOrgs {
		voters = append(voters, node.BridgeVoter)
	}

	governanceCaller2, err := governance2.NewGatewayCaller(common.HexToAddress(l.config.Contracts[task.GATEWAY_CONTRACT]), l.client)
	if err != nil {
		return err
	}

	signatures, err := governanceCaller2.GetBridgeOperatorVotingSignatures(nil, ronEvent.Period, voters)
	if err != nil {
		return err
	}

	governanceTransactor, err := governance.NewGatewayTransactor(common.HexToAddress(l.config.Contracts[task.GATEWAY_CONTRACT]), l.client)
	if err != nil {
		return err
	}

	_, err = l.utilsWrapper.SendContractTransaction(l.EthereumListener.GetValidatorSign(), fromChainId, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return governanceTransactor.RelayBridgeOperators(opts, ronEvent.Period, ronEvent.BridgeOperators, signatures)
	})

	return err
}

func (l *RoninListener) WithdrewCallback(fromChainId *big.Int, tx bridgeCore.Transaction, data []byte) error {
	log.Info("[RoninListener] WithdrewCallback", "tx", tx.GetHash().Hex())
	// Unpack event data
	ethEvent := new(gateway.GatewayWithdrew)
	ethGatewayAbi, err := gateway.GatewayMetaData.GetAbi()
	if err != nil {
		return err
	}

	if err = ethGatewayAbi.UnpackIntoInterface(ethEvent, "Withdrew", data); err != nil {
		return err
	}
	// create caller
	caller, err := gateway2.NewGatewayCaller(common.HexToAddress(l.config.Contracts[task.GATEWAY_CONTRACT]), l.client)
	if err != nil {
		return err
	}
	result, err := caller.MainchainWithdrew(nil, ethEvent.Receipt.Id)
	if err != nil {
		return err
	}
	log.Info("[RoninListener][WithdrewCallback] result of calling MainchainWithdrew function", "result", result, "receiptId", ethEvent.Receipt.Id.Int64(), "tx", tx.GetHash().Hex())
	// create ack task if result is false
	if !result {
		// get chainID
		chainId, err := l.GetChainID()
		if err != nil {
			return err
		}
		return l.bridgeStore.GetTaskStore().Save(&models.Task{
			ChainId:         hexutil.EncodeBig(chainId),
			FromChainId:     hexutil.EncodeBig(fromChainId),
			FromTransaction: tx.GetHash().Hex(),
			Type:            task.ACK_WITHDREW_TASK,
			Data:            common.Bytes2Hex(data),
			Retries:         0,
			Status:          task.STATUS_PENDING,
			LastError:       "",
			CreatedAt:       time.Now().Unix(),
		})
	}
	return nil
}

type RoninCallBackJob struct {
	*EthCallbackJob
}
