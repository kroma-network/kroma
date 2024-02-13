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

// ProtocolVaultMetaData contains all meta data concerning the ProtocolVault contract.
var ProtocolVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MIN_WITHDRAWAL_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RECIPIENT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalProcessed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawToL2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x61012060405234801561001157600080fd5b50604051610b0f380380610b0f8339810160408190526100309161005b565b600060808190526001600160a01b0390911660a052600160c081905260e0919091526101005261008b565b60006020828403121561006d57600080fd5b81516001600160a01b038116811461008457600080fd5b9392505050565b60805160a05160c05160e05161010051610a176100f860003960006103620152600061033901526000610310015260008181608701528181610191015281816102a7015281816103c4015281816104840152610623015260008181610157015261052b0152610a176000f3fe6080604052600436106100695760003560e01c80636ed39f62116100435780636ed39f621461010c57806384411d6514610121578063d3e5792b1461014557600080fd5b80630d9019e1146100755780633ccfd60b146100d357806354fd4d50146100ea57600080fd5b3661007057005b600080fd5b34801561008157600080fd5b506100a97f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156100df57600080fd5b506100e8610179565b005b3480156100f657600080fd5b506100ff610309565b6040516100ca91906108c8565b34801561011857600080fd5b506100e86103ac565b34801561012d57600080fd5b5061013760005481565b6040519081526020016100ca565b34801561015157600080fd5b506101377f000000000000000000000000000000000000000000000000000000000000000081565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610243576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f4665655661756c743a20746865206f6e6c7920726563697069656e742063616e60448201527f2063616c6c00000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b600061024d610527565b604080516020810182526000815290517fe11013dd0000000000000000000000000000000000000000000000000000000081529192507342000000000000000000000000000000000000099163e11013dd9184916102d4917f0000000000000000000000000000000000000000000000000000000000000000916188b891906004016108e2565b6000604051808303818588803b1580156102ed57600080fd5b505af1158015610301573d6000803e3d6000fd5b505050505050565b60606103347f0000000000000000000000000000000000000000000000000000000000000000610693565b61035d7f0000000000000000000000000000000000000000000000000000000000000000610693565b6103867f0000000000000000000000000000000000000000000000000000000000000000610693565b60405160200161039893929190610926565b604051602081830303815290604052905090565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610471576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f4665655661756c743a20746865206f6e6c7920726563697069656e742063616e60448201527f2063616c6c000000000000000000000000000000000000000000000000000000606482015260840161023a565b600061047b610527565b905060006104ba7f00000000000000000000000000000000000000000000000000000000000000005a8460405180602001604052806000815250610751565b905080610523576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f4665655661756c743a20455448207472616e73666572206661696c6564000000604482015260640161023a565b5050565b60007f00000000000000000000000000000000000000000000000000000000000000004710156105ff576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f4665655661756c743a207769746864726177616c20616d6f756e74206d75737460448201527f2062652067726561746572207468616e206d696e696d756d207769746864726160648201527f77616c20616d6f756e7400000000000000000000000000000000000000000000608482015260a40161023a565b600047905080600080828254610615919061099c565b9091555050604080518281527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166020820152338183015290517fc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba9181900360600190a1919050565b606060006106a08361076b565b600101905060008167ffffffffffffffff8111156106c0576106c06109db565b6040519080825280601f01601f1916602001820160405280156106ea576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a85049450846106f457509392505050565b600080600080845160208601878a8af19695505050505050565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f01000000000000000083106107b4577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef810000000083106107e0576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc1000083106107fe57662386f26fc10000830492506010015b6305f5e1008310610816576305f5e100830492506008015b612710831061082a57612710830492506004015b6064831061083c576064830492506002015b600a8310610848576001015b92915050565b60005b83811015610869578181015183820152602001610851565b83811115610878576000848401525b50505050565b6000815180845261089681602086016020860161084e565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006108db602083018461087e565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff8416815263ffffffff8316602082015260606040820152600061091d606083018461087e565b95945050505050565b6000845161093881846020890161084e565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551610974816001850160208a0161084e565b6001920191820152835161098f81600284016020880161084e565b0160020195945050505050565b600082198211156109d6577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea164736f6c634300080f000a",
}

// ProtocolVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use ProtocolVaultMetaData.ABI instead.
var ProtocolVaultABI = ProtocolVaultMetaData.ABI

// ProtocolVaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProtocolVaultMetaData.Bin instead.
var ProtocolVaultBin = ProtocolVaultMetaData.Bin

// DeployProtocolVault deploys a new Ethereum contract, binding an instance of ProtocolVault to it.
func DeployProtocolVault(auth *bind.TransactOpts, backend bind.ContractBackend, _recipient common.Address) (common.Address, *types.Transaction, *ProtocolVault, error) {
	parsed, err := ProtocolVaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProtocolVaultBin), backend, _recipient)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProtocolVault{ProtocolVaultCaller: ProtocolVaultCaller{contract: contract}, ProtocolVaultTransactor: ProtocolVaultTransactor{contract: contract}, ProtocolVaultFilterer: ProtocolVaultFilterer{contract: contract}}, nil
}

