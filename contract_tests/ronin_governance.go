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

// RoninGovernanceSignature is an auto generated low-level Go binding around an user-defined struct.
type RoninGovernanceSignature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// RoninGovernanceMetaData contains all meta data concerning the RoninGovernance contract.
var RoninGovernanceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_voters\",\"type\":\"address[]\"}],\"name\":\"getBridgeOperatorVotingSignatures\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structRoninGovernance.Signature[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_operators\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structRoninGovernance.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"name\":\"voteBridgeOperatorsBySignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"332635be": "getBridgeOperatorVotingSignatures(uint256,address[])",
		"56e237e8": "voteBridgeOperatorsBySignatures(uint256,address[],(uint8,bytes32,bytes32)[])",
	},
	Bin: "0x608060405234801561001057600080fd5b5061057a806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063332635be1461003b57806356e237e814610064575b600080fd5b61004e6100493660046102ab565b610079565b60405161005b91906102f7565b60405180910390f35b6100776100723660046103c3565b6101a5565b005b606060008267ffffffffffffffff81111561009657610096610353565b6040519080825280602002602001820160405280156100e157816020015b60408051606081018252600080825260208083018290529282015282526000199092019101816100b45790505b50905060005b8381101561019c57600086815260208190526040812081878785818110610110576101106104d7565b905060200201602081019061012591906104ed565b6001600160a01b0316815260208082019290925260409081016000208151606081018352815460ff16815260018201549381019390935260020154908201528351909150819084908490811061017d5761017d6104d7565b60200260200101819052505080806101949061051d565b9150506100e7565b50949350505050565b60005b8281101561025857816000815181106101c3576101c36104d7565b602002602001015160008087815260200190815260200160002060008686858181106101f1576101f16104d7565b905060200201602081019061020691906104ed565b6001600160a01b0316815260208082019290925260409081016000208351815460ff191660ff9091161781559183015160018301559190910151600290910155806102508161051d565b9150506101a8565b5050505050565b60008083601f84011261027157600080fd5b50813567ffffffffffffffff81111561028957600080fd5b6020830191508360208260051b85010111156102a457600080fd5b9250929050565b6000806000604084860312156102c057600080fd5b83359250602084013567ffffffffffffffff8111156102de57600080fd5b6102ea8682870161025f565b9497909650939450505050565b602080825282518282018190526000919060409081850190868401855b82811015610346578151805160ff16855286810151878601528501518585015260609093019290850190600101610314565b5091979650505050505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561038c5761038c610353565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156103bb576103bb610353565b604052919050565b60008060008060608086880312156103da57600080fd5b8535945060208087013567ffffffffffffffff808211156103fa57600080fd5b6104068a838b0161025f565b909750955060409150888201358181111561042057600080fd5b8901601f81018b1361043157600080fd5b80358281111561044357610443610353565b610451858260051b01610392565b818152858101935090860282018501908c82111561046e57600080fd5b918501915b818310156104c65786838e03121561048b5760008081fd5b610493610369565b833560ff811681146104a55760008081fd5b81528387013587820152858401358682015284529285019291860191610473565b999c989b5096995050505050505050565b634e487b7160e01b600052603260045260246000fd5b6000602082840312156104ff57600080fd5b81356001600160a01b038116811461051657600080fd5b9392505050565b60006001820161053d57634e487b7160e01b600052601160045260246000fd5b506001019056fea2646970667358221220b955bb832bc91e2eb7b466597b548cb85c2373fee7471579b6a6b782c6348ff664736f6c63430008110033",
}

// RoninGovernanceABI is the input ABI used to generate the binding from.
// Deprecated: Use RoninGovernanceMetaData.ABI instead.
var RoninGovernanceABI = RoninGovernanceMetaData.ABI

// Deprecated: Use RoninGovernanceMetaData.Sigs instead.
// RoninGovernanceFuncSigs maps the 4-byte function signature to its string representation.
var RoninGovernanceFuncSigs = RoninGovernanceMetaData.Sigs

// RoninGovernanceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RoninGovernanceMetaData.Bin instead.
var RoninGovernanceBin = RoninGovernanceMetaData.Bin

// DeployRoninGovernance deploys a new Ethereum contract, binding an instance of RoninGovernance to it.
func DeployRoninGovernance(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RoninGovernance, error) {
	parsed, err := RoninGovernanceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RoninGovernanceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RoninGovernance{RoninGovernanceCaller: RoninGovernanceCaller{contract: contract}, RoninGovernanceTransactor: RoninGovernanceTransactor{contract: contract}, RoninGovernanceFilterer: RoninGovernanceFilterer{contract: contract}}, nil
}

// RoninGovernance is an auto generated Go binding around an Ethereum contract.
type RoninGovernance struct {
	RoninGovernanceCaller     // Read-only binding to the contract
	RoninGovernanceTransactor // Write-only binding to the contract
	RoninGovernanceFilterer   // Log filterer for contract events
}

// RoninGovernanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type RoninGovernanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoninGovernanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RoninGovernanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoninGovernanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoninGovernanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoninGovernanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoninGovernanceSession struct {
	Contract     *RoninGovernance  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoninGovernanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoninGovernanceCallerSession struct {
	Contract *RoninGovernanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// RoninGovernanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoninGovernanceTransactorSession struct {
	Contract     *RoninGovernanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// RoninGovernanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type RoninGovernanceRaw struct {
	Contract *RoninGovernance // Generic contract binding to access the raw methods on
}

// RoninGovernanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoninGovernanceCallerRaw struct {
	Contract *RoninGovernanceCaller // Generic read-only contract binding to access the raw methods on
}

// RoninGovernanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoninGovernanceTransactorRaw struct {
	Contract *RoninGovernanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoninGovernance creates a new instance of RoninGovernance, bound to a specific deployed contract.
func NewRoninGovernance(address common.Address, backend bind.ContractBackend) (*RoninGovernance, error) {
	contract, err := bindRoninGovernance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RoninGovernance{RoninGovernanceCaller: RoninGovernanceCaller{contract: contract}, RoninGovernanceTransactor: RoninGovernanceTransactor{contract: contract}, RoninGovernanceFilterer: RoninGovernanceFilterer{contract: contract}}, nil
}

// NewRoninGovernanceCaller creates a new read-only instance of RoninGovernance, bound to a specific deployed contract.
func NewRoninGovernanceCaller(address common.Address, caller bind.ContractCaller) (*RoninGovernanceCaller, error) {
	contract, err := bindRoninGovernance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoninGovernanceCaller{contract: contract}, nil
}

// NewRoninGovernanceTransactor creates a new write-only instance of RoninGovernance, bound to a specific deployed contract.
func NewRoninGovernanceTransactor(address common.Address, transactor bind.ContractTransactor) (*RoninGovernanceTransactor, error) {
	contract, err := bindRoninGovernance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoninGovernanceTransactor{contract: contract}, nil
}

// NewRoninGovernanceFilterer creates a new log filterer instance of RoninGovernance, bound to a specific deployed contract.
func NewRoninGovernanceFilterer(address common.Address, filterer bind.ContractFilterer) (*RoninGovernanceFilterer, error) {
	contract, err := bindRoninGovernance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoninGovernanceFilterer{contract: contract}, nil
}

// bindRoninGovernance binds a generic wrapper to an already deployed contract.
func bindRoninGovernance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RoninGovernanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoninGovernance *RoninGovernanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoninGovernance.Contract.RoninGovernanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoninGovernance *RoninGovernanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoninGovernance.Contract.RoninGovernanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoninGovernance *RoninGovernanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoninGovernance.Contract.RoninGovernanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoninGovernance *RoninGovernanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoninGovernance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoninGovernance *RoninGovernanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoninGovernance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoninGovernance *RoninGovernanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoninGovernance.Contract.contract.Transact(opts, method, params...)
}

