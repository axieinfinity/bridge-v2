// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gateway

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

// TokenInfo is an auto generated low-level Go binding around an user-defined struct.
type TokenInfo struct {
	Erc      uint8
	Id       *big.Int
	Quantity *big.Int
}

// TokenOwner is an auto generated low-level Go binding around an user-defined struct.
type TokenOwner struct {
	Addr      common.Address
	TokenAddr common.Address
	ChainId   *big.Int
}

// TransferReceipt is an auto generated low-level Go binding around an user-defined struct.
type TransferReceipt struct {
	Id        *big.Int
	Kind      uint8
	Mainchain TokenOwner
	Ronin     TokenOwner
	Info      TokenInfo
}

// TransferRequest is an auto generated low-level Go binding around an user-defined struct.
type TransferRequest struct {
	RecipientAddr common.Address
	TokenAddr common.Address
	Info      TokenInfo
}

// RoninMetaData contains all meta data concerning the Ronin contract.
var RoninMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"receiptHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"threshold\",\"type\":\"uint256[]\"}],\"name\":\"MinimumThresholdsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_numerator\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_denominator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_previousNumerator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_previousDenominator\",\"type\":\"uint256\"}],\"name\":\"ThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"roninTokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"mainchainTokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"chainIds\",\"type\":\"uint256[]\"}],\"name\":\"TokenMapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIWeightedValidator\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ValidatorContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"receiptHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"WithdrawalRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"receiptHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"WithdrawalSignaturesRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWAL_MIGRATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structTransfer.Receipt[]\",\"name\":\"_receipts\",\"type\":\"tuple[]\"}],\"name\":\"bulkDepositFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipientAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structTransfer.Request[]\",\"name\":\"_requests\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"bulkRequestWithdrawalFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_withdrawals\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_signatures\",\"type\":\"bytes[]\"}],\"name\":\"bulkSubmitWithdrawalSignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_voteWeight\",\"type\":\"uint256\"}],\"name\":\"checkThreshold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structTransfer.Receipt\",\"name\":\"_receipt\",\"type\":\"tuple\"}],\"name\":\"depositFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"depositVote\",\"outputs\":[{\"internalType\":\"enumGovernorCore.VoteStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"finalHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_roninToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"getMainchainToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawalId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"}],\"name\":\"getWithdrawalSignatures\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_signatures\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_roleSetter\",\"type\":\"address\"},{\"internalType\":\"contractIWeightedValidator\",\"name\":\"_validatorContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_denominator\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_withdrawalMigrators\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_roninTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_mainchainTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][2]\",\"name\":\"_packedNumbers\",\"type\":\"uint256[][2]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_roninTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_mainchainTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_chainIds\",\"type\":\"uint256[]\"}],\"name\":\"mapTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"markWithdrawalMigrated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipientAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structTransfer.Request[]\",\"name\":\"_requests\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"_requesters\",\"type\":\"address[]\"}],\"name\":\"migrateWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minimumThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumVoteWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipientAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structTransfer.Request\",\"name\":\"_request\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"requestWithdrawalFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawalId\",\"type\":\"uint256\"}],\"name\":\"requestWithdrawalSignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_thresholds\",\"type\":\"uint256[]\"}],\"name\":\"setMinimumThresholds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_denominator\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIWeightedValidator\",\"name\":\"_validatorContract\",\"type\":\"address\"}],\"name\":\"setValidatorContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorContract\",\"outputs\":[{\"internalType\":\"contractIWeightedValidator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalMigrated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RoninABI is the input ABI used to generate the binding from.
// Deprecated: Use RoninMetaData.ABI instead.
var RoninABI = RoninMetaData.ABI

// Ronin is an auto generated Go binding around an Ethereum contract.
type Ronin struct {
	RoninCaller     // Read-only binding to the contract
	RoninTransactor // Write-only binding to the contract
	RoninFilterer   // Log filterer for contract events
}

