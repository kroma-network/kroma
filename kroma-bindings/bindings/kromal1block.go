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

// KromaL1BlockMetaData contains all meta data concerning the KromaL1Block contract.
var KromaL1BlockMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DEPOSITOR_ACCOUNT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"baseFeeScalar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"basefee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batcherHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blobBaseFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blobBaseFeeScalar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l1FeeOverhead\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l1FeeScalar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"number\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sequenceNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setL1BlockValues\",\"inputs\":[{\"name\":\"_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_basefee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_batcherHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_l1FeeOverhead\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_l1FeeScalar\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_validatorRewardScalar\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setL1BlockValuesEcotone\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"timestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorRewardScalar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610624806100206000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80639e8c496611610097578063e81b2c6d11610066578063e81b2c6d14610281578063ed579ad31461028a578063efc674eb14610293578063f8206140146102a657600080fd5b80639e8c4966146101f8578063b80777ea14610201578063c598591814610221578063e591b2821461024157600080fd5b806364ca23ef116100d357806364ca23ef1461017d57806368d5dca6146101aa5780638381f58a146101db5780638b239f73146101ef57600080fd5b806309bd5a6014610105578063440a5e201461012157806354fd4d501461012b5780635cf2496914610174575b600080fd5b61010e60025481565b6040519081526020015b60405180910390f35b6101296102af565b005b6101676040518060400160405280600581526020017f312e312e3000000000000000000000000000000000000000000000000000000081525081565b604051610118919061050b565b61010e60015481565b6003546101919067ffffffffffffffff1681565b60405167ffffffffffffffff9091168152602001610118565b6003546101c69068010000000000000000900463ffffffff1681565b60405163ffffffff9091168152602001610118565b6000546101919067ffffffffffffffff1681565b61010e60055481565b61010e60065481565b6000546101919068010000000000000000900467ffffffffffffffff1681565b6003546101c6906c01000000000000000000000000900463ffffffff1681565b61025c73deaddeaddeaddeaddeaddeaddeaddeaddead000181565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610118565b61010e60045481565b61010e60075481565b6101296102a136600461059b565b61030a565b61010e60085481565b3373deaddeaddeaddeaddeaddeaddeaddeaddead0001146102d857633cc50b456000526004601cfd5b60043560801c60035560143560801c60005560243560015560443560085560643560025560843560045560a435600755565b3373deaddeaddeaddeaddeaddeaddeaddeaddead0001146103b2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603b60248201527f4c31426c6f636b3a206f6e6c7920746865206465706f7369746f72206163636f60448201527f756e742063616e20736574204c3120626c6f636b2076616c756573000000000060648201526084015b60405180910390fd5b61271081111561046a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f4c31426c6f636b3a20746865206d61782076616c7565206f662076616c69646160448201527f746f7220726577617264207363616c617220686173206265656e20657863656560648201527f6465640000000000000000000000000000000000000000000000000000000000608482015260a4016103a9565b6000805467ffffffffffffffff998a1668010000000000000000027fffffffffffffffffffffffffffffffff000000000000000000000000000000009091169a8a169a909a179990991790985560019590955560029390935560038054929095167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009290921691909117909355600492909255600591909155600655600755565b600060208083528351808285015260005b818110156105385785810183015185820160400152820161051c565b8181111561054a576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b803567ffffffffffffffff8116811461059657600080fd5b919050565b60008060008060008060008060006101208a8c0312156105ba57600080fd5b6105c38a61057e565b98506105d160208b0161057e565b975060408a0135965060608a013595506105ed60808b0161057e565b989b979a50959894979660a0860135965060c08601359560e081013595506101000135935091505056fea164736f6c634300080f000a",
}

// KromaL1BlockABI is the input ABI used to generate the binding from.
// Deprecated: Use KromaL1BlockMetaData.ABI instead.
var KromaL1BlockABI = KromaL1BlockMetaData.ABI

// KromaL1BlockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use KromaL1BlockMetaData.Bin instead.
var KromaL1BlockBin = KromaL1BlockMetaData.Bin

// DeployKromaL1Block deploys a new Ethereum contract, binding an instance of KromaL1Block to it.
func DeployKromaL1Block(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KromaL1Block, error) {
	parsed, err := KromaL1BlockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KromaL1BlockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KromaL1Block{KromaL1BlockCaller: KromaL1BlockCaller{contract: contract}, KromaL1BlockTransactor: KromaL1BlockTransactor{contract: contract}, KromaL1BlockFilterer: KromaL1BlockFilterer{contract: contract}}, nil
}

