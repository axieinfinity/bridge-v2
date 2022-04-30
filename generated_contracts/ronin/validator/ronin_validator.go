// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package validator

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

// IWeightedValidatorWeightedValidator is an auto generated low-level Go binding around an user-defined struct.
type IWeightedValidatorWeightedValidator struct {
	Addr   common.Address
	Weight *big.Int
}

// ValidatorMetaData contains all meta data concerning the Validator contract.
var ValidatorMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_numerator\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_denominator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_previousNumerator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_previousDenominator\",\"type\":\"uint256\"}],\"name\":\"ThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIWeightedValidator.WeightedValidator[]\",\"name\":\"_validators\",\"type\":\"tuple[]\"}],\"name\":\"ValidatorsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"}],\"name\":\"ValidatorsRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIWeightedValidator.WeightedValidator[]\",\"name\":\"_validators\",\"type\":\"tuple[]\"}],\"name\":\"ValidatorsUpdated\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"internalType\":\"structIWeightedValidator.WeightedValidator[]\",\"name\":\"_validatorList\",\"type\":\"tuple[]\"}],\"name\":\"addValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_voteWeight\",\"type\":\"uint256\"}],\"name\":\"checkThreshold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"internalType\":\"structIWeightedValidator.WeightedValidator[]\",\"name\":\"_list\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addrList\",\"type\":\"address[]\"}],\"name\":\"getWeights\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"internalType\":\"structIWeightedValidator.WeightedValidator[]\",\"name\":\"_initValidators\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_denominator\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumVoteWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validatorList\",\"type\":\"address[]\"}],\"name\":\"removeValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_denominator\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_previousNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_previousDenom\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalWeights\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"internalType\":\"structIWeightedValidator.WeightedValidator[]\",\"name\":\"_validatorList\",\"type\":\"tuple[]\"}],\"name\":\"updateValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"validators\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"internalType\":\"structIWeightedValidator.WeightedValidator\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ValidatorABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorMetaData.ABI instead.
var ValidatorABI = ValidatorMetaData.ABI

// Validator is an auto generated Go binding around an Ethereum contract.
type Validator struct {
	ValidatorCaller     // Read-only binding to the contract
	ValidatorTransactor // Write-only binding to the contract
	ValidatorFilterer   // Log filterer for contract events
}

// ValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorSession struct {
	Contract     *Validator        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorCallerSession struct {
	Contract *ValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorTransactorSession struct {
	Contract     *ValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorRaw struct {
	Contract *Validator // Generic contract binding to access the raw methods on
}

// ValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorCallerRaw struct {
	Contract *ValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorTransactorRaw struct {
	Contract *ValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidator creates a new instance of Validator, bound to a specific deployed contract.
func NewValidator(address common.Address, backend bind.ContractBackend) (*Validator, error) {
	contract, err := bindValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validator{ValidatorCaller: ValidatorCaller{contract: contract}, ValidatorTransactor: ValidatorTransactor{contract: contract}, ValidatorFilterer: ValidatorFilterer{contract: contract}}, nil
}

// NewValidatorCaller creates a new read-only instance of Validator, bound to a specific deployed contract.
func NewValidatorCaller(address common.Address, caller bind.ContractCaller) (*ValidatorCaller, error) {
	contract, err := bindValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorCaller{contract: contract}, nil
}

// NewValidatorTransactor creates a new write-only instance of Validator, bound to a specific deployed contract.
func NewValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorTransactor, error) {
	contract, err := bindValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorTransactor{contract: contract}, nil
}

// NewValidatorFilterer creates a new log filterer instance of Validator, bound to a specific deployed contract.
func NewValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorFilterer, error) {
	contract, err := bindValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorFilterer{contract: contract}, nil
}

// bindValidator binds a generic wrapper to an already deployed contract.
func bindValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validator *ValidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validator.Contract.ValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validator *ValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.Contract.ValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validator *ValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validator.Contract.ValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validator *ValidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validator *ValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validator *ValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validator.Contract.contract.Transact(opts, method, params...)
}

