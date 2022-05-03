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
	Nonce         *big.Int
	TargetOptions []uint8
	Values        []*big.Int
	Calldatas     [][]byte
}

// ProposalProposalDetail is an auto generated low-level Go binding around an user-defined struct.
type ProposalProposalDetail struct {
	Nonce     *big.Int
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gatewayContract\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_relayers\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"GatewayContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"}],\"indexed\":false,\"internalType\":\"structProposal.ProposalDetail\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"enumBridgeProposal.TargetOption[]\",\"name\":\"targetOptions\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"}],\"indexed\":false,\"internalType\":\"structBridgeProposal.GlobalProposal\",\"name\":\"\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"GlobalProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"}],\"indexed\":false,\"internalType\":\"structProposal.ProposalDetail\",\"name\":\"\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumBallot.VoteType\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"ProposalVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ValidatorContractUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"enumBridgeProposal.TargetOption[]\",\"name\":\"targetOptions\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"}],\"internalType\":\"structBridgeProposal.GlobalProposal\",\"name\":\"_globalProposal\",\"type\":\"tuple\"},{\"internalType\":\"enumBallot.VoteType[]\",\"name\":\"_supports\",\"type\":\"uint8[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"name\":\"castGlobalProposalBySignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"}],\"internalType\":\"structProposal.ProposalDetail\",\"name\":\"_proposal\",\"type\":\"tuple\"},{\"internalType\":\"enumBallot.VoteType[]\",\"name\":\"_supports\",\"type\":\"uint8[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"name\":\"castProposalBySignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeProxyAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gatewayContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"getProxyAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"getProxyImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVotingSignature\",\"outputs\":[{\"internalType\":\"enumBallot.VoteType\",\"name\":\"_support\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature\",\"name\":\"_sig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_calldatas\",\"type\":\"bytes[]\"}],\"name\":\"propose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumBridgeProposal.TargetOption[]\",\"name\":\"_targetOptions\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_calldatas\",\"type\":\"bytes[]\"}],\"name\":\"proposeGlobal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"enumBridgeProposal.TargetOption[]\",\"name\":\"targetOptions\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"}],\"internalType\":\"structBridgeProposal.GlobalProposal\",\"name\":\"_globalProposal\",\"type\":\"tuple\"},{\"internalType\":\"enumBallot.VoteType[]\",\"name\":\"_supports\",\"type\":\"uint8[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"name\":\"proposeGlobalProposalStructAndCastVotes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"}],\"internalType\":\"structProposal.ProposalDetail\",\"name\":\"_proposal\",\"type\":\"tuple\"},{\"internalType\":\"enumBallot.VoteType[]\",\"name\":\"_supports\",\"type\":\"uint8[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"name\":\"proposeProposalStructAndCastVotes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"enumBridgeProposal.TargetOption[]\",\"name\":\"targetOptions\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"}],\"internalType\":\"structBridgeProposal.GlobalProposal\",\"name\":\"_globalProposal\",\"type\":\"tuple\"},{\"internalType\":\"enumBallot.VoteType[]\",\"name\":\"_supports\",\"type\":\"uint8[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"name\":\"relayGlobalProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"}],\"internalType\":\"structProposal.ProposalDetail\",\"name\":\"_proposal\",\"type\":\"tuple\"},{\"internalType\":\"enumBallot.VoteType[]\",\"name\":\"_supports\",\"type\":\"uint8[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"name\":\"relayProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"round\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gatewayContract\",\"type\":\"address\"}],\"name\":\"setGatewayContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorContract\",\"type\":\"address\"}],\"name\":\"setValidatorContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[{\"internalType\":\"enumGovernance.VoteStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"againstVoteWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forVoteWeight\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// GetVotingSignature is a free data retrieval call binding the contract method 0x2e3fab41.
//
// Solidity: function getVotingSignature(uint256 _chainId, uint256 _round, address _voter) view returns(uint8 _support, (uint8,bytes32,bytes32) _sig)
func (_Common *CommonCaller) GetVotingSignature(opts *bind.CallOpts, _chainId *big.Int, _round *big.Int, _voter common.Address) (struct {
	Support uint8
	Sig     SignatureConsumerSignature
}, error) {
	var out []interface{}
	err := _Common.contract.Call(opts, &out, "getVotingSignature", _chainId, _round, _voter)

	outstruct := new(struct {
		Support uint8
		Sig     SignatureConsumerSignature
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Support = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Sig = *abi.ConvertType(out[1], new(SignatureConsumerSignature)).(*SignatureConsumerSignature)

	return *outstruct, err

}

// GetVotingSignature is a free data retrieval call binding the contract method 0x2e3fab41.
//
// Solidity: function getVotingSignature(uint256 _chainId, uint256 _round, address _voter) view returns(uint8 _support, (uint8,bytes32,bytes32) _sig)
func (_Common *CommonSession) GetVotingSignature(_chainId *big.Int, _round *big.Int, _voter common.Address) (struct {
	Support uint8
	Sig     SignatureConsumerSignature
}, error) {
	return _Common.Contract.GetVotingSignature(&_Common.CallOpts, _chainId, _round, _voter)
}

// GetVotingSignature is a free data retrieval call binding the contract method 0x2e3fab41.
//
// Solidity: function getVotingSignature(uint256 _chainId, uint256 _round, address _voter) view returns(uint8 _support, (uint8,bytes32,bytes32) _sig)
func (_Common *CommonCallerSession) GetVotingSignature(_chainId *big.Int, _round *big.Int, _voter common.Address) (struct {
	Support uint8
	Sig     SignatureConsumerSignature
}, error) {
	return _Common.Contract.GetVotingSignature(&_Common.CallOpts, _chainId, _round, _voter)
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

// Round is a free data retrieval call binding the contract method 0x34d5f37b.
//
// Solidity: function round(uint256 ) view returns(uint256)
func (_Common *CommonCaller) Round(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Common.contract.Call(opts, &out, "round", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Round is a free data retrieval call binding the contract method 0x34d5f37b.
//
// Solidity: function round(uint256 ) view returns(uint256)
func (_Common *CommonSession) Round(arg0 *big.Int) (*big.Int, error) {
	return _Common.Contract.Round(&_Common.CallOpts, arg0)
}

// Round is a free data retrieval call binding the contract method 0x34d5f37b.
//
// Solidity: function round(uint256 ) view returns(uint256)
func (_Common *CommonCallerSession) Round(arg0 *big.Int) (*big.Int, error) {
	return _Common.Contract.Round(&_Common.CallOpts, arg0)
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

// Vote is a free data retrieval call binding the contract method 0xb384abef.
//
// Solidity: function vote(uint256 , uint256 ) view returns(uint8 status, bytes32 hash, uint256 againstVoteWeight, uint256 forVoteWeight)
func (_Common *CommonCaller) Vote(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Status            uint8
	Hash              [32]byte
	AgainstVoteWeight *big.Int
	ForVoteWeight     *big.Int
}, error) {
	var out []interface{}
	err := _Common.contract.Call(opts, &out, "vote", arg0, arg1)

	outstruct := new(struct {
		Status            uint8
		Hash              [32]byte
		AgainstVoteWeight *big.Int
		ForVoteWeight     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Hash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.AgainstVoteWeight = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ForVoteWeight = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Vote is a free data retrieval call binding the contract method 0xb384abef.
//
// Solidity: function vote(uint256 , uint256 ) view returns(uint8 status, bytes32 hash, uint256 againstVoteWeight, uint256 forVoteWeight)
func (_Common *CommonSession) Vote(arg0 *big.Int, arg1 *big.Int) (struct {
	Status            uint8
	Hash              [32]byte
	AgainstVoteWeight *big.Int
	ForVoteWeight     *big.Int
}, error) {
	return _Common.Contract.Vote(&_Common.CallOpts, arg0, arg1)
}

// Vote is a free data retrieval call binding the contract method 0xb384abef.
//
// Solidity: function vote(uint256 , uint256 ) view returns(uint8 status, bytes32 hash, uint256 againstVoteWeight, uint256 forVoteWeight)
func (_Common *CommonCallerSession) Vote(arg0 *big.Int, arg1 *big.Int) (struct {
	Status            uint8
	Hash              [32]byte
	AgainstVoteWeight *big.Int
	ForVoteWeight     *big.Int
}, error) {
	return _Common.Contract.Vote(&_Common.CallOpts, arg0, arg1)
}

// CastGlobalProposalBySignatures is a paid mutator transaction binding the contract method 0x647ecb50.
//
// Solidity: function castGlobalProposalBySignatures((uint256,uint8[],uint256[],bytes[]) _globalProposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Common *CommonTransactor) CastGlobalProposalBySignatures(opts *bind.TransactOpts, _globalProposal BridgeProposalGlobalProposal, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "castGlobalProposalBySignatures", _globalProposal, _supports, _signatures)
}

// CastGlobalProposalBySignatures is a paid mutator transaction binding the contract method 0x647ecb50.
//
// Solidity: function castGlobalProposalBySignatures((uint256,uint8[],uint256[],bytes[]) _globalProposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Common *CommonSession) CastGlobalProposalBySignatures(_globalProposal BridgeProposalGlobalProposal, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Common.Contract.CastGlobalProposalBySignatures(&_Common.TransactOpts, _globalProposal, _supports, _signatures)
}

// CastGlobalProposalBySignatures is a paid mutator transaction binding the contract method 0x647ecb50.
//
// Solidity: function castGlobalProposalBySignatures((uint256,uint8[],uint256[],bytes[]) _globalProposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Common *CommonTransactorSession) CastGlobalProposalBySignatures(_globalProposal BridgeProposalGlobalProposal, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Common.Contract.CastGlobalProposalBySignatures(&_Common.TransactOpts, _globalProposal, _supports, _signatures)
}

// CastProposalBySignatures is a paid mutator transaction binding the contract method 0xeb72d5f4.
//
// Solidity: function castProposalBySignatures((uint256,uint256,address[],uint256[],bytes[]) _proposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Common *CommonTransactor) CastProposalBySignatures(opts *bind.TransactOpts, _proposal ProposalProposalDetail, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "castProposalBySignatures", _proposal, _supports, _signatures)
}

// CastProposalBySignatures is a paid mutator transaction binding the contract method 0xeb72d5f4.
//
// Solidity: function castProposalBySignatures((uint256,uint256,address[],uint256[],bytes[]) _proposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Common *CommonSession) CastProposalBySignatures(_proposal ProposalProposalDetail, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Common.Contract.CastProposalBySignatures(&_Common.TransactOpts, _proposal, _supports, _signatures)
}

// CastProposalBySignatures is a paid mutator transaction binding the contract method 0xeb72d5f4.
//
// Solidity: function castProposalBySignatures((uint256,uint256,address[],uint256[],bytes[]) _proposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Common *CommonTransactorSession) CastProposalBySignatures(_proposal ProposalProposalDetail, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Common.Contract.CastProposalBySignatures(&_Common.TransactOpts, _proposal, _supports, _signatures)
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

// Propose is a paid mutator transaction binding the contract method 0x5f9df13f.
//
// Solidity: function propose(uint256 _chainId, address[] _targets, uint256[] _values, bytes[] _calldatas) returns()
func (_Common *CommonTransactor) Propose(opts *bind.TransactOpts, _chainId *big.Int, _targets []common.Address, _values []*big.Int, _calldatas [][]byte) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "propose", _chainId, _targets, _values, _calldatas)
}

// Propose is a paid mutator transaction binding the contract method 0x5f9df13f.
//
// Solidity: function propose(uint256 _chainId, address[] _targets, uint256[] _values, bytes[] _calldatas) returns()
func (_Common *CommonSession) Propose(_chainId *big.Int, _targets []common.Address, _values []*big.Int, _calldatas [][]byte) (*types.Transaction, error) {
	return _Common.Contract.Propose(&_Common.TransactOpts, _chainId, _targets, _values, _calldatas)
}

// Propose is a paid mutator transaction binding the contract method 0x5f9df13f.
//
// Solidity: function propose(uint256 _chainId, address[] _targets, uint256[] _values, bytes[] _calldatas) returns()
func (_Common *CommonTransactorSession) Propose(_chainId *big.Int, _targets []common.Address, _values []*big.Int, _calldatas [][]byte) (*types.Transaction, error) {
	return _Common.Contract.Propose(&_Common.TransactOpts, _chainId, _targets, _values, _calldatas)
}

// ProposeGlobal is a paid mutator transaction binding the contract method 0x92f283c0.
//
// Solidity: function proposeGlobal(uint8[] _targetOptions, uint256[] _values, bytes[] _calldatas) returns()
func (_Common *CommonTransactor) ProposeGlobal(opts *bind.TransactOpts, _targetOptions []uint8, _values []*big.Int, _calldatas [][]byte) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "proposeGlobal", _targetOptions, _values, _calldatas)
}