// KromaL1Block is an auto generated Go binding around an Ethereum contract.
type KromaL1Block struct {
	KromaL1BlockCaller     // Read-only binding to the contract
	KromaL1BlockTransactor // Write-only binding to the contract
	KromaL1BlockFilterer   // Log filterer for contract events
}

// KromaL1BlockCaller is an auto generated read-only Go binding around an Ethereum contract.
type KromaL1BlockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KromaL1BlockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KromaL1BlockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KromaL1BlockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KromaL1BlockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KromaL1BlockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KromaL1BlockSession struct {
	Contract     *KromaL1Block     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KromaL1BlockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KromaL1BlockCallerSession struct {
	Contract *KromaL1BlockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// KromaL1BlockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KromaL1BlockTransactorSession struct {
	Contract     *KromaL1BlockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// KromaL1BlockRaw is an auto generated low-level Go binding around an Ethereum contract.
type KromaL1BlockRaw struct {
	Contract *KromaL1Block // Generic contract binding to access the raw methods on
}

// KromaL1BlockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KromaL1BlockCallerRaw struct {
	Contract *KromaL1BlockCaller // Generic read-only contract binding to access the raw methods on
}

// KromaL1BlockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KromaL1BlockTransactorRaw struct {
	Contract *KromaL1BlockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKromaL1Block creates a new instance of KromaL1Block, bound to a specific deployed contract.
func NewKromaL1Block(address common.Address, backend bind.ContractBackend) (*KromaL1Block, error) {
	contract, err := bindKromaL1Block(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KromaL1Block{KromaL1BlockCaller: KromaL1BlockCaller{contract: contract}, KromaL1BlockTransactor: KromaL1BlockTransactor{contract: contract}, KromaL1BlockFilterer: KromaL1BlockFilterer{contract: contract}}, nil
}

// NewKromaL1BlockCaller creates a new read-only instance of KromaL1Block, bound to a specific deployed contract.
func NewKromaL1BlockCaller(address common.Address, caller bind.ContractCaller) (*KromaL1BlockCaller, error) {
	contract, err := bindKromaL1Block(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KromaL1BlockCaller{contract: contract}, nil
}

// NewKromaL1BlockTransactor creates a new write-only instance of KromaL1Block, bound to a specific deployed contract.
func NewKromaL1BlockTransactor(address common.Address, transactor bind.ContractTransactor) (*KromaL1BlockTransactor, error) {
	contract, err := bindKromaL1Block(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KromaL1BlockTransactor{contract: contract}, nil
}

// NewKromaL1BlockFilterer creates a new log filterer instance of KromaL1Block, bound to a specific deployed contract.
func NewKromaL1BlockFilterer(address common.Address, filterer bind.ContractFilterer) (*KromaL1BlockFilterer, error) {
	contract, err := bindKromaL1Block(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KromaL1BlockFilterer{contract: contract}, nil
}

// bindKromaL1Block binds a generic wrapper to an already deployed contract.
func bindKromaL1Block(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KromaL1BlockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KromaL1Block *KromaL1BlockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KromaL1Block.Contract.KromaL1BlockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KromaL1Block *KromaL1BlockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KromaL1Block.Contract.KromaL1BlockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KromaL1Block *KromaL1BlockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KromaL1Block.Contract.KromaL1BlockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KromaL1Block *KromaL1BlockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KromaL1Block.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KromaL1Block *KromaL1BlockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KromaL1Block.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KromaL1Block *KromaL1BlockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KromaL1Block.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITORACCOUNT is a free data retrieval call binding the contract method 0xe591b282.
//
// Solidity: function DEPOSITOR_ACCOUNT() view returns(address)
func (_KromaL1Block *KromaL1BlockCaller) DEPOSITORACCOUNT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KromaL1Block.contract.Call(opts, &out, "DEPOSITOR_ACCOUNT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEPOSITORACCOUNT is a free data retrieval call binding the contract method 0xe591b282.
//
// Solidity: function DEPOSITOR_ACCOUNT() view returns(address)
func (_KromaL1Block *KromaL1BlockSession) DEPOSITORACCOUNT() (common.Address, error) {
	return _KromaL1Block.Contract.DEPOSITORACCOUNT(&_KromaL1Block.CallOpts)
}

// DEPOSITORACCOUNT is a free data retrieval call binding the contract method 0xe591b282.
//
// Solidity: function DEPOSITOR_ACCOUNT() view returns(address)
func (_KromaL1Block *KromaL1BlockCallerSession) DEPOSITORACCOUNT() (common.Address, error) {
	return _KromaL1Block.Contract.DEPOSITORACCOUNT(&_KromaL1Block.CallOpts)
}

// BaseFeeScalar is a free data retrieval call binding the contract method 0xc5985918.
//
// Solidity: function baseFeeScalar() view returns(uint32)
func (_KromaL1Block *KromaL1BlockCaller) BaseFeeScalar(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _KromaL1Block.contract.Call(opts, &out, "baseFeeScalar")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// BaseFeeScalar is a free data retrieval call binding the contract method 0xc5985918.
//
// Solidity: function baseFeeScalar() view returns(uint32)
func (_KromaL1Block *KromaL1BlockSession) BaseFeeScalar() (uint32, error) {
	return _KromaL1Block.Contract.BaseFeeScalar(&_KromaL1Block.CallOpts)
}

// BaseFeeScalar is a free data retrieval call binding the contract method 0xc5985918.
//
// Solidity: function baseFeeScalar() view returns(uint32)
func (_KromaL1Block *KromaL1BlockCallerSession) BaseFeeScalar() (uint32, error) {
	return _KromaL1Block.Contract.BaseFeeScalar(&_KromaL1Block.CallOpts)
}

// Basefee is a free data retrieval call binding the contract method 0x5cf24969.
//
// Solidity: function basefee() view returns(uint256)
func (_KromaL1Block *KromaL1BlockCaller) Basefee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KromaL1Block.contract.Call(opts, &out, "basefee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Basefee is a free data retrieval call binding the contract method 0x5cf24969.
//
// Solidity: function basefee() view returns(uint256)
func (_KromaL1Block *KromaL1BlockSession) Basefee() (*big.Int, error) {
	return _KromaL1Block.Contract.Basefee(&_KromaL1Block.CallOpts)
}

// Basefee is a free data retrieval call binding the contract method 0x5cf24969.
//
// Solidity: function basefee() view returns(uint256)
func (_KromaL1Block *KromaL1BlockCallerSession) Basefee() (*big.Int, error) {
	return _KromaL1Block.Contract.Basefee(&_KromaL1Block.CallOpts)
}

// BatcherHash is a free data retrieval call binding the contract method 0xe81b2c6d.
//
// Solidity: function batcherHash() view returns(bytes32)
func (_KromaL1Block *KromaL1BlockCaller) BatcherHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KromaL1Block.contract.Call(opts, &out, "batcherHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BatcherHash is a free data retrieval call binding the contract method 0xe81b2c6d.
//
// Solidity: function batcherHash() view returns(bytes32)
func (_KromaL1Block *KromaL1BlockSession) BatcherHash() ([32]byte, error) {
	return _KromaL1Block.Contract.BatcherHash(&_KromaL1Block.CallOpts)
}

// BatcherHash is a free data retrieval call binding the contract method 0xe81b2c6d.
//
// Solidity: function batcherHash() view returns(bytes32)
func (_KromaL1Block *KromaL1BlockCallerSession) BatcherHash() ([32]byte, error) {
	return _KromaL1Block.Contract.BatcherHash(&_KromaL1Block.CallOpts)
}

// BlobBaseFee is a free data retrieval call binding the contract method 0xf8206140.
//
// Solidity: function blobBaseFee() view returns(uint256)
func (_KromaL1Block *KromaL1BlockCaller) BlobBaseFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KromaL1Block.contract.Call(opts, &out, "blobBaseFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlobBaseFee is a free data retrieval call binding the contract method 0xf8206140.
//
// Solidity: function blobBaseFee() view returns(uint256)
func (_KromaL1Block *KromaL1BlockSession) BlobBaseFee() (*big.Int, error) {
	return _KromaL1Block.Contract.BlobBaseFee(&_KromaL1Block.CallOpts)
}

// BlobBaseFee is a free data retrieval call binding the contract method 0xf8206140.
//
// Solidity: function blobBaseFee() view returns(uint256)
func (_KromaL1Block *KromaL1BlockCallerSession) BlobBaseFee() (*big.Int, error) {
	return _KromaL1Block.Contract.BlobBaseFee(&_KromaL1Block.CallOpts)
}

// BlobBaseFeeScalar is a free data retrieval call binding the contract method 0x68d5dca6.
//
// Solidity: function blobBaseFeeScalar() view returns(uint32)
func (_KromaL1Block *KromaL1BlockCaller) BlobBaseFeeScalar(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _KromaL1Block.contract.Call(opts, &out, "blobBaseFeeScalar")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// BlobBaseFeeScalar is a free data retrieval call binding the contract method 0x68d5dca6.
//
// Solidity: function blobBaseFeeScalar() view returns(uint32)
func (_KromaL1Block *KromaL1BlockSession) BlobBaseFeeScalar() (uint32, error) {
	return _KromaL1Block.Contract.BlobBaseFeeScalar(&_KromaL1Block.CallOpts)
}

// BlobBaseFeeScalar is a free data retrieval call binding the contract method 0x68d5dca6.
//
// Solidity: function blobBaseFeeScalar() view returns(uint32)
func (_KromaL1Block *KromaL1BlockCallerSession) BlobBaseFeeScalar() (uint32, error) {
	return _KromaL1Block.Contract.BlobBaseFeeScalar(&_KromaL1Block.CallOpts)
}

// Hash is a free data retrieval call binding the contract method 0x09bd5a60.
//
// Solidity: function hash() view returns(bytes32)
func (_KromaL1Block *KromaL1BlockCaller) Hash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KromaL1Block.contract.Call(opts, &out, "hash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash is a free data retrieval call binding the contract method 0x09bd5a60.
//
// Solidity: function hash() view returns(bytes32)
func (_KromaL1Block *KromaL1BlockSession) Hash() ([32]byte, error) {
	return _KromaL1Block.Contract.Hash(&_KromaL1Block.CallOpts)
}

// Hash is a free data retrieval call binding the contract method 0x09bd5a60.
//
// Solidity: function hash() view returns(bytes32)
func (_KromaL1Block *KromaL1BlockCallerSession) Hash() ([32]byte, error) {
	return _KromaL1Block.Contract.Hash(&_KromaL1Block.CallOpts)
}

// L1FeeOverhead is a free data retrieval call binding the contract method 0x8b239f73.
//
// Solidity: function l1FeeOverhead() view returns(uint256)
func (_KromaL1Block *KromaL1BlockCaller) L1FeeOverhead(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KromaL1Block.contract.Call(opts, &out, "l1FeeOverhead")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L1FeeOverhead is a free data retrieval call binding the contract method 0x8b239f73.
//
// Solidity: function l1FeeOverhead() view returns(uint256)
func (_KromaL1Block *KromaL1BlockSession) L1FeeOverhead() (*big.Int, error) {
	return _KromaL1Block.Contract.L1FeeOverhead(&_KromaL1Block.CallOpts)
}

// L1FeeOverhead is a free data retrieval call binding the contract method 0x8b239f73.
//
// Solidity: function l1FeeOverhead() view returns(uint256)
func (_KromaL1Block *KromaL1BlockCallerSession) L1FeeOverhead() (*big.Int, error) {
	return _KromaL1Block.Contract.L1FeeOverhead(&_KromaL1Block.CallOpts)
}

// L1FeeScalar is a free data retrieval call binding the contract method 0x9e8c4966.
//
// Solidity: function l1FeeScalar() view returns(uint256)
func (_KromaL1Block *KromaL1BlockCaller) L1FeeScalar(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KromaL1Block.contract.Call(opts, &out, "l1FeeScalar")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L1FeeScalar is a free data retrieval call binding the contract method 0x9e8c4966.
//
// Solidity: function l1FeeScalar() view returns(uint256)
func (_KromaL1Block *KromaL1BlockSession) L1FeeScalar() (*big.Int, error) {
	return _KromaL1Block.Contract.L1FeeScalar(&_KromaL1Block.CallOpts)
}

// L1FeeScalar is a free data retrieval call binding the contract method 0x9e8c4966.
//
// Solidity: function l1FeeScalar() view returns(uint256)
func (_KromaL1Block *KromaL1BlockCallerSession) L1FeeScalar() (*big.Int, error) {
	return _KromaL1Block.Contract.L1FeeScalar(&_KromaL1Block.CallOpts)
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint64)
func (_KromaL1Block *KromaL1BlockCaller) Number(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _KromaL1Block.contract.Call(opts, &out, "number")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint64)
func (_KromaL1Block *KromaL1BlockSession) Number() (uint64, error) {
	return _KromaL1Block.Contract.Number(&_KromaL1Block.CallOpts)
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() view returns(uint64)
func (_KromaL1Block *KromaL1BlockCallerSession) Number() (uint64, error) {
	return _KromaL1Block.Contract.Number(&_KromaL1Block.CallOpts)
}

// SequenceNumber is a free data retrieval call binding the contract method 0x64ca23ef.
//
// Solidity: function sequenceNumber() view returns(uint64)
func (_KromaL1Block *KromaL1BlockCaller) SequenceNumber(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _KromaL1Block.contract.Call(opts, &out, "sequenceNumber")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SequenceNumber is a free data retrieval call binding the contract method 0x64ca23ef.
//
// Solidity: function sequenceNumber() view returns(uint64)
func (_KromaL1Block *KromaL1BlockSession) SequenceNumber() (uint64, error) {
	return _KromaL1Block.Contract.SequenceNumber(&_KromaL1Block.CallOpts)
}

// SequenceNumber is a free data retrieval call binding the contract method 0x64ca23ef.
//
// Solidity: function sequenceNumber() view returns(uint64)
func (_KromaL1Block *KromaL1BlockCallerSession) SequenceNumber() (uint64, error) {
	return _KromaL1Block.Contract.SequenceNumber(&_KromaL1Block.CallOpts)
}

// Timestamp is a free data retrieval call binding the contract method 0xb80777ea.
//
// Solidity: function timestamp() view returns(uint64)
func (_KromaL1Block *KromaL1BlockCaller) Timestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _KromaL1Block.contract.Call(opts, &out, "timestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Timestamp is a free data retrieval call binding the contract method 0xb80777ea.
//
// Solidity: function timestamp() view returns(uint64)
func (_KromaL1Block *KromaL1BlockSession) Timestamp() (uint64, error) {
	return _KromaL1Block.Contract.Timestamp(&_KromaL1Block.CallOpts)
}

// Timestamp is a free data retrieval call binding the contract method 0xb80777ea.
//
// Solidity: function timestamp() view returns(uint64)
func (_KromaL1Block *KromaL1BlockCallerSession) Timestamp() (uint64, error) {
	return _KromaL1Block.Contract.Timestamp(&_KromaL1Block.CallOpts)
}

// ValidatorRewardScalar is a free data retrieval call binding the contract method 0xed579ad3.
//
// Solidity: function validatorRewardScalar() view returns(uint256)
func (_KromaL1Block *KromaL1BlockCaller) ValidatorRewardScalar(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KromaL1Block.contract.Call(opts, &out, "validatorRewardScalar")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorRewardScalar is a free data retrieval call binding the contract method 0xed579ad3.
//
// Solidity: function validatorRewardScalar() view returns(uint256)
func (_KromaL1Block *KromaL1BlockSession) ValidatorRewardScalar() (*big.Int, error) {
	return _KromaL1Block.Contract.ValidatorRewardScalar(&_KromaL1Block.CallOpts)
}

// ValidatorRewardScalar is a free data retrieval call binding the contract method 0xed579ad3.
//
// Solidity: function validatorRewardScalar() view returns(uint256)
func (_KromaL1Block *KromaL1BlockCallerSession) ValidatorRewardScalar() (*big.Int, error) {
	return _KromaL1Block.Contract.ValidatorRewardScalar(&_KromaL1Block.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_KromaL1Block *KromaL1BlockCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KromaL1Block.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_KromaL1Block *KromaL1BlockSession) Version() (string, error) {
	return _KromaL1Block.Contract.Version(&_KromaL1Block.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_KromaL1Block *KromaL1BlockCallerSession) Version() (string, error) {
	return _KromaL1Block.Contract.Version(&_KromaL1Block.CallOpts)
}

// SetL1BlockValues is a paid mutator transaction binding the contract method 0xefc674eb.
//
// Solidity: function setL1BlockValues(uint64 _number, uint64 _timestamp, uint256 _basefee, bytes32 _hash, uint64 _sequenceNumber, bytes32 _batcherHash, uint256 _l1FeeOverhead, uint256 _l1FeeScalar, uint256 _validatorRewardScalar) returns()
func (_KromaL1Block *KromaL1BlockTransactor) SetL1BlockValues(opts *bind.TransactOpts, _number uint64, _timestamp uint64, _basefee *big.Int, _hash [32]byte, _sequenceNumber uint64, _batcherHash [32]byte, _l1FeeOverhead *big.Int, _l1FeeScalar *big.Int, _validatorRewardScalar *big.Int) (*types.Transaction, error) {
	return _KromaL1Block.contract.Transact(opts, "setL1BlockValues", _number, _timestamp, _basefee, _hash, _sequenceNumber, _batcherHash, _l1FeeOverhead, _l1FeeScalar, _validatorRewardScalar)
}

// SetL1BlockValues is a paid mutator transaction binding the contract method 0xefc674eb.
//
// Solidity: function setL1BlockValues(uint64 _number, uint64 _timestamp, uint256 _basefee, bytes32 _hash, uint64 _sequenceNumber, bytes32 _batcherHash, uint256 _l1FeeOverhead, uint256 _l1FeeScalar, uint256 _validatorRewardScalar) returns()
func (_KromaL1Block *KromaL1BlockSession) SetL1BlockValues(_number uint64, _timestamp uint64, _basefee *big.Int, _hash [32]byte, _sequenceNumber uint64, _batcherHash [32]byte, _l1FeeOverhead *big.Int, _l1FeeScalar *big.Int, _validatorRewardScalar *big.Int) (*types.Transaction, error) {
	return _KromaL1Block.Contract.SetL1BlockValues(&_KromaL1Block.TransactOpts, _number, _timestamp, _basefee, _hash, _sequenceNumber, _batcherHash, _l1FeeOverhead, _l1FeeScalar, _validatorRewardScalar)
}

// SetL1BlockValues is a paid mutator transaction binding the contract method 0xefc674eb.
//
// Solidity: function setL1BlockValues(uint64 _number, uint64 _timestamp, uint256 _basefee, bytes32 _hash, uint64 _sequenceNumber, bytes32 _batcherHash, uint256 _l1FeeOverhead, uint256 _l1FeeScalar, uint256 _validatorRewardScalar) returns()
func (_KromaL1Block *KromaL1BlockTransactorSession) SetL1BlockValues(_number uint64, _timestamp uint64, _basefee *big.Int, _hash [32]byte, _sequenceNumber uint64, _batcherHash [32]byte, _l1FeeOverhead *big.Int, _l1FeeScalar *big.Int, _validatorRewardScalar *big.Int) (*types.Transaction, error) {
	return _KromaL1Block.Contract.SetL1BlockValues(&_KromaL1Block.TransactOpts, _number, _timestamp, _basefee, _hash, _sequenceNumber, _batcherHash, _l1FeeOverhead, _l1FeeScalar, _validatorRewardScalar)
}

// SetL1BlockValuesEcotone is a paid mutator transaction binding the contract method 0x440a5e20.
//
// Solidity: function setL1BlockValuesEcotone() returns()
func (_KromaL1Block *KromaL1BlockTransactor) SetL1BlockValuesEcotone(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KromaL1Block.contract.Transact(opts, "setL1BlockValuesEcotone")
}

// SetL1BlockValuesEcotone is a paid mutator transaction binding the contract method 0x440a5e20.
//
// Solidity: function setL1BlockValuesEcotone() returns()
func (_KromaL1Block *KromaL1BlockSession) SetL1BlockValuesEcotone() (*types.Transaction, error) {
	return _KromaL1Block.Contract.SetL1BlockValuesEcotone(&_KromaL1Block.TransactOpts)
}

// SetL1BlockValuesEcotone is a paid mutator transaction binding the contract method 0x440a5e20.
//
// Solidity: function setL1BlockValuesEcotone() returns()
func (_KromaL1Block *KromaL1BlockTransactorSession) SetL1BlockValuesEcotone() (*types.Transaction, error) {
	return _KromaL1Block.Contract.SetL1BlockValuesEcotone(&_KromaL1Block.TransactOpts)
}
