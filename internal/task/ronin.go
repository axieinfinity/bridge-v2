package task

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/axieinfinity/bridge-v2/generated_contracts/ethereum/gateway"
	gateway2 "github.com/axieinfinity/bridge-v2/generated_contracts/ronin/gateway"
	"github.com/axieinfinity/bridge-v2/internal/models"
	"github.com/axieinfinity/bridge-v2/internal/types"
	"github.com/axieinfinity/bridge-v2/internal/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"math/big"
	"time"
)

const (
	defaultLimitRecords = 50
	defaultMaxTry       = 5
)

var defaultTaskInterval = time.Minute

type RoninTask struct {
	ctx        context.Context
	cancelFunc context.CancelFunc

	util utils.IUtils

	listener types.IListener
	store    types.IMainStore

	taskInterval    time.Duration
	txCheckInterval time.Duration
	secret          types.Secret

	client    *ethclient.Client
	contracts map[string]string

	validator *ecdsa.PrivateKey
	relayer   *ecdsa.PrivateKey

	chainId *big.Int
}

func NewRoninTask(listener types.IListener, util utils.IUtils) (*RoninTask, error) {
	config := listener.Config()
	client, err := ethclient.Dial(config.RpcUrl)
	if err != nil {
		return nil, err
	}
	chainId, err := listener.GetChainID()
	if err != nil {
		return nil, err
	}
	newCtx, cancelFunc := context.WithCancel(listener.Context())
	task := &RoninTask{
		ctx:             newCtx,
		cancelFunc:      cancelFunc,
		listener:        listener,
		store:           listener.GetStore(),
		taskInterval:    defaultTaskInterval,
		txCheckInterval: defaultTaskInterval,
		secret:          config.Secret,
		contracts:       config.Contracts,
		client:          client,
		chainId:         chainId,
		util:            util,
		validator:       listener.(types.IEthListener).GetValidatorKey(),
		relayer:         listener.(types.IEthListener).GetRelayerKey(),
	}
	if config.TaskInterval > 0 {
		task.taskInterval = config.TaskInterval * time.Second
	}
	if config.TransactionCheckPeriod > 0 {
		task.txCheckInterval = config.TransactionCheckPeriod * time.Second
	}
	return task, nil
}

func (r *RoninTask) Start() {
	for {
		if err := r.process(); err != nil {
			log.Error("[RoninTask] error while process tasks", "err", err)
		}
		select {
		case <-r.ctx.Done():
			r.client.Close()
			r.Close()
			return
		}
	}
}

func (r *RoninTask) Close() {
	r.cancelFunc()
}

func (r *RoninTask) GetListener() types.IListener {
	return r.listener
}

func (r *RoninTask) process() error {
	tasks, err := r.store.GetTaskStore().GetPendingTasks(r.GetListener().GetName(), defaultLimitRecords)
	if err != nil {
		return err
	}

	bulkDepositTask := NewBulkTask(r.client, r.store, r.chainId, r.validator, common.HexToAddress(r.contracts[types.GATEWAY_CONTRACT]), r.txCheckInterval, defaultMaxTry, types.DEPOSIT_TASK, r.util)
	bulkSubmitWithdrawalSignaturesTask := NewBulkTask(r.client, r.store, r.chainId, r.validator, common.HexToAddress(r.contracts[types.GATEWAY_CONTRACT]), r.txCheckInterval, defaultMaxTry, types.WITHDRAWAL_TASK, r.util)
	ackWithdrewTasks := NewBulkTask(r.client, r.store, r.chainId, r.validator, common.HexToAddress(r.contracts[types.GATEWAY_CONTRACT]), r.txCheckInterval, defaultMaxTry, types.ACK_WITHDREW_TASK, r.util)
	for _, task := range tasks {
		// collect tasks for bulk deposits
		bulkDepositTask.collectTask(task)

		// collect tasks for bulk withdrawal signature
		bulkSubmitWithdrawalSignaturesTask.collectTask(task)

		// collect tasks for ...
		ackWithdrewTasks.collectTask(task)
	}
	// wait for bulk deposit finish
	if err = <-bulkDepositTask.send(); err != nil {
		// TODO: log here...
	}

	// wait for bulk submit signatures finish
	if err = <-bulkSubmitWithdrawalSignaturesTask.send(); err != nil {
		// TODO: log here...
	}

	// wait for all ack tasks be finished
	if err = <-ackWithdrewTasks.send(); err != nil {
		// TODO: log here...
	}

	// wait a specific time
	<-time.NewTimer(r.taskInterval).C
	return nil
}

