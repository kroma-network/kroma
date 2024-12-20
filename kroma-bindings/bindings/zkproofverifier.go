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

// ZKProofVerifierMetaData contains all meta data concerning the ZKProofVerifier contract.
var ZKProofVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_zkVerifier\",\"type\":\"address\",\"internalType\":\"contractZKVerifier\"},{\"name\":\"_dummyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_maxTxs\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_zkMerkleTrie\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_sp1Verifier\",\"type\":\"address\",\"internalType\":\"contractISP1Verifier\"},{\"name\":\"_zkVmProgramVKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"dummyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxTxs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sp1Verifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISP1Verifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyZkEvmProof\",\"inputs\":[{\"name\":\"_zkEvmProof\",\"type\":\"tuple\",\"internalType\":\"structTypes.ZkEvmProof\",\"components\":[{\"name\":\"publicInputProof\",\"type\":\"tuple\",\"internalType\":\"structTypes.PublicInputProof\",\"components\":[{\"name\":\"srcOutputRootProof\",\"type\":\"tuple\",\"internalType\":\"structTypes.OutputRootProof\",\"components\":[{\"name\":\"version\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messagePasserStorageRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"latestBlockhash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nextBlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dstOutputRootProof\",\"type\":\"tuple\",\"internalType\":\"structTypes.OutputRootProof\",\"components\":[{\"name\":\"version\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messagePasserStorageRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"latestBlockhash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nextBlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"publicInput\",\"type\":\"tuple\",\"internalType\":\"structTypes.PublicInput\",\"components\":[{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"baseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"transactionsRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"withdrawalsRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"txHashes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blobGasUsed\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"excessBlobGas\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"parentBeaconRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"rlps\",\"type\":\"tuple\",\"internalType\":\"structTypes.BlockHeaderRLP\",\"components\":[{\"name\":\"uncleHash\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"coinbase\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiptsRoot\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"logsBloom\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"difficulty\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasUsed\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"mixHash\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonce\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"l2ToL1MessagePasserBalance\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"l2ToL1MessagePasserCodeHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"merkleProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]},{\"name\":\"proof\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"pair\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"_storedSrcOutput\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_storedDstOutput\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"publicInputHash_\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyZkVmProof\",\"inputs\":[{\"name\":\"_zkVmProof\",\"type\":\"tuple\",\"internalType\":\"structTypes.ZkVmProof\",\"components\":[{\"name\":\"zkVmProgramVKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"publicValues\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proofBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"_storedSrcOutput\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_storedDstOutput\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_storedL1Head\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"publicInputHash_\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"zkMerkleTrie\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"zkVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractZKVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"zkVmProgramVKey\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"BlockHashMismatched\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BlockHashMismatchedBtwSrcAndDst\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DstOutputMatched\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInclusionProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPublicInput\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidZkProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidZkVmVKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SrcOutputMismatched\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StateRootMismatched\",\"inputs\":[]}]",
	Bin: "0x6101406040523480156200001257600080fd5b50604051620023bf380380620023bf83398101604081905262000035916200007e565b6001600160a01b0395861660805260a09490945260c092909252831660e0529091166101005261012052620000ef565b6001600160a01b03811681146200007b57600080fd5b50565b60008060008060008060c087890312156200009857600080fd5b8651620000a58162000065565b8096505060208701519450604087015193506060870151620000c78162000065565b6080880151909350620000da8162000065565b8092505060a087015190509295509295509295565b60805160a05160c05160e051610100516101205161224a620001756000396000818160aa0152818161041f015261055f01526000818160f2015261053801526000818161018201526108530152600081816101a80152818161094b01526109b90152600081816101e10152610985015260008181610207015261033a015261224a6000f3fe608060405234801561001057600080fd5b50600436106100a35760003560e01c80635747274b116100765780639a7ec1361161005b5780639a7ec136146101cc5780639aea2572146101df578063d6df096d1461020557600080fd5b80635747274b14610180578063816bf26d146101a657600080fd5b8063222ce122146100a85780633955d7a1146100dd57806352a07fa3146100f057806354fd4d5014610137575b600080fd5b7f00000000000000000000000000000000000000000000000000000000000000005b6040519081526020015b60405180910390f35b6100ca6100eb366004611504565b61022b565b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100d4565b6101736040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516100d491906115c8565b7f0000000000000000000000000000000000000000000000000000000000000000610112565b7f00000000000000000000000000000000000000000000000000000000000000006100ca565b6100ca6101da3660046115e2565b610419565b7f00000000000000000000000000000000000000000000000000000000000000006100ca565b7f0000000000000000000000000000000000000000000000000000000000000000610112565b6000366102388580611636565b905060808101356101008201351461027c576040517f3f126fab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6102b2848461029861029336869003860186611740565b610616565b6102ad61029336879003870160a08801611740565b6106b8565b6102da60a082016102c76101408401846117b0565b6102d56101608501856117e4565b610736565b6103066102eb6101c0830183611818565b6101808401356101a085013560e086013560c08701356107fb565b610321602082013561031c6101408401846117b0565b610945565b915073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016634292dc3e61036c6020880188611818565b61037960408a018a611818565b876040518663ffffffff1660e01b815260040161039a9594939291906118d6565b602060405180830381865afa1580156103b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103db9190611910565b610411576040517fe1ac453100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b509392505050565b600084357f000000000000000000000000000000000000000000000000000000000000000014610475576040517f2166766900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6104c484846104876020890189611932565b61049691602891600891611997565b61049f916119c1565b6104ac60208a018a611932565b6104bb91605091603091611997565b6102ad916119c1565b816104d26020870187611932565b6104e191607891605891611997565b6104ea916119c1565b14610521576040517f7458ca2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166341493c607f000000000000000000000000000000000000000000000000000000000000000061058b6020890189611932565b61059860408b018b611932565b6040518663ffffffff1660e01b81526004016105b8959493929190611a28565b60006040518083038186803b1580156105d057600080fd5b505afa1580156105e4573d6000803e3d6000fd5b506105f6925050506020860186611932565b604051610604929190611a61565b60405180910390209050949350505050565b60808101516000906106785781516020808401516040808601516060870151915161065b95949192910193845260208401929092526040830152606082015260800190565b604051602081830303815290604052805190602001209050919050565b81516020808401516040808601516060808801516080808a01518551978801989098529386019490945284015282015260a081019190915260c00161065b565b8184146106f1576040517f8b10302800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b821561073057808303610730576040517f4e15341500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50505050565b82602001358260e0013514610777576040517f4d9e774000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006101808301356107a25761079d61078f84611b09565b61079884611c81565b610a02565b6107bc565b6107bc6107ae84611b09565b6107b784611c81565b610aa6565b905080846060013514610730576040517fb033950600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60408051600060208201528082018690526060810185905260808082018590528251808303909101815260a08201928390527f12e64a72000000000000000000000000000000000000000000000000000000009092527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906312e64a72906108c5907f42000000000000000000000000000000000000030000000000000000000000009085908c908c90899060a401611df7565b602060405180830381865afa1580156108e2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109069190611910565b61093c576040517ff35959c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50505050505050565b600060607f0000000000000000000000000000000000000000000000000000000000000000610978610120850185611818565b905010156109e5576109e27f00000000000000000000000000000000000000000000000000000000000000006109b2610120860186611818565b6109dd91507f0000000000000000000000000000000000000000000000000000000000000000611f1f565b610bb7565b90505b6109f8846109f285611b09565b83610c3b565b9150505b92915050565b6040805160118082526102408201909252600091829190816020015b6060815260200190600190039081610a1e579050509050610a40848483610cad565b610a6f846101000151604051602001610a5b91815260200190565b604051602081830303815290604052610f60565b81601081518110610a8257610a82611f36565b6020026020010181905250610a9681610fcb565b8051906020012091505092915050565b6040805160148082526102a08201909252600091829190816020015b6060815260200190600190039081610ac2579050509050610ae4848483610cad565b610aff846101000151604051602001610a5b91815260200190565b81601081518110610b1257610b12611f36565b6020026020010181905250610b3584610140015167ffffffffffffffff16610ff6565b81601181518110610b4857610b48611f36565b6020026020010181905250610b6b84610160015167ffffffffffffffff16610ff6565b81601281518110610b7e57610b7e611f36565b6020026020010181905250610ba4846101800151604051602001610a5b91815260200190565b81601381518110610a8257610a82611f36565b606060008267ffffffffffffffff811115610bd457610bd4611674565b604051908082528060200260200182016040528015610bfd578160200160208202803683370190505b50905060005b838110156104115784828281518110610c1e57610c1e611f36565b602090810291909101015280610c3381611f65565b915050610c03565b6000838360e001516000801b85600001518660200151876060015188604001518960a001518a608001518b6101200151518c61012001518c604051602001610c8e9c9b9a99989796959493929190611fb2565b6040516020818303038152906040528051906020012090509392505050565b610cc78360200151604051602001610a5b91815260200190565b81600081518110610cda57610cda611f36565b6020026020010181905250816000015181600181518110610cfd57610cfd611f36565b6020026020010181905250816020015181600281518110610d2057610d20611f36565b6020026020010181905250610d458360e00151604051602001610a5b91815260200190565b81600381518110610d5857610d58611f36565b6020026020010181905250610d7d8360c00151604051602001610a5b91815260200190565b81600481518110610d9057610d90611f36565b6020026020010181905250816040015181600581518110610db357610db3611f36565b6020026020010181905250816060015181600681518110610dd657610dd6611f36565b6020026020010181905250816080015181600781518110610df957610df9611f36565b6020026020010181905250610e1b836060015167ffffffffffffffff16610ff6565b81600881518110610e2e57610e2e611f36565b6020026020010181905250610e50836080015167ffffffffffffffff16610ff6565b81600981518110610e6357610e63611f36565b60200260200101819052508160a0015181600a81518110610e8657610e86611f36565b6020026020010181905250610ea8836040015167ffffffffffffffff16610ff6565b81600b81518110610ebb57610ebb611f36565b60200260200101819052508160c0015181600c81518110610ede57610ede611f36565b60200260200101819052508160e0015181600d81518110610f0157610f01611f36565b602002602001018190525081610100015181600e81518110610f2557610f25611f36565b6020026020010181905250610f3d8360a00151610ff6565b81600f81518110610f5057610f50611f36565b6020026020010181905250505050565b606081516001148015610f8d5750608082600081518110610f8357610f83611f36565b016020015160f81c105b15610f96575090565b610fa282516080611009565b82604051602001610fb492919061206b565b60405160208183030381529060405290505b919050565b6060610fd6826111fd565b9050610fe4815160c0611009565b81604051602001610fb492919061206b565b60606109fc61100483611332565b610f60565b60606038831015611087576040805160018082528183019092529060208201818036833701905050905061103d828461209a565b60f81b8160008151811061105357611053611f36565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506109fc565b600060015b61109681866120ee565b156110bc57816110a581611f65565b92506110b5905061010082612102565b905061108c565b6110c7826001612121565b67ffffffffffffffff8111156110df576110df611674565b6040519080825280601f01601f191660200182016040528015611109576020820181803683370190505b509250611116848361209a565b61112190603761209a565b60f81b8360008151811061113757611137611f36565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600190505b8181116111f55761010061117f8284611f1f565b61118b9061010061221d565b61119590876120ee565b61119f9190612229565b60f81b8382815181106111b4576111b4611f36565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350806111ed81611f65565b91505061116b565b505092915050565b6060815160000361121c57505060408051600081526020810190915290565b6000805b83518110156112635783818151811061123b5761123b611f36565b6020026020010151518261124f9190612121565b91508061125b81611f65565b915050611220565b8167ffffffffffffffff81111561127c5761127c611674565b6040519080825280601f01601f1916602001820160405280156112a6576020820181803683370190505b50925060009050602083015b845182101561132a5760008583815181106112cf576112cf611f36565b6020026020010151905060006020820190506112ed8382845161148f565b8684815181106112ff576112ff611f36565b602002602001015151836113139190612121565b92505050818061132290611f65565b9250506112b2565b505050919050565b606060008260405160200161134991815260200190565b604051602081830303815290604052905060005b60208110156113b85781818151811061137857611378611f36565b01602001517fff00000000000000000000000000000000000000000000000000000000000000166000036113b857806113b081611f65565b91505061135d565b6113c3816020611f1f565b67ffffffffffffffff8111156113db576113db611674565b6040519080825280601f01601f191660200182016040528015611405576020820181803683370190505b50925060005b835181101561132a57828261141f81611f65565b93508151811061143157611431611f36565b602001015160f81c60f81b84828151811061144e5761144e611f36565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508061148781611f65565b91505061140b565b8282825b602081106114cb57815183526114aa602084612121565b92506114b7602083612121565b91506114c4602082611f1f565b9050611493565b905182516020929092036101000a6000190180199091169116179052505050565b6000606082840312156114fe57600080fd5b50919050565b60008060006060848603121561151957600080fd5b833567ffffffffffffffff81111561153057600080fd5b61153c868287016114ec565b9660208601359650604090950135949350505050565b60005b8381101561156d578181015183820152602001611555565b838111156107305750506000910152565b60008151808452611596816020860160208601611552565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006115db602083018461157e565b9392505050565b600080600080608085870312156115f857600080fd5b843567ffffffffffffffff81111561160f57600080fd5b61161b878288016114ec565b97602087013597506040870135966060013595509350505050565b600082357ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe2183360301811261166a57600080fd5b9190910192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516101a0810167ffffffffffffffff811182821017156116c7576116c7611674565b60405290565b604051610120810167ffffffffffffffff811182821017156116c7576116c7611674565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561173857611738611674565b604052919050565b600060a0828403121561175257600080fd5b60405160a0810181811067ffffffffffffffff8211171561177557611775611674565b806040525082358152602083013560208201526040830135604082015260608301356060820152608083013560808201528091505092915050565b600082357ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe6183360301811261166a57600080fd5b600082357ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee183360301811261166a57600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261184d57600080fd5b83018035915067ffffffffffffffff82111561186857600080fd5b6020019150600581901b360382131561188057600080fd5b9250929050565b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8311156118b957600080fd5b8260051b8083602087013760009401602001938452509192915050565b6060815260006118ea606083018789611887565b82810360208401526118fd818688611887565b9150508260408301529695505050505050565b60006020828403121561192257600080fd5b815180151581146115db57600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261196757600080fd5b83018035915067ffffffffffffffff82111561198257600080fd5b60200191503681900382131561188057600080fd5b600080858511156119a757600080fd5b838611156119b457600080fd5b5050820193919092039150565b803560208310156109fc57600019602084900360031b1b1692915050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b858152606060208201526000611a426060830186886119df565b8281036040840152611a558185876119df565b98975050505050505050565b8183823760009101908152919050565b803567ffffffffffffffff81168114610fc657600080fd5b600082601f830112611a9a57600080fd5b8135602067ffffffffffffffff821115611ab657611ab6611674565b8160051b611ac58282016116f1565b9283528481018201928281019087851115611adf57600080fd5b83870192505b84831015611afe57823582529183019190830190611ae5565b979650505050505050565b60006101a08236031215611b1c57600080fd5b611b246116a3565b8235815260208301356020820152611b3e60408401611a71565b6040820152611b4f60608401611a71565b6060820152611b6060808401611a71565b608082015260a083013560a082015260c083013560c082015260e083013560e08201526101008084013581830152506101208084013567ffffffffffffffff811115611bab57600080fd5b611bb736828701611a89565b828401525050610140611bcb818501611a71565b90820152610160611bdd848201611a71565b9082015261018092830135928101929092525090565b600082601f830112611c0457600080fd5b813567ffffffffffffffff811115611c1e57611c1e611674565b611c4f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016116f1565b818152846020838601011115611c6457600080fd5b816020850160208301376000918101602001919091529392505050565b60006101208236031215611c9457600080fd5b611c9c6116cd565b823567ffffffffffffffff80821115611cb457600080fd5b611cc036838701611bf3565b83526020850135915080821115611cd657600080fd5b611ce236838701611bf3565b60208401526040850135915080821115611cfb57600080fd5b611d0736838701611bf3565b60408401526060850135915080821115611d2057600080fd5b611d2c36838701611bf3565b60608401526080850135915080821115611d4557600080fd5b611d5136838701611bf3565b608084015260a0850135915080821115611d6a57600080fd5b611d7636838701611bf3565b60a084015260c0850135915080821115611d8f57600080fd5b611d9b36838701611bf3565b60c084015260e0850135915080821115611db457600080fd5b611dc036838701611bf3565b60e084015261010091508185013581811115611ddb57600080fd5b611de736828801611bf3565b8385015250505080915050919050565b85815260006020608081840152611e11608084018861157e565b8381036040850152858152818101600587901b820183018860005b89811015611ed7577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085840301845281357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18c3603018112611e8d57600080fd5b8b01868101903567ffffffffffffffff811115611ea957600080fd5b803603821315611eb857600080fd5b611ec38582846119df565b958801959450505090850190600101611e2c565b5050809450505050508260608301529695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015611f3157611f31611ef0565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006000198203611f7857611f78611ef0565b5060010190565b60008151602080840160005b83811015611fa757815187529582019590820190600101611f8b565b509495945050505050565b8c81528b60208201528a604082015289606082015288608082015260007fffffffffffffffff000000000000000000000000000000000000000000000000808a60c01b1660a0840152808960c01b1660a88401528760b0840152808760c01b1660d0840152507fffff0000000000000000000000000000000000000000000000000000000000008560f01b1660d883015261205961205360da840186611f7f565b84611f7f565b9e9d5050505050505050505050505050565b6000835161207d818460208801611552565b835190830190612091818360208801611552565b01949350505050565b600060ff821660ff84168060ff038211156120b7576120b7611ef0565b019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000826120fd576120fd6120bf565b500490565b600081600019048311821515161561211c5761211c611ef0565b500290565b6000821982111561213457612134611ef0565b500190565b600181815b8085111561217457816000190482111561215a5761215a611ef0565b8085161561216757918102915b93841c939080029061213e565b509250929050565b60008261218b575060016109fc565b81612198575060006109fc565b81600181146121ae57600281146121b8576121d4565b60019150506109fc565b60ff8411156121c9576121c9611ef0565b50506001821b6109fc565b5060208310610133831016604e8410600b84101617156121f7575081810a6109fc565b6122018383612139565b806000190482111561221557612215611ef0565b029392505050565b60006115db838361217c565b600082612238576122386120bf565b50069056fea164736f6c634300080f000a",
}

// ZKProofVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use ZKProofVerifierMetaData.ABI instead.
var ZKProofVerifierABI = ZKProofVerifierMetaData.ABI

// ZKProofVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZKProofVerifierMetaData.Bin instead.
var ZKProofVerifierBin = ZKProofVerifierMetaData.Bin

// DeployZKProofVerifier deploys a new Ethereum contract, binding an instance of ZKProofVerifier to it.
func DeployZKProofVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, _zkVerifier common.Address, _dummyHash [32]byte, _maxTxs *big.Int, _zkMerkleTrie common.Address, _sp1Verifier common.Address, _zkVmProgramVKey [32]byte) (common.Address, *types.Transaction, *ZKProofVerifier, error) {
	parsed, err := ZKProofVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZKProofVerifierBin), backend, _zkVerifier, _dummyHash, _maxTxs, _zkMerkleTrie, _sp1Verifier, _zkVmProgramVKey)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZKProofVerifier{ZKProofVerifierCaller: ZKProofVerifierCaller{contract: contract}, ZKProofVerifierTransactor: ZKProofVerifierTransactor{contract: contract}, ZKProofVerifierFilterer: ZKProofVerifierFilterer{contract: contract}}, nil
}

// ZKProofVerifier is an auto generated Go binding around an Ethereum contract.
type ZKProofVerifier struct {
	ZKProofVerifierCaller     // Read-only binding to the contract
	ZKProofVerifierTransactor // Write-only binding to the contract
	ZKProofVerifierFilterer   // Log filterer for contract events
}

// ZKProofVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZKProofVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKProofVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZKProofVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKProofVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZKProofVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKProofVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZKProofVerifierSession struct {
	Contract     *ZKProofVerifier  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZKProofVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZKProofVerifierCallerSession struct {
	Contract *ZKProofVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ZKProofVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZKProofVerifierTransactorSession struct {
	Contract     *ZKProofVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ZKProofVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZKProofVerifierRaw struct {
	Contract *ZKProofVerifier // Generic contract binding to access the raw methods on
}

// ZKProofVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZKProofVerifierCallerRaw struct {
	Contract *ZKProofVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// ZKProofVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZKProofVerifierTransactorRaw struct {
	Contract *ZKProofVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZKProofVerifier creates a new instance of ZKProofVerifier, bound to a specific deployed contract.
func NewZKProofVerifier(address common.Address, backend bind.ContractBackend) (*ZKProofVerifier, error) {
	contract, err := bindZKProofVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZKProofVerifier{ZKProofVerifierCaller: ZKProofVerifierCaller{contract: contract}, ZKProofVerifierTransactor: ZKProofVerifierTransactor{contract: contract}, ZKProofVerifierFilterer: ZKProofVerifierFilterer{contract: contract}}, nil
}

// NewZKProofVerifierCaller creates a new read-only instance of ZKProofVerifier, bound to a specific deployed contract.
func NewZKProofVerifierCaller(address common.Address, caller bind.ContractCaller) (*ZKProofVerifierCaller, error) {
	contract, err := bindZKProofVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZKProofVerifierCaller{contract: contract}, nil
}

// NewZKProofVerifierTransactor creates a new write-only instance of ZKProofVerifier, bound to a specific deployed contract.
func NewZKProofVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*ZKProofVerifierTransactor, error) {
	contract, err := bindZKProofVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZKProofVerifierTransactor{contract: contract}, nil
}

// NewZKProofVerifierFilterer creates a new log filterer instance of ZKProofVerifier, bound to a specific deployed contract.
func NewZKProofVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*ZKProofVerifierFilterer, error) {
	contract, err := bindZKProofVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZKProofVerifierFilterer{contract: contract}, nil
}