// ProposeGlobal is a paid mutator transaction binding the contract method 0x92f283c0.
//
// Solidity: function proposeGlobal(uint8[] _targetOptions, uint256[] _values, bytes[] _calldatas) returns()
func (_Common *CommonSession) ProposeGlobal(_targetOptions []uint8, _values []*big.Int, _calldatas [][]byte) (*types.Transaction, error) {
	return _Common.Contract.ProposeGlobal(&_Common.TransactOpts, _targetOptions, _values, _calldatas)
}

// ProposeGlobal is a paid mutator transaction binding the contract method 0x92f283c0.
//
// Solidity: function proposeGlobal(uint8[] _targetOptions, uint256[] _values, bytes[] _calldatas) returns()
func (_Common *CommonTransactorSession) ProposeGlobal(_targetOptions []uint8, _values []*big.Int, _calldatas [][]byte) (*types.Transaction, error) {
	return _Common.Contract.ProposeGlobal(&_Common.TransactOpts, _targetOptions, _values, _calldatas)
}

// ProposeGlobalProposalStructAndCastVotes is a paid mutator transaction binding the contract method 0x5f844b98.
//
// Solidity: function proposeGlobalProposalStructAndCastVotes((uint256,uint8[],uint256[],bytes[]) _globalProposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Common *CommonTransactor) ProposeGlobalProposalStructAndCastVotes(opts *bind.TransactOpts, _globalProposal BridgeProposalGlobalProposal, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "proposeGlobalProposalStructAndCastVotes", _globalProposal, _supports, _signatures)
}