// GetBridgeOperatorVotingSignatures is a free data retrieval call binding the contract method 0x332635be.
//
// Solidity: function getBridgeOperatorVotingSignatures(uint256 _period, address[] _voters) view returns((uint8,bytes32,bytes32)[])
func (_RoninGovernance *RoninGovernanceCaller) GetBridgeOperatorVotingSignatures(opts *bind.CallOpts, _period *big.Int, _voters []common.Address) ([]RoninGovernanceSignature, error) {
	var out []interface{}
	err := _RoninGovernance.contract.Call(opts, &out, "getBridgeOperatorVotingSignatures", _period, _voters)

	if err != nil {
		return *new([]RoninGovernanceSignature), err
	}

	out0 := *abi.ConvertType(out[0], new([]RoninGovernanceSignature)).(*[]RoninGovernanceSignature)

	return out0, err

}

// GetBridgeOperatorVotingSignatures is a free data retrieval call binding the contract method 0x332635be.
//
// Solidity: function getBridgeOperatorVotingSignatures(uint256 _period, address[] _voters) view returns((uint8,bytes32,bytes32)[])
func (_RoninGovernance *RoninGovernanceSession) GetBridgeOperatorVotingSignatures(_period *big.Int, _voters []common.Address) ([]RoninGovernanceSignature, error) {
	return _RoninGovernance.Contract.GetBridgeOperatorVotingSignatures(&_RoninGovernance.CallOpts, _period, _voters)
}

// GetBridgeOperatorVotingSignatures is a free data retrieval call binding the contract method 0x332635be.
//
// Solidity: function getBridgeOperatorVotingSignatures(uint256 _period, address[] _voters) view returns((uint8,bytes32,bytes32)[])
func (_RoninGovernance *RoninGovernanceCallerSession) GetBridgeOperatorVotingSignatures(_period *big.Int, _voters []common.Address) ([]RoninGovernanceSignature, error) {
	return _RoninGovernance.Contract.GetBridgeOperatorVotingSignatures(&_RoninGovernance.CallOpts, _period, _voters)
}

// VoteBridgeOperatorsBySignatures is a paid mutator transaction binding the contract method 0x56e237e8.
//
// Solidity: function voteBridgeOperatorsBySignatures(uint256 _period, address[] _operators, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_RoninGovernance *RoninGovernanceTransactor) VoteBridgeOperatorsBySignatures(opts *bind.TransactOpts, _period *big.Int, _operators []common.Address, _signatures []RoninGovernanceSignature) (*types.Transaction, error) {
	return _RoninGovernance.contract.Transact(opts, "voteBridgeOperatorsBySignatures", _period, _operators, _signatures)
}

// VoteBridgeOperatorsBySignatures is a paid mutator transaction binding the contract method 0x56e237e8.
//
// Solidity: function voteBridgeOperatorsBySignatures(uint256 _period, address[] _operators, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_RoninGovernance *RoninGovernanceSession) VoteBridgeOperatorsBySignatures(_period *big.Int, _operators []common.Address, _signatures []RoninGovernanceSignature) (*types.Transaction, error) {
	return _RoninGovernance.Contract.VoteBridgeOperatorsBySignatures(&_RoninGovernance.TransactOpts, _period, _operators, _signatures)
}

// VoteBridgeOperatorsBySignatures is a paid mutator transaction binding the contract method 0x56e237e8.
//
// Solidity: function voteBridgeOperatorsBySignatures(uint256 _period, address[] _operators, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_RoninGovernance *RoninGovernanceTransactorSession) VoteBridgeOperatorsBySignatures(_period *big.Int, _operators []common.Address, _signatures []RoninGovernanceSignature) (*types.Transaction, error) {
	return _RoninGovernance.Contract.VoteBridgeOperatorsBySignatures(&_RoninGovernance.TransactOpts, _period, _operators, _signatures)
}
