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

// SignatureConsumerSignature is an auto generated low-level Go binding around an user-defined struct.
type SignatureConsumerSignature struct {
	V uint8
	R [32]byte
	S [32]byte
}

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

// EthereumMetaData contains all meta data concerning the Ethereum contract.
var EthereumMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"limits\",\"type\":\"uint256[]\"}],\"name\":\"DailyWithdrawalLimitsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"receiptHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"DepositRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"thresholds\",\"type\":\"uint256[]\"}],\"name\":\"FullSigsThresholdsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"thresholds\",\"type\":\"uint256[]\"}],\"name\":\"LockedThresholdsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_numerator\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_denominator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_previousNumerator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_previousDenominator\",\"type\":\"uint256\"}],\"name\":\"ThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"mainchainTokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"roninTokens\",\"type\":\"address[]\"}],\"name\":\"TokenMapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIWeightedValidator\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ValidatorContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"WithdrawalLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"WithdrawalUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"receiptHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"Withdrew\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIWETH\",\"name\":\"weth\",\"type\":\"address\"}],\"name\":\"WrappedNativeTokenContractUpdated\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWAL_UNLOCKER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipientAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structTransfer.Request[]\",\"name\":\"_requests\",\"type\":\"tuple[]\"}],\"name\":\"bulkRequestDepositFor\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_voteWeight\",\"type\":\"uint256\"}],\"name\":\"checkThreshold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"dailyWithdrawalLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"fullSigThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mainchainToken\",\"type\":\"address\"}],\"name\":\"getRoninToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_roleSetter\",\"type\":\"address\"},{\"internalType\":\"contractIWETH\",\"name\":\"_wrappedToken\",\"type\":\"address\"},{\"internalType\":\"contractIWeightedValidator\",\"name\":\"_validatorContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_roninChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_denominator\",\"type\":\"uint256\"},{\"internalType\":\"address[][3]\",\"name\":\"_addresses\",\"type\":\"address[][3]\"},{\"internalType\":\"uint256[][3]\",\"name\":\"_thresholds\",\"type\":\"uint256[][3]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastDateSynced\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastSyncedWithdrawal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockedThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_mainchainTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_roninTokens\",\"type\":\"address[]\"}],\"name\":\"mapTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_mainchainTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_roninTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_fullSigThreshold\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_lockedThreshold\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_dailyWithdrawalLimit\",\"type\":\"uint256[]\"}],\"name\":\"mapTokensAndThresholds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumVoteWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"}],\"name\":\"reachedWithdrawalLimit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveEther\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipientAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structTransfer.Request\",\"name\":\"_request\",\"type\":\"tuple\"}],\"name\":\"requestDepositFor\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roninChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_limits\",\"type\":\"uint256[]\"}],\"name\":\"setDailyWithdrawalLimits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_thresholds\",\"type\":\"uint256[]\"}],\"name\":\"setFullSigsThresholds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_thresholds\",\"type\":\"uint256[]\"}],\"name\":\"setLockedThresholds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_denominator\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIWeightedValidator\",\"name\":\"_validatorContract\",\"type\":\"address\"}],\"name\":\"setValidatorContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"_wrappedToken\",\"type\":\"address\"}],\"name\":\"setWrappedNativeTokenContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structTransfer.Receipt\",\"name\":\"_receipt\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"name\":\"submitWithdrawal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_locked\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structTransfer.Receipt\",\"name\":\"_receipt\",\"type\":\"tuple\"}],\"name\":\"unlockWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorContract\",\"outputs\":[{\"internalType\":\"contractIWeightedValidator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawalHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawalLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrappedNativeToken\",\"outputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// EthereumABI is the input ABI used to generate the binding from.
// Deprecated: Use EthereumMetaData.ABI instead.
var EthereumABI = EthereumMetaData.ABI

// Ethereum is an auto generated Go binding around an Ethereum contract.
type Ethereum struct {
	EthereumCaller     // Read-only binding to the contract
	EthereumTransactor // Write-only binding to the contract
	EthereumFilterer   // Log filterer for contract events
}

// EthereumCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthereumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthereumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthereumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthereumSession struct {
	Contract     *Ethereum         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthereumCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthereumCallerSession struct {
	Contract *EthereumCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// EthereumTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthereumTransactorSession struct {
	Contract     *EthereumTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// EthereumRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthereumRaw struct {
	Contract *Ethereum // Generic contract binding to access the raw methods on
}

// EthereumCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthereumCallerRaw struct {
	Contract *EthereumCaller // Generic read-only contract binding to access the raw methods on
}

// EthereumTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthereumTransactorRaw struct {
	Contract *EthereumTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthereum creates a new instance of Ethereum, bound to a specific deployed contract.