// ProposeGlobalProposalStructAndCastVotes is a paid mutator transaction binding the contract method 0x5f844b98.
//
// Solidity: function proposeGlobalProposalStructAndCastVotes((uint256,uint8[],uint256[],bytes[]) _globalProposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Common *CommonSession) ProposeGlobalProposalStructAndCastVotes(_globalProposal BridgeProposalGlobalProposal, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Common.Contract.ProposeGlobalProposalStructAndCastVotes(&_Common.TransactOpts, _globalProposal, _supports, _signatures)
}

// ProposeGlobalProposalStructAndCastVotes is a paid mutator transaction binding the contract method 0x5f844b98.
//
// Solidity: function proposeGlobalProposalStructAndCastVotes((uint256,uint8[],uint256[],bytes[]) _globalProposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Common *CommonTransactorSession) ProposeGlobalProposalStructAndCastVotes(_globalProposal BridgeProposalGlobalProposal, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Common.Contract.ProposeGlobalProposalStructAndCastVotes(&_Common.TransactOpts, _globalProposal, _supports, _signatures)
}

// ProposeProposalStructAndCastVotes is a paid mutator transaction binding the contract method 0x6a35891b.
//
// Solidity: function proposeProposalStructAndCastVotes((uint256,uint256,address[],uint256[],bytes[]) _proposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Common *CommonTransactor) ProposeProposalStructAndCastVotes(opts *bind.TransactOpts, _proposal ProposalProposalDetail, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "proposeProposalStructAndCastVotes", _proposal, _supports, _signatures)
}

