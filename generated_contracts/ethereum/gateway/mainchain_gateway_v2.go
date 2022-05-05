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
	TokenAddr     common.Address
	Info          TokenInfo
}

// GatewayMetaData contains all meta data concerning the Gateway contract.
var GatewayMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"limits\",\"type\":\"uint256[]\"}],\"name\":\"DailyWithdrawalLimitsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"receiptHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"DepositRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"thresholds\",\"type\":\"uint256[]\"}],\"name\":\"FullSigsThresholdsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"thresholds\",\"type\":\"uint256[]\"}],\"name\":\"LockedThresholdsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"numerator\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousNumerator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousDenominator\",\"type\":\"uint256\"}],\"name\":\"ThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"mainchainTokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"roninTokens\",\"type\":\"address[]\"}],\"name\":\"TokenMapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"percentages\",\"type\":\"uint256[]\"}],\"name\":\"UnlockFeePercentagesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIWeightedValidator\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ValidatorContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"receiptHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"WithdrawalLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"receiptHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"WithdrawalUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"receiptHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structTransfer.Receipt\",\"name\":\"receipt\",\"type\":\"tuple\"}],\"name\":\"Withdrew\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIWETH\",\"name\":\"weth\",\"type\":\"address\"}],\"name\":\"WrappedNativeTokenContractUpdated\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWAL_UNLOCKER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_MAX_PERCENTAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipientAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structTransfer.Request[]\",\"name\":\"_requests\",\"type\":\"tuple[]\"}],\"name\":\"bulkRequestDepositFor\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_voteWeight\",\"type\":\"uint256\"}],\"name\":\"checkThreshold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"dailyWithdrawalLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"fullSigThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mainchainToken\",\"type\":\"address\"}],\"name\":\"getRoninToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_roleSetter\",\"type\":\"address\"},{\"internalType\":\"contractIWETH\",\"name\":\"_wrappedToken\",\"type\":\"address\"},{\"internalType\":\"contractIWeightedValidator\",\"name\":\"_validatorContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_roninChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_denominator\",\"type\":\"uint256\"},{\"internalType\":\"address[][3]\",\"name\":\"_addresses\",\"type\":\"address[][3]\"},{\"internalType\":\"uint256[][4]\",\"name\":\"_thresholds\",\"type\":\"uint256[][4]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastDateSynced\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastSyncedWithdrawal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockedThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_mainchainTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_roninTokens\",\"type\":\"address[]\"}],\"name\":\"mapTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_mainchainTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_roninTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_fullSigThreshold\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_lockedThreshold\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_dailyWithdrawalLimit\",\"type\":\"uint256[]\"}],\"name\":\"mapTokensAndThresholds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumVoteWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"}],\"name\":\"reachedWithdrawalLimit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveEther\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipientAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structTransfer.Request\",\"name\":\"_request\",\"type\":\"tuple\"}],\"name\":\"requestDepositFor\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roninChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_limits\",\"type\":\"uint256[]\"}],\"name\":\"setDailyWithdrawalLimits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_thresholds\",\"type\":\"uint256[]\"}],\"name\":\"setFullSigsThresholds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_thresholds\",\"type\":\"uint256[]\"}],\"name\":\"setLockedThresholds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_denominator\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_percentages\",\"type\":\"uint256[]\"}],\"name\":\"setUnlockFeePercentages\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIWeightedValidator\",\"name\":\"_validatorContract\",\"type\":\"address\"}],\"name\":\"setValidatorContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"_wrappedToken\",\"type\":\"address\"}],\"name\":\"setWrappedNativeTokenContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structTransfer.Receipt\",\"name\":\"_receipt\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"name\":\"submitWithdrawal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_locked\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"unlockFeePercentages\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTransfer.Kind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"mainchain\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Owner\",\"name\":\"ronin\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumToken.Standard\",\"name\":\"erc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"internalType\":\"structToken.Info\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structTransfer.Receipt\",\"name\":\"_receipt\",\"type\":\"tuple\"}],\"name\":\"unlockWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorContract\",\"outputs\":[{\"internalType\":\"contractIWeightedValidator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawalHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawalLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrappedNativeToken\",\"outputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Gateway *GatewayCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Gateway *GatewaySession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Gateway.Contract.DOMAINSEPARATOR(&_Gateway.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Gateway *GatewayCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Gateway.Contract.DOMAINSEPARATOR(&_Gateway.CallOpts)
}

// WITHDRAWALUNLOCKERROLE is a free data retrieval call binding the contract method 0x8f34e347.
//
// Solidity: function WITHDRAWAL_UNLOCKER_ROLE() view returns(bytes32)
func (_Gateway *GatewayCaller) WITHDRAWALUNLOCKERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "WITHDRAWAL_UNLOCKER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWALUNLOCKERROLE is a free data retrieval call binding the contract method 0x8f34e347.
//
// Solidity: function WITHDRAWAL_UNLOCKER_ROLE() view returns(bytes32)
func (_Gateway *GatewaySession) WITHDRAWALUNLOCKERROLE() ([32]byte, error) {
	return _Gateway.Contract.WITHDRAWALUNLOCKERROLE(&_Gateway.CallOpts)
}

// WITHDRAWALUNLOCKERROLE is a free data retrieval call binding the contract method 0x8f34e347.
//
// Solidity: function WITHDRAWAL_UNLOCKER_ROLE() view returns(bytes32)
func (_Gateway *GatewayCallerSession) WITHDRAWALUNLOCKERROLE() ([32]byte, error) {
	return _Gateway.Contract.WITHDRAWALUNLOCKERROLE(&_Gateway.CallOpts)
}

// MAXPERCENTAGE is a free data retrieval call binding the contract method 0x302d12db.
//
// Solidity: function _MAX_PERCENTAGE() view returns(uint256)
func (_Gateway *GatewayCaller) MAXPERCENTAGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "_MAX_PERCENTAGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPERCENTAGE is a free data retrieval call binding the contract method 0x302d12db.
//
// Solidity: function _MAX_PERCENTAGE() view returns(uint256)
func (_Gateway *GatewaySession) MAXPERCENTAGE() (*big.Int, error) {
	return _Gateway.Contract.MAXPERCENTAGE(&_Gateway.CallOpts)
}

// MAXPERCENTAGE is a free data retrieval call binding the contract method 0x302d12db.
//
// Solidity: function _MAX_PERCENTAGE() view returns(uint256)
func (_Gateway *GatewayCallerSession) MAXPERCENTAGE() (*big.Int, error) {
	return _Gateway.Contract.MAXPERCENTAGE(&_Gateway.CallOpts)
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

// DailyWithdrawalLimit is a free data retrieval call binding the contract method 0xab796566.
//
// Solidity: function dailyWithdrawalLimit(address ) view returns(uint256)
func (_Gateway *GatewayCaller) DailyWithdrawalLimit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "dailyWithdrawalLimit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DailyWithdrawalLimit is a free data retrieval call binding the contract method 0xab796566.
//
// Solidity: function dailyWithdrawalLimit(address ) view returns(uint256)
func (_Gateway *GatewaySession) DailyWithdrawalLimit(arg0 common.Address) (*big.Int, error) {
	return _Gateway.Contract.DailyWithdrawalLimit(&_Gateway.CallOpts, arg0)
}

// DailyWithdrawalLimit is a free data retrieval call binding the contract method 0xab796566.
//
// Solidity: function dailyWithdrawalLimit(address ) view returns(uint256)
func (_Gateway *GatewayCallerSession) DailyWithdrawalLimit(arg0 common.Address) (*big.Int, error) {
	return _Gateway.Contract.DailyWithdrawalLimit(&_Gateway.CallOpts, arg0)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Gateway *GatewayCaller) DepositCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "depositCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Gateway *GatewaySession) DepositCount() (*big.Int, error) {
	return _Gateway.Contract.DepositCount(&_Gateway.CallOpts)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Gateway *GatewayCallerSession) DepositCount() (*big.Int, error) {
	return _Gateway.Contract.DepositCount(&_Gateway.CallOpts)
}

