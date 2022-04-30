// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package common

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

// BridgeProposalGlobalProposal is an auto generated low-level Go binding around an user-defined struct.
type BridgeProposalGlobalProposal struct {
	TargetOptions []uint8
	Values        []*big.Int
	Calldatas     [][]byte
}

// ProposalProposalDetail is an auto generated low-level Go binding around an user-defined struct.
type ProposalProposalDetail struct {
	ChainId   *big.Int
	Targets   []common.Address
	Values    []*big.Int
	Calldatas [][]byte
}

// SignatureConsumerSignature is an auto generated low-level Go binding around an user-defined struct.
type SignatureConsumerSignature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// CommonMetaData contains all meta data concerning the Common contract.
var CommonMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gatewayContract\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_relayers\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"GatewayContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"enumBridgeProposal.TargetOption[]\",\"name\":\"targetOptions\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"}],\"indexed\":false,\"internalType\":\"structBridgeProposal.GlobalProposal\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"GlobalProposalApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteWeight\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"}],\"indexed\":false,\"internalType\":\"structProposal.ProposalDetail\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"ProposalApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"ProposalVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ValidatorContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"VoteRejected\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeProxyAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gatewayContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"getProxyAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"getProxyImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVotingSignature\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature\",\"name\":\"_sig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rejectCurrentVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"round\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gatewayContract\",\"type\":\"address\"}],\"name\":\"setGatewayContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_num\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_denom\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorContract\",\"type\":\"address\"}],\"name\":\"setValidatorContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[{\"internalType\":\"enumGovernorCore.VoteStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"finalHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"enumBridgeProposal.TargetOption[]\",\"name\":\"targetOptions\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"}],\"internalType\":\"structBridgeProposal.GlobalProposal\",\"name\":\"_bridgeGlobalProposal\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"name\":\"voteGlobalProposalBySignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"}],\"internalType\":\"structProposal.ProposalDetail\",\"name\":\"_proposal\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"name\":\"voteProposalBySignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CommonABI is the input ABI used to generate the binding from.
// Deprecated: Use CommonMetaData.ABI instead.
var CommonABI = CommonMetaData.ABI

// Common is an auto generated Go binding around an Ethereum contract.
type Common struct {
	CommonCaller     // Read-only binding to the contract
	CommonTransactor // Write-only binding to the contract
	CommonFilterer   // Log filterer for contract events
}