// ProposeProposalStructAndCastVotes is a paid mutator transaction binding the contract method 0x6a35891b.
//
// Solidity: function proposeProposalStructAndCastVotes((uint256,uint256,address[],uint256[],bytes[]) _proposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Common *CommonSession) ProposeProposalStructAndCastVotes(_proposal ProposalProposalDetail, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Common.Contract.ProposeProposalStructAndCastVotes(&_Common.TransactOpts, _proposal, _supports, _signatures)
}

// ProposeProposalStructAndCastVotes is a paid mutator transaction binding the contract method 0x6a35891b.
//
// Solidity: function proposeProposalStructAndCastVotes((uint256,uint256,address[],uint256[],bytes[]) _proposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Common *CommonTransactorSession) ProposeProposalStructAndCastVotes(_proposal ProposalProposalDetail, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Common.Contract.ProposeProposalStructAndCastVotes(&_Common.TransactOpts, _proposal, _supports, _signatures)
}

// RelayGlobalProposal is a paid mutator transaction binding the contract method 0xd0b6cd04.
//
// Solidity: function relayGlobalProposal((uint256,uint8[],uint256[],bytes[]) _globalProposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Common *CommonTransactor) RelayGlobalProposal(opts *bind.TransactOpts, _globalProposal BridgeProposalGlobalProposal, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "relayGlobalProposal", _globalProposal, _supports, _signatures)
}