// FullSigThreshold is a free data retrieval call binding the contract method 0xdf5baf45.
//
// Solidity: function fullSigThreshold(address ) view returns(uint256)
func (_Gateway *GatewayCaller) FullSigThreshold(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "fullSigThreshold", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FullSigThreshold is a free data retrieval call binding the contract method 0xdf5baf45.
//
// Solidity: function fullSigThreshold(address ) view returns(uint256)
func (_Gateway *GatewaySession) FullSigThreshold(arg0 common.Address) (*big.Int, error) {
	return _Gateway.Contract.FullSigThreshold(&_Gateway.CallOpts, arg0)
}

// FullSigThreshold is a free data retrieval call binding the contract method 0xdf5baf45.
//
// Solidity: function fullSigThreshold(address ) view returns(uint256)
func (_Gateway *GatewayCallerSession) FullSigThreshold(arg0 common.Address) (*big.Int, error) {
	return _Gateway.Contract.FullSigThreshold(&_Gateway.CallOpts, arg0)
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

// GetRoninToken is a free data retrieval call binding the contract method 0xb2975794.
//
// Solidity: function getRoninToken(address _mainchainToken) view returns(address _addr)
func (_Gateway *GatewayCaller) GetRoninToken(opts *bind.CallOpts, _mainchainToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getRoninToken", _mainchainToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoninToken is a free data retrieval call binding the contract method 0xb2975794.
//
// Solidity: function getRoninToken(address _mainchainToken) view returns(address _addr)
func (_Gateway *GatewaySession) GetRoninToken(_mainchainToken common.Address) (common.Address, error) {
	return _Gateway.Contract.GetRoninToken(&_Gateway.CallOpts, _mainchainToken)
}

// GetRoninToken is a free data retrieval call binding the contract method 0xb2975794.
//
// Solidity: function getRoninToken(address _mainchainToken) view returns(address _addr)
func (_Gateway *GatewayCallerSession) GetRoninToken(_mainchainToken common.Address) (common.Address, error) {
	return _Gateway.Contract.GetRoninToken(&_Gateway.CallOpts, _mainchainToken)
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

// LastDateSynced is a free data retrieval call binding the contract method 0xcbc829c6.
//
// Solidity: function lastDateSynced() view returns(uint256)
func (_Gateway *GatewayCaller) LastDateSynced(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "lastDateSynced")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastDateSynced is a free data retrieval call binding the contract method 0xcbc829c6.
//
// Solidity: function lastDateSynced() view returns(uint256)
func (_Gateway *GatewaySession) LastDateSynced() (*big.Int, error) {
	return _Gateway.Contract.LastDateSynced(&_Gateway.CallOpts)
}

// LastDateSynced is a free data retrieval call binding the contract method 0xcbc829c6.
//
// Solidity: function lastDateSynced() view returns(uint256)
func (_Gateway *GatewayCallerSession) LastDateSynced() (*big.Int, error) {
	return _Gateway.Contract.LastDateSynced(&_Gateway.CallOpts)
}

// LastSyncedWithdrawal is a free data retrieval call binding the contract method 0xd55ed103.
//
// Solidity: function lastSyncedWithdrawal(address ) view returns(uint256)
func (_Gateway *GatewayCaller) LastSyncedWithdrawal(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "lastSyncedWithdrawal", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastSyncedWithdrawal is a free data retrieval call binding the contract method 0xd55ed103.
//
// Solidity: function lastSyncedWithdrawal(address ) view returns(uint256)
func (_Gateway *GatewaySession) LastSyncedWithdrawal(arg0 common.Address) (*big.Int, error) {
	return _Gateway.Contract.LastSyncedWithdrawal(&_Gateway.CallOpts, arg0)
}

// LastSyncedWithdrawal is a free data retrieval call binding the contract method 0xd55ed103.
//
// Solidity: function lastSyncedWithdrawal(address ) view returns(uint256)
func (_Gateway *GatewayCallerSession) LastSyncedWithdrawal(arg0 common.Address) (*big.Int, error) {
	return _Gateway.Contract.LastSyncedWithdrawal(&_Gateway.CallOpts, arg0)
}

// LockedThreshold is a free data retrieval call binding the contract method 0x59122f6b.
//
// Solidity: function lockedThreshold(address ) view returns(uint256)
func (_Gateway *GatewayCaller) LockedThreshold(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "lockedThreshold", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedThreshold is a free data retrieval call binding the contract method 0x59122f6b.
//
// Solidity: function lockedThreshold(address ) view returns(uint256)
func (_Gateway *GatewaySession) LockedThreshold(arg0 common.Address) (*big.Int, error) {
	return _Gateway.Contract.LockedThreshold(&_Gateway.CallOpts, arg0)
}

// LockedThreshold is a free data retrieval call binding the contract method 0x59122f6b.
//
// Solidity: function lockedThreshold(address ) view returns(uint256)
func (_Gateway *GatewayCallerSession) LockedThreshold(arg0 common.Address) (*big.Int, error) {
	return _Gateway.Contract.LockedThreshold(&_Gateway.CallOpts, arg0)
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

// ReachedWithdrawalLimit is a free data retrieval call binding the contract method 0x6c1ce670.
//
// Solidity: function reachedWithdrawalLimit(address _token, uint256 _quantity) view returns(bool)
func (_Gateway *GatewayCaller) ReachedWithdrawalLimit(opts *bind.CallOpts, _token common.Address, _quantity *big.Int) (bool, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "reachedWithdrawalLimit", _token, _quantity)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ReachedWithdrawalLimit is a free data retrieval call binding the contract method 0x6c1ce670.
//
// Solidity: function reachedWithdrawalLimit(address _token, uint256 _quantity) view returns(bool)
func (_Gateway *GatewaySession) ReachedWithdrawalLimit(_token common.Address, _quantity *big.Int) (bool, error) {
	return _Gateway.Contract.ReachedWithdrawalLimit(&_Gateway.CallOpts, _token, _quantity)
}

// ReachedWithdrawalLimit is a free data retrieval call binding the contract method 0x6c1ce670.
//
// Solidity: function reachedWithdrawalLimit(address _token, uint256 _quantity) view returns(bool)
func (_Gateway *GatewayCallerSession) ReachedWithdrawalLimit(_token common.Address, _quantity *big.Int) (bool, error) {
	return _Gateway.Contract.ReachedWithdrawalLimit(&_Gateway.CallOpts, _token, _quantity)
}

// RoninChainId is a free data retrieval call binding the contract method 0x17ce2dd4.
//
// Solidity: function roninChainId() view returns(uint256)
func (_Gateway *GatewayCaller) RoninChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "roninChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoninChainId is a free data retrieval call binding the contract method 0x17ce2dd4.
//
// Solidity: function roninChainId() view returns(uint256)
func (_Gateway *GatewaySession) RoninChainId() (*big.Int, error) {
	return _Gateway.Contract.RoninChainId(&_Gateway.CallOpts)
}

// RoninChainId is a free data retrieval call binding the contract method 0x17ce2dd4.
//
// Solidity: function roninChainId() view returns(uint256)
func (_Gateway *GatewayCallerSession) RoninChainId() (*big.Int, error) {
	return _Gateway.Contract.RoninChainId(&_Gateway.CallOpts)
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

// UnlockFeePercentages is a free data retrieval call binding the contract method 0xd19773d2.
//
// Solidity: function unlockFeePercentages(address ) view returns(uint256)
func (_Gateway *GatewayCaller) UnlockFeePercentages(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "unlockFeePercentages", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnlockFeePercentages is a free data retrieval call binding the contract method 0xd19773d2.
//
// Solidity: function unlockFeePercentages(address ) view returns(uint256)
func (_Gateway *GatewaySession) UnlockFeePercentages(arg0 common.Address) (*big.Int, error) {
	return _Gateway.Contract.UnlockFeePercentages(&_Gateway.CallOpts, arg0)
}

// UnlockFeePercentages is a free data retrieval call binding the contract method 0xd19773d2.
//
// Solidity: function unlockFeePercentages(address ) view returns(uint256)
func (_Gateway *GatewayCallerSession) UnlockFeePercentages(arg0 common.Address) (*big.Int, error) {
	return _Gateway.Contract.UnlockFeePercentages(&_Gateway.CallOpts, arg0)
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

// WithdrawalHash is a free data retrieval call binding the contract method 0x6932be98.
//
// Solidity: function withdrawalHash(uint256 ) view returns(bytes32)
func (_Gateway *GatewayCaller) WithdrawalHash(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "withdrawalHash", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WithdrawalHash is a free data retrieval call binding the contract method 0x6932be98.
//
// Solidity: function withdrawalHash(uint256 ) view returns(bytes32)
func (_Gateway *GatewaySession) WithdrawalHash(arg0 *big.Int) ([32]byte, error) {
	return _Gateway.Contract.WithdrawalHash(&_Gateway.CallOpts, arg0)
}

// WithdrawalHash is a free data retrieval call binding the contract method 0x6932be98.
//
// Solidity: function withdrawalHash(uint256 ) view returns(bytes32)
func (_Gateway *GatewayCallerSession) WithdrawalHash(arg0 *big.Int) ([32]byte, error) {
	return _Gateway.Contract.WithdrawalHash(&_Gateway.CallOpts, arg0)
}

// WithdrawalLocked is a free data retrieval call binding the contract method 0x4d493f4e.
//
// Solidity: function withdrawalLocked(uint256 ) view returns(bool)
func (_Gateway *GatewayCaller) WithdrawalLocked(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "withdrawalLocked", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawalLocked is a free data retrieval call binding the contract method 0x4d493f4e.
//
// Solidity: function withdrawalLocked(uint256 ) view returns(bool)
func (_Gateway *GatewaySession) WithdrawalLocked(arg0 *big.Int) (bool, error) {
	return _Gateway.Contract.WithdrawalLocked(&_Gateway.CallOpts, arg0)
}

// WithdrawalLocked is a free data retrieval call binding the contract method 0x4d493f4e.
//
// Solidity: function withdrawalLocked(uint256 ) view returns(bool)
func (_Gateway *GatewayCallerSession) WithdrawalLocked(arg0 *big.Int) (bool, error) {
	return _Gateway.Contract.WithdrawalLocked(&_Gateway.CallOpts, arg0)
}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_Gateway *GatewayCaller) WrappedNativeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "wrappedNativeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_Gateway *GatewaySession) WrappedNativeToken() (common.Address, error) {
	return _Gateway.Contract.WrappedNativeToken(&_Gateway.CallOpts)
}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_Gateway *GatewayCallerSession) WrappedNativeToken() (common.Address, error) {
	return _Gateway.Contract.WrappedNativeToken(&_Gateway.CallOpts)
}

// BulkRequestDepositFor is a paid mutator transaction binding the contract method 0xd2d9114f.
//
// Solidity: function bulkRequestDepositFor((address,address,(uint8,uint256,uint256))[] _requests) payable returns()
func (_Gateway *GatewayTransactor) BulkRequestDepositFor(opts *bind.TransactOpts, _requests []TransferRequest) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "bulkRequestDepositFor", _requests)
}

// BulkRequestDepositFor is a paid mutator transaction binding the contract method 0xd2d9114f.
//
// Solidity: function bulkRequestDepositFor((address,address,(uint8,uint256,uint256))[] _requests) payable returns()
func (_Gateway *GatewaySession) BulkRequestDepositFor(_requests []TransferRequest) (*types.Transaction, error) {
	return _Gateway.Contract.BulkRequestDepositFor(&_Gateway.TransactOpts, _requests)
}

// BulkRequestDepositFor is a paid mutator transaction binding the contract method 0xd2d9114f.
//
// Solidity: function bulkRequestDepositFor((address,address,(uint8,uint256,uint256))[] _requests) payable returns()
func (_Gateway *GatewayTransactorSession) BulkRequestDepositFor(_requests []TransferRequest) (*types.Transaction, error) {
	return _Gateway.Contract.BulkRequestDepositFor(&_Gateway.TransactOpts, _requests)
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

// Initialize is a paid mutator transaction binding the contract method 0x163d245b.
//
// Solidity: function initialize(address _roleSetter, address _wrappedToken, address _validatorContract, uint256 _roninChainId, uint256 _numerator, uint256 _denominator, address[][3] _addresses, uint256[][4] _thresholds) payable returns()
func (_Gateway *GatewayTransactor) Initialize(opts *bind.TransactOpts, _roleSetter common.Address, _wrappedToken common.Address, _validatorContract common.Address, _roninChainId *big.Int, _numerator *big.Int, _denominator *big.Int, _addresses [3][]common.Address, _thresholds [4][]*big.Int) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "initialize", _roleSetter, _wrappedToken, _validatorContract, _roninChainId, _numerator, _denominator, _addresses, _thresholds)
}

// Initialize is a paid mutator transaction binding the contract method 0x163d245b.
//
// Solidity: function initialize(address _roleSetter, address _wrappedToken, address _validatorContract, uint256 _roninChainId, uint256 _numerator, uint256 _denominator, address[][3] _addresses, uint256[][4] _thresholds) payable returns()
func (_Gateway *GatewaySession) Initialize(_roleSetter common.Address, _wrappedToken common.Address, _validatorContract common.Address, _roninChainId *big.Int, _numerator *big.Int, _denominator *big.Int, _addresses [3][]common.Address, _thresholds [4][]*big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.Initialize(&_Gateway.TransactOpts, _roleSetter, _wrappedToken, _validatorContract, _roninChainId, _numerator, _denominator, _addresses, _thresholds)
}

// Initialize is a paid mutator transaction binding the contract method 0x163d245b.
//
// Solidity: function initialize(address _roleSetter, address _wrappedToken, address _validatorContract, uint256 _roninChainId, uint256 _numerator, uint256 _denominator, address[][3] _addresses, uint256[][4] _thresholds) payable returns()
func (_Gateway *GatewayTransactorSession) Initialize(_roleSetter common.Address, _wrappedToken common.Address, _validatorContract common.Address, _roninChainId *big.Int, _numerator *big.Int, _denominator *big.Int, _addresses [3][]common.Address, _thresholds [4][]*big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.Initialize(&_Gateway.TransactOpts, _roleSetter, _wrappedToken, _validatorContract, _roninChainId, _numerator, _denominator, _addresses, _thresholds)
}

// MapTokens is a paid mutator transaction binding the contract method 0xfc73b5d0.
//
// Solidity: function mapTokens(address[] _mainchainTokens, address[] _roninTokens) returns()
func (_Gateway *GatewayTransactor) MapTokens(opts *bind.TransactOpts, _mainchainTokens []common.Address, _roninTokens []common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "mapTokens", _mainchainTokens, _roninTokens)
}

// MapTokens is a paid mutator transaction binding the contract method 0xfc73b5d0.
//
// Solidity: function mapTokens(address[] _mainchainTokens, address[] _roninTokens) returns()
func (_Gateway *GatewaySession) MapTokens(_mainchainTokens []common.Address, _roninTokens []common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.MapTokens(&_Gateway.TransactOpts, _mainchainTokens, _roninTokens)
}

// MapTokens is a paid mutator transaction binding the contract method 0xfc73b5d0.
//
// Solidity: function mapTokens(address[] _mainchainTokens, address[] _roninTokens) returns()
func (_Gateway *GatewayTransactorSession) MapTokens(_mainchainTokens []common.Address, _roninTokens []common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.MapTokens(&_Gateway.TransactOpts, _mainchainTokens, _roninTokens)
}

// MapTokensAndThresholds is a paid mutator transaction binding the contract method 0x056f7781.
//
// Solidity: function mapTokensAndThresholds(address[] _mainchainTokens, address[] _roninTokens, uint256[] _fullSigThreshold, uint256[] _lockedThreshold, uint256[] _dailyWithdrawalLimit) returns()
func (_Gateway *GatewayTransactor) MapTokensAndThresholds(opts *bind.TransactOpts, _mainchainTokens []common.Address, _roninTokens []common.Address, _fullSigThreshold []*big.Int, _lockedThreshold []*big.Int, _dailyWithdrawalLimit []*big.Int) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "mapTokensAndThresholds", _mainchainTokens, _roninTokens, _fullSigThreshold, _lockedThreshold, _dailyWithdrawalLimit)
}

// MapTokensAndThresholds is a paid mutator transaction binding the contract method 0x056f7781.
//
// Solidity: function mapTokensAndThresholds(address[] _mainchainTokens, address[] _roninTokens, uint256[] _fullSigThreshold, uint256[] _lockedThreshold, uint256[] _dailyWithdrawalLimit) returns()
func (_Gateway *GatewaySession) MapTokensAndThresholds(_mainchainTokens []common.Address, _roninTokens []common.Address, _fullSigThreshold []*big.Int, _lockedThreshold []*big.Int, _dailyWithdrawalLimit []*big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.MapTokensAndThresholds(&_Gateway.TransactOpts, _mainchainTokens, _roninTokens, _fullSigThreshold, _lockedThreshold, _dailyWithdrawalLimit)
}

// MapTokensAndThresholds is a paid mutator transaction binding the contract method 0x056f7781.
//
// Solidity: function mapTokensAndThresholds(address[] _mainchainTokens, address[] _roninTokens, uint256[] _fullSigThreshold, uint256[] _lockedThreshold, uint256[] _dailyWithdrawalLimit) returns()
func (_Gateway *GatewayTransactorSession) MapTokensAndThresholds(_mainchainTokens []common.Address, _roninTokens []common.Address, _fullSigThreshold []*big.Int, _lockedThreshold []*big.Int, _dailyWithdrawalLimit []*big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.MapTokensAndThresholds(&_Gateway.TransactOpts, _mainchainTokens, _roninTokens, _fullSigThreshold, _lockedThreshold, _dailyWithdrawalLimit)
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

// ReceiveEther is a paid mutator transaction binding the contract method 0xa3912ec8.
//
// Solidity: function receiveEther() payable returns()
func (_Gateway *GatewayTransactor) ReceiveEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "receiveEther")
}

// ReceiveEther is a paid mutator transaction binding the contract method 0xa3912ec8.
//
// Solidity: function receiveEther() payable returns()
func (_Gateway *GatewaySession) ReceiveEther() (*types.Transaction, error) {
	return _Gateway.Contract.ReceiveEther(&_Gateway.TransactOpts)
}

// ReceiveEther is a paid mutator transaction binding the contract method 0xa3912ec8.
//
// Solidity: function receiveEther() payable returns()
func (_Gateway *GatewayTransactorSession) ReceiveEther() (*types.Transaction, error) {
	return _Gateway.Contract.ReceiveEther(&_Gateway.TransactOpts)
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

// RequestDepositFor is a paid mutator transaction binding the contract method 0x4b14557e.
//
// Solidity: function requestDepositFor((address,address,(uint8,uint256,uint256)) _request) payable returns()
func (_Gateway *GatewayTransactor) RequestDepositFor(opts *bind.TransactOpts, _request TransferRequest) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "requestDepositFor", _request)
}

// RequestDepositFor is a paid mutator transaction binding the contract method 0x4b14557e.
//
// Solidity: function requestDepositFor((address,address,(uint8,uint256,uint256)) _request) payable returns()
func (_Gateway *GatewaySession) RequestDepositFor(_request TransferRequest) (*types.Transaction, error) {
	return _Gateway.Contract.RequestDepositFor(&_Gateway.TransactOpts, _request)
}

// RequestDepositFor is a paid mutator transaction binding the contract method 0x4b14557e.
//
// Solidity: function requestDepositFor((address,address,(uint8,uint256,uint256)) _request) payable returns()
func (_Gateway *GatewayTransactorSession) RequestDepositFor(_request TransferRequest) (*types.Transaction, error) {
	return _Gateway.Contract.RequestDepositFor(&_Gateway.TransactOpts, _request)
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

// SetDailyWithdrawalLimits is a paid mutator transaction binding the contract method 0xe400327c.
//
// Solidity: function setDailyWithdrawalLimits(address[] _tokens, uint256[] _limits) returns()
func (_Gateway *GatewayTransactor) SetDailyWithdrawalLimits(opts *bind.TransactOpts, _tokens []common.Address, _limits []*big.Int) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "setDailyWithdrawalLimits", _tokens, _limits)
}

