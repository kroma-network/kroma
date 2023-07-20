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

// ZKMerkleTrieMetaData contains all meta data concerning the ZKMerkleTrie contract.
var ZKMerkleTrieMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poseidon2\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"POSEIDON2\",\"outputs\":[{\"internalType\":\"contractIPoseidon2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"_proofs\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_proofs\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"verifyInclusionProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051611b69380380611b6983398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b608051611ac36100a6600039600081816094015281816107ca015281816109e801528181610c3c01526110fd0152611ac36000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806312e64a7214610046578063c423b1e81461006e578063dc8b50381461008f575b600080fd5b61005961005436600461170f565b6100db565b60405190151581526020015b60405180910390f35b61008161007c366004611784565b6101a6565b60405161006592919061184e565b6100b67f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610065565b6040517fc423b1e800000000000000000000000000000000000000000000000000000000815260009081908190309063c423b1e890610122908a9089908990600401611869565b600060405180830381865afa15801561013f573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820160405261018591908101906118f9565b9150915081801561019b575061019b868261074b565b979650505050505050565b600060606002845110156102275760405162461bcd60e51b815260206004820152602960248201527f5a4b4d65726b6c65547269653a2070726f76696465642070726f6f662069732060448201527f746f6f2073686f7274000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b61027e846001865161023991906119ba565b81518110610249576102496119d1565b602002602001015180516020909101207f950654da67865a81bc70e45f3230f5179f08e29c66184bf746f71050f117b3b81490565b6102f05760405162461bcd60e51b815260206004820152602d60248201527f5a4b4d65726b6c65547269653a20746865206c617374206974656d206973206e60448201527f6f74206d61676963206861736800000000000000000000000000000000000000606482015260840161021e565b60006102fb86610767565b9050600061030886610846565b90506103526040805161010081019091528060008152600060208201819052604082018190526060808301829052608083015260a0820181905260c0820181905260e09091015290565b604080516020810190915260008082528351909182918291908290610379906002906119ba565b90505b86818151811061038e5761038e6119d1565b6020026020010151602001519550600060038111156103af576103af611a00565b865160038111156103c2576103c2611a00565b036104a05760006103d38983610939565b90508015610433578660200151861461042e5760405162461bcd60e51b815260206004820152601b60248201527f5a4b4d65726b6c65547269653a20696e76616c6964206b6579204c0000000000604482015260640161021e565b610486565b866040015186146104865760405162461bcd60e51b815260206004820152601b60248201527f5a4b4d65726b6c65547269653a20696e76616c6964206b657920520000000000604482015260640161021e565b61049887602001518860400151610996565b9550506106b5565b6001865160038111156104b5576104b5611a00565b0361061857831580156104c6575082155b6105385760405162461bcd60e51b815260206004820152602660248201527f5a4b4d65726b6c65547269653a206475706c696361746564207465726d696e6160448201527f6c206e6f64650000000000000000000000000000000000000000000000000000606482015260840161021e565b87866060015114935083156107385761056b600160001b87606001516105668960a001518a60800151610a63565b610bff565b6080870151805160208102825260e0890151929750909350839115610611578e8860e00151148061059f5750898860e00151145b6106115760405162461bcd60e51b815260206004820152602260248201527f5a4b4d65726b6c65547269653a20696e76616c6964206b657920707265696d6160448201527f6765000000000000000000000000000000000000000000000000000000000000606482015260840161021e565b50506106b5565b60028651600381111561062d5761062d611a00565b036106b5578315801561063e575082155b6106b05760405162461bcd60e51b815260206004820152602660248201527f5a4b4d65726b6c65547269653a206475706c696361746564207465726d696e6160448201527f6c206e6f64650000000000000000000000000000000000000000000000000000606482015260840161021e565b600192505b80600003610711578a851461070c5760405162461bcd60e51b815260206004820152601960248201527f5a4b4d65726b65547269653a20696e76616c696420726f6f7400000000000000604482015260640161021e565b610738565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161037c565b50919b919a509098505050505050505050565b6000818051906020012083805190602001201490505b92915050565b600080600061077584610d77565b6040805180820182528381526020810183905290517f299e566000000000000000000000000000000000000000000000000000000000815292945090925073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169163299e5660916107fd91600401611a2f565b602060405180830381865afa15801561081a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061083e9190611a60565b949350505050565b805160609060008167ffffffffffffffff81111561086657610866611556565b60405190808252806020026020018201604052801561089f57816020015b61088c6114f7565b8152602001906001900390816108845790505b50905060005b6108b06001846119ba565b8110156109315760006108db8683815181106108ce576108ce6119d1565b6020026020010151610d9f565b905060405180604001604052808784815181106108fa576108fa6119d1565b602002602001015181526020018281525083838151811061091d5761091d6119d1565b6020908102919091010152506001016108a5565b509392505050565b6000610100821061098c5760405162461bcd60e51b815260206004820152601c60248201527f5a4b4d65726b6c65547269653a20746f6f206c6f6e6720646570746800000000604482015260640161021e565b506001901b161590565b6040805180820182528381526020810183905290517f299e566000000000000000000000000000000000000000000000000000000000815260009173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169163299e566091610a1b91600401611a2f565b602060405180830381865afa158015610a38573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a5c9190611a60565b9392505050565b6000600182511015610add5760405162461bcd60e51b815260206004820152602b60248201527f5a4b547269654861736865723a20746f6f206665772076616c75657320666f7260448201527f205f76616c756548617368000000000000000000000000000000000000000000606482015260840161021e565b6000825167ffffffffffffffff811115610af957610af9611556565b604051908082528060200260200182016040528015610b22578160200160208202803683370190505b50905060005b8351811015610bc8576001811b851663ffffffff1615610b8757610b64848281518110610b5757610b576119d1565b6020026020010151610767565b828281518110610b7657610b766119d1565b602002602001018181525050610bc0565b838181518110610b9957610b996119d1565b6020026020010151828281518110610bb357610bb36119d1565b6020026020010181815250505b600101610b28565b50600283511015610bf65780600081518110610be657610be66119d1565b6020026020010151915050610761565b61083e8161107f565b6040805180820182528481526020810184905290517f299e56600000000000000000000000000000000000000000000000000000000081526000917f00000000000000000000000000000000000000000000000000000000000000009173ffffffffffffffffffffffffffffffffffffffff83169163299e566091610c879190600401611a2f565b602060405180830381865afa158015610ca4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cc89190611a60565b6040805180820182528281526020810186905290517f299e566000000000000000000000000000000000000000000000000000000000815291965073ffffffffffffffffffffffffffffffffffffffff83169163299e566091610d2d91600401611a2f565b602060405180830381865afa158015610d4a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d6e9190611a60565b95945050505050565b60008082608081901b610d8a8260801c90565b610d948260801c90565b935093505050915091565b60408051610100810182526000808252602082018190529181018290526060808201839052608082015260a0810182905260c0810182905260e081019190915260408051610100810182526000808252602082018190529181018290526060808201839052608082015260a0810182905260c0810182905260e08101919091526000610e53846040805180820182526060815260006020918201528151808301909252828101825291519181019190915290565b90506000610e6082611269565b60ff16905080610e8b57610e73826112f0565b6020840152610e81826112f0565b604084015261103b565b60018103610f7557610e9c826112f0565b6060840152600080610ead8461138e565b63ffffffff821660a088015290925060ff1690508067ffffffffffffffff811115610eda57610eda611556565b604051908082528060200260200182016040528015610f03578160200160208202803683370190505b50608086015260005b81811015610f4657610f1d856112f0565b86608001518281518110610f3357610f336119d1565b6020908102919091010152600101610f0c565b506000610f5285611269565b60ff1690508015610f6d57610f678582611428565b60e08701525b50505061103b565b6002811461103b5760038103610ff35760405162461bcd60e51b815260206004820152602560248201527f4e6f64655265616465723a20756e657870656374656420726f6f74206e6f646560448201527f2074797065000000000000000000000000000000000000000000000000000000606482015260840161021e565b60405162461bcd60e51b815260206004820152601d60248201527f4e6f64655265616465723a20696e76616c6964206e6f64652074797065000000604482015260640161021e565b80600381111561104d5761104d611a00565b8390600381111561106057611060611a00565b9081600381111561107357611073611a00565b90525091949350505050565b60006004825110156110f95760405162461bcd60e51b815260206004820152602b60248201527f5a4b547269654861736865723a20746f6f206665772076616c75657320666f7260448201527f205f68617368456c656d73000000000000000000000000000000000000000000606482015260840161021e565b81517f00000000000000000000000000000000000000000000000000000000000000009060009081906001906002905b8083101561124157600094505b808510156112335782850193508084101561122a578573ffffffffffffffffffffffffffffffffffffffff1663299e566060405180604001604052808b8981518110611184576111846119d1565b602002602001015181526020018b88815181106111a3576111a36119d1565b60200260200101518152506040518263ffffffff1660e01b81526004016111ca9190611a2f565b602060405180830381865afa1580156111e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061120b9190611a60565b88868151811061121d5761121d6119d1565b6020026020010181815250505b93810193611136565b909150600182901b90611129565b87600081518110611254576112546119d1565b60200260200101519650505050505050919050565b60006001826020015110156112c05760405162461bcd60e51b815260206004820152601f60248201527f4e6f64655265616465723a20746f6f2073686f727420666f722075696e743800604482015260640161021e565b81518051600180830180865260208601805191949360f81c92916112e59083906119ba565b905250949350505050565b600060208260200151101561136d5760405162461bcd60e51b815260206004820152602160248201527f4e6f64655265616465723a20746f6f2073686f727420666f722062797465733360448201527f3200000000000000000000000000000000000000000000000000000000000000606482015260840161021e565b815180516020808301808652818601805191949392916112e59083906119ba565b6000806004836020015110156113e65760405162461bcd60e51b815260206004820181905260248201527f4e6f64655265616465723a20746f6f2073686f727420666f722075696e743332604482015260640161021e565b8251805160048083018087526020870180519194939260f084901c9260f885901c92906114149083906119ba565b90525060ff90911697909650945050505050565b600081836020015110156114a45760405162461bcd60e51b815260206004820152602160248201527f4e6f64655265616465723a20746f6f2073686f727420666f72206e206279746560448201527f7300000000000000000000000000000000000000000000000000000000000000606482015260840161021e565b82516060906000806114b7866008611a79565b6114c3906101006119ba565b8351848801808a5260208a01805191975091831c945091925087916114e99083906119ba565b905250909695505050505050565b6040518060400160405280606081526020016115516040805161010081019091528060008152600060208201819052604082018190526060808301829052608083015260a0820181905260c0820181905260e09091015290565b905290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156115cc576115cc611556565b604052919050565b600067ffffffffffffffff8211156115ee576115ee611556565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f83011261162b57600080fd5b813561163e611639826115d4565b611585565b81815284602083860101111561165357600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f83011261168157600080fd5b8135602067ffffffffffffffff8083111561169e5761169e611556565b8260051b6116ad838201611585565b93845285810183019383810190888611156116c757600080fd5b84880192505b85831015611703578235848111156116e55760008081fd5b6116f38a87838c010161161a565b83525091840191908401906116cd565b98975050505050505050565b6000806000806080858703121561172557600080fd5b84359350602085013567ffffffffffffffff8082111561174457600080fd5b6117508883890161161a565b9450604087013591508082111561176657600080fd5b5061177387828801611670565b949793965093946060013593505050565b60008060006060848603121561179957600080fd5b83359250602084013567ffffffffffffffff8111156117b757600080fd5b6117c386828701611670565b925050604084013590509250925092565b60005b838110156117ef5781810151838201526020016117d7565b838111156117fe576000848401525b50505050565b6000815180845261181c8160208601602086016117d4565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b821515815260406020820152600061083e6040830184611804565b600060608201858352602060608185015281865180845260808601915060808160051b870101935082880160005b828110156118e3577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808887030184526118d1868351611804565b95509284019290840190600101611897565b5050505050604092909201929092529392505050565b6000806040838503121561190c57600080fd5b8251801515811461191c57600080fd5b602084015190925067ffffffffffffffff81111561193957600080fd5b8301601f8101851361194a57600080fd5b8051611958611639826115d4565b81815286602083850101111561196d57600080fd5b61197e8260208301602086016117d4565b8093505050509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000828210156119cc576119cc61198b565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60408101818360005b6002811015611a57578151835260209283019290910190600101611a38565b50505092915050565b600060208284031215611a7257600080fd5b5051919050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611ab157611ab161198b565b50029056fea164736f6c634300080f000a",
}