// RoninCaller is an auto generated read-only Go binding around an Ethereum contract.
type RoninCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoninTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RoninTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoninFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoninFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoninSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoninSession struct {
	Contract     *Ronin            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoninCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoninCallerSession struct {
	Contract *RoninCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RoninTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoninTransactorSession struct {
	Contract     *RoninTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoninRaw is an auto generated low-level Go binding around an Ethereum contract.
type RoninRaw struct {
	Contract *Ronin // Generic contract binding to access the raw methods on
}

// RoninCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoninCallerRaw struct {
	Contract *RoninCaller // Generic read-only contract binding to access the raw methods on
}

// RoninTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoninTransactorRaw struct {
	Contract *RoninTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRonin creates a new instance of Ronin, bound to a specific deployed contract.
func NewRonin(address common.Address, backend bind.ContractBackend) (*Ronin, error) {
	contract, err := bindRonin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ronin{RoninCaller: RoninCaller{contract: contract}, RoninTransactor: RoninTransactor{contract: contract}, RoninFilterer: RoninFilterer{contract: contract}}, nil
}

// NewRoninCaller creates a new read-only instance of Ronin, bound to a specific deployed contract.
func NewRoninCaller(address common.Address, caller bind.ContractCaller) (*RoninCaller, error) {
	contract, err := bindRonin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoninCaller{contract: contract}, nil
}

// NewRoninTransactor creates a new write-only instance of Ronin, bound to a specific deployed contract.
func NewRoninTransactor(address common.Address, transactor bind.ContractTransactor) (*RoninTransactor, error) {
	contract, err := bindRonin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoninTransactor{contract: contract}, nil
}

// NewRoninFilterer creates a new log filterer instance of Ronin, bound to a specific deployed contract.
func NewRoninFilterer(address common.Address, filterer bind.ContractFilterer) (*RoninFilterer, error) {
	contract, err := bindRonin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoninFilterer{contract: contract}, nil
}

// bindRonin binds a generic wrapper to an already deployed contract.
func bindRonin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RoninABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ronin *RoninRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ronin.Contract.RoninCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ronin *RoninRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ronin.Contract.RoninTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ronin *RoninRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ronin.Contract.RoninTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ronin *RoninCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ronin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ronin *RoninTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ronin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ronin *RoninTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ronin.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Ronin *RoninCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ronin.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Ronin *RoninSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Ronin.Contract.DEFAULTADMINROLE(&_Ronin.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Ronin *RoninCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Ronin.Contract.DEFAULTADMINROLE(&_Ronin.CallOpts)
}

// WITHDRAWALMIGRATOR is a free data retrieval call binding the contract method 0xfe90d9c2.
//
// Solidity: function WITHDRAWAL_MIGRATOR() view returns(bytes32)
func (_Ronin *RoninCaller) WITHDRAWALMIGRATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ronin.contract.Call(opts, &out, "WITHDRAWAL_MIGRATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWALMIGRATOR is a free data retrieval call binding the contract method 0xfe90d9c2.
//
// Solidity: function WITHDRAWAL_MIGRATOR() view returns(bytes32)
func (_Ronin *RoninSession) WITHDRAWALMIGRATOR() ([32]byte, error) {
	return _Ronin.Contract.WITHDRAWALMIGRATOR(&_Ronin.CallOpts)
}

// WITHDRAWALMIGRATOR is a free data retrieval call binding the contract method 0xfe90d9c2.
//
// Solidity: function WITHDRAWAL_MIGRATOR() view returns(bytes32)
func (_Ronin *RoninCallerSession) WITHDRAWALMIGRATOR() ([32]byte, error) {
	return _Ronin.Contract.WITHDRAWALMIGRATOR(&_Ronin.CallOpts)
}

// CheckThreshold is a free data retrieval call binding the contract method 0xdafae408.
//
// Solidity: function checkThreshold(uint256 _voteWeight) view returns(bool)
func (_Ronin *RoninCaller) CheckThreshold(opts *bind.CallOpts, _voteWeight *big.Int) (bool, error) {
	var out []interface{}
	err := _Ronin.contract.Call(opts, &out, "checkThreshold", _voteWeight)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckThreshold is a free data retrieval call binding the contract method 0xdafae408.
//
// Solidity: function checkThreshold(uint256 _voteWeight) view returns(bool)
func (_Ronin *RoninSession) CheckThreshold(_voteWeight *big.Int) (bool, error) {
	return _Ronin.Contract.CheckThreshold(&_Ronin.CallOpts, _voteWeight)
}

// CheckThreshold is a free data retrieval call binding the contract method 0xdafae408.
//
// Solidity: function checkThreshold(uint256 _voteWeight) view returns(bool)
func (_Ronin *RoninCallerSession) CheckThreshold(_voteWeight *big.Int) (bool, error) {
	return _Ronin.Contract.CheckThreshold(&_Ronin.CallOpts, _voteWeight)
}

// DepositVote is a free data retrieval call binding the contract method 0x4d92c4f0.
//
// Solidity: function depositVote(uint256 , uint256 ) view returns(uint8 status, bytes32 finalHash)
func (_Ronin *RoninCaller) DepositVote(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Status    uint8
	FinalHash [32]byte
}, error) {
	var out []interface{}
	err := _Ronin.contract.Call(opts, &out, "depositVote", arg0, arg1)

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

// DepositVote is a free data retrieval call binding the contract method 0x4d92c4f0.
//
// Solidity: function depositVote(uint256 , uint256 ) view returns(uint8 status, bytes32 finalHash)
func (_Ronin *RoninSession) DepositVote(arg0 *big.Int, arg1 *big.Int) (struct {
	Status    uint8
	FinalHash [32]byte
}, error) {
	return _Ronin.Contract.DepositVote(&_Ronin.CallOpts, arg0, arg1)
}

// DepositVote is a free data retrieval call binding the contract method 0x4d92c4f0.
//
// Solidity: function depositVote(uint256 , uint256 ) view returns(uint8 status, bytes32 finalHash)
func (_Ronin *RoninCallerSession) DepositVote(arg0 *big.Int, arg1 *big.Int) (struct {
	Status    uint8
	FinalHash [32]byte
}, error) {
	return _Ronin.Contract.DepositVote(&_Ronin.CallOpts, arg0, arg1)
}

// GetMainchainToken is a free data retrieval call binding the contract method 0x5d6a9a90.
//
// Solidity: function getMainchainToken(address _roninToken, uint256 _chainId) view returns(address _addr)
func (_Ronin *RoninCaller) GetMainchainToken(opts *bind.CallOpts, _roninToken common.Address, _chainId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Ronin.contract.Call(opts, &out, "getMainchainToken", _roninToken, _chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMainchainToken is a free data retrieval call binding the contract method 0x5d6a9a90.
//
// Solidity: function getMainchainToken(address _roninToken, uint256 _chainId) view returns(address _addr)
func (_Ronin *RoninSession) GetMainchainToken(_roninToken common.Address, _chainId *big.Int) (common.Address, error) {
	return _Ronin.Contract.GetMainchainToken(&_Ronin.CallOpts, _roninToken, _chainId)
}

// GetMainchainToken is a free data retrieval call binding the contract method 0x5d6a9a90.
//
// Solidity: function getMainchainToken(address _roninToken, uint256 _chainId) view returns(address _addr)
func (_Ronin *RoninCallerSession) GetMainchainToken(_roninToken common.Address, _chainId *big.Int) (common.Address, error) {
	return _Ronin.Contract.GetMainchainToken(&_Ronin.CallOpts, _roninToken, _chainId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Ronin *RoninCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Ronin.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Ronin *RoninSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Ronin.Contract.GetRoleAdmin(&_Ronin.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Ronin *RoninCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Ronin.Contract.GetRoleAdmin(&_Ronin.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Ronin *RoninCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Ronin.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Ronin *RoninSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Ronin.Contract.GetRoleMember(&_Ronin.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Ronin *RoninCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Ronin.Contract.GetRoleMember(&_Ronin.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Ronin *RoninCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Ronin.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Ronin *RoninSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Ronin.Contract.GetRoleMemberCount(&_Ronin.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Ronin *RoninCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Ronin.Contract.GetRoleMemberCount(&_Ronin.CallOpts, role)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256, uint256)
func (_Ronin *RoninCaller) GetThreshold(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Ronin.contract.Call(opts, &out, "getThreshold")

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
func (_Ronin *RoninSession) GetThreshold() (*big.Int, *big.Int, error) {
	return _Ronin.Contract.GetThreshold(&_Ronin.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256, uint256)
func (_Ronin *RoninCallerSession) GetThreshold() (*big.Int, *big.Int, error) {
	return _Ronin.Contract.GetThreshold(&_Ronin.CallOpts)
}

// GetWithdrawalSignatures is a free data retrieval call binding the contract method 0xecc83649.
//
// Solidity: function getWithdrawalSignatures(uint256 _withdrawalId, address[] _validators) view returns(bytes[] _signatures)
func (_Ronin *RoninCaller) GetWithdrawalSignatures(opts *bind.CallOpts, _withdrawalId *big.Int, _validators []common.Address) ([][]byte, error) {
	var out []interface{}
	err := _Ronin.contract.Call(opts, &out, "getWithdrawalSignatures", _withdrawalId, _validators)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetWithdrawalSignatures is a free data retrieval call binding the contract method 0xecc83649.
//
// Solidity: function getWithdrawalSignatures(uint256 _withdrawalId, address[] _validators) view returns(bytes[] _signatures)
func (_Ronin *RoninSession) GetWithdrawalSignatures(_withdrawalId *big.Int, _validators []common.Address) ([][]byte, error) {
	return _Ronin.Contract.GetWithdrawalSignatures(&_Ronin.CallOpts, _withdrawalId, _validators)
}

// GetWithdrawalSignatures is a free data retrieval call binding the contract method 0xecc83649.
//
// Solidity: function getWithdrawalSignatures(uint256 _withdrawalId, address[] _validators) view returns(bytes[] _signatures)
func (_Ronin *RoninCallerSession) GetWithdrawalSignatures(_withdrawalId *big.Int, _validators []common.Address) ([][]byte, error) {
	return _Ronin.Contract.GetWithdrawalSignatures(&_Ronin.CallOpts, _withdrawalId, _validators)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Ronin *RoninCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Ronin.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Ronin *RoninSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Ronin.Contract.HasRole(&_Ronin.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Ronin *RoninCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Ronin.Contract.HasRole(&_Ronin.CallOpts, role, account)
}

// MinimumThreshold is a free data retrieval call binding the contract method 0xbc7f0386.
//
// Solidity: function minimumThreshold(address ) view returns(uint256)
func (_Ronin *RoninCaller) MinimumThreshold(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ronin.contract.Call(opts, &out, "minimumThreshold", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumThreshold is a free data retrieval call binding the contract method 0xbc7f0386.
//
// Solidity: function minimumThreshold(address ) view returns(uint256)
func (_Ronin *RoninSession) MinimumThreshold(arg0 common.Address) (*big.Int, error) {
	return _Ronin.Contract.MinimumThreshold(&_Ronin.CallOpts, arg0)
}

// MinimumThreshold is a free data retrieval call binding the contract method 0xbc7f0386.
//
// Solidity: function minimumThreshold(address ) view returns(uint256)
func (_Ronin *RoninCallerSession) MinimumThreshold(arg0 common.Address) (*big.Int, error) {
	return _Ronin.Contract.MinimumThreshold(&_Ronin.CallOpts, arg0)
}

// MinimumVoteWeight is a free data retrieval call binding the contract method 0x7de5dedd.
//
// Solidity: function minimumVoteWeight() view returns(uint256)
func (_Ronin *RoninCaller) MinimumVoteWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ronin.contract.Call(opts, &out, "minimumVoteWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumVoteWeight is a free data retrieval call binding the contract method 0x7de5dedd.
//
// Solidity: function minimumVoteWeight() view returns(uint256)
func (_Ronin *RoninSession) MinimumVoteWeight() (*big.Int, error) {
	return _Ronin.Contract.MinimumVoteWeight(&_Ronin.CallOpts)
}

// MinimumVoteWeight is a free data retrieval call binding the contract method 0x7de5dedd.
//
// Solidity: function minimumVoteWeight() view returns(uint256)
func (_Ronin *RoninCallerSession) MinimumVoteWeight() (*big.Int, error) {
	return _Ronin.Contract.MinimumVoteWeight(&_Ronin.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Ronin *RoninCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ronin.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Ronin *RoninSession) Nonce() (*big.Int, error) {
	return _Ronin.Contract.Nonce(&_Ronin.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Ronin *RoninCallerSession) Nonce() (*big.Int, error) {
	return _Ronin.Contract.Nonce(&_Ronin.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Ronin *RoninCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Ronin.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Ronin *RoninSession) Paused() (bool, error) {
	return _Ronin.Contract.Paused(&_Ronin.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Ronin *RoninCallerSession) Paused() (bool, error) {
	return _Ronin.Contract.Paused(&_Ronin.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ronin *RoninCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Ronin.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ronin *RoninSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Ronin.Contract.SupportsInterface(&_Ronin.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ronin *RoninCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Ronin.Contract.SupportsInterface(&_Ronin.CallOpts, interfaceId)
}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_Ronin *RoninCaller) ValidatorContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ronin.contract.Call(opts, &out, "validatorContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_Ronin *RoninSession) ValidatorContract() (common.Address, error) {
	return _Ronin.Contract.ValidatorContract(&_Ronin.CallOpts)
}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_Ronin *RoninCallerSession) ValidatorContract() (common.Address, error) {
	return _Ronin.Contract.ValidatorContract(&_Ronin.CallOpts)
}

// Withdrawal is a free data retrieval call binding the contract method 0x835fc6ca.
//
// Solidity: function withdrawal(uint256 ) view returns(uint256 id, uint8 kind, (address,address,uint256) mainchain, (address,address,uint256) ronin, (uint8,uint256,uint256) info)
func (_Ronin *RoninCaller) Withdrawal(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id        *big.Int
	Kind      uint8
	Mainchain TokenOwner
	Ronin     TokenOwner
	Info      TokenInfo
}, error) {
	var out []interface{}
	err := _Ronin.contract.Call(opts, &out, "withdrawal", arg0)

	outstruct := new(struct {
		Id        *big.Int
		Kind      uint8
		Mainchain TokenOwner
		Ronin     TokenOwner
		Info      TokenInfo
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Kind = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Mainchain = *abi.ConvertType(out[2], new(TokenOwner)).(*TokenOwner)
	outstruct.Ronin = *abi.ConvertType(out[3], new(TokenOwner)).(*TokenOwner)
	outstruct.Info = *abi.ConvertType(out[4], new(TokenInfo)).(*TokenInfo)

	return *outstruct, err

}

// Withdrawal is a free data retrieval call binding the contract method 0x835fc6ca.
//
// Solidity: function withdrawal(uint256 ) view returns(uint256 id, uint8 kind, (address,address,uint256) mainchain, (address,address,uint256) ronin, (uint8,uint256,uint256) info)
func (_Ronin *RoninSession) Withdrawal(arg0 *big.Int) (struct {
	Id        *big.Int
	Kind      uint8
	Mainchain TokenOwner
	Ronin     TokenOwner
	Info      TokenInfo
}, error) {
	return _Ronin.Contract.Withdrawal(&_Ronin.CallOpts, arg0)
}

// Withdrawal is a free data retrieval call binding the contract method 0x835fc6ca.
//
// Solidity: function withdrawal(uint256 ) view returns(uint256 id, uint8 kind, (address,address,uint256) mainchain, (address,address,uint256) ronin, (uint8,uint256,uint256) info)
func (_Ronin *RoninCallerSession) Withdrawal(arg0 *big.Int) (struct {
	Id        *big.Int
	Kind      uint8
	Mainchain TokenOwner
	Ronin     TokenOwner
	Info      TokenInfo
}, error) {
	return _Ronin.Contract.Withdrawal(&_Ronin.CallOpts, arg0)
}

// WithdrawalCount is a free data retrieval call binding the contract method 0x71706cbe.
//
// Solidity: function withdrawalCount() view returns(uint256)
func (_Ronin *RoninCaller) WithdrawalCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ronin.contract.Call(opts, &out, "withdrawalCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalCount is a free data retrieval call binding the contract method 0x71706cbe.
//
// Solidity: function withdrawalCount() view returns(uint256)
func (_Ronin *RoninSession) WithdrawalCount() (*big.Int, error) {
	return _Ronin.Contract.WithdrawalCount(&_Ronin.CallOpts)
}

// WithdrawalCount is a free data retrieval call binding the contract method 0x71706cbe.
//
// Solidity: function withdrawalCount() view returns(uint256)
func (_Ronin *RoninCallerSession) WithdrawalCount() (*big.Int, error) {
	return _Ronin.Contract.WithdrawalCount(&_Ronin.CallOpts)
}

// WithdrawalMigrated is a free data retrieval call binding the contract method 0x4f2717c7.
//
// Solidity: function withdrawalMigrated() view returns(bool)
func (_Ronin *RoninCaller) WithdrawalMigrated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Ronin.contract.Call(opts, &out, "withdrawalMigrated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawalMigrated is a free data retrieval call binding the contract method 0x4f2717c7.
//
// Solidity: function withdrawalMigrated() view returns(bool)
func (_Ronin *RoninSession) WithdrawalMigrated() (bool, error) {
	return _Ronin.Contract.WithdrawalMigrated(&_Ronin.CallOpts)
}

// WithdrawalMigrated is a free data retrieval call binding the contract method 0x4f2717c7.
//
// Solidity: function withdrawalMigrated() view returns(bool)
func (_Ronin *RoninCallerSession) WithdrawalMigrated() (bool, error) {
	return _Ronin.Contract.WithdrawalMigrated(&_Ronin.CallOpts)
}

// BulkDepositFor is a paid mutator transaction binding the contract method 0xc544376d.
//
// Solidity: function bulkDepositFor((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256))[] _receipts) returns()
func (_Ronin *RoninTransactor) BulkDepositFor(opts *bind.TransactOpts, _receipts []TransferReceipt) (*types.Transaction, error) {
	return _Ronin.contract.Transact(opts, "bulkDepositFor", _receipts)
}

// BulkDepositFor is a paid mutator transaction binding the contract method 0xc544376d.
//
// Solidity: function bulkDepositFor((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256))[] _receipts) returns()
func (_Ronin *RoninSession) BulkDepositFor(_receipts []TransferReceipt) (*types.Transaction, error) {
	return _Ronin.Contract.BulkDepositFor(&_Ronin.TransactOpts, _receipts)
}

// BulkDepositFor is a paid mutator transaction binding the contract method 0xc544376d.
//
// Solidity: function bulkDepositFor((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256))[] _receipts) returns()
func (_Ronin *RoninTransactorSession) BulkDepositFor(_receipts []TransferReceipt) (*types.Transaction, error) {
	return _Ronin.Contract.BulkDepositFor(&_Ronin.TransactOpts, _receipts)
}

// BulkRequestWithdrawalFor is a paid mutator transaction binding the contract method 0x5a7dd06a.
//
// Solidity: function bulkRequestWithdrawalFor((address,address,(uint8,uint256,uint256))[] _requests, uint256 _chainId) returns()
func (_Ronin *RoninTransactor) BulkRequestWithdrawalFor(opts *bind.TransactOpts, _requests []TransferRequest, _chainId *big.Int) (*types.Transaction, error) {
	return _Ronin.contract.Transact(opts, "bulkRequestWithdrawalFor", _requests, _chainId)
}

// BulkRequestWithdrawalFor is a paid mutator transaction binding the contract method 0x5a7dd06a.
//
// Solidity: function bulkRequestWithdrawalFor((address,address,(uint8,uint256,uint256))[] _requests, uint256 _chainId) returns()
func (_Ronin *RoninSession) BulkRequestWithdrawalFor(_requests []TransferRequest, _chainId *big.Int) (*types.Transaction, error) {
	return _Ronin.Contract.BulkRequestWithdrawalFor(&_Ronin.TransactOpts, _requests, _chainId)
}

// BulkRequestWithdrawalFor is a paid mutator transaction binding the contract method 0x5a7dd06a.
//
// Solidity: function bulkRequestWithdrawalFor((address,address,(uint8,uint256,uint256))[] _requests, uint256 _chainId) returns()
func (_Ronin *RoninTransactorSession) BulkRequestWithdrawalFor(_requests []TransferRequest, _chainId *big.Int) (*types.Transaction, error) {
	return _Ronin.Contract.BulkRequestWithdrawalFor(&_Ronin.TransactOpts, _requests, _chainId)
}

// BulkSubmitWithdrawalSignatures is a paid mutator transaction binding the contract method 0xfa389659.
//
// Solidity: function bulkSubmitWithdrawalSignatures(uint256[] _withdrawals, bytes[] _signatures) returns()
func (_Ronin *RoninTransactor) BulkSubmitWithdrawalSignatures(opts *bind.TransactOpts, _withdrawals []*big.Int, _signatures [][]byte) (*types.Transaction, error) {
	return _Ronin.contract.Transact(opts, "bulkSubmitWithdrawalSignatures", _withdrawals, _signatures)
}

// BulkSubmitWithdrawalSignatures is a paid mutator transaction binding the contract method 0xfa389659.
//
// Solidity: function bulkSubmitWithdrawalSignatures(uint256[] _withdrawals, bytes[] _signatures) returns()
func (_Ronin *RoninSession) BulkSubmitWithdrawalSignatures(_withdrawals []*big.Int, _signatures [][]byte) (*types.Transaction, error) {
	return _Ronin.Contract.BulkSubmitWithdrawalSignatures(&_Ronin.TransactOpts, _withdrawals, _signatures)
}

// BulkSubmitWithdrawalSignatures is a paid mutator transaction binding the contract method 0xfa389659.
//
// Solidity: function bulkSubmitWithdrawalSignatures(uint256[] _withdrawals, bytes[] _signatures) returns()
func (_Ronin *RoninTransactorSession) BulkSubmitWithdrawalSignatures(_withdrawals []*big.Int, _signatures [][]byte) (*types.Transaction, error) {
	return _Ronin.Contract.BulkSubmitWithdrawalSignatures(&_Ronin.TransactOpts, _withdrawals, _signatures)
}

// DepositFor is a paid mutator transaction binding the contract method 0x109679ef.
//
// Solidity: function depositFor((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt) returns()
func (_Ronin *RoninTransactor) DepositFor(opts *bind.TransactOpts, _receipt TransferReceipt) (*types.Transaction, error) {
	return _Ronin.contract.Transact(opts, "depositFor", _receipt)
}

// DepositFor is a paid mutator transaction binding the contract method 0x109679ef.
//
// Solidity: function depositFor((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt) returns()
func (_Ronin *RoninSession) DepositFor(_receipt TransferReceipt) (*types.Transaction, error) {
	return _Ronin.Contract.DepositFor(&_Ronin.TransactOpts, _receipt)
}

// DepositFor is a paid mutator transaction binding the contract method 0x109679ef.
//
// Solidity: function depositFor((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt) returns()
func (_Ronin *RoninTransactorSession) DepositFor(_receipt TransferReceipt) (*types.Transaction, error) {
	return _Ronin.Contract.DepositFor(&_Ronin.TransactOpts, _receipt)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Ronin *RoninTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ronin.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Ronin *RoninSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ronin.Contract.GrantRole(&_Ronin.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Ronin *RoninTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ronin.Contract.GrantRole(&_Ronin.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x1e5bc9ce.
//
// Solidity: function initialize(address _roleSetter, address _validatorContract, uint256 _numerator, uint256 _denominator, address[] _withdrawalMigrators, address[] _roninTokens, address[] _mainchainTokens, uint256[][2] _packedNumbers) returns()
func (_Ronin *RoninTransactor) Initialize(opts *bind.TransactOpts, _roleSetter common.Address, _validatorContract common.Address, _numerator *big.Int, _denominator *big.Int, _withdrawalMigrators []common.Address, _roninTokens []common.Address, _mainchainTokens []common.Address, _packedNumbers [2][]*big.Int) (*types.Transaction, error) {
	return _Ronin.contract.Transact(opts, "initialize", _roleSetter, _validatorContract, _numerator, _denominator, _withdrawalMigrators, _roninTokens, _mainchainTokens, _packedNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0x1e5bc9ce.
//
// Solidity: function initialize(address _roleSetter, address _validatorContract, uint256 _numerator, uint256 _denominator, address[] _withdrawalMigrators, address[] _roninTokens, address[] _mainchainTokens, uint256[][2] _packedNumbers) returns()
func (_Ronin *RoninSession) Initialize(_roleSetter common.Address, _validatorContract common.Address, _numerator *big.Int, _denominator *big.Int, _withdrawalMigrators []common.Address, _roninTokens []common.Address, _mainchainTokens []common.Address, _packedNumbers [2][]*big.Int) (*types.Transaction, error) {
	return _Ronin.Contract.Initialize(&_Ronin.TransactOpts, _roleSetter, _validatorContract, _numerator, _denominator, _withdrawalMigrators, _roninTokens, _mainchainTokens, _packedNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0x1e5bc9ce.
//
// Solidity: function initialize(address _roleSetter, address _validatorContract, uint256 _numerator, uint256 _denominator, address[] _withdrawalMigrators, address[] _roninTokens, address[] _mainchainTokens, uint256[][2] _packedNumbers) returns()
func (_Ronin *RoninTransactorSession) Initialize(_roleSetter common.Address, _validatorContract common.Address, _numerator *big.Int, _denominator *big.Int, _withdrawalMigrators []common.Address, _roninTokens []common.Address, _mainchainTokens []common.Address, _packedNumbers [2][]*big.Int) (*types.Transaction, error) {
	return _Ronin.Contract.Initialize(&_Ronin.TransactOpts, _roleSetter, _validatorContract, _numerator, _denominator, _withdrawalMigrators, _roninTokens, _mainchainTokens, _packedNumbers)
}

// MapTokens is a paid mutator transaction binding the contract method 0x6598c1bd.
//
// Solidity: function mapTokens(address[] _roninTokens, address[] _mainchainTokens, uint256[] _chainIds) returns()
func (_Ronin *RoninTransactor) MapTokens(opts *bind.TransactOpts, _roninTokens []common.Address, _mainchainTokens []common.Address, _chainIds []*big.Int) (*types.Transaction, error) {
	return _Ronin.contract.Transact(opts, "mapTokens", _roninTokens, _mainchainTokens, _chainIds)
}

// MapTokens is a paid mutator transaction binding the contract method 0x6598c1bd.
//
// Solidity: function mapTokens(address[] _roninTokens, address[] _mainchainTokens, uint256[] _chainIds) returns()
func (_Ronin *RoninSession) MapTokens(_roninTokens []common.Address, _mainchainTokens []common.Address, _chainIds []*big.Int) (*types.Transaction, error) {
	return _Ronin.Contract.MapTokens(&_Ronin.TransactOpts, _roninTokens, _mainchainTokens, _chainIds)
}

// MapTokens is a paid mutator transaction binding the contract method 0x6598c1bd.
//
// Solidity: function mapTokens(address[] _roninTokens, address[] _mainchainTokens, uint256[] _chainIds) returns()
func (_Ronin *RoninTransactorSession) MapTokens(_roninTokens []common.Address, _mainchainTokens []common.Address, _chainIds []*big.Int) (*types.Transaction, error) {
	return _Ronin.Contract.MapTokens(&_Ronin.TransactOpts, _roninTokens, _mainchainTokens, _chainIds)
}

// MarkWithdrawalMigrated is a paid mutator transaction binding the contract method 0x3b5afc22.
//
// Solidity: function markWithdrawalMigrated() returns()
func (_Ronin *RoninTransactor) MarkWithdrawalMigrated(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ronin.contract.Transact(opts, "markWithdrawalMigrated")
}

// MarkWithdrawalMigrated is a paid mutator transaction binding the contract method 0x3b5afc22.
//
// Solidity: function markWithdrawalMigrated() returns()
func (_Ronin *RoninSession) MarkWithdrawalMigrated() (*types.Transaction, error) {
	return _Ronin.Contract.MarkWithdrawalMigrated(&_Ronin.TransactOpts)
}

// MarkWithdrawalMigrated is a paid mutator transaction binding the contract method 0x3b5afc22.
//
// Solidity: function markWithdrawalMigrated() returns()
func (_Ronin *RoninTransactorSession) MarkWithdrawalMigrated() (*types.Transaction, error) {
	return _Ronin.Contract.MarkWithdrawalMigrated(&_Ronin.TransactOpts)
}

// MigrateWithdrawals is a paid mutator transaction binding the contract method 0x931ec987.
//
// Solidity: function migrateWithdrawals((address,address,(uint8,uint256,uint256))[] _requests, address[] _requesters) returns()
func (_Ronin *RoninTransactor) MigrateWithdrawals(opts *bind.TransactOpts, _requests []TransferRequest, _requesters []common.Address) (*types.Transaction, error) {
	return _Ronin.contract.Transact(opts, "migrateWithdrawals", _requests, _requesters)
}

// MigrateWithdrawals is a paid mutator transaction binding the contract method 0x931ec987.
//
// Solidity: function migrateWithdrawals((address,address,(uint8,uint256,uint256))[] _requests, address[] _requesters) returns()
func (_Ronin *RoninSession) MigrateWithdrawals(_requests []TransferRequest, _requesters []common.Address) (*types.Transaction, error) {
	return _Ronin.Contract.MigrateWithdrawals(&_Ronin.TransactOpts, _requests, _requesters)
}

// MigrateWithdrawals is a paid mutator transaction binding the contract method 0x931ec987.
//
// Solidity: function migrateWithdrawals((address,address,(uint8,uint256,uint256))[] _requests, address[] _requesters) returns()
func (_Ronin *RoninTransactorSession) MigrateWithdrawals(_requests []TransferRequest, _requesters []common.Address) (*types.Transaction, error) {
	return _Ronin.Contract.MigrateWithdrawals(&_Ronin.TransactOpts, _requests, _requesters)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Ronin *RoninTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ronin.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Ronin *RoninSession) Pause() (*types.Transaction, error) {
	return _Ronin.Contract.Pause(&_Ronin.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Ronin *RoninTransactorSession) Pause() (*types.Transaction, error) {
	return _Ronin.Contract.Pause(&_Ronin.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Ronin *RoninTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ronin.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Ronin *RoninSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ronin.Contract.RenounceRole(&_Ronin.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Ronin *RoninTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ronin.Contract.RenounceRole(&_Ronin.TransactOpts, role, account)
}

// RequestWithdrawalFor is a paid mutator transaction binding the contract method 0x0b1ff17f.
//
// Solidity: function requestWithdrawalFor((address,address,(uint8,uint256,uint256)) _request, uint256 _chainId) returns()
func (_Ronin *RoninTransactor) RequestWithdrawalFor(opts *bind.TransactOpts, _request TransferRequest, _chainId *big.Int) (*types.Transaction, error) {
	return _Ronin.contract.Transact(opts, "requestWithdrawalFor", _request, _chainId)
}

// RequestWithdrawalFor is a paid mutator transaction binding the contract method 0x0b1ff17f.
//
// Solidity: function requestWithdrawalFor((address,address,(uint8,uint256,uint256)) _request, uint256 _chainId) returns()
func (_Ronin *RoninSession) RequestWithdrawalFor(_request TransferRequest, _chainId *big.Int) (*types.Transaction, error) {
	return _Ronin.Contract.RequestWithdrawalFor(&_Ronin.TransactOpts, _request, _chainId)
}

// RequestWithdrawalFor is a paid mutator transaction binding the contract method 0x0b1ff17f.
//
// Solidity: function requestWithdrawalFor((address,address,(uint8,uint256,uint256)) _request, uint256 _chainId) returns()
func (_Ronin *RoninTransactorSession) RequestWithdrawalFor(_request TransferRequest, _chainId *big.Int) (*types.Transaction, error) {
	return _Ronin.Contract.RequestWithdrawalFor(&_Ronin.TransactOpts, _request, _chainId)
}

// RequestWithdrawalSignatures is a paid mutator transaction binding the contract method 0x47b56b2c.
//
// Solidity: function requestWithdrawalSignatures(uint256 _withdrawalId) returns()
func (_Ronin *RoninTransactor) RequestWithdrawalSignatures(opts *bind.TransactOpts, _withdrawalId *big.Int) (*types.Transaction, error) {
	return _Ronin.contract.Transact(opts, "requestWithdrawalSignatures", _withdrawalId)
}

// RequestWithdrawalSignatures is a paid mutator transaction binding the contract method 0x47b56b2c.
//
// Solidity: function requestWithdrawalSignatures(uint256 _withdrawalId) returns()
func (_Ronin *RoninSession) RequestWithdrawalSignatures(_withdrawalId *big.Int) (*types.Transaction, error) {
	return _Ronin.Contract.RequestWithdrawalSignatures(&_Ronin.TransactOpts, _withdrawalId)
}

// RequestWithdrawalSignatures is a paid mutator transaction binding the contract method 0x47b56b2c.
//
// Solidity: function requestWithdrawalSignatures(uint256 _withdrawalId) returns()
func (_Ronin *RoninTransactorSession) RequestWithdrawalSignatures(_withdrawalId *big.Int) (*types.Transaction, error) {
	return _Ronin.Contract.RequestWithdrawalSignatures(&_Ronin.TransactOpts, _withdrawalId)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Ronin *RoninTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ronin.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Ronin *RoninSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ronin.Contract.RevokeRole(&_Ronin.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Ronin *RoninTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ronin.Contract.RevokeRole(&_Ronin.TransactOpts, role, account)
}

// SetMinimumThresholds is a paid mutator transaction binding the contract method 0x64363f78.
//
// Solidity: function setMinimumThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_Ronin *RoninTransactor) SetMinimumThresholds(opts *bind.TransactOpts, _tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _Ronin.contract.Transact(opts, "setMinimumThresholds", _tokens, _thresholds)
}

// SetMinimumThresholds is a paid mutator transaction binding the contract method 0x64363f78.
//
// Solidity: function setMinimumThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_Ronin *RoninSession) SetMinimumThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _Ronin.Contract.SetMinimumThresholds(&_Ronin.TransactOpts, _tokens, _thresholds)
}

// SetMinimumThresholds is a paid mutator transaction binding the contract method 0x64363f78.
//
// Solidity: function setMinimumThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_Ronin *RoninTransactorSession) SetMinimumThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _Ronin.Contract.SetMinimumThresholds(&_Ronin.TransactOpts, _tokens, _thresholds)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 _numerator, uint256 _denominator) returns(uint256, uint256)
func (_Ronin *RoninTransactor) SetThreshold(opts *bind.TransactOpts, _numerator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _Ronin.contract.Transact(opts, "setThreshold", _numerator, _denominator)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 _numerator, uint256 _denominator) returns(uint256, uint256)
func (_Ronin *RoninSession) SetThreshold(_numerator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _Ronin.Contract.SetThreshold(&_Ronin.TransactOpts, _numerator, _denominator)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 _numerator, uint256 _denominator) returns(uint256, uint256)
func (_Ronin *RoninTransactorSession) SetThreshold(_numerator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _Ronin.Contract.SetThreshold(&_Ronin.TransactOpts, _numerator, _denominator)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _validatorContract) returns()
func (_Ronin *RoninTransactor) SetValidatorContract(opts *bind.TransactOpts, _validatorContract common.Address) (*types.Transaction, error) {
	return _Ronin.contract.Transact(opts, "setValidatorContract", _validatorContract)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _validatorContract) returns()
func (_Ronin *RoninSession) SetValidatorContract(_validatorContract common.Address) (*types.Transaction, error) {
	return _Ronin.Contract.SetValidatorContract(&_Ronin.TransactOpts, _validatorContract)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _validatorContract) returns()
func (_Ronin *RoninTransactorSession) SetValidatorContract(_validatorContract common.Address) (*types.Transaction, error) {
	return _Ronin.Contract.SetValidatorContract(&_Ronin.TransactOpts, _validatorContract)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Ronin *RoninTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ronin.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Ronin *RoninSession) Unpause() (*types.Transaction, error) {
	return _Ronin.Contract.Unpause(&_Ronin.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Ronin *RoninTransactorSession) Unpause() (*types.Transaction, error) {
	return _Ronin.Contract.Unpause(&_Ronin.TransactOpts)
}

// RoninDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Ronin contract.
type RoninDepositedIterator struct {
	Event *RoninDeposited // Event containing the contract specifics and raw log

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
func (it *RoninDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninDeposited)
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
		it.Event = new(RoninDeposited)
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
func (it *RoninDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninDeposited represents a Deposited event raised by the Ronin contract.
type RoninDeposited struct {
	ReceiptHash [32]byte
	Receipt     TransferReceipt
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x8d20d8121a34dded9035ff5b43e901c142824f7a22126392992c353c37890524.
//
// Solidity: event Deposited(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_Ronin *RoninFilterer) FilterDeposited(opts *bind.FilterOpts) (*RoninDepositedIterator, error) {

	logs, sub, err := _Ronin.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &RoninDepositedIterator{contract: _Ronin.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x8d20d8121a34dded9035ff5b43e901c142824f7a22126392992c353c37890524.
//
// Solidity: event Deposited(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_Ronin *RoninFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *RoninDeposited) (event.Subscription, error) {

	logs, sub, err := _Ronin.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninDeposited)
				if err := _Ronin.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x8d20d8121a34dded9035ff5b43e901c142824f7a22126392992c353c37890524.
//
// Solidity: event Deposited(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_Ronin *RoninFilterer) ParseDeposited(log types.Log) (*RoninDeposited, error) {
	event := new(RoninDeposited)
	if err := _Ronin.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninMinimumThresholdsUpdatedIterator is returned from FilterMinimumThresholdsUpdated and is used to iterate over the raw logs and unpacked data for MinimumThresholdsUpdated events raised by the Ronin contract.
type RoninMinimumThresholdsUpdatedIterator struct {
	Event *RoninMinimumThresholdsUpdated // Event containing the contract specifics and raw log

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
func (it *RoninMinimumThresholdsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninMinimumThresholdsUpdated)
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
		it.Event = new(RoninMinimumThresholdsUpdated)
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
func (it *RoninMinimumThresholdsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninMinimumThresholdsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninMinimumThresholdsUpdated represents a MinimumThresholdsUpdated event raised by the Ronin contract.
type RoninMinimumThresholdsUpdated struct {
	Tokens    []common.Address
	Threshold []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinimumThresholdsUpdated is a free log retrieval operation binding the contract event 0x6f52f53a938df83439fa4c6055c7df0a6906d621aa6dfa4708187037fdfc41da.
//
// Solidity: event MinimumThresholdsUpdated(address[] tokens, uint256[] threshold)
func (_Ronin *RoninFilterer) FilterMinimumThresholdsUpdated(opts *bind.FilterOpts) (*RoninMinimumThresholdsUpdatedIterator, error) {

	logs, sub, err := _Ronin.contract.FilterLogs(opts, "MinimumThresholdsUpdated")
	if err != nil {
		return nil, err
	}
	return &RoninMinimumThresholdsUpdatedIterator{contract: _Ronin.contract, event: "MinimumThresholdsUpdated", logs: logs, sub: sub}, nil
}

// WatchMinimumThresholdsUpdated is a free log subscription operation binding the contract event 0x6f52f53a938df83439fa4c6055c7df0a6906d621aa6dfa4708187037fdfc41da.
//
// Solidity: event MinimumThresholdsUpdated(address[] tokens, uint256[] threshold)
func (_Ronin *RoninFilterer) WatchMinimumThresholdsUpdated(opts *bind.WatchOpts, sink chan<- *RoninMinimumThresholdsUpdated) (event.Subscription, error) {

	logs, sub, err := _Ronin.contract.WatchLogs(opts, "MinimumThresholdsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninMinimumThresholdsUpdated)
				if err := _Ronin.contract.UnpackLog(event, "MinimumThresholdsUpdated", log); err != nil {
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

// ParseMinimumThresholdsUpdated is a log parse operation binding the contract event 0x6f52f53a938df83439fa4c6055c7df0a6906d621aa6dfa4708187037fdfc41da.
//
// Solidity: event MinimumThresholdsUpdated(address[] tokens, uint256[] threshold)
func (_Ronin *RoninFilterer) ParseMinimumThresholdsUpdated(log types.Log) (*RoninMinimumThresholdsUpdated, error) {
	event := new(RoninMinimumThresholdsUpdated)
	if err := _Ronin.contract.UnpackLog(event, "MinimumThresholdsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Ronin contract.
type RoninPausedIterator struct {
	Event *RoninPaused // Event containing the contract specifics and raw log

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
func (it *RoninPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninPaused)
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
		it.Event = new(RoninPaused)
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
func (it *RoninPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninPaused represents a Paused event raised by the Ronin contract.
type RoninPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Ronin *RoninFilterer) FilterPaused(opts *bind.FilterOpts) (*RoninPausedIterator, error) {

	logs, sub, err := _Ronin.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &RoninPausedIterator{contract: _Ronin.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Ronin *RoninFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *RoninPaused) (event.Subscription, error) {

	logs, sub, err := _Ronin.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninPaused)
				if err := _Ronin.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Ronin *RoninFilterer) ParsePaused(log types.Log) (*RoninPaused, error) {
	event := new(RoninPaused)
	if err := _Ronin.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Ronin contract.
type RoninRoleAdminChangedIterator struct {
	Event *RoninRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *RoninRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninRoleAdminChanged)
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
		it.Event = new(RoninRoleAdminChanged)
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
func (it *RoninRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninRoleAdminChanged represents a RoleAdminChanged event raised by the Ronin contract.
type RoninRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Ronin *RoninFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*RoninRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Ronin.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &RoninRoleAdminChangedIterator{contract: _Ronin.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Ronin *RoninFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *RoninRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Ronin.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninRoleAdminChanged)
				if err := _Ronin.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Ronin *RoninFilterer) ParseRoleAdminChanged(log types.Log) (*RoninRoleAdminChanged, error) {
	event := new(RoninRoleAdminChanged)
	if err := _Ronin.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Ronin contract.
type RoninRoleGrantedIterator struct {
	Event *RoninRoleGranted // Event containing the contract specifics and raw log

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
func (it *RoninRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninRoleGranted)
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
		it.Event = new(RoninRoleGranted)
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
func (it *RoninRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninRoleGranted represents a RoleGranted event raised by the Ronin contract.
type RoninRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Ronin *RoninFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RoninRoleGrantedIterator, error) {

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

	logs, sub, err := _Ronin.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RoninRoleGrantedIterator{contract: _Ronin.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Ronin *RoninFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RoninRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Ronin.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninRoleGranted)
				if err := _Ronin.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Ronin *RoninFilterer) ParseRoleGranted(log types.Log) (*RoninRoleGranted, error) {
	event := new(RoninRoleGranted)
	if err := _Ronin.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Ronin contract.
type RoninRoleRevokedIterator struct {
	Event *RoninRoleRevoked // Event containing the contract specifics and raw log

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
func (it *RoninRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninRoleRevoked)
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
		it.Event = new(RoninRoleRevoked)
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
func (it *RoninRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninRoleRevoked represents a RoleRevoked event raised by the Ronin contract.
type RoninRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Ronin *RoninFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RoninRoleRevokedIterator, error) {

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

	logs, sub, err := _Ronin.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RoninRoleRevokedIterator{contract: _Ronin.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Ronin *RoninFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RoninRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Ronin.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninRoleRevoked)
				if err := _Ronin.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Ronin *RoninFilterer) ParseRoleRevoked(log types.Log) (*RoninRoleRevoked, error) {
	event := new(RoninRoleRevoked)
	if err := _Ronin.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninThresholdUpdatedIterator is returned from FilterThresholdUpdated and is used to iterate over the raw logs and unpacked data for ThresholdUpdated events raised by the Ronin contract.
type RoninThresholdUpdatedIterator struct {
	Event *RoninThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *RoninThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninThresholdUpdated)
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
		it.Event = new(RoninThresholdUpdated)
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
func (it *RoninThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninThresholdUpdated represents a ThresholdUpdated event raised by the Ronin contract.
type RoninThresholdUpdated struct {
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
func (_Ronin *RoninFilterer) FilterThresholdUpdated(opts *bind.FilterOpts, _nonce []*big.Int, _numerator []*big.Int, _denominator []*big.Int) (*RoninThresholdUpdatedIterator, error) {

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

	logs, sub, err := _Ronin.contract.FilterLogs(opts, "ThresholdUpdated", _nonceRule, _numeratorRule, _denominatorRule)
	if err != nil {
		return nil, err
	}
	return &RoninThresholdUpdatedIterator{contract: _Ronin.contract, event: "ThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchThresholdUpdated is a free log subscription operation binding the contract event 0x976f8a9c5bdf8248dec172376d6e2b80a8e3df2f0328e381c6db8e1cf138c0f8.
//
// Solidity: event ThresholdUpdated(uint256 indexed _nonce, uint256 indexed _numerator, uint256 indexed _denominator, uint256 _previousNumerator, uint256 _previousDenominator)
func (_Ronin *RoninFilterer) WatchThresholdUpdated(opts *bind.WatchOpts, sink chan<- *RoninThresholdUpdated, _nonce []*big.Int, _numerator []*big.Int, _denominator []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Ronin.contract.WatchLogs(opts, "ThresholdUpdated", _nonceRule, _numeratorRule, _denominatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninThresholdUpdated)
				if err := _Ronin.contract.UnpackLog(event, "ThresholdUpdated", log); err != nil {
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
func (_Ronin *RoninFilterer) ParseThresholdUpdated(log types.Log) (*RoninThresholdUpdated, error) {
	event := new(RoninThresholdUpdated)
	if err := _Ronin.contract.UnpackLog(event, "ThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninTokenMappedIterator is returned from FilterTokenMapped and is used to iterate over the raw logs and unpacked data for TokenMapped events raised by the Ronin contract.
type RoninTokenMappedIterator struct {
	Event *RoninTokenMapped // Event containing the contract specifics and raw log

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
func (it *RoninTokenMappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninTokenMapped)
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
		it.Event = new(RoninTokenMapped)
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
func (it *RoninTokenMappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninTokenMappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninTokenMapped represents a TokenMapped event raised by the Ronin contract.
type RoninTokenMapped struct {
	RoninTokens     []common.Address
	MainchainTokens []common.Address
	ChainIds        []*big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokenMapped is a free log retrieval operation binding the contract event 0x332d1046c291868dcf6424605d8c4f46a10977395d49f764b86b8d441ba82bd3.
//
// Solidity: event TokenMapped(address[] roninTokens, address[] mainchainTokens, uint256[] chainIds)
func (_Ronin *RoninFilterer) FilterTokenMapped(opts *bind.FilterOpts) (*RoninTokenMappedIterator, error) {

	logs, sub, err := _Ronin.contract.FilterLogs(opts, "TokenMapped")
	if err != nil {
		return nil, err
	}
	return &RoninTokenMappedIterator{contract: _Ronin.contract, event: "TokenMapped", logs: logs, sub: sub}, nil
}

// WatchTokenMapped is a free log subscription operation binding the contract event 0x332d1046c291868dcf6424605d8c4f46a10977395d49f764b86b8d441ba82bd3.
//
// Solidity: event TokenMapped(address[] roninTokens, address[] mainchainTokens, uint256[] chainIds)
func (_Ronin *RoninFilterer) WatchTokenMapped(opts *bind.WatchOpts, sink chan<- *RoninTokenMapped) (event.Subscription, error) {

	logs, sub, err := _Ronin.contract.WatchLogs(opts, "TokenMapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninTokenMapped)
				if err := _Ronin.contract.UnpackLog(event, "TokenMapped", log); err != nil {
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

// ParseTokenMapped is a log parse operation binding the contract event 0x332d1046c291868dcf6424605d8c4f46a10977395d49f764b86b8d441ba82bd3.
//
// Solidity: event TokenMapped(address[] roninTokens, address[] mainchainTokens, uint256[] chainIds)
func (_Ronin *RoninFilterer) ParseTokenMapped(log types.Log) (*RoninTokenMapped, error) {
	event := new(RoninTokenMapped)
	if err := _Ronin.contract.UnpackLog(event, "TokenMapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Ronin contract.
type RoninUnpausedIterator struct {
	Event *RoninUnpaused // Event containing the contract specifics and raw log

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
func (it *RoninUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninUnpaused)
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
		it.Event = new(RoninUnpaused)
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
func (it *RoninUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninUnpaused represents a Unpaused event raised by the Ronin contract.
type RoninUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Ronin *RoninFilterer) FilterUnpaused(opts *bind.FilterOpts) (*RoninUnpausedIterator, error) {

	logs, sub, err := _Ronin.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &RoninUnpausedIterator{contract: _Ronin.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Ronin *RoninFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *RoninUnpaused) (event.Subscription, error) {

	logs, sub, err := _Ronin.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninUnpaused)
				if err := _Ronin.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Ronin *RoninFilterer) ParseUnpaused(log types.Log) (*RoninUnpaused, error) {
	event := new(RoninUnpaused)
	if err := _Ronin.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorContractUpdatedIterator is returned from FilterValidatorContractUpdated and is used to iterate over the raw logs and unpacked data for ValidatorContractUpdated events raised by the Ronin contract.
type RoninValidatorContractUpdatedIterator struct {
	Event *RoninValidatorContractUpdated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorContractUpdated)
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
		it.Event = new(RoninValidatorContractUpdated)
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
func (it *RoninValidatorContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorContractUpdated represents a ValidatorContractUpdated event raised by the Ronin contract.
type RoninValidatorContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterValidatorContractUpdated is a free log retrieval operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_Ronin *RoninFilterer) FilterValidatorContractUpdated(opts *bind.FilterOpts) (*RoninValidatorContractUpdatedIterator, error) {

	logs, sub, err := _Ronin.contract.FilterLogs(opts, "ValidatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorContractUpdatedIterator{contract: _Ronin.contract, event: "ValidatorContractUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorContractUpdated is a free log subscription operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_Ronin *RoninFilterer) WatchValidatorContractUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorContractUpdated) (event.Subscription, error) {

	logs, sub, err := _Ronin.contract.WatchLogs(opts, "ValidatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorContractUpdated)
				if err := _Ronin.contract.UnpackLog(event, "ValidatorContractUpdated", log); err != nil {
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
func (_Ronin *RoninFilterer) ParseValidatorContractUpdated(log types.Log) (*RoninValidatorContractUpdated, error) {
	event := new(RoninValidatorContractUpdated)
	if err := _Ronin.contract.UnpackLog(event, "ValidatorContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninWithdrawalRequestedIterator is returned from FilterWithdrawalRequested and is used to iterate over the raw logs and unpacked data for WithdrawalRequested events raised by the Ronin contract.
type RoninWithdrawalRequestedIterator struct {
	Event *RoninWithdrawalRequested // Event containing the contract specifics and raw log

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
func (it *RoninWithdrawalRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninWithdrawalRequested)
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
		it.Event = new(RoninWithdrawalRequested)
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
func (it *RoninWithdrawalRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninWithdrawalRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninWithdrawalRequested represents a WithdrawalRequested event raised by the Ronin contract.
type RoninWithdrawalRequested struct {
	ReceiptHash [32]byte
	Arg1        TransferReceipt
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalRequested is a free log retrieval operation binding the contract event 0xf313c253a5be72c29d0deb2c8768a9543744ac03d6b3cafd50cc976f1c2632fc.
//
// Solidity: event WithdrawalRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) arg1)
func (_Ronin *RoninFilterer) FilterWithdrawalRequested(opts *bind.FilterOpts) (*RoninWithdrawalRequestedIterator, error) {

	logs, sub, err := _Ronin.contract.FilterLogs(opts, "WithdrawalRequested")
	if err != nil {
		return nil, err
	}
	return &RoninWithdrawalRequestedIterator{contract: _Ronin.contract, event: "WithdrawalRequested", logs: logs, sub: sub}, nil
}

// WatchWithdrawalRequested is a free log subscription operation binding the contract event 0xf313c253a5be72c29d0deb2c8768a9543744ac03d6b3cafd50cc976f1c2632fc.
//
// Solidity: event WithdrawalRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) arg1)
func (_Ronin *RoninFilterer) WatchWithdrawalRequested(opts *bind.WatchOpts, sink chan<- *RoninWithdrawalRequested) (event.Subscription, error) {

	logs, sub, err := _Ronin.contract.WatchLogs(opts, "WithdrawalRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninWithdrawalRequested)
				if err := _Ronin.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
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

// ParseWithdrawalRequested is a log parse operation binding the contract event 0xf313c253a5be72c29d0deb2c8768a9543744ac03d6b3cafd50cc976f1c2632fc.
//
// Solidity: event WithdrawalRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) arg1)
func (_Ronin *RoninFilterer) ParseWithdrawalRequested(log types.Log) (*RoninWithdrawalRequested, error) {
	event := new(RoninWithdrawalRequested)
	if err := _Ronin.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninWithdrawalSignaturesRequestedIterator is returned from FilterWithdrawalSignaturesRequested and is used to iterate over the raw logs and unpacked data for WithdrawalSignaturesRequested events raised by the Ronin contract.
type RoninWithdrawalSignaturesRequestedIterator struct {
	Event *RoninWithdrawalSignaturesRequested // Event containing the contract specifics and raw log

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
func (it *RoninWithdrawalSignaturesRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninWithdrawalSignaturesRequested)
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
		it.Event = new(RoninWithdrawalSignaturesRequested)
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
func (it *RoninWithdrawalSignaturesRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninWithdrawalSignaturesRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninWithdrawalSignaturesRequested represents a WithdrawalSignaturesRequested event raised by the Ronin contract.
type RoninWithdrawalSignaturesRequested struct {
	ReceiptHash [32]byte
	Arg1        TransferReceipt
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalSignaturesRequested is a free log retrieval operation binding the contract event 0x04e8cbd836dea43a2dc7eb19de345cca3a8e6978a2ef5225d924775500f67c7c.
//
// Solidity: event WithdrawalSignaturesRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) arg1)
func (_Ronin *RoninFilterer) FilterWithdrawalSignaturesRequested(opts *bind.FilterOpts) (*RoninWithdrawalSignaturesRequestedIterator, error) {

	logs, sub, err := _Ronin.contract.FilterLogs(opts, "WithdrawalSignaturesRequested")
	if err != nil {
		return nil, err
	}
	return &RoninWithdrawalSignaturesRequestedIterator{contract: _Ronin.contract, event: "WithdrawalSignaturesRequested", logs: logs, sub: sub}, nil
}

// WatchWithdrawalSignaturesRequested is a free log subscription operation binding the contract event 0x04e8cbd836dea43a2dc7eb19de345cca3a8e6978a2ef5225d924775500f67c7c.
//
// Solidity: event WithdrawalSignaturesRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) arg1)
func (_Ronin *RoninFilterer) WatchWithdrawalSignaturesRequested(opts *bind.WatchOpts, sink chan<- *RoninWithdrawalSignaturesRequested) (event.Subscription, error) {

	logs, sub, err := _Ronin.contract.WatchLogs(opts, "WithdrawalSignaturesRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninWithdrawalSignaturesRequested)
				if err := _Ronin.contract.UnpackLog(event, "WithdrawalSignaturesRequested", log); err != nil {
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

// ParseWithdrawalSignaturesRequested is a log parse operation binding the contract event 0x04e8cbd836dea43a2dc7eb19de345cca3a8e6978a2ef5225d924775500f67c7c.
//
// Solidity: event WithdrawalSignaturesRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) arg1)
func (_Ronin *RoninFilterer) ParseWithdrawalSignaturesRequested(log types.Log) (*RoninWithdrawalSignaturesRequested, error) {
	event := new(RoninWithdrawalSignaturesRequested)
	if err := _Ronin.contract.UnpackLog(event, "WithdrawalSignaturesRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