// CommonCaller is an auto generated read-only Go binding around an Ethereum contract.
type CommonCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommonTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CommonTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommonFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CommonFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommonSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CommonSession struct {
	Contract     *Common           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CommonCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CommonCallerSession struct {
	Contract *CommonCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CommonTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CommonTransactorSession struct {
	Contract     *CommonTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CommonRaw is an auto generated low-level Go binding around an Ethereum contract.
type CommonRaw struct {
	Contract *Common // Generic contract binding to access the raw methods on
}

// CommonCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CommonCallerRaw struct {
	Contract *CommonCaller // Generic read-only contract binding to access the raw methods on
}

// CommonTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CommonTransactorRaw struct {
	Contract *CommonTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCommon creates a new instance of Common, bound to a specific deployed contract.
func NewCommon(address common.Address, backend bind.ContractBackend) (*Common, error) {
	contract, err := bindCommon(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Common{CommonCaller: CommonCaller{contract: contract}, CommonTransactor: CommonTransactor{contract: contract}, CommonFilterer: CommonFilterer{contract: contract}}, nil
}

// NewCommonCaller creates a new read-only instance of Common, bound to a specific deployed contract.
func NewCommonCaller(address common.Address, caller bind.ContractCaller) (*CommonCaller, error) {
	contract, err := bindCommon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CommonCaller{contract: contract}, nil
}

// NewCommonTransactor creates a new write-only instance of Common, bound to a specific deployed contract.
func NewCommonTransactor(address common.Address, transactor bind.ContractTransactor) (*CommonTransactor, error) {
	contract, err := bindCommon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CommonTransactor{contract: contract}, nil
}

// NewCommonFilterer creates a new log filterer instance of Common, bound to a specific deployed contract.
func NewCommonFilterer(address common.Address, filterer bind.ContractFilterer) (*CommonFilterer, error) {
	contract, err := bindCommon(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CommonFilterer{contract: contract}, nil
}

// bindCommon binds a generic wrapper to an already deployed contract.
func bindCommon(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CommonABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Common *CommonRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Common.Contract.CommonCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Common *CommonRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Common.Contract.CommonTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Common *CommonRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Common.Contract.CommonTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Common *CommonCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Common.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Common *CommonTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Common.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Common *CommonTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Common.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Common *CommonCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Common.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Common *CommonSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Common.Contract.DEFAULTADMINROLE(&_Common.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Common *CommonCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Common.Contract.DEFAULTADMINROLE(&_Common.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Common *CommonCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Common.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Common *CommonSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Common.Contract.DOMAINSEPARATOR(&_Common.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Common *CommonCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Common.Contract.DOMAINSEPARATOR(&_Common.CallOpts)
}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_Common *CommonCaller) RELAYERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Common.contract.Call(opts, &out, "RELAYER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_Common *CommonSession) RELAYERROLE() ([32]byte, error) {
	return _Common.Contract.RELAYERROLE(&_Common.CallOpts)
}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_Common *CommonCallerSession) RELAYERROLE() ([32]byte, error) {
	return _Common.Contract.RELAYERROLE(&_Common.CallOpts)
}

// GatewayContract is a free data retrieval call binding the contract method 0xeb0cde1d.
//
// Solidity: function gatewayContract() view returns(address)
func (_Common *CommonCaller) GatewayContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Common.contract.Call(opts, &out, "gatewayContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GatewayContract is a free data retrieval call binding the contract method 0xeb0cde1d.
//
// Solidity: function gatewayContract() view returns(address)
func (_Common *CommonSession) GatewayContract() (common.Address, error) {
	return _Common.Contract.GatewayContract(&_Common.CallOpts)
}

// GatewayContract is a free data retrieval call binding the contract method 0xeb0cde1d.
//
// Solidity: function gatewayContract() view returns(address)
func (_Common *CommonCallerSession) GatewayContract() (common.Address, error) {
	return _Common.Contract.GatewayContract(&_Common.CallOpts)
}

// GetProxyAdmin is a free data retrieval call binding the contract method 0xf3b7dead.
//
// Solidity: function getProxyAdmin(address _proxy) view returns(address)
func (_Common *CommonCaller) GetProxyAdmin(opts *bind.CallOpts, _proxy common.Address) (common.Address, error) {
	var out []interface{}
	err := _Common.contract.Call(opts, &out, "getProxyAdmin", _proxy)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyAdmin is a free data retrieval call binding the contract method 0xf3b7dead.
//
// Solidity: function getProxyAdmin(address _proxy) view returns(address)
func (_Common *CommonSession) GetProxyAdmin(_proxy common.Address) (common.Address, error) {
	return _Common.Contract.GetProxyAdmin(&_Common.CallOpts, _proxy)
}

// GetProxyAdmin is a free data retrieval call binding the contract method 0xf3b7dead.
//
// Solidity: function getProxyAdmin(address _proxy) view returns(address)
func (_Common *CommonCallerSession) GetProxyAdmin(_proxy common.Address) (common.Address, error) {
	return _Common.Contract.GetProxyAdmin(&_Common.CallOpts, _proxy)
}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x204e1c7a.
//
// Solidity: function getProxyImplementation(address _proxy) view returns(address)
func (_Common *CommonCaller) GetProxyImplementation(opts *bind.CallOpts, _proxy common.Address) (common.Address, error) {
	var out []interface{}
	err := _Common.contract.Call(opts, &out, "getProxyImplementation", _proxy)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x204e1c7a.
//
// Solidity: function getProxyImplementation(address _proxy) view returns(address)
func (_Common *CommonSession) GetProxyImplementation(_proxy common.Address) (common.Address, error) {
	return _Common.Contract.GetProxyImplementation(&_Common.CallOpts, _proxy)
}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x204e1c7a.
//
// Solidity: function getProxyImplementation(address _proxy) view returns(address)
func (_Common *CommonCallerSession) GetProxyImplementation(_proxy common.Address) (common.Address, error) {
	return _Common.Contract.GetProxyImplementation(&_Common.CallOpts, _proxy)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Common *CommonCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Common.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Common *CommonSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Common.Contract.GetRoleAdmin(&_Common.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Common *CommonCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Common.Contract.GetRoleAdmin(&_Common.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Common *CommonCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Common.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Common *CommonSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Common.Contract.GetRoleMember(&_Common.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Common *CommonCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Common.Contract.GetRoleMember(&_Common.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Common *CommonCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Common.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Common *CommonSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Common.Contract.GetRoleMemberCount(&_Common.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Common *CommonCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Common.Contract.GetRoleMemberCount(&_Common.CallOpts, role)
}

// GetVotingSignature is a free data retrieval call binding the contract method 0xe6883cd3.
//
// Solidity: function getVotingSignature(uint256 _round, address _voter) view returns((uint8,bytes32,bytes32) _sig)
func (_Common *CommonCaller) GetVotingSignature(opts *bind.CallOpts, _round *big.Int, _voter common.Address) (SignatureConsumerSignature, error) {
	var out []interface{}
	err := _Common.contract.Call(opts, &out, "getVotingSignature", _round, _voter)

	if err != nil {
		return *new(SignatureConsumerSignature), err
	}

	out0 := *abi.ConvertType(out[0], new(SignatureConsumerSignature)).(*SignatureConsumerSignature)

	return out0, err

}

// GetVotingSignature is a free data retrieval call binding the contract method 0xe6883cd3.
//
// Solidity: function getVotingSignature(uint256 _round, address _voter) view returns((uint8,bytes32,bytes32) _sig)
func (_Common *CommonSession) GetVotingSignature(_round *big.Int, _voter common.Address) (SignatureConsumerSignature, error) {
	return _Common.Contract.GetVotingSignature(&_Common.CallOpts, _round, _voter)
}

// GetVotingSignature is a free data retrieval call binding the contract method 0xe6883cd3.
//
// Solidity: function getVotingSignature(uint256 _round, address _voter) view returns((uint8,bytes32,bytes32) _sig)
func (_Common *CommonCallerSession) GetVotingSignature(_round *big.Int, _voter common.Address) (SignatureConsumerSignature, error) {
	return _Common.Contract.GetVotingSignature(&_Common.CallOpts, _round, _voter)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Common *CommonCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Common.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Common *CommonSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Common.Contract.HasRole(&_Common.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Common *CommonCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Common.Contract.HasRole(&_Common.CallOpts, role, account)
}

// Round is a free data retrieval call binding the contract method 0x146ca531.
//
// Solidity: function round() view returns(uint256)
func (_Common *CommonCaller) Round(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Common.contract.Call(opts, &out, "round")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Round is a free data retrieval call binding the contract method 0x146ca531.
//
// Solidity: function round() view returns(uint256)
func (_Common *CommonSession) Round() (*big.Int, error) {
	return _Common.Contract.Round(&_Common.CallOpts)
}

// Round is a free data retrieval call binding the contract method 0x146ca531.
//
// Solidity: function round() view returns(uint256)
func (_Common *CommonCallerSession) Round() (*big.Int, error) {
	return _Common.Contract.Round(&_Common.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Common *CommonCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Common.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Common *CommonSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Common.Contract.SupportsInterface(&_Common.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Common *CommonCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Common.Contract.SupportsInterface(&_Common.CallOpts, interfaceId)
}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_Common *CommonCaller) ValidatorContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Common.contract.Call(opts, &out, "validatorContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_Common *CommonSession) ValidatorContract() (common.Address, error) {
	return _Common.Contract.ValidatorContract(&_Common.CallOpts)
}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_Common *CommonCallerSession) ValidatorContract() (common.Address, error) {
	return _Common.Contract.ValidatorContract(&_Common.CallOpts)
}

// Vote is a free data retrieval call binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 ) view returns(uint8 status, bytes32 finalHash)
func (_Common *CommonCaller) Vote(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Status    uint8
	FinalHash [32]byte
}, error) {
	var out []interface{}
	err := _Common.contract.Call(opts, &out, "vote", arg0)

	outstruct := new(struct {
		Status    uint8
		FinalHash [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.FinalHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Vote is a free data retrieval call binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 ) view returns(uint8 status, bytes32 finalHash)
func (_Common *CommonSession) Vote(arg0 *big.Int) (struct {
	Status    uint8
	FinalHash [32]byte
}, error) {
	return _Common.Contract.Vote(&_Common.CallOpts, arg0)
}

// Vote is a free data retrieval call binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 ) view returns(uint8 status, bytes32 finalHash)
func (_Common *CommonCallerSession) Vote(arg0 *big.Int) (struct {
	Status    uint8
	FinalHash [32]byte
}, error) {
	return _Common.Contract.Vote(&_Common.CallOpts, arg0)
}

// ChangeProxyAdmin is a paid mutator transaction binding the contract method 0x7eff275e.
//
// Solidity: function changeProxyAdmin(address _proxy, address _newAdmin) returns()
func (_Common *CommonTransactor) ChangeProxyAdmin(opts *bind.TransactOpts, _proxy common.Address, _newAdmin common.Address) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "changeProxyAdmin", _proxy, _newAdmin)
}

// ChangeProxyAdmin is a paid mutator transaction binding the contract method 0x7eff275e.
//
// Solidity: function changeProxyAdmin(address _proxy, address _newAdmin) returns()
func (_Common *CommonSession) ChangeProxyAdmin(_proxy common.Address, _newAdmin common.Address) (*types.Transaction, error) {
	return _Common.Contract.ChangeProxyAdmin(&_Common.TransactOpts, _proxy, _newAdmin)
}

// ChangeProxyAdmin is a paid mutator transaction binding the contract method 0x7eff275e.
//
// Solidity: function changeProxyAdmin(address _proxy, address _newAdmin) returns()
func (_Common *CommonTransactorSession) ChangeProxyAdmin(_proxy common.Address, _newAdmin common.Address) (*types.Transaction, error) {
	return _Common.Contract.ChangeProxyAdmin(&_Common.TransactOpts, _proxy, _newAdmin)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Common *CommonTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Common *CommonSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Common.Contract.GrantRole(&_Common.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Common *CommonTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Common.Contract.GrantRole(&_Common.TransactOpts, role, account)
}

// RejectCurrentVote is a paid mutator transaction binding the contract method 0xa159974c.
//
// Solidity: function rejectCurrentVote() returns()
func (_Common *CommonTransactor) RejectCurrentVote(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "rejectCurrentVote")
}

// RejectCurrentVote is a paid mutator transaction binding the contract method 0xa159974c.
//
// Solidity: function rejectCurrentVote() returns()
func (_Common *CommonSession) RejectCurrentVote() (*types.Transaction, error) {
	return _Common.Contract.RejectCurrentVote(&_Common.TransactOpts)
}

// RejectCurrentVote is a paid mutator transaction binding the contract method 0xa159974c.
//
// Solidity: function rejectCurrentVote() returns()
func (_Common *CommonTransactorSession) RejectCurrentVote() (*types.Transaction, error) {
	return _Common.Contract.RejectCurrentVote(&_Common.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Common *CommonTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Common *CommonSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Common.Contract.RenounceRole(&_Common.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Common *CommonTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Common.Contract.RenounceRole(&_Common.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Common *CommonTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Common *CommonSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Common.Contract.RevokeRole(&_Common.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Common *CommonTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Common.Contract.RevokeRole(&_Common.TransactOpts, role, account)
}

// SetGatewayContract is a paid mutator transaction binding the contract method 0x50932cb5.
//
// Solidity: function setGatewayContract(address _gatewayContract) returns()
func (_Common *CommonTransactor) SetGatewayContract(opts *bind.TransactOpts, _gatewayContract common.Address) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "setGatewayContract", _gatewayContract)
}

// SetGatewayContract is a paid mutator transaction binding the contract method 0x50932cb5.
//
// Solidity: function setGatewayContract(address _gatewayContract) returns()
func (_Common *CommonSession) SetGatewayContract(_gatewayContract common.Address) (*types.Transaction, error) {
	return _Common.Contract.SetGatewayContract(&_Common.TransactOpts, _gatewayContract)
}

// SetGatewayContract is a paid mutator transaction binding the contract method 0x50932cb5.
//
// Solidity: function setGatewayContract(address _gatewayContract) returns()
func (_Common *CommonTransactorSession) SetGatewayContract(_gatewayContract common.Address) (*types.Transaction, error) {
	return _Common.Contract.SetGatewayContract(&_Common.TransactOpts, _gatewayContract)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xa4e35076.
//
// Solidity: function setThreshold(address _contract, uint256 _num, uint256 _denom) returns()
func (_Common *CommonTransactor) SetThreshold(opts *bind.TransactOpts, _contract common.Address, _num *big.Int, _denom *big.Int) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "setThreshold", _contract, _num, _denom)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xa4e35076.
//
// Solidity: function setThreshold(address _contract, uint256 _num, uint256 _denom) returns()
func (_Common *CommonSession) SetThreshold(_contract common.Address, _num *big.Int, _denom *big.Int) (*types.Transaction, error) {
	return _Common.Contract.SetThreshold(&_Common.TransactOpts, _contract, _num, _denom)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xa4e35076.
//
// Solidity: function setThreshold(address _contract, uint256 _num, uint256 _denom) returns()
func (_Common *CommonTransactorSession) SetThreshold(_contract common.Address, _num *big.Int, _denom *big.Int) (*types.Transaction, error) {
	return _Common.Contract.SetThreshold(&_Common.TransactOpts, _contract, _num, _denom)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _validatorContract) returns()
func (_Common *CommonTransactor) SetValidatorContract(opts *bind.TransactOpts, _validatorContract common.Address) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "setValidatorContract", _validatorContract)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _validatorContract) returns()
func (_Common *CommonSession) SetValidatorContract(_validatorContract common.Address) (*types.Transaction, error) {
	return _Common.Contract.SetValidatorContract(&_Common.TransactOpts, _validatorContract)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _validatorContract) returns()
func (_Common *CommonTransactorSession) SetValidatorContract(_validatorContract common.Address) (*types.Transaction, error) {
	return _Common.Contract.SetValidatorContract(&_Common.TransactOpts, _validatorContract)
}

// UpgradeAndCall is a paid mutator transaction binding the contract method 0x9623609d.
//
// Solidity: function upgradeAndCall(address _proxy, address _implementation, bytes _data) payable returns()
func (_Common *CommonTransactor) UpgradeAndCall(opts *bind.TransactOpts, _proxy common.Address, _implementation common.Address, _data []byte) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "upgradeAndCall", _proxy, _implementation, _data)
}

// UpgradeAndCall is a paid mutator transaction binding the contract method 0x9623609d.
//
// Solidity: function upgradeAndCall(address _proxy, address _implementation, bytes _data) payable returns()
func (_Common *CommonSession) UpgradeAndCall(_proxy common.Address, _implementation common.Address, _data []byte) (*types.Transaction, error) {
	return _Common.Contract.UpgradeAndCall(&_Common.TransactOpts, _proxy, _implementation, _data)
}

// UpgradeAndCall is a paid mutator transaction binding the contract method 0x9623609d.
//
// Solidity: function upgradeAndCall(address _proxy, address _implementation, bytes _data) payable returns()
func (_Common *CommonTransactorSession) UpgradeAndCall(_proxy common.Address, _implementation common.Address, _data []byte) (*types.Transaction, error) {
	return _Common.Contract.UpgradeAndCall(&_Common.TransactOpts, _proxy, _implementation, _data)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x132d807e.
//
// Solidity: function upgradeTo(address _proxy, address _implementation) returns()
func (_Common *CommonTransactor) UpgradeTo(opts *bind.TransactOpts, _proxy common.Address, _implementation common.Address) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "upgradeTo", _proxy, _implementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x132d807e.
//
// Solidity: function upgradeTo(address _proxy, address _implementation) returns()
func (_Common *CommonSession) UpgradeTo(_proxy common.Address, _implementation common.Address) (*types.Transaction, error) {
	return _Common.Contract.UpgradeTo(&_Common.TransactOpts, _proxy, _implementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x132d807e.
//
// Solidity: function upgradeTo(address _proxy, address _implementation) returns()
func (_Common *CommonTransactorSession) UpgradeTo(_proxy common.Address, _implementation common.Address) (*types.Transaction, error) {
	return _Common.Contract.UpgradeTo(&_Common.TransactOpts, _proxy, _implementation)
}

// VoteGlobalProposalBySignatures is a paid mutator transaction binding the contract method 0xca11511e.
//
// Solidity: function voteGlobalProposalBySignatures(uint256 _round, (uint8[],uint256[],bytes[]) _bridgeGlobalProposal, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Common *CommonTransactor) VoteGlobalProposalBySignatures(opts *bind.TransactOpts, _round *big.Int, _bridgeGlobalProposal BridgeProposalGlobalProposal, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "voteGlobalProposalBySignatures", _round, _bridgeGlobalProposal, _signatures)
}

// VoteGlobalProposalBySignatures is a paid mutator transaction binding the contract method 0xca11511e.
//
// Solidity: function voteGlobalProposalBySignatures(uint256 _round, (uint8[],uint256[],bytes[]) _bridgeGlobalProposal, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Common *CommonSession) VoteGlobalProposalBySignatures(_round *big.Int, _bridgeGlobalProposal BridgeProposalGlobalProposal, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Common.Contract.VoteGlobalProposalBySignatures(&_Common.TransactOpts, _round, _bridgeGlobalProposal, _signatures)
}

// VoteGlobalProposalBySignatures is a paid mutator transaction binding the contract method 0xca11511e.
//
// Solidity: function voteGlobalProposalBySignatures(uint256 _round, (uint8[],uint256[],bytes[]) _bridgeGlobalProposal, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Common *CommonTransactorSession) VoteGlobalProposalBySignatures(_round *big.Int, _bridgeGlobalProposal BridgeProposalGlobalProposal, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Common.Contract.VoteGlobalProposalBySignatures(&_Common.TransactOpts, _round, _bridgeGlobalProposal, _signatures)
}

// VoteProposalBySignatures is a paid mutator transaction binding the contract method 0xf956d303.
//
// Solidity: function voteProposalBySignatures(uint256 _round, (uint256,address[],uint256[],bytes[]) _proposal, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Common *CommonTransactor) VoteProposalBySignatures(opts *bind.TransactOpts, _round *big.Int, _proposal ProposalProposalDetail, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "voteProposalBySignatures", _round, _proposal, _signatures)
}

// VoteProposalBySignatures is a paid mutator transaction binding the contract method 0xf956d303.
//
// Solidity: function voteProposalBySignatures(uint256 _round, (uint256,address[],uint256[],bytes[]) _proposal, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Common *CommonSession) VoteProposalBySignatures(_round *big.Int, _proposal ProposalProposalDetail, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Common.Contract.VoteProposalBySignatures(&_Common.TransactOpts, _round, _proposal, _signatures)
}

// VoteProposalBySignatures is a paid mutator transaction binding the contract method 0xf956d303.
//
// Solidity: function voteProposalBySignatures(uint256 _round, (uint256,address[],uint256[],bytes[]) _proposal, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Common *CommonTransactorSession) VoteProposalBySignatures(_round *big.Int, _proposal ProposalProposalDetail, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Common.Contract.VoteProposalBySignatures(&_Common.TransactOpts, _round, _proposal, _signatures)
}

// CommonGatewayContractUpdatedIterator is returned from FilterGatewayContractUpdated and is used to iterate over the raw logs and unpacked data for GatewayContractUpdated events raised by the Common contract.
type CommonGatewayContractUpdatedIterator struct {
	Event *CommonGatewayContractUpdated // Event containing the contract specifics and raw log

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
func (it *CommonGatewayContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommonGatewayContractUpdated)
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
		it.Event = new(CommonGatewayContractUpdated)
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
func (it *CommonGatewayContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommonGatewayContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommonGatewayContractUpdated represents a GatewayContractUpdated event raised by the Common contract.
type CommonGatewayContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterGatewayContractUpdated is a free log retrieval operation binding the contract event 0x7f0264886136f6ac42676074e865c449ef8d885176adb0bf78f47aeb3674e9ac.
//
// Solidity: event GatewayContractUpdated(address arg0)
func (_Common *CommonFilterer) FilterGatewayContractUpdated(opts *bind.FilterOpts) (*CommonGatewayContractUpdatedIterator, error) {

	logs, sub, err := _Common.contract.FilterLogs(opts, "GatewayContractUpdated")
	if err != nil {
		return nil, err
	}
	return &CommonGatewayContractUpdatedIterator{contract: _Common.contract, event: "GatewayContractUpdated", logs: logs, sub: sub}, nil
}

// WatchGatewayContractUpdated is a free log subscription operation binding the contract event 0x7f0264886136f6ac42676074e865c449ef8d885176adb0bf78f47aeb3674e9ac.
//
// Solidity: event GatewayContractUpdated(address arg0)
func (_Common *CommonFilterer) WatchGatewayContractUpdated(opts *bind.WatchOpts, sink chan<- *CommonGatewayContractUpdated) (event.Subscription, error) {

	logs, sub, err := _Common.contract.WatchLogs(opts, "GatewayContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommonGatewayContractUpdated)
				if err := _Common.contract.UnpackLog(event, "GatewayContractUpdated", log); err != nil {
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

// ParseGatewayContractUpdated is a log parse operation binding the contract event 0x7f0264886136f6ac42676074e865c449ef8d885176adb0bf78f47aeb3674e9ac.
//
// Solidity: event GatewayContractUpdated(address arg0)
func (_Common *CommonFilterer) ParseGatewayContractUpdated(log types.Log) (*CommonGatewayContractUpdated, error) {
	event := new(CommonGatewayContractUpdated)
	if err := _Common.contract.UnpackLog(event, "GatewayContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommonGlobalProposalApprovedIterator is returned from FilterGlobalProposalApproved and is used to iterate over the raw logs and unpacked data for GlobalProposalApproved events raised by the Common contract.
type CommonGlobalProposalApprovedIterator struct {
	Event *CommonGlobalProposalApproved // Event containing the contract specifics and raw log

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
func (it *CommonGlobalProposalApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommonGlobalProposalApproved)
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
		it.Event = new(CommonGlobalProposalApproved)
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
func (it *CommonGlobalProposalApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommonGlobalProposalApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommonGlobalProposalApproved represents a GlobalProposalApproved event raised by the Common contract.
type CommonGlobalProposalApproved struct {
	Round *big.Int
	Arg1  BridgeProposalGlobalProposal
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGlobalProposalApproved is a free log retrieval operation binding the contract event 0x4e2691a301841dfedf72fbf93e4913da3cea0002fa205049ce326d14e514b07f.
//
// Solidity: event GlobalProposalApproved(uint256 indexed round, (uint8[],uint256[],bytes[]) arg1)
func (_Common *CommonFilterer) FilterGlobalProposalApproved(opts *bind.FilterOpts, round []*big.Int) (*CommonGlobalProposalApprovedIterator, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _Common.contract.FilterLogs(opts, "GlobalProposalApproved", roundRule)
	if err != nil {
		return nil, err
	}
	return &CommonGlobalProposalApprovedIterator{contract: _Common.contract, event: "GlobalProposalApproved", logs: logs, sub: sub}, nil
}

// WatchGlobalProposalApproved is a free log subscription operation binding the contract event 0x4e2691a301841dfedf72fbf93e4913da3cea0002fa205049ce326d14e514b07f.
//
// Solidity: event GlobalProposalApproved(uint256 indexed round, (uint8[],uint256[],bytes[]) arg1)
func (_Common *CommonFilterer) WatchGlobalProposalApproved(opts *bind.WatchOpts, sink chan<- *CommonGlobalProposalApproved, round []*big.Int) (event.Subscription, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _Common.contract.WatchLogs(opts, "GlobalProposalApproved", roundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommonGlobalProposalApproved)
				if err := _Common.contract.UnpackLog(event, "GlobalProposalApproved", log); err != nil {
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

// ParseGlobalProposalApproved is a log parse operation binding the contract event 0x4e2691a301841dfedf72fbf93e4913da3cea0002fa205049ce326d14e514b07f.
//
// Solidity: event GlobalProposalApproved(uint256 indexed round, (uint8[],uint256[],bytes[]) arg1)
func (_Common *CommonFilterer) ParseGlobalProposalApproved(log types.Log) (*CommonGlobalProposalApproved, error) {
	event := new(CommonGlobalProposalApproved)
	if err := _Common.contract.UnpackLog(event, "GlobalProposalApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommonProposalApprovedIterator is returned from FilterProposalApproved and is used to iterate over the raw logs and unpacked data for ProposalApproved events raised by the Common contract.
type CommonProposalApprovedIterator struct {
	Event *CommonProposalApproved // Event containing the contract specifics and raw log

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
func (it *CommonProposalApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommonProposalApproved)
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
		it.Event = new(CommonProposalApproved)
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
func (it *CommonProposalApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommonProposalApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommonProposalApproved represents a ProposalApproved event raised by the Common contract.
type CommonProposalApproved struct {
	Round      *big.Int
	Hash       [32]byte
	VoteWeight *big.Int
	Arg3       ProposalProposalDetail
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalApproved is a free log retrieval operation binding the contract event 0xbbfc341e57f906c936469ad77d94e3cc961c66e90dbc68529b5fb4be2c97b421.
//
// Solidity: event ProposalApproved(uint256 indexed round, bytes32 hash, uint256 voteWeight, (uint256,address[],uint256[],bytes[]) arg3)
func (_Common *CommonFilterer) FilterProposalApproved(opts *bind.FilterOpts, round []*big.Int) (*CommonProposalApprovedIterator, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _Common.contract.FilterLogs(opts, "ProposalApproved", roundRule)
	if err != nil {
		return nil, err
	}
	return &CommonProposalApprovedIterator{contract: _Common.contract, event: "ProposalApproved", logs: logs, sub: sub}, nil
}

// WatchProposalApproved is a free log subscription operation binding the contract event 0xbbfc341e57f906c936469ad77d94e3cc961c66e90dbc68529b5fb4be2c97b421.
//
// Solidity: event ProposalApproved(uint256 indexed round, bytes32 hash, uint256 voteWeight, (uint256,address[],uint256[],bytes[]) arg3)
func (_Common *CommonFilterer) WatchProposalApproved(opts *bind.WatchOpts, sink chan<- *CommonProposalApproved, round []*big.Int) (event.Subscription, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _Common.contract.WatchLogs(opts, "ProposalApproved", roundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommonProposalApproved)
				if err := _Common.contract.UnpackLog(event, "ProposalApproved", log); err != nil {
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

// ParseProposalApproved is a log parse operation binding the contract event 0xbbfc341e57f906c936469ad77d94e3cc961c66e90dbc68529b5fb4be2c97b421.
//
// Solidity: event ProposalApproved(uint256 indexed round, bytes32 hash, uint256 voteWeight, (uint256,address[],uint256[],bytes[]) arg3)
func (_Common *CommonFilterer) ParseProposalApproved(log types.Log) (*CommonProposalApproved, error) {
	event := new(CommonProposalApproved)
	if err := _Common.contract.UnpackLog(event, "ProposalApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommonProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the Common contract.
type CommonProposalExecutedIterator struct {
	Event *CommonProposalExecuted // Event containing the contract specifics and raw log

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
func (it *CommonProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommonProposalExecuted)
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
		it.Event = new(CommonProposalExecuted)
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
func (it *CommonProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommonProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommonProposalExecuted represents a ProposalExecuted event raised by the Common contract.
type CommonProposalExecuted struct {
	Round *big.Int
	Hash  [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x56a007d3eea04bd347e571f3451382cb2a33ef5fd102b9a63846ff8d787f43cf.
//
// Solidity: event ProposalExecuted(uint256 indexed round, bytes32 hash)
func (_Common *CommonFilterer) FilterProposalExecuted(opts *bind.FilterOpts, round []*big.Int) (*CommonProposalExecutedIterator, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _Common.contract.FilterLogs(opts, "ProposalExecuted", roundRule)
	if err != nil {
		return nil, err
	}
	return &CommonProposalExecutedIterator{contract: _Common.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x56a007d3eea04bd347e571f3451382cb2a33ef5fd102b9a63846ff8d787f43cf.
//
// Solidity: event ProposalExecuted(uint256 indexed round, bytes32 hash)
func (_Common *CommonFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *CommonProposalExecuted, round []*big.Int) (event.Subscription, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _Common.contract.WatchLogs(opts, "ProposalExecuted", roundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommonProposalExecuted)
				if err := _Common.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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

// ParseProposalExecuted is a log parse operation binding the contract event 0x56a007d3eea04bd347e571f3451382cb2a33ef5fd102b9a63846ff8d787f43cf.
//
// Solidity: event ProposalExecuted(uint256 indexed round, bytes32 hash)
func (_Common *CommonFilterer) ParseProposalExecuted(log types.Log) (*CommonProposalExecuted, error) {
	event := new(CommonProposalExecuted)
	if err := _Common.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommonProposalVotedIterator is returned from FilterProposalVoted and is used to iterate over the raw logs and unpacked data for ProposalVoted events raised by the Common contract.
type CommonProposalVotedIterator struct {
	Event *CommonProposalVoted // Event containing the contract specifics and raw log

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
func (it *CommonProposalVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommonProposalVoted)
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
		it.Event = new(CommonProposalVoted)
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
func (it *CommonProposalVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommonProposalVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommonProposalVoted represents a ProposalVoted event raised by the Common contract.
type CommonProposalVoted struct {
	Round  *big.Int
	Voter  common.Address
	Weight *big.Int
	Hash   [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterProposalVoted is a free log retrieval operation binding the contract event 0x013b7e97889f4f50a50bc1f07e7b7921414afd8a4c1a43fe8b9afb60ffbd8544.
//
// Solidity: event ProposalVoted(uint256 indexed round, address indexed voter, uint256 weight, bytes32 hash)
func (_Common *CommonFilterer) FilterProposalVoted(opts *bind.FilterOpts, round []*big.Int, voter []common.Address) (*CommonProposalVotedIterator, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Common.contract.FilterLogs(opts, "ProposalVoted", roundRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &CommonProposalVotedIterator{contract: _Common.contract, event: "ProposalVoted", logs: logs, sub: sub}, nil
}

// WatchProposalVoted is a free log subscription operation binding the contract event 0x013b7e97889f4f50a50bc1f07e7b7921414afd8a4c1a43fe8b9afb60ffbd8544.
//
// Solidity: event ProposalVoted(uint256 indexed round, address indexed voter, uint256 weight, bytes32 hash)
func (_Common *CommonFilterer) WatchProposalVoted(opts *bind.WatchOpts, sink chan<- *CommonProposalVoted, round []*big.Int, voter []common.Address) (event.Subscription, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Common.contract.WatchLogs(opts, "ProposalVoted", roundRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommonProposalVoted)
				if err := _Common.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
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

// ParseProposalVoted is a log parse operation binding the contract event 0x013b7e97889f4f50a50bc1f07e7b7921414afd8a4c1a43fe8b9afb60ffbd8544.
//
// Solidity: event ProposalVoted(uint256 indexed round, address indexed voter, uint256 weight, bytes32 hash)
func (_Common *CommonFilterer) ParseProposalVoted(log types.Log) (*CommonProposalVoted, error) {
	event := new(CommonProposalVoted)
	if err := _Common.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommonRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Common contract.
type CommonRoleAdminChangedIterator struct {
	Event *CommonRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CommonRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommonRoleAdminChanged)
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
		it.Event = new(CommonRoleAdminChanged)
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
func (it *CommonRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommonRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommonRoleAdminChanged represents a RoleAdminChanged event raised by the Common contract.
type CommonRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Common *CommonFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CommonRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Common.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CommonRoleAdminChangedIterator{contract: _Common.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Common *CommonFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CommonRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Common.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommonRoleAdminChanged)
				if err := _Common.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Common *CommonFilterer) ParseRoleAdminChanged(log types.Log) (*CommonRoleAdminChanged, error) {
	event := new(CommonRoleAdminChanged)
	if err := _Common.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommonRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Common contract.
type CommonRoleGrantedIterator struct {
	Event *CommonRoleGranted // Event containing the contract specifics and raw log

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
func (it *CommonRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommonRoleGranted)
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
		it.Event = new(CommonRoleGranted)
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
func (it *CommonRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommonRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommonRoleGranted represents a RoleGranted event raised by the Common contract.
type CommonRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Common *CommonFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CommonRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Common.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CommonRoleGrantedIterator{contract: _Common.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Common *CommonFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CommonRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Common.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommonRoleGranted)
				if err := _Common.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Common *CommonFilterer) ParseRoleGranted(log types.Log) (*CommonRoleGranted, error) {
	event := new(CommonRoleGranted)
	if err := _Common.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommonRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Common contract.
type CommonRoleRevokedIterator struct {
	Event *CommonRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CommonRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommonRoleRevoked)
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
		it.Event = new(CommonRoleRevoked)
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
func (it *CommonRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommonRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommonRoleRevoked represents a RoleRevoked event raised by the Common contract.
type CommonRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Common *CommonFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CommonRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Common.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CommonRoleRevokedIterator{contract: _Common.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Common *CommonFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CommonRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Common.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommonRoleRevoked)
				if err := _Common.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Common *CommonFilterer) ParseRoleRevoked(log types.Log) (*CommonRoleRevoked, error) {
	event := new(CommonRoleRevoked)
	if err := _Common.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommonValidatorContractUpdatedIterator is returned from FilterValidatorContractUpdated and is used to iterate over the raw logs and unpacked data for ValidatorContractUpdated events raised by the Common contract.
type CommonValidatorContractUpdatedIterator struct {
	Event *CommonValidatorContractUpdated // Event containing the contract specifics and raw log

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
func (it *CommonValidatorContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommonValidatorContractUpdated)
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
		it.Event = new(CommonValidatorContractUpdated)
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
func (it *CommonValidatorContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommonValidatorContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommonValidatorContractUpdated represents a ValidatorContractUpdated event raised by the Common contract.
type CommonValidatorContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterValidatorContractUpdated is a free log retrieval operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_Common *CommonFilterer) FilterValidatorContractUpdated(opts *bind.FilterOpts) (*CommonValidatorContractUpdatedIterator, error) {

	logs, sub, err := _Common.contract.FilterLogs(opts, "ValidatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return &CommonValidatorContractUpdatedIterator{contract: _Common.contract, event: "ValidatorContractUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorContractUpdated is a free log subscription operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_Common *CommonFilterer) WatchValidatorContractUpdated(opts *bind.WatchOpts, sink chan<- *CommonValidatorContractUpdated) (event.Subscription, error) {

	logs, sub, err := _Common.contract.WatchLogs(opts, "ValidatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommonValidatorContractUpdated)
				if err := _Common.contract.UnpackLog(event, "ValidatorContractUpdated", log); err != nil {
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

// ParseValidatorContractUpdated is a log parse operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_Common *CommonFilterer) ParseValidatorContractUpdated(log types.Log) (*CommonValidatorContractUpdated, error) {
	event := new(CommonValidatorContractUpdated)
	if err := _Common.contract.UnpackLog(event, "ValidatorContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommonVoteRejectedIterator is returned from FilterVoteRejected and is used to iterate over the raw logs and unpacked data for VoteRejected events raised by the Common contract.
type CommonVoteRejectedIterator struct {
	Event *CommonVoteRejected // Event containing the contract specifics and raw log

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
func (it *CommonVoteRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommonVoteRejected)
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
		it.Event = new(CommonVoteRejected)
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
func (it *CommonVoteRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommonVoteRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommonVoteRejected represents a VoteRejected event raised by the Common contract.
type CommonVoteRejected struct {
	Round *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVoteRejected is a free log retrieval operation binding the contract event 0xfb95c258b2a6900dbb6e469bc9d1162239d1f3df76bbf1c01b0524123b2e5740.
//
// Solidity: event VoteRejected(uint256 indexed round)
func (_Common *CommonFilterer) FilterVoteRejected(opts *bind.FilterOpts, round []*big.Int) (*CommonVoteRejectedIterator, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _Common.contract.FilterLogs(opts, "VoteRejected", roundRule)
	if err != nil {
		return nil, err
	}
	return &CommonVoteRejectedIterator{contract: _Common.contract, event: "VoteRejected", logs: logs, sub: sub}, nil
}

// WatchVoteRejected is a free log subscription operation binding the contract event 0xfb95c258b2a6900dbb6e469bc9d1162239d1f3df76bbf1c01b0524123b2e5740.
//
// Solidity: event VoteRejected(uint256 indexed round)
func (_Common *CommonFilterer) WatchVoteRejected(opts *bind.WatchOpts, sink chan<- *CommonVoteRejected, round []*big.Int) (event.Subscription, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _Common.contract.WatchLogs(opts, "VoteRejected", roundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommonVoteRejected)
				if err := _Common.contract.UnpackLog(event, "VoteRejected", log); err != nil {
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

// ParseVoteRejected is a log parse operation binding the contract event 0xfb95c258b2a6900dbb6e469bc9d1162239d1f3df76bbf1c01b0524123b2e5740.
//
// Solidity: event VoteRejected(uint256 indexed round)
func (_Common *CommonFilterer) ParseVoteRejected(log types.Log) (*CommonVoteRejected, error) {
	event := new(CommonVoteRejected)
	if err := _Common.contract.UnpackLog(event, "VoteRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
