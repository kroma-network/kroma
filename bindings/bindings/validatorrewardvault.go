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

// ValidatorRewardVaultMetaData contains all meta data concerning the ValidatorRewardVault contract.
var ValidatorRewardVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardDivider\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2BlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Rewarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MIN_WITHDRAWAL_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RECIPIENT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_DIVIDER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_POOL\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_penaltyNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_penaltyDenom\",\"type\":\"uint256\"}],\"name\":\"reward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalProcessed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalReserved\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x61016060405234801561001157600080fd5b50604051610ef4380380610ef48339810160408190526100309161006a565b678ac7230489e80000608052600060a081905260c0819052600160e052610100526001600160a01b039190911661012052610140526100a4565b6000806040838503121561007d57600080fd5b82516001600160a01b038116811461009457600080fd5b6020939093015192949293505050565b60805160a05160c05160e051610100516101205161014051610de26101126000396000818161016a0152610a2a0152600081816102050152610530015260006104cf015260006104a60152600061047d0152600060d301526000818161026f01526102a30152610de26000f3fe6080604052600436106100b55760003560e01c806384411d6511610069578063c5a3487c1161004e578063c5a3487c14610227578063c71b0e1c14610247578063d3e5792b1461025d57600080fd5b806384411d65146101dd578063b98debbf146101f357600080fd5b806354fd4d501161009a57806354fd4d501461013657806362aba76b1461015857806370a082311461019a57600080fd5b80630d9019e1146100c15780633ccfd60b1461011f57600080fd5b366100bc57005b600080fd5b3480156100cd57600080fd5b506100f57f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561012b57600080fd5b50610134610291565b005b34801561014257600080fd5b5061014b610476565b6040516101169190610afc565b34801561016457600080fd5b5061018c7f000000000000000000000000000000000000000000000000000000000000000081565b604051908152602001610116565b3480156101a657600080fd5b5061018c6101b5366004610b3f565b73ffffffffffffffffffffffffffffffffffffffff1660009081526001602052604090205490565b3480156101e957600080fd5b5061018c60005481565b3480156101ff57600080fd5b506100f57f000000000000000000000000000000000000000000000000000000000000000081565b34801561023357600080fd5b50610134610242366004610b5a565b610519565b34801561025357600080fd5b5061018c60035481565b34801561026957600080fd5b5061018c7f000000000000000000000000000000000000000000000000000000000000000081565b336000908152600160205260409020547f000000000000000000000000000000000000000000000000000000000000000081101561037c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605660248201527f56616c696461746f725265776172645661756c743a207769746864726177616c60448201527f20616d6f756e74206d7573742062652067726561746572207468616e206d696e60648201527f696d756d207769746864726177616c20616d6f756e7400000000000000000000608482015260a4015b60405180910390fd5b33600081815260016020908152604080832083905560038054869003905582548501909255815184815290810183905280820192909252517fc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba9181900360600190a1604080516020810182526000815290517fe11013dd0000000000000000000000000000000000000000000000000000000081527342000000000000000000000000000000000000099163e11013dd9184916104419133916188b891600401610b93565b6000604051808303818588803b15801561045a57600080fd5b505af115801561046e573d6000803e3d6000fd5b505050505050565b60606104a17f0000000000000000000000000000000000000000000000000000000000000000610857565b6104ca7f0000000000000000000000000000000000000000000000000000000000000000610857565b6104f37f0000000000000000000000000000000000000000000000000000000000000000610857565b60405160200161050593929190610bd7565b604051602081830303815290604052905090565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000167fffffffffffffffffffffffffeeeeffffffffffffffffffffffffffffffffeeef330173ffffffffffffffffffffffffffffffffffffffff161461063c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604860248201527f56616c696461746f725265776172645661756c743a2066756e6374696f6e206360448201527f616e206f6e6c792062652063616c6c65642066726f6d207468652056616c696460648201527f61746f72506f6f6c000000000000000000000000000000000000000000000000608482015260a401610373565b73ffffffffffffffffffffffffffffffffffffffff84166106df576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603360248201527f56616c696461746f725265776172645661756c743a2076616c696461746f722060448201527f616464726573732063616e6e6f742062652030000000000000000000000000006064820152608401610373565b60008381526002602052604090205460ff16156107a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604e60248201527f56616c696461746f725265776172645661756c743a207468652072657761726460448201527f2068617320616c7265616479206265656e207061696420666f7220746865204c60648201527f3220626c6f636b206e756d626572000000000000000000000000000000000000608482015260a401610373565b60006107b08383610994565b600380548201905573ffffffffffffffffffffffffffffffffffffffff86166000818152600160208181526040808420805487019055898452600282529283902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016909217909155905183815292935086927f291e8ba3c0f4b0bd86e6e2346fcee1e7ca0975b1cc1886bfbc722d34f3f24d91910160405180910390a35050505050565b60608160000361089a57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156108c457806108ae81610c7c565b91506108bd9050600a83610ce3565b915061089e565b60008167ffffffffffffffff8111156108df576108df610cf7565b6040519080825280601f01601f191660200182016040528015610909576020820181803683370190505b5090505b841561098c5761091e600183610d26565b915061092b600a86610d3d565b610936906030610d51565b60f81b81838151811061094b5761094b610d69565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610985600a86610ce3565b945061090d565b949350505050565b600081600003610a26576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f56616c696461746f725265776172645661756c743a2064656e6f6d696e61746f60448201527f722063616e6e6f742062652030000000000000000000000000000000000000006064820152608401610373565b60007f000000000000000000000000000000000000000000000000000000000000000060035447610a579190610d26565b610a619190610ce3565b905082610a6e8582610d26565b610a789083610d98565b61098c9190610ce3565b60005b83811015610a9d578181015183820152602001610a85565b83811115610aac576000848401525b50505050565b60008151808452610aca816020860160208601610a82565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610b0f6020830184610ab2565b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610b3a57600080fd5b919050565b600060208284031215610b5157600080fd5b610b0f82610b16565b60008060008060808587031215610b7057600080fd5b610b7985610b16565b966020860135965060408601359560600135945092505050565b73ffffffffffffffffffffffffffffffffffffffff8416815263ffffffff83166020820152606060408201526000610bce6060830184610ab2565b95945050505050565b60008451610be9818460208901610a82565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551610c25816001850160208a01610a82565b60019201918201528351610c40816002840160208801610a82565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610cad57610cad610c4d565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082610cf257610cf2610cb4565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082821015610d3857610d38610c4d565b500390565b600082610d4c57610d4c610cb4565b500690565b60008219821115610d6457610d64610c4d565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615610dd057610dd0610c4d565b50029056fea164736f6c634300080f000a",
}

// ValidatorRewardVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorRewardVaultMetaData.ABI instead.
var ValidatorRewardVaultABI = ValidatorRewardVaultMetaData.ABI

// ValidatorRewardVaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValidatorRewardVaultMetaData.Bin instead.
var ValidatorRewardVaultBin = ValidatorRewardVaultMetaData.Bin

// DeployValidatorRewardVault deploys a new Ethereum contract, binding an instance of ValidatorRewardVault to it.
func DeployValidatorRewardVault(auth *bind.TransactOpts, backend bind.ContractBackend, _validatorPool common.Address, _rewardDivider *big.Int) (common.Address, *types.Transaction, *ValidatorRewardVault, error) {
	parsed, err := ValidatorRewardVaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValidatorRewardVaultBin), backend, _validatorPool, _rewardDivider)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidatorRewardVault{ValidatorRewardVaultCaller: ValidatorRewardVaultCaller{contract: contract}, ValidatorRewardVaultTransactor: ValidatorRewardVaultTransactor{contract: contract}, ValidatorRewardVaultFilterer: ValidatorRewardVaultFilterer{contract: contract}}, nil
}

// ValidatorRewardVault is an auto generated Go binding around an Ethereum contract.
type ValidatorRewardVault struct {
	ValidatorRewardVaultCaller     // Read-only binding to the contract
	ValidatorRewardVaultTransactor // Write-only binding to the contract
	ValidatorRewardVaultFilterer   // Log filterer for contract events
}

// ValidatorRewardVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorRewardVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorRewardVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorRewardVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorRewardVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorRewardVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorRewardVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorRewardVaultSession struct {
	Contract     *ValidatorRewardVault // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ValidatorRewardVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorRewardVaultCallerSession struct {
	Contract *ValidatorRewardVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ValidatorRewardVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorRewardVaultTransactorSession struct {
	Contract     *ValidatorRewardVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ValidatorRewardVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorRewardVaultRaw struct {
	Contract *ValidatorRewardVault // Generic contract binding to access the raw methods on
}

// ValidatorRewardVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorRewardVaultCallerRaw struct {
	Contract *ValidatorRewardVaultCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorRewardVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorRewardVaultTransactorRaw struct {
	Contract *ValidatorRewardVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorRewardVault creates a new instance of ValidatorRewardVault, bound to a specific deployed contract.
func NewValidatorRewardVault(address common.Address, backend bind.ContractBackend) (*ValidatorRewardVault, error) {
	contract, err := bindValidatorRewardVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorRewardVault{ValidatorRewardVaultCaller: ValidatorRewardVaultCaller{contract: contract}, ValidatorRewardVaultTransactor: ValidatorRewardVaultTransactor{contract: contract}, ValidatorRewardVaultFilterer: ValidatorRewardVaultFilterer{contract: contract}}, nil
}

// NewValidatorRewardVaultCaller creates a new read-only instance of ValidatorRewardVault, bound to a specific deployed contract.
func NewValidatorRewardVaultCaller(address common.Address, caller bind.ContractCaller) (*ValidatorRewardVaultCaller, error) {
	contract, err := bindValidatorRewardVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorRewardVaultCaller{contract: contract}, nil
}

// NewValidatorRewardVaultTransactor creates a new write-only instance of ValidatorRewardVault, bound to a specific deployed contract.
func NewValidatorRewardVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorRewardVaultTransactor, error) {
	contract, err := bindValidatorRewardVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorRewardVaultTransactor{contract: contract}, nil
}

// NewValidatorRewardVaultFilterer creates a new log filterer instance of ValidatorRewardVault, bound to a specific deployed contract.
func NewValidatorRewardVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorRewardVaultFilterer, error) {
	contract, err := bindValidatorRewardVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorRewardVaultFilterer{contract: contract}, nil
}

// bindValidatorRewardVault binds a generic wrapper to an already deployed contract.
func bindValidatorRewardVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValidatorRewardVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorRewardVault *ValidatorRewardVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorRewardVault.Contract.ValidatorRewardVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorRewardVault *ValidatorRewardVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorRewardVault.Contract.ValidatorRewardVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorRewardVault *ValidatorRewardVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorRewardVault.Contract.ValidatorRewardVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorRewardVault *ValidatorRewardVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorRewardVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorRewardVault *ValidatorRewardVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorRewardVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorRewardVault *ValidatorRewardVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorRewardVault.Contract.contract.Transact(opts, method, params...)
}

// MINWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0xd3e5792b.
//
// Solidity: function MIN_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_ValidatorRewardVault *ValidatorRewardVaultCaller) MINWITHDRAWALAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorRewardVault.contract.Call(opts, &out, "MIN_WITHDRAWAL_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0xd3e5792b.
//
// Solidity: function MIN_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_ValidatorRewardVault *ValidatorRewardVaultSession) MINWITHDRAWALAMOUNT() (*big.Int, error) {
	return _ValidatorRewardVault.Contract.MINWITHDRAWALAMOUNT(&_ValidatorRewardVault.CallOpts)
}

// MINWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0xd3e5792b.
//
// Solidity: function MIN_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_ValidatorRewardVault *ValidatorRewardVaultCallerSession) MINWITHDRAWALAMOUNT() (*big.Int, error) {
	return _ValidatorRewardVault.Contract.MINWITHDRAWALAMOUNT(&_ValidatorRewardVault.CallOpts)
}

// RECIPIENT is a free data retrieval call binding the contract method 0x0d9019e1.
//
// Solidity: function RECIPIENT() view returns(address)
func (_ValidatorRewardVault *ValidatorRewardVaultCaller) RECIPIENT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorRewardVault.contract.Call(opts, &out, "RECIPIENT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RECIPIENT is a free data retrieval call binding the contract method 0x0d9019e1.
//
// Solidity: function RECIPIENT() view returns(address)
func (_ValidatorRewardVault *ValidatorRewardVaultSession) RECIPIENT() (common.Address, error) {
	return _ValidatorRewardVault.Contract.RECIPIENT(&_ValidatorRewardVault.CallOpts)
}

// RECIPIENT is a free data retrieval call binding the contract method 0x0d9019e1.
//
// Solidity: function RECIPIENT() view returns(address)
func (_ValidatorRewardVault *ValidatorRewardVaultCallerSession) RECIPIENT() (common.Address, error) {
	return _ValidatorRewardVault.Contract.RECIPIENT(&_ValidatorRewardVault.CallOpts)
}

// REWARDDIVIDER is a free data retrieval call binding the contract method 0x62aba76b.
//
// Solidity: function REWARD_DIVIDER() view returns(uint256)
func (_ValidatorRewardVault *ValidatorRewardVaultCaller) REWARDDIVIDER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorRewardVault.contract.Call(opts, &out, "REWARD_DIVIDER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REWARDDIVIDER is a free data retrieval call binding the contract method 0x62aba76b.
//
// Solidity: function REWARD_DIVIDER() view returns(uint256)
func (_ValidatorRewardVault *ValidatorRewardVaultSession) REWARDDIVIDER() (*big.Int, error) {
	return _ValidatorRewardVault.Contract.REWARDDIVIDER(&_ValidatorRewardVault.CallOpts)
}

// REWARDDIVIDER is a free data retrieval call binding the contract method 0x62aba76b.
//
// Solidity: function REWARD_DIVIDER() view returns(uint256)
func (_ValidatorRewardVault *ValidatorRewardVaultCallerSession) REWARDDIVIDER() (*big.Int, error) {
	return _ValidatorRewardVault.Contract.REWARDDIVIDER(&_ValidatorRewardVault.CallOpts)
}

// VALIDATORPOOL is a free data retrieval call binding the contract method 0xb98debbf.
//
// Solidity: function VALIDATOR_POOL() view returns(address)
func (_ValidatorRewardVault *ValidatorRewardVaultCaller) VALIDATORPOOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorRewardVault.contract.Call(opts, &out, "VALIDATOR_POOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORPOOL is a free data retrieval call binding the contract method 0xb98debbf.
//
// Solidity: function VALIDATOR_POOL() view returns(address)
func (_ValidatorRewardVault *ValidatorRewardVaultSession) VALIDATORPOOL() (common.Address, error) {
	return _ValidatorRewardVault.Contract.VALIDATORPOOL(&_ValidatorRewardVault.CallOpts)
}

// VALIDATORPOOL is a free data retrieval call binding the contract method 0xb98debbf.
//
// Solidity: function VALIDATOR_POOL() view returns(address)
func (_ValidatorRewardVault *ValidatorRewardVaultCallerSession) VALIDATORPOOL() (common.Address, error) {
	return _ValidatorRewardVault.Contract.VALIDATORPOOL(&_ValidatorRewardVault.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _addr) view returns(uint256)
func (_ValidatorRewardVault *ValidatorRewardVaultCaller) BalanceOf(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorRewardVault.contract.Call(opts, &out, "balanceOf", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _addr) view returns(uint256)
func (_ValidatorRewardVault *ValidatorRewardVaultSession) BalanceOf(_addr common.Address) (*big.Int, error) {
	return _ValidatorRewardVault.Contract.BalanceOf(&_ValidatorRewardVault.CallOpts, _addr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _addr) view returns(uint256)
func (_ValidatorRewardVault *ValidatorRewardVaultCallerSession) BalanceOf(_addr common.Address) (*big.Int, error) {
	return _ValidatorRewardVault.Contract.BalanceOf(&_ValidatorRewardVault.CallOpts, _addr)
}

// TotalProcessed is a free data retrieval call binding the contract method 0x84411d65.
//
// Solidity: function totalProcessed() view returns(uint256)
func (_ValidatorRewardVault *ValidatorRewardVaultCaller) TotalProcessed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorRewardVault.contract.Call(opts, &out, "totalProcessed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalProcessed is a free data retrieval call binding the contract method 0x84411d65.
//
// Solidity: function totalProcessed() view returns(uint256)
func (_ValidatorRewardVault *ValidatorRewardVaultSession) TotalProcessed() (*big.Int, error) {
	return _ValidatorRewardVault.Contract.TotalProcessed(&_ValidatorRewardVault.CallOpts)
}

// TotalProcessed is a free data retrieval call binding the contract method 0x84411d65.
//
// Solidity: function totalProcessed() view returns(uint256)
func (_ValidatorRewardVault *ValidatorRewardVaultCallerSession) TotalProcessed() (*big.Int, error) {
	return _ValidatorRewardVault.Contract.TotalProcessed(&_ValidatorRewardVault.CallOpts)
}

// TotalReserved is a free data retrieval call binding the contract method 0xc71b0e1c.
//
// Solidity: function totalReserved() view returns(uint256)
func (_ValidatorRewardVault *ValidatorRewardVaultCaller) TotalReserved(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorRewardVault.contract.Call(opts, &out, "totalReserved")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReserved is a free data retrieval call binding the contract method 0xc71b0e1c.
//
// Solidity: function totalReserved() view returns(uint256)
func (_ValidatorRewardVault *ValidatorRewardVaultSession) TotalReserved() (*big.Int, error) {
	return _ValidatorRewardVault.Contract.TotalReserved(&_ValidatorRewardVault.CallOpts)
}

// TotalReserved is a free data retrieval call binding the contract method 0xc71b0e1c.
//
// Solidity: function totalReserved() view returns(uint256)
func (_ValidatorRewardVault *ValidatorRewardVaultCallerSession) TotalReserved() (*big.Int, error) {
	return _ValidatorRewardVault.Contract.TotalReserved(&_ValidatorRewardVault.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ValidatorRewardVault *ValidatorRewardVaultCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ValidatorRewardVault.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ValidatorRewardVault *ValidatorRewardVaultSession) Version() (string, error) {
	return _ValidatorRewardVault.Contract.Version(&_ValidatorRewardVault.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ValidatorRewardVault *ValidatorRewardVaultCallerSession) Version() (string, error) {
	return _ValidatorRewardVault.Contract.Version(&_ValidatorRewardVault.CallOpts)
}

// Reward is a paid mutator transaction binding the contract method 0xc5a3487c.
//
// Solidity: function reward(address _validator, uint256 _l2BlockNumber, uint256 _penaltyNum, uint256 _penaltyDenom) returns()
func (_ValidatorRewardVault *ValidatorRewardVaultTransactor) Reward(opts *bind.TransactOpts, _validator common.Address, _l2BlockNumber *big.Int, _penaltyNum *big.Int, _penaltyDenom *big.Int) (*types.Transaction, error) {
	return _ValidatorRewardVault.contract.Transact(opts, "reward", _validator, _l2BlockNumber, _penaltyNum, _penaltyDenom)
}

// Reward is a paid mutator transaction binding the contract method 0xc5a3487c.
//
// Solidity: function reward(address _validator, uint256 _l2BlockNumber, uint256 _penaltyNum, uint256 _penaltyDenom) returns()
func (_ValidatorRewardVault *ValidatorRewardVaultSession) Reward(_validator common.Address, _l2BlockNumber *big.Int, _penaltyNum *big.Int, _penaltyDenom *big.Int) (*types.Transaction, error) {
	return _ValidatorRewardVault.Contract.Reward(&_ValidatorRewardVault.TransactOpts, _validator, _l2BlockNumber, _penaltyNum, _penaltyDenom)
}

// Reward is a paid mutator transaction binding the contract method 0xc5a3487c.
//
// Solidity: function reward(address _validator, uint256 _l2BlockNumber, uint256 _penaltyNum, uint256 _penaltyDenom) returns()
func (_ValidatorRewardVault *ValidatorRewardVaultTransactorSession) Reward(_validator common.Address, _l2BlockNumber *big.Int, _penaltyNum *big.Int, _penaltyDenom *big.Int) (*types.Transaction, error) {
	return _ValidatorRewardVault.Contract.Reward(&_ValidatorRewardVault.TransactOpts, _validator, _l2BlockNumber, _penaltyNum, _penaltyDenom)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ValidatorRewardVault *ValidatorRewardVaultTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorRewardVault.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ValidatorRewardVault *ValidatorRewardVaultSession) Withdraw() (*types.Transaction, error) {
	return _ValidatorRewardVault.Contract.Withdraw(&_ValidatorRewardVault.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ValidatorRewardVault *ValidatorRewardVaultTransactorSession) Withdraw() (*types.Transaction, error) {
	return _ValidatorRewardVault.Contract.Withdraw(&_ValidatorRewardVault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ValidatorRewardVault *ValidatorRewardVaultTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorRewardVault.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ValidatorRewardVault *ValidatorRewardVaultSession) Receive() (*types.Transaction, error) {
	return _ValidatorRewardVault.Contract.Receive(&_ValidatorRewardVault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ValidatorRewardVault *ValidatorRewardVaultTransactorSession) Receive() (*types.Transaction, error) {
	return _ValidatorRewardVault.Contract.Receive(&_ValidatorRewardVault.TransactOpts)
}

// ValidatorRewardVaultRewardedIterator is returned from FilterRewarded and is used to iterate over the raw logs and unpacked data for Rewarded events raised by the ValidatorRewardVault contract.
type ValidatorRewardVaultRewardedIterator struct {
	Event *ValidatorRewardVaultRewarded // Event containing the contract specifics and raw log

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
func (it *ValidatorRewardVaultRewardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorRewardVaultRewarded)
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
		it.Event = new(ValidatorRewardVaultRewarded)
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
func (it *ValidatorRewardVaultRewardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorRewardVaultRewardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorRewardVaultRewarded represents a Rewarded event raised by the ValidatorRewardVault contract.
type ValidatorRewardVaultRewarded struct {
	Validator     common.Address
	L2BlockNumber *big.Int
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRewarded is a free log retrieval operation binding the contract event 0x291e8ba3c0f4b0bd86e6e2346fcee1e7ca0975b1cc1886bfbc722d34f3f24d91.
//
// Solidity: event Rewarded(address indexed validator, uint256 indexed l2BlockNumber, uint256 amount)
func (_ValidatorRewardVault *ValidatorRewardVaultFilterer) FilterRewarded(opts *bind.FilterOpts, validator []common.Address, l2BlockNumber []*big.Int) (*ValidatorRewardVaultRewardedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var l2BlockNumberRule []interface{}
	for _, l2BlockNumberItem := range l2BlockNumber {
		l2BlockNumberRule = append(l2BlockNumberRule, l2BlockNumberItem)
	}

	logs, sub, err := _ValidatorRewardVault.contract.FilterLogs(opts, "Rewarded", validatorRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorRewardVaultRewardedIterator{contract: _ValidatorRewardVault.contract, event: "Rewarded", logs: logs, sub: sub}, nil
}

// WatchRewarded is a free log subscription operation binding the contract event 0x291e8ba3c0f4b0bd86e6e2346fcee1e7ca0975b1cc1886bfbc722d34f3f24d91.
//
// Solidity: event Rewarded(address indexed validator, uint256 indexed l2BlockNumber, uint256 amount)
func (_ValidatorRewardVault *ValidatorRewardVaultFilterer) WatchRewarded(opts *bind.WatchOpts, sink chan<- *ValidatorRewardVaultRewarded, validator []common.Address, l2BlockNumber []*big.Int) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var l2BlockNumberRule []interface{}
	for _, l2BlockNumberItem := range l2BlockNumber {
		l2BlockNumberRule = append(l2BlockNumberRule, l2BlockNumberItem)
	}

	logs, sub, err := _ValidatorRewardVault.contract.WatchLogs(opts, "Rewarded", validatorRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorRewardVaultRewarded)
				if err := _ValidatorRewardVault.contract.UnpackLog(event, "Rewarded", log); err != nil {
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

// ParseRewarded is a log parse operation binding the contract event 0x291e8ba3c0f4b0bd86e6e2346fcee1e7ca0975b1cc1886bfbc722d34f3f24d91.
//
// Solidity: event Rewarded(address indexed validator, uint256 indexed l2BlockNumber, uint256 amount)
func (_ValidatorRewardVault *ValidatorRewardVaultFilterer) ParseRewarded(log types.Log) (*ValidatorRewardVaultRewarded, error) {
	event := new(ValidatorRewardVaultRewarded)
	if err := _ValidatorRewardVault.contract.UnpackLog(event, "Rewarded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorRewardVaultWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the ValidatorRewardVault contract.
type ValidatorRewardVaultWithdrawalIterator struct {
	Event *ValidatorRewardVaultWithdrawal // Event containing the contract specifics and raw log

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
func (it *ValidatorRewardVaultWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorRewardVaultWithdrawal)
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
		it.Event = new(ValidatorRewardVaultWithdrawal)
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
func (it *ValidatorRewardVaultWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorRewardVaultWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorRewardVaultWithdrawal represents a Withdrawal event raised by the ValidatorRewardVault contract.
type ValidatorRewardVaultWithdrawal struct {
	Value *big.Int
	To    common.Address
	From  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0xc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba.
//
// Solidity: event Withdrawal(uint256 value, address to, address from)
func (_ValidatorRewardVault *ValidatorRewardVaultFilterer) FilterWithdrawal(opts *bind.FilterOpts) (*ValidatorRewardVaultWithdrawalIterator, error) {

	logs, sub, err := _ValidatorRewardVault.contract.FilterLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return &ValidatorRewardVaultWithdrawalIterator{contract: _ValidatorRewardVault.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0xc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba.
//
// Solidity: event Withdrawal(uint256 value, address to, address from)
func (_ValidatorRewardVault *ValidatorRewardVaultFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *ValidatorRewardVaultWithdrawal) (event.Subscription, error) {

	logs, sub, err := _ValidatorRewardVault.contract.WatchLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorRewardVaultWithdrawal)
				if err := _ValidatorRewardVault.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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
func (_ValidatorRewardVault *ValidatorRewardVaultFilterer) ParseWithdrawal(log types.Log) (*ValidatorRewardVaultWithdrawal, error) {
	event := new(ValidatorRewardVaultWithdrawal)
	if err := _ValidatorRewardVault.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
