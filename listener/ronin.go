package listener

import (
	"context"
	"errors"
	"github.com/axieinfinity/bridge-v2/contract"
	"math/big"
	"time"

	bridgeUtils "github.com/axieinfinity/bridge-v2/utils"

	ethGovernance "github.com/axieinfinity/bridge-contracts/generated_contracts/ethereum/governance"
	"github.com/ethereum/go-ethereum/crypto"

	roninTrustedOrganization "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/trusted_organization"

	"github.com/axieinfinity/bridge-contracts/generated_contracts/ethereum/gateway"
	gateway2 "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/gateway"
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

func NewRoninListener(ctx context.Context, cfg *bridgeCore.LsConfig, helpers utils.Utils, store bridgeCoreStores.MainStore, pool *bridgeCore.Pool) (*RoninListener, error) {
	listener, err := NewEthereumListener(ctx, cfg, helpers, store, pool)
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

	if err = l.utilsWrapper.UnpackLog(*ronGatewayAbi, ronEvent, "MainchainWithdrew", data); err != nil {
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

func (l *RoninListener) provideReceiptSignature(fromChainId *big.Int, tx bridgeCore.Transaction, data []byte, isAgain bool) error {
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

func (l *RoninListener) ProvideReceiptSignatureCallback(fromChainId *big.Int, tx bridgeCore.Transaction, data []byte) error {
	return l.provideReceiptSignature(fromChainId, tx, data, false)
}

func (l *RoninListener) ProvideReceiptSignatureAgainCallback(fromChainId *big.Int, tx bridgeCore.Transaction, data []byte) error {
	return l.provideReceiptSignature(fromChainId, tx, data, true)
}

func (l *RoninListener) DepositRequestedCallback(fromChainId *big.Int, tx bridgeCore.Transaction, data []byte) error {
	log.Info("[RoninListener] DepositRequestedCallback", "tx", tx.GetHash().Hex())
	// get chainID
	chainId, err := l.GetChainID()
	if err != nil {
		return err
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

func (l *RoninListener) isRelayerNode() (bool, error) {
	relayerSign := l.GetRelayerSign()
	if relayerSign == nil {
		log.Warn("The current node is not set relayer key")
		return false, nil
	}

	ethClient := l.GetListener(bridgeUtils.Ethereum).GetEthClient()
	ethGovernanceCaller, err := ethGovernance.NewGovernanceCaller(common.HexToAddress(l.config.Contracts[task.ETH_GOVERNANCE_CONTRACT]), ethClient)
	if err != nil {
		return false, err
	}

	var ret [32]byte
	copy(ret[:], crypto.Keccak256([]byte("RELAYER_ROLE")))
	addr := relayerSign.GetAddress()
	isRelayer, err := ethGovernanceCaller.HasRole(nil, ret, addr)
	if err != nil {
		return false, err
	}

	return isRelayer, nil
}

func (l *RoninListener) isTrustedNode() error {
	roninTrustedCaller, err := roninTrustedOrganization.NewTrustedOrganizationCaller(common.HexToAddress(l.config.Contracts[task.TRUSTED_ORGANIZATION_CONTRACT]), l.client)
	if err != nil {
		return err
	}

	addr := l.GetVoterSign().GetAddress()
	weight, err := roninTrustedCaller.GetBridgeVoterWeight(nil, addr)
	if err != nil {
		return err
	}
	log.Debug("[RoninListener][isTrustedNode] Trusted node info", "weight", weight)

	if weight.Cmp(big.NewInt(0)) == 1 {
		return nil
	}

	return errors.New("current node is not trusted node")
}

func (l *RoninListener) BridgeOperatorSetUpdatedCallback(fromChainId *big.Int, tx bridgeCore.Transaction, data []byte) error {
	log.Info("[RoninListener][BridgeOperatorSetUpdatedCallback] Received new event", "tx", tx.GetHash().Hex())

	if l.GetVoterSign() == nil {
		log.Info("[RoninListener][BridgeOperatorSetUpdatedCallback] Voter key is missing")
		return nil
	}

	if err := l.isTrustedNode(); err != nil {
		log.Info("[RoninListener][BridgeOperatorSetUpdatedCallback] The current node is not trusted node")
		return nil
	}

	// Get chainID
	chainId, err := l.GetChainID()
	if err != nil {
		return err
	}

	bridgeOperatorsUpdated := &models.Task{
		ChainId:         hexutil.EncodeBig(chainId),
		FromChainId:     hexutil.EncodeBig(fromChainId),
		FromTransaction: tx.GetHash().Hex(),
		Type:            task.VOTE_BRIDGE_OPERATORS_TASK,
		Data:            common.Bytes2Hex(data),
		Retries:         0,
		Status:          task.STATUS_PENDING,
		LastError:       "",
		CreatedAt:       time.Now().Unix(),
	}

	return l.bridgeStore.GetTaskStore().Save(bridgeOperatorsUpdated)
}

func (l *RoninListener) BridgeOperatorsApprovedCallback(fromChainId *big.Int, tx bridgeCore.Transaction, data []byte) error {
	log.Info("[RoninListener][BridgeOperatorsApprovedCallback] Received new event", "tx", tx.GetHash().Hex())

	isRelayer, err := l.isRelayerNode()
	if err != nil {
		return err
	}
	if !isRelayer {
		log.Info("[RoninListener][BridgeOperatorsApprovedCallback] The current node is not relayer")
		return nil
	}

	if l.GetRelayerSign() == nil {
		log.Error("[RoninListener][BridgeOperatorSetUpdatedCallback] Relayer key is missing")
		return nil
	}

	// Get chainID
	chainId, err := l.GetChainID()
	if err != nil {
		return err
	}

	bridgeOperatorsApproved := &models.Task{
		ChainId:         hexutil.EncodeBig(chainId),
		FromChainId:     hexutil.EncodeBig(fromChainId),
		FromTransaction: tx.GetHash().Hex(),
		Type:            task.RELAY_BRIDGE_OPERATORS_TASK,
		Data:            common.Bytes2Hex(data),
		Retries:         0,
		Status:          task.STATUS_PENDING,
		LastError:       "",
		CreatedAt:       time.Now().Unix(),
	}

	return l.bridgeStore.GetTaskStore().Save(bridgeOperatorsApproved)
}

func (l *RoninListener) WithdrewCallback(fromChainId *big.Int, tx bridgeCore.Transaction, data []byte) error {
	log.Info("[RoninListener] WithdrewCallback", "tx", tx.GetHash().Hex())
	// Unpack event data
	ethEvent := new(gateway.GatewayWithdrew)
	ethGatewayAbi, err := gateway.GatewayMetaData.GetAbi()
	if err != nil {
		return err
	}

	if err = l.utilsWrapper.UnpackLog(*ethGatewayAbi, ethEvent, "Withdrew", data); err != nil {
		return err
	}
	log.Info("[RoninListener][WithdrewCallback] result of calling MainchainWithdrew function", "receiptId", ethEvent.Receipt.Id.Int64(), "tx", tx.GetHash().Hex())
	// get chainID
	chainId, err := l.GetChainID()
	if err != nil {
		return err
	}
	ackWithdrewTask := &models.Task{
		ChainId:         hexutil.EncodeBig(chainId),
		FromChainId:     hexutil.EncodeBig(fromChainId),
		FromTransaction: tx.GetHash().Hex(),
		Type:            task.ACK_WITHDREW_TASK,
		Data:            common.Bytes2Hex(data),
		Retries:         0,
		Status:          task.STATUS_PENDING,
		LastError:       "",
		CreatedAt:       time.Now().Unix(),
	}
	return l.bridgeStore.GetTaskStore().Save(ackWithdrewTask)
}

func (l *RoninListener) RandomSeedRequestedCallback(fromChainId *big.Int, tx bridgeCore.Transaction, data []byte) error {
	log.Info("[RoninListener] RandomSeedRequestedCallback", "tx", tx.GetHash().Hex())
	// if vrf is not enabled, we don't need to create task for it to prevent further errors or panics
	if task.VRFConfig == nil {
		return nil
	}
	event := new(contract.RoninVRFCoordinatorRandomSeedRequested)
	contractAbi, err := contract.RoninVRFCoordinatorMetaData.GetAbi()
	if err != nil {
		return err
	}
	if err = l.utilsWrapper.UnpackLog(*contractAbi, event, "RandomSeedRequested", data); err != nil {
		return err
	}
	// get chainID
	chainId, err := l.GetChainID()
	if err != nil {
		return err
	}
	vrfTask := &models.Task{
		ChainId:         hexutil.EncodeBig(chainId),
		FromChainId:     hexutil.EncodeBig(fromChainId),
		FromTransaction: tx.GetHash().Hex(),
		Type:            task.VRF_RANDOM_SEED_REQUEST,
		Data:            common.Bytes2Hex(data),
		Retries:         0,
		Status:          task.STATUS_PENDING,
		LastError:       "",
		CreatedAt:       time.Now().Unix(),
	}
	return l.bridgeStore.GetTaskStore().Save(vrfTask)
}

type RoninCallBackJob struct {
	*EthCallbackJob
}