func NewEthereum(address common.Address, backend bind.ContractBackend) (*Ethereum, error) {
	contract, err := bindEthereum(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ethereum{EthereumCaller: EthereumCaller{contract: contract}, EthereumTransactor: EthereumTransactor{contract: contract}, EthereumFilterer: EthereumFilterer{contract: contract}}, nil
}

// NewEthereumCaller creates a new read-only instance of Ethereum, bound to a specific deployed contract.
func NewEthereumCaller(address common.Address, caller bind.ContractCaller) (*EthereumCaller, error) {
	contract, err := bindEthereum(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthereumCaller{contract: contract}, nil
}

// NewEthereumTransactor creates a new write-only instance of Ethereum, bound to a specific deployed contract.
func NewEthereumTransactor(address common.Address, transactor bind.ContractTransactor) (*EthereumTransactor, error) {
	contract, err := bindEthereum(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthereumTransactor{contract: contract}, nil
}

// NewEthereumFilterer creates a new log filterer instance of Ethereum, bound to a specific deployed contract.
func NewEthereumFilterer(address common.Address, filterer bind.ContractFilterer) (*EthereumFilterer, error) {
	contract, err := bindEthereum(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthereumFilterer{contract: contract}, nil
}

// bindEthereum binds a generic wrapper to an already deployed contract.
func bindEthereum(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthereumABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethereum *EthereumRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ethereum.Contract.EthereumCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethereum *EthereumRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethereum.Contract.EthereumTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethereum *EthereumRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethereum.Contract.EthereumTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethereum *EthereumCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ethereum.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethereum *EthereumTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethereum.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethereum *EthereumTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethereum.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Ethereum *EthereumCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Ethereum *EthereumSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Ethereum.Contract.DEFAULTADMINROLE(&_Ethereum.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Ethereum *EthereumCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Ethereum.Contract.DEFAULTADMINROLE(&_Ethereum.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Ethereum *EthereumCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Ethereum *EthereumSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Ethereum.Contract.DOMAINSEPARATOR(&_Ethereum.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Ethereum *EthereumCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Ethereum.Contract.DOMAINSEPARATOR(&_Ethereum.CallOpts)
}

// WITHDRAWALUNLOCKERROLE is a free data retrieval call binding the contract method 0x8f34e347.
//
// Solidity: function WITHDRAWAL_UNLOCKER_ROLE() view returns(bytes32)
func (_Ethereum *EthereumCaller) WITHDRAWALUNLOCKERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "WITHDRAWAL_UNLOCKER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWALUNLOCKERROLE is a free data retrieval call binding the contract method 0x8f34e347.
//
// Solidity: function WITHDRAWAL_UNLOCKER_ROLE() view returns(bytes32)
func (_Ethereum *EthereumSession) WITHDRAWALUNLOCKERROLE() ([32]byte, error) {
	return _Ethereum.Contract.WITHDRAWALUNLOCKERROLE(&_Ethereum.CallOpts)
}

// WITHDRAWALUNLOCKERROLE is a free data retrieval call binding the contract method 0x8f34e347.
//
// Solidity: function WITHDRAWAL_UNLOCKER_ROLE() view returns(bytes32)
func (_Ethereum *EthereumCallerSession) WITHDRAWALUNLOCKERROLE() ([32]byte, error) {
	return _Ethereum.Contract.WITHDRAWALUNLOCKERROLE(&_Ethereum.CallOpts)
}

// CheckThreshold is a free data retrieval call binding the contract method 0xdafae408.
//
// Solidity: function checkThreshold(uint256 _voteWeight) view returns(bool)
func (_Ethereum *EthereumCaller) CheckThreshold(opts *bind.CallOpts, _voteWeight *big.Int) (bool, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "checkThreshold", _voteWeight)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckThreshold is a free data retrieval call binding the contract method 0xdafae408.
//
// Solidity: function checkThreshold(uint256 _voteWeight) view returns(bool)
func (_Ethereum *EthereumSession) CheckThreshold(_voteWeight *big.Int) (bool, error) {
	return _Ethereum.Contract.CheckThreshold(&_Ethereum.CallOpts, _voteWeight)
}

// CheckThreshold is a free data retrieval call binding the contract method 0xdafae408.
//
// Solidity: function checkThreshold(uint256 _voteWeight) view returns(bool)
func (_Ethereum *EthereumCallerSession) CheckThreshold(_voteWeight *big.Int) (bool, error) {
	return _Ethereum.Contract.CheckThreshold(&_Ethereum.CallOpts, _voteWeight)
}

// DailyWithdrawalLimit is a free data retrieval call binding the contract method 0xab796566.
//
// Solidity: function dailyWithdrawalLimit(address ) view returns(uint256)
func (_Ethereum *EthereumCaller) DailyWithdrawalLimit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "dailyWithdrawalLimit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DailyWithdrawalLimit is a free data retrieval call binding the contract method 0xab796566.
//
// Solidity: function dailyWithdrawalLimit(address ) view returns(uint256)
func (_Ethereum *EthereumSession) DailyWithdrawalLimit(arg0 common.Address) (*big.Int, error) {
	return _Ethereum.Contract.DailyWithdrawalLimit(&_Ethereum.CallOpts, arg0)
}

// DailyWithdrawalLimit is a free data retrieval call binding the contract method 0xab796566.
//
// Solidity: function dailyWithdrawalLimit(address ) view returns(uint256)
func (_Ethereum *EthereumCallerSession) DailyWithdrawalLimit(arg0 common.Address) (*big.Int, error) {
	return _Ethereum.Contract.DailyWithdrawalLimit(&_Ethereum.CallOpts, arg0)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Ethereum *EthereumCaller) DepositCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "depositCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Ethereum *EthereumSession) DepositCount() (*big.Int, error) {
	return _Ethereum.Contract.DepositCount(&_Ethereum.CallOpts)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Ethereum *EthereumCallerSession) DepositCount() (*big.Int, error) {
	return _Ethereum.Contract.DepositCount(&_Ethereum.CallOpts)
}

// FullSigThreshold is a free data retrieval call binding the contract method 0xdf5baf45.
//
// Solidity: function fullSigThreshold(address ) view returns(uint256)
func (_Ethereum *EthereumCaller) FullSigThreshold(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "fullSigThreshold", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FullSigThreshold is a free data retrieval call binding the contract method 0xdf5baf45.
//
// Solidity: function fullSigThreshold(address ) view returns(uint256)
func (_Ethereum *EthereumSession) FullSigThreshold(arg0 common.Address) (*big.Int, error) {
	return _Ethereum.Contract.FullSigThreshold(&_Ethereum.CallOpts, arg0)
}

// FullSigThreshold is a free data retrieval call binding the contract method 0xdf5baf45.
//
// Solidity: function fullSigThreshold(address ) view returns(uint256)
func (_Ethereum *EthereumCallerSession) FullSigThreshold(arg0 common.Address) (*big.Int, error) {
	return _Ethereum.Contract.FullSigThreshold(&_Ethereum.CallOpts, arg0)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Ethereum *EthereumCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Ethereum *EthereumSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Ethereum.Contract.GetRoleAdmin(&_Ethereum.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Ethereum *EthereumCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Ethereum.Contract.GetRoleAdmin(&_Ethereum.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Ethereum *EthereumCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Ethereum *EthereumSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Ethereum.Contract.GetRoleMember(&_Ethereum.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Ethereum *EthereumCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Ethereum.Contract.GetRoleMember(&_Ethereum.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Ethereum *EthereumCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Ethereum *EthereumSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Ethereum.Contract.GetRoleMemberCount(&_Ethereum.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Ethereum *EthereumCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Ethereum.Contract.GetRoleMemberCount(&_Ethereum.CallOpts, role)
}

// GetRoninToken is a free data retrieval call binding the contract method 0xb2975794.
//
// Solidity: function getRoninToken(address _mainchainToken) view returns(address _addr)
func (_Ethereum *EthereumCaller) GetRoninToken(opts *bind.CallOpts, _mainchainToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "getRoninToken", _mainchainToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoninToken is a free data retrieval call binding the contract method 0xb2975794.
//
// Solidity: function getRoninToken(address _mainchainToken) view returns(address _addr)
func (_Ethereum *EthereumSession) GetRoninToken(_mainchainToken common.Address) (common.Address, error) {
	return _Ethereum.Contract.GetRoninToken(&_Ethereum.CallOpts, _mainchainToken)
}

// GetRoninToken is a free data retrieval call binding the contract method 0xb2975794.
//
// Solidity: function getRoninToken(address _mainchainToken) view returns(address _addr)
func (_Ethereum *EthereumCallerSession) GetRoninToken(_mainchainToken common.Address) (common.Address, error) {
	return _Ethereum.Contract.GetRoninToken(&_Ethereum.CallOpts, _mainchainToken)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256, uint256)
func (_Ethereum *EthereumCaller) GetThreshold(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "getThreshold")

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
func (_Ethereum *EthereumSession) GetThreshold() (*big.Int, *big.Int, error) {
	return _Ethereum.Contract.GetThreshold(&_Ethereum.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256, uint256)
func (_Ethereum *EthereumCallerSession) GetThreshold() (*big.Int, *big.Int, error) {
	return _Ethereum.Contract.GetThreshold(&_Ethereum.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Ethereum *EthereumCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Ethereum *EthereumSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Ethereum.Contract.HasRole(&_Ethereum.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Ethereum *EthereumCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Ethereum.Contract.HasRole(&_Ethereum.CallOpts, role, account)
}

// LastDateSynced is a free data retrieval call binding the contract method 0xcbc829c6.
//
// Solidity: function lastDateSynced() view returns(uint256)
func (_Ethereum *EthereumCaller) LastDateSynced(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "lastDateSynced")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastDateSynced is a free data retrieval call binding the contract method 0xcbc829c6.
//
// Solidity: function lastDateSynced() view returns(uint256)
func (_Ethereum *EthereumSession) LastDateSynced() (*big.Int, error) {
	return _Ethereum.Contract.LastDateSynced(&_Ethereum.CallOpts)
}

// LastDateSynced is a free data retrieval call binding the contract method 0xcbc829c6.
//
// Solidity: function lastDateSynced() view returns(uint256)
func (_Ethereum *EthereumCallerSession) LastDateSynced() (*big.Int, error) {
	return _Ethereum.Contract.LastDateSynced(&_Ethereum.CallOpts)
}

// LastSyncedWithdrawal is a free data retrieval call binding the contract method 0xd55ed103.
//
// Solidity: function lastSyncedWithdrawal(address ) view returns(uint256)
func (_Ethereum *EthereumCaller) LastSyncedWithdrawal(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "lastSyncedWithdrawal", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastSyncedWithdrawal is a free data retrieval call binding the contract method 0xd55ed103.
//
// Solidity: function lastSyncedWithdrawal(address ) view returns(uint256)
func (_Ethereum *EthereumSession) LastSyncedWithdrawal(arg0 common.Address) (*big.Int, error) {
	return _Ethereum.Contract.LastSyncedWithdrawal(&_Ethereum.CallOpts, arg0)
}

// LastSyncedWithdrawal is a free data retrieval call binding the contract method 0xd55ed103.
//
// Solidity: function lastSyncedWithdrawal(address ) view returns(uint256)
func (_Ethereum *EthereumCallerSession) LastSyncedWithdrawal(arg0 common.Address) (*big.Int, error) {
	return _Ethereum.Contract.LastSyncedWithdrawal(&_Ethereum.CallOpts, arg0)
}

// LockedThreshold is a free data retrieval call binding the contract method 0x59122f6b.
//
// Solidity: function lockedThreshold(address ) view returns(uint256)
func (_Ethereum *EthereumCaller) LockedThreshold(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "lockedThreshold", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedThreshold is a free data retrieval call binding the contract method 0x59122f6b.
//
// Solidity: function lockedThreshold(address ) view returns(uint256)
func (_Ethereum *EthereumSession) LockedThreshold(arg0 common.Address) (*big.Int, error) {
	return _Ethereum.Contract.LockedThreshold(&_Ethereum.CallOpts, arg0)
}

// LockedThreshold is a free data retrieval call binding the contract method 0x59122f6b.
//
// Solidity: function lockedThreshold(address ) view returns(uint256)
func (_Ethereum *EthereumCallerSession) LockedThreshold(arg0 common.Address) (*big.Int, error) {
	return _Ethereum.Contract.LockedThreshold(&_Ethereum.CallOpts, arg0)
}

// MinimumVoteWeight is a free data retrieval call binding the contract method 0x7de5dedd.
//
// Solidity: function minimumVoteWeight() view returns(uint256)
func (_Ethereum *EthereumCaller) MinimumVoteWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "minimumVoteWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumVoteWeight is a free data retrieval call binding the contract method 0x7de5dedd.
//
// Solidity: function minimumVoteWeight() view returns(uint256)
func (_Ethereum *EthereumSession) MinimumVoteWeight() (*big.Int, error) {
	return _Ethereum.Contract.MinimumVoteWeight(&_Ethereum.CallOpts)
}

// MinimumVoteWeight is a free data retrieval call binding the contract method 0x7de5dedd.
//
// Solidity: function minimumVoteWeight() view returns(uint256)
func (_Ethereum *EthereumCallerSession) MinimumVoteWeight() (*big.Int, error) {
	return _Ethereum.Contract.MinimumVoteWeight(&_Ethereum.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Ethereum *EthereumCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Ethereum *EthereumSession) Nonce() (*big.Int, error) {
	return _Ethereum.Contract.Nonce(&_Ethereum.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Ethereum *EthereumCallerSession) Nonce() (*big.Int, error) {
	return _Ethereum.Contract.Nonce(&_Ethereum.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Ethereum *EthereumCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Ethereum *EthereumSession) Paused() (bool, error) {
	return _Ethereum.Contract.Paused(&_Ethereum.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Ethereum *EthereumCallerSession) Paused() (bool, error) {
	return _Ethereum.Contract.Paused(&_Ethereum.CallOpts)
}

// ReachedWithdrawalLimit is a free data retrieval call binding the contract method 0x6c1ce670.
//
// Solidity: function reachedWithdrawalLimit(address _token, uint256 _quantity) view returns(bool)
func (_Ethereum *EthereumCaller) ReachedWithdrawalLimit(opts *bind.CallOpts, _token common.Address, _quantity *big.Int) (bool, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "reachedWithdrawalLimit", _token, _quantity)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ReachedWithdrawalLimit is a free data retrieval call binding the contract method 0x6c1ce670.
//
// Solidity: function reachedWithdrawalLimit(address _token, uint256 _quantity) view returns(bool)
func (_Ethereum *EthereumSession) ReachedWithdrawalLimit(_token common.Address, _quantity *big.Int) (bool, error) {
	return _Ethereum.Contract.ReachedWithdrawalLimit(&_Ethereum.CallOpts, _token, _quantity)
}

// ReachedWithdrawalLimit is a free data retrieval call binding the contract method 0x6c1ce670.
//
// Solidity: function reachedWithdrawalLimit(address _token, uint256 _quantity) view returns(bool)
func (_Ethereum *EthereumCallerSession) ReachedWithdrawalLimit(_token common.Address, _quantity *big.Int) (bool, error) {
	return _Ethereum.Contract.ReachedWithdrawalLimit(&_Ethereum.CallOpts, _token, _quantity)
}

// RoninChainId is a free data retrieval call binding the contract method 0x17ce2dd4.
//
// Solidity: function roninChainId() view returns(uint256)
func (_Ethereum *EthereumCaller) RoninChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "roninChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoninChainId is a free data retrieval call binding the contract method 0x17ce2dd4.
//
// Solidity: function roninChainId() view returns(uint256)
func (_Ethereum *EthereumSession) RoninChainId() (*big.Int, error) {
	return _Ethereum.Contract.RoninChainId(&_Ethereum.CallOpts)
}

// RoninChainId is a free data retrieval call binding the contract method 0x17ce2dd4.
//
// Solidity: function roninChainId() view returns(uint256)
func (_Ethereum *EthereumCallerSession) RoninChainId() (*big.Int, error) {
	return _Ethereum.Contract.RoninChainId(&_Ethereum.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ethereum *EthereumCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ethereum *EthereumSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Ethereum.Contract.SupportsInterface(&_Ethereum.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ethereum *EthereumCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Ethereum.Contract.SupportsInterface(&_Ethereum.CallOpts, interfaceId)
}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_Ethereum *EthereumCaller) ValidatorContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "validatorContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_Ethereum *EthereumSession) ValidatorContract() (common.Address, error) {
	return _Ethereum.Contract.ValidatorContract(&_Ethereum.CallOpts)
}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_Ethereum *EthereumCallerSession) ValidatorContract() (common.Address, error) {
	return _Ethereum.Contract.ValidatorContract(&_Ethereum.CallOpts)
}

// WithdrawalHash is a free data retrieval call binding the contract method 0x6932be98.
//
// Solidity: function withdrawalHash(uint256 ) view returns(bytes32)
func (_Ethereum *EthereumCaller) WithdrawalHash(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "withdrawalHash", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WithdrawalHash is a free data retrieval call binding the contract method 0x6932be98.
//
// Solidity: function withdrawalHash(uint256 ) view returns(bytes32)
func (_Ethereum *EthereumSession) WithdrawalHash(arg0 *big.Int) ([32]byte, error) {
	return _Ethereum.Contract.WithdrawalHash(&_Ethereum.CallOpts, arg0)
}

// WithdrawalHash is a free data retrieval call binding the contract method 0x6932be98.
//
// Solidity: function withdrawalHash(uint256 ) view returns(bytes32)
func (_Ethereum *EthereumCallerSession) WithdrawalHash(arg0 *big.Int) ([32]byte, error) {
	return _Ethereum.Contract.WithdrawalHash(&_Ethereum.CallOpts, arg0)
}

// WithdrawalLocked is a free data retrieval call binding the contract method 0x4d493f4e.
//
// Solidity: function withdrawalLocked(uint256 ) view returns(bool)
func (_Ethereum *EthereumCaller) WithdrawalLocked(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "withdrawalLocked", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawalLocked is a free data retrieval call binding the contract method 0x4d493f4e.
//
// Solidity: function withdrawalLocked(uint256 ) view returns(bool)
func (_Ethereum *EthereumSession) WithdrawalLocked(arg0 *big.Int) (bool, error) {
	return _Ethereum.Contract.WithdrawalLocked(&_Ethereum.CallOpts, arg0)
}

// WithdrawalLocked is a free data retrieval call binding the contract method 0x4d493f4e.
//
// Solidity: function withdrawalLocked(uint256 ) view returns(bool)
func (_Ethereum *EthereumCallerSession) WithdrawalLocked(arg0 *big.Int) (bool, error) {
	return _Ethereum.Contract.WithdrawalLocked(&_Ethereum.CallOpts, arg0)
}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_Ethereum *EthereumCaller) WrappedNativeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "wrappedNativeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_Ethereum *EthereumSession) WrappedNativeToken() (common.Address, error) {
	return _Ethereum.Contract.WrappedNativeToken(&_Ethereum.CallOpts)
}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_Ethereum *EthereumCallerSession) WrappedNativeToken() (common.Address, error) {
	return _Ethereum.Contract.WrappedNativeToken(&_Ethereum.CallOpts)
}

// BulkRequestDepositFor is a paid mutator transaction binding the contract method 0xd2d9114f.
//
// Solidity: function bulkRequestDepositFor((address,address,(uint8,uint256,uint256))[] _requests) payable returns()
func (_Ethereum *EthereumTransactor) BulkRequestDepositFor(opts *bind.TransactOpts, _requests []TransferRequest) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "bulkRequestDepositFor", _requests)
}

// BulkRequestDepositFor is a paid mutator transaction binding the contract method 0xd2d9114f.
//
// Solidity: function bulkRequestDepositFor((address,address,(uint8,uint256,uint256))[] _requests) payable returns()
func (_Ethereum *EthereumSession) BulkRequestDepositFor(_requests []TransferRequest) (*types.Transaction, error) {
	return _Ethereum.Contract.BulkRequestDepositFor(&_Ethereum.TransactOpts, _requests)
}

// BulkRequestDepositFor is a paid mutator transaction binding the contract method 0xd2d9114f.
//
// Solidity: function bulkRequestDepositFor((address,address,(uint8,uint256,uint256))[] _requests) payable returns()
func (_Ethereum *EthereumTransactorSession) BulkRequestDepositFor(_requests []TransferRequest) (*types.Transaction, error) {
	return _Ethereum.Contract.BulkRequestDepositFor(&_Ethereum.TransactOpts, _requests)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Ethereum *EthereumTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Ethereum *EthereumSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ethereum.Contract.GrantRole(&_Ethereum.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Ethereum *EthereumTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ethereum.Contract.GrantRole(&_Ethereum.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc48d999a.
//
// Solidity: function initialize(address _roleSetter, address _wrappedToken, address _validatorContract, uint256 _roninChainId, uint256 _numerator, uint256 _denominator, address[][3] _addresses, uint256[][3] _thresholds) payable returns()
func (_Ethereum *EthereumTransactor) Initialize(opts *bind.TransactOpts, _roleSetter common.Address, _wrappedToken common.Address, _validatorContract common.Address, _roninChainId *big.Int, _numerator *big.Int, _denominator *big.Int, _addresses [3][]common.Address, _thresholds [3][]*big.Int) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "initialize", _roleSetter, _wrappedToken, _validatorContract, _roninChainId, _numerator, _denominator, _addresses, _thresholds)
}

// Initialize is a paid mutator transaction binding the contract method 0xc48d999a.
//
// Solidity: function initialize(address _roleSetter, address _wrappedToken, address _validatorContract, uint256 _roninChainId, uint256 _numerator, uint256 _denominator, address[][3] _addresses, uint256[][3] _thresholds) payable returns()
func (_Ethereum *EthereumSession) Initialize(_roleSetter common.Address, _wrappedToken common.Address, _validatorContract common.Address, _roninChainId *big.Int, _numerator *big.Int, _denominator *big.Int, _addresses [3][]common.Address, _thresholds [3][]*big.Int) (*types.Transaction, error) {
	return _Ethereum.Contract.Initialize(&_Ethereum.TransactOpts, _roleSetter, _wrappedToken, _validatorContract, _roninChainId, _numerator, _denominator, _addresses, _thresholds)
}

// Initialize is a paid mutator transaction binding the contract method 0xc48d999a.
//
// Solidity: function initialize(address _roleSetter, address _wrappedToken, address _validatorContract, uint256 _roninChainId, uint256 _numerator, uint256 _denominator, address[][3] _addresses, uint256[][3] _thresholds) payable returns()
func (_Ethereum *EthereumTransactorSession) Initialize(_roleSetter common.Address, _wrappedToken common.Address, _validatorContract common.Address, _roninChainId *big.Int, _numerator *big.Int, _denominator *big.Int, _addresses [3][]common.Address, _thresholds [3][]*big.Int) (*types.Transaction, error) {
	return _Ethereum.Contract.Initialize(&_Ethereum.TransactOpts, _roleSetter, _wrappedToken, _validatorContract, _roninChainId, _numerator, _denominator, _addresses, _thresholds)
}

// MapTokens is a paid mutator transaction binding the contract method 0xfc73b5d0.
//
// Solidity: function mapTokens(address[] _mainchainTokens, address[] _roninTokens) returns()
func (_Ethereum *EthereumTransactor) MapTokens(opts *bind.TransactOpts, _mainchainTokens []common.Address, _roninTokens []common.Address) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "mapTokens", _mainchainTokens, _roninTokens)
}

// MapTokens is a paid mutator transaction binding the contract method 0xfc73b5d0.
//
// Solidity: function mapTokens(address[] _mainchainTokens, address[] _roninTokens) returns()
func (_Ethereum *EthereumSession) MapTokens(_mainchainTokens []common.Address, _roninTokens []common.Address) (*types.Transaction, error) {
	return _Ethereum.Contract.MapTokens(&_Ethereum.TransactOpts, _mainchainTokens, _roninTokens)
}

// MapTokens is a paid mutator transaction binding the contract method 0xfc73b5d0.
//
// Solidity: function mapTokens(address[] _mainchainTokens, address[] _roninTokens) returns()
func (_Ethereum *EthereumTransactorSession) MapTokens(_mainchainTokens []common.Address, _roninTokens []common.Address) (*types.Transaction, error) {
	return _Ethereum.Contract.MapTokens(&_Ethereum.TransactOpts, _mainchainTokens, _roninTokens)
}

// MapTokensAndThresholds is a paid mutator transaction binding the contract method 0x056f7781.
//
// Solidity: function mapTokensAndThresholds(address[] _mainchainTokens, address[] _roninTokens, uint256[] _fullSigThreshold, uint256[] _lockedThreshold, uint256[] _dailyWithdrawalLimit) returns()
func (_Ethereum *EthereumTransactor) MapTokensAndThresholds(opts *bind.TransactOpts, _mainchainTokens []common.Address, _roninTokens []common.Address, _fullSigThreshold []*big.Int, _lockedThreshold []*big.Int, _dailyWithdrawalLimit []*big.Int) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "mapTokensAndThresholds", _mainchainTokens, _roninTokens, _fullSigThreshold, _lockedThreshold, _dailyWithdrawalLimit)
}

// MapTokensAndThresholds is a paid mutator transaction binding the contract method 0x056f7781.
//
// Solidity: function mapTokensAndThresholds(address[] _mainchainTokens, address[] _roninTokens, uint256[] _fullSigThreshold, uint256[] _lockedThreshold, uint256[] _dailyWithdrawalLimit) returns()
func (_Ethereum *EthereumSession) MapTokensAndThresholds(_mainchainTokens []common.Address, _roninTokens []common.Address, _fullSigThreshold []*big.Int, _lockedThreshold []*big.Int, _dailyWithdrawalLimit []*big.Int) (*types.Transaction, error) {
	return _Ethereum.Contract.MapTokensAndThresholds(&_Ethereum.TransactOpts, _mainchainTokens, _roninTokens, _fullSigThreshold, _lockedThreshold, _dailyWithdrawalLimit)
}

// MapTokensAndThresholds is a paid mutator transaction binding the contract method 0x056f7781.
//
// Solidity: function mapTokensAndThresholds(address[] _mainchainTokens, address[] _roninTokens, uint256[] _fullSigThreshold, uint256[] _lockedThreshold, uint256[] _dailyWithdrawalLimit) returns()
func (_Ethereum *EthereumTransactorSession) MapTokensAndThresholds(_mainchainTokens []common.Address, _roninTokens []common.Address, _fullSigThreshold []*big.Int, _lockedThreshold []*big.Int, _dailyWithdrawalLimit []*big.Int) (*types.Transaction, error) {
	return _Ethereum.Contract.MapTokensAndThresholds(&_Ethereum.TransactOpts, _mainchainTokens, _roninTokens, _fullSigThreshold, _lockedThreshold, _dailyWithdrawalLimit)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Ethereum *EthereumTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Ethereum *EthereumSession) Pause() (*types.Transaction, error) {
	return _Ethereum.Contract.Pause(&_Ethereum.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Ethereum *EthereumTransactorSession) Pause() (*types.Transaction, error) {
	return _Ethereum.Contract.Pause(&_Ethereum.TransactOpts)
}

// ReceiveEther is a paid mutator transaction binding the contract method 0xa3912ec8.
//
// Solidity: function receiveEther() payable returns()
func (_Ethereum *EthereumTransactor) ReceiveEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "receiveEther")
}

// ReceiveEther is a paid mutator transaction binding the contract method 0xa3912ec8.
//
// Solidity: function receiveEther() payable returns()
func (_Ethereum *EthereumSession) ReceiveEther() (*types.Transaction, error) {
	return _Ethereum.Contract.ReceiveEther(&_Ethereum.TransactOpts)
}

// ReceiveEther is a paid mutator transaction binding the contract method 0xa3912ec8.
//
// Solidity: function receiveEther() payable returns()
func (_Ethereum *EthereumTransactorSession) ReceiveEther() (*types.Transaction, error) {
	return _Ethereum.Contract.ReceiveEther(&_Ethereum.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Ethereum *EthereumTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Ethereum *EthereumSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ethereum.Contract.RenounceRole(&_Ethereum.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Ethereum *EthereumTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ethereum.Contract.RenounceRole(&_Ethereum.TransactOpts, role, account)
}

// RequestDepositFor is a paid mutator transaction binding the contract method 0x4b14557e.
//
// Solidity: function requestDepositFor((address,address,(uint8,uint256,uint256)) _request) payable returns()
func (_Ethereum *EthereumTransactor) RequestDepositFor(opts *bind.TransactOpts, _request TransferRequest) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "requestDepositFor", _request)
}

// RequestDepositFor is a paid mutator transaction binding the contract method 0x4b14557e.
//
// Solidity: function requestDepositFor((address,address,(uint8,uint256,uint256)) _request) payable returns()
func (_Ethereum *EthereumSession) RequestDepositFor(_request TransferRequest) (*types.Transaction, error) {
	return _Ethereum.Contract.RequestDepositFor(&_Ethereum.TransactOpts, _request)
}

// RequestDepositFor is a paid mutator transaction binding the contract method 0x4b14557e.
//
// Solidity: function requestDepositFor((address,address,(uint8,uint256,uint256)) _request) payable returns()
func (_Ethereum *EthereumTransactorSession) RequestDepositFor(_request TransferRequest) (*types.Transaction, error) {
	return _Ethereum.Contract.RequestDepositFor(&_Ethereum.TransactOpts, _request)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Ethereum *EthereumTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Ethereum *EthereumSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ethereum.Contract.RevokeRole(&_Ethereum.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Ethereum *EthereumTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ethereum.Contract.RevokeRole(&_Ethereum.TransactOpts, role, account)
}

// SetDailyWithdrawalLimits is a paid mutator transaction binding the contract method 0xe400327c.
//
// Solidity: function setDailyWithdrawalLimits(address[] _tokens, uint256[] _limits) returns()
func (_Ethereum *EthereumTransactor) SetDailyWithdrawalLimits(opts *bind.TransactOpts, _tokens []common.Address, _limits []*big.Int) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "setDailyWithdrawalLimits", _tokens, _limits)
}

// SetDailyWithdrawalLimits is a paid mutator transaction binding the contract method 0xe400327c.
//
// Solidity: function setDailyWithdrawalLimits(address[] _tokens, uint256[] _limits) returns()
func (_Ethereum *EthereumSession) SetDailyWithdrawalLimits(_tokens []common.Address, _limits []*big.Int) (*types.Transaction, error) {
	return _Ethereum.Contract.SetDailyWithdrawalLimits(&_Ethereum.TransactOpts, _tokens, _limits)
}

// SetDailyWithdrawalLimits is a paid mutator transaction binding the contract method 0xe400327c.
//
// Solidity: function setDailyWithdrawalLimits(address[] _tokens, uint256[] _limits) returns()
func (_Ethereum *EthereumTransactorSession) SetDailyWithdrawalLimits(_tokens []common.Address, _limits []*big.Int) (*types.Transaction, error) {
	return _Ethereum.Contract.SetDailyWithdrawalLimits(&_Ethereum.TransactOpts, _tokens, _limits)
}

// SetFullSigsThresholds is a paid mutator transaction binding the contract method 0xd8356de7.
//
// Solidity: function setFullSigsThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_Ethereum *EthereumTransactor) SetFullSigsThresholds(opts *bind.TransactOpts, _tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "setFullSigsThresholds", _tokens, _thresholds)
}

// SetFullSigsThresholds is a paid mutator transaction binding the contract method 0xd8356de7.
//
// Solidity: function setFullSigsThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_Ethereum *EthereumSession) SetFullSigsThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _Ethereum.Contract.SetFullSigsThresholds(&_Ethereum.TransactOpts, _tokens, _thresholds)
}

// SetFullSigsThresholds is a paid mutator transaction binding the contract method 0xd8356de7.
//
// Solidity: function setFullSigsThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_Ethereum *EthereumTransactorSession) SetFullSigsThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _Ethereum.Contract.SetFullSigsThresholds(&_Ethereum.TransactOpts, _tokens, _thresholds)
}

// SetLockedThresholds is a paid mutator transaction binding the contract method 0x1a8e55b0.
//
// Solidity: function setLockedThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_Ethereum *EthereumTransactor) SetLockedThresholds(opts *bind.TransactOpts, _tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "setLockedThresholds", _tokens, _thresholds)
}

// SetLockedThresholds is a paid mutator transaction binding the contract method 0x1a8e55b0.
//
// Solidity: function setLockedThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_Ethereum *EthereumSession) SetLockedThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _Ethereum.Contract.SetLockedThresholds(&_Ethereum.TransactOpts, _tokens, _thresholds)
}

// SetLockedThresholds is a paid mutator transaction binding the contract method 0x1a8e55b0.
//
// Solidity: function setLockedThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_Ethereum *EthereumTransactorSession) SetLockedThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _Ethereum.Contract.SetLockedThresholds(&_Ethereum.TransactOpts, _tokens, _thresholds)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 _numerator, uint256 _denominator) returns(uint256, uint256)
func (_Ethereum *EthereumTransactor) SetThreshold(opts *bind.TransactOpts, _numerator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "setThreshold", _numerator, _denominator)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 _numerator, uint256 _denominator) returns(uint256, uint256)
func (_Ethereum *EthereumSession) SetThreshold(_numerator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _Ethereum.Contract.SetThreshold(&_Ethereum.TransactOpts, _numerator, _denominator)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 _numerator, uint256 _denominator) returns(uint256, uint256)
func (_Ethereum *EthereumTransactorSession) SetThreshold(_numerator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _Ethereum.Contract.SetThreshold(&_Ethereum.TransactOpts, _numerator, _denominator)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _validatorContract) returns()
func (_Ethereum *EthereumTransactor) SetValidatorContract(opts *bind.TransactOpts, _validatorContract common.Address) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "setValidatorContract", _validatorContract)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _validatorContract) returns()
func (_Ethereum *EthereumSession) SetValidatorContract(_validatorContract common.Address) (*types.Transaction, error) {
	return _Ethereum.Contract.SetValidatorContract(&_Ethereum.TransactOpts, _validatorContract)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _validatorContract) returns()
func (_Ethereum *EthereumTransactorSession) SetValidatorContract(_validatorContract common.Address) (*types.Transaction, error) {
	return _Ethereum.Contract.SetValidatorContract(&_Ethereum.TransactOpts, _validatorContract)
}

// SetWrappedNativeTokenContract is a paid mutator transaction binding the contract method 0xd64af2a6.
//
// Solidity: function setWrappedNativeTokenContract(address _wrappedToken) returns()
func (_Ethereum *EthereumTransactor) SetWrappedNativeTokenContract(opts *bind.TransactOpts, _wrappedToken common.Address) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "setWrappedNativeTokenContract", _wrappedToken)
}

// SetWrappedNativeTokenContract is a paid mutator transaction binding the contract method 0xd64af2a6.
//
// Solidity: function setWrappedNativeTokenContract(address _wrappedToken) returns()
func (_Ethereum *EthereumSession) SetWrappedNativeTokenContract(_wrappedToken common.Address) (*types.Transaction, error) {
	return _Ethereum.Contract.SetWrappedNativeTokenContract(&_Ethereum.TransactOpts, _wrappedToken)
}

// SetWrappedNativeTokenContract is a paid mutator transaction binding the contract method 0xd64af2a6.
//
// Solidity: function setWrappedNativeTokenContract(address _wrappedToken) returns()
func (_Ethereum *EthereumTransactorSession) SetWrappedNativeTokenContract(_wrappedToken common.Address) (*types.Transaction, error) {
	return _Ethereum.Contract.SetWrappedNativeTokenContract(&_Ethereum.TransactOpts, _wrappedToken)
}

// SubmitWithdrawal is a paid mutator transaction binding the contract method 0x4d0d6673.
//
// Solidity: function submitWithdrawal((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt, (uint8,bytes32,bytes32)[] _signatures) returns(bool _locked)
func (_Ethereum *EthereumTransactor) SubmitWithdrawal(opts *bind.TransactOpts, _receipt TransferReceipt, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "submitWithdrawal", _receipt, _signatures)
}

// SubmitWithdrawal is a paid mutator transaction binding the contract method 0x4d0d6673.
//
// Solidity: function submitWithdrawal((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt, (uint8,bytes32,bytes32)[] _signatures) returns(bool _locked)
func (_Ethereum *EthereumSession) SubmitWithdrawal(_receipt TransferReceipt, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Ethereum.Contract.SubmitWithdrawal(&_Ethereum.TransactOpts, _receipt, _signatures)
}

// SubmitWithdrawal is a paid mutator transaction binding the contract method 0x4d0d6673.
//
// Solidity: function submitWithdrawal((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt, (uint8,bytes32,bytes32)[] _signatures) returns(bool _locked)
func (_Ethereum *EthereumTransactorSession) SubmitWithdrawal(_receipt TransferReceipt, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Ethereum.Contract.SubmitWithdrawal(&_Ethereum.TransactOpts, _receipt, _signatures)
}

// UnlockWithdrawal is a paid mutator transaction binding the contract method 0x9157921c.
//
// Solidity: function unlockWithdrawal((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt) returns()
func (_Ethereum *EthereumTransactor) UnlockWithdrawal(opts *bind.TransactOpts, _receipt TransferReceipt) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "unlockWithdrawal", _receipt)
}