// ZKMerkleTrieABI is the input ABI used to generate the binding from.
// Deprecated: Use ZKMerkleTrieMetaData.ABI instead.
var ZKMerkleTrieABI = ZKMerkleTrieMetaData.ABI

// ZKMerkleTrieBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZKMerkleTrieMetaData.Bin instead.
var ZKMerkleTrieBin = ZKMerkleTrieMetaData.Bin

// DeployZKMerkleTrie deploys a new Ethereum contract, binding an instance of ZKMerkleTrie to it.
func DeployZKMerkleTrie(auth *bind.TransactOpts, backend bind.ContractBackend, _poseidon2 common.Address) (common.Address, *types.Transaction, *ZKMerkleTrie, error) {
	parsed, err := ZKMerkleTrieMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZKMerkleTrieBin), backend, _poseidon2)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZKMerkleTrie{ZKMerkleTrieCaller: ZKMerkleTrieCaller{contract: contract}, ZKMerkleTrieTransactor: ZKMerkleTrieTransactor{contract: contract}, ZKMerkleTrieFilterer: ZKMerkleTrieFilterer{contract: contract}}, nil
}

// ZKMerkleTrie is an auto generated Go binding around an Ethereum contract.
type ZKMerkleTrie struct {
	ZKMerkleTrieCaller     // Read-only binding to the contract
	ZKMerkleTrieTransactor // Write-only binding to the contract
	ZKMerkleTrieFilterer   // Log filterer for contract events
}

// ZKMerkleTrieCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZKMerkleTrieCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKMerkleTrieTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZKMerkleTrieTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKMerkleTrieFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZKMerkleTrieFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKMerkleTrieSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZKMerkleTrieSession struct {
	Contract     *ZKMerkleTrie     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZKMerkleTrieCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZKMerkleTrieCallerSession struct {
	Contract *ZKMerkleTrieCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ZKMerkleTrieTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZKMerkleTrieTransactorSession struct {
	Contract     *ZKMerkleTrieTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ZKMerkleTrieRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZKMerkleTrieRaw struct {
	Contract *ZKMerkleTrie // Generic contract binding to access the raw methods on
}

// ZKMerkleTrieCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZKMerkleTrieCallerRaw struct {
	Contract *ZKMerkleTrieCaller // Generic read-only contract binding to access the raw methods on
}

// ZKMerkleTrieTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZKMerkleTrieTransactorRaw struct {
	Contract *ZKMerkleTrieTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZKMerkleTrie creates a new instance of ZKMerkleTrie, bound to a specific deployed contract.
func NewZKMerkleTrie(address common.Address, backend bind.ContractBackend) (*ZKMerkleTrie, error) {
	contract, err := bindZKMerkleTrie(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZKMerkleTrie{ZKMerkleTrieCaller: ZKMerkleTrieCaller{contract: contract}, ZKMerkleTrieTransactor: ZKMerkleTrieTransactor{contract: contract}, ZKMerkleTrieFilterer: ZKMerkleTrieFilterer{contract: contract}}, nil
}

// NewZKMerkleTrieCaller creates a new read-only instance of ZKMerkleTrie, bound to a specific deployed contract.
func NewZKMerkleTrieCaller(address common.Address, caller bind.ContractCaller) (*ZKMerkleTrieCaller, error) {
	contract, err := bindZKMerkleTrie(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZKMerkleTrieCaller{contract: contract}, nil
}

// NewZKMerkleTrieTransactor creates a new write-only instance of ZKMerkleTrie, bound to a specific deployed contract.
func NewZKMerkleTrieTransactor(address common.Address, transactor bind.ContractTransactor) (*ZKMerkleTrieTransactor, error) {
	contract, err := bindZKMerkleTrie(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZKMerkleTrieTransactor{contract: contract}, nil
}

// NewZKMerkleTrieFilterer creates a new log filterer instance of ZKMerkleTrie, bound to a specific deployed contract.
func NewZKMerkleTrieFilterer(address common.Address, filterer bind.ContractFilterer) (*ZKMerkleTrieFilterer, error) {
	contract, err := bindZKMerkleTrie(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZKMerkleTrieFilterer{contract: contract}, nil
}

// bindZKMerkleTrie binds a generic wrapper to an already deployed contract.
func bindZKMerkleTrie(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZKMerkleTrieMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKMerkleTrie *ZKMerkleTrieRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKMerkleTrie.Contract.ZKMerkleTrieCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKMerkleTrie *ZKMerkleTrieRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKMerkleTrie.Contract.ZKMerkleTrieTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKMerkleTrie *ZKMerkleTrieRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKMerkleTrie.Contract.ZKMerkleTrieTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKMerkleTrie *ZKMerkleTrieCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKMerkleTrie.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKMerkleTrie *ZKMerkleTrieTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKMerkleTrie.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKMerkleTrie *ZKMerkleTrieTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKMerkleTrie.Contract.contract.Transact(opts, method, params...)
}

// POSEIDON2 is a free data retrieval call binding the contract method 0xdc8b5038.
//
// Solidity: function POSEIDON2() view returns(address)
func (_ZKMerkleTrie *ZKMerkleTrieCaller) POSEIDON2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZKMerkleTrie.contract.Call(opts, &out, "POSEIDON2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// POSEIDON2 is a free data retrieval call binding the contract method 0xdc8b5038.
//
// Solidity: function POSEIDON2() view returns(address)
func (_ZKMerkleTrie *ZKMerkleTrieSession) POSEIDON2() (common.Address, error) {
	return _ZKMerkleTrie.Contract.POSEIDON2(&_ZKMerkleTrie.CallOpts)
}

// POSEIDON2 is a free data retrieval call binding the contract method 0xdc8b5038.
//
// Solidity: function POSEIDON2() view returns(address)
func (_ZKMerkleTrie *ZKMerkleTrieCallerSession) POSEIDON2() (common.Address, error) {
	return _ZKMerkleTrie.Contract.POSEIDON2(&_ZKMerkleTrie.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0xc423b1e8.
//
// Solidity: function get(bytes32 _key, bytes[] _proofs, bytes32 _root) view returns(bool, bytes)
func (_ZKMerkleTrie *ZKMerkleTrieCaller) Get(opts *bind.CallOpts, _key [32]byte, _proofs [][]byte, _root [32]byte) (bool, []byte, error) {
	var out []interface{}
	err := _ZKMerkleTrie.contract.Call(opts, &out, "get", _key, _proofs, _root)

	if err != nil {
		return *new(bool), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// Get is a free data retrieval call binding the contract method 0xc423b1e8.
//
// Solidity: function get(bytes32 _key, bytes[] _proofs, bytes32 _root) view returns(bool, bytes)
func (_ZKMerkleTrie *ZKMerkleTrieSession) Get(_key [32]byte, _proofs [][]byte, _root [32]byte) (bool, []byte, error) {
	return _ZKMerkleTrie.Contract.Get(&_ZKMerkleTrie.CallOpts, _key, _proofs, _root)
}

// Get is a free data retrieval call binding the contract method 0xc423b1e8.
//
// Solidity: function get(bytes32 _key, bytes[] _proofs, bytes32 _root) view returns(bool, bytes)
func (_ZKMerkleTrie *ZKMerkleTrieCallerSession) Get(_key [32]byte, _proofs [][]byte, _root [32]byte) (bool, []byte, error) {
	return _ZKMerkleTrie.Contract.Get(&_ZKMerkleTrie.CallOpts, _key, _proofs, _root)
}

// VerifyInclusionProof is a free data retrieval call binding the contract method 0x12e64a72.
//
// Solidity: function verifyInclusionProof(bytes32 _key, bytes _value, bytes[] _proofs, bytes32 _root) view returns(bool)
func (_ZKMerkleTrie *ZKMerkleTrieCaller) VerifyInclusionProof(opts *bind.CallOpts, _key [32]byte, _value []byte, _proofs [][]byte, _root [32]byte) (bool, error) {
	var out []interface{}
	err := _ZKMerkleTrie.contract.Call(opts, &out, "verifyInclusionProof", _key, _value, _proofs, _root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyInclusionProof is a free data retrieval call binding the contract method 0x12e64a72.
//
// Solidity: function verifyInclusionProof(bytes32 _key, bytes _value, bytes[] _proofs, bytes32 _root) view returns(bool)
func (_ZKMerkleTrie *ZKMerkleTrieSession) VerifyInclusionProof(_key [32]byte, _value []byte, _proofs [][]byte, _root [32]byte) (bool, error) {
	return _ZKMerkleTrie.Contract.VerifyInclusionProof(&_ZKMerkleTrie.CallOpts, _key, _value, _proofs, _root)
}

// VerifyInclusionProof is a free data retrieval call binding the contract method 0x12e64a72.
//
// Solidity: function verifyInclusionProof(bytes32 _key, bytes _value, bytes[] _proofs, bytes32 _root) view returns(bool)
func (_ZKMerkleTrie *ZKMerkleTrieCallerSession) VerifyInclusionProof(_key [32]byte, _value []byte, _proofs [][]byte, _root [32]byte) (bool, error) {
	return _ZKMerkleTrie.Contract.VerifyInclusionProof(&_ZKMerkleTrie.CallOpts, _key, _value, _proofs, _root)
}
