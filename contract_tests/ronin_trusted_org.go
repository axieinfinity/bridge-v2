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

// RoninTrustedOrganizationTrustedOrganization is an auto generated low-level Go binding around an user-defined struct.
type RoninTrustedOrganizationTrustedOrganization struct {
	BridgeVoter common.Address
}

// RoninTrustedOrganizationMetaData contains all meta data concerning the RoninTrustedOrganization contract.
var RoninTrustedOrganizationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getAllTrustedOrganizations\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"bridgeVoter\",\"type\":\"address\"}],\"internalType\":\"structRoninTrustedOrganization.TrustedOrganization[]\",\"name\":\"_trustedOrganizations\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"trusted\",\"type\":\"address[]\"}],\"name\":\"setTrustedOrganizations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"15074005": "getAllTrustedOrganizations()",
		"856bd9d8": "setTrustedOrganizations(address[])",
	},
	Bin: "0x608060405234801561001057600080fd5b50610331806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063150740051461003b578063856bd9d814610059575b600080fd5b61004361006e565b6040516100509190610179565b60405180910390f35b61006c6100673660046101f9565b6100d5565b005b60606000805480602002602001604051908101604052809291908181526020016000905b828210156100cc57600084815260209081902060408051808401909152908401546001600160a01b03168152825260019092019101610092565b50505050905090565b60005b81518110156101755760006040518060200160405280848481518110610100576101006102be565b60209081029190910101516001600160a01b039081169091526000805460018101825590805291517f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56390920180546001600160a01b03191692909116919091179055508061016d816102d4565b9150506100d8565b5050565b6020808252825182820181905260009190848201906040850190845b818110156101bb578351516001600160a01b031683529284019291840191600101610195565b50909695505050505050565b634e487b7160e01b600052604160045260246000fd5b80356001600160a01b03811681146101f457600080fd5b919050565b6000602080838503121561020c57600080fd5b823567ffffffffffffffff8082111561022457600080fd5b818501915085601f83011261023857600080fd5b81358181111561024a5761024a6101c7565b8060051b604051601f19603f8301168101818110858211171561026f5761026f6101c7565b60405291825284820192508381018501918883111561028d57600080fd5b938501935b828510156102b2576102a3856101dd565b84529385019392850192610292565b98975050505050505050565b634e487b7160e01b600052603260045260246000fd5b6000600182016102f457634e487b7160e01b600052601160045260246000fd5b506001019056fea26469706673582212202134d86653305a16183265c9cda1f1cac7d87633db35b40ba63b9cc7249d390b64736f6c63430008110033",
}

// RoninTrustedOrganizationABI is the input ABI used to generate the binding from.
// Deprecated: Use RoninTrustedOrganizationMetaData.ABI instead.
var RoninTrustedOrganizationABI = RoninTrustedOrganizationMetaData.ABI

// Deprecated: Use RoninTrustedOrganizationMetaData.Sigs instead.
// RoninTrustedOrganizationFuncSigs maps the 4-byte function signature to its string representation.
var RoninTrustedOrganizationFuncSigs = RoninTrustedOrganizationMetaData.Sigs

// RoninTrustedOrganizationBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RoninTrustedOrganizationMetaData.Bin instead.
var RoninTrustedOrganizationBin = RoninTrustedOrganizationMetaData.Bin

// DeployRoninTrustedOrganization deploys a new Ethereum contract, binding an instance of RoninTrustedOrganization to it.
func DeployRoninTrustedOrganization(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RoninTrustedOrganization, error) {
	parsed, err := RoninTrustedOrganizationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RoninTrustedOrganizationBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RoninTrustedOrganization{RoninTrustedOrganizationCaller: RoninTrustedOrganizationCaller{contract: contract}, RoninTrustedOrganizationTransactor: RoninTrustedOrganizationTransactor{contract: contract}, RoninTrustedOrganizationFilterer: RoninTrustedOrganizationFilterer{contract: contract}}, nil
}

