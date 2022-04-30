package listener

import (
	"context"
	"fmt"
	"github.com/axieinfinity/bridge-v2/generated_contracts/ethereum/gateway"
	"github.com/axieinfinity/bridge-v2/internal/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"math/big"
	"reflect"
	"testing"
	"time"
)

type mockUtils struct{}

func (m *mockUtils) Invoke(any interface{}, name string, args ...interface{}) (reflect.Value, error) {
	return reflect.Value{}, nil
}

func (m *mockUtils) LoadAbi(path string) (*abi.ABI, error) {
	return nil, nil
}

func (m *mockUtils) GetArguments(a abi.ABI, name string, data []byte, isInput bool) (abi.Arguments, error) {
	return nil, nil
}
func (m *mockUtils) UnpackToInterface(a abi.ABI, name string, data []byte, isInput bool, v interface{}) error {
	return nil
}

func (m *mockUtils) Title(text string) string {
	return ""
}

var ch = make(chan struct{}, 2)

type TestListener struct {
	name   string
	period time.Duration
	ctx    context.Context
}

func (c Controller) InitTestListener(ctx context.Context, lsConfig *LsConfig) *TestListener {
	return &TestListener{
		name:   lsConfig.Name,
		period: lsConfig.LoadInterval,
		ctx:    ctx,
	}
}

func (c Controller) InitTestListener1(ctx context.Context, lsConfig *LsConfig) *TestListener {
	return &TestListener{
		name:   lsConfig.Name,
		period: lsConfig.LoadInterval,
		ctx:    ctx,
	}
}

func (t *TestListener) Context() context.Context {
	return t.ctx
}

func (t *TestListener) Period() time.Duration {
	return t.period
}

func (t *TestListener) FetchData() {
	println(fmt.Sprintf("calling fetch data on %s", t.name))
	ch <- struct{}{}
}

func TestController(t *testing.T) {
	c, err := New(&Config{
		Listeners: map[string]*LsConfig{
			"TestListener": {
				Name:           "firstListener",
				RpcUrl:         "",
				LoadInterval:   1 * time.Second,
				SafeBlockRange: uint64(0),
				Subscriptions: map[string]*Subscribe{
					"test1": {
						From: "0xabc",
						To:   "0xdef",
						Type: TxEvent,
						//Handler:  &Handler{
						//	ABI:    "",
						//	Method: "",
						//},
						//CallBack: &Handler{
						//
						//},
					},
				},
			},
			"TestListener1": {
				Name:           "secondListener",
				RpcUrl:         "",
				LoadInterval:   1 * time.Second,
				SafeBlockRange: uint64(0),
			},
		}})
	c.utilWrapper = &mockUtils{}

	require.NoError(t, err)
	println("starting controller")
	err = c.Start()
	require.NoError(t, err)
	<-ch
	<-ch
	c.Close()
}

func TestUnpackSmartContractData(t *testing.T) {
	u := &utils.Utils{}
	a, err := u.LoadAbi("/Users/mac/coding/go/src/github.com/axieinfinity/bridge-v2/contracts/ethereum/MainchainGatewayV2.abi")
	//a, err := gateway.EthereumMetaData.GetAbi()
	require.NoError(t, err)
	//contract := bind.NewBoundContract(common.Address{}, a, nil, nil, nil)
	req := []gateway.TransferRequest{
		{
			RecipientAddr: common.HexToAddress("0x123"),
			TokenAddr:     common.HexToAddress("0x234"),
			Info: gateway.TokenInfo{
				Erc:      0,
				Id:       big.NewInt(0),
				Quantity: big.NewInt(1000),
			},
		},
	}
	data, err := a.Pack("bulkRequestDepositFor", req)
	require.NoError(t, err)
	println(common.Bytes2Hex(data))

	var unpacked interface{}
	err = u.UnpackToInterface(*a, "bulkRequestDepositFor", data[4:], true, &unpacked)
	require.NoError(t, err)
	println(reflect.TypeOf(unpacked).Elem().String())
	out := make([]interface{}, 0)
	rv := reflect.ValueOf(unpacked)
	if rv.Kind() == reflect.Slice {
		for i := 0; i < rv.Len(); i++ {
			println(rv.Index(i).Type().String())
			out = append(out, rv.Index(i).Interface())
		}
	}
	for i := 0; i < reflect.TypeOf(out[0]).NumField(); i++ {
		println(fmt.Sprintf("%s:%v", reflect.TypeOf(out[0]).Field(i).Name, reflect.ValueOf(out[0]).Field(i).Interface()))
	}
}
