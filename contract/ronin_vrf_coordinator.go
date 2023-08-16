// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// SLARandomRequest is an auto generated low-level Go binding around an user-defined struct.
type SLARandomRequest struct {
	BlockNumber      *big.Int
	CallbackGasLimit *big.Int
	GasPrice         *big.Int
	GasFee           *big.Int
	Requester        common.Address
	Consumer         common.Address
	RefundAddr       common.Address
	Nonce            *big.Int
	ConstantFee      *big.Int
}

// VRFProof is an auto generated low-level Go binding around an user-defined struct.
type VRFProof struct {
	Pk            [2]*big.Int
	Gamma         [2]*big.Int
	C             *big.Int
	S             *big.Int
	Seed          *big.Int
	UWitness      common.Address
	CGammaWitness [2]*big.Int
	SHashWitness  [2]*big.Int
	ZInv          *big.Int
}

// RoninVRFCoordinatorMetaData contains all meta data concerning the RoninVRFCoordinator contract.
var RoninVRFCoordinatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerIsNotAConsumer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrUnauthorizedConsumer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGasPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProofSeed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"InvalidRefundAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RONTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RandomRequestAlreadyFinalized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequestHashNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongRandomRequest\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"AddressWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_constantFee\",\"type\":\"uint256\"}],\"name\":\"ConstantFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_gasCost\",\"type\":\"uint256\"}],\"name\":\"GasAfterPaymentCalculationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_gasCost\",\"type\":\"uint256\"}],\"name\":\"GasToEstimateRandomFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"keyHashList\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"oracleAddrList\",\"type\":\"address[]\"}],\"name\":\"OraclesAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"keyHashList\",\"type\":\"bytes32[]\"}],\"name\":\"OraclesRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"randomSeed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paymentAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"constantFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"callbackResult\",\"type\":\"bool\"}],\"name\":\"RandomSeedFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"reqHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callbackGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"constantFee\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSLA.RandomRequest\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"RandomSeedRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"}],\"name\":\"TreasuryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"WhitelistAllChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_keyHashList\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_oracleAddrList\",\"type\":\"address[]\"}],\"name\":\"addOracles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"constantFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_callbackGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasPrice\",\"type\":\"uint256\"}],\"name\":\"estimateRequestRandomFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"pk\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"gamma\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uWitness\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"cGammaWitness\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"sHashWitness\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"zInv\",\"type\":\"uint256\"}],\"internalType\":\"structVRF.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callbackGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"constantFee\",\"type\":\"uint256\"}],\"internalType\":\"structSLA.RandomRequest\",\"name\":\"_req\",\"type\":\"tuple\"}],\"name\":\"fulfillRandomSeed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_paymentAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasAfterPaymentCalculation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasToEstimateRandomFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consumer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"getRandomRequestHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consumer\",\"type\":\"address\"}],\"name\":\"getRandomRequestNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gasToEstimateRandomFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasAfterPaymentCalculation\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gasToEstimateRandomFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasAfterPaymentCalculation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_constantFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"}],\"name\":\"initializeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"publicKey\",\"type\":\"uint256[2]\"}],\"name\":\"keyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyHash\",\"type\":\"bytes32\"}],\"name\":\"oracleAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_keyHashList\",\"type\":\"bytes32[]\"}],\"name\":\"removeOracles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_reqHash\",\"type\":\"bytes32\"}],\"name\":\"requestFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_callbackGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_consumer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"requestRandomSeed\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_reqHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_constantFee\",\"type\":\"uint256\"}],\"name\":\"setConstantFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gas\",\"type\":\"uint256\"}],\"name\":\"setGasAfterPaymentCalculation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gas\",\"type\":\"uint256\"}],\"name\":\"setGasToEstimateRandomFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"whitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"whitelistAllAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistedAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RoninVRFCoordinatorABI is the input ABI used to generate the binding from.
// Deprecated: Use RoninVRFCoordinatorMetaData.ABI instead.
var RoninVRFCoordinatorABI = RoninVRFCoordinatorMetaData.ABI

// RoninVRFCoordinator is an auto generated Go binding around an Ethereum contract.
type RoninVRFCoordinator struct {
	RoninVRFCoordinatorCaller     // Read-only binding to the contract
	RoninVRFCoordinatorTransactor // Write-only binding to the contract
	RoninVRFCoordinatorFilterer   // Log filterer for contract events
}

// RoninVRFCoordinatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type RoninVRFCoordinatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoninVRFCoordinatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RoninVRFCoordinatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoninVRFCoordinatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoninVRFCoordinatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoninVRFCoordinatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoninVRFCoordinatorSession struct {
	Contract     *RoninVRFCoordinator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RoninVRFCoordinatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoninVRFCoordinatorCallerSession struct {
	Contract *RoninVRFCoordinatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// RoninVRFCoordinatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoninVRFCoordinatorTransactorSession struct {
	Contract     *RoninVRFCoordinatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// RoninVRFCoordinatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type RoninVRFCoordinatorRaw struct {
	Contract *RoninVRFCoordinator // Generic contract binding to access the raw methods on
}

// RoninVRFCoordinatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoninVRFCoordinatorCallerRaw struct {
	Contract *RoninVRFCoordinatorCaller // Generic read-only contract binding to access the raw methods on
}

// RoninVRFCoordinatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoninVRFCoordinatorTransactorRaw struct {
	Contract *RoninVRFCoordinatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoninVRFCoordinator creates a new instance of RoninVRFCoordinator, bound to a specific deployed contract.
func NewRoninVRFCoordinator(address common.Address, backend bind.ContractBackend) (*RoninVRFCoordinator, error) {
	contract, err := bindRoninVRFCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RoninVRFCoordinator{RoninVRFCoordinatorCaller: RoninVRFCoordinatorCaller{contract: contract}, RoninVRFCoordinatorTransactor: RoninVRFCoordinatorTransactor{contract: contract}, RoninVRFCoordinatorFilterer: RoninVRFCoordinatorFilterer{contract: contract}}, nil
}

// NewRoninVRFCoordinatorCaller creates a new read-only instance of RoninVRFCoordinator, bound to a specific deployed contract.
func NewRoninVRFCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*RoninVRFCoordinatorCaller, error) {
	contract, err := bindRoninVRFCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoninVRFCoordinatorCaller{contract: contract}, nil
}

// NewRoninVRFCoordinatorTransactor creates a new write-only instance of RoninVRFCoordinator, bound to a specific deployed contract.
func NewRoninVRFCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*RoninVRFCoordinatorTransactor, error) {
	contract, err := bindRoninVRFCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoninVRFCoordinatorTransactor{contract: contract}, nil
}

// NewRoninVRFCoordinatorFilterer creates a new log filterer instance of RoninVRFCoordinator, bound to a specific deployed contract.
func NewRoninVRFCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*RoninVRFCoordinatorFilterer, error) {
	contract, err := bindRoninVRFCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoninVRFCoordinatorFilterer{contract: contract}, nil
}

// bindRoninVRFCoordinator binds a generic wrapper to an already deployed contract.
func bindRoninVRFCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RoninVRFCoordinatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoninVRFCoordinator *RoninVRFCoordinatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoninVRFCoordinator.Contract.RoninVRFCoordinatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoninVRFCoordinator *RoninVRFCoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.RoninVRFCoordinatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoninVRFCoordinator *RoninVRFCoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.RoninVRFCoordinatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoninVRFCoordinator *RoninVRFCoordinatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoninVRFCoordinator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoninVRFCoordinator *RoninVRFCoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoninVRFCoordinator *RoninVRFCoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.contract.Transact(opts, method, params...)
}

