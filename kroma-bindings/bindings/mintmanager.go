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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_governanceToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_recipients\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"GOVERNANCE_TOKEN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractGovernanceToken\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MINT_CAP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SHARE_DENOMINATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"distribute\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"minted\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recipients\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnershipOfToken\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shareOf\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnershipOfToken\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x60a06040523480156200001157600080fd5b506040516200153a3803806200153a83398101604081905262000034916200050e565b6200003f33620002e2565b6001600160a01b038416608052620000578362000332565b8051825114620000ba5760405162461bcd60e51b8152602060048201526024808201527f4d696e744d616e616765723a20696e76616c6964206c656e677468206f6620616044820152637272617960e01b60648201526084015b60405180910390fd5b6000805b83518110156200024f576000848281518110620000df57620000df62000604565b6020026020010151905060006001600160a01b0316816001600160a01b031603620001605760405162461bcd60e51b815260206004820152602a60248201527f4d696e744d616e616765723a20726563697069656e74206164647265737320636044820152690616e6e6f7420626520360b41b6064820152608401620000b1565b600084838151811062000177576200017762000604565b6020026020010151905080600003620001d35760405162461bcd60e51b815260206004820152601e60248201527f4d696e744d616e616765723a2073686172652063616e6e6f74206265203000006044820152606401620000b1565b620001df818562000630565b600180548082019091557fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180546001600160a01b039094166001600160a01b0319909416841790556000928352600260205260409092205591508062000246816200064b565b915050620000be565b506064811115620002d75760405162461bcd60e51b8152602060048201526044602482018190527f4d696e744d616e616765723a206d617820746f74616c20736861726520697320908201527f657175616c206f72206c657373207468616e2053484152455f44454e4f4d494e60648201526320aa27a960e11b608482015260a401620000b1565b505050505062000667565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6200033c620003b1565b6001600160a01b038116620003a35760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401620000b1565b620003ae81620002e2565b50565b6000546001600160a01b031633146200040d5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401620000b1565b565b80516001600160a01b03811681146200042757600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b03811182821017156200046d576200046d6200042c565b604052919050565b60006001600160401b038211156200049157620004916200042c565b5060051b60200190565b600082601f830112620004ad57600080fd5b81516020620004c6620004c08362000475565b62000442565b82815260059290921b84018101918181019086841115620004e657600080fd5b8286015b84811015620005035780518352918301918301620004ea565b509695505050505050565b600080600080608085870312156200052557600080fd5b62000530856200040f565b93506020620005418187016200040f565b60408701519094506001600160401b03808211156200055f57600080fd5b818801915088601f8301126200057457600080fd5b815162000585620004c08262000475565b81815260059190911b8301840190848101908b831115620005a557600080fd5b938501935b82851015620005ce57620005be856200040f565b82529385019390850190620005aa565b60608b01519097509450505080831115620005e857600080fd5b5050620005f8878288016200049b565b91505092959194509250565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600082198211156200064657620006466200061a565b500190565b6000600182016200066057620006606200061a565b5060010190565b608051610e86620006b460003960008181610126015281816102e00152818161045001528181610551015281816105ce01528181610691015281816107f501526108ad0152610e866000f3fe608060405234801561001057600080fd5b50600436106100df5760003560e01c80637eb118451161008c578063baee5ed411610066578063baee5ed4146101ee578063d1bc76a1146101f6578063e4fc6b6d14610209578063f2fde38b1461021157600080fd5b80637eb11845146101bd5780638da5cb5b146101c557806398f1312e146101e357600080fd5b8063457c3977116100bd578063457c39771461016d5780634f02c42014610180578063715018a6146101b557600080fd5b80631249c58b146100e457806321e5e2c4146100ee5780632efd46d614610121575b600080fd5b6100ec610224565b005b61010e6100fc366004610b6e565b60026020526000908152604090205481565b6040519081526020015b60405180910390f35b6101487f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610118565b6100ec61017b366004610b6e565b610504565b6000546101a59074010000000000000000000000000000000000000000900460ff1681565b6040519015158152602001610118565b6100ec6105b0565b61010e606481565b60005473ffffffffffffffffffffffffffffffffffffffff16610148565b61010e633b9aca0081565b6100ec6105c4565b610148610204366004610bab565b61064e565b6100ec610685565b6100ec61021f366004610b6e565b6109c1565b61022c610a78565b60005474010000000000000000000000000000000000000000900460ff16156102dc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f4d696e744d616e616765723a20616c7265616479206d696e746564206f6e207460448201527f68697320636861696e000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015610349573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061036d9190610bc4565b61037890600a610d38565b61038690633b9aca00610d47565b90506000805b60015481101561041a576000600182815481106103ab576103ab610d84565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff16808352600290915260408220549092509060646103ec8388610d47565b6103f69190610db3565b90506104028186610dee565b9450505050808061041290610e06565b91505061038c565b506040517f40c10f19000000000000000000000000000000000000000000000000000000008152306004820152602481018290527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906340c10f1990604401600060405180830381600087803b1580156104a957600080fd5b505af11580156104bd573d6000803e3d6000fd5b5050600080547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000017905550505050565b61050c610a78565b6040517ff2fde38b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063f2fde38b90602401600060405180830381600087803b15801561059557600080fd5b505af11580156105a9573d6000803e3d6000fd5b5050505050565b6105b8610a78565b6105c26000610af9565b565b6105cc610a78565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663715018a66040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561063457600080fd5b505af1158015610648573d6000803e3d6000fd5b50505050565b6001818154811061065e57600080fd5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b61068d610a78565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106fa573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061071e9190610bc4565b61072990600a610d38565b61073790633b9aca00610d47565b905060005b60015481101561087b5760006001828154811061075b5761075b610d84565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff168083526002909152604082205490925090606461079c8387610d47565b6107a69190610db3565b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018390529192507f00000000000000000000000000000000000000000000000000000000000000009091169063a9059cbb906044016020604051808303816000875af1158015610840573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108649190610e3e565b50505050808061087390610e06565b91505061073c565b506040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa158015610909573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061092d9190610e60565b905080156109bd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4d696e744d616e616765723a20746f6b656e732072656d61696e20616674657260448201527f20646973747269627574696f6e0000000000000000000000000000000000000060648201526084016102d3565b5050565b6109c9610a78565b73ffffffffffffffffffffffffffffffffffffffff8116610a6c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016102d3565b610a7581610af9565b50565b60005473ffffffffffffffffffffffffffffffffffffffff1633146105c2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102d3565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600060208284031215610b8057600080fd5b813573ffffffffffffffffffffffffffffffffffffffff81168114610ba457600080fd5b9392505050565b600060208284031215610bbd57600080fd5b5035919050565b600060208284031215610bd657600080fd5b815160ff81168114610ba457600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600181815b80851115610c6f57817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115610c5557610c55610be7565b80851615610c6257918102915b93841c9390800290610c1b565b509250929050565b600082610c8657506001610d32565b81610c9357506000610d32565b8160018114610ca95760028114610cb357610ccf565b6001915050610d32565b60ff841115610cc457610cc4610be7565b50506001821b610d32565b5060208310610133831016604e8410600b8410161715610cf2575081810a610d32565b610cfc8383610c16565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115610d2e57610d2e610be7565b0290505b92915050565b6000610ba460ff841683610c77565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615610d7f57610d7f610be7565b500290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082610de9577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60008219821115610e0157610e01610be7565b500190565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610e3757610e37610be7565b5060010190565b600060208284031215610e5057600080fd5b81518015158114610ba457600080fd5b600060208284031215610e7257600080fd5b505191905056fea164736f6c634300080f000a",
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
// Solidity: function transferOwnershipOfToken(address newOwner) returns()
func (_MintManager *MintManagerTransactor) TransferOwnershipOfToken(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "transferOwnershipOfToken", newOwner)
}

// TransferOwnershipOfToken is a paid mutator transaction binding the contract method 0x457c3977.
//
// Solidity: function transferOwnershipOfToken(address newOwner) returns()
func (_MintManager *MintManagerSession) TransferOwnershipOfToken(newOwner common.Address) (*types.Transaction, error) {
	return _MintManager.Contract.TransferOwnershipOfToken(&_MintManager.TransactOpts, newOwner)
}

// TransferOwnershipOfToken is a paid mutator transaction binding the contract method 0x457c3977.
//
// Solidity: function transferOwnershipOfToken(address newOwner) returns()
func (_MintManager *MintManagerTransactorSession) TransferOwnershipOfToken(newOwner common.Address) (*types.Transaction, error) {
	return _MintManager.Contract.TransferOwnershipOfToken(&_MintManager.TransactOpts, newOwner)
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
