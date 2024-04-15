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

// KromaMintableERC20FactoryMetaData contains all meta data concerning the KromaMintableERC20Factory contract.
var KromaMintableERC20FactoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_bridge\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BRIDGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createKromaMintableERC20\",\"inputs\":[{\"name\":\"_remoteToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"KromaMintableERC20Created\",\"inputs\":[{\"name\":\"localToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"remoteToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"deployer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161198338038061198383398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b6080516118f26100916000396000818160d101526101a001526118f26000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80635269aa1b1461004657806354fd4d5014610083578063ee9a31a2146100cc575b600080fd5b61005961005436600461033a565b6100f3565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100bf6040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525081565b60405161007a9190610434565b6100597f000000000000000000000000000000000000000000000000000000000000000081565b600073ffffffffffffffffffffffffffffffffffffffff841661019c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603c60248201527f4b726f6d614d696e7461626c654552433230466163746f72793a206d7573742060448201527f70726f766964652072656d6f746520746f6b656e206164647265737300000000606482015260840160405180910390fd5b60007f00000000000000000000000000000000000000000000000000000000000000008585856040516101ce90610253565b6101db949392919061044e565b604051809103906000f0801580156101f7573d6000803e3d6000fd5b5060405133815290915073ffffffffffffffffffffffffffffffffffffffff80871691908316907f16f14001f89df9d8ecc68e7cbb61373ece9025038b9df30bea3635fc0e4701a99060200160405180910390a3949350505050565b611441806104a583390190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f8301126102a057600080fd5b813567ffffffffffffffff808211156102bb576102bb610260565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190828211818310171561030157610301610260565b8160405283815286602085880101111561031a57600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060006060848603121561034f57600080fd5b833573ffffffffffffffffffffffffffffffffffffffff8116811461037357600080fd5b9250602084013567ffffffffffffffff8082111561039057600080fd5b61039c8783880161028f565b935060408601359150808211156103b257600080fd5b506103bf8682870161028f565b9150509250925092565b6000815180845260005b818110156103ef576020818501810151868301820152016103d3565b81811115610401576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061044760208301846103c9565b9392505050565b600073ffffffffffffffffffffffffffffffffffffffff80871683528086166020840152506080604083015261048760808301856103c9565b828103606084015261049981856103c9565b97965050505050505056fe60c06040523480156200001157600080fd5b50604051620014413803806200144183398101604081905262000034916200015a565b8181600362000044838262000279565b50600462000053828262000279565b5050506001600160a01b0392831660805250501660a05262000345565b80516001600160a01b03811681146200008857600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112620000b557600080fd5b81516001600160401b0380821115620000d257620000d26200008d565b604051601f8301601f19908116603f01168101908282118183101715620000fd57620000fd6200008d565b816040528381526020925086838588010111156200011a57600080fd5b600091505b838210156200013e57858201830151818301840152908201906200011f565b83821115620001505760008385830101525b9695505050505050565b600080600080608085870312156200017157600080fd5b6200017c8562000070565b93506200018c6020860162000070565b60408601519093506001600160401b0380821115620001aa57600080fd5b620001b888838901620000a3565b93506060870151915080821115620001cf57600080fd5b50620001de87828801620000a3565b91505092959194509250565b600181811c90821680620001ff57607f821691505b6020821081036200022057634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200027457600081815260208120601f850160051c810160208610156200024f5750805b601f850160051c820191505b8181101562000270578281556001016200025b565b5050505b505050565b81516001600160401b038111156200029557620002956200008d565b620002ad81620002a68454620001ea565b8462000226565b602080601f831160018114620002e55760008415620002cc5750858301515b600019600386901b1c1916600185901b17855562000270565b600085815260208120601f198616915b828110156200031657888601518255948401946001909101908401620002f5565b5085821015620003355787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60805160a0516110c862000379600039600081816103160152818161050a01526106270152600061014d01526110c86000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c806340c10f19116100b25780639dc29fac11610081578063a9059cbb11610066578063a9059cbb146102b8578063dd62ed3e146102cb578063ee9a31a21461031157600080fd5b80639dc29fac14610292578063a457c2d7146102a557600080fd5b806340c10f191461020357806354fd4d501461021857806370a082311461025457806395d89b411461028a57600080fd5b806318160ddd116100ee57806318160ddd146101bc57806323b872dd146101ce578063313ce567146101e157806339509351146101f057600080fd5b806301ffc9a714610120578063033964be1461014857806306fdde0314610194578063095ea7b3146101a9575b600080fd5b61013361012e366004610e90565b610338565b60405190151581526020015b60405180910390f35b61016f7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161013f565b61019c6103d8565b60405161013f9190610ed9565b6101336101b7366004610f75565b61046a565b6002545b60405190815260200161013f565b6101336101dc366004610f9f565b610482565b6040516012815260200161013f565b6101336101fe366004610f75565b6104a6565b610216610211366004610f75565b6104f2565b005b61019c6040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525081565b6101c0610262366004610fdb565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b61019c610600565b6102166102a0366004610f75565b61060f565b6101336102b3366004610f75565b61070c565b6101336102c6366004610f75565b6107c3565b6101c06102d9366004610ff6565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b61016f7f000000000000000000000000000000000000000000000000000000000000000081565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007f30a0c5a9000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000084168214806103d057507fffffffff00000000000000000000000000000000000000000000000000000000848116908216145b949350505050565b6060600380546103e790611029565b80601f016020809104026020016040519081016040528092919081815260200182805461041390611029565b80156104605780601f1061043557610100808354040283529160200191610460565b820191906000526020600020905b81548152906001019060200180831161044357829003601f168201915b5050505050905090565b6000336104788185856107d1565b5060019392505050565b600033610490858285610951565b61049b858585610a0e565b506001949350505050565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919061047890829086906104ed90879061107c565b6107d1565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146105a25760405162461bcd60e51b815260206004820152603160248201527f4b726f6d614d696e7461626c6545524332303a206f6e6c79206272696467652060448201527f63616e206d696e7420616e64206275726e00000000000000000000000000000060648201526084015b60405180910390fd5b6105ac8282610c2f565b8173ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885826040516105f491815260200190565b60405180910390a25050565b6060600480546103e790611029565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146106ba5760405162461bcd60e51b815260206004820152603160248201527f4b726f6d614d696e7461626c6545524332303a206f6e6c79206272696467652060448201527f63616e206d696e7420616e64206275726e0000000000000000000000000000006064820152608401610599565b6106c48282610d08565b8173ffffffffffffffffffffffffffffffffffffffff167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5826040516105f491815260200190565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152812054909190838110156107b65760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f0000000000000000000000000000000000000000000000000000006064820152608401610599565b61049b82868684036107d1565b600033610478818585610a0e565b73ffffffffffffffffffffffffffffffffffffffff83166108595760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152608401610599565b73ffffffffffffffffffffffffffffffffffffffff82166108e25760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f73730000000000000000000000000000000000000000000000000000000000006064820152608401610599565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610a0857818110156109fb5760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610599565b610a0884848484036107d1565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610a975760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f64726573730000000000000000000000000000000000000000000000000000006064820152608401610599565b73ffffffffffffffffffffffffffffffffffffffff8216610b205760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152608401610599565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610bbc5760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e636500000000000000000000000000000000000000000000000000006064820152608401610599565b73ffffffffffffffffffffffffffffffffffffffff848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3610a08565b73ffffffffffffffffffffffffffffffffffffffff8216610c925760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610599565b8060026000828254610ca4919061107c565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b73ffffffffffffffffffffffffffffffffffffffff8216610d915760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f73000000000000000000000000000000000000000000000000000000000000006064820152608401610599565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090205481811015610e2d5760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f63650000000000000000000000000000000000000000000000000000000000006064820152608401610599565b73ffffffffffffffffffffffffffffffffffffffff83166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9101610944565b600060208284031215610ea257600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610ed257600080fd5b9392505050565b600060208083528351808285015260005b81811015610f0657858101830151858201604001528201610eea565b81811115610f18576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610f7057600080fd5b919050565b60008060408385031215610f8857600080fd5b610f9183610f4c565b946020939093013593505050565b600080600060608486031215610fb457600080fd5b610fbd84610f4c565b9250610fcb60208501610f4c565b9150604084013590509250925092565b600060208284031215610fed57600080fd5b610ed282610f4c565b6000806040838503121561100957600080fd5b61101283610f4c565b915061102060208401610f4c565b90509250929050565b600181811c9082168061103d57607f821691505b602082108103611076577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b600082198211156110b6577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b50019056fea164736f6c634300080f000aa164736f6c634300080f000a",
}