// ConstantFee is a free data retrieval call binding the contract method 0x75cc48a8.
//
// Solidity: function constantFee() view returns(uint256)
func (_RoninVRFCoordinator *RoninVRFCoordinatorCaller) ConstantFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninVRFCoordinator.contract.Call(opts, &out, "constantFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConstantFee is a free data retrieval call binding the contract method 0x75cc48a8.
//
// Solidity: function constantFee() view returns(uint256)
func (_RoninVRFCoordinator *RoninVRFCoordinatorSession) ConstantFee() (*big.Int, error) {
	return _RoninVRFCoordinator.Contract.ConstantFee(&_RoninVRFCoordinator.CallOpts)
}

// ConstantFee is a free data retrieval call binding the contract method 0x75cc48a8.
//
// Solidity: function constantFee() view returns(uint256)
func (_RoninVRFCoordinator *RoninVRFCoordinatorCallerSession) ConstantFee() (*big.Int, error) {
	return _RoninVRFCoordinator.Contract.ConstantFee(&_RoninVRFCoordinator.CallOpts)
}

// EstimateRequestRandomFee is a free data retrieval call binding the contract method 0x5109c096.
//
// Solidity: function estimateRequestRandomFee(uint256 _callbackGasLimit, uint256 _gasPrice) view returns(uint256)
func (_RoninVRFCoordinator *RoninVRFCoordinatorCaller) EstimateRequestRandomFee(opts *bind.CallOpts, _callbackGasLimit *big.Int, _gasPrice *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RoninVRFCoordinator.contract.Call(opts, &out, "estimateRequestRandomFee", _callbackGasLimit, _gasPrice)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateRequestRandomFee is a free data retrieval call binding the contract method 0x5109c096.
//
// Solidity: function estimateRequestRandomFee(uint256 _callbackGasLimit, uint256 _gasPrice) view returns(uint256)
func (_RoninVRFCoordinator *RoninVRFCoordinatorSession) EstimateRequestRandomFee(_callbackGasLimit *big.Int, _gasPrice *big.Int) (*big.Int, error) {
	return _RoninVRFCoordinator.Contract.EstimateRequestRandomFee(&_RoninVRFCoordinator.CallOpts, _callbackGasLimit, _gasPrice)
}

// EstimateRequestRandomFee is a free data retrieval call binding the contract method 0x5109c096.
//
// Solidity: function estimateRequestRandomFee(uint256 _callbackGasLimit, uint256 _gasPrice) view returns(uint256)
func (_RoninVRFCoordinator *RoninVRFCoordinatorCallerSession) EstimateRequestRandomFee(_callbackGasLimit *big.Int, _gasPrice *big.Int) (*big.Int, error) {
	return _RoninVRFCoordinator.Contract.EstimateRequestRandomFee(&_RoninVRFCoordinator.CallOpts, _callbackGasLimit, _gasPrice)
}

// GasAfterPaymentCalculation is a free data retrieval call binding the contract method 0x1199b2cd.
//
// Solidity: function gasAfterPaymentCalculation() view returns(uint256)
func (_RoninVRFCoordinator *RoninVRFCoordinatorCaller) GasAfterPaymentCalculation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninVRFCoordinator.contract.Call(opts, &out, "gasAfterPaymentCalculation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasAfterPaymentCalculation is a free data retrieval call binding the contract method 0x1199b2cd.
//
// Solidity: function gasAfterPaymentCalculation() view returns(uint256)
func (_RoninVRFCoordinator *RoninVRFCoordinatorSession) GasAfterPaymentCalculation() (*big.Int, error) {
	return _RoninVRFCoordinator.Contract.GasAfterPaymentCalculation(&_RoninVRFCoordinator.CallOpts)
}

// GasAfterPaymentCalculation is a free data retrieval call binding the contract method 0x1199b2cd.
//
// Solidity: function gasAfterPaymentCalculation() view returns(uint256)
func (_RoninVRFCoordinator *RoninVRFCoordinatorCallerSession) GasAfterPaymentCalculation() (*big.Int, error) {
	return _RoninVRFCoordinator.Contract.GasAfterPaymentCalculation(&_RoninVRFCoordinator.CallOpts)
}

// GasToEstimateRandomFee is a free data retrieval call binding the contract method 0xe7d1125b.
//
// Solidity: function gasToEstimateRandomFee() view returns(uint256)
func (_RoninVRFCoordinator *RoninVRFCoordinatorCaller) GasToEstimateRandomFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninVRFCoordinator.contract.Call(opts, &out, "gasToEstimateRandomFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasToEstimateRandomFee is a free data retrieval call binding the contract method 0xe7d1125b.
//
// Solidity: function gasToEstimateRandomFee() view returns(uint256)
func (_RoninVRFCoordinator *RoninVRFCoordinatorSession) GasToEstimateRandomFee() (*big.Int, error) {
	return _RoninVRFCoordinator.Contract.GasToEstimateRandomFee(&_RoninVRFCoordinator.CallOpts)
}

// GasToEstimateRandomFee is a free data retrieval call binding the contract method 0xe7d1125b.
//
// Solidity: function gasToEstimateRandomFee() view returns(uint256)
func (_RoninVRFCoordinator *RoninVRFCoordinatorCallerSession) GasToEstimateRandomFee() (*big.Int, error) {
	return _RoninVRFCoordinator.Contract.GasToEstimateRandomFee(&_RoninVRFCoordinator.CallOpts)
}

// GetRandomRequestHash is a free data retrieval call binding the contract method 0x3991c1fe.
//
// Solidity: function getRandomRequestHash(address _consumer, uint256 _nonce) view returns(bytes32 _hash)
func (_RoninVRFCoordinator *RoninVRFCoordinatorCaller) GetRandomRequestHash(opts *bind.CallOpts, _consumer common.Address, _nonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _RoninVRFCoordinator.contract.Call(opts, &out, "getRandomRequestHash", _consumer, _nonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRandomRequestHash is a free data retrieval call binding the contract method 0x3991c1fe.
//
// Solidity: function getRandomRequestHash(address _consumer, uint256 _nonce) view returns(bytes32 _hash)
func (_RoninVRFCoordinator *RoninVRFCoordinatorSession) GetRandomRequestHash(_consumer common.Address, _nonce *big.Int) ([32]byte, error) {
	return _RoninVRFCoordinator.Contract.GetRandomRequestHash(&_RoninVRFCoordinator.CallOpts, _consumer, _nonce)
}

// GetRandomRequestHash is a free data retrieval call binding the contract method 0x3991c1fe.
//
// Solidity: function getRandomRequestHash(address _consumer, uint256 _nonce) view returns(bytes32 _hash)
func (_RoninVRFCoordinator *RoninVRFCoordinatorCallerSession) GetRandomRequestHash(_consumer common.Address, _nonce *big.Int) ([32]byte, error) {
	return _RoninVRFCoordinator.Contract.GetRandomRequestHash(&_RoninVRFCoordinator.CallOpts, _consumer, _nonce)
}

// GetRandomRequestNonce is a free data retrieval call binding the contract method 0x4d2a22e2.
//
// Solidity: function getRandomRequestNonce(address _consumer) view returns(uint256 nonce)
func (_RoninVRFCoordinator *RoninVRFCoordinatorCaller) GetRandomRequestNonce(opts *bind.CallOpts, _consumer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RoninVRFCoordinator.contract.Call(opts, &out, "getRandomRequestNonce", _consumer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRandomRequestNonce is a free data retrieval call binding the contract method 0x4d2a22e2.
//
// Solidity: function getRandomRequestNonce(address _consumer) view returns(uint256 nonce)
func (_RoninVRFCoordinator *RoninVRFCoordinatorSession) GetRandomRequestNonce(_consumer common.Address) (*big.Int, error) {
	return _RoninVRFCoordinator.Contract.GetRandomRequestNonce(&_RoninVRFCoordinator.CallOpts, _consumer)
}

// GetRandomRequestNonce is a free data retrieval call binding the contract method 0x4d2a22e2.
//
// Solidity: function getRandomRequestNonce(address _consumer) view returns(uint256 nonce)
func (_RoninVRFCoordinator *RoninVRFCoordinatorCallerSession) GetRandomRequestNonce(_consumer common.Address) (*big.Int, error) {
	return _RoninVRFCoordinator.Contract.GetRandomRequestNonce(&_RoninVRFCoordinator.CallOpts, _consumer)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address _address) view returns(bool)
func (_RoninVRFCoordinator *RoninVRFCoordinatorCaller) IsWhitelisted(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _RoninVRFCoordinator.contract.Call(opts, &out, "isWhitelisted", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address _address) view returns(bool)
func (_RoninVRFCoordinator *RoninVRFCoordinatorSession) IsWhitelisted(_address common.Address) (bool, error) {
	return _RoninVRFCoordinator.Contract.IsWhitelisted(&_RoninVRFCoordinator.CallOpts, _address)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address _address) view returns(bool)
func (_RoninVRFCoordinator *RoninVRFCoordinatorCallerSession) IsWhitelisted(_address common.Address) (bool, error) {
	return _RoninVRFCoordinator.Contract.IsWhitelisted(&_RoninVRFCoordinator.CallOpts, _address)
}

// KeyHash is a free data retrieval call binding the contract method 0xffb9cee4.
//
// Solidity: function keyHash(uint256[2] publicKey) pure returns(bytes32)
func (_RoninVRFCoordinator *RoninVRFCoordinatorCaller) KeyHash(opts *bind.CallOpts, publicKey [2]*big.Int) ([32]byte, error) {
	var out []interface{}
	err := _RoninVRFCoordinator.contract.Call(opts, &out, "keyHash", publicKey)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// KeyHash is a free data retrieval call binding the contract method 0xffb9cee4.
//
// Solidity: function keyHash(uint256[2] publicKey) pure returns(bytes32)
func (_RoninVRFCoordinator *RoninVRFCoordinatorSession) KeyHash(publicKey [2]*big.Int) ([32]byte, error) {
	return _RoninVRFCoordinator.Contract.KeyHash(&_RoninVRFCoordinator.CallOpts, publicKey)
}

// KeyHash is a free data retrieval call binding the contract method 0xffb9cee4.
//
// Solidity: function keyHash(uint256[2] publicKey) pure returns(bytes32)
func (_RoninVRFCoordinator *RoninVRFCoordinatorCallerSession) KeyHash(publicKey [2]*big.Int) ([32]byte, error) {
	return _RoninVRFCoordinator.Contract.KeyHash(&_RoninVRFCoordinator.CallOpts, publicKey)
}

// OracleAddress is a free data retrieval call binding the contract method 0x00d4e44e.
//
// Solidity: function oracleAddress(bytes32 _keyHash) view returns(address)
func (_RoninVRFCoordinator *RoninVRFCoordinatorCaller) OracleAddress(opts *bind.CallOpts, _keyHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _RoninVRFCoordinator.contract.Call(opts, &out, "oracleAddress", _keyHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OracleAddress is a free data retrieval call binding the contract method 0x00d4e44e.
//
// Solidity: function oracleAddress(bytes32 _keyHash) view returns(address)
func (_RoninVRFCoordinator *RoninVRFCoordinatorSession) OracleAddress(_keyHash [32]byte) (common.Address, error) {
	return _RoninVRFCoordinator.Contract.OracleAddress(&_RoninVRFCoordinator.CallOpts, _keyHash)
}

// OracleAddress is a free data retrieval call binding the contract method 0x00d4e44e.
//
// Solidity: function oracleAddress(bytes32 _keyHash) view returns(address)
func (_RoninVRFCoordinator *RoninVRFCoordinatorCallerSession) OracleAddress(_keyHash [32]byte) (common.Address, error) {
	return _RoninVRFCoordinator.Contract.OracleAddress(&_RoninVRFCoordinator.CallOpts, _keyHash)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RoninVRFCoordinator *RoninVRFCoordinatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RoninVRFCoordinator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RoninVRFCoordinator *RoninVRFCoordinatorSession) Owner() (common.Address, error) {
	return _RoninVRFCoordinator.Contract.Owner(&_RoninVRFCoordinator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RoninVRFCoordinator *RoninVRFCoordinatorCallerSession) Owner() (common.Address, error) {
	return _RoninVRFCoordinator.Contract.Owner(&_RoninVRFCoordinator.CallOpts)
}

// RequestFinalized is a free data retrieval call binding the contract method 0x37a21e9d.
//
// Solidity: function requestFinalized(bytes32 _reqHash) view returns(bool)
func (_RoninVRFCoordinator *RoninVRFCoordinatorCaller) RequestFinalized(opts *bind.CallOpts, _reqHash [32]byte) (bool, error) {
	var out []interface{}
	err := _RoninVRFCoordinator.contract.Call(opts, &out, "requestFinalized", _reqHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequestFinalized is a free data retrieval call binding the contract method 0x37a21e9d.
//
// Solidity: function requestFinalized(bytes32 _reqHash) view returns(bool)
func (_RoninVRFCoordinator *RoninVRFCoordinatorSession) RequestFinalized(_reqHash [32]byte) (bool, error) {
	return _RoninVRFCoordinator.Contract.RequestFinalized(&_RoninVRFCoordinator.CallOpts, _reqHash)
}

// RequestFinalized is a free data retrieval call binding the contract method 0x37a21e9d.
//
// Solidity: function requestFinalized(bytes32 _reqHash) view returns(bool)
func (_RoninVRFCoordinator *RoninVRFCoordinatorCallerSession) RequestFinalized(_reqHash [32]byte) (bool, error) {
	return _RoninVRFCoordinator.Contract.RequestFinalized(&_RoninVRFCoordinator.CallOpts, _reqHash)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_RoninVRFCoordinator *RoninVRFCoordinatorCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RoninVRFCoordinator.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_RoninVRFCoordinator *RoninVRFCoordinatorSession) Treasury() (common.Address, error) {
	return _RoninVRFCoordinator.Contract.Treasury(&_RoninVRFCoordinator.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_RoninVRFCoordinator *RoninVRFCoordinatorCallerSession) Treasury() (common.Address, error) {
	return _RoninVRFCoordinator.Contract.Treasury(&_RoninVRFCoordinator.CallOpts)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_RoninVRFCoordinator *RoninVRFCoordinatorCaller) Whitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _RoninVRFCoordinator.contract.Call(opts, &out, "whitelisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_RoninVRFCoordinator *RoninVRFCoordinatorSession) Whitelisted(arg0 common.Address) (bool, error) {
	return _RoninVRFCoordinator.Contract.Whitelisted(&_RoninVRFCoordinator.CallOpts, arg0)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_RoninVRFCoordinator *RoninVRFCoordinatorCallerSession) Whitelisted(arg0 common.Address) (bool, error) {
	return _RoninVRFCoordinator.Contract.Whitelisted(&_RoninVRFCoordinator.CallOpts, arg0)
}

// WhitelistedAll is a free data retrieval call binding the contract method 0x141f5331.
//
// Solidity: function whitelistedAll() view returns(bool)
func (_RoninVRFCoordinator *RoninVRFCoordinatorCaller) WhitelistedAll(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RoninVRFCoordinator.contract.Call(opts, &out, "whitelistedAll")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistedAll is a free data retrieval call binding the contract method 0x141f5331.
//
// Solidity: function whitelistedAll() view returns(bool)
func (_RoninVRFCoordinator *RoninVRFCoordinatorSession) WhitelistedAll() (bool, error) {
	return _RoninVRFCoordinator.Contract.WhitelistedAll(&_RoninVRFCoordinator.CallOpts)
}

// WhitelistedAll is a free data retrieval call binding the contract method 0x141f5331.
//
// Solidity: function whitelistedAll() view returns(bool)
func (_RoninVRFCoordinator *RoninVRFCoordinatorCallerSession) WhitelistedAll() (bool, error) {
	return _RoninVRFCoordinator.Contract.WhitelistedAll(&_RoninVRFCoordinator.CallOpts)
}

// AddOracles is a paid mutator transaction binding the contract method 0xf0571b9c.
//
// Solidity: function addOracles(bytes32[] _keyHashList, address[] _oracleAddrList) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorTransactor) AddOracles(opts *bind.TransactOpts, _keyHashList [][32]byte, _oracleAddrList []common.Address) (*types.Transaction, error) {
	return _RoninVRFCoordinator.contract.Transact(opts, "addOracles", _keyHashList, _oracleAddrList)
}

// AddOracles is a paid mutator transaction binding the contract method 0xf0571b9c.
//
// Solidity: function addOracles(bytes32[] _keyHashList, address[] _oracleAddrList) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorSession) AddOracles(_keyHashList [][32]byte, _oracleAddrList []common.Address) (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.AddOracles(&_RoninVRFCoordinator.TransactOpts, _keyHashList, _oracleAddrList)
}

// AddOracles is a paid mutator transaction binding the contract method 0xf0571b9c.
//
// Solidity: function addOracles(bytes32[] _keyHashList, address[] _oracleAddrList) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorTransactorSession) AddOracles(_keyHashList [][32]byte, _oracleAddrList []common.Address) (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.AddOracles(&_RoninVRFCoordinator.TransactOpts, _keyHashList, _oracleAddrList)
}

// FulfillRandomSeed is a paid mutator transaction binding the contract method 0x6ae53af4.
//
// Solidity: function fulfillRandomSeed((uint256[2],uint256[2],uint256,uint256,uint256,address,uint256[2],uint256[2],uint256) _proof, (uint256,uint256,uint256,uint256,address,address,address,uint256,uint256) _req) returns(uint256 _paymentAmount)
func (_RoninVRFCoordinator *RoninVRFCoordinatorTransactor) FulfillRandomSeed(opts *bind.TransactOpts, _proof VRFProof, _req SLARandomRequest) (*types.Transaction, error) {
	return _RoninVRFCoordinator.contract.Transact(opts, "fulfillRandomSeed", _proof, _req)
}

// FulfillRandomSeed is a paid mutator transaction binding the contract method 0x6ae53af4.
//
// Solidity: function fulfillRandomSeed((uint256[2],uint256[2],uint256,uint256,uint256,address,uint256[2],uint256[2],uint256) _proof, (uint256,uint256,uint256,uint256,address,address,address,uint256,uint256) _req) returns(uint256 _paymentAmount)
func (_RoninVRFCoordinator *RoninVRFCoordinatorSession) FulfillRandomSeed(_proof VRFProof, _req SLARandomRequest) (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.FulfillRandomSeed(&_RoninVRFCoordinator.TransactOpts, _proof, _req)
}

// FulfillRandomSeed is a paid mutator transaction binding the contract method 0x6ae53af4.
//
// Solidity: function fulfillRandomSeed((uint256[2],uint256[2],uint256,uint256,uint256,address,uint256[2],uint256[2],uint256) _proof, (uint256,uint256,uint256,uint256,address,address,address,uint256,uint256) _req) returns(uint256 _paymentAmount)
func (_RoninVRFCoordinator *RoninVRFCoordinatorTransactorSession) FulfillRandomSeed(_proof VRFProof, _req SLARandomRequest) (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.FulfillRandomSeed(&_RoninVRFCoordinator.TransactOpts, _proof, _req)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address _admin, uint256 _gasToEstimateRandomFee, uint256 _gasAfterPaymentCalculation) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, _gasToEstimateRandomFee *big.Int, _gasAfterPaymentCalculation *big.Int) (*types.Transaction, error) {
	return _RoninVRFCoordinator.contract.Transact(opts, "initialize", _admin, _gasToEstimateRandomFee, _gasAfterPaymentCalculation)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address _admin, uint256 _gasToEstimateRandomFee, uint256 _gasAfterPaymentCalculation) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorSession) Initialize(_admin common.Address, _gasToEstimateRandomFee *big.Int, _gasAfterPaymentCalculation *big.Int) (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.Initialize(&_RoninVRFCoordinator.TransactOpts, _admin, _gasToEstimateRandomFee, _gasAfterPaymentCalculation)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address _admin, uint256 _gasToEstimateRandomFee, uint256 _gasAfterPaymentCalculation) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorTransactorSession) Initialize(_admin common.Address, _gasToEstimateRandomFee *big.Int, _gasAfterPaymentCalculation *big.Int) (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.Initialize(&_RoninVRFCoordinator.TransactOpts, _admin, _gasToEstimateRandomFee, _gasAfterPaymentCalculation)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x04a49f6f.
//
// Solidity: function initializeV2(uint256 _gasToEstimateRandomFee, uint256 _gasAfterPaymentCalculation, uint256 _constantFee, address _treasury) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorTransactor) InitializeV2(opts *bind.TransactOpts, _gasToEstimateRandomFee *big.Int, _gasAfterPaymentCalculation *big.Int, _constantFee *big.Int, _treasury common.Address) (*types.Transaction, error) {
	return _RoninVRFCoordinator.contract.Transact(opts, "initializeV2", _gasToEstimateRandomFee, _gasAfterPaymentCalculation, _constantFee, _treasury)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x04a49f6f.
//
// Solidity: function initializeV2(uint256 _gasToEstimateRandomFee, uint256 _gasAfterPaymentCalculation, uint256 _constantFee, address _treasury) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorSession) InitializeV2(_gasToEstimateRandomFee *big.Int, _gasAfterPaymentCalculation *big.Int, _constantFee *big.Int, _treasury common.Address) (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.InitializeV2(&_RoninVRFCoordinator.TransactOpts, _gasToEstimateRandomFee, _gasAfterPaymentCalculation, _constantFee, _treasury)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x04a49f6f.
//
// Solidity: function initializeV2(uint256 _gasToEstimateRandomFee, uint256 _gasAfterPaymentCalculation, uint256 _constantFee, address _treasury) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorTransactorSession) InitializeV2(_gasToEstimateRandomFee *big.Int, _gasAfterPaymentCalculation *big.Int, _constantFee *big.Int, _treasury common.Address) (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.InitializeV2(&_RoninVRFCoordinator.TransactOpts, _gasToEstimateRandomFee, _gasAfterPaymentCalculation, _constantFee, _treasury)
}

// RemoveOracles is a paid mutator transaction binding the contract method 0x8bb7772e.
//
// Solidity: function removeOracles(bytes32[] _keyHashList) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorTransactor) RemoveOracles(opts *bind.TransactOpts, _keyHashList [][32]byte) (*types.Transaction, error) {
	return _RoninVRFCoordinator.contract.Transact(opts, "removeOracles", _keyHashList)
}

// RemoveOracles is a paid mutator transaction binding the contract method 0x8bb7772e.
//
// Solidity: function removeOracles(bytes32[] _keyHashList) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorSession) RemoveOracles(_keyHashList [][32]byte) (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.RemoveOracles(&_RoninVRFCoordinator.TransactOpts, _keyHashList)
}

// RemoveOracles is a paid mutator transaction binding the contract method 0x8bb7772e.
//
// Solidity: function removeOracles(bytes32[] _keyHashList) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorTransactorSession) RemoveOracles(_keyHashList [][32]byte) (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.RemoveOracles(&_RoninVRFCoordinator.TransactOpts, _keyHashList)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoninVRFCoordinator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.RenounceOwnership(&_RoninVRFCoordinator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.RenounceOwnership(&_RoninVRFCoordinator.TransactOpts)
}

// RequestRandomSeed is a paid mutator transaction binding the contract method 0x9725d768.
//
// Solidity: function requestRandomSeed(uint256 _callbackGasLimit, uint256 _gasPrice, address _consumer, address _refundAddress) payable returns(bytes32 _reqHash)
func (_RoninVRFCoordinator *RoninVRFCoordinatorTransactor) RequestRandomSeed(opts *bind.TransactOpts, _callbackGasLimit *big.Int, _gasPrice *big.Int, _consumer common.Address, _refundAddress common.Address) (*types.Transaction, error) {
	return _RoninVRFCoordinator.contract.Transact(opts, "requestRandomSeed", _callbackGasLimit, _gasPrice, _consumer, _refundAddress)
}

// RequestRandomSeed is a paid mutator transaction binding the contract method 0x9725d768.
//
// Solidity: function requestRandomSeed(uint256 _callbackGasLimit, uint256 _gasPrice, address _consumer, address _refundAddress) payable returns(bytes32 _reqHash)
func (_RoninVRFCoordinator *RoninVRFCoordinatorSession) RequestRandomSeed(_callbackGasLimit *big.Int, _gasPrice *big.Int, _consumer common.Address, _refundAddress common.Address) (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.RequestRandomSeed(&_RoninVRFCoordinator.TransactOpts, _callbackGasLimit, _gasPrice, _consumer, _refundAddress)
}

// RequestRandomSeed is a paid mutator transaction binding the contract method 0x9725d768.
//
// Solidity: function requestRandomSeed(uint256 _callbackGasLimit, uint256 _gasPrice, address _consumer, address _refundAddress) payable returns(bytes32 _reqHash)
func (_RoninVRFCoordinator *RoninVRFCoordinatorTransactorSession) RequestRandomSeed(_callbackGasLimit *big.Int, _gasPrice *big.Int, _consumer common.Address, _refundAddress common.Address) (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.RequestRandomSeed(&_RoninVRFCoordinator.TransactOpts, _callbackGasLimit, _gasPrice, _consumer, _refundAddress)
}

// SetConstantFee is a paid mutator transaction binding the contract method 0x67e09bef.
//
// Solidity: function setConstantFee(uint256 _constantFee) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorTransactor) SetConstantFee(opts *bind.TransactOpts, _constantFee *big.Int) (*types.Transaction, error) {
	return _RoninVRFCoordinator.contract.Transact(opts, "setConstantFee", _constantFee)
}

// SetConstantFee is a paid mutator transaction binding the contract method 0x67e09bef.
//
// Solidity: function setConstantFee(uint256 _constantFee) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorSession) SetConstantFee(_constantFee *big.Int) (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.SetConstantFee(&_RoninVRFCoordinator.TransactOpts, _constantFee)
}

// SetConstantFee is a paid mutator transaction binding the contract method 0x67e09bef.
//
// Solidity: function setConstantFee(uint256 _constantFee) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorTransactorSession) SetConstantFee(_constantFee *big.Int) (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.SetConstantFee(&_RoninVRFCoordinator.TransactOpts, _constantFee)
}

// SetGasAfterPaymentCalculation is a paid mutator transaction binding the contract method 0xe52f2496.
//
// Solidity: function setGasAfterPaymentCalculation(uint256 _gas) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorTransactor) SetGasAfterPaymentCalculation(opts *bind.TransactOpts, _gas *big.Int) (*types.Transaction, error) {
	return _RoninVRFCoordinator.contract.Transact(opts, "setGasAfterPaymentCalculation", _gas)
}

// SetGasAfterPaymentCalculation is a paid mutator transaction binding the contract method 0xe52f2496.
//
// Solidity: function setGasAfterPaymentCalculation(uint256 _gas) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorSession) SetGasAfterPaymentCalculation(_gas *big.Int) (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.SetGasAfterPaymentCalculation(&_RoninVRFCoordinator.TransactOpts, _gas)
}

// SetGasAfterPaymentCalculation is a paid mutator transaction binding the contract method 0xe52f2496.
//
// Solidity: function setGasAfterPaymentCalculation(uint256 _gas) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorTransactorSession) SetGasAfterPaymentCalculation(_gas *big.Int) (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.SetGasAfterPaymentCalculation(&_RoninVRFCoordinator.TransactOpts, _gas)
}

// SetGasToEstimateRandomFee is a paid mutator transaction binding the contract method 0xf3612fe9.
//
// Solidity: function setGasToEstimateRandomFee(uint256 _gas) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorTransactor) SetGasToEstimateRandomFee(opts *bind.TransactOpts, _gas *big.Int) (*types.Transaction, error) {
	return _RoninVRFCoordinator.contract.Transact(opts, "setGasToEstimateRandomFee", _gas)
}

// SetGasToEstimateRandomFee is a paid mutator transaction binding the contract method 0xf3612fe9.
//
// Solidity: function setGasToEstimateRandomFee(uint256 _gas) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorSession) SetGasToEstimateRandomFee(_gas *big.Int) (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.SetGasToEstimateRandomFee(&_RoninVRFCoordinator.TransactOpts, _gas)
}