// CheckThreshold is a free data retrieval call binding the contract method 0xdafae408.
//
// Solidity: function checkThreshold(uint256 _voteWeight) view returns(bool)
func (_Validator *ValidatorCaller) CheckThreshold(opts *bind.CallOpts, _voteWeight *big.Int) (bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "checkThreshold", _voteWeight)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckThreshold is a free data retrieval call binding the contract method 0xdafae408.
//
// Solidity: function checkThreshold(uint256 _voteWeight) view returns(bool)
func (_Validator *ValidatorSession) CheckThreshold(_voteWeight *big.Int) (bool, error) {
	return _Validator.Contract.CheckThreshold(&_Validator.CallOpts, _voteWeight)
}

// CheckThreshold is a free data retrieval call binding the contract method 0xdafae408.
//
// Solidity: function checkThreshold(uint256 _voteWeight) view returns(bool)
func (_Validator *ValidatorCallerSession) CheckThreshold(_voteWeight *big.Int) (bool, error) {
	return _Validator.Contract.CheckThreshold(&_Validator.CallOpts, _voteWeight)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256, uint256)
func (_Validator *ValidatorCaller) GetThreshold(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "getThreshold")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256, uint256)
func (_Validator *ValidatorSession) GetThreshold() (*big.Int, *big.Int, error) {
	return _Validator.Contract.GetThreshold(&_Validator.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256, uint256)
func (_Validator *ValidatorCallerSession) GetThreshold() (*big.Int, *big.Int, error) {
	return _Validator.Contract.GetThreshold(&_Validator.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns((address,uint256)[] _list)
func (_Validator *ValidatorCaller) GetValidators(opts *bind.CallOpts) ([]IWeightedValidatorWeightedValidator, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "getValidators")

	if err != nil {
		return *new([]IWeightedValidatorWeightedValidator), err
	}

	out0 := *abi.ConvertType(out[0], new([]IWeightedValidatorWeightedValidator)).(*[]IWeightedValidatorWeightedValidator)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns((address,uint256)[] _list)
func (_Validator *ValidatorSession) GetValidators() ([]IWeightedValidatorWeightedValidator, error) {
	return _Validator.Contract.GetValidators(&_Validator.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns((address,uint256)[] _list)
func (_Validator *ValidatorCallerSession) GetValidators() ([]IWeightedValidatorWeightedValidator, error) {
	return _Validator.Contract.GetValidators(&_Validator.CallOpts)
}

// GetWeight is a free data retrieval call binding the contract method 0xac6c5251.
//
// Solidity: function getWeight(address _addr) view returns(uint256)
func (_Validator *ValidatorCaller) GetWeight(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "getWeight", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWeight is a free data retrieval call binding the contract method 0xac6c5251.
//
// Solidity: function getWeight(address _addr) view returns(uint256)
func (_Validator *ValidatorSession) GetWeight(_addr common.Address) (*big.Int, error) {
	return _Validator.Contract.GetWeight(&_Validator.CallOpts, _addr)
}

// GetWeight is a free data retrieval call binding the contract method 0xac6c5251.
//
// Solidity: function getWeight(address _addr) view returns(uint256)
func (_Validator *ValidatorCallerSession) GetWeight(_addr common.Address) (*big.Int, error) {
	return _Validator.Contract.GetWeight(&_Validator.CallOpts, _addr)
}

// GetWeights is a free data retrieval call binding the contract method 0x965b9ff1.
//
// Solidity: function getWeights(address[] _addrList) view returns(uint256 _weight)
func (_Validator *ValidatorCaller) GetWeights(opts *bind.CallOpts, _addrList []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "getWeights", _addrList)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWeights is a free data retrieval call binding the contract method 0x965b9ff1.
//
// Solidity: function getWeights(address[] _addrList) view returns(uint256 _weight)
func (_Validator *ValidatorSession) GetWeights(_addrList []common.Address) (*big.Int, error) {
	return _Validator.Contract.GetWeights(&_Validator.CallOpts, _addrList)
}

// GetWeights is a free data retrieval call binding the contract method 0x965b9ff1.
//
// Solidity: function getWeights(address[] _addrList) view returns(uint256 _weight)
func (_Validator *ValidatorCallerSession) GetWeights(_addrList []common.Address) (*big.Int, error) {
	return _Validator.Contract.GetWeights(&_Validator.CallOpts, _addrList)
}

// MinimumVoteWeight is a free data retrieval call binding the contract method 0x7de5dedd.
//
// Solidity: function minimumVoteWeight() view returns(uint256)
func (_Validator *ValidatorCaller) MinimumVoteWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "minimumVoteWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumVoteWeight is a free data retrieval call binding the contract method 0x7de5dedd.
//
// Solidity: function minimumVoteWeight() view returns(uint256)
func (_Validator *ValidatorSession) MinimumVoteWeight() (*big.Int, error) {
	return _Validator.Contract.MinimumVoteWeight(&_Validator.CallOpts)
}

// MinimumVoteWeight is a free data retrieval call binding the contract method 0x7de5dedd.
//
// Solidity: function minimumVoteWeight() view returns(uint256)
func (_Validator *ValidatorCallerSession) MinimumVoteWeight() (*big.Int, error) {
	return _Validator.Contract.MinimumVoteWeight(&_Validator.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Validator *ValidatorCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Validator *ValidatorSession) Nonce() (*big.Int, error) {
	return _Validator.Contract.Nonce(&_Validator.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Validator *ValidatorCallerSession) Nonce() (*big.Int, error) {
	return _Validator.Contract.Nonce(&_Validator.CallOpts)
}

// TotalValidators is a free data retrieval call binding the contract method 0xc81b356b.
//
// Solidity: function totalValidators() view returns(uint256)
func (_Validator *ValidatorCaller) TotalValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "totalValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalValidators is a free data retrieval call binding the contract method 0xc81b356b.
//
// Solidity: function totalValidators() view returns(uint256)
func (_Validator *ValidatorSession) TotalValidators() (*big.Int, error) {
	return _Validator.Contract.TotalValidators(&_Validator.CallOpts)
}

// TotalValidators is a free data retrieval call binding the contract method 0xc81b356b.
//
// Solidity: function totalValidators() view returns(uint256)
func (_Validator *ValidatorCallerSession) TotalValidators() (*big.Int, error) {
	return _Validator.Contract.TotalValidators(&_Validator.CallOpts)
}

// TotalWeights is a free data retrieval call binding the contract method 0x926323d5.
//
// Solidity: function totalWeights() view returns(uint256)
func (_Validator *ValidatorCaller) TotalWeights(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "totalWeights")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalWeights is a free data retrieval call binding the contract method 0x926323d5.
//
// Solidity: function totalWeights() view returns(uint256)
func (_Validator *ValidatorSession) TotalWeights() (*big.Int, error) {
	return _Validator.Contract.TotalWeights(&_Validator.CallOpts)
}

// TotalWeights is a free data retrieval call binding the contract method 0x926323d5.
//
// Solidity: function totalWeights() view returns(uint256)
func (_Validator *ValidatorCallerSession) TotalWeights() (*big.Int, error) {
	return _Validator.Contract.TotalWeights(&_Validator.CallOpts)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 _index) view returns((address,uint256))
func (_Validator *ValidatorCaller) Validators(opts *bind.CallOpts, _index *big.Int) (IWeightedValidatorWeightedValidator, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "validators", _index)

	if err != nil {
		return *new(IWeightedValidatorWeightedValidator), err
	}

	out0 := *abi.ConvertType(out[0], new(IWeightedValidatorWeightedValidator)).(*IWeightedValidatorWeightedValidator)

	return out0, err

}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 _index) view returns((address,uint256))
func (_Validator *ValidatorSession) Validators(_index *big.Int) (IWeightedValidatorWeightedValidator, error) {
	return _Validator.Contract.Validators(&_Validator.CallOpts, _index)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 _index) view returns((address,uint256))
func (_Validator *ValidatorCallerSession) Validators(_index *big.Int) (IWeightedValidatorWeightedValidator, error) {
	return _Validator.Contract.Validators(&_Validator.CallOpts, _index)
}

// AddValidators is a paid mutator transaction binding the contract method 0x23aff661.
//
// Solidity: function addValidators((address,uint256)[] _validatorList) returns()
func (_Validator *ValidatorTransactor) AddValidators(opts *bind.TransactOpts, _validatorList []IWeightedValidatorWeightedValidator) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "addValidators", _validatorList)
}

// AddValidators is a paid mutator transaction binding the contract method 0x23aff661.
//
// Solidity: function addValidators((address,uint256)[] _validatorList) returns()
func (_Validator *ValidatorSession) AddValidators(_validatorList []IWeightedValidatorWeightedValidator) (*types.Transaction, error) {
	return _Validator.Contract.AddValidators(&_Validator.TransactOpts, _validatorList)
}

// AddValidators is a paid mutator transaction binding the contract method 0x23aff661.
//
// Solidity: function addValidators((address,uint256)[] _validatorList) returns()
func (_Validator *ValidatorTransactorSession) AddValidators(_validatorList []IWeightedValidatorWeightedValidator) (*types.Transaction, error) {
	return _Validator.Contract.AddValidators(&_Validator.TransactOpts, _validatorList)
}

// Initialize is a paid mutator transaction binding the contract method 0x92605829.
//
// Solidity: function initialize((address,uint256)[] _initValidators, uint256 _numerator, uint256 _denominator) returns()
func (_Validator *ValidatorTransactor) Initialize(opts *bind.TransactOpts, _initValidators []IWeightedValidatorWeightedValidator, _numerator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "initialize", _initValidators, _numerator, _denominator)
}

// Initialize is a paid mutator transaction binding the contract method 0x92605829.
//
// Solidity: function initialize((address,uint256)[] _initValidators, uint256 _numerator, uint256 _denominator) returns()
func (_Validator *ValidatorSession) Initialize(_initValidators []IWeightedValidatorWeightedValidator, _numerator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.Initialize(&_Validator.TransactOpts, _initValidators, _numerator, _denominator)
}

// Initialize is a paid mutator transaction binding the contract method 0x92605829.
//
// Solidity: function initialize((address,uint256)[] _initValidators, uint256 _numerator, uint256 _denominator) returns()
func (_Validator *ValidatorTransactorSession) Initialize(_initValidators []IWeightedValidatorWeightedValidator, _numerator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.Initialize(&_Validator.TransactOpts, _initValidators, _numerator, _denominator)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0x1d40f0d8.
//
// Solidity: function removeValidators(address[] _validatorList) returns()
func (_Validator *ValidatorTransactor) RemoveValidators(opts *bind.TransactOpts, _validatorList []common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "removeValidators", _validatorList)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0x1d40f0d8.
//
// Solidity: function removeValidators(address[] _validatorList) returns()
func (_Validator *ValidatorSession) RemoveValidators(_validatorList []common.Address) (*types.Transaction, error) {
	return _Validator.Contract.RemoveValidators(&_Validator.TransactOpts, _validatorList)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0x1d40f0d8.
//
// Solidity: function removeValidators(address[] _validatorList) returns()
func (_Validator *ValidatorTransactorSession) RemoveValidators(_validatorList []common.Address) (*types.Transaction, error) {
	return _Validator.Contract.RemoveValidators(&_Validator.TransactOpts, _validatorList)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 _numerator, uint256 _denominator) returns(uint256 _previousNum, uint256 _previousDenom)
func (_Validator *ValidatorTransactor) SetThreshold(opts *bind.TransactOpts, _numerator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "setThreshold", _numerator, _denominator)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 _numerator, uint256 _denominator) returns(uint256 _previousNum, uint256 _previousDenom)
func (_Validator *ValidatorSession) SetThreshold(_numerator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.SetThreshold(&_Validator.TransactOpts, _numerator, _denominator)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 _numerator, uint256 _denominator) returns(uint256 _previousNum, uint256 _previousDenom)
func (_Validator *ValidatorTransactorSession) SetThreshold(_numerator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.SetThreshold(&_Validator.TransactOpts, _numerator, _denominator)
}

// UpdateValidators is a paid mutator transaction binding the contract method 0x98b3ab61.
//
// Solidity: function updateValidators((address,uint256)[] _validatorList) returns()
func (_Validator *ValidatorTransactor) UpdateValidators(opts *bind.TransactOpts, _validatorList []IWeightedValidatorWeightedValidator) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "updateValidators", _validatorList)
}

// UpdateValidators is a paid mutator transaction binding the contract method 0x98b3ab61.
//
// Solidity: function updateValidators((address,uint256)[] _validatorList) returns()
func (_Validator *ValidatorSession) UpdateValidators(_validatorList []IWeightedValidatorWeightedValidator) (*types.Transaction, error) {
	return _Validator.Contract.UpdateValidators(&_Validator.TransactOpts, _validatorList)
}

// UpdateValidators is a paid mutator transaction binding the contract method 0x98b3ab61.
//
// Solidity: function updateValidators((address,uint256)[] _validatorList) returns()
func (_Validator *ValidatorTransactorSession) UpdateValidators(_validatorList []IWeightedValidatorWeightedValidator) (*types.Transaction, error) {
	return _Validator.Contract.UpdateValidators(&_Validator.TransactOpts, _validatorList)
}

// ValidatorThresholdUpdatedIterator is returned from FilterThresholdUpdated and is used to iterate over the raw logs and unpacked data for ThresholdUpdated events raised by the Validator contract.
type ValidatorThresholdUpdatedIterator struct {
	Event *ValidatorThresholdUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorThresholdUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorThresholdUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorThresholdUpdated represents a ThresholdUpdated event raised by the Validator contract.
type ValidatorThresholdUpdated struct {
	Nonce               *big.Int
	Numerator           *big.Int
	Denominator         *big.Int
	PreviousNumerator   *big.Int
	PreviousDenominator *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterThresholdUpdated is a free log retrieval operation binding the contract event 0x976f8a9c5bdf8248dec172376d6e2b80a8e3df2f0328e381c6db8e1cf138c0f8.
//
// Solidity: event ThresholdUpdated(uint256 indexed _nonce, uint256 indexed _numerator, uint256 indexed _denominator, uint256 _previousNumerator, uint256 _previousDenominator)
func (_Validator *ValidatorFilterer) FilterThresholdUpdated(opts *bind.FilterOpts, _nonce []*big.Int, _numerator []*big.Int, _denominator []*big.Int) (*ValidatorThresholdUpdatedIterator, error) {

	var _nonceRule []interface{}
	for _, _nonceItem := range _nonce {
		_nonceRule = append(_nonceRule, _nonceItem)
	}
	var _numeratorRule []interface{}
	for _, _numeratorItem := range _numerator {
		_numeratorRule = append(_numeratorRule, _numeratorItem)
	}
	var _denominatorRule []interface{}
	for _, _denominatorItem := range _denominator {
		_denominatorRule = append(_denominatorRule, _denominatorItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "ThresholdUpdated", _nonceRule, _numeratorRule, _denominatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorThresholdUpdatedIterator{contract: _Validator.contract, event: "ThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchThresholdUpdated is a free log subscription operation binding the contract event 0x976f8a9c5bdf8248dec172376d6e2b80a8e3df2f0328e381c6db8e1cf138c0f8.
//
// Solidity: event ThresholdUpdated(uint256 indexed _nonce, uint256 indexed _numerator, uint256 indexed _denominator, uint256 _previousNumerator, uint256 _previousDenominator)
func (_Validator *ValidatorFilterer) WatchThresholdUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorThresholdUpdated, _nonce []*big.Int, _numerator []*big.Int, _denominator []*big.Int) (event.Subscription, error) {

	var _nonceRule []interface{}
	for _, _nonceItem := range _nonce {
		_nonceRule = append(_nonceRule, _nonceItem)
	}
	var _numeratorRule []interface{}
	for _, _numeratorItem := range _numerator {
		_numeratorRule = append(_numeratorRule, _numeratorItem)
	}
	var _denominatorRule []interface{}
	for _, _denominatorItem := range _denominator {
		_denominatorRule = append(_denominatorRule, _denominatorItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "ThresholdUpdated", _nonceRule, _numeratorRule, _denominatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorThresholdUpdated)
				if err := _Validator.contract.UnpackLog(event, "ThresholdUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseThresholdUpdated is a log parse operation binding the contract event 0x976f8a9c5bdf8248dec172376d6e2b80a8e3df2f0328e381c6db8e1cf138c0f8.
//
// Solidity: event ThresholdUpdated(uint256 indexed _nonce, uint256 indexed _numerator, uint256 indexed _denominator, uint256 _previousNumerator, uint256 _previousDenominator)
func (_Validator *ValidatorFilterer) ParseThresholdUpdated(log types.Log) (*ValidatorThresholdUpdated, error) {
	event := new(ValidatorThresholdUpdated)
	if err := _Validator.contract.UnpackLog(event, "ThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorValidatorsAddedIterator is returned from FilterValidatorsAdded and is used to iterate over the raw logs and unpacked data for ValidatorsAdded events raised by the Validator contract.
type ValidatorValidatorsAddedIterator struct {
	Event *ValidatorValidatorsAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorValidatorsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorValidatorsAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorValidatorsAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorValidatorsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorValidatorsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorValidatorsAdded represents a ValidatorsAdded event raised by the Validator contract.
type ValidatorValidatorsAdded struct {
	Nonce      *big.Int
	Validators []IWeightedValidatorWeightedValidator
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorsAdded is a free log retrieval operation binding the contract event 0x023add4282789a911fbc45b180f3ac9e120db393149dd03e11ce2b32447f4b6c.
//
// Solidity: event ValidatorsAdded(uint256 indexed _nonce, (address,uint256)[] _validators)
func (_Validator *ValidatorFilterer) FilterValidatorsAdded(opts *bind.FilterOpts, _nonce []*big.Int) (*ValidatorValidatorsAddedIterator, error) {

	var _nonceRule []interface{}
	for _, _nonceItem := range _nonce {
		_nonceRule = append(_nonceRule, _nonceItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "ValidatorsAdded", _nonceRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorValidatorsAddedIterator{contract: _Validator.contract, event: "ValidatorsAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorsAdded is a free log subscription operation binding the contract event 0x023add4282789a911fbc45b180f3ac9e120db393149dd03e11ce2b32447f4b6c.
//
// Solidity: event ValidatorsAdded(uint256 indexed _nonce, (address,uint256)[] _validators)
func (_Validator *ValidatorFilterer) WatchValidatorsAdded(opts *bind.WatchOpts, sink chan<- *ValidatorValidatorsAdded, _nonce []*big.Int) (event.Subscription, error) {

	var _nonceRule []interface{}
	for _, _nonceItem := range _nonce {
		_nonceRule = append(_nonceRule, _nonceItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "ValidatorsAdded", _nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorValidatorsAdded)
				if err := _Validator.contract.UnpackLog(event, "ValidatorsAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorsAdded is a log parse operation binding the contract event 0x023add4282789a911fbc45b180f3ac9e120db393149dd03e11ce2b32447f4b6c.
//
// Solidity: event ValidatorsAdded(uint256 indexed _nonce, (address,uint256)[] _validators)
func (_Validator *ValidatorFilterer) ParseValidatorsAdded(log types.Log) (*ValidatorValidatorsAdded, error) {
	event := new(ValidatorValidatorsAdded)
	if err := _Validator.contract.UnpackLog(event, "ValidatorsAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorValidatorsRemovedIterator is returned from FilterValidatorsRemoved and is used to iterate over the raw logs and unpacked data for ValidatorsRemoved events raised by the Validator contract.
type ValidatorValidatorsRemovedIterator struct {
	Event *ValidatorValidatorsRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorValidatorsRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorValidatorsRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorValidatorsRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorValidatorsRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorValidatorsRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorValidatorsRemoved represents a ValidatorsRemoved event raised by the Validator contract.
type ValidatorValidatorsRemoved struct {
	Nonce      *big.Int
	Validators []common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorsRemoved is a free log retrieval operation binding the contract event 0x100be0d21e2a3eeb62d3a8e929368953ff0bfd628dbe935d27006acb1c7c4772.
//
// Solidity: event ValidatorsRemoved(uint256 indexed _nonce, address[] _validators)
func (_Validator *ValidatorFilterer) FilterValidatorsRemoved(opts *bind.FilterOpts, _nonce []*big.Int) (*ValidatorValidatorsRemovedIterator, error) {

	var _nonceRule []interface{}
	for _, _nonceItem := range _nonce {
		_nonceRule = append(_nonceRule, _nonceItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "ValidatorsRemoved", _nonceRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorValidatorsRemovedIterator{contract: _Validator.contract, event: "ValidatorsRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorsRemoved is a free log subscription operation binding the contract event 0x100be0d21e2a3eeb62d3a8e929368953ff0bfd628dbe935d27006acb1c7c4772.
//
// Solidity: event ValidatorsRemoved(uint256 indexed _nonce, address[] _validators)
func (_Validator *ValidatorFilterer) WatchValidatorsRemoved(opts *bind.WatchOpts, sink chan<- *ValidatorValidatorsRemoved, _nonce []*big.Int) (event.Subscription, error) {

	var _nonceRule []interface{}
	for _, _nonceItem := range _nonce {
		_nonceRule = append(_nonceRule, _nonceItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "ValidatorsRemoved", _nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorValidatorsRemoved)
				if err := _Validator.contract.UnpackLog(event, "ValidatorsRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorsRemoved is a log parse operation binding the contract event 0x100be0d21e2a3eeb62d3a8e929368953ff0bfd628dbe935d27006acb1c7c4772.
//
// Solidity: event ValidatorsRemoved(uint256 indexed _nonce, address[] _validators)
func (_Validator *ValidatorFilterer) ParseValidatorsRemoved(log types.Log) (*ValidatorValidatorsRemoved, error) {
	event := new(ValidatorValidatorsRemoved)
	if err := _Validator.contract.UnpackLog(event, "ValidatorsRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorValidatorsUpdatedIterator is returned from FilterValidatorsUpdated and is used to iterate over the raw logs and unpacked data for ValidatorsUpdated events raised by the Validator contract.
type ValidatorValidatorsUpdatedIterator struct {
	Event *ValidatorValidatorsUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ValidatorValidatorsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorValidatorsUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ValidatorValidatorsUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ValidatorValidatorsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorValidatorsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorValidatorsUpdated represents a ValidatorsUpdated event raised by the Validator contract.
type ValidatorValidatorsUpdated struct {
	Nonce      *big.Int
	Validators []IWeightedValidatorWeightedValidator
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorsUpdated is a free log retrieval operation binding the contract event 0xbe828cc9c0bead3b0fae154a56ceef44df9f828833e44e5415972865e76bb81a.
//
// Solidity: event ValidatorsUpdated(uint256 indexed _nonce, (address,uint256)[] _validators)
func (_Validator *ValidatorFilterer) FilterValidatorsUpdated(opts *bind.FilterOpts, _nonce []*big.Int) (*ValidatorValidatorsUpdatedIterator, error) {

	var _nonceRule []interface{}
	for _, _nonceItem := range _nonce {
		_nonceRule = append(_nonceRule, _nonceItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "ValidatorsUpdated", _nonceRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorValidatorsUpdatedIterator{contract: _Validator.contract, event: "ValidatorsUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorsUpdated is a free log subscription operation binding the contract event 0xbe828cc9c0bead3b0fae154a56ceef44df9f828833e44e5415972865e76bb81a.
//
// Solidity: event ValidatorsUpdated(uint256 indexed _nonce, (address,uint256)[] _validators)
func (_Validator *ValidatorFilterer) WatchValidatorsUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorValidatorsUpdated, _nonce []*big.Int) (event.Subscription, error) {

	var _nonceRule []interface{}
	for _, _nonceItem := range _nonce {
		_nonceRule = append(_nonceRule, _nonceItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "ValidatorsUpdated", _nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorValidatorsUpdated)
				if err := _Validator.contract.UnpackLog(event, "ValidatorsUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorsUpdated is a log parse operation binding the contract event 0xbe828cc9c0bead3b0fae154a56ceef44df9f828833e44e5415972865e76bb81a.
//
// Solidity: event ValidatorsUpdated(uint256 indexed _nonce, (address,uint256)[] _validators)
func (_Validator *ValidatorFilterer) ParseValidatorsUpdated(log types.Log) (*ValidatorValidatorsUpdated, error) {
	event := new(ValidatorValidatorsUpdated)
	if err := _Validator.contract.UnpackLog(event, "ValidatorsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
