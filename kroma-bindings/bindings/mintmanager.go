// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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
	_ = abi.ConvertType
)

// MintManagerMetaData contains all meta data concerning the MintManager contract.
var MintManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractGovernanceToken\",\"name\":\"_governanceToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_initMintPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slidingWindowBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_decayingFactor\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DECAYING_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DECAYING_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FLOOR_UNIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOVERNANCE_TOKEN\",\"outputs\":[{\"internalType\":\"contractGovernanceToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_MINT_PER_BLOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_TOKEN_CALLER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHARE_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLIDING_WINDOW_BLOCKS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_shares\",\"type\":\"uint256[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"mintAmountPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"shareOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x61010060405234801561001157600080fd5b506040516110ad3803806110ad8339810160408190526100309161004f565b6001600160a01b0390931660805260a09190915260c05260e05261009a565b6000806000806080858703121561006557600080fd5b84516001600160a01b038116811461007c57600080fd5b60208601516040870151606090970151919890975090945092505050565b60805160a05160c05160e051610f9461011960003960008181610260015281816102c60152610c0d01526000818161018501528181610aec01528181610b2501528181610b630152610bd3015260008181610201015281816102860152610b940152600081816101390152818161046c01526105df0152610f946000f3fe608060405234801561001057600080fd5b50600436106100df5760003560e01c806354fd4d501161008c5780637fbbe46f116100665780637fbbe46f1461022d578063894bee6114610223578063a103a2dd14610240578063e2da80901461025b57600080fd5b806354fd4d50146101b35780637b6a4cda146101fc5780637eb118451461022357600080fd5b80632efd46d6116100bd5780632efd46d61461013457806349ed1e6914610180578063544a54e3146101a757600080fd5b8063062459e6146100e45780631249c58b1461010a57806321e5e2c414610114575b600080fd5b6100f76100f2366004610c94565b610282565b6040519081526020015b60405180910390f35b61011261032e565b005b6100f7610122366004610cad565b60026020526000908152604090205481565b61015b7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610101565b6100f77f000000000000000000000000000000000000000000000000000000000000000081565b6100f76402540be40081565b6101ef6040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516101019190610cea565b6100f77f000000000000000000000000000000000000000000000000000000000000000081565b6100f7620186a081565b61011261023b366004610da9565b61065c565b61015b73deaddeaddeaddeaddeaddeaddeaddeaddead007081565b6100f77f000000000000000000000000000000000000000000000000000000000000000081565b60007f0000000000000000000000000000000000000000000000000000000000000000816102af84610ae3565b50905060015b8181101561032557620186a06102eb7f000000000000000000000000000000000000000000000000000000000000000085610e44565b6102f59190610eb0565b92506402540be4006103078185610eb0565b6103119190610e44565b92508061031d81610ec4565b9150506102b5565b50909392505050565b3373deaddeaddeaddeaddeaddeaddeaddeaddead0070146103d6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4d696e744d616e616765723a206f6e6c7920746865206d696e742063616c6c6560448201527f722069732061636365707465640000000000000000000000000000000000000060648201526084015b60405180910390fd5b4360035403610467576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f4d696e744d616e616765723a20746f6b656e73206861766520616c726561647960448201527f206265656e206d696e74656420696e207468697320626c6f636b00000000000060648201526084016103cd565b6000807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104f99190610efc565b111561050f5761050843610282565b905061051b565b61051843610b8f565b90505b80156106595760005b6001548110156106535760006001828154811061054357610543610f15565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff168083526002909152604082205490925090620186a06105868387610e44565b6105909190610eb0565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018390529192507f0000000000000000000000000000000000000000000000000000000000000000909116906340c10f1990604401600060405180830381600087803b15801561062557600080fd5b505af1158015610639573d6000803e3d6000fd5b50505050505050808061064b90610ec4565b915050610524565b50436003555b50565b600054610100900460ff161580801561067c5750600054600160ff909116105b806106965750303b158015610696575060005460ff166001145b610722576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016103cd565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561078057600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b83821461080e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f4d696e744d616e616765723a20696e76616c6964206c656e677468206f66206160448201527f727261790000000000000000000000000000000000000000000000000000000060648201526084016103cd565b6000805b85811015610a0b57600087878381811061082e5761082e610f15565b90506020020160208101906108439190610cad565b905073ffffffffffffffffffffffffffffffffffffffff81166108e8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f4d696e744d616e616765723a20726563697069656e742061646472657373206360448201527f616e6e6f7420626520300000000000000000000000000000000000000000000060648201526084016103cd565b60008686848181106108fc576108fc610f15565b9050602002013590508060000361096f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f4d696e744d616e616765723a2073686172652063616e6e6f742062652030000060448201526064016103cd565b6109798185610f44565b600180548082019091557fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf601805473ffffffffffffffffffffffffffffffffffffffff9094167fffffffffffffffffffffffff00000000000000000000000000000000000000009094168417905560009283526002602052604090922055915080610a0381610ec4565b915050610812565b50620186a08114610a78576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f4d696e744d616e616765723a20696e76616c696420736861726573000000000060448201526064016103cd565b508015610adc57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b60008080610b117f000000000000000000000000000000000000000000000000000000000000000085610eb0565b610b1c906001610f44565b90506000610b4a7f000000000000000000000000000000000000000000000000000000000000000086610f5c565b905080600003610b8557610b5f600183610f70565b91507f000000000000000000000000000000000000000000000000000000000000000090505b9094909350915050565b6000807f00000000000000000000000000000000000000000000000000000000000000008180610bbe86610ae3565b909250905060015b82811015610c6c57610bf87f000000000000000000000000000000000000000000000000000000000000000085610e44565b610c029086610f44565b9450620186a0610c327f000000000000000000000000000000000000000000000000000000000000000086610e44565b610c3c9190610eb0565b93506402540be400610c4e8186610eb0565b610c589190610e44565b935080610c6481610ec4565b915050610bc6565b508015610c8a57610c7d8184610e44565b610c879085610f44565b93505b5091949350505050565b600060208284031215610ca657600080fd5b5035919050565b600060208284031215610cbf57600080fd5b813573ffffffffffffffffffffffffffffffffffffffff81168114610ce357600080fd5b9392505050565b600060208083528351808285015260005b81811015610d1757858101830151858201604001528201610cfb565b81811115610d29576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b60008083601f840112610d6f57600080fd5b50813567ffffffffffffffff811115610d8757600080fd5b6020830191508360208260051b8501011115610da257600080fd5b9250929050565b60008060008060408587031215610dbf57600080fd5b843567ffffffffffffffff80821115610dd757600080fd5b610de388838901610d5d565b90965094506020870135915080821115610dfc57600080fd5b50610e0987828801610d5d565b95989497509550505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615610e7c57610e7c610e15565b500290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082610ebf57610ebf610e81565b500490565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610ef557610ef5610e15565b5060010190565b600060208284031215610f0e57600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008219821115610f5757610f57610e15565b500190565b600082610f6b57610f6b610e81565b500690565b600082821015610f8257610f82610e15565b50039056fea164736f6c634300080f000a",
}

// MintManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use MintManagerMetaData.ABI instead.
var MintManagerABI = MintManagerMetaData.ABI

// MintManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MintManagerMetaData.Bin instead.
var MintManagerBin = MintManagerMetaData.Bin

// DeployMintManager deploys a new Ethereum contract, binding an instance of MintManager to it.
func DeployMintManager(auth *bind.TransactOpts, backend bind.ContractBackend, _governanceToken common.Address, _initMintPerBlock *big.Int, _slidingWindowBlocks *big.Int, _decayingFactor *big.Int) (common.Address, *types.Transaction, *MintManager, error) {
	parsed, err := MintManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MintManagerBin), backend, _governanceToken, _initMintPerBlock, _slidingWindowBlocks, _decayingFactor)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MintManager{MintManagerCaller: MintManagerCaller{contract: contract}, MintManagerTransactor: MintManagerTransactor{contract: contract}, MintManagerFilterer: MintManagerFilterer{contract: contract}}, nil
}

// MintManager is an auto generated Go binding around an Ethereum contract.
type MintManager struct {
	MintManagerCaller     // Read-only binding to the contract
	MintManagerTransactor // Write-only binding to the contract
	MintManagerFilterer   // Log filterer for contract events
}

// MintManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MintManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MintManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MintManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MintManagerSession struct {
	Contract     *MintManager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MintManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MintManagerCallerSession struct {
	Contract *MintManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MintManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MintManagerTransactorSession struct {
	Contract     *MintManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MintManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MintManagerRaw struct {
	Contract *MintManager // Generic contract binding to access the raw methods on
}

// MintManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MintManagerCallerRaw struct {
	Contract *MintManagerCaller // Generic read-only contract binding to access the raw methods on
}

// MintManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MintManagerTransactorRaw struct {
	Contract *MintManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMintManager creates a new instance of MintManager, bound to a specific deployed contract.
func NewMintManager(address common.Address, backend bind.ContractBackend) (*MintManager, error) {
	contract, err := bindMintManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MintManager{MintManagerCaller: MintManagerCaller{contract: contract}, MintManagerTransactor: MintManagerTransactor{contract: contract}, MintManagerFilterer: MintManagerFilterer{contract: contract}}, nil
}

// NewMintManagerCaller creates a new read-only instance of MintManager, bound to a specific deployed contract.
func NewMintManagerCaller(address common.Address, caller bind.ContractCaller) (*MintManagerCaller, error) {
	contract, err := bindMintManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MintManagerCaller{contract: contract}, nil
}

// NewMintManagerTransactor creates a new write-only instance of MintManager, bound to a specific deployed contract.
func NewMintManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*MintManagerTransactor, error) {
	contract, err := bindMintManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MintManagerTransactor{contract: contract}, nil
}