// SetGasToEstimateRandomFee is a paid mutator transaction binding the contract method 0xf3612fe9.
//
// Solidity: function setGasToEstimateRandomFee(uint256 _gas) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorTransactorSession) SetGasToEstimateRandomFee(_gas *big.Int) (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.SetGasToEstimateRandomFee(&_RoninVRFCoordinator.TransactOpts, _gas)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorTransactor) SetTreasury(opts *bind.TransactOpts, _treasury common.Address) (*types.Transaction, error) {
	return _RoninVRFCoordinator.contract.Transact(opts, "setTreasury", _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorSession) SetTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.SetTreasury(&_RoninVRFCoordinator.TransactOpts, _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorTransactorSession) SetTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.SetTreasury(&_RoninVRFCoordinator.TransactOpts, _treasury)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RoninVRFCoordinator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.TransferOwnership(&_RoninVRFCoordinator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.TransferOwnership(&_RoninVRFCoordinator.TransactOpts, newOwner)
}

// Whitelist is a paid mutator transaction binding the contract method 0xf59c3708.
//
// Solidity: function whitelist(address _address, bool _status) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorTransactor) Whitelist(opts *bind.TransactOpts, _address common.Address, _status bool) (*types.Transaction, error) {
	return _RoninVRFCoordinator.contract.Transact(opts, "whitelist", _address, _status)
}