// SetDailyWithdrawalLimits is a paid mutator transaction binding the contract method 0xe400327c.
//
// Solidity: function setDailyWithdrawalLimits(address[] _tokens, uint256[] _limits) returns()
func (_Gateway *GatewaySession) SetDailyWithdrawalLimits(_tokens []common.Address, _limits []*big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.SetDailyWithdrawalLimits(&_Gateway.TransactOpts, _tokens, _limits)
}

// SetDailyWithdrawalLimits is a paid mutator transaction binding the contract method 0xe400327c.
//
// Solidity: function setDailyWithdrawalLimits(address[] _tokens, uint256[] _limits) returns()
func (_Gateway *GatewayTransactorSession) SetDailyWithdrawalLimits(_tokens []common.Address, _limits []*big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.SetDailyWithdrawalLimits(&_Gateway.TransactOpts, _tokens, _limits)
}

// SetFullSigsThresholds is a paid mutator transaction binding the contract method 0xd8356de7.
//
// Solidity: function setFullSigsThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_Gateway *GatewayTransactor) SetFullSigsThresholds(opts *bind.TransactOpts, _tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "setFullSigsThresholds", _tokens, _thresholds)
}

// SetFullSigsThresholds is a paid mutator transaction binding the contract method 0xd8356de7.
//
// Solidity: function setFullSigsThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_Gateway *GatewaySession) SetFullSigsThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.SetFullSigsThresholds(&_Gateway.TransactOpts, _tokens, _thresholds)
}