// NewMintManagerFilterer creates a new log filterer instance of MintManager, bound to a specific deployed contract.
func NewMintManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*MintManagerFilterer, error) {
	contract, err := bindMintManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MintManagerFilterer{contract: contract}, nil
}

// bindMintManager binds a generic wrapper to an already deployed contract.
func bindMintManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MintManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintManager *MintManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MintManager.Contract.MintManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintManager *MintManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintManager.Contract.MintManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintManager *MintManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintManager.Contract.MintManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintManager *MintManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MintManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintManager *MintManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintManager *MintManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintManager.Contract.contract.Transact(opts, method, params...)
}

// DECAYINGDENOMINATOR is a free data retrieval call binding the contract method 0x894bee61.
//
// Solidity: function DECAYING_DENOMINATOR() view returns(uint256)
func (_MintManager *MintManagerCaller) DECAYINGDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "DECAYING_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DECAYINGDENOMINATOR is a free data retrieval call binding the contract method 0x894bee61.
//
// Solidity: function DECAYING_DENOMINATOR() view returns(uint256)
func (_MintManager *MintManagerSession) DECAYINGDENOMINATOR() (*big.Int, error) {
	return _MintManager.Contract.DECAYINGDENOMINATOR(&_MintManager.CallOpts)
}

// DECAYINGDENOMINATOR is a free data retrieval call binding the contract method 0x894bee61.
//
// Solidity: function DECAYING_DENOMINATOR() view returns(uint256)
func (_MintManager *MintManagerCallerSession) DECAYINGDENOMINATOR() (*big.Int, error) {
	return _MintManager.Contract.DECAYINGDENOMINATOR(&_MintManager.CallOpts)
}

