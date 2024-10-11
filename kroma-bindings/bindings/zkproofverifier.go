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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_zkVerifier\",\"type\":\"address\",\"internalType\":\"contractZKVerifier\"},{\"name\":\"_dummyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_maxTxs\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_zkMerkleTrie\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_sp1Verifier\",\"type\":\"address\",\"internalType\":\"contractISP1Verifier\"},{\"name\":\"_zkVmProgramVKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"dummyHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxTxs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sp1Verifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISP1Verifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyZkEvmProof\",\"inputs\":[{\"name\":\"_zkEvmProof\",\"type\":\"tuple\",\"internalType\":\"structTypes.ZkEvmProof\",\"components\":[{\"name\":\"publicInputProof\",\"type\":\"tuple\",\"internalType\":\"structTypes.PublicInputProof\",\"components\":[{\"name\":\"srcOutputRootProof\",\"type\":\"tuple\",\"internalType\":\"structTypes.OutputRootProof\",\"components\":[{\"name\":\"version\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messagePasserStorageRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"latestBlockhash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nextBlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dstOutputRootProof\",\"type\":\"tuple\",\"internalType\":\"structTypes.OutputRootProof\",\"components\":[{\"name\":\"version\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messagePasserStorageRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"latestBlockhash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nextBlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"publicInput\",\"type\":\"tuple\",\"internalType\":\"structTypes.PublicInput\",\"components\":[{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"baseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"transactionsRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"withdrawalsRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"txHashes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blobGasUsed\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"excessBlobGas\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"parentBeaconRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"rlps\",\"type\":\"tuple\",\"internalType\":\"structTypes.BlockHeaderRLP\",\"components\":[{\"name\":\"uncleHash\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"coinbase\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiptsRoot\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"logsBloom\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"difficulty\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasUsed\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"mixHash\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonce\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"l2ToL1MessagePasserBalance\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"l2ToL1MessagePasserCodeHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"merkleProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]},{\"name\":\"proof\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"pair\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"_storedSrcOutput\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_storedDstOutput\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"publicInputHash_\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyZkVmProof\",\"inputs\":[{\"name\":\"_zkVmProof\",\"type\":\"tuple\",\"internalType\":\"structTypes.ZkVmProof\",\"components\":[{\"name\":\"publicValues\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proofBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"_storedSrcOutput\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_storedDstOutput\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_storedL1Head\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"publicInputHash_\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"zkMerkleTrie\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"zkVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractZKVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"zkVmProgramVKey\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"BlockHashMismatched\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BlockHashMismatchedBtwSrcAndDst\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DstOutputMatched\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInclusionProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPublicInput\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidZkProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SrcOutputMismatched\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StateRootMismatched\",\"inputs\":[]}]",
	Bin: "0x6101406040523480156200001257600080fd5b50604051620023433803806200234383398101604081905262000035916200007e565b6001600160a01b0395861660805260a09490945260c092909252831660e0529091166101005261012052620000ef565b6001600160a01b03811681146200007b57600080fd5b50565b60008060008060008060c087890312156200009857600080fd5b8651620000a58162000065565b8096505060208701519450604087015193506060870151620000c78162000065565b6080880151909350620000da8162000065565b8092505060a087015190509295509295509295565b60805160a05160c05160e05161010051610120516121d56200016e6000396000818160aa01526104fc01526000818160f201526104d501526000818161018201526107ec0152600081816101a8015281816108e401526109520152600081816101ce015261091e0152600081816101f4015261033a01526121d56000f3fe608060405234801561001057600080fd5b50600436106100a35760003560e01c80635747274b116100765780639aea25721161005b5780639aea2572146101cc578063d6df096d146101f2578063e674f0e01461021857600080fd5b80635747274b14610180578063816bf26d146101a657600080fd5b8063222ce122146100a85780633955d7a1146100dd57806352a07fa3146100f057806354fd4d5014610137575b600080fd5b7f00000000000000000000000000000000000000000000000000000000000000005b6040519081526020015b60405180910390f35b6100ca6100eb366004611485565b61022b565b7f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100d4565b6101736040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516100d4919061154e565b7f0000000000000000000000000000000000000000000000000000000000000000610112565b7f00000000000000000000000000000000000000000000000000000000000000006100ca565b7f00000000000000000000000000000000000000000000000000000000000000006100ca565b7f0000000000000000000000000000000000000000000000000000000000000000610112565b6100ca610226366004611568565b610419565b60003661023885806115c1565b905060808101356101008201351461027c576040517f3f126fab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6102b28484610298610293368690038601866116cb565b6105af565b6102ad61029336879003870160a088016116cb565b610651565b6102da60a082016102c761014084018461173b565b6102d561016085018561176f565b6106cf565b6103066102eb6101c08301836117a3565b6101808401356101a085013560e086013560c0870135610794565b610321602082013561031c61014084018461173b565b6108de565b915073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016634292dc3e61036c60208801886117a3565b61037960408a018a6117a3565b876040518663ffffffff1660e01b815260040161039a959493929190611861565b602060405180830381865afa1580156103b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103db919061189b565b610411576040517fe1ac453100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b509392505050565b6000610464848461042a88806118bd565b61043991602091600091611922565b6104429161194c565b61044c89806118bd565b61045b91604091602091611922565b6102ad9161194c565b8161046f86806118bd565b61047e91606091604091611922565b6104879161194c565b146104be576040517f7458ca2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166341493c607f000000000000000000000000000000000000000000000000000000000000000061052588806118bd565b61053260208b018b6118bd565b6040518663ffffffff1660e01b81526004016105529594939291906119b3565b60006040518083038186803b15801561056a57600080fd5b505afa15801561057e573d6000803e3d6000fd5b5061058f92508791508190506118bd565b60405161059d9291906119ec565b60405180910390209050949350505050565b6080810151600090610611578151602080840151604080860151606087015191516105f495949192910193845260208401929092526040830152606082015260800190565b604051602081830303815290604052805190602001209050919050565b81516020808401516040808601516060808801516080808a01518551978801989098529386019490945284015282015260a081019190915260c0016105f4565b81841461068a576040517f8b10302800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82156106c9578083036106c9576040517f4e15341500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50505050565b82602001358260e0013514610710576040517f4d9e774000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061018083013561073b5761073661072884611a94565b61073184611c0c565b61099b565b610755565b61075561074784611a94565b61075084611c0c565b610a3f565b9050808460600135146106c9576040517fb033950600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60408051600060208201528082018690526060810185905260808082018590528251808303909101815260a08201928390527f12e64a72000000000000000000000000000000000000000000000000000000009092527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906312e64a729061085e907f42000000000000000000000000000000000000030000000000000000000000009085908c908c90899060a401611d82565b602060405180830381865afa15801561087b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061089f919061189b565b6108d5576040517ff35959c000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50505050505050565b600060607f00000000000000000000000000000000000000000000000000000000000000006109116101208501856117a3565b9050101561097e5761097b7f000000000000000000000000000000000000000000000000000000000000000061094b6101208601866117a3565b61097691507f0000000000000000000000000000000000000000000000000000000000000000611eaa565b610b50565b90505b6109918461098b85611a94565b83610bd4565b9150505b92915050565b6040805160118082526102408201909252600091829190816020015b60608152602001906001900390816109b75790505090506109d9848483610c46565b610a088461010001516040516020016109f491815260200190565b604051602081830303815290604052610ef9565b81601081518110610a1b57610a1b611ec1565b6020026020010181905250610a2f81610f64565b8051906020012091505092915050565b6040805160148082526102a08201909252600091829190816020015b6060815260200190600190039081610a5b579050509050610a7d848483610c46565b610a988461010001516040516020016109f491815260200190565b81601081518110610aab57610aab611ec1565b6020026020010181905250610ace84610140015167ffffffffffffffff16610f8f565b81601181518110610ae157610ae1611ec1565b6020026020010181905250610b0484610160015167ffffffffffffffff16610f8f565b81601281518110610b1757610b17611ec1565b6020026020010181905250610b3d8461018001516040516020016109f491815260200190565b81601381518110610a1b57610a1b611ec1565b606060008267ffffffffffffffff811115610b6d57610b6d6115ff565b604051908082528060200260200182016040528015610b96578160200160208202803683370190505b50905060005b838110156104115784828281518110610bb757610bb7611ec1565b602090810291909101015280610bcc81611ef0565b915050610b9c565b6000838360e001516000801b85600001518660200151876060015188604001518960a001518a608001518b6101200151518c61012001518c604051602001610c279c9b9a99989796959493929190611f3d565b6040516020818303038152906040528051906020012090509392505050565b610c6083602001516040516020016109f491815260200190565b81600081518110610c7357610c73611ec1565b6020026020010181905250816000015181600181518110610c9657610c96611ec1565b6020026020010181905250816020015181600281518110610cb957610cb9611ec1565b6020026020010181905250610cde8360e001516040516020016109f491815260200190565b81600381518110610cf157610cf1611ec1565b6020026020010181905250610d168360c001516040516020016109f491815260200190565b81600481518110610d2957610d29611ec1565b6020026020010181905250816040015181600581518110610d4c57610d4c611ec1565b6020026020010181905250816060015181600681518110610d6f57610d6f611ec1565b6020026020010181905250816080015181600781518110610d9257610d92611ec1565b6020026020010181905250610db4836060015167ffffffffffffffff16610f8f565b81600881518110610dc757610dc7611ec1565b6020026020010181905250610de9836080015167ffffffffffffffff16610f8f565b81600981518110610dfc57610dfc611ec1565b60200260200101819052508160a0015181600a81518110610e1f57610e1f611ec1565b6020026020010181905250610e41836040015167ffffffffffffffff16610f8f565b81600b81518110610e5457610e54611ec1565b60200260200101819052508160c0015181600c81518110610e7757610e77611ec1565b60200260200101819052508160e0015181600d81518110610e9a57610e9a611ec1565b602002602001018190525081610100015181600e81518110610ebe57610ebe611ec1565b6020026020010181905250610ed68360a00151610f8f565b81600f81518110610ee957610ee9611ec1565b6020026020010181905250505050565b606081516001148015610f265750608082600081518110610f1c57610f1c611ec1565b016020015160f81c105b15610f2f575090565b610f3b82516080610fa2565b82604051602001610f4d929190611ff6565b60405160208183030381529060405290505b919050565b6060610f6f82611196565b9050610f7d815160c0610fa2565b81604051602001610f4d929190611ff6565b6060610995610f9d836112cb565b610ef9565b606060388310156110205760408051600180825281830190925290602082018180368337019050509050610fd68284612025565b60f81b81600081518110610fec57610fec611ec1565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610995565b600060015b61102f8186612079565b15611055578161103e81611ef0565b925061104e90506101008261208d565b9050611025565b6110608260016120ac565b67ffffffffffffffff811115611078576110786115ff565b6040519080825280601f01601f1916602001820160405280156110a2576020820181803683370190505b5092506110af8483612025565b6110ba906037612025565b60f81b836000815181106110d0576110d0611ec1565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600190505b81811161118e576101006111188284611eaa565b611124906101006121a8565b61112e9087612079565b61113891906121b4565b60f81b83828151811061114d5761114d611ec1565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508061118681611ef0565b915050611104565b505092915050565b606081516000036111b557505060408051600081526020810190915290565b6000805b83518110156111fc578381815181106111d4576111d4611ec1565b602002602001015151826111e891906120ac565b9150806111f481611ef0565b9150506111b9565b8167ffffffffffffffff811115611215576112156115ff565b6040519080825280601f01601f19166020018201604052801561123f576020820181803683370190505b50925060009050602083015b84518210156112c357600085838151811061126857611268611ec1565b60200260200101519050600060208201905061128683828451611428565b86848151811061129857611298611ec1565b602002602001015151836112ac91906120ac565b9250505081806112bb90611ef0565b92505061124b565b505050919050565b60606000826040516020016112e291815260200190565b604051602081830303815290604052905060005b60208110156113515781818151811061131157611311611ec1565b01602001517fff0000000000000000000000000000000000000000000000000000000000000016600003611351578061134981611ef0565b9150506112f6565b61135c816020611eaa565b67ffffffffffffffff811115611374576113746115ff565b6040519080825280601f01601f19166020018201604052801561139e576020820181803683370190505b50925060005b83518110156112c35782826113b881611ef0565b9350815181106113ca576113ca611ec1565b602001015160f81c60f81b8482815181106113e7576113e7611ec1565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508061142081611ef0565b9150506113a4565b8282825b6020811061146457815183526114436020846120ac565b92506114506020836120ac565b915061145d602082611eaa565b905061142c565b905182516020929092036101000a6000190180199091169116179052505050565b60008060006060848603121561149a57600080fd5b833567ffffffffffffffff8111156114b157600080fd5b8401606081870312156114c357600080fd5b95602085013595506040909401359392505050565b60005b838110156114f35781810151838201526020016114db565b838111156106c95750506000910152565b6000815180845261151c8160208601602086016114d8565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006115616020830184611504565b9392505050565b6000806000806080858703121561157e57600080fd5b843567ffffffffffffffff81111561159557600080fd5b8501604081880312156115a757600080fd5b966020860135965060408601359560600135945092505050565b600082357ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe218336030181126115f557600080fd5b9190910192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516101a0810167ffffffffffffffff81118282101715611652576116526115ff565b60405290565b604051610120810167ffffffffffffffff81118282101715611652576116526115ff565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156116c3576116c36115ff565b604052919050565b600060a082840312156116dd57600080fd5b60405160a0810181811067ffffffffffffffff82111715611700576117006115ff565b806040525082358152602083013560208201526040830135604082015260608301356060820152608083013560808201528091505092915050565b600082357ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe618336030181126115f557600080fd5b600082357ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffee18336030181126115f557600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126117d857600080fd5b83018035915067ffffffffffffffff8211156117f357600080fd5b6020019150600581901b360382131561180b57600080fd5b9250929050565b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83111561184457600080fd5b8260051b8083602087013760009401602001938452509192915050565b606081526000611875606083018789611812565b8281036020840152611888818688611812565b9150508260408301529695505050505050565b6000602082840312156118ad57600080fd5b8151801515811461156157600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126118f257600080fd5b83018035915067ffffffffffffffff82111561190d57600080fd5b60200191503681900382131561180b57600080fd5b6000808585111561193257600080fd5b8386111561193f57600080fd5b5050820193919092039150565b8035602083101561099557600019602084900360031b1b1692915050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b8581526060602082015260006119cd60608301868861196a565b82810360408401526119e081858761196a565b98975050505050505050565b8183823760009101908152919050565b803567ffffffffffffffff81168114610f5f57600080fd5b600082601f830112611a2557600080fd5b8135602067ffffffffffffffff821115611a4157611a416115ff565b8160051b611a5082820161167c565b9283528481018201928281019087851115611a6a57600080fd5b83870192505b84831015611a8957823582529183019190830190611a70565b979650505050505050565b60006101a08236031215611aa757600080fd5b611aaf61162e565b8235815260208301356020820152611ac9604084016119fc565b6040820152611ada606084016119fc565b6060820152611aeb608084016119fc565b608082015260a083013560a082015260c083013560c082015260e083013560e08201526101008084013581830152506101208084013567ffffffffffffffff811115611b3657600080fd5b611b4236828701611a14565b828401525050610140611b568185016119fc565b90820152610160611b688482016119fc565b9082015261018092830135928101929092525090565b600082601f830112611b8f57600080fd5b813567ffffffffffffffff811115611ba957611ba96115ff565b611bda60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8401160161167c565b818152846020838601011115611bef57600080fd5b816020850160208301376000918101602001919091529392505050565b60006101208236031215611c1f57600080fd5b611c27611658565b823567ffffffffffffffff80821115611c3f57600080fd5b611c4b36838701611b7e565b83526020850135915080821115611c6157600080fd5b611c6d36838701611b7e565b60208401526040850135915080821115611c8657600080fd5b611c9236838701611b7e565b60408401526060850135915080821115611cab57600080fd5b611cb736838701611b7e565b60608401526080850135915080821115611cd057600080fd5b611cdc36838701611b7e565b608084015260a0850135915080821115611cf557600080fd5b611d0136838701611b7e565b60a084015260c0850135915080821115611d1a57600080fd5b611d2636838701611b7e565b60c084015260e0850135915080821115611d3f57600080fd5b611d4b36838701611b7e565b60e084015261010091508185013581811115611d6657600080fd5b611d7236828801611b7e565b8385015250505080915050919050565b85815260006020608081840152611d9c6080840188611504565b8381036040850152858152818101600587901b820183018860005b89811015611e62577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085840301845281357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18c3603018112611e1857600080fd5b8b01868101903567ffffffffffffffff811115611e3457600080fd5b803603821315611e4357600080fd5b611e4e85828461196a565b958801959450505090850190600101611db7565b5050809450505050508260608301529695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015611ebc57611ebc611e7b565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006000198203611f0357611f03611e7b565b5060010190565b60008151602080840160005b83811015611f3257815187529582019590820190600101611f16565b509495945050505050565b8c81528b60208201528a604082015289606082015288608082015260007fffffffffffffffff000000000000000000000000000000000000000000000000808a60c01b1660a0840152808960c01b1660a88401528760b0840152808760c01b1660d0840152507fffff0000000000000000000000000000000000000000000000000000000000008560f01b1660d8830152611fe4611fde60da840186611f0a565b84611f0a565b9e9d5050505050505050505050505050565b600083516120088184602088016114d8565b83519083019061201c8183602088016114d8565b01949350505050565b600060ff821660ff84168060ff0382111561204257612042611e7b565b019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000826120885761208861204a565b500490565b60008160001904831182151516156120a7576120a7611e7b565b500290565b600082198211156120bf576120bf611e7b565b500190565b600181815b808511156120ff5781600019048211156120e5576120e5611e7b565b808516156120f257918102915b93841c93908002906120c9565b509250929050565b60008261211657506001610995565b8161212357506000610995565b816001811461213957600281146121435761215f565b6001915050610995565b60ff84111561215457612154611e7b565b50506001821b610995565b5060208310610133831016604e8410600b8410161715612182575081810a610995565b61218c83836120c4565b80600019048211156121a0576121a0611e7b565b029392505050565b60006115618383612107565b6000826121c3576121c361204a565b50069056fea164736f6c634300080f000a",
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

// VerifyZkVmProof is a free data retrieval call binding the contract method 0xe674f0e0.
//
// Solidity: function verifyZkVmProof((bytes,bytes) _zkVmProof, bytes32 _storedSrcOutput, bytes32 _storedDstOutput, bytes32 _storedL1Head) view returns(bytes32 publicInputHash_)
func (_ZKProofVerifier *ZKProofVerifierCaller) VerifyZkVmProof(opts *bind.CallOpts, _zkVmProof TypesZkVmProof, _storedSrcOutput [32]byte, _storedDstOutput [32]byte, _storedL1Head [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ZKProofVerifier.contract.Call(opts, &out, "verifyZkVmProof", _zkVmProof, _storedSrcOutput, _storedDstOutput, _storedL1Head)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VerifyZkVmProof is a free data retrieval call binding the contract method 0xe674f0e0.
//
// Solidity: function verifyZkVmProof((bytes,bytes) _zkVmProof, bytes32 _storedSrcOutput, bytes32 _storedDstOutput, bytes32 _storedL1Head) view returns(bytes32 publicInputHash_)
func (_ZKProofVerifier *ZKProofVerifierSession) VerifyZkVmProof(_zkVmProof TypesZkVmProof, _storedSrcOutput [32]byte, _storedDstOutput [32]byte, _storedL1Head [32]byte) ([32]byte, error) {
	return _ZKProofVerifier.Contract.VerifyZkVmProof(&_ZKProofVerifier.CallOpts, _zkVmProof, _storedSrcOutput, _storedDstOutput, _storedL1Head)
}

// VerifyZkVmProof is a free data retrieval call binding the contract method 0xe674f0e0.
//
// Solidity: function verifyZkVmProof((bytes,bytes) _zkVmProof, bytes32 _storedSrcOutput, bytes32 _storedDstOutput, bytes32 _storedL1Head) view returns(bytes32 publicInputHash_)
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