// RelayGlobalProposal is a paid mutator transaction binding the contract method 0xd0b6cd04.
//
// Solidity: function relayGlobalProposal((uint256,uint8[],uint256[],bytes[]) _globalProposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Common *CommonSession) RelayGlobalProposal(_globalProposal BridgeProposalGlobalProposal, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Common.Contract.RelayGlobalProposal(&_Common.TransactOpts, _globalProposal, _supports, _signatures)
}

// RelayGlobalProposal is a paid mutator transaction binding the contract method 0xd0b6cd04.
//
// Solidity: function relayGlobalProposal((uint256,uint8[],uint256[],bytes[]) _globalProposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Common *CommonTransactorSession) RelayGlobalProposal(_globalProposal BridgeProposalGlobalProposal, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Common.Contract.RelayGlobalProposal(&_Common.TransactOpts, _globalProposal, _supports, _signatures)
}

// RelayProposal is a paid mutator transaction binding the contract method 0xeb45b978.
//
// Solidity: function relayProposal((uint256,uint256,address[],uint256[],bytes[]) _proposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Common *CommonTransactor) RelayProposal(opts *bind.TransactOpts, _proposal ProposalProposalDetail, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Common.contract.Transact(opts, "relayProposal", _proposal, _supports, _signatures)
}

// RelayProposal is a paid mutator transaction binding the contract method 0xeb45b978.
//
// Solidity: function relayProposal((uint256,uint256,address[],uint256[],bytes[]) _proposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Common *CommonSession) RelayProposal(_proposal ProposalProposalDetail, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Common.Contract.RelayProposal(&_Common.TransactOpts, _proposal, _supports, _signatures)
}

// RelayProposal is a paid mutator transaction binding the contract method 0xeb45b978.
//
// Solidity: function relayProposal((uint256,uint256,address[],uint256[],bytes[]) _proposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Common *CommonTransactorSession) RelayProposal(_proposal ProposalProposalDetail, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Common.Contract.RelayProposal(&_Common.TransactOpts, _proposal, _supports, _signatures)
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