// bindZKProofVerifier binds a generic wrapper to an already deployed contract.
func bindZKProofVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZKProofVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKProofVerifier *ZKProofVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKProofVerifier.Contract.ZKProofVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKProofVerifier *ZKProofVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKProofVerifier.Contract.ZKProofVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKProofVerifier *ZKProofVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKProofVerifier.Contract.ZKProofVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKProofVerifier *ZKProofVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKProofVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKProofVerifier *ZKProofVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKProofVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKProofVerifier *ZKProofVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKProofVerifier.Contract.contract.Transact(opts, method, params...)
}

// DummyHash is a free data retrieval call binding the contract method 0x9aea2572.
//
// Solidity: function dummyHash() view returns(bytes32)
func (_ZKProofVerifier *ZKProofVerifierCaller) DummyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKProofVerifier.contract.Call(opts, &out, "dummyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DummyHash is a free data retrieval call binding the contract method 0x9aea2572.
//
// Solidity: function dummyHash() view returns(bytes32)
func (_ZKProofVerifier *ZKProofVerifierSession) DummyHash() ([32]byte, error) {
	return _ZKProofVerifier.Contract.DummyHash(&_ZKProofVerifier.CallOpts)
}

// DummyHash is a free data retrieval call binding the contract method 0x9aea2572.
//
// Solidity: function dummyHash() view returns(bytes32)
func (_ZKProofVerifier *ZKProofVerifierCallerSession) DummyHash() ([32]byte, error) {
	return _ZKProofVerifier.Contract.DummyHash(&_ZKProofVerifier.CallOpts)
}

// MaxTxs is a free data retrieval call binding the contract method 0x816bf26d.
//
// Solidity: function maxTxs() view returns(uint256)
func (_ZKProofVerifier *ZKProofVerifierCaller) MaxTxs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZKProofVerifier.contract.Call(opts, &out, "maxTxs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTxs is a free data retrieval call binding the contract method 0x816bf26d.
//
// Solidity: function maxTxs() view returns(uint256)
func (_ZKProofVerifier *ZKProofVerifierSession) MaxTxs() (*big.Int, error) {
	return _ZKProofVerifier.Contract.MaxTxs(&_ZKProofVerifier.CallOpts)
}

// MaxTxs is a free data retrieval call binding the contract method 0x816bf26d.
//
// Solidity: function maxTxs() view returns(uint256)
func (_ZKProofVerifier *ZKProofVerifierCallerSession) MaxTxs() (*big.Int, error) {
	return _ZKProofVerifier.Contract.MaxTxs(&_ZKProofVerifier.CallOpts)
}

// Sp1Verifier is a free data retrieval call binding the contract method 0x52a07fa3.
//
// Solidity: function sp1Verifier() view returns(address)
func (_ZKProofVerifier *ZKProofVerifierCaller) Sp1Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZKProofVerifier.contract.Call(opts, &out, "sp1Verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sp1Verifier is a free data retrieval call binding the contract method 0x52a07fa3.
//
// Solidity: function sp1Verifier() view returns(address)
func (_ZKProofVerifier *ZKProofVerifierSession) Sp1Verifier() (common.Address, error) {
	return _ZKProofVerifier.Contract.Sp1Verifier(&_ZKProofVerifier.CallOpts)
}

// Sp1Verifier is a free data retrieval call binding the contract method 0x52a07fa3.
//
// Solidity: function sp1Verifier() view returns(address)
func (_ZKProofVerifier *ZKProofVerifierCallerSession) Sp1Verifier() (common.Address, error) {
	return _ZKProofVerifier.Contract.Sp1Verifier(&_ZKProofVerifier.CallOpts)
}

// VerifyZkEvmProof is a free data retrieval call binding the contract method 0x3955d7a1.
//
// Solidity: function verifyZkEvmProof((((bytes32,bytes32,bytes32,bytes32,bytes32),(bytes32,bytes32,bytes32,bytes32,bytes32),(bytes32,bytes32,uint64,uint64,uint64,uint256,bytes32,bytes32,bytes32,bytes32[],uint64,uint64,bytes32),(bytes,bytes,bytes,bytes,bytes,bytes,bytes,bytes,bytes),bytes32,bytes32,bytes[]),uint256[],uint256[]) _zkEvmProof, bytes32 _storedSrcOutput, bytes32 _storedDstOutput) view returns(bytes32 publicInputHash_)
func (_ZKProofVerifier *ZKProofVerifierCaller) VerifyZkEvmProof(opts *bind.CallOpts, _zkEvmProof TypesZkEvmProof, _storedSrcOutput [32]byte, _storedDstOutput [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ZKProofVerifier.contract.Call(opts, &out, "verifyZkEvmProof", _zkEvmProof, _storedSrcOutput, _storedDstOutput)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VerifyZkEvmProof is a free data retrieval call binding the contract method 0x3955d7a1.
//
// Solidity: function verifyZkEvmProof((((bytes32,bytes32,bytes32,bytes32,bytes32),(bytes32,bytes32,bytes32,bytes32,bytes32),(bytes32,bytes32,uint64,uint64,uint64,uint256,bytes32,bytes32,bytes32,bytes32[],uint64,uint64,bytes32),(bytes,bytes,bytes,bytes,bytes,bytes,bytes,bytes,bytes),bytes32,bytes32,bytes[]),uint256[],uint256[]) _zkEvmProof, bytes32 _storedSrcOutput, bytes32 _storedDstOutput) view returns(bytes32 publicInputHash_)
func (_ZKProofVerifier *ZKProofVerifierSession) VerifyZkEvmProof(_zkEvmProof TypesZkEvmProof, _storedSrcOutput [32]byte, _storedDstOutput [32]byte) ([32]byte, error) {
	return _ZKProofVerifier.Contract.VerifyZkEvmProof(&_ZKProofVerifier.CallOpts, _zkEvmProof, _storedSrcOutput, _storedDstOutput)
}

// VerifyZkEvmProof is a free data retrieval call binding the contract method 0x3955d7a1.
//
// Solidity: function verifyZkEvmProof((((bytes32,bytes32,bytes32,bytes32,bytes32),(bytes32,bytes32,bytes32,bytes32,bytes32),(bytes32,bytes32,uint64,uint64,uint64,uint256,bytes32,bytes32,bytes32,bytes32[],uint64,uint64,bytes32),(bytes,bytes,bytes,bytes,bytes,bytes,bytes,bytes,bytes),bytes32,bytes32,bytes[]),uint256[],uint256[]) _zkEvmProof, bytes32 _storedSrcOutput, bytes32 _storedDstOutput) view returns(bytes32 publicInputHash_)
func (_ZKProofVerifier *ZKProofVerifierCallerSession) VerifyZkEvmProof(_zkEvmProof TypesZkEvmProof, _storedSrcOutput [32]byte, _storedDstOutput [32]byte) ([32]byte, error) {
	return _ZKProofVerifier.Contract.VerifyZkEvmProof(&_ZKProofVerifier.CallOpts, _zkEvmProof, _storedSrcOutput, _storedDstOutput)
}

// VerifyZkVmProof is a free data retrieval call binding the contract method 0x9a7ec136.
//
// Solidity: function verifyZkVmProof((bytes32,bytes,bytes) _zkVmProof, bytes32 _storedSrcOutput, bytes32 _storedDstOutput, bytes32 _storedL1Head) view returns(bytes32 publicInputHash_)
func (_ZKProofVerifier *ZKProofVerifierCaller) VerifyZkVmProof(opts *bind.CallOpts, _zkVmProof TypesZkVmProof, _storedSrcOutput [32]byte, _storedDstOutput [32]byte, _storedL1Head [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ZKProofVerifier.contract.Call(opts, &out, "verifyZkVmProof", _zkVmProof, _storedSrcOutput, _storedDstOutput, _storedL1Head)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VerifyZkVmProof is a free data retrieval call binding the contract method 0x9a7ec136.
//
// Solidity: function verifyZkVmProof((bytes32,bytes,bytes) _zkVmProof, bytes32 _storedSrcOutput, bytes32 _storedDstOutput, bytes32 _storedL1Head) view returns(bytes32 publicInputHash_)
func (_ZKProofVerifier *ZKProofVerifierSession) VerifyZkVmProof(_zkVmProof TypesZkVmProof, _storedSrcOutput [32]byte, _storedDstOutput [32]byte, _storedL1Head [32]byte) ([32]byte, error) {
	return _ZKProofVerifier.Contract.VerifyZkVmProof(&_ZKProofVerifier.CallOpts, _zkVmProof, _storedSrcOutput, _storedDstOutput, _storedL1Head)
}

// VerifyZkVmProof is a free data retrieval call binding the contract method 0x9a7ec136.
//
// Solidity: function verifyZkVmProof((bytes32,bytes,bytes) _zkVmProof, bytes32 _storedSrcOutput, bytes32 _storedDstOutput, bytes32 _storedL1Head) view returns(bytes32 publicInputHash_)
func (_ZKProofVerifier *ZKProofVerifierCallerSession) VerifyZkVmProof(_zkVmProof TypesZkVmProof, _storedSrcOutput [32]byte, _storedDstOutput [32]byte, _storedL1Head [32]byte) ([32]byte, error) {
	return _ZKProofVerifier.Contract.VerifyZkVmProof(&_ZKProofVerifier.CallOpts, _zkVmProof, _storedSrcOutput, _storedDstOutput, _storedL1Head)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ZKProofVerifier *ZKProofVerifierCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZKProofVerifier.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ZKProofVerifier *ZKProofVerifierSession) Version() (string, error) {
	return _ZKProofVerifier.Contract.Version(&_ZKProofVerifier.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ZKProofVerifier *ZKProofVerifierCallerSession) Version() (string, error) {
	return _ZKProofVerifier.Contract.Version(&_ZKProofVerifier.CallOpts)
}

// ZkMerkleTrie is a free data retrieval call binding the contract method 0x5747274b.
//
// Solidity: function zkMerkleTrie() view returns(address)
func (_ZKProofVerifier *ZKProofVerifierCaller) ZkMerkleTrie(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZKProofVerifier.contract.Call(opts, &out, "zkMerkleTrie")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZkMerkleTrie is a free data retrieval call binding the contract method 0x5747274b.
//
// Solidity: function zkMerkleTrie() view returns(address)
func (_ZKProofVerifier *ZKProofVerifierSession) ZkMerkleTrie() (common.Address, error) {
	return _ZKProofVerifier.Contract.ZkMerkleTrie(&_ZKProofVerifier.CallOpts)
}

// ZkMerkleTrie is a free data retrieval call binding the contract method 0x5747274b.
//
// Solidity: function zkMerkleTrie() view returns(address)
func (_ZKProofVerifier *ZKProofVerifierCallerSession) ZkMerkleTrie() (common.Address, error) {
	return _ZKProofVerifier.Contract.ZkMerkleTrie(&_ZKProofVerifier.CallOpts)
}

// ZkVerifier is a free data retrieval call binding the contract method 0xd6df096d.
//
// Solidity: function zkVerifier() view returns(address)
func (_ZKProofVerifier *ZKProofVerifierCaller) ZkVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZKProofVerifier.contract.Call(opts, &out, "zkVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZkVerifier is a free data retrieval call binding the contract method 0xd6df096d.
//
// Solidity: function zkVerifier() view returns(address)
func (_ZKProofVerifier *ZKProofVerifierSession) ZkVerifier() (common.Address, error) {
	return _ZKProofVerifier.Contract.ZkVerifier(&_ZKProofVerifier.CallOpts)
}

// ZkVerifier is a free data retrieval call binding the contract method 0xd6df096d.
//
// Solidity: function zkVerifier() view returns(address)
func (_ZKProofVerifier *ZKProofVerifierCallerSession) ZkVerifier() (common.Address, error) {
	return _ZKProofVerifier.Contract.ZkVerifier(&_ZKProofVerifier.CallOpts)
}

// ZkVmProgramVKey is a free data retrieval call binding the contract method 0x222ce122.
//
// Solidity: function zkVmProgramVKey() view returns(bytes32)
func (_ZKProofVerifier *ZKProofVerifierCaller) ZkVmProgramVKey(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKProofVerifier.contract.Call(opts, &out, "zkVmProgramVKey")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ZkVmProgramVKey is a free data retrieval call binding the contract method 0x222ce122.
//
// Solidity: function zkVmProgramVKey() view returns(bytes32)
func (_ZKProofVerifier *ZKProofVerifierSession) ZkVmProgramVKey() ([32]byte, error) {
	return _ZKProofVerifier.Contract.ZkVmProgramVKey(&_ZKProofVerifier.CallOpts)
}

// ZkVmProgramVKey is a free data retrieval call binding the contract method 0x222ce122.
//
// Solidity: function zkVmProgramVKey() view returns(bytes32)
func (_ZKProofVerifier *ZKProofVerifierCallerSession) ZkVmProgramVKey() ([32]byte, error) {
	return _ZKProofVerifier.Contract.ZkVmProgramVKey(&_ZKProofVerifier.CallOpts)
}