// DECAYINGFACTOR is a free data retrieval call binding the contract method 0xe2da8090.
//
// Solidity: function DECAYING_FACTOR() view returns(uint256)
func (_MintManager *MintManagerCaller) DECAYINGFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "DECAYING_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DECAYINGFACTOR is a free data retrieval call binding the contract method 0xe2da8090.
//
// Solidity: function DECAYING_FACTOR() view returns(uint256)
func (_MintManager *MintManagerSession) DECAYINGFACTOR() (*big.Int, error) {
	return _MintManager.Contract.DECAYINGFACTOR(&_MintManager.CallOpts)
}

// DECAYINGFACTOR is a free data retrieval call binding the contract method 0xe2da8090.
//
// Solidity: function DECAYING_FACTOR() view returns(uint256)
func (_MintManager *MintManagerCallerSession) DECAYINGFACTOR() (*big.Int, error) {
	return _MintManager.Contract.DECAYINGFACTOR(&_MintManager.CallOpts)
}

// FLOORUNIT is a free data retrieval call binding the contract method 0x544a54e3.
//
// Solidity: function FLOOR_UNIT() view returns(uint256)
func (_MintManager *MintManagerCaller) FLOORUNIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "FLOOR_UNIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLOORUNIT is a free data retrieval call binding the contract method 0x544a54e3.
//
// Solidity: function FLOOR_UNIT() view returns(uint256)
func (_MintManager *MintManagerSession) FLOORUNIT() (*big.Int, error) {
	return _MintManager.Contract.FLOORUNIT(&_MintManager.CallOpts)
}

// FLOORUNIT is a free data retrieval call binding the contract method 0x544a54e3.
//
// Solidity: function FLOOR_UNIT() view returns(uint256)
func (_MintManager *MintManagerCallerSession) FLOORUNIT() (*big.Int, error) {
	return _MintManager.Contract.FLOORUNIT(&_MintManager.CallOpts)
}

// GOVERNANCETOKEN is a free data retrieval call binding the contract method 0x2efd46d6.
//
// Solidity: function GOVERNANCE_TOKEN() view returns(address)
func (_MintManager *MintManagerCaller) GOVERNANCETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "GOVERNANCE_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOVERNANCETOKEN is a free data retrieval call binding the contract method 0x2efd46d6.
//
// Solidity: function GOVERNANCE_TOKEN() view returns(address)
func (_MintManager *MintManagerSession) GOVERNANCETOKEN() (common.Address, error) {
	return _MintManager.Contract.GOVERNANCETOKEN(&_MintManager.CallOpts)
}

// GOVERNANCETOKEN is a free data retrieval call binding the contract method 0x2efd46d6.
//
// Solidity: function GOVERNANCE_TOKEN() view returns(address)
func (_MintManager *MintManagerCallerSession) GOVERNANCETOKEN() (common.Address, error) {
	return _MintManager.Contract.GOVERNANCETOKEN(&_MintManager.CallOpts)
}

// INITMINTPERBLOCK is a free data retrieval call binding the contract method 0x7b6a4cda.
//
// Solidity: function INIT_MINT_PER_BLOCK() view returns(uint256)
func (_MintManager *MintManagerCaller) INITMINTPERBLOCK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "INIT_MINT_PER_BLOCK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITMINTPERBLOCK is a free data retrieval call binding the contract method 0x7b6a4cda.
//
// Solidity: function INIT_MINT_PER_BLOCK() view returns(uint256)
func (_MintManager *MintManagerSession) INITMINTPERBLOCK() (*big.Int, error) {
	return _MintManager.Contract.INITMINTPERBLOCK(&_MintManager.CallOpts)
}

// INITMINTPERBLOCK is a free data retrieval call binding the contract method 0x7b6a4cda.
//
// Solidity: function INIT_MINT_PER_BLOCK() view returns(uint256)
func (_MintManager *MintManagerCallerSession) INITMINTPERBLOCK() (*big.Int, error) {
	return _MintManager.Contract.INITMINTPERBLOCK(&_MintManager.CallOpts)
}