// KromaMintableERC20FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use KromaMintableERC20FactoryMetaData.ABI instead.
var KromaMintableERC20FactoryABI = KromaMintableERC20FactoryMetaData.ABI

// KromaMintableERC20FactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use KromaMintableERC20FactoryMetaData.Bin instead.
var KromaMintableERC20FactoryBin = KromaMintableERC20FactoryMetaData.Bin

// DeployKromaMintableERC20Factory deploys a new Ethereum contract, binding an instance of KromaMintableERC20Factory to it.
func DeployKromaMintableERC20Factory(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address) (common.Address, *types.Transaction, *KromaMintableERC20Factory, error) {
	parsed, err := KromaMintableERC20FactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KromaMintableERC20FactoryBin), backend, _bridge)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KromaMintableERC20Factory{KromaMintableERC20FactoryCaller: KromaMintableERC20FactoryCaller{contract: contract}, KromaMintableERC20FactoryTransactor: KromaMintableERC20FactoryTransactor{contract: contract}, KromaMintableERC20FactoryFilterer: KromaMintableERC20FactoryFilterer{contract: contract}}, nil
}

// KromaMintableERC20Factory is an auto generated Go binding around an Ethereum contract.
type KromaMintableERC20Factory struct {
	KromaMintableERC20FactoryCaller     // Read-only binding to the contract
	KromaMintableERC20FactoryTransactor // Write-only binding to the contract
	KromaMintableERC20FactoryFilterer   // Log filterer for contract events
}

