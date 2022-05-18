package listener

import (
	"context"
	"github.com/axieinfinity/bridge-v2/generated_contracts/ethereum/gateway"
	gateway2 "github.com/axieinfinity/bridge-v2/generated_contracts/ronin/gateway"
	"github.com/axieinfinity/bridge-v2/internal/models"
	"github.com/axieinfinity/bridge-v2/internal/task"
	"github.com/axieinfinity/bridge-v2/internal/types"
	"github.com/axieinfinity/bridge-v2/internal/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"time"
)

type RoninListener struct {
	*EthereumListener
	task types.ITask
}

func NewRoninListener(ctx context.Context, cfg *types.LsConfig, helpers utils.IUtils, store types.IMainStore, prepareJobChan chan types.IJob) (*RoninListener, error) {
	listener, err := NewEthereumListener(ctx, cfg, helpers, store, prepareJobChan)
	if err != nil {
		panic(err)
	}
	l := &RoninListener{EthereumListener: listener}
	l.task, err = task.NewRoninTask(l, l.utilsWrapper)
	if err != nil {
		return nil, err
	}
	return l, nil
}

func (l *RoninListener) Start() {
	go l.task.Start()
}

func (l *RoninListener) GetTask() types.ITask {
	return l.task
}

// StoreMainchainWithdrawCallback stores the receipt to own database for future check from ProvideReceiptSignatureCallback
func (l *RoninListener) StoreMainchainWithdrawCallback(fromChainId *big.Int, tx types.ITransaction, data []byte) error {
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
	return l.store.GetWithdrawalStore().Save(&models.Withdrawal{
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

func (l *RoninListener) ProvideReceiptSignatureCallback(fromChainId *big.Int, tx types.ITransaction, data []byte) error {
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
	withdrawal, _ := l.store.GetWithdrawalStore().GetWithdrawalById(receipt.Id.Int64())
	if withdrawal != nil && withdrawal.ID > 0 {
		return nil
	}
	// try checking on smart contract
	// create caller
	caller, err := gateway2.NewGatewayCaller(common.HexToAddress(l.config.Contracts[types.GATEWAY_CONTRACT]), l.client)
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
			Type:            types.WITHDRAWAL_TASK,
			Data:            common.Bytes2Hex(data),
			Retries:         0,
			Status:          types.STATUS_PENDING,
			LastError:       "",
			CreatedAt:       time.Now().Unix(),
		}
		return l.store.GetTaskStore().Save(withdrawalTask)
	}
	return nil
}

func (l *RoninListener) DepositRequestedCallback(fromChainId *big.Int, tx types.ITransaction, data []byte) error {
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
	caller, err := gateway2.NewGatewayCaller(common.HexToAddress(l.config.Contracts[types.GATEWAY_CONTRACT]), l.client)
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
	if result.Status == types.VoteStatusExecuted {
		return nil
	}

	// check if current validator has been voted for this deposit or not
	voted, err := caller.DepositVoted(nil, ethEvent.Receipt.Mainchain.ChainId, ethEvent.Receipt.Id, crypto.PubkeyToAddress(l.validatorKey.PublicKey))
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
		Type:            types.DEPOSIT_TASK,
		Data:            common.Bytes2Hex(data),
		Retries:         0,
		Status:          types.STATUS_PENDING,
		LastError:       "",
		CreatedAt:       time.Now().Unix(),
	}
	return l.store.GetTaskStore().Save(depositTask)
}

func (l *RoninListener) WithdrewCallback(fromChainId *big.Int, tx types.ITransaction, data []byte) error {
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
	caller, err := gateway2.NewGatewayCaller(common.HexToAddress(l.config.Contracts[types.GATEWAY_CONTRACT]), l.client)
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
		return l.store.GetTaskStore().Save(&models.Task{
			ChainId:         hexutil.EncodeBig(chainId),
			FromChainId:     hexutil.EncodeBig(fromChainId),
			FromTransaction: tx.GetHash().Hex(),
			Type:            types.ACK_WITHDREW_TASK,
			Data:            common.Bytes2Hex(data),
			Retries:         0,
			Status:          types.STATUS_PENDING,
			LastError:       "",
			CreatedAt:       time.Now().Unix(),
		})
	}
	return nil
}

type RoninCallBackJob struct {
	*EthCallbackJob
}
