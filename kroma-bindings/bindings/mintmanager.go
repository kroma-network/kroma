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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintActivatedBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_initMintPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slidingWindowBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_decayingFactor\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DECAYING_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DECAYING_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FLOOR_UNIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOVERNANCE_TOKEN\",\"outputs\":[{\"internalType\":\"contractGovernanceToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_MINT_PER_BLOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_ACTIVATED_BLOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHARE_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLIDING_WINDOW_BLOCKS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_shares\",\"type\":\"uint256[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"mintAmountPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"shareOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x61012060405234801561001157600080fd5b5060405161102638038061102683398101604081905261003091610060565b60a09390935260c09190915260e052610100527342000000000000000000000000000000000000ff608052610096565b6000806000806080858703121561007657600080fd5b505082516020840151604085015160609095015191969095509092509050565b60805160a05160c05160e05161010051610f0d6101196000396000818161026c015281816102d20152610b9f01526000818161018501528181610a9101528181610ad60152610b65015260008181610201015281816102920152610b2601526000818161024501526103e501526000818161013901526105860152610f0d6000f3fe608060405234801561001057600080fd5b50600436106100df5760003560e01c806354fd4d501161008c5780637fbbe46f116100665780637fbbe46f1461022d57806381e9090e14610240578063894bee6114610223578063e2da80901461026757600080fd5b806354fd4d50146101b35780637b6a4cda146101fc5780637eb118451461022357600080fd5b80632efd46d6116100bd5780632efd46d61461013457806349ed1e6914610180578063544a54e3146101a757600080fd5b8063062459e6146100e45780631249c58b1461010a57806321e5e2c414610114575b600080fd5b6100f76100f2366004610c26565b61028e565b6040519081526020015b60405180910390f35b61011261033a565b005b6100f7610122366004610c3f565b60026020526000908152604090205481565b61015b7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610101565b6100f77f000000000000000000000000000000000000000000000000000000000000000081565b6100f76402540be40081565b6101ef6040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516101019190610c7c565b6100f77f000000000000000000000000000000000000000000000000000000000000000081565b6100f7620186a081565b61011261023b366004610d3b565b610604565b6100f77f000000000000000000000000000000000000000000000000000000000000000081565b6100f77f000000000000000000000000000000000000000000000000000000000000000081565b60007f0000000000000000000000000000000000000000000000000000000000000000816102bb84610a8b565b50905060015b8181101561033157620186a06102f77f000000000000000000000000000000000000000000000000000000000000000085610dd6565b6103019190610e42565b92506402540be4006103138185610e42565b61031d9190610dd6565b92508061032981610e56565b9150506102c1565b50909392505050565b33734200000000000000000000000000000000000002146103e2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603460248201527f4d696e744d616e616765723a206f6e6c7920746865204c31426c6f636b20636160448201527f6e2063616c6c20746869732066756e6374696f6e00000000000000000000000060648201526084015b60405180910390fd5b437f00000000000000000000000000000000000000000000000000000000000000001161060257436003540361049a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f4d696e744d616e616765723a20746f6b656e73206861766520616c726561647960448201527f206265656e206d696e74656420696e207468697320626c6f636b00000000000060648201526084016103d9565b60006003546000036104b6576104af43610b21565b90506104c2565b6104bf4361028e565b90505b80156106005760005b6001548110156105fa576000600182815481106104ea576104ea610e8e565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff168083526002909152604082205490925090620186a061052d8387610dd6565b6105379190610e42565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018390529192507f0000000000000000000000000000000000000000000000000000000000000000909116906340c10f1990604401600060405180830381600087803b1580156105cc57600080fd5b505af11580156105e0573d6000803e3d6000fd5b5050505050505080806105f290610e56565b9150506104cb565b50436003555b505b565b600054610100900460ff16158080156106245750600054600160ff909116105b8061063e5750303b15801561063e575060005460ff166001145b6106ca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016103d9565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561072857600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b8382146107b6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f4d696e744d616e616765723a20696e76616c6964206c656e677468206f66206160448201527f727261790000000000000000000000000000000000000000000000000000000060648201526084016103d9565b6000805b858110156109b35760008787838181106107d6576107d6610e8e565b90506020020160208101906107eb9190610c3f565b905073ffffffffffffffffffffffffffffffffffffffff8116610890576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f4d696e744d616e616765723a20726563697069656e742061646472657373206360448201527f616e6e6f7420626520300000000000000000000000000000000000000000000060648201526084016103d9565b60008686848181106108a4576108a4610e8e565b90506020020135905080600003610917576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f4d696e744d616e616765723a2073686172652063616e6e6f742062652030000060448201526064016103d9565b6109218185610ebd565b600180548082019091557fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf601805473ffffffffffffffffffffffffffffffffffffffff9094167fffffffffffffffffffffffff000000000000000000000000000000000000000090941684179055600092835260026020526040909220559150806109ab81610e56565b9150506107ba565b50620186a08114610a20576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f4d696e744d616e616765723a20696e76616c696420736861726573000000000060448201526064016103d9565b508015610a8457600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b600080807f0000000000000000000000000000000000000000000000000000000000000000610abb600186610ed5565b610ac59190610e42565b610ad0906001610ebd565b905060007f0000000000000000000000000000000000000000000000000000000000000000610b00600187610ed5565b610b0a9190610eec565b610b15906001610ebd565b91959194509092505050565b6000807f00000000000000000000000000000000000000000000000000000000000000008180610b5086610a8b565b909250905060015b82811015610bfe57610b8a7f000000000000000000000000000000000000000000000000000000000000000085610dd6565b610b949086610ebd565b9450620186a0610bc47f000000000000000000000000000000000000000000000000000000000000000086610dd6565b610bce9190610e42565b93506402540be400610be08186610e42565b610bea9190610dd6565b935080610bf681610e56565b915050610b58565b508015610c1c57610c0f8184610dd6565b610c199085610ebd565b93505b5091949350505050565b600060208284031215610c3857600080fd5b5035919050565b600060208284031215610c5157600080fd5b813573ffffffffffffffffffffffffffffffffffffffff81168114610c7557600080fd5b9392505050565b600060208083528351808285015260005b81811015610ca957858101830151858201604001528201610c8d565b81811115610cbb576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b60008083601f840112610d0157600080fd5b50813567ffffffffffffffff811115610d1957600080fd5b6020830191508360208260051b8501011115610d3457600080fd5b9250929050565b60008060008060408587031215610d5157600080fd5b843567ffffffffffffffff80821115610d6957600080fd5b610d7588838901610cef565b90965094506020870135915080821115610d8e57600080fd5b50610d9b87828801610cef565b95989497509550505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615610e0e57610e0e610da7565b500290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082610e5157610e51610e13565b500490565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610e8757610e87610da7565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008219821115610ed057610ed0610da7565b500190565b600082821015610ee757610ee7610da7565b500390565b600082610efb57610efb610e13565b50069056fea164736f6c634300080f000a",
}

// MintManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use MintManagerMetaData.ABI instead.
var MintManagerABI = MintManagerMetaData.ABI

// MintManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MintManagerMetaData.Bin instead.
var MintManagerBin = MintManagerMetaData.Bin

// DeployMintManager deploys a new Ethereum contract, binding an instance of MintManager to it.
func DeployMintManager(auth *bind.TransactOpts, backend bind.ContractBackend, _mintActivatedBlock *big.Int, _initMintPerBlock *big.Int, _slidingWindowBlocks *big.Int, _decayingFactor *big.Int) (common.Address, *types.Transaction, *MintManager, error) {
	parsed, err := MintManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MintManagerBin), backend, _mintActivatedBlock, _initMintPerBlock, _slidingWindowBlocks, _decayingFactor)
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

// MINTACTIVATEDBLOCK is a free data retrieval call binding the contract method 0x81e9090e.
//
// Solidity: function MINT_ACTIVATED_BLOCK() view returns(uint256)
func (_MintManager *MintManagerCaller) MINTACTIVATEDBLOCK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "MINT_ACTIVATED_BLOCK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTACTIVATEDBLOCK is a free data retrieval call binding the contract method 0x81e9090e.
//
// Solidity: function MINT_ACTIVATED_BLOCK() view returns(uint256)
func (_MintManager *MintManagerSession) MINTACTIVATEDBLOCK() (*big.Int, error) {
	return _MintManager.Contract.MINTACTIVATEDBLOCK(&_MintManager.CallOpts)
}

// MINTACTIVATEDBLOCK is a free data retrieval call binding the contract method 0x81e9090e.
//
// Solidity: function MINT_ACTIVATED_BLOCK() view returns(uint256)
func (_MintManager *MintManagerCallerSession) MINTACTIVATEDBLOCK() (*big.Int, error) {
	return _MintManager.Contract.MINTACTIVATEDBLOCK(&_MintManager.CallOpts)
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
