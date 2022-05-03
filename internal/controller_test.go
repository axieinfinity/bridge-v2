package internal

import (
	"context"
	"github.com/axieinfinity/bridge-v2/internal/listener"
	utilsmocks "github.com/axieinfinity/bridge-v2/internal/mocks/utils"
	"github.com/axieinfinity/bridge-v2/internal/types"
	"github.com/axieinfinity/bridge-v2/internal/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/sha3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"hash"
	"math/big"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

var (
	_, b, _, _ = runtime.Caller(0)
	abiPath    = filepath.Join(filepath.Dir(b), "../contracts/common/BridgeAdmin.abi")
)
var sampleAbi abi.ABI

func init() {
	u := &utils.Utils{}
	a, _ := u.LoadAbi(abiPath)
	sampleAbi = *a
}

type mockUtils struct {
	*utilsmocks.IUtils
	real *utils.Utils
}

func NewMockUtils(t *testing.T) *mockUtils {
	return &mockUtils{utilsmocks.NewIUtils(t), &utils.Utils{}}
}

func (m *mockUtils) Invoke(any interface{}, name string, args ...interface{}) (reflect.Value, error) {
	return m.real.Invoke(any, name, args...)
}

func (m *mockUtils) LoadAbi(path string) (*abi.ABI, error) {
	return m.real.LoadAbi(path)
}

func LoadControllerConfig() *types.Config {
	return &types.Config{
		Listeners: map[string]*types.LsConfig{
			"TestListener": {
				Name:           "firstListener",
				RpcUrl:         "",
				LoadInterval:   1,
				SafeBlockRange: uint64(0),
				Secret: types.Secret{
					Validator: "8843ebcb1021b00ae9a644db6617f9c6d870e5fd53624cefe374c1d2d710fd06",
					Relayer:   "8843ebcb1021b00ae9a644db6617f9c6d870e5fd53624cefe374c1d2d710fd06",
				},
				Subscriptions: map[string]*types.Subscribe{
					"test1": {
						From: "0x0000000000000000000000000000000000000abc",
						To:   "0x0000000000000000000000000000000000000deF",
						Type: types.LogEvent,
						Handler: &types.Handler{
							ABIPath: abiPath,
							Name:    "ProposalExecuted",
						},
						CallBacks: map[string]string{
							"TestListener": "CallFuncA",
						},
					},
				},
			},
		},
	}
}

var ch = make(chan struct{}, 2)

type TestListener struct {
	*listener.EthereumListener
}

func (c *Controller) InitTestListener(ctx context.Context, lsConfig *types.LsConfig, store types.IMainStore) *TestListener {
	l, err := listener.NewEthereumListener(ctx, lsConfig, c.utilWrapper, store)
	if err != nil {
		panic(err)
	}
	if l == nil {
		panic("cannot get listener")
	}
	return &TestListener{l}
}

func (l *TestListener) Start() {
	println("TestListener start")
}

func (l *TestListener) CallFuncA(args ...interface{}) error {
	ch <- struct{}{}
	return nil
}

/**
{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"proposalHash","type":"bytes32"}],"name":"ProposalExecuted","type":"event"}
*/
func TestListenEvent(t *testing.T) {
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlInfo, log.StreamHandler(os.Stderr, log.TerminalFormat(true))))
	// create event data from abi
	event := sampleAbi.Events["ProposalExecuted"]
	var key [32]byte
	copy(key[:], "foo")
	args, err := event.Inputs.Pack(key)
	if err != nil {
		t.Fatal(err)
	}
	input := append(event.ID.Bytes(), args...)
	receipt := &ethtypes.Receipt{
		Status: 1,
		Logs: []*ethtypes.Log{
			{
				Address: common.Address{},
				Data:    input,
			},
		},
	}

	cfg := LoadControllerConfig()

	// rpcMocks mocks events get ethClient, BlockByNumber, TransactionReceipt and ChainID from rpc
	rpcMocks := utilsmocks.NewEthClient(t)
	firstBlock := makeBlock(1000)
	var nilHeight *big.Int
	rpcMocks.On("BlockByNumber", mock.Anything, nilHeight).Return(firstBlock, nil).Twice()
	rpcMocks.On("BlockByNumber", mock.Anything, big.NewInt(1001)).Return(makeBlock(1001), nil).Once()

	rpcMocks.On("TransactionReceipt", mock.Anything, mock.Anything).Return(receipt, nil)
	rpcMocks.On("ChainID", mock.Anything).Return(big.NewInt(1), nil)

	// config helpers
	helpers := NewMockUtils(t)
	helpers.On("NewEthClient", mock.Anything).Return(rpcMocks, nil)

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	// init controller
	controller, err := New(cfg, db, helpers)
	if err != nil {
		t.Fatal(err)
	}
	err = controller.Start()
	if err != nil {
		t.Fatal(err)
	}

	for {
		select {
		case <-ch:
			controller.Close()
			return
		}
	}
}

// testHasher is the helper tool for transaction/receipt list hashing.
// The original hasher is trie, in order to get rid of import cycle,
// use the testing hasher instead.
type testHasher struct {
	hasher hash.Hash
}

func newHasher() *testHasher {
	return &testHasher{hasher: sha3.NewLegacyKeccak256()}
}

func (h *testHasher) Reset() {
	h.hasher.Reset()
}

func (h *testHasher) Update(key, val []byte) {
	h.hasher.Write(key)
	h.hasher.Write(val)
}

func (h *testHasher) Hash() common.Hash {
	return common.BytesToHash(h.hasher.Sum(nil))
}

func makeBlock(height int64) *ethtypes.Block {
	var (
		key, _   = crypto.GenerateKey()
		txs      = make([]*ethtypes.Transaction, 1)
		receipts = make([]*ethtypes.Receipt, len(txs))
		signer   = ethtypes.LatestSigner(params.TestChainConfig)
	)
	number := big.NewInt(height)
	header := &ethtypes.Header{
		Difficulty: big.NewInt(3),
		Number:     number,
		GasLimit:   12345678,
		GasUsed:    1476322,
		Time:       9876543,
		Extra:      []byte("coolest block on chain"),
	}
	for i := range txs {
		amount := math.BigPow(2, int64(i))
		price := big.NewInt(300000)
		data := make([]byte, 100)
		tx := ethtypes.NewTransaction(uint64(i), common.HexToAddress("0xdef"), amount, 123457, price, data)
		signedTx, err := ethtypes.SignTx(tx, signer, key)
		if err != nil {
			panic(err)
		}
		txs[i] = signedTx
		receipts[i] = ethtypes.NewReceipt(make([]byte, 32), false, tx.Gas())
	}
	return ethtypes.NewBlock(header, txs, nil, nil, newHasher())
}
