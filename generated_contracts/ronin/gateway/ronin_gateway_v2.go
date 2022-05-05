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
	TokenAddr     common.Address
	Info          TokenInfo
}

// GatewayMetaData contains all meta data concerning the Gateway contract.
var GatewayMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"receiptHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"receiptHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"MainchainWithdrew\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"threshold\",\"type\":\"uint256[]\"}],\"name\":\"MinimumThresholdsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"numerator\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousNumerator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousDenominator\",\"type\":\"uint256\"}],\"name\":\"ThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"roninTokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"mainchainTokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"chainIds\",\"type\":\"uint256[]\"}],\"name\":\"TokenMapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIWeightedValidator\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ValidatorContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"receiptHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"WithdrawalRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"receiptHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"WithdrawalSignaturesRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWAL_MIGRATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_withdrawalIds\",\"type\":\"uint256[]\"}],\"name\":\"bulkAcknowledgeMainchainWithdrew\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structTransfer.Receipt[]\",\"name\":\"_receipts\",\"type\":\"tuple[]\"}],\"name\":\"bulkDepositFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipientAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structTransfer.Request[]\",\"name\":\"_requests\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"bulkRequestWithdrawalFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_withdrawals\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_signatures\",\"type\":\"bytes[]\"}],\"name\":\"bulkSubmitWithdrawalSignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_voteWeight\",\"type\":\"uint256\"}],\"name\":\"checkThreshold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structTransfer.Receipt\",\"name\":\"_receipt\",\"type\":\"tuple\"}],\"name\":\"depositFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"depositVote\",\"outputs\":[{\"internalType\":\"enumGatewayGovernance.VoteStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"finalHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_roninToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"getMainchainToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawalId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"}],\"name\":\"getWithdrawalSignatures\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_signatures\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_roleSetter\",\"type\":\"address\"},{\"internalType\":\"contractIWeightedValidator\",\"name\":\"_validatorContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_denominator\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_withdrawalMigrators\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_roninTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_mainchainTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[][2]\",\"name\":\"_packedNumbers\",\"type\":\"uint256[][2]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mainchainWithdrew\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_roninTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_mainchainTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_chainIds\",\"type\":\"uint256[]\"}],\"name\":\"mapTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"markWithdrawalMigrated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipientAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structTransfer.Request[]\",\"name\":\"_requests\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"_requesters\",\"type\":\"address[]\"}],\"name\":\"migrateWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minimumThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumVoteWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipientAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structTransfer.Request\",\"name\":\"_request\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"requestWithdrawalFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawalId\",\"type\":\"uint256\"}],\"name\":\"requestWithdrawalSignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_thresholds\",\"type\":\"uint256[]\"}],\"name\":\"setMinimumThresholds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_denominator\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIWeightedValidator\",\"name\":\"_validatorContract\",\"type\":\"address\"}],\"name\":\"setValidatorContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorContract\",\"outputs\":[{\"internalType\":\"contractIWeightedValidator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalMigrated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// GatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayMetaData.ABI instead.
var GatewayABI = GatewayMetaData.ABI

// Gateway is an auto generated Go binding around an Ethereum contract.
type Gateway struct {
	GatewayCaller     // Read-only binding to the contract
	GatewayTransactor // Write-only binding to the contract
	GatewayFilterer   // Log filterer for contract events
}

// GatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewaySession struct {
	Contract     *Gateway          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayCallerSession struct {
	Contract *GatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// GatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayTransactorSession struct {
	Contract     *GatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayRaw struct {
	Contract *Gateway // Generic contract binding to access the raw methods on
}

// GatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayCallerRaw struct {
	Contract *GatewayCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayTransactorRaw struct {
	Contract *GatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGateway creates a new instance of Gateway, bound to a specific deployed contract.
func NewGateway(address common.Address, backend bind.ContractBackend) (*Gateway, error) {
	contract, err := bindGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gateway{GatewayCaller: GatewayCaller{contract: contract}, GatewayTransactor: GatewayTransactor{contract: contract}, GatewayFilterer: GatewayFilterer{contract: contract}}, nil
}

// NewGatewayCaller creates a new read-only instance of Gateway, bound to a specific deployed contract.
func NewGatewayCaller(address common.Address, caller bind.ContractCaller) (*GatewayCaller, error) {
	contract, err := bindGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayCaller{contract: contract}, nil
}

// NewGatewayTransactor creates a new write-only instance of Gateway, bound to a specific deployed contract.
func NewGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayTransactor, error) {
	contract, err := bindGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayTransactor{contract: contract}, nil
}

// NewGatewayFilterer creates a new log filterer instance of Gateway, bound to a specific deployed contract.
func NewGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayFilterer, error) {
	contract, err := bindGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayFilterer{contract: contract}, nil
}

// bindGateway binds a generic wrapper to an already deployed contract.
func bindGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GatewayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gateway *GatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gateway.Contract.GatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gateway *GatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.Contract.GatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gateway *GatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gateway.Contract.GatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gateway *GatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gateway *GatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gateway *GatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gateway.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Gateway *GatewayCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Gateway *GatewaySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Gateway.Contract.DEFAULTADMINROLE(&_Gateway.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Gateway *GatewayCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Gateway.Contract.DEFAULTADMINROLE(&_Gateway.CallOpts)
}

// WITHDRAWALMIGRATOR is a free data retrieval call binding the contract method 0xfe90d9c2.
//
// Solidity: function WITHDRAWAL_MIGRATOR() view returns(bytes32)
func (_Gateway *GatewayCaller) WITHDRAWALMIGRATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "WITHDRAWAL_MIGRATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWALMIGRATOR is a free data retrieval call binding the contract method 0xfe90d9c2.
//
// Solidity: function WITHDRAWAL_MIGRATOR() view returns(bytes32)
func (_Gateway *GatewaySession) WITHDRAWALMIGRATOR() ([32]byte, error) {
	return _Gateway.Contract.WITHDRAWALMIGRATOR(&_Gateway.CallOpts)
}

// WITHDRAWALMIGRATOR is a free data retrieval call binding the contract method 0xfe90d9c2.
//
// Solidity: function WITHDRAWAL_MIGRATOR() view returns(bytes32)
func (_Gateway *GatewayCallerSession) WITHDRAWALMIGRATOR() ([32]byte, error) {
	return _Gateway.Contract.WITHDRAWALMIGRATOR(&_Gateway.CallOpts)
}

// CheckThreshold is a free data retrieval call binding the contract method 0xdafae408.
//
// Solidity: function checkThreshold(uint256 _voteWeight) view returns(bool)
func (_Gateway *GatewayCaller) CheckThreshold(opts *bind.CallOpts, _voteWeight *big.Int) (bool, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "checkThreshold", _voteWeight)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckThreshold is a free data retrieval call binding the contract method 0xdafae408.
//
// Solidity: function checkThreshold(uint256 _voteWeight) view returns(bool)
func (_Gateway *GatewaySession) CheckThreshold(_voteWeight *big.Int) (bool, error) {
	return _Gateway.Contract.CheckThreshold(&_Gateway.CallOpts, _voteWeight)
}

// CheckThreshold is a free data retrieval call binding the contract method 0xdafae408.
//
// Solidity: function checkThreshold(uint256 _voteWeight) view returns(bool)
func (_Gateway *GatewayCallerSession) CheckThreshold(_voteWeight *big.Int) (bool, error) {
	return _Gateway.Contract.CheckThreshold(&_Gateway.CallOpts, _voteWeight)
}

