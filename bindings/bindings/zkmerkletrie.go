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
	Bin: "0x60a060405234801561001057600080fd5b50604051611bd6380380611bd683398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b608051611b306100a66000396000818160940152818161083701528181610a5501528181610ca9015261116a0152611b306000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806312e64a7214610046578063c423b1e81461006e578063dc8b50381461008f575b600080fd5b61005961005436600461177c565b6100db565b60405190151581526020015b60405180910390f35b61008161007c3660046117f1565b6101a6565b6040516100659291906118bb565b6100b67f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610065565b6040517fc423b1e800000000000000000000000000000000000000000000000000000000815260009081908190309063c423b1e890610122908a90899089906004016118d6565b600060405180830381865afa15801561013f573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526101859190810190611966565b9150915081801561019b575061019b86826107b8565b979650505050505050565b600060606002845110156102275760405162461bcd60e51b815260206004820152602960248201527f5a4b4d65726b6c65547269653a2070726f76696465642070726f6f662069732060448201527f746f6f2073686f7274000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b61027e84600186516102399190611a27565b8151811061024957610249611a3e565b602002602001015180516020909101207f950654da67865a81bc70e45f3230f5179f08e29c66184bf746f71050f117b3b81490565b6102f05760405162461bcd60e51b815260206004820152602d60248201527f5a4b4d65726b6c65547269653a20746865206c617374206974656d206973206e60448201527f6f74206d61676963206861736800000000000000000000000000000000000000606482015260840161021e565b60006102fb866107d4565b90506000610308866108b3565b90506103526040805161010081019091528060008152600060208201819052604082018190526060808301829052608083015260a0820181905260c0820181905260e09091015290565b60408051602081019091526000808252835190918291829190829061037990600290611a27565b90505b86818151811061038e5761038e611a3e565b6020026020010151602001519550600060038111156103af576103af611a6d565b865160038111156103c2576103c2611a6d565b036104a05760006103d389836109a6565b90508015610433578660200151861461042e5760405162461bcd60e51b815260206004820152601b60248201527f5a4b4d65726b6c65547269653a20696e76616c6964206b6579204c0000000000604482015260640161021e565b610486565b866040015186146104865760405162461bcd60e51b815260206004820152601b60248201527f5a4b4d65726b6c65547269653a20696e76616c6964206b657920520000000000604482015260640161021e565b61049887602001518860400151610a03565b955050610722565b6001865160038111156104b5576104b5611a6d565b0361068f57831561052e5760405162461bcd60e51b815260206004820152602260248201527f5a4b4d65726b6c65547269653a206475706c696361746564206c656166206e6f60448201527f6465000000000000000000000000000000000000000000000000000000000000606482015260840161021e565b60019350610556600160001b87606001516105518960a001518a60800151610ad0565b610c6c565b9450856080015151602061056a9190611a9c565b67ffffffffffffffff811115610582576105826115c3565b6040519080825280601f01601f1916602001820160405280156105ac576020820181803683370190505b509150602060005b8760800151518110156105f7576000886080015182815181106105d9576105d9611a3e565b602090810291909101810151868501529290920191506001016105b4565b5060e087015115610689578d8760e0015114806106175750888760e00151145b6106895760405162461bcd60e51b815260206004820152602260248201527f5a4b4d65726b6c65547269653a20696e76616c6964206b657920707265696d6160448201527f6765000000000000000000000000000000000000000000000000000000000000606482015260840161021e565b50610722565b6002865160038111156106a4576106a4611a6d565b0361072257821561071d5760405162461bcd60e51b815260206004820152602360248201527f5a4b4d65726b6c65547269653a206475706c69636174656420656d707479206e60448201527f6f64650000000000000000000000000000000000000000000000000000000000606482015260840161021e565b600192505b8060000361077e578a85146107795760405162461bcd60e51b815260206004820152601960248201527f5a4b4d65726b65547269653a20696e76616c696420726f6f7400000000000000604482015260640161021e565b6107a5565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161037c565b50919b919a509098505050505050505050565b6000818051906020012083805190602001201490505b92915050565b60008060006107e284610de4565b6040805180820182528381526020810183905290517f299e566000000000000000000000000000000000000000000000000000000000815292945090925073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169163299e56609161086a91600401611ad9565b602060405180830381865afa158015610887573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108ab9190611b0a565b949350505050565b805160609060008167ffffffffffffffff8111156108d3576108d36115c3565b60405190808252806020026020018201604052801561090c57816020015b6108f9611564565b8152602001906001900390816108f15790505b50905060005b61091d600184611a27565b81101561099e57600061094886838151811061093b5761093b611a3e565b6020026020010151610e0c565b9050604051806040016040528087848151811061096757610967611a3e565b602002602001015181526020018281525083838151811061098a5761098a611a3e565b602090810291909101015250600101610912565b509392505050565b600061010082106109f95760405162461bcd60e51b815260206004820152601c60248201527f5a4b4d65726b6c65547269653a20746f6f206c6f6e6720646570746800000000604482015260640161021e565b506001901b161590565b6040805180820182528381526020810183905290517f299e566000000000000000000000000000000000000000000000000000000000815260009173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169163299e566091610a8891600401611ad9565b602060405180830381865afa158015610aa5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ac99190611b0a565b9392505050565b6000600182511015610b4a5760405162461bcd60e51b815260206004820152602b60248201527f5a4b547269654861736865723a20746f6f206665772076616c75657320666f7260448201527f205f76616c756548617368000000000000000000000000000000000000000000606482015260840161021e565b6000825167ffffffffffffffff811115610b6657610b666115c3565b604051908082528060200260200182016040528015610b8f578160200160208202803683370190505b50905060005b8351811015610c35576001811b851663ffffffff1615610bf457610bd1848281518110610bc457610bc4611a3e565b60200260200101516107d4565b828281518110610be357610be3611a3e565b602002602001018181525050610c2d565b838181518110610c0657610c06611a3e565b6020026020010151828281518110610c2057610c20611a3e565b6020026020010181815250505b600101610b95565b50600283511015610c635780600081518110610c5357610c53611a3e565b60200260200101519150506107ce565b6108ab816110ec565b6040805180820182528481526020810184905290517f299e56600000000000000000000000000000000000000000000000000000000081526000917f00000000000000000000000000000000000000000000000000000000000000009173ffffffffffffffffffffffffffffffffffffffff83169163299e566091610cf49190600401611ad9565b602060405180830381865afa158015610d11573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d359190611b0a565b6040805180820182528281526020810186905290517f299e566000000000000000000000000000000000000000000000000000000000815291965073ffffffffffffffffffffffffffffffffffffffff83169163299e566091610d9a91600401611ad9565b602060405180830381865afa158015610db7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ddb9190611b0a565b95945050505050565b60008082608081901b610df78260801c90565b610e018260801c90565b935093505050915091565b60408051610100810182526000808252602082018190529181018290526060808201839052608082015260a0810182905260c0810182905260e081019190915260408051610100810182526000808252602082018190529181018290526060808201839052608082015260a0810182905260c0810182905260e08101919091526000610ec0846040805180820182526060815260006020918201528151808301909252828101825291519181019190915290565b90506000610ecd826112d6565b60ff16905080610ef857610ee08261135d565b6020840152610eee8261135d565b60408401526110a8565b60018103610fe257610f098261135d565b6060840152600080610f1a846113fb565b63ffffffff821660a088015290925060ff1690508067ffffffffffffffff811115610f4757610f476115c3565b604051908082528060200260200182016040528015610f70578160200160208202803683370190505b50608086015260005b81811015610fb357610f8a8561135d565b86608001518281518110610fa057610fa0611a3e565b6020908102919091010152600101610f79565b506000610fbf856112d6565b60ff1690508015610fda57610fd48582611495565b60e08701525b5050506110a8565b600281146110a857600381036110605760405162461bcd60e51b815260206004820152602560248201527f4e6f64655265616465723a20756e657870656374656420726f6f74206e6f646560448201527f2074797065000000000000000000000000000000000000000000000000000000606482015260840161021e565b60405162461bcd60e51b815260206004820152601d60248201527f4e6f64655265616465723a20696e76616c6964206e6f64652074797065000000604482015260640161021e565b8060038111156110ba576110ba611a6d565b839060038111156110cd576110cd611a6d565b908160038111156110e0576110e0611a6d565b90525091949350505050565b60006004825110156111665760405162461bcd60e51b815260206004820152602b60248201527f5a4b547269654861736865723a20746f6f206665772076616c75657320666f7260448201527f205f68617368456c656d73000000000000000000000000000000000000000000606482015260840161021e565b81517f00000000000000000000000000000000000000000000000000000000000000009060009081906001906002905b808310156112ae57600094505b808510156112a057828501935080841015611297578573ffffffffffffffffffffffffffffffffffffffff1663299e566060405180604001604052808b89815181106111f1576111f1611a3e565b602002602001015181526020018b888151811061121057611210611a3e565b60200260200101518152506040518263ffffffff1660e01b81526004016112379190611ad9565b602060405180830381865afa158015611254573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112789190611b0a565b88868151811061128a5761128a611a3e565b6020026020010181815250505b938101936111a3565b909150600182901b90611196565b876000815181106112c1576112c1611a3e565b60200260200101519650505050505050919050565b600060018260200151101561132d5760405162461bcd60e51b815260206004820152601f60248201527f4e6f64655265616465723a20746f6f2073686f727420666f722075696e743800604482015260640161021e565b81518051600180830180865260208601805191949360f81c9291611352908390611a27565b905250949350505050565b60006020826020015110156113da5760405162461bcd60e51b815260206004820152602160248201527f4e6f64655265616465723a20746f6f2073686f727420666f722062797465733360448201527f3200000000000000000000000000000000000000000000000000000000000000606482015260840161021e565b81518051602080830180865281860180519194939291611352908390611a27565b6000806004836020015110156114535760405162461bcd60e51b815260206004820181905260248201527f4e6f64655265616465723a20746f6f2073686f727420666f722075696e743332604482015260640161021e565b8251805160048083018087526020870180519194939260f084901c9260f885901c9290611481908390611a27565b90525060ff90911697909650945050505050565b600081836020015110156115115760405162461bcd60e51b815260206004820152602160248201527f4e6f64655265616465723a20746f6f2073686f727420666f72206e206279746560448201527f7300000000000000000000000000000000000000000000000000000000000000606482015260840161021e565b8251606090600080611524866008611a9c565b61153090610100611a27565b8351848801808a5260208a01805191975091831c94509192508791611556908390611a27565b905250909695505050505050565b6040518060400160405280606081526020016115be6040805161010081019091528060008152600060208201819052604082018190526060808301829052608083015260a0820181905260c0820181905260e09091015290565b905290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611639576116396115c3565b604052919050565b600067ffffffffffffffff82111561165b5761165b6115c3565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f83011261169857600080fd5b81356116ab6116a682611641565b6115f2565b8181528460208386010111156116c057600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f8301126116ee57600080fd5b8135602067ffffffffffffffff8083111561170b5761170b6115c3565b8260051b61171a8382016115f2565b938452858101830193838101908886111561173457600080fd5b84880192505b85831015611770578235848111156117525760008081fd5b6117608a87838c0101611687565b835250918401919084019061173a565b98975050505050505050565b6000806000806080858703121561179257600080fd5b84359350602085013567ffffffffffffffff808211156117b157600080fd5b6117bd88838901611687565b945060408701359150808211156117d357600080fd5b506117e0878288016116dd565b949793965093946060013593505050565b60008060006060848603121561180657600080fd5b83359250602084013567ffffffffffffffff81111561182457600080fd5b611830868287016116dd565b925050604084013590509250925092565b60005b8381101561185c578181015183820152602001611844565b8381111561186b576000848401525b50505050565b60008151808452611889816020860160208601611841565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b82151581526040602082015260006108ab6040830184611871565b600060608201858352602060608185015281865180845260808601915060808160051b870101935082880160005b82811015611950577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8088870301845261193e868351611871565b95509284019290840190600101611904565b5050505050604092909201929092529392505050565b6000806040838503121561197957600080fd5b8251801515811461198957600080fd5b602084015190925067ffffffffffffffff8111156119a657600080fd5b8301601f810185136119b757600080fd5b80516119c56116a682611641565b8181528660208385010111156119da57600080fd5b6119eb826020830160208601611841565b8093505050509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015611a3957611a396119f8565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611ad457611ad46119f8565b500290565b60408101818360005b6002811015611b01578151835260209283019290910190600101611ae2565b50505092915050565b600060208284031215611b1c57600080fd5b505191905056fea164736f6c634300080f000a",
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
