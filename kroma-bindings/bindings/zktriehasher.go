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

// ZKTrieHasherMetaData contains all meta data concerning the ZKTrieHasher contract.
var ZKTrieHasherMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_poseidon2\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"POSEIDON2\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPoseidon2\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161011138038061011183398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b608051608961008860003960006031015260896000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063dc8b503814602d575b600080fd5b60537f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f3fea164736f6c634300080f000a",
}

// ZKTrieHasherABI is the input ABI used to generate the binding from.
// Deprecated: Use ZKTrieHasherMetaData.ABI instead.
var ZKTrieHasherABI = ZKTrieHasherMetaData.ABI

// ZKTrieHasherBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZKTrieHasherMetaData.Bin instead.
var ZKTrieHasherBin = ZKTrieHasherMetaData.Bin

// DeployZKTrieHasher deploys a new Ethereum contract, binding an instance of ZKTrieHasher to it.
func DeployZKTrieHasher(auth *bind.TransactOpts, backend bind.ContractBackend, _poseidon2 common.Address) (common.Address, *types.Transaction, *ZKTrieHasher, error) {
	parsed, err := ZKTrieHasherMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZKTrieHasherBin), backend, _poseidon2)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZKTrieHasher{ZKTrieHasherCaller: ZKTrieHasherCaller{contract: contract}, ZKTrieHasherTransactor: ZKTrieHasherTransactor{contract: contract}, ZKTrieHasherFilterer: ZKTrieHasherFilterer{contract: contract}}, nil
}

// ZKTrieHasher is an auto generated Go binding around an Ethereum contract.
type ZKTrieHasher struct {
	ZKTrieHasherCaller     // Read-only binding to the contract
	ZKTrieHasherTransactor // Write-only binding to the contract
	ZKTrieHasherFilterer   // Log filterer for contract events
}

// ZKTrieHasherCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZKTrieHasherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKTrieHasherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZKTrieHasherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKTrieHasherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZKTrieHasherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKTrieHasherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZKTrieHasherSession struct {
	Contract     *ZKTrieHasher     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZKTrieHasherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZKTrieHasherCallerSession struct {
	Contract *ZKTrieHasherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ZKTrieHasherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZKTrieHasherTransactorSession struct {
	Contract     *ZKTrieHasherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ZKTrieHasherRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZKTrieHasherRaw struct {
	Contract *ZKTrieHasher // Generic contract binding to access the raw methods on
}

// ZKTrieHasherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZKTrieHasherCallerRaw struct {
	Contract *ZKTrieHasherCaller // Generic read-only contract binding to access the raw methods on
}

// ZKTrieHasherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZKTrieHasherTransactorRaw struct {
	Contract *ZKTrieHasherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZKTrieHasher creates a new instance of ZKTrieHasher, bound to a specific deployed contract.
func NewZKTrieHasher(address common.Address, backend bind.ContractBackend) (*ZKTrieHasher, error) {
	contract, err := bindZKTrieHasher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZKTrieHasher{ZKTrieHasherCaller: ZKTrieHasherCaller{contract: contract}, ZKTrieHasherTransactor: ZKTrieHasherTransactor{contract: contract}, ZKTrieHasherFilterer: ZKTrieHasherFilterer{contract: contract}}, nil
}

// NewZKTrieHasherCaller creates a new read-only instance of ZKTrieHasher, bound to a specific deployed contract.
func NewZKTrieHasherCaller(address common.Address, caller bind.ContractCaller) (*ZKTrieHasherCaller, error) {
	contract, err := bindZKTrieHasher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZKTrieHasherCaller{contract: contract}, nil
}

// NewZKTrieHasherTransactor creates a new write-only instance of ZKTrieHasher, bound to a specific deployed contract.
func NewZKTrieHasherTransactor(address common.Address, transactor bind.ContractTransactor) (*ZKTrieHasherTransactor, error) {
	contract, err := bindZKTrieHasher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZKTrieHasherTransactor{contract: contract}, nil
}

// NewZKTrieHasherFilterer creates a new log filterer instance of ZKTrieHasher, bound to a specific deployed contract.
func NewZKTrieHasherFilterer(address common.Address, filterer bind.ContractFilterer) (*ZKTrieHasherFilterer, error) {
	contract, err := bindZKTrieHasher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZKTrieHasherFilterer{contract: contract}, nil
}

// bindZKTrieHasher binds a generic wrapper to an already deployed contract.
func bindZKTrieHasher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZKTrieHasherMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKTrieHasher *ZKTrieHasherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKTrieHasher.Contract.ZKTrieHasherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKTrieHasher *ZKTrieHasherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKTrieHasher.Contract.ZKTrieHasherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKTrieHasher *ZKTrieHasherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKTrieHasher.Contract.ZKTrieHasherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKTrieHasher *ZKTrieHasherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKTrieHasher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKTrieHasher *ZKTrieHasherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKTrieHasher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKTrieHasher *ZKTrieHasherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKTrieHasher.Contract.contract.Transact(opts, method, params...)
}

// POSEIDON2 is a free data retrieval call binding the contract method 0xdc8b5038.
//
// Solidity: function POSEIDON2() view returns(address)
func (_ZKTrieHasher *ZKTrieHasherCaller) POSEIDON2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZKTrieHasher.contract.Call(opts, &out, "POSEIDON2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// POSEIDON2 is a free data retrieval call binding the contract method 0xdc8b5038.
//
// Solidity: function POSEIDON2() view returns(address)
func (_ZKTrieHasher *ZKTrieHasherSession) POSEIDON2() (common.Address, error) {
	return _ZKTrieHasher.Contract.POSEIDON2(&_ZKTrieHasher.CallOpts)
}

// POSEIDON2 is a free data retrieval call binding the contract method 0xdc8b5038.
//
// Solidity: function POSEIDON2() view returns(address)
func (_ZKTrieHasher *ZKTrieHasherCallerSession) POSEIDON2() (common.Address, error) {
	return _ZKTrieHasher.Contract.POSEIDON2(&_ZKTrieHasher.CallOpts)
}
