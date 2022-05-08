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

// StoreMainchainWithdrawCallback stores the receipt to own database for future check from ProvideReceiptSignatureCallback
func (l *RoninListener) StoreMainchainWithdrawCallback(tx types.ITransaction, data []byte) error {
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
		RoninAddress:         receipt.Ronin.Addr.Hex(),
		RoninTokenAddress:    receipt.Ronin.TokenAddr.Hex(),
		TokenErc:             receipt.Info.Erc,
		TokenId:              receipt.Info.Id.Int64(),
		// TODO: should we hex this value, since int64 might not cover big number more than 10^18
		TokenQuantity: receipt.Info.Quantity.Int64(),
		BlockNumber:   int64(ronEvent.Raw.BlockNumber),
	})
}

func (l *RoninListener) ProvideReceiptSignatureCallback(tx types.ITransaction, data []byte) error {
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
	// otherwise, create a task for submitting signature
	// get chainID
	chainId, err := l.GetChainID()
	if err != nil {
		return err
	}
	// create task and store to database
	withdrawalTask := &models.Task{
		ChainId:   hexutil.EncodeBig(chainId),
		Type:      types.WITHDRAWAL_TASK,
		Data:      common.Bytes2Hex(data),
		Retries:   0,
		Status:    types.STATUS_PENDING,
		LastError: "",
		CreatedAt: time.Now().Unix(),
	}
	return l.store.GetTaskStore().Save(withdrawalTask)
}

func (l *RoninListener) DepositRequestedCallback(tx types.ITransaction, data []byte) error {
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
	result, err := caller.DepositVote(nil, chainId, ethEvent.Receipt.Id)
	if err != nil {
		return err
	}
	if result.Status == types.VoteStatusExecuted {
		return nil
	}
	// create task and store to database
	depositTask := &models.Task{
		ChainId:   hexutil.EncodeBig(chainId),
		Type:      types.DEPOSIT_TASK,
		Data:      common.Bytes2Hex(data),
		Retries:   0,
		Status:    types.STATUS_PENDING,
		LastError: "",
		CreatedAt: time.Now().Unix(),
	}
	return l.store.GetTaskStore().Save(depositTask)
}

func (l *RoninListener) WithdrewCallback(tx types.ITransaction, data []byte) error {
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
	// create ack task if result is false
	if !result {
		// get chainID
		chainId, err := l.GetChainID()
		if err != nil {
			return err
		}
		return l.store.GetTaskStore().Save(&models.Task{
			ChainId:   hexutil.EncodeBig(chainId),
			Type:      types.ACK_WITHDREW_TASK,
			Data:      common.Bytes2Hex(data),
			Retries:   0,
			Status:    types.STATUS_PENDING,
			LastError: "",
			CreatedAt: time.Now().Unix(),
		})
	}
	return nil
}

type RoninCallBackJob struct {
	*EthCallbackJob
}