// SetFullSigsThresholds is a paid mutator transaction binding the contract method 0xd8356de7.
//
// Solidity: function setFullSigsThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_Gateway *GatewayTransactorSession) SetFullSigsThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.SetFullSigsThresholds(&_Gateway.TransactOpts, _tokens, _thresholds)
}

// SetLockedThresholds is a paid mutator transaction binding the contract method 0x1a8e55b0.
//
// Solidity: function setLockedThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_Gateway *GatewayTransactor) SetLockedThresholds(opts *bind.TransactOpts, _tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "setLockedThresholds", _tokens, _thresholds)
}

// SetLockedThresholds is a paid mutator transaction binding the contract method 0x1a8e55b0.
//
// Solidity: function setLockedThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_Gateway *GatewaySession) SetLockedThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.SetLockedThresholds(&_Gateway.TransactOpts, _tokens, _thresholds)
}

// SetLockedThresholds is a paid mutator transaction binding the contract method 0x1a8e55b0.
//
// Solidity: function setLockedThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_Gateway *GatewayTransactorSession) SetLockedThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.SetLockedThresholds(&_Gateway.TransactOpts, _tokens, _thresholds)
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

// SetUnlockFeePercentages is a paid mutator transaction binding the contract method 0xb1a2567e.
//
// Solidity: function setUnlockFeePercentages(address[] _tokens, uint256[] _percentages) returns()
func (_Gateway *GatewayTransactor) SetUnlockFeePercentages(opts *bind.TransactOpts, _tokens []common.Address, _percentages []*big.Int) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "setUnlockFeePercentages", _tokens, _percentages)
}