// ProtocolVault is an auto generated Go binding around an Ethereum contract.
type ProtocolVault struct {
	ProtocolVaultCaller     // Read-only binding to the contract
	ProtocolVaultTransactor // Write-only binding to the contract
	ProtocolVaultFilterer   // Log filterer for contract events
}

// ProtocolVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProtocolVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProtocolVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProtocolVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProtocolVaultSession struct {
	Contract     *ProtocolVault    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProtocolVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProtocolVaultCallerSession struct {
	Contract *ProtocolVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ProtocolVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProtocolVaultTransactorSession struct {
	Contract     *ProtocolVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ProtocolVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProtocolVaultRaw struct {
	Contract *ProtocolVault // Generic contract binding to access the raw methods on
}

// ProtocolVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProtocolVaultCallerRaw struct {
	Contract *ProtocolVaultCaller // Generic read-only contract binding to access the raw methods on
}

// ProtocolVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProtocolVaultTransactorRaw struct {
	Contract *ProtocolVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProtocolVault creates a new instance of ProtocolVault, bound to a specific deployed contract.
func NewProtocolVault(address common.Address, backend bind.ContractBackend) (*ProtocolVault, error) {
	contract, err := bindProtocolVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProtocolVault{ProtocolVaultCaller: ProtocolVaultCaller{contract: contract}, ProtocolVaultTransactor: ProtocolVaultTransactor{contract: contract}, ProtocolVaultFilterer: ProtocolVaultFilterer{contract: contract}}, nil
}

// NewProtocolVaultCaller creates a new read-only instance of ProtocolVault, bound to a specific deployed contract.
func NewProtocolVaultCaller(address common.Address, caller bind.ContractCaller) (*ProtocolVaultCaller, error) {
	contract, err := bindProtocolVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProtocolVaultCaller{contract: contract}, nil
}

// NewProtocolVaultTransactor creates a new write-only instance of ProtocolVault, bound to a specific deployed contract.
func NewProtocolVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*ProtocolVaultTransactor, error) {
	contract, err := bindProtocolVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProtocolVaultTransactor{contract: contract}, nil
}

// NewProtocolVaultFilterer creates a new log filterer instance of ProtocolVault, bound to a specific deployed contract.
func NewProtocolVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*ProtocolVaultFilterer, error) {
	contract, err := bindProtocolVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProtocolVaultFilterer{contract: contract}, nil
}

// bindProtocolVault binds a generic wrapper to an already deployed contract.
func bindProtocolVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProtocolVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProtocolVault *ProtocolVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProtocolVault.Contract.ProtocolVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProtocolVault *ProtocolVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtocolVault.Contract.ProtocolVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProtocolVault *ProtocolVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProtocolVault.Contract.ProtocolVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProtocolVault *ProtocolVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProtocolVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProtocolVault *ProtocolVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtocolVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProtocolVault *ProtocolVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProtocolVault.Contract.contract.Transact(opts, method, params...)
}

// MINWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0xd3e5792b.
//
// Solidity: function MIN_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_ProtocolVault *ProtocolVaultCaller) MINWITHDRAWALAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProtocolVault.contract.Call(opts, &out, "MIN_WITHDRAWAL_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0xd3e5792b.
//
// Solidity: function MIN_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_ProtocolVault *ProtocolVaultSession) MINWITHDRAWALAMOUNT() (*big.Int, error) {
	return _ProtocolVault.Contract.MINWITHDRAWALAMOUNT(&_ProtocolVault.CallOpts)
}

// MINWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0xd3e5792b.
//
// Solidity: function MIN_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_ProtocolVault *ProtocolVaultCallerSession) MINWITHDRAWALAMOUNT() (*big.Int, error) {
	return _ProtocolVault.Contract.MINWITHDRAWALAMOUNT(&_ProtocolVault.CallOpts)
}