// MINTTOKENCALLER is a free data retrieval call binding the contract method 0xa103a2dd.
//
// Solidity: function MINT_TOKEN_CALLER() view returns(address)
func (_MintManager *MintManagerCaller) MINTTOKENCALLER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "MINT_TOKEN_CALLER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MINTTOKENCALLER is a free data retrieval call binding the contract method 0xa103a2dd.
//
// Solidity: function MINT_TOKEN_CALLER() view returns(address)
func (_MintManager *MintManagerSession) MINTTOKENCALLER() (common.Address, error) {
	return _MintManager.Contract.MINTTOKENCALLER(&_MintManager.CallOpts)
}

// MINTTOKENCALLER is a free data retrieval call binding the contract method 0xa103a2dd.
//
// Solidity: function MINT_TOKEN_CALLER() view returns(address)
func (_MintManager *MintManagerCallerSession) MINTTOKENCALLER() (common.Address, error) {
	return _MintManager.Contract.MINTTOKENCALLER(&_MintManager.CallOpts)
}

// SHAREDENOMINATOR is a free data retrieval call binding the contract method 0x7eb11845.
//
// Solidity: function SHARE_DENOMINATOR() view returns(uint256)
func (_MintManager *MintManagerCaller) SHAREDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "SHARE_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SHAREDENOMINATOR is a free data retrieval call binding the contract method 0x7eb11845.
//
// Solidity: function SHARE_DENOMINATOR() view returns(uint256)
func (_MintManager *MintManagerSession) SHAREDENOMINATOR() (*big.Int, error) {
	return _MintManager.Contract.SHAREDENOMINATOR(&_MintManager.CallOpts)
}

// SHAREDENOMINATOR is a free data retrieval call binding the contract method 0x7eb11845.
//
// Solidity: function SHARE_DENOMINATOR() view returns(uint256)
func (_MintManager *MintManagerCallerSession) SHAREDENOMINATOR() (*big.Int, error) {
	return _MintManager.Contract.SHAREDENOMINATOR(&_MintManager.CallOpts)
}

// SLIDINGWINDOWBLOCKS is a free data retrieval call binding the contract method 0x49ed1e69.
//
// Solidity: function SLIDING_WINDOW_BLOCKS() view returns(uint256)
func (_MintManager *MintManagerCaller) SLIDINGWINDOWBLOCKS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "SLIDING_WINDOW_BLOCKS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SLIDINGWINDOWBLOCKS is a free data retrieval call binding the contract method 0x49ed1e69.
//
// Solidity: function SLIDING_WINDOW_BLOCKS() view returns(uint256)
func (_MintManager *MintManagerSession) SLIDINGWINDOWBLOCKS() (*big.Int, error) {
	return _MintManager.Contract.SLIDINGWINDOWBLOCKS(&_MintManager.CallOpts)
}

// SLIDINGWINDOWBLOCKS is a free data retrieval call binding the contract method 0x49ed1e69.
//
// Solidity: function SLIDING_WINDOW_BLOCKS() view returns(uint256)
func (_MintManager *MintManagerCallerSession) SLIDINGWINDOWBLOCKS() (*big.Int, error) {
	return _MintManager.Contract.SLIDINGWINDOWBLOCKS(&_MintManager.CallOpts)
}