// SetUnlockFeePercentages is a paid mutator transaction binding the contract method 0xb1a2567e.
//
// Solidity: function setUnlockFeePercentages(address[] _tokens, uint256[] _percentages) returns()
func (_Gateway *GatewaySession) SetUnlockFeePercentages(_tokens []common.Address, _percentages []*big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.SetUnlockFeePercentages(&_Gateway.TransactOpts, _tokens, _percentages)
}

// SetUnlockFeePercentages is a paid mutator transaction binding the contract method 0xb1a2567e.
//
// Solidity: function setUnlockFeePercentages(address[] _tokens, uint256[] _percentages) returns()
func (_Gateway *GatewayTransactorSession) SetUnlockFeePercentages(_tokens []common.Address, _percentages []*big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.SetUnlockFeePercentages(&_Gateway.TransactOpts, _tokens, _percentages)
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

// SetWrappedNativeTokenContract is a paid mutator transaction binding the contract method 0xd64af2a6.
//
// Solidity: function setWrappedNativeTokenContract(address _wrappedToken) returns()
func (_Gateway *GatewayTransactor) SetWrappedNativeTokenContract(opts *bind.TransactOpts, _wrappedToken common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "setWrappedNativeTokenContract", _wrappedToken)
}

// SetWrappedNativeTokenContract is a paid mutator transaction binding the contract method 0xd64af2a6.
//
// Solidity: function setWrappedNativeTokenContract(address _wrappedToken) returns()
func (_Gateway *GatewaySession) SetWrappedNativeTokenContract(_wrappedToken common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.SetWrappedNativeTokenContract(&_Gateway.TransactOpts, _wrappedToken)
}

// SetWrappedNativeTokenContract is a paid mutator transaction binding the contract method 0xd64af2a6.
//
// Solidity: function setWrappedNativeTokenContract(address _wrappedToken) returns()
func (_Gateway *GatewayTransactorSession) SetWrappedNativeTokenContract(_wrappedToken common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.SetWrappedNativeTokenContract(&_Gateway.TransactOpts, _wrappedToken)
}

// SubmitWithdrawal is a paid mutator transaction binding the contract method 0x4d0d6673.
//
// Solidity: function submitWithdrawal((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt, (uint8,bytes32,bytes32)[] _signatures) returns(bool _locked)
func (_Gateway *GatewayTransactor) SubmitWithdrawal(opts *bind.TransactOpts, _receipt TransferReceipt, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "submitWithdrawal", _receipt, _signatures)
}

// SubmitWithdrawal is a paid mutator transaction binding the contract method 0x4d0d6673.
//
// Solidity: function submitWithdrawal((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt, (uint8,bytes32,bytes32)[] _signatures) returns(bool _locked)
func (_Gateway *GatewaySession) SubmitWithdrawal(_receipt TransferReceipt, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Gateway.Contract.SubmitWithdrawal(&_Gateway.TransactOpts, _receipt, _signatures)
}

// SubmitWithdrawal is a paid mutator transaction binding the contract method 0x4d0d6673.
//
// Solidity: function submitWithdrawal((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt, (uint8,bytes32,bytes32)[] _signatures) returns(bool _locked)
func (_Gateway *GatewayTransactorSession) SubmitWithdrawal(_receipt TransferReceipt, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Gateway.Contract.SubmitWithdrawal(&_Gateway.TransactOpts, _receipt, _signatures)
}

// UnlockWithdrawal is a paid mutator transaction binding the contract method 0x9157921c.
//
// Solidity: function unlockWithdrawal((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt) returns()
func (_Gateway *GatewayTransactor) UnlockWithdrawal(opts *bind.TransactOpts, _receipt TransferReceipt) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "unlockWithdrawal", _receipt)
}

// UnlockWithdrawal is a paid mutator transaction binding the contract method 0x9157921c.
//
// Solidity: function unlockWithdrawal((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt) returns()
func (_Gateway *GatewaySession) UnlockWithdrawal(_receipt TransferReceipt) (*types.Transaction, error) {
	return _Gateway.Contract.UnlockWithdrawal(&_Gateway.TransactOpts, _receipt)
}

