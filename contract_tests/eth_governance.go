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
	Bin: "0x608060405234801561001057600080fd5b5061050b806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063332635be1461003b578063d5918e6114610064575b600080fd5b61004e61004936600461023c565b610079565b60405161005b9190610288565b60405180910390f35b610077610072366004610354565b610137565b005b60608060005b8381101561012e576000868152602081905260408120908686848181106100a8576100a8610468565b90506020020160208101906100bd919061047e565b6001600160a01b0316815260208082019290925260409081016000208151606081018352815460ff1681526001820154938101939093526002015490820152825183908390811061011057610110610468565b60200260200101819052508080610126906104ae565b91505061007f565b50949350505050565b60005b828110156101e95781818151811061015457610154610468565b6020026020010151600080878152602001908152602001600020600086868581811061018257610182610468565b9050602002016020810190610197919061047e565b6001600160a01b0316815260208082019290925260409081016000208351815460ff191660ff9091161781559183015160018301559190910151600290910155806101e1816104ae565b91505061013a565b5050505050565b60008083601f84011261020257600080fd5b50813567ffffffffffffffff81111561021a57600080fd5b6020830191508360208260051b850101111561023557600080fd5b9250929050565b60008060006040848603121561025157600080fd5b83359250602084013567ffffffffffffffff81111561026f57600080fd5b61027b868287016101f0565b9497909650939450505050565b602080825282518282018190526000919060409081850190868401855b828110156102d7578151805160ff168552868101518786015285015185850152606090930192908501906001016102a5565b5091979650505050505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561031d5761031d6102e4565b60405290565b604051601f8201601f1916810167ffffffffffffffff8111828210171561034c5761034c6102e4565b604052919050565b600080600080606080868803121561036b57600080fd5b8535945060208087013567ffffffffffffffff8082111561038b57600080fd5b6103978a838b016101f0565b90975095506040915088820135818111156103b157600080fd5b8901601f81018b136103c257600080fd5b8035828111156103d4576103d46102e4565b6103e2858260051b01610323565b818152858101935090860282018501908c8211156103ff57600080fd5b918501915b818310156104575786838e03121561041c5760008081fd5b6104246102fa565b833560ff811681146104365760008081fd5b81528387013587820152858401358682015284529285019291860191610404565b999c989b5096995050505050505050565b634e487b7160e01b600052603260045260246000fd5b60006020828403121561049057600080fd5b81356001600160a01b03811681146104a757600080fd5b9392505050565b6000600182016104ce57634e487b7160e01b600052601160045260246000fd5b506001019056fea264697066735822122050842fb0ccd1e91f4a0d24384ce26ea8e9022b309fa1c719b1df4becf08ff1c464736f6c63430008110033",
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