// DepositVote is a free data retrieval call binding the contract method 0x4d92c4f0.
//
// Solidity: function depositVote(uint256 , uint256 ) view returns(uint8 status, bytes32 finalHash)
func (_Gateway *GatewayCaller) DepositVote(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Status    uint8
	FinalHash [32]byte
}, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "depositVote", arg0, arg1)

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
func (_Gateway *GatewaySession) DepositVote(arg0 *big.Int, arg1 *big.Int) (struct {
	Status    uint8
	FinalHash [32]byte
}, error) {
	return _Gateway.Contract.DepositVote(&_Gateway.CallOpts, arg0, arg1)
}

// DepositVote is a free data retrieval call binding the contract method 0x4d92c4f0.
//
// Solidity: function depositVote(uint256 , uint256 ) view returns(uint8 status, bytes32 finalHash)
func (_Gateway *GatewayCallerSession) DepositVote(arg0 *big.Int, arg1 *big.Int) (struct {
	Status    uint8
	FinalHash [32]byte
}, error) {
	return _Gateway.Contract.DepositVote(&_Gateway.CallOpts, arg0, arg1)
}

// GetMainchainToken is a free data retrieval call binding the contract method 0x5d6a9a90.
//
// Solidity: function getMainchainToken(address _roninToken, uint256 _chainId) view returns(address _addr)
func (_Gateway *GatewayCaller) GetMainchainToken(opts *bind.CallOpts, _roninToken common.Address, _chainId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getMainchainToken", _roninToken, _chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMainchainToken is a free data retrieval call binding the contract method 0x5d6a9a90.
//
// Solidity: function getMainchainToken(address _roninToken, uint256 _chainId) view returns(address _addr)
func (_Gateway *GatewaySession) GetMainchainToken(_roninToken common.Address, _chainId *big.Int) (common.Address, error) {
	return _Gateway.Contract.GetMainchainToken(&_Gateway.CallOpts, _roninToken, _chainId)
}

// GetMainchainToken is a free data retrieval call binding the contract method 0x5d6a9a90.
//
// Solidity: function getMainchainToken(address _roninToken, uint256 _chainId) view returns(address _addr)
func (_Gateway *GatewayCallerSession) GetMainchainToken(_roninToken common.Address, _chainId *big.Int) (common.Address, error) {
	return _Gateway.Contract.GetMainchainToken(&_Gateway.CallOpts, _roninToken, _chainId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Gateway *GatewayCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Gateway *GatewaySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Gateway.Contract.GetRoleAdmin(&_Gateway.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Gateway *GatewayCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Gateway.Contract.GetRoleAdmin(&_Gateway.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Gateway *GatewayCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Gateway *GatewaySession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Gateway.Contract.GetRoleMember(&_Gateway.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Gateway *GatewayCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Gateway.Contract.GetRoleMember(&_Gateway.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Gateway *GatewayCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Gateway *GatewaySession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Gateway.Contract.GetRoleMemberCount(&_Gateway.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Gateway *GatewayCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Gateway.Contract.GetRoleMemberCount(&_Gateway.CallOpts, role)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256, uint256)
func (_Gateway *GatewayCaller) GetThreshold(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getThreshold")

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
func (_Gateway *GatewaySession) GetThreshold() (*big.Int, *big.Int, error) {
	return _Gateway.Contract.GetThreshold(&_Gateway.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256, uint256)
func (_Gateway *GatewayCallerSession) GetThreshold() (*big.Int, *big.Int, error) {
	return _Gateway.Contract.GetThreshold(&_Gateway.CallOpts)
}

// GetWithdrawalSignatures is a free data retrieval call binding the contract method 0xecc83649.
//
// Solidity: function getWithdrawalSignatures(uint256 _withdrawalId, address[] _validators) view returns(bytes[] _signatures)
func (_Gateway *GatewayCaller) GetWithdrawalSignatures(opts *bind.CallOpts, _withdrawalId *big.Int, _validators []common.Address) ([][]byte, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getWithdrawalSignatures", _withdrawalId, _validators)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetWithdrawalSignatures is a free data retrieval call binding the contract method 0xecc83649.
//
// Solidity: function getWithdrawalSignatures(uint256 _withdrawalId, address[] _validators) view returns(bytes[] _signatures)
func (_Gateway *GatewaySession) GetWithdrawalSignatures(_withdrawalId *big.Int, _validators []common.Address) ([][]byte, error) {
	return _Gateway.Contract.GetWithdrawalSignatures(&_Gateway.CallOpts, _withdrawalId, _validators)
}

// GetWithdrawalSignatures is a free data retrieval call binding the contract method 0xecc83649.
//
// Solidity: function getWithdrawalSignatures(uint256 _withdrawalId, address[] _validators) view returns(bytes[] _signatures)
func (_Gateway *GatewayCallerSession) GetWithdrawalSignatures(_withdrawalId *big.Int, _validators []common.Address) ([][]byte, error) {
	return _Gateway.Contract.GetWithdrawalSignatures(&_Gateway.CallOpts, _withdrawalId, _validators)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Gateway *GatewayCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Gateway *GatewaySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Gateway.Contract.HasRole(&_Gateway.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Gateway *GatewayCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Gateway.Contract.HasRole(&_Gateway.CallOpts, role, account)
}

// MainchainWithdrew is a free data retrieval call binding the contract method 0xf668214a.
//
// Solidity: function mainchainWithdrew(uint256 ) view returns(bool)
func (_Gateway *GatewayCaller) MainchainWithdrew(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "mainchainWithdrew", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MainchainWithdrew is a free data retrieval call binding the contract method 0xf668214a.
//
// Solidity: function mainchainWithdrew(uint256 ) view returns(bool)
func (_Gateway *GatewaySession) MainchainWithdrew(arg0 *big.Int) (bool, error) {
	return _Gateway.Contract.MainchainWithdrew(&_Gateway.CallOpts, arg0)
}

// MainchainWithdrew is a free data retrieval call binding the contract method 0xf668214a.
//
// Solidity: function mainchainWithdrew(uint256 ) view returns(bool)
func (_Gateway *GatewayCallerSession) MainchainWithdrew(arg0 *big.Int) (bool, error) {
	return _Gateway.Contract.MainchainWithdrew(&_Gateway.CallOpts, arg0)
}

// MinimumThreshold is a free data retrieval call binding the contract method 0xbc7f0386.
//
// Solidity: function minimumThreshold(address ) view returns(uint256)
func (_Gateway *GatewayCaller) MinimumThreshold(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "minimumThreshold", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumThreshold is a free data retrieval call binding the contract method 0xbc7f0386.
//
// Solidity: function minimumThreshold(address ) view returns(uint256)
func (_Gateway *GatewaySession) MinimumThreshold(arg0 common.Address) (*big.Int, error) {
	return _Gateway.Contract.MinimumThreshold(&_Gateway.CallOpts, arg0)
}

// MinimumThreshold is a free data retrieval call binding the contract method 0xbc7f0386.
//
// Solidity: function minimumThreshold(address ) view returns(uint256)
func (_Gateway *GatewayCallerSession) MinimumThreshold(arg0 common.Address) (*big.Int, error) {
	return _Gateway.Contract.MinimumThreshold(&_Gateway.CallOpts, arg0)
}

// MinimumVoteWeight is a free data retrieval call binding the contract method 0x7de5dedd.
//
// Solidity: function minimumVoteWeight() view returns(uint256)
func (_Gateway *GatewayCaller) MinimumVoteWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "minimumVoteWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumVoteWeight is a free data retrieval call binding the contract method 0x7de5dedd.
//
// Solidity: function minimumVoteWeight() view returns(uint256)
func (_Gateway *GatewaySession) MinimumVoteWeight() (*big.Int, error) {
	return _Gateway.Contract.MinimumVoteWeight(&_Gateway.CallOpts)
}

// MinimumVoteWeight is a free data retrieval call binding the contract method 0x7de5dedd.
//
// Solidity: function minimumVoteWeight() view returns(uint256)
func (_Gateway *GatewayCallerSession) MinimumVoteWeight() (*big.Int, error) {
	return _Gateway.Contract.MinimumVoteWeight(&_Gateway.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Gateway *GatewayCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Gateway *GatewaySession) Nonce() (*big.Int, error) {
	return _Gateway.Contract.Nonce(&_Gateway.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Gateway *GatewayCallerSession) Nonce() (*big.Int, error) {
	return _Gateway.Contract.Nonce(&_Gateway.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Gateway *GatewayCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Gateway *GatewaySession) Paused() (bool, error) {
	return _Gateway.Contract.Paused(&_Gateway.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Gateway *GatewayCallerSession) Paused() (bool, error) {
	return _Gateway.Contract.Paused(&_Gateway.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Gateway *GatewayCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Gateway *GatewaySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Gateway.Contract.SupportsInterface(&_Gateway.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Gateway *GatewayCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Gateway.Contract.SupportsInterface(&_Gateway.CallOpts, interfaceId)
}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_Gateway *GatewayCaller) ValidatorContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "validatorContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_Gateway *GatewaySession) ValidatorContract() (common.Address, error) {
	return _Gateway.Contract.ValidatorContract(&_Gateway.CallOpts)
}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_Gateway *GatewayCallerSession) ValidatorContract() (common.Address, error) {
	return _Gateway.Contract.ValidatorContract(&_Gateway.CallOpts)
}

// Withdrawal is a free data retrieval call binding the contract method 0x835fc6ca.
//
// Solidity: function withdrawal(uint256 ) view returns(uint256 id, uint8 kind, (address,address,uint256) mainchain, (address,address,uint256) ronin, (uint8,uint256,uint256) info)
func (_Gateway *GatewayCaller) Withdrawal(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id        *big.Int
	Kind      uint8
	Mainchain TokenOwner
	Ronin     TokenOwner
	Info      TokenInfo
}, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "withdrawal", arg0)

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
func (_Gateway *GatewaySession) Withdrawal(arg0 *big.Int) (struct {
	Id        *big.Int
	Kind      uint8
	Mainchain TokenOwner
	Ronin     TokenOwner
	Info      TokenInfo
}, error) {
	return _Gateway.Contract.Withdrawal(&_Gateway.CallOpts, arg0)
}

// Withdrawal is a free data retrieval call binding the contract method 0x835fc6ca.
//
// Solidity: function withdrawal(uint256 ) view returns(uint256 id, uint8 kind, (address,address,uint256) mainchain, (address,address,uint256) ronin, (uint8,uint256,uint256) info)
func (_Gateway *GatewayCallerSession) Withdrawal(arg0 *big.Int) (struct {
	Id        *big.Int
	Kind      uint8
	Mainchain TokenOwner
	Ronin     TokenOwner
	Info      TokenInfo
}, error) {
	return _Gateway.Contract.Withdrawal(&_Gateway.CallOpts, arg0)
}

// WithdrawalCount is a free data retrieval call binding the contract method 0x71706cbe.
//
// Solidity: function withdrawalCount() view returns(uint256)
func (_Gateway *GatewayCaller) WithdrawalCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "withdrawalCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalCount is a free data retrieval call binding the contract method 0x71706cbe.
//
// Solidity: function withdrawalCount() view returns(uint256)
func (_Gateway *GatewaySession) WithdrawalCount() (*big.Int, error) {
	return _Gateway.Contract.WithdrawalCount(&_Gateway.CallOpts)
}

// WithdrawalCount is a free data retrieval call binding the contract method 0x71706cbe.
//
// Solidity: function withdrawalCount() view returns(uint256)
func (_Gateway *GatewayCallerSession) WithdrawalCount() (*big.Int, error) {
	return _Gateway.Contract.WithdrawalCount(&_Gateway.CallOpts)
}

// WithdrawalMigrated is a free data retrieval call binding the contract method 0x4f2717c7.
//
// Solidity: function withdrawalMigrated() view returns(bool)
func (_Gateway *GatewayCaller) WithdrawalMigrated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "withdrawalMigrated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawalMigrated is a free data retrieval call binding the contract method 0x4f2717c7.
//
// Solidity: function withdrawalMigrated() view returns(bool)
func (_Gateway *GatewaySession) WithdrawalMigrated() (bool, error) {
	return _Gateway.Contract.WithdrawalMigrated(&_Gateway.CallOpts)
}

// WithdrawalMigrated is a free data retrieval call binding the contract method 0x4f2717c7.
//
// Solidity: function withdrawalMigrated() view returns(bool)
func (_Gateway *GatewayCallerSession) WithdrawalMigrated() (bool, error) {
	return _Gateway.Contract.WithdrawalMigrated(&_Gateway.CallOpts)
}

// BulkAcknowledgeMainchainWithdrew is a paid mutator transaction binding the contract method 0x51050683.
//
// Solidity: function bulkAcknowledgeMainchainWithdrew(uint256[] _withdrawalIds) returns()
func (_Gateway *GatewayTransactor) BulkAcknowledgeMainchainWithdrew(opts *bind.TransactOpts, _withdrawalIds []*big.Int) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "bulkAcknowledgeMainchainWithdrew", _withdrawalIds)
}

// BulkAcknowledgeMainchainWithdrew is a paid mutator transaction binding the contract method 0x51050683.
//
// Solidity: function bulkAcknowledgeMainchainWithdrew(uint256[] _withdrawalIds) returns()
func (_Gateway *GatewaySession) BulkAcknowledgeMainchainWithdrew(_withdrawalIds []*big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.BulkAcknowledgeMainchainWithdrew(&_Gateway.TransactOpts, _withdrawalIds)
}

// BulkAcknowledgeMainchainWithdrew is a paid mutator transaction binding the contract method 0x51050683.
//
// Solidity: function bulkAcknowledgeMainchainWithdrew(uint256[] _withdrawalIds) returns()
func (_Gateway *GatewayTransactorSession) BulkAcknowledgeMainchainWithdrew(_withdrawalIds []*big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.BulkAcknowledgeMainchainWithdrew(&_Gateway.TransactOpts, _withdrawalIds)
}

// BulkDepositFor is a paid mutator transaction binding the contract method 0xc544376d.
//
// Solidity: function bulkDepositFor((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256))[] _receipts) returns()
func (_Gateway *GatewayTransactor) BulkDepositFor(opts *bind.TransactOpts, _receipts []TransferReceipt) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "bulkDepositFor", _receipts)
}

// BulkDepositFor is a paid mutator transaction binding the contract method 0xc544376d.
//
// Solidity: function bulkDepositFor((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256))[] _receipts) returns()
func (_Gateway *GatewaySession) BulkDepositFor(_receipts []TransferReceipt) (*types.Transaction, error) {
	return _Gateway.Contract.BulkDepositFor(&_Gateway.TransactOpts, _receipts)
}

// BulkDepositFor is a paid mutator transaction binding the contract method 0xc544376d.
//
// Solidity: function bulkDepositFor((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256))[] _receipts) returns()
func (_Gateway *GatewayTransactorSession) BulkDepositFor(_receipts []TransferReceipt) (*types.Transaction, error) {
	return _Gateway.Contract.BulkDepositFor(&_Gateway.TransactOpts, _receipts)
}

// BulkRequestWithdrawalFor is a paid mutator transaction binding the contract method 0x5a7dd06a.
//
// Solidity: function bulkRequestWithdrawalFor((address,address,(uint8,uint256,uint256))[] _requests, uint256 _chainId) returns()
func (_Gateway *GatewayTransactor) BulkRequestWithdrawalFor(opts *bind.TransactOpts, _requests []TransferRequest, _chainId *big.Int) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "bulkRequestWithdrawalFor", _requests, _chainId)
}

// BulkRequestWithdrawalFor is a paid mutator transaction binding the contract method 0x5a7dd06a.
//
// Solidity: function bulkRequestWithdrawalFor((address,address,(uint8,uint256,uint256))[] _requests, uint256 _chainId) returns()
func (_Gateway *GatewaySession) BulkRequestWithdrawalFor(_requests []TransferRequest, _chainId *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.BulkRequestWithdrawalFor(&_Gateway.TransactOpts, _requests, _chainId)
}

// BulkRequestWithdrawalFor is a paid mutator transaction binding the contract method 0x5a7dd06a.
//
// Solidity: function bulkRequestWithdrawalFor((address,address,(uint8,uint256,uint256))[] _requests, uint256 _chainId) returns()
func (_Gateway *GatewayTransactorSession) BulkRequestWithdrawalFor(_requests []TransferRequest, _chainId *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.BulkRequestWithdrawalFor(&_Gateway.TransactOpts, _requests, _chainId)
}

// BulkSubmitWithdrawalSignatures is a paid mutator transaction binding the contract method 0xfa389659.
//
// Solidity: function bulkSubmitWithdrawalSignatures(uint256[] _withdrawals, bytes[] _signatures) returns()
func (_Gateway *GatewayTransactor) BulkSubmitWithdrawalSignatures(opts *bind.TransactOpts, _withdrawals []*big.Int, _signatures [][]byte) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "bulkSubmitWithdrawalSignatures", _withdrawals, _signatures)
}

// BulkSubmitWithdrawalSignatures is a paid mutator transaction binding the contract method 0xfa389659.
//
// Solidity: function bulkSubmitWithdrawalSignatures(uint256[] _withdrawals, bytes[] _signatures) returns()
func (_Gateway *GatewaySession) BulkSubmitWithdrawalSignatures(_withdrawals []*big.Int, _signatures [][]byte) (*types.Transaction, error) {
	return _Gateway.Contract.BulkSubmitWithdrawalSignatures(&_Gateway.TransactOpts, _withdrawals, _signatures)
}

// BulkSubmitWithdrawalSignatures is a paid mutator transaction binding the contract method 0xfa389659.
//
// Solidity: function bulkSubmitWithdrawalSignatures(uint256[] _withdrawals, bytes[] _signatures) returns()
func (_Gateway *GatewayTransactorSession) BulkSubmitWithdrawalSignatures(_withdrawals []*big.Int, _signatures [][]byte) (*types.Transaction, error) {
	return _Gateway.Contract.BulkSubmitWithdrawalSignatures(&_Gateway.TransactOpts, _withdrawals, _signatures)
}

// DepositFor is a paid mutator transaction binding the contract method 0x109679ef.
//
// Solidity: function depositFor((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt) returns()
func (_Gateway *GatewayTransactor) DepositFor(opts *bind.TransactOpts, _receipt TransferReceipt) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "depositFor", _receipt)
}

// DepositFor is a paid mutator transaction binding the contract method 0x109679ef.
//
// Solidity: function depositFor((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt) returns()
func (_Gateway *GatewaySession) DepositFor(_receipt TransferReceipt) (*types.Transaction, error) {
	return _Gateway.Contract.DepositFor(&_Gateway.TransactOpts, _receipt)
}

// DepositFor is a paid mutator transaction binding the contract method 0x109679ef.
//
// Solidity: function depositFor((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt) returns()
func (_Gateway *GatewayTransactorSession) DepositFor(_receipt TransferReceipt) (*types.Transaction, error) {
	return _Gateway.Contract.DepositFor(&_Gateway.TransactOpts, _receipt)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Gateway *GatewayTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Gateway *GatewaySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.GrantRole(&_Gateway.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Gateway *GatewayTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.GrantRole(&_Gateway.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x1e5bc9ce.
//
// Solidity: function initialize(address _roleSetter, address _validatorContract, uint256 _numerator, uint256 _denominator, address[] _withdrawalMigrators, address[] _roninTokens, address[] _mainchainTokens, uint256[][2] _packedNumbers) returns()
func (_Gateway *GatewayTransactor) Initialize(opts *bind.TransactOpts, _roleSetter common.Address, _validatorContract common.Address, _numerator *big.Int, _denominator *big.Int, _withdrawalMigrators []common.Address, _roninTokens []common.Address, _mainchainTokens []common.Address, _packedNumbers [2][]*big.Int) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "initialize", _roleSetter, _validatorContract, _numerator, _denominator, _withdrawalMigrators, _roninTokens, _mainchainTokens, _packedNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0x1e5bc9ce.
//
// Solidity: function initialize(address _roleSetter, address _validatorContract, uint256 _numerator, uint256 _denominator, address[] _withdrawalMigrators, address[] _roninTokens, address[] _mainchainTokens, uint256[][2] _packedNumbers) returns()
func (_Gateway *GatewaySession) Initialize(_roleSetter common.Address, _validatorContract common.Address, _numerator *big.Int, _denominator *big.Int, _withdrawalMigrators []common.Address, _roninTokens []common.Address, _mainchainTokens []common.Address, _packedNumbers [2][]*big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.Initialize(&_Gateway.TransactOpts, _roleSetter, _validatorContract, _numerator, _denominator, _withdrawalMigrators, _roninTokens, _mainchainTokens, _packedNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0x1e5bc9ce.
//
// Solidity: function initialize(address _roleSetter, address _validatorContract, uint256 _numerator, uint256 _denominator, address[] _withdrawalMigrators, address[] _roninTokens, address[] _mainchainTokens, uint256[][2] _packedNumbers) returns()
func (_Gateway *GatewayTransactorSession) Initialize(_roleSetter common.Address, _validatorContract common.Address, _numerator *big.Int, _denominator *big.Int, _withdrawalMigrators []common.Address, _roninTokens []common.Address, _mainchainTokens []common.Address, _packedNumbers [2][]*big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.Initialize(&_Gateway.TransactOpts, _roleSetter, _validatorContract, _numerator, _denominator, _withdrawalMigrators, _roninTokens, _mainchainTokens, _packedNumbers)
}

// MapTokens is a paid mutator transaction binding the contract method 0x6598c1bd.
//
// Solidity: function mapTokens(address[] _roninTokens, address[] _mainchainTokens, uint256[] _chainIds) returns()
func (_Gateway *GatewayTransactor) MapTokens(opts *bind.TransactOpts, _roninTokens []common.Address, _mainchainTokens []common.Address, _chainIds []*big.Int) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "mapTokens", _roninTokens, _mainchainTokens, _chainIds)
}

// MapTokens is a paid mutator transaction binding the contract method 0x6598c1bd.
//
// Solidity: function mapTokens(address[] _roninTokens, address[] _mainchainTokens, uint256[] _chainIds) returns()
func (_Gateway *GatewaySession) MapTokens(_roninTokens []common.Address, _mainchainTokens []common.Address, _chainIds []*big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.MapTokens(&_Gateway.TransactOpts, _roninTokens, _mainchainTokens, _chainIds)
}

// MapTokens is a paid mutator transaction binding the contract method 0x6598c1bd.
//
// Solidity: function mapTokens(address[] _roninTokens, address[] _mainchainTokens, uint256[] _chainIds) returns()
func (_Gateway *GatewayTransactorSession) MapTokens(_roninTokens []common.Address, _mainchainTokens []common.Address, _chainIds []*big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.MapTokens(&_Gateway.TransactOpts, _roninTokens, _mainchainTokens, _chainIds)
}

// MarkWithdrawalMigrated is a paid mutator transaction binding the contract method 0x3b5afc22.
//
// Solidity: function markWithdrawalMigrated() returns()
func (_Gateway *GatewayTransactor) MarkWithdrawalMigrated(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "markWithdrawalMigrated")
}

// MarkWithdrawalMigrated is a paid mutator transaction binding the contract method 0x3b5afc22.
//
// Solidity: function markWithdrawalMigrated() returns()
func (_Gateway *GatewaySession) MarkWithdrawalMigrated() (*types.Transaction, error) {
	return _Gateway.Contract.MarkWithdrawalMigrated(&_Gateway.TransactOpts)
}

// MarkWithdrawalMigrated is a paid mutator transaction binding the contract method 0x3b5afc22.
//
// Solidity: function markWithdrawalMigrated() returns()
func (_Gateway *GatewayTransactorSession) MarkWithdrawalMigrated() (*types.Transaction, error) {
	return _Gateway.Contract.MarkWithdrawalMigrated(&_Gateway.TransactOpts)
}

// MigrateWithdrawals is a paid mutator transaction binding the contract method 0x931ec987.
//
// Solidity: function migrateWithdrawals((address,address,(uint8,uint256,uint256))[] _requests, address[] _requesters) returns()
func (_Gateway *GatewayTransactor) MigrateWithdrawals(opts *bind.TransactOpts, _requests []TransferRequest, _requesters []common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "migrateWithdrawals", _requests, _requesters)
}

// MigrateWithdrawals is a paid mutator transaction binding the contract method 0x931ec987.
//
// Solidity: function migrateWithdrawals((address,address,(uint8,uint256,uint256))[] _requests, address[] _requesters) returns()
func (_Gateway *GatewaySession) MigrateWithdrawals(_requests []TransferRequest, _requesters []common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.MigrateWithdrawals(&_Gateway.TransactOpts, _requests, _requesters)
}

// MigrateWithdrawals is a paid mutator transaction binding the contract method 0x931ec987.
//
// Solidity: function migrateWithdrawals((address,address,(uint8,uint256,uint256))[] _requests, address[] _requesters) returns()
func (_Gateway *GatewayTransactorSession) MigrateWithdrawals(_requests []TransferRequest, _requesters []common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.MigrateWithdrawals(&_Gateway.TransactOpts, _requests, _requesters)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Gateway *GatewayTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Gateway *GatewaySession) Pause() (*types.Transaction, error) {
	return _Gateway.Contract.Pause(&_Gateway.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Gateway *GatewayTransactorSession) Pause() (*types.Transaction, error) {
	return _Gateway.Contract.Pause(&_Gateway.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Gateway *GatewayTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Gateway *GatewaySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.RenounceRole(&_Gateway.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Gateway *GatewayTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.RenounceRole(&_Gateway.TransactOpts, role, account)
}

// RequestWithdrawalFor is a paid mutator transaction binding the contract method 0x0b1ff17f.
//
// Solidity: function requestWithdrawalFor((address,address,(uint8,uint256,uint256)) _request, uint256 _chainId) returns()
func (_Gateway *GatewayTransactor) RequestWithdrawalFor(opts *bind.TransactOpts, _request TransferRequest, _chainId *big.Int) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "requestWithdrawalFor", _request, _chainId)
}

// RequestWithdrawalFor is a paid mutator transaction binding the contract method 0x0b1ff17f.
//
// Solidity: function requestWithdrawalFor((address,address,(uint8,uint256,uint256)) _request, uint256 _chainId) returns()
func (_Gateway *GatewaySession) RequestWithdrawalFor(_request TransferRequest, _chainId *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.RequestWithdrawalFor(&_Gateway.TransactOpts, _request, _chainId)
}

// RequestWithdrawalFor is a paid mutator transaction binding the contract method 0x0b1ff17f.
//
// Solidity: function requestWithdrawalFor((address,address,(uint8,uint256,uint256)) _request, uint256 _chainId) returns()
func (_Gateway *GatewayTransactorSession) RequestWithdrawalFor(_request TransferRequest, _chainId *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.RequestWithdrawalFor(&_Gateway.TransactOpts, _request, _chainId)
}

// RequestWithdrawalSignatures is a paid mutator transaction binding the contract method 0x47b56b2c.
//
// Solidity: function requestWithdrawalSignatures(uint256 _withdrawalId) returns()
func (_Gateway *GatewayTransactor) RequestWithdrawalSignatures(opts *bind.TransactOpts, _withdrawalId *big.Int) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "requestWithdrawalSignatures", _withdrawalId)
}

// RequestWithdrawalSignatures is a paid mutator transaction binding the contract method 0x47b56b2c.
//
// Solidity: function requestWithdrawalSignatures(uint256 _withdrawalId) returns()
func (_Gateway *GatewaySession) RequestWithdrawalSignatures(_withdrawalId *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.RequestWithdrawalSignatures(&_Gateway.TransactOpts, _withdrawalId)
}

// RequestWithdrawalSignatures is a paid mutator transaction binding the contract method 0x47b56b2c.
//
// Solidity: function requestWithdrawalSignatures(uint256 _withdrawalId) returns()
func (_Gateway *GatewayTransactorSession) RequestWithdrawalSignatures(_withdrawalId *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.RequestWithdrawalSignatures(&_Gateway.TransactOpts, _withdrawalId)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Gateway *GatewayTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Gateway *GatewaySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.RevokeRole(&_Gateway.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Gateway *GatewayTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.RevokeRole(&_Gateway.TransactOpts, role, account)
}

// SetMinimumThresholds is a paid mutator transaction binding the contract method 0x64363f78.
//
// Solidity: function setMinimumThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_Gateway *GatewayTransactor) SetMinimumThresholds(opts *bind.TransactOpts, _tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "setMinimumThresholds", _tokens, _thresholds)
}

// SetMinimumThresholds is a paid mutator transaction binding the contract method 0x64363f78.
//
// Solidity: function setMinimumThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_Gateway *GatewaySession) SetMinimumThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.SetMinimumThresholds(&_Gateway.TransactOpts, _tokens, _thresholds)
}

// SetMinimumThresholds is a paid mutator transaction binding the contract method 0x64363f78.
//
// Solidity: function setMinimumThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_Gateway *GatewayTransactorSession) SetMinimumThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.SetMinimumThresholds(&_Gateway.TransactOpts, _tokens, _thresholds)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 _numerator, uint256 _denominator) returns(uint256, uint256)
func (_Gateway *GatewayTransactor) SetThreshold(opts *bind.TransactOpts, _numerator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "setThreshold", _numerator, _denominator)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 _numerator, uint256 _denominator) returns(uint256, uint256)
func (_Gateway *GatewaySession) SetThreshold(_numerator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.SetThreshold(&_Gateway.TransactOpts, _numerator, _denominator)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 _numerator, uint256 _denominator) returns(uint256, uint256)
func (_Gateway *GatewayTransactorSession) SetThreshold(_numerator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.SetThreshold(&_Gateway.TransactOpts, _numerator, _denominator)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _validatorContract) returns()
func (_Gateway *GatewayTransactor) SetValidatorContract(opts *bind.TransactOpts, _validatorContract common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "setValidatorContract", _validatorContract)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _validatorContract) returns()
func (_Gateway *GatewaySession) SetValidatorContract(_validatorContract common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.SetValidatorContract(&_Gateway.TransactOpts, _validatorContract)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _validatorContract) returns()
func (_Gateway *GatewayTransactorSession) SetValidatorContract(_validatorContract common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.SetValidatorContract(&_Gateway.TransactOpts, _validatorContract)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Gateway *GatewayTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Gateway *GatewaySession) Unpause() (*types.Transaction, error) {
	return _Gateway.Contract.Unpause(&_Gateway.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Gateway *GatewayTransactorSession) Unpause() (*types.Transaction, error) {
	return _Gateway.Contract.Unpause(&_Gateway.TransactOpts)
}

// GatewayDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Gateway contract.
type GatewayDepositedIterator struct {
	Event *GatewayDeposited // Event containing the contract specifics and raw log

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
func (it *GatewayDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayDeposited)
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
		it.Event = new(GatewayDeposited)
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
func (it *GatewayDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayDeposited represents a Deposited event raised by the Gateway contract.
type GatewayDeposited struct {
	ReceiptHash [32]byte
	Receipt     TransferReceipt
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x8d20d8121a34dded9035ff5b43e901c142824f7a22126392992c353c37890524.
//
// Solidity: event Deposited(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_Gateway *GatewayFilterer) FilterDeposited(opts *bind.FilterOpts) (*GatewayDepositedIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &GatewayDepositedIterator{contract: _Gateway.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x8d20d8121a34dded9035ff5b43e901c142824f7a22126392992c353c37890524.
//
// Solidity: event Deposited(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_Gateway *GatewayFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *GatewayDeposited) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayDeposited)
				if err := _Gateway.contract.UnpackLog(event, "Deposited", log); err != nil {
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
func (_Gateway *GatewayFilterer) ParseDeposited(log types.Log) (*GatewayDeposited, error) {
	event := new(GatewayDeposited)
	if err := _Gateway.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMainchainWithdrewIterator is returned from FilterMainchainWithdrew and is used to iterate over the raw logs and unpacked data for MainchainWithdrew events raised by the Gateway contract.
type GatewayMainchainWithdrewIterator struct {
	Event *GatewayMainchainWithdrew // Event containing the contract specifics and raw log

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
func (it *GatewayMainchainWithdrewIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMainchainWithdrew)
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
		it.Event = new(GatewayMainchainWithdrew)
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
func (it *GatewayMainchainWithdrewIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMainchainWithdrewIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMainchainWithdrew represents a MainchainWithdrew event raised by the Gateway contract.
type GatewayMainchainWithdrew struct {
	ReceiptHash [32]byte
	Receipt     TransferReceipt
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMainchainWithdrew is a free log retrieval operation binding the contract event 0x62520d049932cdee872e9b3c59c0f6073637147e5e9bc8b050b062430eaf5c9f.
//
// Solidity: event MainchainWithdrew(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_Gateway *GatewayFilterer) FilterMainchainWithdrew(opts *bind.FilterOpts) (*GatewayMainchainWithdrewIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "MainchainWithdrew")
	if err != nil {
		return nil, err
	}
	return &GatewayMainchainWithdrewIterator{contract: _Gateway.contract, event: "MainchainWithdrew", logs: logs, sub: sub}, nil
}

// WatchMainchainWithdrew is a free log subscription operation binding the contract event 0x62520d049932cdee872e9b3c59c0f6073637147e5e9bc8b050b062430eaf5c9f.
//
// Solidity: event MainchainWithdrew(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_Gateway *GatewayFilterer) WatchMainchainWithdrew(opts *bind.WatchOpts, sink chan<- *GatewayMainchainWithdrew) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "MainchainWithdrew")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMainchainWithdrew)
				if err := _Gateway.contract.UnpackLog(event, "MainchainWithdrew", log); err != nil {
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

// ParseMainchainWithdrew is a log parse operation binding the contract event 0x62520d049932cdee872e9b3c59c0f6073637147e5e9bc8b050b062430eaf5c9f.
//
// Solidity: event MainchainWithdrew(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_Gateway *GatewayFilterer) ParseMainchainWithdrew(log types.Log) (*GatewayMainchainWithdrew, error) {
	event := new(GatewayMainchainWithdrew)
	if err := _Gateway.contract.UnpackLog(event, "MainchainWithdrew", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMinimumThresholdsUpdatedIterator is returned from FilterMinimumThresholdsUpdated and is used to iterate over the raw logs and unpacked data for MinimumThresholdsUpdated events raised by the Gateway contract.
type GatewayMinimumThresholdsUpdatedIterator struct {
	Event *GatewayMinimumThresholdsUpdated // Event containing the contract specifics and raw log

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
func (it *GatewayMinimumThresholdsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMinimumThresholdsUpdated)
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
		it.Event = new(GatewayMinimumThresholdsUpdated)
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
func (it *GatewayMinimumThresholdsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMinimumThresholdsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMinimumThresholdsUpdated represents a MinimumThresholdsUpdated event raised by the Gateway contract.
type GatewayMinimumThresholdsUpdated struct {
	Tokens    []common.Address
	Threshold []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinimumThresholdsUpdated is a free log retrieval operation binding the contract event 0x6f52f53a938df83439fa4c6055c7df0a6906d621aa6dfa4708187037fdfc41da.
//
// Solidity: event MinimumThresholdsUpdated(address[] tokens, uint256[] threshold)
func (_Gateway *GatewayFilterer) FilterMinimumThresholdsUpdated(opts *bind.FilterOpts) (*GatewayMinimumThresholdsUpdatedIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "MinimumThresholdsUpdated")
	if err != nil {
		return nil, err
	}
	return &GatewayMinimumThresholdsUpdatedIterator{contract: _Gateway.contract, event: "MinimumThresholdsUpdated", logs: logs, sub: sub}, nil
}

// WatchMinimumThresholdsUpdated is a free log subscription operation binding the contract event 0x6f52f53a938df83439fa4c6055c7df0a6906d621aa6dfa4708187037fdfc41da.
//
// Solidity: event MinimumThresholdsUpdated(address[] tokens, uint256[] threshold)
func (_Gateway *GatewayFilterer) WatchMinimumThresholdsUpdated(opts *bind.WatchOpts, sink chan<- *GatewayMinimumThresholdsUpdated) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "MinimumThresholdsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMinimumThresholdsUpdated)
				if err := _Gateway.contract.UnpackLog(event, "MinimumThresholdsUpdated", log); err != nil {
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
func (_Gateway *GatewayFilterer) ParseMinimumThresholdsUpdated(log types.Log) (*GatewayMinimumThresholdsUpdated, error) {
	event := new(GatewayMinimumThresholdsUpdated)
	if err := _Gateway.contract.UnpackLog(event, "MinimumThresholdsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Gateway contract.
type GatewayPausedIterator struct {
	Event *GatewayPaused // Event containing the contract specifics and raw log

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
func (it *GatewayPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayPaused)
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
		it.Event = new(GatewayPaused)
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
func (it *GatewayPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayPaused represents a Paused event raised by the Gateway contract.
type GatewayPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Gateway *GatewayFilterer) FilterPaused(opts *bind.FilterOpts) (*GatewayPausedIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &GatewayPausedIterator{contract: _Gateway.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Gateway *GatewayFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *GatewayPaused) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayPaused)
				if err := _Gateway.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Gateway *GatewayFilterer) ParsePaused(log types.Log) (*GatewayPaused, error) {
	event := new(GatewayPaused)
	if err := _Gateway.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Gateway contract.
type GatewayRoleAdminChangedIterator struct {
	Event *GatewayRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *GatewayRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayRoleAdminChanged)
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
		it.Event = new(GatewayRoleAdminChanged)
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
func (it *GatewayRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayRoleAdminChanged represents a RoleAdminChanged event raised by the Gateway contract.
type GatewayRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Gateway *GatewayFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*GatewayRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &GatewayRoleAdminChangedIterator{contract: _Gateway.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Gateway *GatewayFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *GatewayRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayRoleAdminChanged)
				if err := _Gateway.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Gateway *GatewayFilterer) ParseRoleAdminChanged(log types.Log) (*GatewayRoleAdminChanged, error) {
	event := new(GatewayRoleAdminChanged)
	if err := _Gateway.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Gateway contract.
type GatewayRoleGrantedIterator struct {
	Event *GatewayRoleGranted // Event containing the contract specifics and raw log

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
func (it *GatewayRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayRoleGranted)
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
		it.Event = new(GatewayRoleGranted)
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
func (it *GatewayRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayRoleGranted represents a RoleGranted event raised by the Gateway contract.
type GatewayRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Gateway *GatewayFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GatewayRoleGrantedIterator, error) {

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

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GatewayRoleGrantedIterator{contract: _Gateway.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Gateway *GatewayFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *GatewayRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayRoleGranted)
				if err := _Gateway.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Gateway *GatewayFilterer) ParseRoleGranted(log types.Log) (*GatewayRoleGranted, error) {
	event := new(GatewayRoleGranted)
	if err := _Gateway.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Gateway contract.
type GatewayRoleRevokedIterator struct {
	Event *GatewayRoleRevoked // Event containing the contract specifics and raw log

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
func (it *GatewayRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayRoleRevoked)
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
		it.Event = new(GatewayRoleRevoked)
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
func (it *GatewayRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayRoleRevoked represents a RoleRevoked event raised by the Gateway contract.
type GatewayRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Gateway *GatewayFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GatewayRoleRevokedIterator, error) {

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

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GatewayRoleRevokedIterator{contract: _Gateway.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Gateway *GatewayFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *GatewayRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayRoleRevoked)
				if err := _Gateway.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Gateway *GatewayFilterer) ParseRoleRevoked(log types.Log) (*GatewayRoleRevoked, error) {
	event := new(GatewayRoleRevoked)
	if err := _Gateway.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayThresholdUpdatedIterator is returned from FilterThresholdUpdated and is used to iterate over the raw logs and unpacked data for ThresholdUpdated events raised by the Gateway contract.
type GatewayThresholdUpdatedIterator struct {
	Event *GatewayThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *GatewayThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayThresholdUpdated)
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
		it.Event = new(GatewayThresholdUpdated)
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
func (it *GatewayThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayThresholdUpdated represents a ThresholdUpdated event raised by the Gateway contract.
type GatewayThresholdUpdated struct {
	Nonce               *big.Int
	Numerator           *big.Int
	Denominator         *big.Int
	PreviousNumerator   *big.Int
	PreviousDenominator *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterThresholdUpdated is a free log retrieval operation binding the contract event 0x976f8a9c5bdf8248dec172376d6e2b80a8e3df2f0328e381c6db8e1cf138c0f8.
//
// Solidity: event ThresholdUpdated(uint256 indexed nonce, uint256 indexed numerator, uint256 indexed denominator, uint256 previousNumerator, uint256 previousDenominator)
func (_Gateway *GatewayFilterer) FilterThresholdUpdated(opts *bind.FilterOpts, nonce []*big.Int, numerator []*big.Int, denominator []*big.Int) (*GatewayThresholdUpdatedIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var numeratorRule []interface{}
	for _, numeratorItem := range numerator {
		numeratorRule = append(numeratorRule, numeratorItem)
	}
	var denominatorRule []interface{}
	for _, denominatorItem := range denominator {
		denominatorRule = append(denominatorRule, denominatorItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "ThresholdUpdated", nonceRule, numeratorRule, denominatorRule)
	if err != nil {
		return nil, err
	}
	return &GatewayThresholdUpdatedIterator{contract: _Gateway.contract, event: "ThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchThresholdUpdated is a free log subscription operation binding the contract event 0x976f8a9c5bdf8248dec172376d6e2b80a8e3df2f0328e381c6db8e1cf138c0f8.
//
// Solidity: event ThresholdUpdated(uint256 indexed nonce, uint256 indexed numerator, uint256 indexed denominator, uint256 previousNumerator, uint256 previousDenominator)
func (_Gateway *GatewayFilterer) WatchThresholdUpdated(opts *bind.WatchOpts, sink chan<- *GatewayThresholdUpdated, nonce []*big.Int, numerator []*big.Int, denominator []*big.Int) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var numeratorRule []interface{}
	for _, numeratorItem := range numerator {
		numeratorRule = append(numeratorRule, numeratorItem)
	}
	var denominatorRule []interface{}
	for _, denominatorItem := range denominator {
		denominatorRule = append(denominatorRule, denominatorItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "ThresholdUpdated", nonceRule, numeratorRule, denominatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayThresholdUpdated)
				if err := _Gateway.contract.UnpackLog(event, "ThresholdUpdated", log); err != nil {
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
// Solidity: event ThresholdUpdated(uint256 indexed nonce, uint256 indexed numerator, uint256 indexed denominator, uint256 previousNumerator, uint256 previousDenominator)
func (_Gateway *GatewayFilterer) ParseThresholdUpdated(log types.Log) (*GatewayThresholdUpdated, error) {
	event := new(GatewayThresholdUpdated)
	if err := _Gateway.contract.UnpackLog(event, "ThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayTokenMappedIterator is returned from FilterTokenMapped and is used to iterate over the raw logs and unpacked data for TokenMapped events raised by the Gateway contract.
type GatewayTokenMappedIterator struct {
	Event *GatewayTokenMapped // Event containing the contract specifics and raw log

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
func (it *GatewayTokenMappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayTokenMapped)
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
		it.Event = new(GatewayTokenMapped)
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
func (it *GatewayTokenMappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayTokenMappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayTokenMapped represents a TokenMapped event raised by the Gateway contract.
type GatewayTokenMapped struct {
	RoninTokens     []common.Address
	MainchainTokens []common.Address
	ChainIds        []*big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokenMapped is a free log retrieval operation binding the contract event 0x332d1046c291868dcf6424605d8c4f46a10977395d49f764b86b8d441ba82bd3.
//
// Solidity: event TokenMapped(address[] roninTokens, address[] mainchainTokens, uint256[] chainIds)
func (_Gateway *GatewayFilterer) FilterTokenMapped(opts *bind.FilterOpts) (*GatewayTokenMappedIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "TokenMapped")
	if err != nil {
		return nil, err
	}
	return &GatewayTokenMappedIterator{contract: _Gateway.contract, event: "TokenMapped", logs: logs, sub: sub}, nil
}

// WatchTokenMapped is a free log subscription operation binding the contract event 0x332d1046c291868dcf6424605d8c4f46a10977395d49f764b86b8d441ba82bd3.
//
// Solidity: event TokenMapped(address[] roninTokens, address[] mainchainTokens, uint256[] chainIds)
func (_Gateway *GatewayFilterer) WatchTokenMapped(opts *bind.WatchOpts, sink chan<- *GatewayTokenMapped) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "TokenMapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayTokenMapped)
				if err := _Gateway.contract.UnpackLog(event, "TokenMapped", log); err != nil {
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
func (_Gateway *GatewayFilterer) ParseTokenMapped(log types.Log) (*GatewayTokenMapped, error) {
	event := new(GatewayTokenMapped)
	if err := _Gateway.contract.UnpackLog(event, "TokenMapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Gateway contract.
type GatewayUnpausedIterator struct {
	Event *GatewayUnpaused // Event containing the contract specifics and raw log

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
func (it *GatewayUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayUnpaused)
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
		it.Event = new(GatewayUnpaused)
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
func (it *GatewayUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayUnpaused represents a Unpaused event raised by the Gateway contract.
type GatewayUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Gateway *GatewayFilterer) FilterUnpaused(opts *bind.FilterOpts) (*GatewayUnpausedIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &GatewayUnpausedIterator{contract: _Gateway.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Gateway *GatewayFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *GatewayUnpaused) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayUnpaused)
				if err := _Gateway.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Gateway *GatewayFilterer) ParseUnpaused(log types.Log) (*GatewayUnpaused, error) {
	event := new(GatewayUnpaused)
	if err := _Gateway.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayValidatorContractUpdatedIterator is returned from FilterValidatorContractUpdated and is used to iterate over the raw logs and unpacked data for ValidatorContractUpdated events raised by the Gateway contract.
type GatewayValidatorContractUpdatedIterator struct {
	Event *GatewayValidatorContractUpdated // Event containing the contract specifics and raw log

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
func (it *GatewayValidatorContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayValidatorContractUpdated)
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
		it.Event = new(GatewayValidatorContractUpdated)
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
func (it *GatewayValidatorContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayValidatorContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayValidatorContractUpdated represents a ValidatorContractUpdated event raised by the Gateway contract.
type GatewayValidatorContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterValidatorContractUpdated is a free log retrieval operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_Gateway *GatewayFilterer) FilterValidatorContractUpdated(opts *bind.FilterOpts) (*GatewayValidatorContractUpdatedIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "ValidatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return &GatewayValidatorContractUpdatedIterator{contract: _Gateway.contract, event: "ValidatorContractUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorContractUpdated is a free log subscription operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_Gateway *GatewayFilterer) WatchValidatorContractUpdated(opts *bind.WatchOpts, sink chan<- *GatewayValidatorContractUpdated) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "ValidatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayValidatorContractUpdated)
				if err := _Gateway.contract.UnpackLog(event, "ValidatorContractUpdated", log); err != nil {
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
func (_Gateway *GatewayFilterer) ParseValidatorContractUpdated(log types.Log) (*GatewayValidatorContractUpdated, error) {
	event := new(GatewayValidatorContractUpdated)
	if err := _Gateway.contract.UnpackLog(event, "ValidatorContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayWithdrawalRequestedIterator is returned from FilterWithdrawalRequested and is used to iterate over the raw logs and unpacked data for WithdrawalRequested events raised by the Gateway contract.
type GatewayWithdrawalRequestedIterator struct {
	Event *GatewayWithdrawalRequested // Event containing the contract specifics and raw log

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
func (it *GatewayWithdrawalRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWithdrawalRequested)
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
		it.Event = new(GatewayWithdrawalRequested)
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
func (it *GatewayWithdrawalRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWithdrawalRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWithdrawalRequested represents a WithdrawalRequested event raised by the Gateway contract.
type GatewayWithdrawalRequested struct {
	ReceiptHash [32]byte
	Arg1        TransferReceipt
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalRequested is a free log retrieval operation binding the contract event 0xf313c253a5be72c29d0deb2c8768a9543744ac03d6b3cafd50cc976f1c2632fc.
//
// Solidity: event WithdrawalRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) arg1)
func (_Gateway *GatewayFilterer) FilterWithdrawalRequested(opts *bind.FilterOpts) (*GatewayWithdrawalRequestedIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "WithdrawalRequested")
	if err != nil {
		return nil, err
	}
	return &GatewayWithdrawalRequestedIterator{contract: _Gateway.contract, event: "WithdrawalRequested", logs: logs, sub: sub}, nil
}

// WatchWithdrawalRequested is a free log subscription operation binding the contract event 0xf313c253a5be72c29d0deb2c8768a9543744ac03d6b3cafd50cc976f1c2632fc.
//
// Solidity: event WithdrawalRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) arg1)
func (_Gateway *GatewayFilterer) WatchWithdrawalRequested(opts *bind.WatchOpts, sink chan<- *GatewayWithdrawalRequested) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "WithdrawalRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWithdrawalRequested)
				if err := _Gateway.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
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
func (_Gateway *GatewayFilterer) ParseWithdrawalRequested(log types.Log) (*GatewayWithdrawalRequested, error) {
	event := new(GatewayWithdrawalRequested)
	if err := _Gateway.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayWithdrawalSignaturesRequestedIterator is returned from FilterWithdrawalSignaturesRequested and is used to iterate over the raw logs and unpacked data for WithdrawalSignaturesRequested events raised by the Gateway contract.
type GatewayWithdrawalSignaturesRequestedIterator struct {
	Event *GatewayWithdrawalSignaturesRequested // Event containing the contract specifics and raw log

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
func (it *GatewayWithdrawalSignaturesRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWithdrawalSignaturesRequested)
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
		it.Event = new(GatewayWithdrawalSignaturesRequested)
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
func (it *GatewayWithdrawalSignaturesRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWithdrawalSignaturesRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWithdrawalSignaturesRequested represents a WithdrawalSignaturesRequested event raised by the Gateway contract.
type GatewayWithdrawalSignaturesRequested struct {
	ReceiptHash [32]byte
	Arg1        TransferReceipt
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalSignaturesRequested is a free log retrieval operation binding the contract event 0x04e8cbd836dea43a2dc7eb19de345cca3a8e6978a2ef5225d924775500f67c7c.
//
// Solidity: event WithdrawalSignaturesRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) arg1)
func (_Gateway *GatewayFilterer) FilterWithdrawalSignaturesRequested(opts *bind.FilterOpts) (*GatewayWithdrawalSignaturesRequestedIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "WithdrawalSignaturesRequested")
	if err != nil {
		return nil, err
	}
	return &GatewayWithdrawalSignaturesRequestedIterator{contract: _Gateway.contract, event: "WithdrawalSignaturesRequested", logs: logs, sub: sub}, nil
}

// WatchWithdrawalSignaturesRequested is a free log subscription operation binding the contract event 0x04e8cbd836dea43a2dc7eb19de345cca3a8e6978a2ef5225d924775500f67c7c.
//
// Solidity: event WithdrawalSignaturesRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) arg1)
func (_Gateway *GatewayFilterer) WatchWithdrawalSignaturesRequested(opts *bind.WatchOpts, sink chan<- *GatewayWithdrawalSignaturesRequested) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "WithdrawalSignaturesRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWithdrawalSignaturesRequested)
				if err := _Gateway.contract.UnpackLog(event, "WithdrawalSignaturesRequested", log); err != nil {
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
func (_Gateway *GatewayFilterer) ParseWithdrawalSignaturesRequested(log types.Log) (*GatewayWithdrawalSignaturesRequested, error) {
	event := new(GatewayWithdrawalSignaturesRequested)
	if err := _Gateway.contract.UnpackLog(event, "WithdrawalSignaturesRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