// RoninTrustedOrganization is an auto generated Go binding around an Ethereum contract.
type RoninTrustedOrganization struct {
	RoninTrustedOrganizationCaller     // Read-only binding to the contract
	RoninTrustedOrganizationTransactor // Write-only binding to the contract
	RoninTrustedOrganizationFilterer   // Log filterer for contract events
}

// RoninTrustedOrganizationCaller is an auto generated read-only Go binding around an Ethereum contract.
type RoninTrustedOrganizationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoninTrustedOrganizationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RoninTrustedOrganizationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoninTrustedOrganizationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoninTrustedOrganizationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoninTrustedOrganizationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoninTrustedOrganizationSession struct {
	Contract     *RoninTrustedOrganization // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// RoninTrustedOrganizationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoninTrustedOrganizationCallerSession struct {
	Contract *RoninTrustedOrganizationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// RoninTrustedOrganizationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoninTrustedOrganizationTransactorSession struct {
	Contract     *RoninTrustedOrganizationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// RoninTrustedOrganizationRaw is an auto generated low-level Go binding around an Ethereum contract.
type RoninTrustedOrganizationRaw struct {
	Contract *RoninTrustedOrganization // Generic contract binding to access the raw methods on
}

// RoninTrustedOrganizationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoninTrustedOrganizationCallerRaw struct {
	Contract *RoninTrustedOrganizationCaller // Generic read-only contract binding to access the raw methods on
}

// RoninTrustedOrganizationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoninTrustedOrganizationTransactorRaw struct {
	Contract *RoninTrustedOrganizationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoninTrustedOrganization creates a new instance of RoninTrustedOrganization, bound to a specific deployed contract.