// KromaMintableERC20FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type KromaMintableERC20FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KromaMintableERC20FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KromaMintableERC20FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KromaMintableERC20FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KromaMintableERC20FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KromaMintableERC20FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KromaMintableERC20FactorySession struct {
	Contract     *KromaMintableERC20Factory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// KromaMintableERC20FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KromaMintableERC20FactoryCallerSession struct {
	Contract *KromaMintableERC20FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// KromaMintableERC20FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KromaMintableERC20FactoryTransactorSession struct {
	Contract     *KromaMintableERC20FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// KromaMintableERC20FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type KromaMintableERC20FactoryRaw struct {
	Contract *KromaMintableERC20Factory // Generic contract binding to access the raw methods on
}

// KromaMintableERC20FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KromaMintableERC20FactoryCallerRaw struct {
	Contract *KromaMintableERC20FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// KromaMintableERC20FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KromaMintableERC20FactoryTransactorRaw struct {
	Contract *KromaMintableERC20FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKromaMintableERC20Factory creates a new instance of KromaMintableERC20Factory, bound to a specific deployed contract.
func NewKromaMintableERC20Factory(address common.Address, backend bind.ContractBackend) (*KromaMintableERC20Factory, error) {
	contract, err := bindKromaMintableERC20Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KromaMintableERC20Factory{KromaMintableERC20FactoryCaller: KromaMintableERC20FactoryCaller{contract: contract}, KromaMintableERC20FactoryTransactor: KromaMintableERC20FactoryTransactor{contract: contract}, KromaMintableERC20FactoryFilterer: KromaMintableERC20FactoryFilterer{contract: contract}}, nil
}

// NewKromaMintableERC20FactoryCaller creates a new read-only instance of KromaMintableERC20Factory, bound to a specific deployed contract.
func NewKromaMintableERC20FactoryCaller(address common.Address, caller bind.ContractCaller) (*KromaMintableERC20FactoryCaller, error) {
	contract, err := bindKromaMintableERC20Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KromaMintableERC20FactoryCaller{contract: contract}, nil
}