type BulkTask struct {
	util            utils.IUtils
	tasks           []*models.Task
	store           types.IMainStore
	validator       *ecdsa.PrivateKey
	client          *ethclient.Client
	contractAddress common.Address
	chainId         *big.Int
	ticker          *time.Ticker
	maxTry          int
	taskType        string
}

func NewBulkTask(client *ethclient.Client, store types.IMainStore, chainId *big.Int, validator *ecdsa.PrivateKey, contractAddress common.Address, ticker time.Duration, maxTry int, taskType string, util utils.IUtils) *BulkTask {
	return &BulkTask{
		util:            util,
		tasks:           make([]*models.Task, 0),
		store:           store,
		validator:       validator,
		client:          client,
		contractAddress: contractAddress,
		chainId:         chainId,
		ticker:          time.NewTicker(ticker),
		maxTry:          maxTry,
		taskType:        taskType,
	}
}

func (r *BulkTask) collectTask(t *models.Task) {
	if t.Type == r.taskType {
		r.tasks = append(r.tasks, t)
	}
}

func (r *BulkTask) send() (errCh chan error) {
	var (
		successTasks, failedTasks []*models.Task
		err                       error
	)
	switch r.taskType {
	case types.DEPOSIT_TASK:
		successTasks, failedTasks = r.sendDepositTransaction()
	case types.WITHDRAWAL_TASK:
		successTasks, failedTasks = r.sendWithdrawalSignaturesTransaction()
	case types.ACK_WITHDREW_TASK:
		successTasks, failedTasks = r.SendAckTransactions()
	}
	errCh <- err
	if err = updateTasks(r.store, successTasks, types.STATUS_DONE); err != nil {
		// TODO: log here...
	}
	if err = updateTasks(r.store, failedTasks, types.STATUS_FAILED); err != nil {
		// TODO: log here...
	}
	return
}

func (r *BulkTask) sendDepositTransaction() ([]*models.Task, []*models.Task) {
	var (
		successTasks, failedTasks []*models.Task
		receipts                  []gateway2.TransferReceipt
	)

	for _, t := range r.tasks {
		ethEvent := new(gateway.GatewayDepositRequested)
		ethGatewayAbi, err := gateway.GatewayMetaData.GetAbi()
		if err != nil {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
			continue
		}

		data := common.Hex2Bytes(t.Data)
		if err = ethGatewayAbi.UnpackIntoInterface(ethEvent, "DepositRequested", data); err != nil {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
			continue
		}
		// append task into success tasks
		successTasks = append(successTasks, t)

		// append new receipt into receipts slice
		receipts = append(receipts, gateway2.TransferReceipt{
			Id:   ethEvent.Receipt.Id,
			Kind: ethEvent.Receipt.Kind,
			Mainchain: gateway2.TokenOwner{
				Addr:      ethEvent.Receipt.Mainchain.Addr,
				TokenAddr: ethEvent.Receipt.Mainchain.TokenAddr,
				ChainId:   ethEvent.Receipt.Mainchain.ChainId,
			},
			Ronin: gateway2.TokenOwner{
				Addr:      ethEvent.Receipt.Ronin.Addr,
				TokenAddr: ethEvent.Receipt.Ronin.TokenAddr,
				ChainId:   ethEvent.Receipt.Ronin.ChainId,
			},
			Info: gateway2.TokenInfo{
				Erc:      ethEvent.Receipt.Info.Erc,
				Id:       ethEvent.Receipt.Info.Id,
				Quantity: ethEvent.Receipt.Info.Quantity,
			},
		})
	}
	// send tx with above receipts
	c, err := gateway2.NewGatewayTransactor(r.contractAddress, r.client)
	if err != nil {
		// append all success tasks into failed tasks
		for _, t := range successTasks {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
		}
		return nil, failedTasks
	}
	tx, err := r.util.SendContractTransaction(r.validator, r.chainId, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return c.BulkDepositFor(opts, receipts)
	})
	if err != nil {
		for _, t := range successTasks {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
		}
		return nil, failedTasks
	}
	if err = r.util.SubscribeTransactionReceipt(r.client, tx, r.ticker, r.maxTry); err != nil {
		for _, t := range successTasks {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
		}
		return nil, failedTasks
	}
	return successTasks, failedTasks
}

