package task

import (
	"context"
	"math/big"
	"time"

	bridgeCore "github.com/axieinfinity/bridge-core"
	bridgeCoreModels "github.com/axieinfinity/bridge-core/models"
	bridgeCoreStore "github.com/axieinfinity/bridge-core/stores"
	bridgeCoreStores "github.com/axieinfinity/bridge-core/stores"
	"github.com/axieinfinity/bridge-core/utils"
	bridgeCoreUtils "github.com/axieinfinity/bridge-core/utils"
	"github.com/axieinfinity/bridge-v2/stores"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
)

type MockRoninTask struct {
	listener       bridgeCore.Listener
	client         *backends.SimulatedBackend
	store          stores.BridgeStore
	chainId        *big.Int
	contracts      map[string]string
	releaseTasksCh chan int
	util           utils.Utils
}

func newMockRoninTask(client *backends.SimulatedBackend, listener bridgeCore.Listener, db *gorm.DB, util utils.Utils) *MockRoninTask {
	config := listener.Config()
	chainId, _ := listener.GetChainID()

	return &MockRoninTask{
		listener:       listener,
		client:         client,
		store:          stores.NewBridgeStore(db),
		chainId:        chainId,
		util:           util,
		releaseTasksCh: make(chan int, defaultLimitRecords),
		contracts:      config.Contracts,
	}
}

func (m MockRoninTask) Start() {
	//TODO implement me
	panic("implement me")
}

func (m MockRoninTask) Close() {
	//TODO implement me
	panic("implement me")
}

func (m MockRoninTask) GetListener() bridgeCore.Listener {
	//TODO implement me
	panic("implement me")
}

func (m MockRoninTask) SetLimitQuery(limit int) {
	//TODO implement me
	panic("implement me")
}

type MockListener struct {
	ctx           context.Context
	utilsWrapper  utils.Utils
	bridgeStore   stores.BridgeStore
	chainId       *big.Int
	config        *bridgeCore.LsConfig
	client        *backends.SimulatedBackend
	validatorSign bridgeCoreUtils.ISign
}

func (m MockListener) GetVoterSign() bridgeCoreUtils.ISign {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) AddListeners(m2 map[string]bridgeCore.Listener) {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) GetListener(s string) bridgeCore.Listener {
	//TODO implement me
	panic("implement me")
}

func newMockListener(client *backends.SimulatedBackend, ctx context.Context, cfg *bridgeCore.LsConfig, helpers utils.Utils, store bridgeCoreStores.MainStore) *MockListener {
	l := &MockListener{
		ctx:          ctx,
		bridgeStore:  stores.NewBridgeStore(store.GetDB()),
		utilsWrapper: helpers,
		chainId:      hexutil.MustDecodeBig(cfg.ChainId),
		config:       cfg,
		client:       client,
	}

	if cfg.Secret.BridgeOperator != nil {
		validatorSign, err := bridgeCoreUtils.NewSignMethod(cfg.Secret.BridgeOperator)
		if err != nil {
			panic(err)
		}
		l.validatorSign = validatorSign
	}

	return l
}

func (m MockListener) GetName() string {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) GetStore() bridgeCoreStore.MainStore {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) Config() *bridgeCore.LsConfig {
	//TODO implement me
	return m.config
}

func (m MockListener) Period() time.Duration {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) GetSafeBlockRange() uint64 {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) GetCurrentBlock() bridgeCore.Block {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) GetLatestBlock() (bridgeCore.Block, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) GetLatestBlockHeight() (uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) GetBlock(height uint64) (bridgeCore.Block, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) GetBlockWithLogs(height uint64) (bridgeCore.Block, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) GetChainID() (*big.Int, error) {
	//TODO implement me
	return m.chainId, nil
}

func (m MockListener) GetReceipt(hash common.Hash) (*ethtypes.Receipt, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) Context() context.Context {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) GetSubscriptions() map[string]*bridgeCore.Subscribe {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) UpdateCurrentBlock(block bridgeCore.Block) error {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) SaveCurrentBlockToDB() error {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) SaveTransactionsToDB(txs []bridgeCore.Transaction) error {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) GetListenHandleJob(subscriptionName string, tx bridgeCore.Transaction, eventId string, data []byte) bridgeCore.JobHandler {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) SendCallbackJobs(listeners map[string]bridgeCore.Listener, subscriptionName string, tx bridgeCore.Transaction, inputData []byte) {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) NewJobFromDB(job *bridgeCoreModels.Job) (bridgeCore.JobHandler, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) Start() {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) Close() {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) IsDisabled() bool {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) SetInitHeight(u uint64) {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) GetInitHeight() uint64 {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) GetEthClient() utils.EthClient {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) GetTasks() []bridgeCore.TaskHandler {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) GetTask(index int) bridgeCore.TaskHandler {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) AddTask(handler bridgeCore.TaskHandler) {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) IsUpTodate() bool {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) SetPrepareJobChan(handlers chan bridgeCore.JobHandler) {
	//TODO implement me
	panic("implement me")
}

func (m MockListener) GetValidatorSign() utils.ISign {
	return m.validatorSign
}

func (m MockListener) GetRelayerSign() utils.ISign {
	//TODO implement me
	panic("implement me")
}