// RECIPIENT is a free data retrieval call binding the contract method 0x0d9019e1.
//
// Solidity: function RECIPIENT() view returns(address)
func (_ProtocolVault *ProtocolVaultCaller) RECIPIENT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProtocolVault.contract.Call(opts, &out, "RECIPIENT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RECIPIENT is a free data retrieval call binding the contract method 0x0d9019e1.
//
// Solidity: function RECIPIENT() view returns(address)
func (_ProtocolVault *ProtocolVaultSession) RECIPIENT() (common.Address, error) {
	return _ProtocolVault.Contract.RECIPIENT(&_ProtocolVault.CallOpts)
}

// RECIPIENT is a free data retrieval call binding the contract method 0x0d9019e1.
//
// Solidity: function RECIPIENT() view returns(address)
func (_ProtocolVault *ProtocolVaultCallerSession) RECIPIENT() (common.Address, error) {
	return _ProtocolVault.Contract.RECIPIENT(&_ProtocolVault.CallOpts)
}

// TotalProcessed is a free data retrieval call binding the contract method 0x84411d65.
//
// Solidity: function totalProcessed() view returns(uint256)
func (_ProtocolVault *ProtocolVaultCaller) TotalProcessed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProtocolVault.contract.Call(opts, &out, "totalProcessed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalProcessed is a free data retrieval call binding the contract method 0x84411d65.
//
// Solidity: function totalProcessed() view returns(uint256)
func (_ProtocolVault *ProtocolVaultSession) TotalProcessed() (*big.Int, error) {
	return _ProtocolVault.Contract.TotalProcessed(&_ProtocolVault.CallOpts)
}

// TotalProcessed is a free data retrieval call binding the contract method 0x84411d65.
//
// Solidity: function totalProcessed() view returns(uint256)
func (_ProtocolVault *ProtocolVaultCallerSession) TotalProcessed() (*big.Int, error) {
	return _ProtocolVault.Contract.TotalProcessed(&_ProtocolVault.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ProtocolVault *ProtocolVaultCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ProtocolVault.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ProtocolVault *ProtocolVaultSession) Version() (string, error) {
	return _ProtocolVault.Contract.Version(&_ProtocolVault.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ProtocolVault *ProtocolVaultCallerSession) Version() (string, error) {
	return _ProtocolVault.Contract.Version(&_ProtocolVault.CallOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ProtocolVault *ProtocolVaultTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtocolVault.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ProtocolVault *ProtocolVaultSession) Withdraw() (*types.Transaction, error) {
	return _ProtocolVault.Contract.Withdraw(&_ProtocolVault.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ProtocolVault *ProtocolVaultTransactorSession) Withdraw() (*types.Transaction, error) {
	return _ProtocolVault.Contract.Withdraw(&_ProtocolVault.TransactOpts)
}

// WithdrawToL2 is a paid mutator transaction binding the contract method 0x6ed39f62.
//
// Solidity: function withdrawToL2() returns()
func (_ProtocolVault *ProtocolVaultTransactor) WithdrawToL2(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtocolVault.contract.Transact(opts, "withdrawToL2")
}

// WithdrawToL2 is a paid mutator transaction binding the contract method 0x6ed39f62.
//
// Solidity: function withdrawToL2() returns()
func (_ProtocolVault *ProtocolVaultSession) WithdrawToL2() (*types.Transaction, error) {
	return _ProtocolVault.Contract.WithdrawToL2(&_ProtocolVault.TransactOpts)
}

// WithdrawToL2 is a paid mutator transaction binding the contract method 0x6ed39f62.
//
// Solidity: function withdrawToL2() returns()
func (_ProtocolVault *ProtocolVaultTransactorSession) WithdrawToL2() (*types.Transaction, error) {
	return _ProtocolVault.Contract.WithdrawToL2(&_ProtocolVault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ProtocolVault *ProtocolVaultTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtocolVault.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ProtocolVault *ProtocolVaultSession) Receive() (*types.Transaction, error) {
	return _ProtocolVault.Contract.Receive(&_ProtocolVault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ProtocolVault *ProtocolVaultTransactorSession) Receive() (*types.Transaction, error) {
	return _ProtocolVault.Contract.Receive(&_ProtocolVault.TransactOpts)
}

// ProtocolVaultWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the ProtocolVault contract.
type ProtocolVaultWithdrawalIterator struct {
	Event *ProtocolVaultWithdrawal // Event containing the contract specifics and raw log

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
func (it *ProtocolVaultWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolVaultWithdrawal)
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
		it.Event = new(ProtocolVaultWithdrawal)
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
func (it *ProtocolVaultWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolVaultWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolVaultWithdrawal represents a Withdrawal event raised by the ProtocolVault contract.
type ProtocolVaultWithdrawal struct {
	Value *big.Int
	To    common.Address
	From  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0xc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba.
//
// Solidity: event Withdrawal(uint256 value, address to, address from)
func (_ProtocolVault *ProtocolVaultFilterer) FilterWithdrawal(opts *bind.FilterOpts) (*ProtocolVaultWithdrawalIterator, error) {

	logs, sub, err := _ProtocolVault.contract.FilterLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return &ProtocolVaultWithdrawalIterator{contract: _ProtocolVault.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0xc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba.
//
// Solidity: event Withdrawal(uint256 value, address to, address from)
func (_ProtocolVault *ProtocolVaultFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *ProtocolVaultWithdrawal) (event.Subscription, error) {

	logs, sub, err := _ProtocolVault.contract.WatchLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolVaultWithdrawal)
				if err := _ProtocolVault.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0xc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba.
//
// Solidity: event Withdrawal(uint256 value, address to, address from)
func (_ProtocolVault *ProtocolVaultFilterer) ParseWithdrawal(log types.Log) (*ProtocolVaultWithdrawal, error) {
	event := new(ProtocolVaultWithdrawal)
	if err := _ProtocolVault.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
