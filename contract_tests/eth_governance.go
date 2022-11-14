// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract_tests

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EthGovernanceSignature is an auto generated low-level Go binding around an user-defined struct.
type EthGovernanceSignature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// EthGovernanceMetaData contains all meta data concerning the EthGovernance contract.
var EthGovernanceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_voters\",\"type\":\"address[]\"}],\"name\":\"getBridgeOperatorVotingSignatures\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structEthGovernance.Signature[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_operators\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structEthGovernance.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"name\":\"relayBridgeOperators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"332635be": "getBridgeOperatorVotingSignatures(uint256,address[])",
		"d5918e61": "relayBridgeOperators(uint256,address[],(uint8,bytes32,bytes32)[])",
	},
	Bin: "0x608060405234801561001057600080fd5b50610573806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063332635be1461003b578063d5918e6114610064575b600080fd5b61004e6100493660046102a4565b610079565b60405161005b91906102f0565b60405180910390f35b6100776100723660046103bc565b61019f565b005b606060008267ffffffffffffffff8111156100965761009661034c565b6040519080825280602002602001820160405280156100e157816020015b60408051606081018252600080825260208083018290529282015282526000199092019101816100b45790505b50905060005b8381101561019657600086815260208190526040812090868684818110610110576101106104d0565b905060200201602081019061012591906104e6565b6001600160a01b0316815260208082019290925260409081016000208151606081018352815460ff16815260018201549381019390935260020154908201528251839083908110610178576101786104d0565b6020026020010181905250808061018e90610516565b9150506100e7565b50949350505050565b60005b82811015610251578181815181106101bc576101bc6104d0565b602002602001015160008087815260200190815260200160002060008686858181106101ea576101ea6104d0565b90506020020160208101906101ff91906104e6565b6001600160a01b0316815260208082019290925260409081016000208351815460ff191660ff90911617815591830151600183015591909101516002909101558061024981610516565b9150506101a2565b5050505050565b60008083601f84011261026a57600080fd5b50813567ffffffffffffffff81111561028257600080fd5b6020830191508360208260051b850101111561029d57600080fd5b9250929050565b6000806000604084860312156102b957600080fd5b83359250602084013567ffffffffffffffff8111156102d757600080fd5b6102e386828701610258565b9497909650939450505050565b602080825282518282018190526000919060409081850190868401855b8281101561033f578151805160ff1685528681015187860152850151858501526060909301929085019060010161030d565b5091979650505050505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff811182821017156103855761038561034c565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156103b4576103b461034c565b604052919050565b60008060008060608086880312156103d357600080fd5b8535945060208087013567ffffffffffffffff808211156103f357600080fd5b6103ff8a838b01610258565b909750955060409150888201358181111561041957600080fd5b8901601f81018b1361042a57600080fd5b80358281111561043c5761043c61034c565b61044a858260051b0161038b565b818152858101935090860282018501908c82111561046757600080fd5b918501915b818310156104bf5786838e0312156104845760008081fd5b61048c610362565b833560ff8116811461049e5760008081fd5b8152838701358782015285840135868201528452928501929186019161046c565b999c989b5096995050505050505050565b634e487b7160e01b600052603260045260246000fd5b6000602082840312156104f857600080fd5b81356001600160a01b038116811461050f57600080fd5b9392505050565b60006001820161053657634e487b7160e01b600052601160045260246000fd5b506001019056fea2646970667358221220b55919f75ff0f17fb796cba140c7187904cbb503d0daedac9005abddf05d26ec64736f6c63430008110033",
}

// EthGovernanceABI is the input ABI used to generate the binding from.
// Deprecated: Use EthGovernanceMetaData.ABI instead.
var EthGovernanceABI = EthGovernanceMetaData.ABI

// Deprecated: Use EthGovernanceMetaData.Sigs instead.
// EthGovernanceFuncSigs maps the 4-byte function signature to its string representation.
var EthGovernanceFuncSigs = EthGovernanceMetaData.Sigs

// EthGovernanceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthGovernanceMetaData.Bin instead.
var EthGovernanceBin = EthGovernanceMetaData.Bin

// DeployEthGovernance deploys a new Ethereum contract, binding an instance of EthGovernance to it.
func DeployEthGovernance(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EthGovernance, error) {
	parsed, err := EthGovernanceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthGovernanceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthGovernance{EthGovernanceCaller: EthGovernanceCaller{contract: contract}, EthGovernanceTransactor: EthGovernanceTransactor{contract: contract}, EthGovernanceFilterer: EthGovernanceFilterer{contract: contract}}, nil
}

// EthGovernance is an auto generated Go binding around an Ethereum contract.
type EthGovernance struct {
	EthGovernanceCaller     // Read-only binding to the contract
	EthGovernanceTransactor // Write-only binding to the contract
	EthGovernanceFilterer   // Log filterer for contract events
}

// EthGovernanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthGovernanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthGovernanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthGovernanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthGovernanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthGovernanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthGovernanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthGovernanceSession struct {
	Contract     *EthGovernance    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthGovernanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthGovernanceCallerSession struct {
	Contract *EthGovernanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EthGovernanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthGovernanceTransactorSession struct {
	Contract     *EthGovernanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EthGovernanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthGovernanceRaw struct {
	Contract *EthGovernance // Generic contract binding to access the raw methods on
}

// EthGovernanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthGovernanceCallerRaw struct {
	Contract *EthGovernanceCaller // Generic read-only contract binding to access the raw methods on
}

// EthGovernanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthGovernanceTransactorRaw struct {
	Contract *EthGovernanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthGovernance creates a new instance of EthGovernance, bound to a specific deployed contract.
func NewEthGovernance(address common.Address, backend bind.ContractBackend) (*EthGovernance, error) {
	contract, err := bindEthGovernance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthGovernance{EthGovernanceCaller: EthGovernanceCaller{contract: contract}, EthGovernanceTransactor: EthGovernanceTransactor{contract: contract}, EthGovernanceFilterer: EthGovernanceFilterer{contract: contract}}, nil
}

// NewEthGovernanceCaller creates a new read-only instance of EthGovernance, bound to a specific deployed contract.
func NewEthGovernanceCaller(address common.Address, caller bind.ContractCaller) (*EthGovernanceCaller, error) {
	contract, err := bindEthGovernance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthGovernanceCaller{contract: contract}, nil
}

// NewEthGovernanceTransactor creates a new write-only instance of EthGovernance, bound to a specific deployed contract.
func NewEthGovernanceTransactor(address common.Address, transactor bind.ContractTransactor) (*EthGovernanceTransactor, error) {
	contract, err := bindEthGovernance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthGovernanceTransactor{contract: contract}, nil
}

// NewEthGovernanceFilterer creates a new log filterer instance of EthGovernance, bound to a specific deployed contract.
func NewEthGovernanceFilterer(address common.Address, filterer bind.ContractFilterer) (*EthGovernanceFilterer, error) {
	contract, err := bindEthGovernance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthGovernanceFilterer{contract: contract}, nil
}

// bindEthGovernance binds a generic wrapper to an already deployed contract.
func bindEthGovernance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthGovernanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthGovernance *EthGovernanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthGovernance.Contract.EthGovernanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthGovernance *EthGovernanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthGovernance.Contract.EthGovernanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthGovernance *EthGovernanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthGovernance.Contract.EthGovernanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthGovernance *EthGovernanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthGovernance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthGovernance *EthGovernanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthGovernance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthGovernance *EthGovernanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthGovernance.Contract.contract.Transact(opts, method, params...)
}

// GetBridgeOperatorVotingSignatures is a free data retrieval call binding the contract method 0x332635be.
//
// Solidity: function getBridgeOperatorVotingSignatures(uint256 _period, address[] _voters) view returns((uint8,bytes32,bytes32)[])
func (_EthGovernance *EthGovernanceCaller) GetBridgeOperatorVotingSignatures(opts *bind.CallOpts, _period *big.Int, _voters []common.Address) ([]EthGovernanceSignature, error) {
	var out []interface{}
	err := _EthGovernance.contract.Call(opts, &out, "getBridgeOperatorVotingSignatures", _period, _voters)

	if err != nil {
		return *new([]EthGovernanceSignature), err
	}

	out0 := *abi.ConvertType(out[0], new([]EthGovernanceSignature)).(*[]EthGovernanceSignature)

	return out0, err

}

// GetBridgeOperatorVotingSignatures is a free data retrieval call binding the contract method 0x332635be.
//
// Solidity: function getBridgeOperatorVotingSignatures(uint256 _period, address[] _voters) view returns((uint8,bytes32,bytes32)[])
func (_EthGovernance *EthGovernanceSession) GetBridgeOperatorVotingSignatures(_period *big.Int, _voters []common.Address) ([]EthGovernanceSignature, error) {
	return _EthGovernance.Contract.GetBridgeOperatorVotingSignatures(&_EthGovernance.CallOpts, _period, _voters)
}

// GetBridgeOperatorVotingSignatures is a free data retrieval call binding the contract method 0x332635be.
//
// Solidity: function getBridgeOperatorVotingSignatures(uint256 _period, address[] _voters) view returns((uint8,bytes32,bytes32)[])
func (_EthGovernance *EthGovernanceCallerSession) GetBridgeOperatorVotingSignatures(_period *big.Int, _voters []common.Address) ([]EthGovernanceSignature, error) {
	return _EthGovernance.Contract.GetBridgeOperatorVotingSignatures(&_EthGovernance.CallOpts, _period, _voters)
}

// RelayBridgeOperators is a paid mutator transaction binding the contract method 0xd5918e61.
//
// Solidity: function relayBridgeOperators(uint256 _period, address[] _operators, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_EthGovernance *EthGovernanceTransactor) RelayBridgeOperators(opts *bind.TransactOpts, _period *big.Int, _operators []common.Address, _signatures []EthGovernanceSignature) (*types.Transaction, error) {
	return _EthGovernance.contract.Transact(opts, "relayBridgeOperators", _period, _operators, _signatures)
}

// RelayBridgeOperators is a paid mutator transaction binding the contract method 0xd5918e61.
//
// Solidity: function relayBridgeOperators(uint256 _period, address[] _operators, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_EthGovernance *EthGovernanceSession) RelayBridgeOperators(_period *big.Int, _operators []common.Address, _signatures []EthGovernanceSignature) (*types.Transaction, error) {
	return _EthGovernance.Contract.RelayBridgeOperators(&_EthGovernance.TransactOpts, _period, _operators, _signatures)
}

// RelayBridgeOperators is a paid mutator transaction binding the contract method 0xd5918e61.
//
// Solidity: function relayBridgeOperators(uint256 _period, address[] _operators, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_EthGovernance *EthGovernanceTransactorSession) RelayBridgeOperators(_period *big.Int, _operators []common.Address, _signatures []EthGovernanceSignature) (*types.Transaction, error) {
	return _EthGovernance.Contract.RelayBridgeOperators(&_EthGovernance.TransactOpts, _period, _operators, _signatures)
}