// NewKromaMintableERC20FactoryTransactor creates a new write-only instance of KromaMintableERC20Factory, bound to a specific deployed contract.
func NewKromaMintableERC20FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*KromaMintableERC20FactoryTransactor, error) {
	contract, err := bindKromaMintableERC20Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KromaMintableERC20FactoryTransactor{contract: contract}, nil
}

// NewKromaMintableERC20FactoryFilterer creates a new log filterer instance of KromaMintableERC20Factory, bound to a specific deployed contract.
func NewKromaMintableERC20FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*KromaMintableERC20FactoryFilterer, error) {
	contract, err := bindKromaMintableERC20Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KromaMintableERC20FactoryFilterer{contract: contract}, nil
}

// bindKromaMintableERC20Factory binds a generic wrapper to an already deployed contract.
func bindKromaMintableERC20Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KromaMintableERC20FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KromaMintableERC20Factory *KromaMintableERC20FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KromaMintableERC20Factory.Contract.KromaMintableERC20FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KromaMintableERC20Factory *KromaMintableERC20FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KromaMintableERC20Factory.Contract.KromaMintableERC20FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KromaMintableERC20Factory *KromaMintableERC20FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KromaMintableERC20Factory.Contract.KromaMintableERC20FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KromaMintableERC20Factory *KromaMintableERC20FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KromaMintableERC20Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KromaMintableERC20Factory *KromaMintableERC20FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KromaMintableERC20Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KromaMintableERC20Factory *KromaMintableERC20FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KromaMintableERC20Factory.Contract.contract.Transact(opts, method, params...)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_KromaMintableERC20Factory *KromaMintableERC20FactoryCaller) BRIDGE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KromaMintableERC20Factory.contract.Call(opts, &out, "BRIDGE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_KromaMintableERC20Factory *KromaMintableERC20FactorySession) BRIDGE() (common.Address, error) {
	return _KromaMintableERC20Factory.Contract.BRIDGE(&_KromaMintableERC20Factory.CallOpts)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_KromaMintableERC20Factory *KromaMintableERC20FactoryCallerSession) BRIDGE() (common.Address, error) {
	return _KromaMintableERC20Factory.Contract.BRIDGE(&_KromaMintableERC20Factory.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_KromaMintableERC20Factory *KromaMintableERC20FactoryCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KromaMintableERC20Factory.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_KromaMintableERC20Factory *KromaMintableERC20FactorySession) Version() (string, error) {
	return _KromaMintableERC20Factory.Contract.Version(&_KromaMintableERC20Factory.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_KromaMintableERC20Factory *KromaMintableERC20FactoryCallerSession) Version() (string, error) {
	return _KromaMintableERC20Factory.Contract.Version(&_KromaMintableERC20Factory.CallOpts)
}

// CreateKromaMintableERC20 is a paid mutator transaction binding the contract method 0x5269aa1b.
//
// Solidity: function createKromaMintableERC20(address _remoteToken, string _name, string _symbol) returns(address)
func (_KromaMintableERC20Factory *KromaMintableERC20FactoryTransactor) CreateKromaMintableERC20(opts *bind.TransactOpts, _remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _KromaMintableERC20Factory.contract.Transact(opts, "createKromaMintableERC20", _remoteToken, _name, _symbol)
}

// CreateKromaMintableERC20 is a paid mutator transaction binding the contract method 0x5269aa1b.
//
// Solidity: function createKromaMintableERC20(address _remoteToken, string _name, string _symbol) returns(address)
func (_KromaMintableERC20Factory *KromaMintableERC20FactorySession) CreateKromaMintableERC20(_remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _KromaMintableERC20Factory.Contract.CreateKromaMintableERC20(&_KromaMintableERC20Factory.TransactOpts, _remoteToken, _name, _symbol)
}

// CreateKromaMintableERC20 is a paid mutator transaction binding the contract method 0x5269aa1b.
//
// Solidity: function createKromaMintableERC20(address _remoteToken, string _name, string _symbol) returns(address)
func (_KromaMintableERC20Factory *KromaMintableERC20FactoryTransactorSession) CreateKromaMintableERC20(_remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _KromaMintableERC20Factory.Contract.CreateKromaMintableERC20(&_KromaMintableERC20Factory.TransactOpts, _remoteToken, _name, _symbol)
}

// KromaMintableERC20FactoryKromaMintableERC20CreatedIterator is returned from FilterKromaMintableERC20Created and is used to iterate over the raw logs and unpacked data for KromaMintableERC20Created events raised by the KromaMintableERC20Factory contract.
type KromaMintableERC20FactoryKromaMintableERC20CreatedIterator struct {
	Event *KromaMintableERC20FactoryKromaMintableERC20Created // Event containing the contract specifics and raw log

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
func (it *KromaMintableERC20FactoryKromaMintableERC20CreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KromaMintableERC20FactoryKromaMintableERC20Created)
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
		it.Event = new(KromaMintableERC20FactoryKromaMintableERC20Created)
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
func (it *KromaMintableERC20FactoryKromaMintableERC20CreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KromaMintableERC20FactoryKromaMintableERC20CreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KromaMintableERC20FactoryKromaMintableERC20Created represents a KromaMintableERC20Created event raised by the KromaMintableERC20Factory contract.
type KromaMintableERC20FactoryKromaMintableERC20Created struct {
	LocalToken  common.Address
	RemoteToken common.Address
	Deployer    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterKromaMintableERC20Created is a free log retrieval operation binding the contract event 0x16f14001f89df9d8ecc68e7cbb61373ece9025038b9df30bea3635fc0e4701a9.
//
// Solidity: event KromaMintableERC20Created(address indexed localToken, address indexed remoteToken, address deployer)
func (_KromaMintableERC20Factory *KromaMintableERC20FactoryFilterer) FilterKromaMintableERC20Created(opts *bind.FilterOpts, localToken []common.Address, remoteToken []common.Address) (*KromaMintableERC20FactoryKromaMintableERC20CreatedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}

	logs, sub, err := _KromaMintableERC20Factory.contract.FilterLogs(opts, "KromaMintableERC20Created", localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return &KromaMintableERC20FactoryKromaMintableERC20CreatedIterator{contract: _KromaMintableERC20Factory.contract, event: "KromaMintableERC20Created", logs: logs, sub: sub}, nil
}

// WatchKromaMintableERC20Created is a free log subscription operation binding the contract event 0x16f14001f89df9d8ecc68e7cbb61373ece9025038b9df30bea3635fc0e4701a9.
//
// Solidity: event KromaMintableERC20Created(address indexed localToken, address indexed remoteToken, address deployer)
func (_KromaMintableERC20Factory *KromaMintableERC20FactoryFilterer) WatchKromaMintableERC20Created(opts *bind.WatchOpts, sink chan<- *KromaMintableERC20FactoryKromaMintableERC20Created, localToken []common.Address, remoteToken []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}

	logs, sub, err := _KromaMintableERC20Factory.contract.WatchLogs(opts, "KromaMintableERC20Created", localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KromaMintableERC20FactoryKromaMintableERC20Created)
				if err := _KromaMintableERC20Factory.contract.UnpackLog(event, "KromaMintableERC20Created", log); err != nil {
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

// ParseKromaMintableERC20Created is a log parse operation binding the contract event 0x16f14001f89df9d8ecc68e7cbb61373ece9025038b9df30bea3635fc0e4701a9.
//
// Solidity: event KromaMintableERC20Created(address indexed localToken, address indexed remoteToken, address deployer)
func (_KromaMintableERC20Factory *KromaMintableERC20FactoryFilterer) ParseKromaMintableERC20Created(log types.Log) (*KromaMintableERC20FactoryKromaMintableERC20Created, error) {
	event := new(KromaMintableERC20FactoryKromaMintableERC20Created)
	if err := _KromaMintableERC20Factory.contract.UnpackLog(event, "KromaMintableERC20Created", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