// MintAmountPerBlock is a free data retrieval call binding the contract method 0x062459e6.
//
// Solidity: function mintAmountPerBlock(uint256 _blockNumber) view returns(uint256)
func (_MintManager *MintManagerCaller) MintAmountPerBlock(opts *bind.CallOpts, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "mintAmountPerBlock", _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintAmountPerBlock is a free data retrieval call binding the contract method 0x062459e6.
//
// Solidity: function mintAmountPerBlock(uint256 _blockNumber) view returns(uint256)
func (_MintManager *MintManagerSession) MintAmountPerBlock(_blockNumber *big.Int) (*big.Int, error) {
	return _MintManager.Contract.MintAmountPerBlock(&_MintManager.CallOpts, _blockNumber)
}

// MintAmountPerBlock is a free data retrieval call binding the contract method 0x062459e6.
//
// Solidity: function mintAmountPerBlock(uint256 _blockNumber) view returns(uint256)
func (_MintManager *MintManagerCallerSession) MintAmountPerBlock(_blockNumber *big.Int) (*big.Int, error) {
	return _MintManager.Contract.MintAmountPerBlock(&_MintManager.CallOpts, _blockNumber)
}

// ShareOf is a free data retrieval call binding the contract method 0x21e5e2c4.
//
// Solidity: function shareOf(address ) view returns(uint256)
func (_MintManager *MintManagerCaller) ShareOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "shareOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ShareOf is a free data retrieval call binding the contract method 0x21e5e2c4.
//
// Solidity: function shareOf(address ) view returns(uint256)
func (_MintManager *MintManagerSession) ShareOf(arg0 common.Address) (*big.Int, error) {
	return _MintManager.Contract.ShareOf(&_MintManager.CallOpts, arg0)
}

// ShareOf is a free data retrieval call binding the contract method 0x21e5e2c4.
//
// Solidity: function shareOf(address ) view returns(uint256)
func (_MintManager *MintManagerCallerSession) ShareOf(arg0 common.Address) (*big.Int, error) {
	return _MintManager.Contract.ShareOf(&_MintManager.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_MintManager *MintManagerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_MintManager *MintManagerSession) Version() (string, error) {
	return _MintManager.Contract.Version(&_MintManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_MintManager *MintManagerCallerSession) Version() (string, error) {
	return _MintManager.Contract.Version(&_MintManager.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x7fbbe46f.
//
// Solidity: function initialize(address[] _recipients, uint256[] _shares) returns()
func (_MintManager *MintManagerTransactor) Initialize(opts *bind.TransactOpts, _recipients []common.Address, _shares []*big.Int) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "initialize", _recipients, _shares)
}

// Initialize is a paid mutator transaction binding the contract method 0x7fbbe46f.
//
// Solidity: function initialize(address[] _recipients, uint256[] _shares) returns()
func (_MintManager *MintManagerSession) Initialize(_recipients []common.Address, _shares []*big.Int) (*types.Transaction, error) {
	return _MintManager.Contract.Initialize(&_MintManager.TransactOpts, _recipients, _shares)
}

// Initialize is a paid mutator transaction binding the contract method 0x7fbbe46f.
//
// Solidity: function initialize(address[] _recipients, uint256[] _shares) returns()
func (_MintManager *MintManagerTransactorSession) Initialize(_recipients []common.Address, _shares []*big.Int) (*types.Transaction, error) {
	return _MintManager.Contract.Initialize(&_MintManager.TransactOpts, _recipients, _shares)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns()
func (_MintManager *MintManagerTransactor) Mint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "mint")
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns()
func (_MintManager *MintManagerSession) Mint() (*types.Transaction, error) {
	return _MintManager.Contract.Mint(&_MintManager.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns()
func (_MintManager *MintManagerTransactorSession) Mint() (*types.Transaction, error) {
	return _MintManager.Contract.Mint(&_MintManager.TransactOpts)
}

// MintManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MintManager contract.
type MintManagerInitializedIterator struct {
	Event *MintManagerInitialized // Event containing the contract specifics and raw log

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
func (it *MintManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintManagerInitialized)
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
		it.Event = new(MintManagerInitialized)
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
func (it *MintManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintManagerInitialized represents a Initialized event raised by the MintManager contract.
type MintManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MintManager *MintManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*MintManagerInitializedIterator, error) {

	logs, sub, err := _MintManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MintManagerInitializedIterator{contract: _MintManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MintManager *MintManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MintManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _MintManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintManagerInitialized)
				if err := _MintManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_MintManager *MintManagerFilterer) ParseInitialized(log types.Log) (*MintManagerInitialized, error) {
	event := new(MintManagerInitialized)
	if err := _MintManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