// UnlockWithdrawal is a paid mutator transaction binding the contract method 0x9157921c.
//
// Solidity: function unlockWithdrawal((uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) _receipt) returns()
func (_Gateway *GatewayTransactorSession) UnlockWithdrawal(_receipt TransferReceipt) (*types.Transaction, error) {
	return _Gateway.Contract.UnlockWithdrawal(&_Gateway.TransactOpts, _receipt)
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

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Gateway *GatewayTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Gateway.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Gateway *GatewaySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Gateway.Contract.Fallback(&_Gateway.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Gateway *GatewayTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Gateway.Contract.Fallback(&_Gateway.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Gateway *GatewayTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Gateway *GatewaySession) Receive() (*types.Transaction, error) {
	return _Gateway.Contract.Receive(&_Gateway.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Gateway *GatewayTransactorSession) Receive() (*types.Transaction, error) {
	return _Gateway.Contract.Receive(&_Gateway.TransactOpts)
}

// GatewayDailyWithdrawalLimitsUpdatedIterator is returned from FilterDailyWithdrawalLimitsUpdated and is used to iterate over the raw logs and unpacked data for DailyWithdrawalLimitsUpdated events raised by the Gateway contract.
type GatewayDailyWithdrawalLimitsUpdatedIterator struct {
	Event *GatewayDailyWithdrawalLimitsUpdated // Event containing the contract specifics and raw log

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
func (it *GatewayDailyWithdrawalLimitsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayDailyWithdrawalLimitsUpdated)
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
		it.Event = new(GatewayDailyWithdrawalLimitsUpdated)
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
func (it *GatewayDailyWithdrawalLimitsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayDailyWithdrawalLimitsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayDailyWithdrawalLimitsUpdated represents a DailyWithdrawalLimitsUpdated event raised by the Gateway contract.
type GatewayDailyWithdrawalLimitsUpdated struct {
	Tokens []common.Address
	Limits []*big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDailyWithdrawalLimitsUpdated is a free log retrieval operation binding the contract event 0xb5d2963614d72181b4df1f993d45b83edf42fa19710f0204217ba1b3e183bb73.
//
// Solidity: event DailyWithdrawalLimitsUpdated(address[] tokens, uint256[] limits)
func (_Gateway *GatewayFilterer) FilterDailyWithdrawalLimitsUpdated(opts *bind.FilterOpts) (*GatewayDailyWithdrawalLimitsUpdatedIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "DailyWithdrawalLimitsUpdated")
	if err != nil {
		return nil, err
	}
	return &GatewayDailyWithdrawalLimitsUpdatedIterator{contract: _Gateway.contract, event: "DailyWithdrawalLimitsUpdated", logs: logs, sub: sub}, nil
}

// WatchDailyWithdrawalLimitsUpdated is a free log subscription operation binding the contract event 0xb5d2963614d72181b4df1f993d45b83edf42fa19710f0204217ba1b3e183bb73.
//
// Solidity: event DailyWithdrawalLimitsUpdated(address[] tokens, uint256[] limits)
func (_Gateway *GatewayFilterer) WatchDailyWithdrawalLimitsUpdated(opts *bind.WatchOpts, sink chan<- *GatewayDailyWithdrawalLimitsUpdated) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "DailyWithdrawalLimitsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayDailyWithdrawalLimitsUpdated)
				if err := _Gateway.contract.UnpackLog(event, "DailyWithdrawalLimitsUpdated", log); err != nil {
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
func (_Gateway *GatewayFilterer) ParseDailyWithdrawalLimitsUpdated(log types.Log) (*GatewayDailyWithdrawalLimitsUpdated, error) {
	event := new(GatewayDailyWithdrawalLimitsUpdated)
	if err := _Gateway.contract.UnpackLog(event, "DailyWithdrawalLimitsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayDepositRequestedIterator is returned from FilterDepositRequested and is used to iterate over the raw logs and unpacked data for DepositRequested events raised by the Gateway contract.
type GatewayDepositRequestedIterator struct {
	Event *GatewayDepositRequested // Event containing the contract specifics and raw log

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
func (it *GatewayDepositRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayDepositRequested)
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
		it.Event = new(GatewayDepositRequested)
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
func (it *GatewayDepositRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayDepositRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayDepositRequested represents a DepositRequested event raised by the Gateway contract.
type GatewayDepositRequested struct {
	ReceiptHash [32]byte
	Receipt     TransferReceipt
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDepositRequested is a free log retrieval operation binding the contract event 0xd7b25068d9dc8d00765254cfb7f5070f98d263c8d68931d937c7362fa738048b.
//
// Solidity: event DepositRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_Gateway *GatewayFilterer) FilterDepositRequested(opts *bind.FilterOpts) (*GatewayDepositRequestedIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "DepositRequested")
	if err != nil {
		return nil, err
	}
	return &GatewayDepositRequestedIterator{contract: _Gateway.contract, event: "DepositRequested", logs: logs, sub: sub}, nil
}

// WatchDepositRequested is a free log subscription operation binding the contract event 0xd7b25068d9dc8d00765254cfb7f5070f98d263c8d68931d937c7362fa738048b.
//
// Solidity: event DepositRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_Gateway *GatewayFilterer) WatchDepositRequested(opts *bind.WatchOpts, sink chan<- *GatewayDepositRequested) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "DepositRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayDepositRequested)
				if err := _Gateway.contract.UnpackLog(event, "DepositRequested", log); err != nil {
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
// Solidity: event DepositRequested(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_Gateway *GatewayFilterer) ParseDepositRequested(log types.Log) (*GatewayDepositRequested, error) {
	event := new(GatewayDepositRequested)
	if err := _Gateway.contract.UnpackLog(event, "DepositRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayFullSigsThresholdsUpdatedIterator is returned from FilterFullSigsThresholdsUpdated and is used to iterate over the raw logs and unpacked data for FullSigsThresholdsUpdated events raised by the Gateway contract.
type GatewayFullSigsThresholdsUpdatedIterator struct {
	Event *GatewayFullSigsThresholdsUpdated // Event containing the contract specifics and raw log

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
func (it *GatewayFullSigsThresholdsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayFullSigsThresholdsUpdated)
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
		it.Event = new(GatewayFullSigsThresholdsUpdated)
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
func (it *GatewayFullSigsThresholdsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayFullSigsThresholdsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayFullSigsThresholdsUpdated represents a FullSigsThresholdsUpdated event raised by the Gateway contract.
type GatewayFullSigsThresholdsUpdated struct {
	Tokens     []common.Address
	Thresholds []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFullSigsThresholdsUpdated is a free log retrieval operation binding the contract event 0x30f30fe53f33a6b009d6d0446c37f11eff8aa1033a9a92df9e8fda478cb768f7.
//
// Solidity: event FullSigsThresholdsUpdated(address[] tokens, uint256[] thresholds)
func (_Gateway *GatewayFilterer) FilterFullSigsThresholdsUpdated(opts *bind.FilterOpts) (*GatewayFullSigsThresholdsUpdatedIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "FullSigsThresholdsUpdated")
	if err != nil {
		return nil, err
	}
	return &GatewayFullSigsThresholdsUpdatedIterator{contract: _Gateway.contract, event: "FullSigsThresholdsUpdated", logs: logs, sub: sub}, nil
}

// WatchFullSigsThresholdsUpdated is a free log subscription operation binding the contract event 0x30f30fe53f33a6b009d6d0446c37f11eff8aa1033a9a92df9e8fda478cb768f7.
//
// Solidity: event FullSigsThresholdsUpdated(address[] tokens, uint256[] thresholds)
func (_Gateway *GatewayFilterer) WatchFullSigsThresholdsUpdated(opts *bind.WatchOpts, sink chan<- *GatewayFullSigsThresholdsUpdated) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "FullSigsThresholdsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayFullSigsThresholdsUpdated)
				if err := _Gateway.contract.UnpackLog(event, "FullSigsThresholdsUpdated", log); err != nil {
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
func (_Gateway *GatewayFilterer) ParseFullSigsThresholdsUpdated(log types.Log) (*GatewayFullSigsThresholdsUpdated, error) {
	event := new(GatewayFullSigsThresholdsUpdated)
	if err := _Gateway.contract.UnpackLog(event, "FullSigsThresholdsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayLockedThresholdsUpdatedIterator is returned from FilterLockedThresholdsUpdated and is used to iterate over the raw logs and unpacked data for LockedThresholdsUpdated events raised by the Gateway contract.
type GatewayLockedThresholdsUpdatedIterator struct {
	Event *GatewayLockedThresholdsUpdated // Event containing the contract specifics and raw log

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
func (it *GatewayLockedThresholdsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayLockedThresholdsUpdated)
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
		it.Event = new(GatewayLockedThresholdsUpdated)
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
func (it *GatewayLockedThresholdsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayLockedThresholdsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayLockedThresholdsUpdated represents a LockedThresholdsUpdated event raised by the Gateway contract.
type GatewayLockedThresholdsUpdated struct {
	Tokens     []common.Address
	Thresholds []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLockedThresholdsUpdated is a free log retrieval operation binding the contract event 0x64557254143204d91ba2d95acb9fda1e5fea55f77efd028685765bc1e94dd4b5.
//
// Solidity: event LockedThresholdsUpdated(address[] tokens, uint256[] thresholds)
func (_Gateway *GatewayFilterer) FilterLockedThresholdsUpdated(opts *bind.FilterOpts) (*GatewayLockedThresholdsUpdatedIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "LockedThresholdsUpdated")
	if err != nil {
		return nil, err
	}
	return &GatewayLockedThresholdsUpdatedIterator{contract: _Gateway.contract, event: "LockedThresholdsUpdated", logs: logs, sub: sub}, nil
}

// WatchLockedThresholdsUpdated is a free log subscription operation binding the contract event 0x64557254143204d91ba2d95acb9fda1e5fea55f77efd028685765bc1e94dd4b5.
//
// Solidity: event LockedThresholdsUpdated(address[] tokens, uint256[] thresholds)
func (_Gateway *GatewayFilterer) WatchLockedThresholdsUpdated(opts *bind.WatchOpts, sink chan<- *GatewayLockedThresholdsUpdated) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "LockedThresholdsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayLockedThresholdsUpdated)
				if err := _Gateway.contract.UnpackLog(event, "LockedThresholdsUpdated", log); err != nil {
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
func (_Gateway *GatewayFilterer) ParseLockedThresholdsUpdated(log types.Log) (*GatewayLockedThresholdsUpdated, error) {
	event := new(GatewayLockedThresholdsUpdated)
	if err := _Gateway.contract.UnpackLog(event, "LockedThresholdsUpdated", log); err != nil {
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
	MainchainTokens []common.Address
	RoninTokens     []common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokenMapped is a free log retrieval operation binding the contract event 0x1b7fd57807f95645ee3e0d5a22377a7779e93655cd5f6871151ae7129df015b3.
//
// Solidity: event TokenMapped(address[] mainchainTokens, address[] roninTokens)
func (_Gateway *GatewayFilterer) FilterTokenMapped(opts *bind.FilterOpts) (*GatewayTokenMappedIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "TokenMapped")
	if err != nil {
		return nil, err
	}
	return &GatewayTokenMappedIterator{contract: _Gateway.contract, event: "TokenMapped", logs: logs, sub: sub}, nil
}

// WatchTokenMapped is a free log subscription operation binding the contract event 0x1b7fd57807f95645ee3e0d5a22377a7779e93655cd5f6871151ae7129df015b3.
//
// Solidity: event TokenMapped(address[] mainchainTokens, address[] roninTokens)
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

// ParseTokenMapped is a log parse operation binding the contract event 0x1b7fd57807f95645ee3e0d5a22377a7779e93655cd5f6871151ae7129df015b3.
//
// Solidity: event TokenMapped(address[] mainchainTokens, address[] roninTokens)
func (_Gateway *GatewayFilterer) ParseTokenMapped(log types.Log) (*GatewayTokenMapped, error) {
	event := new(GatewayTokenMapped)
	if err := _Gateway.contract.UnpackLog(event, "TokenMapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayUnlockFeePercentagesUpdatedIterator is returned from FilterUnlockFeePercentagesUpdated and is used to iterate over the raw logs and unpacked data for UnlockFeePercentagesUpdated events raised by the Gateway contract.
type GatewayUnlockFeePercentagesUpdatedIterator struct {
	Event *GatewayUnlockFeePercentagesUpdated // Event containing the contract specifics and raw log

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
func (it *GatewayUnlockFeePercentagesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayUnlockFeePercentagesUpdated)
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
		it.Event = new(GatewayUnlockFeePercentagesUpdated)
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
func (it *GatewayUnlockFeePercentagesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayUnlockFeePercentagesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayUnlockFeePercentagesUpdated represents a UnlockFeePercentagesUpdated event raised by the Gateway contract.
type GatewayUnlockFeePercentagesUpdated struct {
	Tokens      []common.Address
	Percentages []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnlockFeePercentagesUpdated is a free log retrieval operation binding the contract event 0xb05f5de88ae0294ebb6f67c5af2fcbbd593cc6bdfe543e2869794a4c8ce3ea50.
//
// Solidity: event UnlockFeePercentagesUpdated(address[] tokens, uint256[] percentages)
func (_Gateway *GatewayFilterer) FilterUnlockFeePercentagesUpdated(opts *bind.FilterOpts) (*GatewayUnlockFeePercentagesUpdatedIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "UnlockFeePercentagesUpdated")
	if err != nil {
		return nil, err
	}
	return &GatewayUnlockFeePercentagesUpdatedIterator{contract: _Gateway.contract, event: "UnlockFeePercentagesUpdated", logs: logs, sub: sub}, nil
}

// WatchUnlockFeePercentagesUpdated is a free log subscription operation binding the contract event 0xb05f5de88ae0294ebb6f67c5af2fcbbd593cc6bdfe543e2869794a4c8ce3ea50.
//
// Solidity: event UnlockFeePercentagesUpdated(address[] tokens, uint256[] percentages)
func (_Gateway *GatewayFilterer) WatchUnlockFeePercentagesUpdated(opts *bind.WatchOpts, sink chan<- *GatewayUnlockFeePercentagesUpdated) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "UnlockFeePercentagesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayUnlockFeePercentagesUpdated)
				if err := _Gateway.contract.UnpackLog(event, "UnlockFeePercentagesUpdated", log); err != nil {
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

// ParseUnlockFeePercentagesUpdated is a log parse operation binding the contract event 0xb05f5de88ae0294ebb6f67c5af2fcbbd593cc6bdfe543e2869794a4c8ce3ea50.
//
// Solidity: event UnlockFeePercentagesUpdated(address[] tokens, uint256[] percentages)
func (_Gateway *GatewayFilterer) ParseUnlockFeePercentagesUpdated(log types.Log) (*GatewayUnlockFeePercentagesUpdated, error) {
	event := new(GatewayUnlockFeePercentagesUpdated)
	if err := _Gateway.contract.UnpackLog(event, "UnlockFeePercentagesUpdated", log); err != nil {
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

// GatewayWithdrawalLockedIterator is returned from FilterWithdrawalLocked and is used to iterate over the raw logs and unpacked data for WithdrawalLocked events raised by the Gateway contract.
type GatewayWithdrawalLockedIterator struct {
	Event *GatewayWithdrawalLocked // Event containing the contract specifics and raw log

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
func (it *GatewayWithdrawalLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWithdrawalLocked)
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
		it.Event = new(GatewayWithdrawalLocked)
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
func (it *GatewayWithdrawalLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWithdrawalLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWithdrawalLocked represents a WithdrawalLocked event raised by the Gateway contract.
type GatewayWithdrawalLocked struct {
	ReceiptHash [32]byte
	Receipt     TransferReceipt
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalLocked is a free log retrieval operation binding the contract event 0x89e52969465b1f1866fc5d46fd62de953962e9cb33552443cd999eba05bd20dc.
//
// Solidity: event WithdrawalLocked(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_Gateway *GatewayFilterer) FilterWithdrawalLocked(opts *bind.FilterOpts) (*GatewayWithdrawalLockedIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "WithdrawalLocked")
	if err != nil {
		return nil, err
	}
	return &GatewayWithdrawalLockedIterator{contract: _Gateway.contract, event: "WithdrawalLocked", logs: logs, sub: sub}, nil
}

// WatchWithdrawalLocked is a free log subscription operation binding the contract event 0x89e52969465b1f1866fc5d46fd62de953962e9cb33552443cd999eba05bd20dc.
//
// Solidity: event WithdrawalLocked(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_Gateway *GatewayFilterer) WatchWithdrawalLocked(opts *bind.WatchOpts, sink chan<- *GatewayWithdrawalLocked) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "WithdrawalLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWithdrawalLocked)
				if err := _Gateway.contract.UnpackLog(event, "WithdrawalLocked", log); err != nil {
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

// ParseWithdrawalLocked is a log parse operation binding the contract event 0x89e52969465b1f1866fc5d46fd62de953962e9cb33552443cd999eba05bd20dc.
//
// Solidity: event WithdrawalLocked(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_Gateway *GatewayFilterer) ParseWithdrawalLocked(log types.Log) (*GatewayWithdrawalLocked, error) {
	event := new(GatewayWithdrawalLocked)
	if err := _Gateway.contract.UnpackLog(event, "WithdrawalLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayWithdrawalUnlockedIterator is returned from FilterWithdrawalUnlocked and is used to iterate over the raw logs and unpacked data for WithdrawalUnlocked events raised by the Gateway contract.
type GatewayWithdrawalUnlockedIterator struct {
	Event *GatewayWithdrawalUnlocked // Event containing the contract specifics and raw log

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
func (it *GatewayWithdrawalUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWithdrawalUnlocked)
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
		it.Event = new(GatewayWithdrawalUnlocked)
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
func (it *GatewayWithdrawalUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWithdrawalUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWithdrawalUnlocked represents a WithdrawalUnlocked event raised by the Gateway contract.
type GatewayWithdrawalUnlocked struct {
	ReceiptHash [32]byte
	Receipt     TransferReceipt
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalUnlocked is a free log retrieval operation binding the contract event 0xd639511b37b3b002cca6cfe6bca0d833945a5af5a045578a0627fc43b79b2630.
//
// Solidity: event WithdrawalUnlocked(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_Gateway *GatewayFilterer) FilterWithdrawalUnlocked(opts *bind.FilterOpts) (*GatewayWithdrawalUnlockedIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "WithdrawalUnlocked")
	if err != nil {
		return nil, err
	}
	return &GatewayWithdrawalUnlockedIterator{contract: _Gateway.contract, event: "WithdrawalUnlocked", logs: logs, sub: sub}, nil
}

// WatchWithdrawalUnlocked is a free log subscription operation binding the contract event 0xd639511b37b3b002cca6cfe6bca0d833945a5af5a045578a0627fc43b79b2630.
//
// Solidity: event WithdrawalUnlocked(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_Gateway *GatewayFilterer) WatchWithdrawalUnlocked(opts *bind.WatchOpts, sink chan<- *GatewayWithdrawalUnlocked) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "WithdrawalUnlocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWithdrawalUnlocked)
				if err := _Gateway.contract.UnpackLog(event, "WithdrawalUnlocked", log); err != nil {
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

// ParseWithdrawalUnlocked is a log parse operation binding the contract event 0xd639511b37b3b002cca6cfe6bca0d833945a5af5a045578a0627fc43b79b2630.
//
// Solidity: event WithdrawalUnlocked(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_Gateway *GatewayFilterer) ParseWithdrawalUnlocked(log types.Log) (*GatewayWithdrawalUnlocked, error) {
	event := new(GatewayWithdrawalUnlocked)
	if err := _Gateway.contract.UnpackLog(event, "WithdrawalUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayWithdrewIterator is returned from FilterWithdrew and is used to iterate over the raw logs and unpacked data for Withdrew events raised by the Gateway contract.
type GatewayWithdrewIterator struct {
	Event *GatewayWithdrew // Event containing the contract specifics and raw log

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
func (it *GatewayWithdrewIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWithdrew)
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
		it.Event = new(GatewayWithdrew)
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
func (it *GatewayWithdrewIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWithdrewIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWithdrew represents a Withdrew event raised by the Gateway contract.
type GatewayWithdrew struct {
	ReceiptHash [32]byte
	Receipt     TransferReceipt
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrew is a free log retrieval operation binding the contract event 0x21e88e956aa3e086f6388e899965cef814688f99ad8bb29b08d396571016372d.
//
// Solidity: event Withdrew(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_Gateway *GatewayFilterer) FilterWithdrew(opts *bind.FilterOpts) (*GatewayWithdrewIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "Withdrew")
	if err != nil {
		return nil, err
	}
	return &GatewayWithdrewIterator{contract: _Gateway.contract, event: "Withdrew", logs: logs, sub: sub}, nil
}

// WatchWithdrew is a free log subscription operation binding the contract event 0x21e88e956aa3e086f6388e899965cef814688f99ad8bb29b08d396571016372d.
//
// Solidity: event Withdrew(bytes32 receiptHash, (uint256,uint8,(address,address,uint256),(address,address,uint256),(uint8,uint256,uint256)) receipt)
func (_Gateway *GatewayFilterer) WatchWithdrew(opts *bind.WatchOpts, sink chan<- *GatewayWithdrew) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "Withdrew")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWithdrew)
				if err := _Gateway.contract.UnpackLog(event, "Withdrew", log); err != nil {
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
func (_Gateway *GatewayFilterer) ParseWithdrew(log types.Log) (*GatewayWithdrew, error) {
	event := new(GatewayWithdrew)
	if err := _Gateway.contract.UnpackLog(event, "Withdrew", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayWrappedNativeTokenContractUpdatedIterator is returned from FilterWrappedNativeTokenContractUpdated and is used to iterate over the raw logs and unpacked data for WrappedNativeTokenContractUpdated events raised by the Gateway contract.
type GatewayWrappedNativeTokenContractUpdatedIterator struct {
	Event *GatewayWrappedNativeTokenContractUpdated // Event containing the contract specifics and raw log

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
func (it *GatewayWrappedNativeTokenContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayWrappedNativeTokenContractUpdated)
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
		it.Event = new(GatewayWrappedNativeTokenContractUpdated)
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
func (it *GatewayWrappedNativeTokenContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayWrappedNativeTokenContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayWrappedNativeTokenContractUpdated represents a WrappedNativeTokenContractUpdated event raised by the Gateway contract.
type GatewayWrappedNativeTokenContractUpdated struct {
	Weth common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWrappedNativeTokenContractUpdated is a free log retrieval operation binding the contract event 0x9d2334c23be647e994f27a72c5eee42a43d5bdcfe15bb88e939103c2b114cbaf.
//
// Solidity: event WrappedNativeTokenContractUpdated(address weth)
func (_Gateway *GatewayFilterer) FilterWrappedNativeTokenContractUpdated(opts *bind.FilterOpts) (*GatewayWrappedNativeTokenContractUpdatedIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "WrappedNativeTokenContractUpdated")
	if err != nil {
		return nil, err
	}
	return &GatewayWrappedNativeTokenContractUpdatedIterator{contract: _Gateway.contract, event: "WrappedNativeTokenContractUpdated", logs: logs, sub: sub}, nil
}

// WatchWrappedNativeTokenContractUpdated is a free log subscription operation binding the contract event 0x9d2334c23be647e994f27a72c5eee42a43d5bdcfe15bb88e939103c2b114cbaf.
//
// Solidity: event WrappedNativeTokenContractUpdated(address weth)
func (_Gateway *GatewayFilterer) WatchWrappedNativeTokenContractUpdated(opts *bind.WatchOpts, sink chan<- *GatewayWrappedNativeTokenContractUpdated) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "WrappedNativeTokenContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayWrappedNativeTokenContractUpdated)
				if err := _Gateway.contract.UnpackLog(event, "WrappedNativeTokenContractUpdated", log); err != nil {
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
func (_Gateway *GatewayFilterer) ParseWrappedNativeTokenContractUpdated(log types.Log) (*GatewayWrappedNativeTokenContractUpdated, error) {
	event := new(GatewayWrappedNativeTokenContractUpdated)
	if err := _Gateway.contract.UnpackLog(event, "WrappedNativeTokenContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