func NewRoninTrustedOrganization(address common.Address, backend bind.ContractBackend) (*RoninTrustedOrganization, error) {
	contract, err := bindRoninTrustedOrganization(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RoninTrustedOrganization{RoninTrustedOrganizationCaller: RoninTrustedOrganizationCaller{contract: contract}, RoninTrustedOrganizationTransactor: RoninTrustedOrganizationTransactor{contract: contract}, RoninTrustedOrganizationFilterer: RoninTrustedOrganizationFilterer{contract: contract}}, nil
}

// NewRoninTrustedOrganizationCaller creates a new read-only instance of RoninTrustedOrganization, bound to a specific deployed contract.
func NewRoninTrustedOrganizationCaller(address common.Address, caller bind.ContractCaller) (*RoninTrustedOrganizationCaller, error) {
	contract, err := bindRoninTrustedOrganization(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoninTrustedOrganizationCaller{contract: contract}, nil
}

// NewRoninTrustedOrganizationTransactor creates a new write-only instance of RoninTrustedOrganization, bound to a specific deployed contract.
func NewRoninTrustedOrganizationTransactor(address common.Address, transactor bind.ContractTransactor) (*RoninTrustedOrganizationTransactor, error) {
	contract, err := bindRoninTrustedOrganization(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoninTrustedOrganizationTransactor{contract: contract}, nil
}

// NewRoninTrustedOrganizationFilterer creates a new log filterer instance of RoninTrustedOrganization, bound to a specific deployed contract.
func NewRoninTrustedOrganizationFilterer(address common.Address, filterer bind.ContractFilterer) (*RoninTrustedOrganizationFilterer, error) {
	contract, err := bindRoninTrustedOrganization(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoninTrustedOrganizationFilterer{contract: contract}, nil
}

// bindRoninTrustedOrganization binds a generic wrapper to an already deployed contract.
func bindRoninTrustedOrganization(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RoninTrustedOrganizationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoninTrustedOrganization *RoninTrustedOrganizationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoninTrustedOrganization.Contract.RoninTrustedOrganizationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoninTrustedOrganization *RoninTrustedOrganizationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoninTrustedOrganization.Contract.RoninTrustedOrganizationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoninTrustedOrganization *RoninTrustedOrganizationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoninTrustedOrganization.Contract.RoninTrustedOrganizationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoninTrustedOrganization *RoninTrustedOrganizationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoninTrustedOrganization.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoninTrustedOrganization *RoninTrustedOrganizationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoninTrustedOrganization.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoninTrustedOrganization *RoninTrustedOrganizationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoninTrustedOrganization.Contract.contract.Transact(opts, method, params...)
}

// GetAllTrustedOrganizations is a free data retrieval call binding the contract method 0x15074005.
//
// Solidity: function getAllTrustedOrganizations() view returns((address)[] _trustedOrganizations)
func (_RoninTrustedOrganization *RoninTrustedOrganizationCaller) GetAllTrustedOrganizations(opts *bind.CallOpts) ([]RoninTrustedOrganizationTrustedOrganization, error) {
	var out []interface{}
	err := _RoninTrustedOrganization.contract.Call(opts, &out, "getAllTrustedOrganizations")

	if err != nil {
		return *new([]RoninTrustedOrganizationTrustedOrganization), err
	}

	out0 := *abi.ConvertType(out[0], new([]RoninTrustedOrganizationTrustedOrganization)).(*[]RoninTrustedOrganizationTrustedOrganization)

	return out0, err

}

// GetAllTrustedOrganizations is a free data retrieval call binding the contract method 0x15074005.
//
// Solidity: function getAllTrustedOrganizations() view returns((address)[] _trustedOrganizations)
func (_RoninTrustedOrganization *RoninTrustedOrganizationSession) GetAllTrustedOrganizations() ([]RoninTrustedOrganizationTrustedOrganization, error) {
	return _RoninTrustedOrganization.Contract.GetAllTrustedOrganizations(&_RoninTrustedOrganization.CallOpts)
}

// GetAllTrustedOrganizations is a free data retrieval call binding the contract method 0x15074005.
//
// Solidity: function getAllTrustedOrganizations() view returns((address)[] _trustedOrganizations)
func (_RoninTrustedOrganization *RoninTrustedOrganizationCallerSession) GetAllTrustedOrganizations() ([]RoninTrustedOrganizationTrustedOrganization, error) {
	return _RoninTrustedOrganization.Contract.GetAllTrustedOrganizations(&_RoninTrustedOrganization.CallOpts)
}

// SetTrustedOrganizations is a paid mutator transaction binding the contract method 0x856bd9d8.
//
// Solidity: function setTrustedOrganizations(address[] trusted) returns()
func (_RoninTrustedOrganization *RoninTrustedOrganizationTransactor) SetTrustedOrganizations(opts *bind.TransactOpts, trusted []common.Address) (*types.Transaction, error) {
	return _RoninTrustedOrganization.contract.Transact(opts, "setTrustedOrganizations", trusted)
}

// SetTrustedOrganizations is a paid mutator transaction binding the contract method 0x856bd9d8.
//
// Solidity: function setTrustedOrganizations(address[] trusted) returns()
func (_RoninTrustedOrganization *RoninTrustedOrganizationSession) SetTrustedOrganizations(trusted []common.Address) (*types.Transaction, error) {
	return _RoninTrustedOrganization.Contract.SetTrustedOrganizations(&_RoninTrustedOrganization.TransactOpts, trusted)
}

// SetTrustedOrganizations is a paid mutator transaction binding the contract method 0x856bd9d8.
//
// Solidity: function setTrustedOrganizations(address[] trusted) returns()
func (_RoninTrustedOrganization *RoninTrustedOrganizationTransactorSession) SetTrustedOrganizations(trusted []common.Address) (*types.Transaction, error) {
	return _RoninTrustedOrganization.Contract.SetTrustedOrganizations(&_RoninTrustedOrganization.TransactOpts, trusted)
}