// Whitelist is a paid mutator transaction binding the contract method 0xf59c3708.
//
// Solidity: function whitelist(address _address, bool _status) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorSession) Whitelist(_address common.Address, _status bool) (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.Whitelist(&_RoninVRFCoordinator.TransactOpts, _address, _status)
}

// Whitelist is a paid mutator transaction binding the contract method 0xf59c3708.
//
// Solidity: function whitelist(address _address, bool _status) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorTransactorSession) Whitelist(_address common.Address, _status bool) (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.Whitelist(&_RoninVRFCoordinator.TransactOpts, _address, _status)
}

// WhitelistAllAddresses is a paid mutator transaction binding the contract method 0x47a3ed0c.
//
// Solidity: function whitelistAllAddresses(bool _status) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorTransactor) WhitelistAllAddresses(opts *bind.TransactOpts, _status bool) (*types.Transaction, error) {
	return _RoninVRFCoordinator.contract.Transact(opts, "whitelistAllAddresses", _status)
}

// WhitelistAllAddresses is a paid mutator transaction binding the contract method 0x47a3ed0c.
//
// Solidity: function whitelistAllAddresses(bool _status) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorSession) WhitelistAllAddresses(_status bool) (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.WhitelistAllAddresses(&_RoninVRFCoordinator.TransactOpts, _status)
}

