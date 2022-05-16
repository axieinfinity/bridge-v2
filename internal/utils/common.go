package utils

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"math/big"
	"os"
	"reflect"
)

type EthClient interface {
	ethereum.ChainReader
	ethereum.TransactionReader
	ethereum.ChainStateReader
	ethereum.ContractCaller
	ethereum.LogFilterer

	BlockNumber(ctx context.Context) (uint64, error)
	ChainID(ctx context.Context) (*big.Int, error)
	Close()
}

type IUtils interface {
	Invoke(any interface{}, name string, args ...interface{}) (reflect.Value, error)
	LoadAbi(path string) (*abi.ABI, error)
	GetArguments(a abi.ABI, name string, data []byte, isInput bool) (abi.Arguments, error)
	UnpackToInterface(a abi.ABI, name string, data []byte, isInput bool, v interface{}) error
	Title(text string) string
	NewEthClient(url string) (EthClient, error)
	SendContractTransaction(key *ecdsa.PrivateKey, chainId *big.Int, fn func(opts *bind.TransactOpts) (*types.Transaction, error)) (*types.Transaction, error)
	SignTypedData(typedData apitypes.TypedData, privateKey *ecdsa.PrivateKey) (hexutil.Bytes, error)
	FilterLogs(client EthClient, opts *bind.FilterOpts, contractAddresses []common.Address, filteredMethods map[*abi.ABI]map[string]struct{}) ([]types.Log, error)
}

type Utils struct{}

func (u *Utils) Invoke(any interface{}, name string, args ...interface{}) (reflect.Value, error) {
	method := reflect.ValueOf(any).MethodByName(name)
	methodType := method.Type()
	numIn := methodType.NumIn()
	if numIn > len(args) {
		return reflect.ValueOf(nil), fmt.Errorf("method %s must have minimum %d params. Have %d", name, numIn, len(args))
	}
	if numIn != len(args) && !methodType.IsVariadic() {
		return reflect.ValueOf(nil), fmt.Errorf("method %s must have %d params. Have %d", name, numIn, len(args))
	}
	in := make([]reflect.Value, len(args))
	for i := 0; i < len(args); i++ {
		var inType reflect.Type
		if methodType.IsVariadic() && i >= numIn-1 {
			inType = methodType.In(numIn - 1).Elem()
		} else {
			inType = methodType.In(i)
		}
		argValue := reflect.ValueOf(args[i])
		if !argValue.IsValid() {
			return reflect.ValueOf(nil), fmt.Errorf("Method %s. Param[%d] must be %s. Have %s", name, i, inType, argValue.String())
		}
		argType := argValue.Type()
		if argType.ConvertibleTo(inType) {
			in[i] = argValue.Convert(inType)
		} else {
			return reflect.ValueOf(nil), fmt.Errorf("Method %s. Param[%d] must be %s. Have %s", name, i, inType, argType)
		}
	}
	return method.Call(in)[0], nil
}

func (u *Utils) LoadAbi(path string) (*abi.ABI, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	a, err := abi.JSON(f)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (u *Utils) GetArguments(a abi.ABI, name string, data []byte, isInput bool) (abi.Arguments, error) {
	// since there can't be naming collisions with contracts and events,
	// we need to decide whether we're calling a method or an event
	var args abi.Arguments
	if method, ok := a.Methods[name]; ok {
		if len(data)%32 != 0 {
			return nil, fmt.Errorf("abi: improperly formatted output: %s - Bytes: [%+v]", string(data), data)
		}
		if isInput {
			args = method.Inputs
		} else {
			args = method.Outputs
		}
	}
	if event, ok := a.Events[name]; ok {
		args = event.Inputs
	}
	if args == nil {
		return nil, errors.New("abi: could not locate named method or event")
	}
	return args, nil
}

func (u *Utils) UnpackToInterface(a abi.ABI, name string, data []byte, isInput bool, v interface{}) error {
	args, err := u.GetArguments(a, name, data, isInput)
	if err != nil {
		return err
	}
	unpacked, err := args.Unpack(data)
	if err != nil {
		return err
	}
	return args.Copy(v, unpacked)
}

func (u *Utils) Title(text string) string {
	c := cases.Title(language.English)
	return c.String(text)
}

func (u *Utils) NewEthClient(url string) (EthClient, error) {
	return ethclient.Dial(url)
}

func (u *Utils) SendContractTransaction(key *ecdsa.PrivateKey, chainId *big.Int, fn func(opts *bind.TransactOpts) (*types.Transaction, error)) (*types.Transaction, error) {
	if key == nil {
		return nil, nil
	}
	opts, err := bind.NewKeyedTransactorWithChainID(key, chainId)
	if err != nil {
		return nil, err
	}
	return fn(opts)
}

// SignTypedData signs EIP-712 conformant typed data
// hash = keccak256("\x19${byteVersion}${domainSeparator}${hashStruct(message)}")
// It returns
// - the signature,
// - and/or any error
func (u *Utils) SignTypedData(typedData apitypes.TypedData, privateKey *ecdsa.PrivateKey) (hexutil.Bytes, error) {
	return u.signTypedData(typedData, privateKey)
}

// signTypedData is identical to the capitalized version
func (u *Utils) signTypedData(typedData apitypes.TypedData, privateKey *ecdsa.PrivateKey) (hexutil.Bytes, error) {
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return nil, err
	}
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return nil, err
	}
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	signature, err := crypto.Sign(crypto.Keccak256(rawData), privateKey)
	if err != nil {
		return nil, err
	}
	signature[64] += 27 // Transform V from 0/1 to 27/28 according to the yellow paper
	return signature, nil
}

func (u *Utils) FilterLogs(client EthClient, opts *bind.FilterOpts, contractAddresses []common.Address, filteredMethods map[*abi.ABI]map[string]struct{}) ([]types.Log, error) {
	// Don't crash on a lazy user
	if opts == nil {
		opts = new(bind.FilterOpts)
	}
	var query [][]interface{}
	for contractAbi, methods := range filteredMethods {
		var events []interface{}
		for method, _ := range methods {
			if _, ok := contractAbi.Events[method]; ok {
				events = append(events, contractAbi.Events[method].ID)
			}
		}
		query = append(query, events)
	}
	topics, err := abi.MakeTopics(query...)
	if err != nil {
		return nil, err
	}
	config := ethereum.FilterQuery{
		Addresses: contractAddresses,
		Topics:    topics,
		FromBlock: new(big.Int).SetUint64(opts.Start),
	}
	if opts.End != nil {
		config.ToBlock = new(big.Int).SetUint64(*opts.End)
	}
	return client.FilterLogs(opts.Context, config)
}