func (r *BulkTask) sendWithdrawalSignaturesTransaction() (successTasks []*models.Task, failedTasks []*models.Task) {

	var (
		ids        []*big.Int
		signatures [][]byte
	)

	for _, t := range r.tasks {
		// Unpack event from data
		ronEvent := new(gateway2.GatewayMainchainWithdrew)
		ronGatewayAbi, err := gateway2.GatewayMetaData.GetAbi()
		if err != nil {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
			continue
		}
		if err = ronGatewayAbi.UnpackIntoInterface(ronEvent, "MainchainWithdrew", common.Hex2Bytes(t.Data)); err != nil {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
			continue
		}
		receipt := ronEvent.Receipt
		typedData := apitypes.TypedData{
			Types: apitypes.Types{
				"EIP712Domain": []apitypes.Type{{Name: "verifyingContract", Type: "address"}},
				"SubmitWithdrawalSignatures": []apitypes.Type{
					{Name: "_withdrawal", Type: "uint256"},
				},
			},
			Domain: apitypes.TypedDataDomain{
				VerifyingContract: r.contractAddress.Hex(),
			},
			PrimaryType: "SubmitWithdrawalSignatures",
			Message: apitypes.TypedDataMessage{
				"_withdrawals": fmt.Sprintf("%#d", receipt.Id),
			},
		}
		sigs, err := r.util.SignTypedData(typedData, r.validator)
		if err != nil {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
			continue
		}
		signatures = append(signatures, sigs)
		ids = append(ids, receipt.Id)
		successTasks = append(successTasks, t)
	}
	c, err := gateway2.NewGatewayTransactor(r.contractAddress, r.client)
	if err != nil {
		// append all success tasks into failed tasks
		for _, t := range successTasks {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
		}
		return nil, failedTasks
	}
	tx, err := r.util.SendContractTransaction(r.validator, r.chainId, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return c.BulkSubmitWithdrawalSignatures(opts, ids, signatures)
	})
	if err != nil {
		// append all success tasks into failed tasks
		for _, t := range successTasks {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
		}
		return nil, failedTasks
	}
	if err = r.util.SubscribeTransactionReceipt(r.client, tx, r.ticker, r.maxTry); err != nil {
		// append all success tasks into failed tasks
		for _, t := range successTasks {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
		}
		return nil, failedTasks
	}
	return
}

func (r *BulkTask) SendAckTransactions() (successTasks []*models.Task, failedTasks []*models.Task) {
	var ids []*big.Int
	c, err := gateway2.NewGatewayTransactor(r.contractAddress, r.client)
	if err != nil {
		for _, t := range r.tasks {
			t.LastError = err.Error()
		}
		return nil, r.tasks
	}
	for _, t := range r.tasks {
		// Unpack event data
		ethEvent := new(gateway.GatewayWithdrew)
		ethGatewayAbi, err := gateway.GatewayMetaData.GetAbi()
		if err != nil {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
			continue
		}

		if err = ethGatewayAbi.UnpackIntoInterface(ethEvent, "Withdrew", common.Hex2Bytes(t.Data)); err != nil {
			// TODO: log here...
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
			continue
		}
		ids = append(ids, ethEvent.Receipt.Id)
		successTasks = append(successTasks, t)
	}
	tx, err := r.util.SendContractTransaction(r.validator, r.chainId, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return c.BulkAcknowledgeMainchainWithdrew(nil, ids)
	})
	if err != nil {
		// append all success tasks into failed tasks
		for _, t := range successTasks {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
		}
		return nil, failedTasks
	}
	if err = r.util.SubscribeTransactionReceipt(r.client, tx, r.ticker, r.maxTry); err != nil {
		// append all success tasks into failed tasks
		for _, t := range successTasks {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
		}
		return nil, failedTasks
	}
	return
}

func updateTasks(store types.IMainStore, tasks []*models.Task, status string) error {
	// update tasks with given status
	// note: if task.retries < 10 then retries++ and status still be processing
	for _, t := range tasks {
		if status == types.STATUS_FAILED {
			if t.Retries+1 >= 10 {
				t.Status = status
			} else {
				t.Retries += 1
			}
		} else {
			t.Status = status
		}
		if err := store.GetTaskStore().Update(t); err != nil {
			log.Error("error while update task", "id", t.ID)
		}
	}
	return nil
}