// WhitelistAllAddresses is a paid mutator transaction binding the contract method 0x47a3ed0c.
//
// Solidity: function whitelistAllAddresses(bool _status) returns()
func (_RoninVRFCoordinator *RoninVRFCoordinatorTransactorSession) WhitelistAllAddresses(_status bool) (*types.Transaction, error) {
	return _RoninVRFCoordinator.Contract.WhitelistAllAddresses(&_RoninVRFCoordinator.TransactOpts, _status)
}

// RoninVRFCoordinatorAddressWhitelistedIterator is returned from FilterAddressWhitelisted and is used to iterate over the raw logs and unpacked data for AddressWhitelisted events raised by the RoninVRFCoordinator contract.
type RoninVRFCoordinatorAddressWhitelistedIterator struct {
	Event *RoninVRFCoordinatorAddressWhitelisted // Event containing the contract specifics and raw log

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
func (it *RoninVRFCoordinatorAddressWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninVRFCoordinatorAddressWhitelisted)
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
		it.Event = new(RoninVRFCoordinatorAddressWhitelisted)
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
func (it *RoninVRFCoordinatorAddressWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninVRFCoordinatorAddressWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninVRFCoordinatorAddressWhitelisted represents a AddressWhitelisted event raised by the RoninVRFCoordinator contract.
type RoninVRFCoordinatorAddressWhitelisted struct {
	Address common.Address
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddressWhitelisted is a free log retrieval operation binding the contract event 0xaf367c7d20ce5b2ab6da56afd0c9c39b00ba995263c60292a3e1ee3781fd4885.
//
// Solidity: event AddressWhitelisted(address indexed _address, bool indexed _status)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) FilterAddressWhitelisted(opts *bind.FilterOpts, _address []common.Address, _status []bool) (*RoninVRFCoordinatorAddressWhitelistedIterator, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}
	var _statusRule []interface{}
	for _, _statusItem := range _status {
		_statusRule = append(_statusRule, _statusItem)
	}

	logs, sub, err := _RoninVRFCoordinator.contract.FilterLogs(opts, "AddressWhitelisted", _addressRule, _statusRule)
	if err != nil {
		return nil, err
	}
	return &RoninVRFCoordinatorAddressWhitelistedIterator{contract: _RoninVRFCoordinator.contract, event: "AddressWhitelisted", logs: logs, sub: sub}, nil
}

// WatchAddressWhitelisted is a free log subscription operation binding the contract event 0xaf367c7d20ce5b2ab6da56afd0c9c39b00ba995263c60292a3e1ee3781fd4885.
//
// Solidity: event AddressWhitelisted(address indexed _address, bool indexed _status)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) WatchAddressWhitelisted(opts *bind.WatchOpts, sink chan<- *RoninVRFCoordinatorAddressWhitelisted, _address []common.Address, _status []bool) (event.Subscription, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}
	var _statusRule []interface{}
	for _, _statusItem := range _status {
		_statusRule = append(_statusRule, _statusItem)
	}

	logs, sub, err := _RoninVRFCoordinator.contract.WatchLogs(opts, "AddressWhitelisted", _addressRule, _statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninVRFCoordinatorAddressWhitelisted)
				if err := _RoninVRFCoordinator.contract.UnpackLog(event, "AddressWhitelisted", log); err != nil {
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

// ParseAddressWhitelisted is a log parse operation binding the contract event 0xaf367c7d20ce5b2ab6da56afd0c9c39b00ba995263c60292a3e1ee3781fd4885.
//
// Solidity: event AddressWhitelisted(address indexed _address, bool indexed _status)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) ParseAddressWhitelisted(log types.Log) (*RoninVRFCoordinatorAddressWhitelisted, error) {
	event := new(RoninVRFCoordinatorAddressWhitelisted)
	if err := _RoninVRFCoordinator.contract.UnpackLog(event, "AddressWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninVRFCoordinatorConstantFeeUpdatedIterator is returned from FilterConstantFeeUpdated and is used to iterate over the raw logs and unpacked data for ConstantFeeUpdated events raised by the RoninVRFCoordinator contract.
type RoninVRFCoordinatorConstantFeeUpdatedIterator struct {
	Event *RoninVRFCoordinatorConstantFeeUpdated // Event containing the contract specifics and raw log

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
func (it *RoninVRFCoordinatorConstantFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninVRFCoordinatorConstantFeeUpdated)
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
		it.Event = new(RoninVRFCoordinatorConstantFeeUpdated)
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
func (it *RoninVRFCoordinatorConstantFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninVRFCoordinatorConstantFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninVRFCoordinatorConstantFeeUpdated represents a ConstantFeeUpdated event raised by the RoninVRFCoordinator contract.
type RoninVRFCoordinatorConstantFeeUpdated struct {
	ConstantFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConstantFeeUpdated is a free log retrieval operation binding the contract event 0x7a5f2d652e8508c29f4c2adfe783c653f1787dc27ef4bdefa758ca2c45ce439e.
//
// Solidity: event ConstantFeeUpdated(uint256 _constantFee)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) FilterConstantFeeUpdated(opts *bind.FilterOpts) (*RoninVRFCoordinatorConstantFeeUpdatedIterator, error) {

	logs, sub, err := _RoninVRFCoordinator.contract.FilterLogs(opts, "ConstantFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &RoninVRFCoordinatorConstantFeeUpdatedIterator{contract: _RoninVRFCoordinator.contract, event: "ConstantFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchConstantFeeUpdated is a free log subscription operation binding the contract event 0x7a5f2d652e8508c29f4c2adfe783c653f1787dc27ef4bdefa758ca2c45ce439e.
//
// Solidity: event ConstantFeeUpdated(uint256 _constantFee)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) WatchConstantFeeUpdated(opts *bind.WatchOpts, sink chan<- *RoninVRFCoordinatorConstantFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _RoninVRFCoordinator.contract.WatchLogs(opts, "ConstantFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninVRFCoordinatorConstantFeeUpdated)
				if err := _RoninVRFCoordinator.contract.UnpackLog(event, "ConstantFeeUpdated", log); err != nil {
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

// ParseConstantFeeUpdated is a log parse operation binding the contract event 0x7a5f2d652e8508c29f4c2adfe783c653f1787dc27ef4bdefa758ca2c45ce439e.
//
// Solidity: event ConstantFeeUpdated(uint256 _constantFee)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) ParseConstantFeeUpdated(log types.Log) (*RoninVRFCoordinatorConstantFeeUpdated, error) {
	event := new(RoninVRFCoordinatorConstantFeeUpdated)
	if err := _RoninVRFCoordinator.contract.UnpackLog(event, "ConstantFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninVRFCoordinatorGasAfterPaymentCalculationUpdatedIterator is returned from FilterGasAfterPaymentCalculationUpdated and is used to iterate over the raw logs and unpacked data for GasAfterPaymentCalculationUpdated events raised by the RoninVRFCoordinator contract.
type RoninVRFCoordinatorGasAfterPaymentCalculationUpdatedIterator struct {
	Event *RoninVRFCoordinatorGasAfterPaymentCalculationUpdated // Event containing the contract specifics and raw log

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
func (it *RoninVRFCoordinatorGasAfterPaymentCalculationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninVRFCoordinatorGasAfterPaymentCalculationUpdated)
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
		it.Event = new(RoninVRFCoordinatorGasAfterPaymentCalculationUpdated)
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
func (it *RoninVRFCoordinatorGasAfterPaymentCalculationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninVRFCoordinatorGasAfterPaymentCalculationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninVRFCoordinatorGasAfterPaymentCalculationUpdated represents a GasAfterPaymentCalculationUpdated event raised by the RoninVRFCoordinator contract.
type RoninVRFCoordinatorGasAfterPaymentCalculationUpdated struct {
	GasCost *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGasAfterPaymentCalculationUpdated is a free log retrieval operation binding the contract event 0x410fd5aad09c7ebf0e275cad89eb8faa9e43280bbf0320483401858e0820a43b.
//
// Solidity: event GasAfterPaymentCalculationUpdated(uint256 _gasCost)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) FilterGasAfterPaymentCalculationUpdated(opts *bind.FilterOpts) (*RoninVRFCoordinatorGasAfterPaymentCalculationUpdatedIterator, error) {

	logs, sub, err := _RoninVRFCoordinator.contract.FilterLogs(opts, "GasAfterPaymentCalculationUpdated")
	if err != nil {
		return nil, err
	}
	return &RoninVRFCoordinatorGasAfterPaymentCalculationUpdatedIterator{contract: _RoninVRFCoordinator.contract, event: "GasAfterPaymentCalculationUpdated", logs: logs, sub: sub}, nil
}

// WatchGasAfterPaymentCalculationUpdated is a free log subscription operation binding the contract event 0x410fd5aad09c7ebf0e275cad89eb8faa9e43280bbf0320483401858e0820a43b.
//
// Solidity: event GasAfterPaymentCalculationUpdated(uint256 _gasCost)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) WatchGasAfterPaymentCalculationUpdated(opts *bind.WatchOpts, sink chan<- *RoninVRFCoordinatorGasAfterPaymentCalculationUpdated) (event.Subscription, error) {

	logs, sub, err := _RoninVRFCoordinator.contract.WatchLogs(opts, "GasAfterPaymentCalculationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninVRFCoordinatorGasAfterPaymentCalculationUpdated)
				if err := _RoninVRFCoordinator.contract.UnpackLog(event, "GasAfterPaymentCalculationUpdated", log); err != nil {
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

// ParseGasAfterPaymentCalculationUpdated is a log parse operation binding the contract event 0x410fd5aad09c7ebf0e275cad89eb8faa9e43280bbf0320483401858e0820a43b.
//
// Solidity: event GasAfterPaymentCalculationUpdated(uint256 _gasCost)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) ParseGasAfterPaymentCalculationUpdated(log types.Log) (*RoninVRFCoordinatorGasAfterPaymentCalculationUpdated, error) {
	event := new(RoninVRFCoordinatorGasAfterPaymentCalculationUpdated)
	if err := _RoninVRFCoordinator.contract.UnpackLog(event, "GasAfterPaymentCalculationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninVRFCoordinatorGasToEstimateRandomFeeUpdatedIterator is returned from FilterGasToEstimateRandomFeeUpdated and is used to iterate over the raw logs and unpacked data for GasToEstimateRandomFeeUpdated events raised by the RoninVRFCoordinator contract.
type RoninVRFCoordinatorGasToEstimateRandomFeeUpdatedIterator struct {
	Event *RoninVRFCoordinatorGasToEstimateRandomFeeUpdated // Event containing the contract specifics and raw log

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
func (it *RoninVRFCoordinatorGasToEstimateRandomFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninVRFCoordinatorGasToEstimateRandomFeeUpdated)
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
		it.Event = new(RoninVRFCoordinatorGasToEstimateRandomFeeUpdated)
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
func (it *RoninVRFCoordinatorGasToEstimateRandomFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninVRFCoordinatorGasToEstimateRandomFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninVRFCoordinatorGasToEstimateRandomFeeUpdated represents a GasToEstimateRandomFeeUpdated event raised by the RoninVRFCoordinator contract.
type RoninVRFCoordinatorGasToEstimateRandomFeeUpdated struct {
	GasCost *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGasToEstimateRandomFeeUpdated is a free log retrieval operation binding the contract event 0x7127b6643bb3a0e9c5a3ecf52d548002150de8b2bdaf6af0b41382a549c42139.
//
// Solidity: event GasToEstimateRandomFeeUpdated(uint256 _gasCost)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) FilterGasToEstimateRandomFeeUpdated(opts *bind.FilterOpts) (*RoninVRFCoordinatorGasToEstimateRandomFeeUpdatedIterator, error) {

	logs, sub, err := _RoninVRFCoordinator.contract.FilterLogs(opts, "GasToEstimateRandomFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &RoninVRFCoordinatorGasToEstimateRandomFeeUpdatedIterator{contract: _RoninVRFCoordinator.contract, event: "GasToEstimateRandomFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchGasToEstimateRandomFeeUpdated is a free log subscription operation binding the contract event 0x7127b6643bb3a0e9c5a3ecf52d548002150de8b2bdaf6af0b41382a549c42139.
//
// Solidity: event GasToEstimateRandomFeeUpdated(uint256 _gasCost)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) WatchGasToEstimateRandomFeeUpdated(opts *bind.WatchOpts, sink chan<- *RoninVRFCoordinatorGasToEstimateRandomFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _RoninVRFCoordinator.contract.WatchLogs(opts, "GasToEstimateRandomFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninVRFCoordinatorGasToEstimateRandomFeeUpdated)
				if err := _RoninVRFCoordinator.contract.UnpackLog(event, "GasToEstimateRandomFeeUpdated", log); err != nil {
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

// ParseGasToEstimateRandomFeeUpdated is a log parse operation binding the contract event 0x7127b6643bb3a0e9c5a3ecf52d548002150de8b2bdaf6af0b41382a549c42139.
//
// Solidity: event GasToEstimateRandomFeeUpdated(uint256 _gasCost)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) ParseGasToEstimateRandomFeeUpdated(log types.Log) (*RoninVRFCoordinatorGasToEstimateRandomFeeUpdated, error) {
	event := new(RoninVRFCoordinatorGasToEstimateRandomFeeUpdated)
	if err := _RoninVRFCoordinator.contract.UnpackLog(event, "GasToEstimateRandomFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninVRFCoordinatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RoninVRFCoordinator contract.
type RoninVRFCoordinatorInitializedIterator struct {
	Event *RoninVRFCoordinatorInitialized // Event containing the contract specifics and raw log

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
func (it *RoninVRFCoordinatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninVRFCoordinatorInitialized)
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
		it.Event = new(RoninVRFCoordinatorInitialized)
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
func (it *RoninVRFCoordinatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninVRFCoordinatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninVRFCoordinatorInitialized represents a Initialized event raised by the RoninVRFCoordinator contract.
type RoninVRFCoordinatorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*RoninVRFCoordinatorInitializedIterator, error) {

	logs, sub, err := _RoninVRFCoordinator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RoninVRFCoordinatorInitializedIterator{contract: _RoninVRFCoordinator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RoninVRFCoordinatorInitialized) (event.Subscription, error) {

	logs, sub, err := _RoninVRFCoordinator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninVRFCoordinatorInitialized)
				if err := _RoninVRFCoordinator.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) ParseInitialized(log types.Log) (*RoninVRFCoordinatorInitialized, error) {
	event := new(RoninVRFCoordinatorInitialized)
	if err := _RoninVRFCoordinator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninVRFCoordinatorOraclesAddedIterator is returned from FilterOraclesAdded and is used to iterate over the raw logs and unpacked data for OraclesAdded events raised by the RoninVRFCoordinator contract.
type RoninVRFCoordinatorOraclesAddedIterator struct {
	Event *RoninVRFCoordinatorOraclesAdded // Event containing the contract specifics and raw log

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
func (it *RoninVRFCoordinatorOraclesAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninVRFCoordinatorOraclesAdded)
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
		it.Event = new(RoninVRFCoordinatorOraclesAdded)
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
func (it *RoninVRFCoordinatorOraclesAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninVRFCoordinatorOraclesAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninVRFCoordinatorOraclesAdded represents a OraclesAdded event raised by the RoninVRFCoordinator contract.
type RoninVRFCoordinatorOraclesAdded struct {
	KeyHashList    [][32]byte
	OracleAddrList []common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOraclesAdded is a free log retrieval operation binding the contract event 0x49f41c0f90f9a743fdb3b4a76e6058551c171be7175499e9cae77122d9676f38.
//
// Solidity: event OraclesAdded(bytes32[] keyHashList, address[] oracleAddrList)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) FilterOraclesAdded(opts *bind.FilterOpts) (*RoninVRFCoordinatorOraclesAddedIterator, error) {

	logs, sub, err := _RoninVRFCoordinator.contract.FilterLogs(opts, "OraclesAdded")
	if err != nil {
		return nil, err
	}
	return &RoninVRFCoordinatorOraclesAddedIterator{contract: _RoninVRFCoordinator.contract, event: "OraclesAdded", logs: logs, sub: sub}, nil
}

// WatchOraclesAdded is a free log subscription operation binding the contract event 0x49f41c0f90f9a743fdb3b4a76e6058551c171be7175499e9cae77122d9676f38.
//
// Solidity: event OraclesAdded(bytes32[] keyHashList, address[] oracleAddrList)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) WatchOraclesAdded(opts *bind.WatchOpts, sink chan<- *RoninVRFCoordinatorOraclesAdded) (event.Subscription, error) {

	logs, sub, err := _RoninVRFCoordinator.contract.WatchLogs(opts, "OraclesAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninVRFCoordinatorOraclesAdded)
				if err := _RoninVRFCoordinator.contract.UnpackLog(event, "OraclesAdded", log); err != nil {
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

// ParseOraclesAdded is a log parse operation binding the contract event 0x49f41c0f90f9a743fdb3b4a76e6058551c171be7175499e9cae77122d9676f38.
//
// Solidity: event OraclesAdded(bytes32[] keyHashList, address[] oracleAddrList)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) ParseOraclesAdded(log types.Log) (*RoninVRFCoordinatorOraclesAdded, error) {
	event := new(RoninVRFCoordinatorOraclesAdded)
	if err := _RoninVRFCoordinator.contract.UnpackLog(event, "OraclesAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninVRFCoordinatorOraclesRemovedIterator is returned from FilterOraclesRemoved and is used to iterate over the raw logs and unpacked data for OraclesRemoved events raised by the RoninVRFCoordinator contract.
type RoninVRFCoordinatorOraclesRemovedIterator struct {
	Event *RoninVRFCoordinatorOraclesRemoved // Event containing the contract specifics and raw log

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
func (it *RoninVRFCoordinatorOraclesRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninVRFCoordinatorOraclesRemoved)
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
		it.Event = new(RoninVRFCoordinatorOraclesRemoved)
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
func (it *RoninVRFCoordinatorOraclesRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninVRFCoordinatorOraclesRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninVRFCoordinatorOraclesRemoved represents a OraclesRemoved event raised by the RoninVRFCoordinator contract.
type RoninVRFCoordinatorOraclesRemoved struct {
	KeyHashList [][32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOraclesRemoved is a free log retrieval operation binding the contract event 0x8d8481139eeb4930797ed1169695e10526672c8de87c1f24c3a94664b4007afb.
//
// Solidity: event OraclesRemoved(bytes32[] keyHashList)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) FilterOraclesRemoved(opts *bind.FilterOpts) (*RoninVRFCoordinatorOraclesRemovedIterator, error) {

	logs, sub, err := _RoninVRFCoordinator.contract.FilterLogs(opts, "OraclesRemoved")
	if err != nil {
		return nil, err
	}
	return &RoninVRFCoordinatorOraclesRemovedIterator{contract: _RoninVRFCoordinator.contract, event: "OraclesRemoved", logs: logs, sub: sub}, nil
}

// WatchOraclesRemoved is a free log subscription operation binding the contract event 0x8d8481139eeb4930797ed1169695e10526672c8de87c1f24c3a94664b4007afb.
//
// Solidity: event OraclesRemoved(bytes32[] keyHashList)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) WatchOraclesRemoved(opts *bind.WatchOpts, sink chan<- *RoninVRFCoordinatorOraclesRemoved) (event.Subscription, error) {

	logs, sub, err := _RoninVRFCoordinator.contract.WatchLogs(opts, "OraclesRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninVRFCoordinatorOraclesRemoved)
				if err := _RoninVRFCoordinator.contract.UnpackLog(event, "OraclesRemoved", log); err != nil {
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

// ParseOraclesRemoved is a log parse operation binding the contract event 0x8d8481139eeb4930797ed1169695e10526672c8de87c1f24c3a94664b4007afb.
//
// Solidity: event OraclesRemoved(bytes32[] keyHashList)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) ParseOraclesRemoved(log types.Log) (*RoninVRFCoordinatorOraclesRemoved, error) {
	event := new(RoninVRFCoordinatorOraclesRemoved)
	if err := _RoninVRFCoordinator.contract.UnpackLog(event, "OraclesRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninVRFCoordinatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RoninVRFCoordinator contract.
type RoninVRFCoordinatorOwnershipTransferredIterator struct {
	Event *RoninVRFCoordinatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RoninVRFCoordinatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninVRFCoordinatorOwnershipTransferred)
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
		it.Event = new(RoninVRFCoordinatorOwnershipTransferred)
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
func (it *RoninVRFCoordinatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninVRFCoordinatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninVRFCoordinatorOwnershipTransferred represents a OwnershipTransferred event raised by the RoninVRFCoordinator contract.
type RoninVRFCoordinatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RoninVRFCoordinatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RoninVRFCoordinator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RoninVRFCoordinatorOwnershipTransferredIterator{contract: _RoninVRFCoordinator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RoninVRFCoordinatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RoninVRFCoordinator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninVRFCoordinatorOwnershipTransferred)
				if err := _RoninVRFCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) ParseOwnershipTransferred(log types.Log) (*RoninVRFCoordinatorOwnershipTransferred, error) {
	event := new(RoninVRFCoordinatorOwnershipTransferred)
	if err := _RoninVRFCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninVRFCoordinatorRandomSeedFulfilledIterator is returned from FilterRandomSeedFulfilled and is used to iterate over the raw logs and unpacked data for RandomSeedFulfilled events raised by the RoninVRFCoordinator contract.
type RoninVRFCoordinatorRandomSeedFulfilledIterator struct {
	Event *RoninVRFCoordinatorRandomSeedFulfilled // Event containing the contract specifics and raw log

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
func (it *RoninVRFCoordinatorRandomSeedFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninVRFCoordinatorRandomSeedFulfilled)
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
		it.Event = new(RoninVRFCoordinatorRandomSeedFulfilled)
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
func (it *RoninVRFCoordinatorRandomSeedFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninVRFCoordinatorRandomSeedFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninVRFCoordinatorRandomSeedFulfilled represents a RandomSeedFulfilled event raised by the RoninVRFCoordinator contract.
type RoninVRFCoordinatorRandomSeedFulfilled struct {
	RequestHash    [32]byte
	RandomSeed     *big.Int
	RequestValue   *big.Int
	RefundAmount   *big.Int
	PaymentAmount  *big.Int
	ConstantFee    *big.Int
	CallbackResult bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRandomSeedFulfilled is a free log retrieval operation binding the contract event 0xad8205703644511eef3e682fe71baaf5a0d2a611157f3ba57cc89226e3729cd1.
//
// Solidity: event RandomSeedFulfilled(bytes32 indexed requestHash, uint256 randomSeed, uint256 requestValue, uint256 refundAmount, uint256 paymentAmount, uint256 constantFee, bool callbackResult)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) FilterRandomSeedFulfilled(opts *bind.FilterOpts, requestHash [][32]byte) (*RoninVRFCoordinatorRandomSeedFulfilledIterator, error) {

	var requestHashRule []interface{}
	for _, requestHashItem := range requestHash {
		requestHashRule = append(requestHashRule, requestHashItem)
	}

	logs, sub, err := _RoninVRFCoordinator.contract.FilterLogs(opts, "RandomSeedFulfilled", requestHashRule)
	if err != nil {
		return nil, err
	}
	return &RoninVRFCoordinatorRandomSeedFulfilledIterator{contract: _RoninVRFCoordinator.contract, event: "RandomSeedFulfilled", logs: logs, sub: sub}, nil
}

// WatchRandomSeedFulfilled is a free log subscription operation binding the contract event 0xad8205703644511eef3e682fe71baaf5a0d2a611157f3ba57cc89226e3729cd1.
//
// Solidity: event RandomSeedFulfilled(bytes32 indexed requestHash, uint256 randomSeed, uint256 requestValue, uint256 refundAmount, uint256 paymentAmount, uint256 constantFee, bool callbackResult)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) WatchRandomSeedFulfilled(opts *bind.WatchOpts, sink chan<- *RoninVRFCoordinatorRandomSeedFulfilled, requestHash [][32]byte) (event.Subscription, error) {

	var requestHashRule []interface{}
	for _, requestHashItem := range requestHash {
		requestHashRule = append(requestHashRule, requestHashItem)
	}

	logs, sub, err := _RoninVRFCoordinator.contract.WatchLogs(opts, "RandomSeedFulfilled", requestHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninVRFCoordinatorRandomSeedFulfilled)
				if err := _RoninVRFCoordinator.contract.UnpackLog(event, "RandomSeedFulfilled", log); err != nil {
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

// ParseRandomSeedFulfilled is a log parse operation binding the contract event 0xad8205703644511eef3e682fe71baaf5a0d2a611157f3ba57cc89226e3729cd1.
//
// Solidity: event RandomSeedFulfilled(bytes32 indexed requestHash, uint256 randomSeed, uint256 requestValue, uint256 refundAmount, uint256 paymentAmount, uint256 constantFee, bool callbackResult)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) ParseRandomSeedFulfilled(log types.Log) (*RoninVRFCoordinatorRandomSeedFulfilled, error) {
	event := new(RoninVRFCoordinatorRandomSeedFulfilled)
	if err := _RoninVRFCoordinator.contract.UnpackLog(event, "RandomSeedFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninVRFCoordinatorRandomSeedRequestedIterator is returned from FilterRandomSeedRequested and is used to iterate over the raw logs and unpacked data for RandomSeedRequested events raised by the RoninVRFCoordinator contract.
type RoninVRFCoordinatorRandomSeedRequestedIterator struct {
	Event *RoninVRFCoordinatorRandomSeedRequested // Event containing the contract specifics and raw log

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
func (it *RoninVRFCoordinatorRandomSeedRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninVRFCoordinatorRandomSeedRequested)
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
		it.Event = new(RoninVRFCoordinatorRandomSeedRequested)
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
func (it *RoninVRFCoordinatorRandomSeedRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninVRFCoordinatorRandomSeedRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninVRFCoordinatorRandomSeedRequested represents a RandomSeedRequested event raised by the RoninVRFCoordinator contract.
type RoninVRFCoordinatorRandomSeedRequested struct {
	ReqHash [32]byte
	Request SLARandomRequest
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRandomSeedRequested is a free log retrieval operation binding the contract event 0xc3d3bb585055665a04b8b1f0e5a4831575a84b2c21f533a8eb88286dbeebfa19.
//
// Solidity: event RandomSeedRequested(bytes32 indexed reqHash, (uint256,uint256,uint256,uint256,address,address,address,uint256,uint256) request)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) FilterRandomSeedRequested(opts *bind.FilterOpts, reqHash [][32]byte) (*RoninVRFCoordinatorRandomSeedRequestedIterator, error) {

	var reqHashRule []interface{}
	for _, reqHashItem := range reqHash {
		reqHashRule = append(reqHashRule, reqHashItem)
	}

	logs, sub, err := _RoninVRFCoordinator.contract.FilterLogs(opts, "RandomSeedRequested", reqHashRule)
	if err != nil {
		return nil, err
	}
	return &RoninVRFCoordinatorRandomSeedRequestedIterator{contract: _RoninVRFCoordinator.contract, event: "RandomSeedRequested", logs: logs, sub: sub}, nil
}

// WatchRandomSeedRequested is a free log subscription operation binding the contract event 0xc3d3bb585055665a04b8b1f0e5a4831575a84b2c21f533a8eb88286dbeebfa19.
//
// Solidity: event RandomSeedRequested(bytes32 indexed reqHash, (uint256,uint256,uint256,uint256,address,address,address,uint256,uint256) request)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) WatchRandomSeedRequested(opts *bind.WatchOpts, sink chan<- *RoninVRFCoordinatorRandomSeedRequested, reqHash [][32]byte) (event.Subscription, error) {

	var reqHashRule []interface{}
	for _, reqHashItem := range reqHash {
		reqHashRule = append(reqHashRule, reqHashItem)
	}

	logs, sub, err := _RoninVRFCoordinator.contract.WatchLogs(opts, "RandomSeedRequested", reqHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninVRFCoordinatorRandomSeedRequested)
				if err := _RoninVRFCoordinator.contract.UnpackLog(event, "RandomSeedRequested", log); err != nil {
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

// ParseRandomSeedRequested is a log parse operation binding the contract event 0xc3d3bb585055665a04b8b1f0e5a4831575a84b2c21f533a8eb88286dbeebfa19.
//
// Solidity: event RandomSeedRequested(bytes32 indexed reqHash, (uint256,uint256,uint256,uint256,address,address,address,uint256,uint256) request)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) ParseRandomSeedRequested(log types.Log) (*RoninVRFCoordinatorRandomSeedRequested, error) {
	event := new(RoninVRFCoordinatorRandomSeedRequested)
	if err := _RoninVRFCoordinator.contract.UnpackLog(event, "RandomSeedRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninVRFCoordinatorTreasuryUpdatedIterator is returned from FilterTreasuryUpdated and is used to iterate over the raw logs and unpacked data for TreasuryUpdated events raised by the RoninVRFCoordinator contract.
type RoninVRFCoordinatorTreasuryUpdatedIterator struct {
	Event *RoninVRFCoordinatorTreasuryUpdated // Event containing the contract specifics and raw log

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
func (it *RoninVRFCoordinatorTreasuryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninVRFCoordinatorTreasuryUpdated)
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
		it.Event = new(RoninVRFCoordinatorTreasuryUpdated)
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
func (it *RoninVRFCoordinatorTreasuryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninVRFCoordinatorTreasuryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninVRFCoordinatorTreasuryUpdated represents a TreasuryUpdated event raised by the RoninVRFCoordinator contract.
type RoninVRFCoordinatorTreasuryUpdated struct {
	Treasury common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTreasuryUpdated is a free log retrieval operation binding the contract event 0x7dae230f18360d76a040c81f050aa14eb9d6dc7901b20fc5d855e2a20fe814d1.
//
// Solidity: event TreasuryUpdated(address _treasury)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) FilterTreasuryUpdated(opts *bind.FilterOpts) (*RoninVRFCoordinatorTreasuryUpdatedIterator, error) {

	logs, sub, err := _RoninVRFCoordinator.contract.FilterLogs(opts, "TreasuryUpdated")
	if err != nil {
		return nil, err
	}
	return &RoninVRFCoordinatorTreasuryUpdatedIterator{contract: _RoninVRFCoordinator.contract, event: "TreasuryUpdated", logs: logs, sub: sub}, nil
}

// WatchTreasuryUpdated is a free log subscription operation binding the contract event 0x7dae230f18360d76a040c81f050aa14eb9d6dc7901b20fc5d855e2a20fe814d1.
//
// Solidity: event TreasuryUpdated(address _treasury)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) WatchTreasuryUpdated(opts *bind.WatchOpts, sink chan<- *RoninVRFCoordinatorTreasuryUpdated) (event.Subscription, error) {

	logs, sub, err := _RoninVRFCoordinator.contract.WatchLogs(opts, "TreasuryUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninVRFCoordinatorTreasuryUpdated)
				if err := _RoninVRFCoordinator.contract.UnpackLog(event, "TreasuryUpdated", log); err != nil {
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

// ParseTreasuryUpdated is a log parse operation binding the contract event 0x7dae230f18360d76a040c81f050aa14eb9d6dc7901b20fc5d855e2a20fe814d1.
//
// Solidity: event TreasuryUpdated(address _treasury)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) ParseTreasuryUpdated(log types.Log) (*RoninVRFCoordinatorTreasuryUpdated, error) {
	event := new(RoninVRFCoordinatorTreasuryUpdated)
	if err := _RoninVRFCoordinator.contract.UnpackLog(event, "TreasuryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninVRFCoordinatorWhitelistAllChangedIterator is returned from FilterWhitelistAllChanged and is used to iterate over the raw logs and unpacked data for WhitelistAllChanged events raised by the RoninVRFCoordinator contract.
type RoninVRFCoordinatorWhitelistAllChangedIterator struct {
	Event *RoninVRFCoordinatorWhitelistAllChanged // Event containing the contract specifics and raw log

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
func (it *RoninVRFCoordinatorWhitelistAllChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninVRFCoordinatorWhitelistAllChanged)
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
		it.Event = new(RoninVRFCoordinatorWhitelistAllChanged)
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
func (it *RoninVRFCoordinatorWhitelistAllChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninVRFCoordinatorWhitelistAllChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninVRFCoordinatorWhitelistAllChanged represents a WhitelistAllChanged event raised by the RoninVRFCoordinator contract.
type RoninVRFCoordinatorWhitelistAllChanged struct {
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAllChanged is a free log retrieval operation binding the contract event 0xf00cf1f0ddd6809f07fa6905b1b604c167671d7164140306bba41f906a637ff1.
//
// Solidity: event WhitelistAllChanged(bool indexed _status)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) FilterWhitelistAllChanged(opts *bind.FilterOpts, _status []bool) (*RoninVRFCoordinatorWhitelistAllChangedIterator, error) {

	var _statusRule []interface{}
	for _, _statusItem := range _status {
		_statusRule = append(_statusRule, _statusItem)
	}

	logs, sub, err := _RoninVRFCoordinator.contract.FilterLogs(opts, "WhitelistAllChanged", _statusRule)
	if err != nil {
		return nil, err
	}
	return &RoninVRFCoordinatorWhitelistAllChangedIterator{contract: _RoninVRFCoordinator.contract, event: "WhitelistAllChanged", logs: logs, sub: sub}, nil
}

// WatchWhitelistAllChanged is a free log subscription operation binding the contract event 0xf00cf1f0ddd6809f07fa6905b1b604c167671d7164140306bba41f906a637ff1.
//
// Solidity: event WhitelistAllChanged(bool indexed _status)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) WatchWhitelistAllChanged(opts *bind.WatchOpts, sink chan<- *RoninVRFCoordinatorWhitelistAllChanged, _status []bool) (event.Subscription, error) {

	var _statusRule []interface{}
	for _, _statusItem := range _status {
		_statusRule = append(_statusRule, _statusItem)
	}

	logs, sub, err := _RoninVRFCoordinator.contract.WatchLogs(opts, "WhitelistAllChanged", _statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninVRFCoordinatorWhitelistAllChanged)
				if err := _RoninVRFCoordinator.contract.UnpackLog(event, "WhitelistAllChanged", log); err != nil {
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

// ParseWhitelistAllChanged is a log parse operation binding the contract event 0xf00cf1f0ddd6809f07fa6905b1b604c167671d7164140306bba41f906a637ff1.
//
// Solidity: event WhitelistAllChanged(bool indexed _status)
func (_RoninVRFCoordinator *RoninVRFCoordinatorFilterer) ParseWhitelistAllChanged(log types.Log) (*RoninVRFCoordinatorWhitelistAllChanged, error) {
	event := new(RoninVRFCoordinatorWhitelistAllChanged)
	if err := _RoninVRFCoordinator.contract.UnpackLog(event, "WhitelistAllChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