// CommonGlobalProposalCreatedIterator is returned from FilterGlobalProposalCreated and is used to iterate over the raw logs and unpacked data for GlobalProposalCreated events raised by the Common contract.
type CommonGlobalProposalCreatedIterator struct {
	Event *CommonGlobalProposalCreated // Event containing the contract specifics and raw log

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
func (it *CommonGlobalProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommonGlobalProposalCreated)
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
		it.Event = new(CommonGlobalProposalCreated)
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
func (it *CommonGlobalProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommonGlobalProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommonGlobalProposalCreated represents a GlobalProposalCreated event raised by the Common contract.
type CommonGlobalProposalCreated struct {
	Round        *big.Int
	ProposalHash [32]byte
	Arg2         ProposalProposalDetail
	Arg3         BridgeProposalGlobalProposal
	Creator      common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGlobalProposalCreated is a free log retrieval operation binding the contract event 0x82ee0ad0758ee805d42e19ed61c32161ac3890848c4711fbe4f35b5081607217.
//
// Solidity: event GlobalProposalCreated(uint256 indexed round, bytes32 indexed proposalHash, (uint256,uint256,address[],uint256[],bytes[]) arg2, (uint256,uint8[],uint256[],bytes[]) arg3, address creator)
func (_Common *CommonFilterer) FilterGlobalProposalCreated(opts *bind.FilterOpts, round []*big.Int, proposalHash [][32]byte) (*CommonGlobalProposalCreatedIterator, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _Common.contract.FilterLogs(opts, "GlobalProposalCreated", roundRule, proposalHashRule)
	if err != nil {
		return nil, err
	}
	return &CommonGlobalProposalCreatedIterator{contract: _Common.contract, event: "GlobalProposalCreated", logs: logs, sub: sub}, nil
}

// WatchGlobalProposalCreated is a free log subscription operation binding the contract event 0x82ee0ad0758ee805d42e19ed61c32161ac3890848c4711fbe4f35b5081607217.
//
// Solidity: event GlobalProposalCreated(uint256 indexed round, bytes32 indexed proposalHash, (uint256,uint256,address[],uint256[],bytes[]) arg2, (uint256,uint8[],uint256[],bytes[]) arg3, address creator)
func (_Common *CommonFilterer) WatchGlobalProposalCreated(opts *bind.WatchOpts, sink chan<- *CommonGlobalProposalCreated, round []*big.Int, proposalHash [][32]byte) (event.Subscription, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _Common.contract.WatchLogs(opts, "GlobalProposalCreated", roundRule, proposalHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommonGlobalProposalCreated)
				if err := _Common.contract.UnpackLog(event, "GlobalProposalCreated", log); err != nil {
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

// ParseGlobalProposalCreated is a log parse operation binding the contract event 0x82ee0ad0758ee805d42e19ed61c32161ac3890848c4711fbe4f35b5081607217.
//
// Solidity: event GlobalProposalCreated(uint256 indexed round, bytes32 indexed proposalHash, (uint256,uint256,address[],uint256[],bytes[]) arg2, (uint256,uint8[],uint256[],bytes[]) arg3, address creator)
func (_Common *CommonFilterer) ParseGlobalProposalCreated(log types.Log) (*CommonGlobalProposalCreated, error) {
	event := new(CommonGlobalProposalCreated)
	if err := _Common.contract.UnpackLog(event, "GlobalProposalCreated", log); err != nil {
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
	ProposalHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProposalApproved is a free log retrieval operation binding the contract event 0x5c819725ea53655a3b898f3df59b66489761935454e9212ca1e5ebd759953d0b.
//
// Solidity: event ProposalApproved(bytes32 indexed proposalHash)
func (_Common *CommonFilterer) FilterProposalApproved(opts *bind.FilterOpts, proposalHash [][32]byte) (*CommonProposalApprovedIterator, error) {

	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _Common.contract.FilterLogs(opts, "ProposalApproved", proposalHashRule)
	if err != nil {
		return nil, err
	}
	return &CommonProposalApprovedIterator{contract: _Common.contract, event: "ProposalApproved", logs: logs, sub: sub}, nil
}

// WatchProposalApproved is a free log subscription operation binding the contract event 0x5c819725ea53655a3b898f3df59b66489761935454e9212ca1e5ebd759953d0b.
//
// Solidity: event ProposalApproved(bytes32 indexed proposalHash)
func (_Common *CommonFilterer) WatchProposalApproved(opts *bind.WatchOpts, sink chan<- *CommonProposalApproved, proposalHash [][32]byte) (event.Subscription, error) {

	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _Common.contract.WatchLogs(opts, "ProposalApproved", proposalHashRule)
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

// ParseProposalApproved is a log parse operation binding the contract event 0x5c819725ea53655a3b898f3df59b66489761935454e9212ca1e5ebd759953d0b.
//
// Solidity: event ProposalApproved(bytes32 indexed proposalHash)
func (_Common *CommonFilterer) ParseProposalApproved(log types.Log) (*CommonProposalApproved, error) {
	event := new(CommonProposalApproved)
	if err := _Common.contract.UnpackLog(event, "ProposalApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommonProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the Common contract.
type CommonProposalCreatedIterator struct {
	Event *CommonProposalCreated // Event containing the contract specifics and raw log

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
func (it *CommonProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommonProposalCreated)
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
		it.Event = new(CommonProposalCreated)
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
func (it *CommonProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommonProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommonProposalCreated represents a ProposalCreated event raised by the Common contract.
type CommonProposalCreated struct {
	ChainId      *big.Int
	Round        *big.Int
	ProposalHash [32]byte
	Arg3         ProposalProposalDetail
	Creator      common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0x3d53769dd1253e37ceefb20fe16fbc7ff25d98e2d0f8c4730236e18500ca9b8c.
//
// Solidity: event ProposalCreated(uint256 indexed chainId, uint256 indexed round, bytes32 indexed proposalHash, (uint256,uint256,address[],uint256[],bytes[]) arg3, address creator)
func (_Common *CommonFilterer) FilterProposalCreated(opts *bind.FilterOpts, chainId []*big.Int, round []*big.Int, proposalHash [][32]byte) (*CommonProposalCreatedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _Common.contract.FilterLogs(opts, "ProposalCreated", chainIdRule, roundRule, proposalHashRule)
	if err != nil {
		return nil, err
	}
	return &CommonProposalCreatedIterator{contract: _Common.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0x3d53769dd1253e37ceefb20fe16fbc7ff25d98e2d0f8c4730236e18500ca9b8c.
//
// Solidity: event ProposalCreated(uint256 indexed chainId, uint256 indexed round, bytes32 indexed proposalHash, (uint256,uint256,address[],uint256[],bytes[]) arg3, address creator)
func (_Common *CommonFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *CommonProposalCreated, chainId []*big.Int, round []*big.Int, proposalHash [][32]byte) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _Common.contract.WatchLogs(opts, "ProposalCreated", chainIdRule, roundRule, proposalHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommonProposalCreated)
				if err := _Common.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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

// ParseProposalCreated is a log parse operation binding the contract event 0x3d53769dd1253e37ceefb20fe16fbc7ff25d98e2d0f8c4730236e18500ca9b8c.
//
// Solidity: event ProposalCreated(uint256 indexed chainId, uint256 indexed round, bytes32 indexed proposalHash, (uint256,uint256,address[],uint256[],bytes[]) arg3, address creator)
func (_Common *CommonFilterer) ParseProposalCreated(log types.Log) (*CommonProposalCreated, error) {
	event := new(CommonProposalCreated)
	if err := _Common.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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
	ProposalHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x7b1bcf1ccf901a11589afff5504d59fd0a53780eed2a952adade0348985139e0.
//
// Solidity: event ProposalExecuted(bytes32 indexed proposalHash)
func (_Common *CommonFilterer) FilterProposalExecuted(opts *bind.FilterOpts, proposalHash [][32]byte) (*CommonProposalExecutedIterator, error) {

	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _Common.contract.FilterLogs(opts, "ProposalExecuted", proposalHashRule)
	if err != nil {
		return nil, err
	}
	return &CommonProposalExecutedIterator{contract: _Common.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x7b1bcf1ccf901a11589afff5504d59fd0a53780eed2a952adade0348985139e0.
//
// Solidity: event ProposalExecuted(bytes32 indexed proposalHash)
func (_Common *CommonFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *CommonProposalExecuted, proposalHash [][32]byte) (event.Subscription, error) {

	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _Common.contract.WatchLogs(opts, "ProposalExecuted", proposalHashRule)
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

// ParseProposalExecuted is a log parse operation binding the contract event 0x7b1bcf1ccf901a11589afff5504d59fd0a53780eed2a952adade0348985139e0.
//
// Solidity: event ProposalExecuted(bytes32 indexed proposalHash)
func (_Common *CommonFilterer) ParseProposalExecuted(log types.Log) (*CommonProposalExecuted, error) {
	event := new(CommonProposalExecuted)
	if err := _Common.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommonProposalRejectedIterator is returned from FilterProposalRejected and is used to iterate over the raw logs and unpacked data for ProposalRejected events raised by the Common contract.
type CommonProposalRejectedIterator struct {
	Event *CommonProposalRejected // Event containing the contract specifics and raw log

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
func (it *CommonProposalRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommonProposalRejected)
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
		it.Event = new(CommonProposalRejected)
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
func (it *CommonProposalRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommonProposalRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommonProposalRejected represents a ProposalRejected event raised by the Common contract.
type CommonProposalRejected struct {
	ProposalHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProposalRejected is a free log retrieval operation binding the contract event 0x55295d4ce992922fa2e5ffbf3a3dcdb367de0a15e125ace083456017fd22060f.
//
// Solidity: event ProposalRejected(bytes32 indexed proposalHash)
func (_Common *CommonFilterer) FilterProposalRejected(opts *bind.FilterOpts, proposalHash [][32]byte) (*CommonProposalRejectedIterator, error) {

	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _Common.contract.FilterLogs(opts, "ProposalRejected", proposalHashRule)
	if err != nil {
		return nil, err
	}
	return &CommonProposalRejectedIterator{contract: _Common.contract, event: "ProposalRejected", logs: logs, sub: sub}, nil
}

// WatchProposalRejected is a free log subscription operation binding the contract event 0x55295d4ce992922fa2e5ffbf3a3dcdb367de0a15e125ace083456017fd22060f.
//
// Solidity: event ProposalRejected(bytes32 indexed proposalHash)
func (_Common *CommonFilterer) WatchProposalRejected(opts *bind.WatchOpts, sink chan<- *CommonProposalRejected, proposalHash [][32]byte) (event.Subscription, error) {

	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _Common.contract.WatchLogs(opts, "ProposalRejected", proposalHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommonProposalRejected)
				if err := _Common.contract.UnpackLog(event, "ProposalRejected", log); err != nil {
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

// ParseProposalRejected is a log parse operation binding the contract event 0x55295d4ce992922fa2e5ffbf3a3dcdb367de0a15e125ace083456017fd22060f.
//
// Solidity: event ProposalRejected(bytes32 indexed proposalHash)
func (_Common *CommonFilterer) ParseProposalRejected(log types.Log) (*CommonProposalRejected, error) {
	event := new(CommonProposalRejected)
	if err := _Common.contract.UnpackLog(event, "ProposalRejected", log); err != nil {
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
	ProposalHash [32]byte
	Voter        common.Address
	Support      uint8
	Weight       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProposalVoted is a free log retrieval operation binding the contract event 0x1203f9e81c814a35f5f4cc24087b2a24c6fb7986a9f1406b68a9484882c93a23.
//
// Solidity: event ProposalVoted(bytes32 indexed proposalHash, address indexed voter, uint8 support, uint256 weight)
func (_Common *CommonFilterer) FilterProposalVoted(opts *bind.FilterOpts, proposalHash [][32]byte, voter []common.Address) (*CommonProposalVotedIterator, error) {

	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Common.contract.FilterLogs(opts, "ProposalVoted", proposalHashRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &CommonProposalVotedIterator{contract: _Common.contract, event: "ProposalVoted", logs: logs, sub: sub}, nil
}

// WatchProposalVoted is a free log subscription operation binding the contract event 0x1203f9e81c814a35f5f4cc24087b2a24c6fb7986a9f1406b68a9484882c93a23.
//
// Solidity: event ProposalVoted(bytes32 indexed proposalHash, address indexed voter, uint8 support, uint256 weight)
func (_Common *CommonFilterer) WatchProposalVoted(opts *bind.WatchOpts, sink chan<- *CommonProposalVoted, proposalHash [][32]byte, voter []common.Address) (event.Subscription, error) {

	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Common.contract.WatchLogs(opts, "ProposalVoted", proposalHashRule, voterRule)
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

// ParseProposalVoted is a log parse operation binding the contract event 0x1203f9e81c814a35f5f4cc24087b2a24c6fb7986a9f1406b68a9484882c93a23.
//
// Solidity: event ProposalVoted(bytes32 indexed proposalHash, address indexed voter, uint8 support, uint256 weight)
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