// UnlockWithdrawal is a paid mutator transaction binding the contract method 0x9157921c.
//
// Solidity: function unlockWithdrawal((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt) returns()
func (_Ethereum *EthereumSession) UnlockWithdrawal(_receipt TransferReceipt) (*types.Transaction, error) {
	return _Ethereum.Contract.UnlockWithdrawal(&_Ethereum.TransactOpts, _receipt)
}

// UnlockWithdrawal is a paid mutator transaction binding the contract method 0x9157921c.
//
// Solidity: function unlockWithdrawal((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt) returns()
func (_Ethereum *EthereumTransactorSession) UnlockWithdrawal(_receipt TransferReceipt) (*types.Transaction, error) {
	return _Ethereum.Contract.UnlockWithdrawal(&_Ethereum.TransactOpts, _receipt)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Ethereum *EthereumTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Ethereum *EthereumSession) Unpause() (*types.Transaction, error) {
	return _Ethereum.Contract.Unpause(&_Ethereum.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Ethereum *EthereumTransactorSession) Unpause() (*types.Transaction, error) {
	return _Ethereum.Contract.Unpause(&_Ethereum.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Ethereum *EthereumTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Ethereum.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Ethereum *EthereumSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Ethereum.Contract.Fallback(&_Ethereum.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Ethereum *EthereumTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Ethereum.Contract.Fallback(&_Ethereum.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ethereum *EthereumTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethereum.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ethereum *EthereumSession) Receive() (*types.Transaction, error) {
	return _Ethereum.Contract.Receive(&_Ethereum.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ethereum *EthereumTransactorSession) Receive() (*types.Transaction, error) {
	return _Ethereum.Contract.Receive(&_Ethereum.TransactOpts)
}

// EthereumDailyWithdrawalLimitsUpdatedIterator is returned from FilterDailyWithdrawalLimitsUpdated and is used to iterate over the raw logs and unpacked data for DailyWithdrawalLimitsUpdated events raised by the Ethereum contract.
type EthereumDailyWithdrawalLimitsUpdatedIterator struct {
	Event *EthereumDailyWithdrawalLimitsUpdated // Event containing the contract specifics and raw log

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
func (it *EthereumDailyWithdrawalLimitsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumDailyWithdrawalLimitsUpdated)
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
		it.Event = new(EthereumDailyWithdrawalLimitsUpdated)
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
func (it *EthereumDailyWithdrawalLimitsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumDailyWithdrawalLimitsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumDailyWithdrawalLimitsUpdated represents a DailyWithdrawalLimitsUpdated event raised by the Ethereum contract.
type EthereumDailyWithdrawalLimitsUpdated struct {
	Tokens []common.Address
	Limits []*big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDailyWithdrawalLimitsUpdated is a free log retrieval operation binding the contract event 0xb5d2963614d72181b4df1f993d45b83edf42fa19710f0204217ba1b3e183bb73.
//
// Solidity: event DailyWithdrawalLimitsUpdated(address[] tokens, uint256[] limits)
func (_Ethereum *EthereumFilterer) FilterDailyWithdrawalLimitsUpdated(opts *bind.FilterOpts) (*EthereumDailyWithdrawalLimitsUpdatedIterator, error) {

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "DailyWithdrawalLimitsUpdated")
	if err != nil {
		return nil, err
	}
	return &EthereumDailyWithdrawalLimitsUpdatedIterator{contract: _Ethereum.contract, event: "DailyWithdrawalLimitsUpdated", logs: logs, sub: sub}, nil
}

// WatchDailyWithdrawalLimitsUpdated is a free log subscription operation binding the contract event 0xb5d2963614d72181b4df1f993d45b83edf42fa19710f0204217ba1b3e183bb73.
//
// Solidity: event DailyWithdrawalLimitsUpdated(address[] tokens, uint256[] limits)
func (_Ethereum *EthereumFilterer) WatchDailyWithdrawalLimitsUpdated(opts *bind.WatchOpts, sink chan<- *EthereumDailyWithdrawalLimitsUpdated) (event.Subscription, error) {

	logs, sub, err := _Ethereum.contract.WatchLogs(opts, "DailyWithdrawalLimitsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumDailyWithdrawalLimitsUpdated)
				if err := _Ethereum.contract.UnpackLog(event, "DailyWithdrawalLimitsUpdated", log); err != nil {
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

// ParseDailyWithdrawalLimitsUpdated is a log parse operation binding the contract event 0xb5d2963614d72181b4df1f993d45b83edf42fa19710f0204217ba1b3e183bb73.
//
// Solidity: event DailyWithdrawalLimitsUpdated(address[] tokens, uint256[] limits)
func (_Ethereum *EthereumFilterer) ParseDailyWithdrawalLimitsUpdated(log types.Log) (*EthereumDailyWithdrawalLimitsUpdated, error) {
	event := new(EthereumDailyWithdrawalLimitsUpdated)
	if err := _Ethereum.contract.UnpackLog(event, "DailyWithdrawalLimitsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumDepositRequestedIterator is returned from FilterDepositRequested and is used to iterate over the raw logs and unpacked data for DepositRequested events raised by the Ethereum contract.
type EthereumDepositRequestedIterator struct {
	Event *EthereumDepositRequested // Event containing the contract specifics and raw log

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
func (it *EthereumDepositRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumDepositRequested)
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
		it.Event = new(EthereumDepositRequested)
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
func (it *EthereumDepositRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumDepositRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumDepositRequested represents a DepositRequested event raised by the Ethereum contract.
type EthereumDepositRequested struct {
	ReceiptHash [32]byte
	Arg1        TransferReceipt
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDepositRequested is a free log retrieval operation binding the contract event 0xd7b25068d9dc8d00765254cfb7f5070f98d263c8d68931d937c7362fa738048b.
//
// Solidity: event DepositRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) arg1)
func (_Ethereum *EthereumFilterer) FilterDepositRequested(opts *bind.FilterOpts) (*EthereumDepositRequestedIterator, error) {

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "DepositRequested")
	if err != nil {
		return nil, err
	}
	return &EthereumDepositRequestedIterator{contract: _Ethereum.contract, event: "DepositRequested", logs: logs, sub: sub}, nil
}

// WatchDepositRequested is a free log subscription operation binding the contract event 0xd7b25068d9dc8d00765254cfb7f5070f98d263c8d68931d937c7362fa738048b.
//
// Solidity: event DepositRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) arg1)
func (_Ethereum *EthereumFilterer) WatchDepositRequested(opts *bind.WatchOpts, sink chan<- *EthereumDepositRequested) (event.Subscription, error) {

	logs, sub, err := _Ethereum.contract.WatchLogs(opts, "DepositRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumDepositRequested)
				if err := _Ethereum.contract.UnpackLog(event, "DepositRequested", log); err != nil {
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

// ParseDepositRequested is a log parse operation binding the contract event 0xd7b25068d9dc8d00765254cfb7f5070f98d263c8d68931d937c7362fa738048b.
//
// Solidity: event DepositRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) arg1)
func (_Ethereum *EthereumFilterer) ParseDepositRequested(log types.Log) (*EthereumDepositRequested, error) {
	event := new(EthereumDepositRequested)
	if err := _Ethereum.contract.UnpackLog(event, "DepositRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumFullSigsThresholdsUpdatedIterator is returned from FilterFullSigsThresholdsUpdated and is used to iterate over the raw logs and unpacked data for FullSigsThresholdsUpdated events raised by the Ethereum contract.
type EthereumFullSigsThresholdsUpdatedIterator struct {
	Event *EthereumFullSigsThresholdsUpdated // Event containing the contract specifics and raw log

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
func (it *EthereumFullSigsThresholdsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumFullSigsThresholdsUpdated)
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
		it.Event = new(EthereumFullSigsThresholdsUpdated)
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
func (it *EthereumFullSigsThresholdsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumFullSigsThresholdsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumFullSigsThresholdsUpdated represents a FullSigsThresholdsUpdated event raised by the Ethereum contract.
type EthereumFullSigsThresholdsUpdated struct {
	Tokens     []common.Address
	Thresholds []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFullSigsThresholdsUpdated is a free log retrieval operation binding the contract event 0x30f30fe53f33a6b009d6d0446c37f11eff8aa1033a9a92df9e8fda478cb768f7.
//
// Solidity: event FullSigsThresholdsUpdated(address[] tokens, uint256[] thresholds)
func (_Ethereum *EthereumFilterer) FilterFullSigsThresholdsUpdated(opts *bind.FilterOpts) (*EthereumFullSigsThresholdsUpdatedIterator, error) {

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "FullSigsThresholdsUpdated")
	if err != nil {
		return nil, err
	}
	return &EthereumFullSigsThresholdsUpdatedIterator{contract: _Ethereum.contract, event: "FullSigsThresholdsUpdated", logs: logs, sub: sub}, nil
}

// WatchFullSigsThresholdsUpdated is a free log subscription operation binding the contract event 0x30f30fe53f33a6b009d6d0446c37f11eff8aa1033a9a92df9e8fda478cb768f7.
//
// Solidity: event FullSigsThresholdsUpdated(address[] tokens, uint256[] thresholds)
func (_Ethereum *EthereumFilterer) WatchFullSigsThresholdsUpdated(opts *bind.WatchOpts, sink chan<- *EthereumFullSigsThresholdsUpdated) (event.Subscription, error) {

	logs, sub, err := _Ethereum.contract.WatchLogs(opts, "FullSigsThresholdsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumFullSigsThresholdsUpdated)
				if err := _Ethereum.contract.UnpackLog(event, "FullSigsThresholdsUpdated", log); err != nil {
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

// ParseFullSigsThresholdsUpdated is a log parse operation binding the contract event 0x30f30fe53f33a6b009d6d0446c37f11eff8aa1033a9a92df9e8fda478cb768f7.
//
// Solidity: event FullSigsThresholdsUpdated(address[] tokens, uint256[] thresholds)
func (_Ethereum *EthereumFilterer) ParseFullSigsThresholdsUpdated(log types.Log) (*EthereumFullSigsThresholdsUpdated, error) {
	event := new(EthereumFullSigsThresholdsUpdated)
	if err := _Ethereum.contract.UnpackLog(event, "FullSigsThresholdsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumLockedThresholdsUpdatedIterator is returned from FilterLockedThresholdsUpdated and is used to iterate over the raw logs and unpacked data for LockedThresholdsUpdated events raised by the Ethereum contract.
type EthereumLockedThresholdsUpdatedIterator struct {
	Event *EthereumLockedThresholdsUpdated // Event containing the contract specifics and raw log

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
func (it *EthereumLockedThresholdsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumLockedThresholdsUpdated)
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
		it.Event = new(EthereumLockedThresholdsUpdated)
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
func (it *EthereumLockedThresholdsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumLockedThresholdsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumLockedThresholdsUpdated represents a LockedThresholdsUpdated event raised by the Ethereum contract.
type EthereumLockedThresholdsUpdated struct {
	Tokens     []common.Address
	Thresholds []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLockedThresholdsUpdated is a free log retrieval operation binding the contract event 0x64557254143204d91ba2d95acb9fda1e5fea55f77efd028685765bc1e94dd4b5.
//
// Solidity: event LockedThresholdsUpdated(address[] tokens, uint256[] thresholds)
func (_Ethereum *EthereumFilterer) FilterLockedThresholdsUpdated(opts *bind.FilterOpts) (*EthereumLockedThresholdsUpdatedIterator, error) {

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "LockedThresholdsUpdated")
	if err != nil {
		return nil, err
	}
	return &EthereumLockedThresholdsUpdatedIterator{contract: _Ethereum.contract, event: "LockedThresholdsUpdated", logs: logs, sub: sub}, nil
}

// WatchLockedThresholdsUpdated is a free log subscription operation binding the contract event 0x64557254143204d91ba2d95acb9fda1e5fea55f77efd028685765bc1e94dd4b5.
//
// Solidity: event LockedThresholdsUpdated(address[] tokens, uint256[] thresholds)
func (_Ethereum *EthereumFilterer) WatchLockedThresholdsUpdated(opts *bind.WatchOpts, sink chan<- *EthereumLockedThresholdsUpdated) (event.Subscription, error) {

	logs, sub, err := _Ethereum.contract.WatchLogs(opts, "LockedThresholdsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumLockedThresholdsUpdated)
				if err := _Ethereum.contract.UnpackLog(event, "LockedThresholdsUpdated", log); err != nil {
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

// ParseLockedThresholdsUpdated is a log parse operation binding the contract event 0x64557254143204d91ba2d95acb9fda1e5fea55f77efd028685765bc1e94dd4b5.
//
// Solidity: event LockedThresholdsUpdated(address[] tokens, uint256[] thresholds)
func (_Ethereum *EthereumFilterer) ParseLockedThresholdsUpdated(log types.Log) (*EthereumLockedThresholdsUpdated, error) {
	event := new(EthereumLockedThresholdsUpdated)
	if err := _Ethereum.contract.UnpackLog(event, "LockedThresholdsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Ethereum contract.
type EthereumPausedIterator struct {
	Event *EthereumPaused // Event containing the contract specifics and raw log

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
func (it *EthereumPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumPaused)
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
		it.Event = new(EthereumPaused)
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
func (it *EthereumPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumPaused represents a Paused event raised by the Ethereum contract.
type EthereumPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Ethereum *EthereumFilterer) FilterPaused(opts *bind.FilterOpts) (*EthereumPausedIterator, error) {

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EthereumPausedIterator{contract: _Ethereum.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Ethereum *EthereumFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EthereumPaused) (event.Subscription, error) {

	logs, sub, err := _Ethereum.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumPaused)
				if err := _Ethereum.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Ethereum *EthereumFilterer) ParsePaused(log types.Log) (*EthereumPaused, error) {
	event := new(EthereumPaused)
	if err := _Ethereum.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Ethereum contract.
type EthereumRoleAdminChangedIterator struct {
	Event *EthereumRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *EthereumRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumRoleAdminChanged)
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
		it.Event = new(EthereumRoleAdminChanged)
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
func (it *EthereumRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumRoleAdminChanged represents a RoleAdminChanged event raised by the Ethereum contract.
type EthereumRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Ethereum *EthereumFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*EthereumRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &EthereumRoleAdminChangedIterator{contract: _Ethereum.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Ethereum *EthereumFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *EthereumRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Ethereum.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumRoleAdminChanged)
				if err := _Ethereum.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Ethereum *EthereumFilterer) ParseRoleAdminChanged(log types.Log) (*EthereumRoleAdminChanged, error) {
	event := new(EthereumRoleAdminChanged)
	if err := _Ethereum.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Ethereum contract.
type EthereumRoleGrantedIterator struct {
	Event *EthereumRoleGranted // Event containing the contract specifics and raw log

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
func (it *EthereumRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumRoleGranted)
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
		it.Event = new(EthereumRoleGranted)
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
func (it *EthereumRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumRoleGranted represents a RoleGranted event raised by the Ethereum contract.
type EthereumRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Ethereum *EthereumFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*EthereumRoleGrantedIterator, error) {

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

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EthereumRoleGrantedIterator{contract: _Ethereum.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Ethereum *EthereumFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *EthereumRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Ethereum.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumRoleGranted)
				if err := _Ethereum.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Ethereum *EthereumFilterer) ParseRoleGranted(log types.Log) (*EthereumRoleGranted, error) {
	event := new(EthereumRoleGranted)
	if err := _Ethereum.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Ethereum contract.
type EthereumRoleRevokedIterator struct {
	Event *EthereumRoleRevoked // Event containing the contract specifics and raw log

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
func (it *EthereumRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumRoleRevoked)
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
		it.Event = new(EthereumRoleRevoked)
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
func (it *EthereumRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumRoleRevoked represents a RoleRevoked event raised by the Ethereum contract.
type EthereumRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Ethereum *EthereumFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*EthereumRoleRevokedIterator, error) {

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

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EthereumRoleRevokedIterator{contract: _Ethereum.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Ethereum *EthereumFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *EthereumRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Ethereum.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumRoleRevoked)
				if err := _Ethereum.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Ethereum *EthereumFilterer) ParseRoleRevoked(log types.Log) (*EthereumRoleRevoked, error) {
	event := new(EthereumRoleRevoked)
	if err := _Ethereum.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumThresholdUpdatedIterator is returned from FilterThresholdUpdated and is used to iterate over the raw logs and unpacked data for ThresholdUpdated events raised by the Ethereum contract.
type EthereumThresholdUpdatedIterator struct {
	Event *EthereumThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *EthereumThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumThresholdUpdated)
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
		it.Event = new(EthereumThresholdUpdated)
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
func (it *EthereumThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumThresholdUpdated represents a ThresholdUpdated event raised by the Ethereum contract.
type EthereumThresholdUpdated struct {
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
func (_Ethereum *EthereumFilterer) FilterThresholdUpdated(opts *bind.FilterOpts, _nonce []*big.Int, _numerator []*big.Int, _denominator []*big.Int) (*EthereumThresholdUpdatedIterator, error) {

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

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "ThresholdUpdated", _nonceRule, _numeratorRule, _denominatorRule)
	if err != nil {
		return nil, err
	}
	return &EthereumThresholdUpdatedIterator{contract: _Ethereum.contract, event: "ThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchThresholdUpdated is a free log subscription operation binding the contract event 0x976f8a9c5bdf8248dec172376d6e2b80a8e3df2f0328e381c6db8e1cf138c0f8.
//
// Solidity: event ThresholdUpdated(uint256 indexed _nonce, uint256 indexed _numerator, uint256 indexed _denominator, uint256 _previousNumerator, uint256 _previousDenominator)
func (_Ethereum *EthereumFilterer) WatchThresholdUpdated(opts *bind.WatchOpts, sink chan<- *EthereumThresholdUpdated, _nonce []*big.Int, _numerator []*big.Int, _denominator []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Ethereum.contract.WatchLogs(opts, "ThresholdUpdated", _nonceRule, _numeratorRule, _denominatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumThresholdUpdated)
				if err := _Ethereum.contract.UnpackLog(event, "ThresholdUpdated", log); err != nil {
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
func (_Ethereum *EthereumFilterer) ParseThresholdUpdated(log types.Log) (*EthereumThresholdUpdated, error) {
	event := new(EthereumThresholdUpdated)
	if err := _Ethereum.contract.UnpackLog(event, "ThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumTokenMappedIterator is returned from FilterTokenMapped and is used to iterate over the raw logs and unpacked data for TokenMapped events raised by the Ethereum contract.
type EthereumTokenMappedIterator struct {
	Event *EthereumTokenMapped // Event containing the contract specifics and raw log

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
func (it *EthereumTokenMappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumTokenMapped)
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
		it.Event = new(EthereumTokenMapped)
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
func (it *EthereumTokenMappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumTokenMappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumTokenMapped represents a TokenMapped event raised by the Ethereum contract.
type EthereumTokenMapped struct {
	MainchainTokens []common.Address
	RoninTokens     []common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokenMapped is a free log retrieval operation binding the contract event 0x1b7fd57807f95645ee3e0d5a22377a7779e93655cd5f6871151ae7129df015b3.
//
// Solidity: event TokenMapped(address[] mainchainTokens, address[] roninTokens)
func (_Ethereum *EthereumFilterer) FilterTokenMapped(opts *bind.FilterOpts) (*EthereumTokenMappedIterator, error) {

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "TokenMapped")
	if err != nil {
		return nil, err
	}
	return &EthereumTokenMappedIterator{contract: _Ethereum.contract, event: "TokenMapped", logs: logs, sub: sub}, nil
}

// WatchTokenMapped is a free log subscription operation binding the contract event 0x1b7fd57807f95645ee3e0d5a22377a7779e93655cd5f6871151ae7129df015b3.
//
// Solidity: event TokenMapped(address[] mainchainTokens, address[] roninTokens)
func (_Ethereum *EthereumFilterer) WatchTokenMapped(opts *bind.WatchOpts, sink chan<- *EthereumTokenMapped) (event.Subscription, error) {

	logs, sub, err := _Ethereum.contract.WatchLogs(opts, "TokenMapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumTokenMapped)
				if err := _Ethereum.contract.UnpackLog(event, "TokenMapped", log); err != nil {
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

// ParseTokenMapped is a log parse operation binding the contract event 0x1b7fd57807f95645ee3e0d5a22377a7779e93655cd5f6871151ae7129df015b3.
//
// Solidity: event TokenMapped(address[] mainchainTokens, address[] roninTokens)
func (_Ethereum *EthereumFilterer) ParseTokenMapped(log types.Log) (*EthereumTokenMapped, error) {
	event := new(EthereumTokenMapped)
	if err := _Ethereum.contract.UnpackLog(event, "TokenMapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Ethereum contract.
type EthereumUnpausedIterator struct {
	Event *EthereumUnpaused // Event containing the contract specifics and raw log

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
func (it *EthereumUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumUnpaused)
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
		it.Event = new(EthereumUnpaused)
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
func (it *EthereumUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumUnpaused represents a Unpaused event raised by the Ethereum contract.
type EthereumUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Ethereum *EthereumFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EthereumUnpausedIterator, error) {

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EthereumUnpausedIterator{contract: _Ethereum.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Ethereum *EthereumFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EthereumUnpaused) (event.Subscription, error) {

	logs, sub, err := _Ethereum.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumUnpaused)
				if err := _Ethereum.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Ethereum *EthereumFilterer) ParseUnpaused(log types.Log) (*EthereumUnpaused, error) {
	event := new(EthereumUnpaused)
	if err := _Ethereum.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumValidatorContractUpdatedIterator is returned from FilterValidatorContractUpdated and is used to iterate over the raw logs and unpacked data for ValidatorContractUpdated events raised by the Ethereum contract.
type EthereumValidatorContractUpdatedIterator struct {
	Event *EthereumValidatorContractUpdated // Event containing the contract specifics and raw log

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
func (it *EthereumValidatorContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumValidatorContractUpdated)
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
		it.Event = new(EthereumValidatorContractUpdated)
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
func (it *EthereumValidatorContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumValidatorContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumValidatorContractUpdated represents a ValidatorContractUpdated event raised by the Ethereum contract.
type EthereumValidatorContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterValidatorContractUpdated is a free log retrieval operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_Ethereum *EthereumFilterer) FilterValidatorContractUpdated(opts *bind.FilterOpts) (*EthereumValidatorContractUpdatedIterator, error) {

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "ValidatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return &EthereumValidatorContractUpdatedIterator{contract: _Ethereum.contract, event: "ValidatorContractUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorContractUpdated is a free log subscription operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_Ethereum *EthereumFilterer) WatchValidatorContractUpdated(opts *bind.WatchOpts, sink chan<- *EthereumValidatorContractUpdated) (event.Subscription, error) {

	logs, sub, err := _Ethereum.contract.WatchLogs(opts, "ValidatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumValidatorContractUpdated)
				if err := _Ethereum.contract.UnpackLog(event, "ValidatorContractUpdated", log); err != nil {
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
func (_Ethereum *EthereumFilterer) ParseValidatorContractUpdated(log types.Log) (*EthereumValidatorContractUpdated, error) {
	event := new(EthereumValidatorContractUpdated)
	if err := _Ethereum.contract.UnpackLog(event, "ValidatorContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumWithdrawalLockedIterator is returned from FilterWithdrawalLocked and is used to iterate over the raw logs and unpacked data for WithdrawalLocked events raised by the Ethereum contract.
type EthereumWithdrawalLockedIterator struct {
	Event *EthereumWithdrawalLocked // Event containing the contract specifics and raw log

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
func (it *EthereumWithdrawalLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumWithdrawalLocked)
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
		it.Event = new(EthereumWithdrawalLocked)
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
func (it *EthereumWithdrawalLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumWithdrawalLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumWithdrawalLocked represents a WithdrawalLocked event raised by the Ethereum contract.
type EthereumWithdrawalLocked struct {
	Arg0 TransferReceipt
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalLocked is a free log retrieval operation binding the contract event 0xf770f1e32df4cff7799445eb849b642389647242bf1083d138a1f41bafce8ccb.
//
// Solidity: event WithdrawalLocked((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) arg0)
func (_Ethereum *EthereumFilterer) FilterWithdrawalLocked(opts *bind.FilterOpts) (*EthereumWithdrawalLockedIterator, error) {

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "WithdrawalLocked")
	if err != nil {
		return nil, err
	}
	return &EthereumWithdrawalLockedIterator{contract: _Ethereum.contract, event: "WithdrawalLocked", logs: logs, sub: sub}, nil
}

// WatchWithdrawalLocked is a free log subscription operation binding the contract event 0xf770f1e32df4cff7799445eb849b642389647242bf1083d138a1f41bafce8ccb.
//
// Solidity: event WithdrawalLocked((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) arg0)
func (_Ethereum *EthereumFilterer) WatchWithdrawalLocked(opts *bind.WatchOpts, sink chan<- *EthereumWithdrawalLocked) (event.Subscription, error) {

	logs, sub, err := _Ethereum.contract.WatchLogs(opts, "WithdrawalLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumWithdrawalLocked)
				if err := _Ethereum.contract.UnpackLog(event, "WithdrawalLocked", log); err != nil {
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

// ParseWithdrawalLocked is a log parse operation binding the contract event 0xf770f1e32df4cff7799445eb849b642389647242bf1083d138a1f41bafce8ccb.
//
// Solidity: event WithdrawalLocked((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) arg0)
func (_Ethereum *EthereumFilterer) ParseWithdrawalLocked(log types.Log) (*EthereumWithdrawalLocked, error) {
	event := new(EthereumWithdrawalLocked)
	if err := _Ethereum.contract.UnpackLog(event, "WithdrawalLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumWithdrawalUnlockedIterator is returned from FilterWithdrawalUnlocked and is used to iterate over the raw logs and unpacked data for WithdrawalUnlocked events raised by the Ethereum contract.
type EthereumWithdrawalUnlockedIterator struct {
	Event *EthereumWithdrawalUnlocked // Event containing the contract specifics and raw log

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
func (it *EthereumWithdrawalUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumWithdrawalUnlocked)
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
		it.Event = new(EthereumWithdrawalUnlocked)
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
func (it *EthereumWithdrawalUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumWithdrawalUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumWithdrawalUnlocked represents a WithdrawalUnlocked event raised by the Ethereum contract.
type EthereumWithdrawalUnlocked struct {
	Arg0 TransferReceipt
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalUnlocked is a free log retrieval operation binding the contract event 0x54aaae2a58b79359da16804ec55aeb9cc2d064a167cb23748efacfb6dbb7da32.
//
// Solidity: event WithdrawalUnlocked((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) arg0)
func (_Ethereum *EthereumFilterer) FilterWithdrawalUnlocked(opts *bind.FilterOpts) (*EthereumWithdrawalUnlockedIterator, error) {

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "WithdrawalUnlocked")
	if err != nil {
		return nil, err
	}
	return &EthereumWithdrawalUnlockedIterator{contract: _Ethereum.contract, event: "WithdrawalUnlocked", logs: logs, sub: sub}, nil
}

// WatchWithdrawalUnlocked is a free log subscription operation binding the contract event 0x54aaae2a58b79359da16804ec55aeb9cc2d064a167cb23748efacfb6dbb7da32.
//
// Solidity: event WithdrawalUnlocked((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) arg0)
func (_Ethereum *EthereumFilterer) WatchWithdrawalUnlocked(opts *bind.WatchOpts, sink chan<- *EthereumWithdrawalUnlocked) (event.Subscription, error) {

	logs, sub, err := _Ethereum.contract.WatchLogs(opts, "WithdrawalUnlocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumWithdrawalUnlocked)
				if err := _Ethereum.contract.UnpackLog(event, "WithdrawalUnlocked", log); err != nil {
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

// ParseWithdrawalUnlocked is a log parse operation binding the contract event 0x54aaae2a58b79359da16804ec55aeb9cc2d064a167cb23748efacfb6dbb7da32.
//
// Solidity: event WithdrawalUnlocked((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) arg0)
func (_Ethereum *EthereumFilterer) ParseWithdrawalUnlocked(log types.Log) (*EthereumWithdrawalUnlocked, error) {
	event := new(EthereumWithdrawalUnlocked)
	if err := _Ethereum.contract.UnpackLog(event, "WithdrawalUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumWithdrewIterator is returned from FilterWithdrew and is used to iterate over the raw logs and unpacked data for Withdrew events raised by the Ethereum contract.
type EthereumWithdrewIterator struct {
	Event *EthereumWithdrew // Event containing the contract specifics and raw log

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
func (it *EthereumWithdrewIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumWithdrew)
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
		it.Event = new(EthereumWithdrew)
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
func (it *EthereumWithdrewIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumWithdrewIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumWithdrew represents a Withdrew event raised by the Ethereum contract.
type EthereumWithdrew struct {
	ReceiptHash [32]byte
	Receipt     TransferReceipt
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrew is a free log retrieval operation binding the contract event 0x21e88e956aa3e086f6388e899965cef814688f99ad8bb29b08d396571016372d.
//
// Solidity: event Withdrew(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_Ethereum *EthereumFilterer) FilterWithdrew(opts *bind.FilterOpts) (*EthereumWithdrewIterator, error) {

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "Withdrew")
	if err != nil {
		return nil, err
	}
	return &EthereumWithdrewIterator{contract: _Ethereum.contract, event: "Withdrew", logs: logs, sub: sub}, nil
}

// WatchWithdrew is a free log subscription operation binding the contract event 0x21e88e956aa3e086f6388e899965cef814688f99ad8bb29b08d396571016372d.
//
// Solidity: event Withdrew(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_Ethereum *EthereumFilterer) WatchWithdrew(opts *bind.WatchOpts, sink chan<- *EthereumWithdrew) (event.Subscription, error) {

	logs, sub, err := _Ethereum.contract.WatchLogs(opts, "Withdrew")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumWithdrew)
				if err := _Ethereum.contract.UnpackLog(event, "Withdrew", log); err != nil {
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

// ParseWithdrew is a log parse operation binding the contract event 0x21e88e956aa3e086f6388e899965cef814688f99ad8bb29b08d396571016372d.
//
// Solidity: event Withdrew(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_Ethereum *EthereumFilterer) ParseWithdrew(log types.Log) (*EthereumWithdrew, error) {
	event := new(EthereumWithdrew)
	if err := _Ethereum.contract.UnpackLog(event, "Withdrew", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthereumWrappedNativeTokenContractUpdatedIterator is returned from FilterWrappedNativeTokenContractUpdated and is used to iterate over the raw logs and unpacked data for WrappedNativeTokenContractUpdated events raised by the Ethereum contract.
type EthereumWrappedNativeTokenContractUpdatedIterator struct {
	Event *EthereumWrappedNativeTokenContractUpdated // Event containing the contract specifics and raw log

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
func (it *EthereumWrappedNativeTokenContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthereumWrappedNativeTokenContractUpdated)
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
		it.Event = new(EthereumWrappedNativeTokenContractUpdated)
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
func (it *EthereumWrappedNativeTokenContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthereumWrappedNativeTokenContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthereumWrappedNativeTokenContractUpdated represents a WrappedNativeTokenContractUpdated event raised by the Ethereum contract.
type EthereumWrappedNativeTokenContractUpdated struct {
	Weth common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWrappedNativeTokenContractUpdated is a free log retrieval operation binding the contract event 0x9d2334c23be647e994f27a72c5eee42a43d5bdcfe15bb88e939103c2b114cbaf.
//
// Solidity: event WrappedNativeTokenContractUpdated(address weth)
func (_Ethereum *EthereumFilterer) FilterWrappedNativeTokenContractUpdated(opts *bind.FilterOpts) (*EthereumWrappedNativeTokenContractUpdatedIterator, error) {

	logs, sub, err := _Ethereum.contract.FilterLogs(opts, "WrappedNativeTokenContractUpdated")
	if err != nil {
		return nil, err
	}
	return &EthereumWrappedNativeTokenContractUpdatedIterator{contract: _Ethereum.contract, event: "WrappedNativeTokenContractUpdated", logs: logs, sub: sub}, nil
}

// WatchWrappedNativeTokenContractUpdated is a free log subscription operation binding the contract event 0x9d2334c23be647e994f27a72c5eee42a43d5bdcfe15bb88e939103c2b114cbaf.
//
// Solidity: event WrappedNativeTokenContractUpdated(address weth)
func (_Ethereum *EthereumFilterer) WatchWrappedNativeTokenContractUpdated(opts *bind.WatchOpts, sink chan<- *EthereumWrappedNativeTokenContractUpdated) (event.Subscription, error) {

	logs, sub, err := _Ethereum.contract.WatchLogs(opts, "WrappedNativeTokenContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthereumWrappedNativeTokenContractUpdated)
				if err := _Ethereum.contract.UnpackLog(event, "WrappedNativeTokenContractUpdated", log); err != nil {
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

// ParseWrappedNativeTokenContractUpdated is a log parse operation binding the contract event 0x9d2334c23be647e994f27a72c5eee42a43d5bdcfe15bb88e939103c2b114cbaf.
//
// Solidity: event WrappedNativeTokenContractUpdated(address weth)
func (_Ethereum *EthereumFilterer) ParseWrappedNativeTokenContractUpdated(log types.Log) (*EthereumWrappedNativeTokenContractUpdated, error) {
	event := new(EthereumWrappedNativeTokenContractUpdated)
	if err := _Ethereum.contract.UnpackLog(event, "WrappedNativeTokenContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
