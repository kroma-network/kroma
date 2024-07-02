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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_governanceToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_recipients\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"GOVERNANCE_TOKEN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractGovernanceToken\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MINT_CAP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SHARE_DENOMINATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnershipOfToken\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"distribute\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"minted\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recipients\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnershipOfToken\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shareOf\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnershipOfToken\",\"inputs\":[{\"name\":\"newMintManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x60a06040523480156200001157600080fd5b506040516200169338038062001693833981016040819052620000349162000574565b6200003f336200032a565b6001600160a01b038416608052620000578362000354565b8051825114620000ba5760405162461bcd60e51b8152602060048201526024808201527f4d696e744d616e616765723a20696e76616c6964206c656e677468206f6620616044820152637272617960e01b60648201526084015b60405180910390fd5b6000805b835181101562000295576000848281518110620000df57620000df6200066a565b6020026020010151905060006001600160a01b0316816001600160a01b031603620001605760405162461bcd60e51b815260206004820152602a60248201527f4d696e744d616e616765723a20726563697069656e74206164647265737320636044820152690616e6e6f7420626520360b41b6064820152608401620000b1565b60008483815181106200017757620001776200066a565b6020026020010151905080600003620001d35760405162461bcd60e51b815260206004820152601e60248201527f4d696e744d616e616765723a2073686172652063616e6e6f74206265203000006044820152606401620000b1565b6001600160a01b03821660009081526003602052604081205490036200023f57600280546001810182556000919091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0180546001600160a01b0319166001600160a01b0384161790555b6001600160a01b038216600090815260036020526040812080548392906200026990849062000696565b909155506200027b9050818562000696565b9350505080806200028c90620006b1565b915050620000be565b50620186a08111156200031f5760405162461bcd60e51b8152602060048201526044602482018190527f4d696e744d616e616765723a206d617820746f74616c20736861726520697320908201527f657175616c206f72206c657373207468616e2053484152455f44454e4f4d494e60648201526320aa27a960e11b608482015260a401620000b1565b5050505050620006cd565b600180546001600160a01b03191690556200035181620003c7602090811b62000b5317901c565b50565b6200035e62000417565b600180546001600160a01b0319166001600160a01b0383169081179091556200038f6000546001600160a01b031690565b6001600160a01b03167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000546001600160a01b03163314620004735760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401620000b1565b565b80516001600160a01b03811681146200048d57600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715620004d357620004d362000492565b604052919050565b60006001600160401b03821115620004f757620004f762000492565b5060051b60200190565b600082601f8301126200051357600080fd5b815160206200052c6200052683620004db565b620004a8565b82815260059290921b840181019181810190868411156200054c57600080fd5b8286015b8481101562000569578051835291830191830162000550565b509695505050505050565b600080600080608085870312156200058b57600080fd5b620005968562000475565b93506020620005a781870162000475565b60408701519094506001600160401b0380821115620005c557600080fd5b818801915088601f830112620005da57600080fd5b8151620005eb6200052682620004db565b81815260059190911b8301840190848101908b8311156200060b57600080fd5b938501935b828510156200063457620006248562000475565b8252938501939085019062000610565b60608b015190975094505050808311156200064e57600080fd5b50506200065e8782880162000501565b91505092959194509250565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60008219821115620006ac57620006ac62000680565b500190565b600060018201620006c657620006c662000680565b5060010190565b608051610f796200071a6000396000818161014f01528181610331015281816104a3015281816105610152818161062e0152818161080a015281816108b30152610a190152610f796000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80637eb1184511610097578063d1bc76a111610066578063d1bc76a114610229578063e30c39781461023c578063e4fc6b6d1461025a578063f2fde38b1461026257600080fd5b80637eb11845146101ee5780638da5cb5b146101f857806398f1312e14610216578063baee5ed41461022157600080fd5b8063457c3977116100d3578063457c3977146101965780634f02c420146101a9578063715018a6146101de57806379ba5097146101e657600080fd5b80631249c58b1461010557806321e5e2c41461010f57806328df3f18146101425780632efd46d61461014a575b600080fd5b61010d610275565b005b61012f61011d366004610c7a565b60036020526000908152604090205481565b6040519081526020015b60405180910390f35b61010d610557565b6101717f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610139565b61010d6101a4366004610c7a565b6105e1565b6001546101ce9074010000000000000000000000000000000000000000900460ff1681565b6040519015158152602001610139565b61010d61068d565b61010d6106a1565b61012f620186a081565b60005473ffffffffffffffffffffffffffffffffffffffff16610171565b61012f633b9aca0081565b61010d610756565b610171610237366004610cb7565b610870565b60015473ffffffffffffffffffffffffffffffffffffffff16610171565b61010d6108a7565b61010d610270366004610c7a565b610aa3565b61027d610bc8565b60015474010000000000000000000000000000000000000000900460ff161561032d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f4d696e744d616e616765723a20616c7265616479206d696e746564206f6e207460448201527f68697320636861696e000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa15801561039a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103be9190610cd0565b6103c990600a610e44565b6103d790633b9aca00610e53565b90506000805b60025481101561046d576000600282815481106103fc576103fc610e90565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff168083526003909152604082205490925090620186a061043f8388610e53565b6104499190610ebf565b90506104558186610efa565b9450505050808061046590610f12565b9150506103dd565b506040517f40c10f19000000000000000000000000000000000000000000000000000000008152306004820152602481018290527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906340c10f1990604401600060405180830381600087803b1580156104fc57600080fd5b505af1158015610510573d6000803e3d6000fd5b5050600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000017905550505050565b61055f610bc8565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166379ba50976040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156105c757600080fd5b505af11580156105db573d6000803e3d6000fd5b50505050565b6105e9610bc8565b6040517ff2fde38b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063f2fde38b90602401600060405180830381600087803b15801561067257600080fd5b505af1158015610686573d6000803e3d6000fd5b5050505050565b610695610bc8565b61069f6000610c49565b565b600154339073ffffffffffffffffffffffffffffffffffffffff16811461074a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f4f776e61626c6532537465703a2063616c6c6572206973206e6f74207468652060448201527f6e6577206f776e657200000000000000000000000000000000000000000000006064820152608401610324565b61075381610c49565b50565b61075e610bc8565b60015474010000000000000000000000000000000000000000900460ff16610808576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f4d696e744d616e616765723a206e6f74206d696e746564206265666f7265207260448201527f656e6f756e6365206f776e6572736869700000000000000000000000000000006064820152608401610324565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663715018a66040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156105c757600080fd5b6002818154811061088057600080fd5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b6108af610bc8565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa15801561091c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109409190610cd0565b61094b90600a610e44565b61095990633b9aca00610e53565b905060005b600254811015610a9f5760006002828154811061097d5761097d610e90565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff168083526003909152604082205490925090620186a06109c08387610e53565b6109ca9190610ebf565b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018390529192507f00000000000000000000000000000000000000000000000000000000000000009091169063a9059cbb906044016020604051808303816000875af1158015610a64573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a889190610f4a565b505050508080610a9790610f12565b91505061095e565b5050565b610aab610bc8565b6001805473ffffffffffffffffffffffffffffffffffffffff83167fffffffffffffffffffffffff00000000000000000000000000000000000000009091168117909155610b0e60005473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60005473ffffffffffffffffffffffffffffffffffffffff16331461069f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610324565b600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905561075381610b53565b600060208284031215610c8c57600080fd5b813573ffffffffffffffffffffffffffffffffffffffff81168114610cb057600080fd5b9392505050565b600060208284031215610cc957600080fd5b5035919050565b600060208284031215610ce257600080fd5b815160ff81168114610cb057600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600181815b80851115610d7b57817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115610d6157610d61610cf3565b80851615610d6e57918102915b93841c9390800290610d27565b509250929050565b600082610d9257506001610e3e565b81610d9f57506000610e3e565b8160018114610db55760028114610dbf57610ddb565b6001915050610e3e565b60ff841115610dd057610dd0610cf3565b50506001821b610e3e565b5060208310610133831016604e8410600b8410161715610dfe575081810a610e3e565b610e088383610d22565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115610e3a57610e3a610cf3565b0290505b92915050565b6000610cb060ff841683610d83565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615610e8b57610e8b610cf3565b500290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082610ef5577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60008219821115610f0d57610f0d610cf3565b500190565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610f4357610f43610cf3565b5060010190565b600060208284031215610f5c57600080fd5b81518015158114610cb057600080fdfea164736f6c634300080f000a",
}

// MintManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use MintManagerMetaData.ABI instead.
var MintManagerABI = MintManagerMetaData.ABI

// MintManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MintManagerMetaData.Bin instead.
var MintManagerBin = MintManagerMetaData.Bin

// DeployMintManager deploys a new Ethereum contract, binding an instance of MintManager to it.
func DeployMintManager(auth *bind.TransactOpts, backend bind.ContractBackend, _governanceToken common.Address, _owner common.Address, _recipients []common.Address, _shares []*big.Int) (common.Address, *types.Transaction, *MintManager, error) {
	parsed, err := MintManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MintManagerBin), backend, _governanceToken, _owner, _recipients, _shares)
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

// MINTCAP is a free data retrieval call binding the contract method 0x98f1312e.
//
// Solidity: function MINT_CAP() view returns(uint256)
func (_MintManager *MintManagerCaller) MINTCAP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "MINT_CAP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTCAP is a free data retrieval call binding the contract method 0x98f1312e.
//
// Solidity: function MINT_CAP() view returns(uint256)
func (_MintManager *MintManagerSession) MINTCAP() (*big.Int, error) {
	return _MintManager.Contract.MINTCAP(&_MintManager.CallOpts)
}

// MINTCAP is a free data retrieval call binding the contract method 0x98f1312e.
//
// Solidity: function MINT_CAP() view returns(uint256)
func (_MintManager *MintManagerCallerSession) MINTCAP() (*big.Int, error) {
	return _MintManager.Contract.MINTCAP(&_MintManager.CallOpts)
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

// Minted is a free data retrieval call binding the contract method 0x4f02c420.
//
// Solidity: function minted() view returns(bool)
func (_MintManager *MintManagerCaller) Minted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "minted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Minted is a free data retrieval call binding the contract method 0x4f02c420.
//
// Solidity: function minted() view returns(bool)
func (_MintManager *MintManagerSession) Minted() (bool, error) {
	return _MintManager.Contract.Minted(&_MintManager.CallOpts)
}

// Minted is a free data retrieval call binding the contract method 0x4f02c420.
//
// Solidity: function minted() view returns(bool)
func (_MintManager *MintManagerCallerSession) Minted() (bool, error) {
	return _MintManager.Contract.Minted(&_MintManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MintManager *MintManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MintManager *MintManagerSession) Owner() (common.Address, error) {
	return _MintManager.Contract.Owner(&_MintManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MintManager *MintManagerCallerSession) Owner() (common.Address, error) {
	return _MintManager.Contract.Owner(&_MintManager.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_MintManager *MintManagerCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_MintManager *MintManagerSession) PendingOwner() (common.Address, error) {
	return _MintManager.Contract.PendingOwner(&_MintManager.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_MintManager *MintManagerCallerSession) PendingOwner() (common.Address, error) {
	return _MintManager.Contract.PendingOwner(&_MintManager.CallOpts)
}

// Recipients is a free data retrieval call binding the contract method 0xd1bc76a1.
//
// Solidity: function recipients(uint256 ) view returns(address)
func (_MintManager *MintManagerCaller) Recipients(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "recipients", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Recipients is a free data retrieval call binding the contract method 0xd1bc76a1.
//
// Solidity: function recipients(uint256 ) view returns(address)
func (_MintManager *MintManagerSession) Recipients(arg0 *big.Int) (common.Address, error) {
	return _MintManager.Contract.Recipients(&_MintManager.CallOpts, arg0)
}

// Recipients is a free data retrieval call binding the contract method 0xd1bc76a1.
//
// Solidity: function recipients(uint256 ) view returns(address)
func (_MintManager *MintManagerCallerSession) Recipients(arg0 *big.Int) (common.Address, error) {
	return _MintManager.Contract.Recipients(&_MintManager.CallOpts, arg0)
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

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MintManager *MintManagerTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MintManager *MintManagerSession) AcceptOwnership() (*types.Transaction, error) {
	return _MintManager.Contract.AcceptOwnership(&_MintManager.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MintManager *MintManagerTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _MintManager.Contract.AcceptOwnership(&_MintManager.TransactOpts)
}

// AcceptOwnershipOfToken is a paid mutator transaction binding the contract method 0x28df3f18.
//
// Solidity: function acceptOwnershipOfToken() returns()
func (_MintManager *MintManagerTransactor) AcceptOwnershipOfToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "acceptOwnershipOfToken")
}

// AcceptOwnershipOfToken is a paid mutator transaction binding the contract method 0x28df3f18.
//
// Solidity: function acceptOwnershipOfToken() returns()
func (_MintManager *MintManagerSession) AcceptOwnershipOfToken() (*types.Transaction, error) {
	return _MintManager.Contract.AcceptOwnershipOfToken(&_MintManager.TransactOpts)
}

// AcceptOwnershipOfToken is a paid mutator transaction binding the contract method 0x28df3f18.
//
// Solidity: function acceptOwnershipOfToken() returns()
func (_MintManager *MintManagerTransactorSession) AcceptOwnershipOfToken() (*types.Transaction, error) {
	return _MintManager.Contract.AcceptOwnershipOfToken(&_MintManager.TransactOpts)
}

// Distribute is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns()
func (_MintManager *MintManagerTransactor) Distribute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "distribute")
}

// Distribute is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns()
func (_MintManager *MintManagerSession) Distribute() (*types.Transaction, error) {
	return _MintManager.Contract.Distribute(&_MintManager.TransactOpts)
}

// Distribute is a paid mutator transaction binding the contract method 0xe4fc6b6d.
//
// Solidity: function distribute() returns()
func (_MintManager *MintManagerTransactorSession) Distribute() (*types.Transaction, error) {
	return _MintManager.Contract.Distribute(&_MintManager.TransactOpts)
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

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MintManager *MintManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MintManager *MintManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _MintManager.Contract.RenounceOwnership(&_MintManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MintManager *MintManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MintManager.Contract.RenounceOwnership(&_MintManager.TransactOpts)
}

// RenounceOwnershipOfToken is a paid mutator transaction binding the contract method 0xbaee5ed4.
//
// Solidity: function renounceOwnershipOfToken() returns()
func (_MintManager *MintManagerTransactor) RenounceOwnershipOfToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "renounceOwnershipOfToken")
}

// RenounceOwnershipOfToken is a paid mutator transaction binding the contract method 0xbaee5ed4.
//
// Solidity: function renounceOwnershipOfToken() returns()
func (_MintManager *MintManagerSession) RenounceOwnershipOfToken() (*types.Transaction, error) {
	return _MintManager.Contract.RenounceOwnershipOfToken(&_MintManager.TransactOpts)
}

// RenounceOwnershipOfToken is a paid mutator transaction binding the contract method 0xbaee5ed4.
//
// Solidity: function renounceOwnershipOfToken() returns()
func (_MintManager *MintManagerTransactorSession) RenounceOwnershipOfToken() (*types.Transaction, error) {
	return _MintManager.Contract.RenounceOwnershipOfToken(&_MintManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MintManager *MintManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MintManager *MintManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MintManager.Contract.TransferOwnership(&_MintManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MintManager *MintManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MintManager.Contract.TransferOwnership(&_MintManager.TransactOpts, newOwner)
}

// TransferOwnershipOfToken is a paid mutator transaction binding the contract method 0x457c3977.
//
// Solidity: function transferOwnershipOfToken(address newMintManager) returns()
func (_MintManager *MintManagerTransactor) TransferOwnershipOfToken(opts *bind.TransactOpts, newMintManager common.Address) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "transferOwnershipOfToken", newMintManager)
}

// TransferOwnershipOfToken is a paid mutator transaction binding the contract method 0x457c3977.
//
// Solidity: function transferOwnershipOfToken(address newMintManager) returns()
func (_MintManager *MintManagerSession) TransferOwnershipOfToken(newMintManager common.Address) (*types.Transaction, error) {
	return _MintManager.Contract.TransferOwnershipOfToken(&_MintManager.TransactOpts, newMintManager)
}

// TransferOwnershipOfToken is a paid mutator transaction binding the contract method 0x457c3977.
//
// Solidity: function transferOwnershipOfToken(address newMintManager) returns()
func (_MintManager *MintManagerTransactorSession) TransferOwnershipOfToken(newMintManager common.Address) (*types.Transaction, error) {
	return _MintManager.Contract.TransferOwnershipOfToken(&_MintManager.TransactOpts, newMintManager)
}

// MintManagerOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the MintManager contract.
type MintManagerOwnershipTransferStartedIterator struct {
	Event *MintManagerOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *MintManagerOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintManagerOwnershipTransferStarted)
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
		it.Event = new(MintManagerOwnershipTransferStarted)
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
func (it *MintManagerOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintManagerOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintManagerOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the MintManager contract.
type MintManagerOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_MintManager *MintManagerFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MintManagerOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MintManager.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MintManagerOwnershipTransferStartedIterator{contract: _MintManager.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_MintManager *MintManagerFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *MintManagerOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MintManager.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintManagerOwnershipTransferStarted)
				if err := _MintManager.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_MintManager *MintManagerFilterer) ParseOwnershipTransferStarted(log types.Log) (*MintManagerOwnershipTransferStarted, error) {
	event := new(MintManagerOwnershipTransferStarted)
	if err := _MintManager.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MintManager contract.
type MintManagerOwnershipTransferredIterator struct {
	Event *MintManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MintManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintManagerOwnershipTransferred)
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
		it.Event = new(MintManagerOwnershipTransferred)
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
func (it *MintManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintManagerOwnershipTransferred represents a OwnershipTransferred event raised by the MintManager contract.
type MintManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MintManager *MintManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MintManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MintManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MintManagerOwnershipTransferredIterator{contract: _MintManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MintManager *MintManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MintManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MintManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintManagerOwnershipTransferred)
				if err := _MintManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MintManager *MintManagerFilterer) ParseOwnershipTransferred(log types.Log) (*MintManagerOwnershipTransferred, error) {
	event := new(MintManagerOwnershipTransferred)
	if err := _MintManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
