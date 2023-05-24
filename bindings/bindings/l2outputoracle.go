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

// L2OutputOracleMetaData contains all meta data concerning the L2OutputOracle contract.
var L2OutputOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_submissionInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2BlockTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startingTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"contractValidatorPool\",\"name\":\"_validatorPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_colosseum\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_finalizationPeriodSeconds\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"outputIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newOutputRoot\",\"type\":\"bytes32\"}],\"name\":\"OutputReplaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2BlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"l1Timestamp\",\"type\":\"uint256\"}],\"name\":\"OutputSubmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"COLOSSEUM\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FINALIZATION_PERIOD_SECONDS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2_BLOCK_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUBMISSION_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_POOL\",\"outputs\":[{\"internalType\":\"contractValidatorPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"}],\"name\":\"computeL2Timestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2OutputIndex\",\"type\":\"uint256\"}],\"name\":\"getL2Output\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"internalType\":\"structTypes.CheckpointOutput\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"}],\"name\":\"getL2OutputAfter\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"internalType\":\"structTypes.CheckpointOutput\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"}],\"name\":\"getL2OutputIndexAfter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_outputIndex\",\"type\":\"uint256\"}],\"name\":\"getSubmitter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startingTimestamp\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestOutputIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextOutputIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2OutputIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_newOutputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_submitter\",\"type\":\"address\"}],\"name\":\"replaceL2Output\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startingBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startingTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_l1BlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_l1BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bondAmount\",\"type\":\"uint256\"}],\"name\":\"submitL2Output\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101806040523480156200001257600080fd5b5060405162001e7238038062001e72833981016040819052620000359162000352565b60006080819052600160a05260c05285620000bd5760405162461bcd60e51b815260206004820152603460248201527f4c324f75747075744f7261636c653a204c3220626c6f636b2074696d65206d7560448201527f73742062652067726561746572207468616e203000000000000000000000000060648201526084015b60405180910390fd5b60008711620001355760405162461bcd60e51b815260206004820152603a60248201527f4c324f75747075744f7261636c653a207375626d697373696f6e20696e74657260448201527f76616c206d7573742062652067726561746572207468616e20300000000000006064820152608401620000b4565b60e08790526101008690526001600160a01b038084166101405282166101205261016081905262000167858562000174565b50505050505050620003c1565b600054610100900460ff1615808015620001955750600054600160ff909116105b80620001c55750620001b2306200032a60201b6200153d1760201c565b158015620001c5575060005460ff166001145b6200022a5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401620000b4565b6000805460ff1916600117905580156200024e576000805461ff0019166101001790555b42821115620002d45760405162461bcd60e51b8152602060048201526044602482018190527f4c324f75747075744f7261636c653a207374617274696e67204c322074696d65908201527f7374616d70206d757374206265206c657373207468616e2063757272656e742060648201526374696d6560e01b608482015260a401620000b4565b60028290556001839055801562000325576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b6001600160a01b03163b151590565b6001600160a01b03811681146200034f57600080fd5b50565b600080600080600080600060e0888a0312156200036e57600080fd5b875196506020880151955060408801519450606088015193506080880151620003978162000339565b60a0890151909350620003aa8162000339565b8092505060c0880151905092959891949750929550565b60805160a05160c05160e05161010051610120516101405161016051611a1a620004586000396000818161047301528181610cde01526113770152600081816103aa015281816107a40152610cb20152600081816102b201526111960152600081816101750152610ed50152600081816101d10152610f2c01526000610561015260006105380152600061050f0152611a1a6000f3fe60806040526004361061015e5760003560e01c80639e45e8f4116100c0578063d1de856c11610074578063e4a3011611610059578063e4a3011614610421578063e664672314610441578063f4daa2911461046157600080fd5b8063d1de856c146103ec578063dcec33481461040c57600080fd5b8063b0ea09a8116100a5578063b0ea09a814610378578063b98debbf14610398578063cf8e5cf0146103cc57600080fd5b80639e45e8f4146102a0578063a25ae557146102f957600080fd5b80636abcf563116101175780637f006420116100fc5780637f0064201461025557806380e95ef014610275578063887862721461028a57600080fd5b80636abcf5631461022a57806370872aa51461023f57600080fd5b8063529933df11610148578063529933df146101bf57806354fd4d50146101f357806369f16eec1461021557600080fd5b80622134cc146101635780634599c788146101aa575b600080fd5b34801561016f57600080fd5b506101977f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020015b60405180910390f35b3480156101b657600080fd5b50610197610495565b3480156101cb57600080fd5b506101977f000000000000000000000000000000000000000000000000000000000000000081565b3480156101ff57600080fd5b50610208610508565b6040516101a191906116c6565b34801561022157600080fd5b506101976105ab565b34801561023657600080fd5b50600354610197565b34801561024b57600080fd5b5061019760015481565b34801561026157600080fd5b50610197610270366004611717565b6105bd565b610288610283366004611730565b6107a2565b005b34801561029657600080fd5b5061019760025481565b3480156102ac57600080fd5b506102d47f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101a1565b34801561030557600080fd5b50610319610314366004611717565b610d93565b60408051825173ffffffffffffffffffffffffffffffffffffffff16815260208084015190820152828201516fffffffffffffffffffffffffffffffff90811692820192909252606092830151909116918101919091526080016101a1565b34801561038457600080fd5b506102d4610393366004611717565b610e50565b3480156103a457600080fd5b506102d47f000000000000000000000000000000000000000000000000000000000000000081565b3480156103d857600080fd5b506103196103e7366004611717565b610e92565b3480156103f857600080fd5b50610197610407366004611717565b610ed1565b34801561041857600080fd5b50610197610f1f565b34801561042d57600080fd5b5061028861043c36600461176b565b610f65565b34801561044d57600080fd5b5061028861045c3660046117b2565b61117e565b34801561046d57600080fd5b506101977f000000000000000000000000000000000000000000000000000000000000000081565b600354600090156104ff57600380546104b09060019061181a565b815481106104c0576104c0611831565b600091825260209091206003909102016002015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16919050565b6001545b905090565b60606105337f0000000000000000000000000000000000000000000000000000000000000000611559565b61055c7f0000000000000000000000000000000000000000000000000000000000000000611559565b6105857f0000000000000000000000000000000000000000000000000000000000000000611559565b60405160200161059793929190611860565b604051602081830303815290604052905090565b6003546000906105039060019061181a565b60006105c7610495565b8211156106675760405162461bcd60e51b815260206004820152604960248201527f4c324f75747075744f7261636c653a2063616e6e6f7420676574206f7574707560448201527f7420666f72206120626c6f636b207468617420686173206e6f74206265656e2060648201527f7375626d69747465640000000000000000000000000000000000000000000000608482015260a4015b60405180910390fd5b6003546107025760405162461bcd60e51b815260206004820152604760248201527f4c324f75747075744f7261636c653a2063616e6e6f7420676574206f7574707560448201527f74206173206e6f206f7574707574732068617665206265656e207375626d697460648201527f7465642079657400000000000000000000000000000000000000000000000000608482015260a40161065e565b6003546000905b8082101561079b576000600261071f83856118d6565b610729919061191d565b9050846003828154811061073f5761073f611831565b600091825260209091206003909102016002015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff1610156107915761078a8160016118d6565b9250610795565b8091505b50610709565b5092915050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633a5490466040518163ffffffff1660e01b8152600401602060405180830381865afa15801561080d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108319190611931565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108f75760405162461bcd60e51b815260206004820152604260248201527f4c324f75747075744f7261636c653a206f6e6c7920746865206e65787420736560448201527f6c65637465642076616c696461746f722063616e207375626d6974206f75747060648201527f7574000000000000000000000000000000000000000000000000000000000000608482015260a40161065e565b6108ff610f1f565b84146109995760405162461bcd60e51b815260206004820152604860248201527f4c324f75747075744f7261636c653a20626c6f636b206e756d626572206d757360448201527f7420626520657175616c20746f206e65787420657870656374656420626c6f6360648201527f6b206e756d626572000000000000000000000000000000000000000000000000608482015260a40161065e565b426109a385610ed1565b10610a165760405162461bcd60e51b815260206004820152603560248201527f4c324f75747075744f7261636c653a2063616e6e6f74207375626d6974204c3260448201527f206f757470757420696e20746865206675747572650000000000000000000000606482015260840161065e565b84610a895760405162461bcd60e51b815260206004820152603c60248201527f4c324f75747075744f7261636c653a204c3220636865636b706f696e74206f7560448201527f747075742063616e6e6f7420626520746865207a65726f206861736800000000606482015260840161065e565b8215610b2b5782824014610b2b5760405162461bcd60e51b815260206004820152604960248201527f4c324f75747075744f7261636c653a20626c6f636b206861736820646f65732060448201527f6e6f74206d61746368207468652068617368206174207468652065787065637460648201527f6564206865696768740000000000000000000000000000000000000000000000608482015260a40161065e565b6000610b3660035490565b60408051608081018252338152602081018981526fffffffffffffffffffffffffffffffff428181168486019081528b831660608601908152600380546001810182556000829052965196027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b8101805473ffffffffffffffffffffffffffffffffffffffff989098167fffffffffffffffffffffffff00000000000000000000000000000000000000009098169790971790965593517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85c86015551925182167001000000000000000000000000000000000292909116919091177fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85d9092019190915590519192508691839189917f457b4388026260019ae0b0b4f16c98235d74fe7359be469bdcba16e6d0d4968991610c939190815260200190565b60405180910390a473ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166396946f758284610d037f0000000000000000000000000000000000000000000000000000000000000000426118d6565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b16815260048101939093526fffffffffffffffffffffffffffffffff9182166024840152166044820152606401600060405180830381600087803b158015610d7357600080fd5b505af1158015610d87573d6000803e3d6000fd5b50505050505050505050565b60408051608081018252600080825260208201819052918101829052606081019190915260038281548110610dca57610dca611831565b6000918252602091829020604080516080810182526003909302909101805473ffffffffffffffffffffffffffffffffffffffff1683526001810154938301939093526002909201546fffffffffffffffffffffffffffffffff808216938301939093527001000000000000000000000000000000009004909116606082015292915050565b600060038281548110610e6557610e65611831565b600091825260209091206003909102015473ffffffffffffffffffffffffffffffffffffffff1692915050565b6040805160808101825260008082526020820181905291810182905260608101919091526003610ec1836105bd565b81548110610dca57610dca611831565b60007f000000000000000000000000000000000000000000000000000000000000000060015483610f02919061181a565b610f0c9190611955565b600254610f1991906118d6565b92915050565b60035460009015610f5d577f0000000000000000000000000000000000000000000000000000000000000000610f53610495565b61050391906118d6565b610503610495565b600054610100900460ff1615808015610f855750600054600160ff909116105b80610f9f5750303b158015610f9f575060005460ff166001145b6110115760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161065e565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561106f57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b4282111561110c5760405162461bcd60e51b8152602060048201526044602482018190527f4c324f75747075744f7261636c653a207374617274696e67204c322074696d65908201527f7374616d70206d757374206265206c657373207468616e2063757272656e742060648201527f74696d6500000000000000000000000000000000000000000000000000000000608482015260a40161065e565b60028290556001839055801561117957600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461124f5760405162461bcd60e51b815260206004820152604160248201527f4c324f75747075744f7261636c653a206f6e6c792074686520636f6c6f73736560448201527f756d20636f6e74726163742063616e207265706c61636520616e206f7574707560648201527f7400000000000000000000000000000000000000000000000000000000000000608482015260a40161065e565b73ffffffffffffffffffffffffffffffffffffffff81166112d85760405162461bcd60e51b815260206004820152603060248201527f4c324f75747075744f7261636c653a207375626d69747465722061646472657360448201527f732063616e6e6f74206265207a65726f00000000000000000000000000000000606482015260840161065e565b60035483106113755760405162461bcd60e51b815260206004820152604660248201527f4c324f75747075744f7261636c653a2063616e6e6f74207265706c616365206160448201527f6e206f757470757420616674657220746865206c6174657374206f757470757460648201527f20696e6465780000000000000000000000000000000000000000000000000000608482015260a40161065e565b7f0000000000000000000000000000000000000000000000000000000000000000600384815481106113a9576113a9611831565b60009182526020909120600260039092020101546113d9906fffffffffffffffffffffffffffffffff164261181a565b106114725760405162461bcd60e51b815260206004820152604860248201527f4c324f75747075744f7261636c653a2063616e6e6f74207265706c616365206160448201527f6e206f757470757420746861742068617320616c7265616479206265656e206660648201527f696e616c697a6564000000000000000000000000000000000000000000000000608482015260a40161065e565b816003848154811061148657611486611831565b90600052602060002090600302016001018190555080600384815481106114af576114af611831565b906000526020600020906003020160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550827fa1b831bb8b6b242db6d0988a6d21f869c610de9f703a5e45e1b7d3dc3137b9068360405161153091815260200190565b60405180910390a2505050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b60608160000361159c57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156115c657806115b081611992565b91506115bf9050600a8361191d565b91506115a0565b60008167ffffffffffffffff8111156115e1576115e16119ca565b6040519080825280601f01601f19166020018201604052801561160b576020820181803683370190505b5090505b841561168e5761162060018361181a565b915061162d600a866119f9565b6116389060306118d6565b60f81b81838151811061164d5761164d611831565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350611687600a8661191d565b945061160f565b949350505050565b60005b838110156116b1578181015183820152602001611699565b838111156116c0576000848401525b50505050565b60208152600082518060208401526116e5816040850160208701611696565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b60006020828403121561172957600080fd5b5035919050565b600080600080600060a0868803121561174857600080fd5b505083359560208501359550604085013594606081013594506080013592509050565b6000806040838503121561177e57600080fd5b50508035926020909101359150565b73ffffffffffffffffffffffffffffffffffffffff811681146117af57600080fd5b50565b6000806000606084860312156117c757600080fd5b833592506020840135915060408401356117e08161178d565b809150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561182c5761182c6117eb565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008451611872818460208901611696565b80830190507f2e0000000000000000000000000000000000000000000000000000000000000080825285516118ae816001850160208a01611696565b600192019182015283516118c9816002840160208801611696565b0160020195945050505050565b600082198211156118e9576118e96117eb565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60008261192c5761192c6118ee565b500490565b60006020828403121561194357600080fd5b815161194e8161178d565b9392505050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561198d5761198d6117eb565b500290565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036119c3576119c36117eb565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082611a0857611a086118ee565b50069056fea164736f6c634300080f000a",
}

// L2OutputOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use L2OutputOracleMetaData.ABI instead.
var L2OutputOracleABI = L2OutputOracleMetaData.ABI

// L2OutputOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2OutputOracleMetaData.Bin instead.
var L2OutputOracleBin = L2OutputOracleMetaData.Bin

// DeployL2OutputOracle deploys a new Ethereum contract, binding an instance of L2OutputOracle to it.
func DeployL2OutputOracle(auth *bind.TransactOpts, backend bind.ContractBackend, _submissionInterval *big.Int, _l2BlockTime *big.Int, _startingBlockNumber *big.Int, _startingTimestamp *big.Int, _validatorPool common.Address, _colosseum common.Address, _finalizationPeriodSeconds *big.Int) (common.Address, *types.Transaction, *L2OutputOracle, error) {
	parsed, err := L2OutputOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2OutputOracleBin), backend, _submissionInterval, _l2BlockTime, _startingBlockNumber, _startingTimestamp, _validatorPool, _colosseum, _finalizationPeriodSeconds)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2OutputOracle{L2OutputOracleCaller: L2OutputOracleCaller{contract: contract}, L2OutputOracleTransactor: L2OutputOracleTransactor{contract: contract}, L2OutputOracleFilterer: L2OutputOracleFilterer{contract: contract}}, nil
}

// L2OutputOracle is an auto generated Go binding around an Ethereum contract.
type L2OutputOracle struct {
	L2OutputOracleCaller     // Read-only binding to the contract
	L2OutputOracleTransactor // Write-only binding to the contract
	L2OutputOracleFilterer   // Log filterer for contract events
}

// L2OutputOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2OutputOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2OutputOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2OutputOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2OutputOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2OutputOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2OutputOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2OutputOracleSession struct {
	Contract     *L2OutputOracle   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2OutputOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2OutputOracleCallerSession struct {
	Contract *L2OutputOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// L2OutputOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2OutputOracleTransactorSession struct {
	Contract     *L2OutputOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// L2OutputOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2OutputOracleRaw struct {
	Contract *L2OutputOracle // Generic contract binding to access the raw methods on
}

// L2OutputOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2OutputOracleCallerRaw struct {
	Contract *L2OutputOracleCaller // Generic read-only contract binding to access the raw methods on
}

// L2OutputOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2OutputOracleTransactorRaw struct {
	Contract *L2OutputOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2OutputOracle creates a new instance of L2OutputOracle, bound to a specific deployed contract.
func NewL2OutputOracle(address common.Address, backend bind.ContractBackend) (*L2OutputOracle, error) {
	contract, err := bindL2OutputOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2OutputOracle{L2OutputOracleCaller: L2OutputOracleCaller{contract: contract}, L2OutputOracleTransactor: L2OutputOracleTransactor{contract: contract}, L2OutputOracleFilterer: L2OutputOracleFilterer{contract: contract}}, nil
}

// NewL2OutputOracleCaller creates a new read-only instance of L2OutputOracle, bound to a specific deployed contract.
func NewL2OutputOracleCaller(address common.Address, caller bind.ContractCaller) (*L2OutputOracleCaller, error) {
	contract, err := bindL2OutputOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2OutputOracleCaller{contract: contract}, nil
}

// NewL2OutputOracleTransactor creates a new write-only instance of L2OutputOracle, bound to a specific deployed contract.
func NewL2OutputOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*L2OutputOracleTransactor, error) {
	contract, err := bindL2OutputOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2OutputOracleTransactor{contract: contract}, nil
}

// NewL2OutputOracleFilterer creates a new log filterer instance of L2OutputOracle, bound to a specific deployed contract.
func NewL2OutputOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*L2OutputOracleFilterer, error) {
	contract, err := bindL2OutputOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2OutputOracleFilterer{contract: contract}, nil
}

// bindL2OutputOracle binds a generic wrapper to an already deployed contract.
func bindL2OutputOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2OutputOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2OutputOracle *L2OutputOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2OutputOracle.Contract.L2OutputOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2OutputOracle *L2OutputOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2OutputOracle.Contract.L2OutputOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2OutputOracle *L2OutputOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2OutputOracle.Contract.L2OutputOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2OutputOracle *L2OutputOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2OutputOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2OutputOracle *L2OutputOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2OutputOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2OutputOracle *L2OutputOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2OutputOracle.Contract.contract.Transact(opts, method, params...)
}

// COLOSSEUM is a free data retrieval call binding the contract method 0x9e45e8f4.
//
// Solidity: function COLOSSEUM() view returns(address)
func (_L2OutputOracle *L2OutputOracleCaller) COLOSSEUM(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2OutputOracle.contract.Call(opts, &out, "COLOSSEUM")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// COLOSSEUM is a free data retrieval call binding the contract method 0x9e45e8f4.
//
// Solidity: function COLOSSEUM() view returns(address)
func (_L2OutputOracle *L2OutputOracleSession) COLOSSEUM() (common.Address, error) {
	return _L2OutputOracle.Contract.COLOSSEUM(&_L2OutputOracle.CallOpts)
}

// COLOSSEUM is a free data retrieval call binding the contract method 0x9e45e8f4.
//
// Solidity: function COLOSSEUM() view returns(address)
func (_L2OutputOracle *L2OutputOracleCallerSession) COLOSSEUM() (common.Address, error) {
	return _L2OutputOracle.Contract.COLOSSEUM(&_L2OutputOracle.CallOpts)
}

// FINALIZATIONPERIODSECONDS is a free data retrieval call binding the contract method 0xf4daa291.
//
// Solidity: function FINALIZATION_PERIOD_SECONDS() view returns(uint256)
func (_L2OutputOracle *L2OutputOracleCaller) FINALIZATIONPERIODSECONDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2OutputOracle.contract.Call(opts, &out, "FINALIZATION_PERIOD_SECONDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FINALIZATIONPERIODSECONDS is a free data retrieval call binding the contract method 0xf4daa291.
//
// Solidity: function FINALIZATION_PERIOD_SECONDS() view returns(uint256)
func (_L2OutputOracle *L2OutputOracleSession) FINALIZATIONPERIODSECONDS() (*big.Int, error) {
	return _L2OutputOracle.Contract.FINALIZATIONPERIODSECONDS(&_L2OutputOracle.CallOpts)
}

// FINALIZATIONPERIODSECONDS is a free data retrieval call binding the contract method 0xf4daa291.
//
// Solidity: function FINALIZATION_PERIOD_SECONDS() view returns(uint256)
func (_L2OutputOracle *L2OutputOracleCallerSession) FINALIZATIONPERIODSECONDS() (*big.Int, error) {
	return _L2OutputOracle.Contract.FINALIZATIONPERIODSECONDS(&_L2OutputOracle.CallOpts)
}

// L2BLOCKTIME is a free data retrieval call binding the contract method 0x002134cc.
//
// Solidity: function L2_BLOCK_TIME() view returns(uint256)
func (_L2OutputOracle *L2OutputOracleCaller) L2BLOCKTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2OutputOracle.contract.Call(opts, &out, "L2_BLOCK_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2BLOCKTIME is a free data retrieval call binding the contract method 0x002134cc.
//
// Solidity: function L2_BLOCK_TIME() view returns(uint256)
func (_L2OutputOracle *L2OutputOracleSession) L2BLOCKTIME() (*big.Int, error) {
	return _L2OutputOracle.Contract.L2BLOCKTIME(&_L2OutputOracle.CallOpts)
}

// L2BLOCKTIME is a free data retrieval call binding the contract method 0x002134cc.
//
// Solidity: function L2_BLOCK_TIME() view returns(uint256)
func (_L2OutputOracle *L2OutputOracleCallerSession) L2BLOCKTIME() (*big.Int, error) {
	return _L2OutputOracle.Contract.L2BLOCKTIME(&_L2OutputOracle.CallOpts)
}

// SUBMISSIONINTERVAL is a free data retrieval call binding the contract method 0x529933df.
//
// Solidity: function SUBMISSION_INTERVAL() view returns(uint256)
func (_L2OutputOracle *L2OutputOracleCaller) SUBMISSIONINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2OutputOracle.contract.Call(opts, &out, "SUBMISSION_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SUBMISSIONINTERVAL is a free data retrieval call binding the contract method 0x529933df.
//
// Solidity: function SUBMISSION_INTERVAL() view returns(uint256)
func (_L2OutputOracle *L2OutputOracleSession) SUBMISSIONINTERVAL() (*big.Int, error) {
	return _L2OutputOracle.Contract.SUBMISSIONINTERVAL(&_L2OutputOracle.CallOpts)
}

// SUBMISSIONINTERVAL is a free data retrieval call binding the contract method 0x529933df.
//
// Solidity: function SUBMISSION_INTERVAL() view returns(uint256)
func (_L2OutputOracle *L2OutputOracleCallerSession) SUBMISSIONINTERVAL() (*big.Int, error) {
	return _L2OutputOracle.Contract.SUBMISSIONINTERVAL(&_L2OutputOracle.CallOpts)
}

// VALIDATORPOOL is a free data retrieval call binding the contract method 0xb98debbf.
//
// Solidity: function VALIDATOR_POOL() view returns(address)
func (_L2OutputOracle *L2OutputOracleCaller) VALIDATORPOOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2OutputOracle.contract.Call(opts, &out, "VALIDATOR_POOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORPOOL is a free data retrieval call binding the contract method 0xb98debbf.
//
// Solidity: function VALIDATOR_POOL() view returns(address)
func (_L2OutputOracle *L2OutputOracleSession) VALIDATORPOOL() (common.Address, error) {
	return _L2OutputOracle.Contract.VALIDATORPOOL(&_L2OutputOracle.CallOpts)
}

// VALIDATORPOOL is a free data retrieval call binding the contract method 0xb98debbf.
//
// Solidity: function VALIDATOR_POOL() view returns(address)
func (_L2OutputOracle *L2OutputOracleCallerSession) VALIDATORPOOL() (common.Address, error) {
	return _L2OutputOracle.Contract.VALIDATORPOOL(&_L2OutputOracle.CallOpts)
}

// ComputeL2Timestamp is a free data retrieval call binding the contract method 0xd1de856c.
//
// Solidity: function computeL2Timestamp(uint256 _l2BlockNumber) view returns(uint256)
func (_L2OutputOracle *L2OutputOracleCaller) ComputeL2Timestamp(opts *bind.CallOpts, _l2BlockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _L2OutputOracle.contract.Call(opts, &out, "computeL2Timestamp", _l2BlockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeL2Timestamp is a free data retrieval call binding the contract method 0xd1de856c.
//
// Solidity: function computeL2Timestamp(uint256 _l2BlockNumber) view returns(uint256)
func (_L2OutputOracle *L2OutputOracleSession) ComputeL2Timestamp(_l2BlockNumber *big.Int) (*big.Int, error) {
	return _L2OutputOracle.Contract.ComputeL2Timestamp(&_L2OutputOracle.CallOpts, _l2BlockNumber)
}

// ComputeL2Timestamp is a free data retrieval call binding the contract method 0xd1de856c.
//
// Solidity: function computeL2Timestamp(uint256 _l2BlockNumber) view returns(uint256)
func (_L2OutputOracle *L2OutputOracleCallerSession) ComputeL2Timestamp(_l2BlockNumber *big.Int) (*big.Int, error) {
	return _L2OutputOracle.Contract.ComputeL2Timestamp(&_L2OutputOracle.CallOpts, _l2BlockNumber)
}

// GetL2Output is a free data retrieval call binding the contract method 0xa25ae557.
//
// Solidity: function getL2Output(uint256 _l2OutputIndex) view returns((address,bytes32,uint128,uint128))
func (_L2OutputOracle *L2OutputOracleCaller) GetL2Output(opts *bind.CallOpts, _l2OutputIndex *big.Int) (TypesCheckpointOutput, error) {
	var out []interface{}
	err := _L2OutputOracle.contract.Call(opts, &out, "getL2Output", _l2OutputIndex)

	if err != nil {
		return *new(TypesCheckpointOutput), err
	}

	out0 := *abi.ConvertType(out[0], new(TypesCheckpointOutput)).(*TypesCheckpointOutput)

	return out0, err

}

// GetL2Output is a free data retrieval call binding the contract method 0xa25ae557.
//
// Solidity: function getL2Output(uint256 _l2OutputIndex) view returns((address,bytes32,uint128,uint128))
func (_L2OutputOracle *L2OutputOracleSession) GetL2Output(_l2OutputIndex *big.Int) (TypesCheckpointOutput, error) {
	return _L2OutputOracle.Contract.GetL2Output(&_L2OutputOracle.CallOpts, _l2OutputIndex)
}

// GetL2Output is a free data retrieval call binding the contract method 0xa25ae557.
//
// Solidity: function getL2Output(uint256 _l2OutputIndex) view returns((address,bytes32,uint128,uint128))
func (_L2OutputOracle *L2OutputOracleCallerSession) GetL2Output(_l2OutputIndex *big.Int) (TypesCheckpointOutput, error) {
	return _L2OutputOracle.Contract.GetL2Output(&_L2OutputOracle.CallOpts, _l2OutputIndex)
}

// GetL2OutputAfter is a free data retrieval call binding the contract method 0xcf8e5cf0.
//
// Solidity: function getL2OutputAfter(uint256 _l2BlockNumber) view returns((address,bytes32,uint128,uint128))
func (_L2OutputOracle *L2OutputOracleCaller) GetL2OutputAfter(opts *bind.CallOpts, _l2BlockNumber *big.Int) (TypesCheckpointOutput, error) {
	var out []interface{}
	err := _L2OutputOracle.contract.Call(opts, &out, "getL2OutputAfter", _l2BlockNumber)

	if err != nil {
		return *new(TypesCheckpointOutput), err
	}

	out0 := *abi.ConvertType(out[0], new(TypesCheckpointOutput)).(*TypesCheckpointOutput)

	return out0, err

}

// GetL2OutputAfter is a free data retrieval call binding the contract method 0xcf8e5cf0.
//
// Solidity: function getL2OutputAfter(uint256 _l2BlockNumber) view returns((address,bytes32,uint128,uint128))
func (_L2OutputOracle *L2OutputOracleSession) GetL2OutputAfter(_l2BlockNumber *big.Int) (TypesCheckpointOutput, error) {
	return _L2OutputOracle.Contract.GetL2OutputAfter(&_L2OutputOracle.CallOpts, _l2BlockNumber)
}

// GetL2OutputAfter is a free data retrieval call binding the contract method 0xcf8e5cf0.
//
// Solidity: function getL2OutputAfter(uint256 _l2BlockNumber) view returns((address,bytes32,uint128,uint128))
func (_L2OutputOracle *L2OutputOracleCallerSession) GetL2OutputAfter(_l2BlockNumber *big.Int) (TypesCheckpointOutput, error) {
	return _L2OutputOracle.Contract.GetL2OutputAfter(&_L2OutputOracle.CallOpts, _l2BlockNumber)
}

// GetL2OutputIndexAfter is a free data retrieval call binding the contract method 0x7f006420.
//
// Solidity: function getL2OutputIndexAfter(uint256 _l2BlockNumber) view returns(uint256)
func (_L2OutputOracle *L2OutputOracleCaller) GetL2OutputIndexAfter(opts *bind.CallOpts, _l2BlockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _L2OutputOracle.contract.Call(opts, &out, "getL2OutputIndexAfter", _l2BlockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetL2OutputIndexAfter is a free data retrieval call binding the contract method 0x7f006420.
//
// Solidity: function getL2OutputIndexAfter(uint256 _l2BlockNumber) view returns(uint256)
func (_L2OutputOracle *L2OutputOracleSession) GetL2OutputIndexAfter(_l2BlockNumber *big.Int) (*big.Int, error) {
	return _L2OutputOracle.Contract.GetL2OutputIndexAfter(&_L2OutputOracle.CallOpts, _l2BlockNumber)
}

// GetL2OutputIndexAfter is a free data retrieval call binding the contract method 0x7f006420.
//
// Solidity: function getL2OutputIndexAfter(uint256 _l2BlockNumber) view returns(uint256)
func (_L2OutputOracle *L2OutputOracleCallerSession) GetL2OutputIndexAfter(_l2BlockNumber *big.Int) (*big.Int, error) {
	return _L2OutputOracle.Contract.GetL2OutputIndexAfter(&_L2OutputOracle.CallOpts, _l2BlockNumber)
}

// GetSubmitter is a free data retrieval call binding the contract method 0xb0ea09a8.
//
// Solidity: function getSubmitter(uint256 _outputIndex) view returns(address)
func (_L2OutputOracle *L2OutputOracleCaller) GetSubmitter(opts *bind.CallOpts, _outputIndex *big.Int) (common.Address, error) {
	var out []interface{}
	err := _L2OutputOracle.contract.Call(opts, &out, "getSubmitter", _outputIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSubmitter is a free data retrieval call binding the contract method 0xb0ea09a8.
//
// Solidity: function getSubmitter(uint256 _outputIndex) view returns(address)
func (_L2OutputOracle *L2OutputOracleSession) GetSubmitter(_outputIndex *big.Int) (common.Address, error) {
	return _L2OutputOracle.Contract.GetSubmitter(&_L2OutputOracle.CallOpts, _outputIndex)
}

// GetSubmitter is a free data retrieval call binding the contract method 0xb0ea09a8.
//
// Solidity: function getSubmitter(uint256 _outputIndex) view returns(address)
func (_L2OutputOracle *L2OutputOracleCallerSession) GetSubmitter(_outputIndex *big.Int) (common.Address, error) {
	return _L2OutputOracle.Contract.GetSubmitter(&_L2OutputOracle.CallOpts, _outputIndex)
}

// LatestBlockNumber is a free data retrieval call binding the contract method 0x4599c788.
//
// Solidity: function latestBlockNumber() view returns(uint256)
func (_L2OutputOracle *L2OutputOracleCaller) LatestBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2OutputOracle.contract.Call(opts, &out, "latestBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestBlockNumber is a free data retrieval call binding the contract method 0x4599c788.
//
// Solidity: function latestBlockNumber() view returns(uint256)
func (_L2OutputOracle *L2OutputOracleSession) LatestBlockNumber() (*big.Int, error) {
	return _L2OutputOracle.Contract.LatestBlockNumber(&_L2OutputOracle.CallOpts)
}

// LatestBlockNumber is a free data retrieval call binding the contract method 0x4599c788.
//
// Solidity: function latestBlockNumber() view returns(uint256)
func (_L2OutputOracle *L2OutputOracleCallerSession) LatestBlockNumber() (*big.Int, error) {
	return _L2OutputOracle.Contract.LatestBlockNumber(&_L2OutputOracle.CallOpts)
}

// LatestOutputIndex is a free data retrieval call binding the contract method 0x69f16eec.
//
// Solidity: function latestOutputIndex() view returns(uint256)
func (_L2OutputOracle *L2OutputOracleCaller) LatestOutputIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2OutputOracle.contract.Call(opts, &out, "latestOutputIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestOutputIndex is a free data retrieval call binding the contract method 0x69f16eec.
//
// Solidity: function latestOutputIndex() view returns(uint256)
func (_L2OutputOracle *L2OutputOracleSession) LatestOutputIndex() (*big.Int, error) {
	return _L2OutputOracle.Contract.LatestOutputIndex(&_L2OutputOracle.CallOpts)
}

// LatestOutputIndex is a free data retrieval call binding the contract method 0x69f16eec.
//
// Solidity: function latestOutputIndex() view returns(uint256)
func (_L2OutputOracle *L2OutputOracleCallerSession) LatestOutputIndex() (*big.Int, error) {
	return _L2OutputOracle.Contract.LatestOutputIndex(&_L2OutputOracle.CallOpts)
}

// NextBlockNumber is a free data retrieval call binding the contract method 0xdcec3348.
//
// Solidity: function nextBlockNumber() view returns(uint256)
func (_L2OutputOracle *L2OutputOracleCaller) NextBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2OutputOracle.contract.Call(opts, &out, "nextBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextBlockNumber is a free data retrieval call binding the contract method 0xdcec3348.
//
// Solidity: function nextBlockNumber() view returns(uint256)
func (_L2OutputOracle *L2OutputOracleSession) NextBlockNumber() (*big.Int, error) {
	return _L2OutputOracle.Contract.NextBlockNumber(&_L2OutputOracle.CallOpts)
}

// NextBlockNumber is a free data retrieval call binding the contract method 0xdcec3348.
//
// Solidity: function nextBlockNumber() view returns(uint256)
func (_L2OutputOracle *L2OutputOracleCallerSession) NextBlockNumber() (*big.Int, error) {
	return _L2OutputOracle.Contract.NextBlockNumber(&_L2OutputOracle.CallOpts)
}

// NextOutputIndex is a free data retrieval call binding the contract method 0x6abcf563.
//
// Solidity: function nextOutputIndex() view returns(uint256)
func (_L2OutputOracle *L2OutputOracleCaller) NextOutputIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2OutputOracle.contract.Call(opts, &out, "nextOutputIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextOutputIndex is a free data retrieval call binding the contract method 0x6abcf563.
//
// Solidity: function nextOutputIndex() view returns(uint256)
func (_L2OutputOracle *L2OutputOracleSession) NextOutputIndex() (*big.Int, error) {
	return _L2OutputOracle.Contract.NextOutputIndex(&_L2OutputOracle.CallOpts)
}

// NextOutputIndex is a free data retrieval call binding the contract method 0x6abcf563.
//
// Solidity: function nextOutputIndex() view returns(uint256)
func (_L2OutputOracle *L2OutputOracleCallerSession) NextOutputIndex() (*big.Int, error) {
	return _L2OutputOracle.Contract.NextOutputIndex(&_L2OutputOracle.CallOpts)
}

// StartingBlockNumber is a free data retrieval call binding the contract method 0x70872aa5.
//
// Solidity: function startingBlockNumber() view returns(uint256)
func (_L2OutputOracle *L2OutputOracleCaller) StartingBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2OutputOracle.contract.Call(opts, &out, "startingBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartingBlockNumber is a free data retrieval call binding the contract method 0x70872aa5.
//
// Solidity: function startingBlockNumber() view returns(uint256)
func (_L2OutputOracle *L2OutputOracleSession) StartingBlockNumber() (*big.Int, error) {
	return _L2OutputOracle.Contract.StartingBlockNumber(&_L2OutputOracle.CallOpts)
}

// StartingBlockNumber is a free data retrieval call binding the contract method 0x70872aa5.
//
// Solidity: function startingBlockNumber() view returns(uint256)
func (_L2OutputOracle *L2OutputOracleCallerSession) StartingBlockNumber() (*big.Int, error) {
	return _L2OutputOracle.Contract.StartingBlockNumber(&_L2OutputOracle.CallOpts)
}

// StartingTimestamp is a free data retrieval call binding the contract method 0x88786272.
//
// Solidity: function startingTimestamp() view returns(uint256)
func (_L2OutputOracle *L2OutputOracleCaller) StartingTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2OutputOracle.contract.Call(opts, &out, "startingTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartingTimestamp is a free data retrieval call binding the contract method 0x88786272.
//
// Solidity: function startingTimestamp() view returns(uint256)
func (_L2OutputOracle *L2OutputOracleSession) StartingTimestamp() (*big.Int, error) {
	return _L2OutputOracle.Contract.StartingTimestamp(&_L2OutputOracle.CallOpts)
}

// StartingTimestamp is a free data retrieval call binding the contract method 0x88786272.
//
// Solidity: function startingTimestamp() view returns(uint256)
func (_L2OutputOracle *L2OutputOracleCallerSession) StartingTimestamp() (*big.Int, error) {
	return _L2OutputOracle.Contract.StartingTimestamp(&_L2OutputOracle.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2OutputOracle *L2OutputOracleCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L2OutputOracle.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2OutputOracle *L2OutputOracleSession) Version() (string, error) {
	return _L2OutputOracle.Contract.Version(&_L2OutputOracle.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2OutputOracle *L2OutputOracleCallerSession) Version() (string, error) {
	return _L2OutputOracle.Contract.Version(&_L2OutputOracle.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xe4a30116.
//
// Solidity: function initialize(uint256 _startingBlockNumber, uint256 _startingTimestamp) returns()
func (_L2OutputOracle *L2OutputOracleTransactor) Initialize(opts *bind.TransactOpts, _startingBlockNumber *big.Int, _startingTimestamp *big.Int) (*types.Transaction, error) {
	return _L2OutputOracle.contract.Transact(opts, "initialize", _startingBlockNumber, _startingTimestamp)
}

// Initialize is a paid mutator transaction binding the contract method 0xe4a30116.
//
// Solidity: function initialize(uint256 _startingBlockNumber, uint256 _startingTimestamp) returns()
func (_L2OutputOracle *L2OutputOracleSession) Initialize(_startingBlockNumber *big.Int, _startingTimestamp *big.Int) (*types.Transaction, error) {
	return _L2OutputOracle.Contract.Initialize(&_L2OutputOracle.TransactOpts, _startingBlockNumber, _startingTimestamp)
}

// Initialize is a paid mutator transaction binding the contract method 0xe4a30116.
//
// Solidity: function initialize(uint256 _startingBlockNumber, uint256 _startingTimestamp) returns()
func (_L2OutputOracle *L2OutputOracleTransactorSession) Initialize(_startingBlockNumber *big.Int, _startingTimestamp *big.Int) (*types.Transaction, error) {
	return _L2OutputOracle.Contract.Initialize(&_L2OutputOracle.TransactOpts, _startingBlockNumber, _startingTimestamp)
}

// ReplaceL2Output is a paid mutator transaction binding the contract method 0xe6646723.
//
// Solidity: function replaceL2Output(uint256 _l2OutputIndex, bytes32 _newOutputRoot, address _submitter) returns()
func (_L2OutputOracle *L2OutputOracleTransactor) ReplaceL2Output(opts *bind.TransactOpts, _l2OutputIndex *big.Int, _newOutputRoot [32]byte, _submitter common.Address) (*types.Transaction, error) {
	return _L2OutputOracle.contract.Transact(opts, "replaceL2Output", _l2OutputIndex, _newOutputRoot, _submitter)
}

// ReplaceL2Output is a paid mutator transaction binding the contract method 0xe6646723.
//
// Solidity: function replaceL2Output(uint256 _l2OutputIndex, bytes32 _newOutputRoot, address _submitter) returns()
func (_L2OutputOracle *L2OutputOracleSession) ReplaceL2Output(_l2OutputIndex *big.Int, _newOutputRoot [32]byte, _submitter common.Address) (*types.Transaction, error) {
	return _L2OutputOracle.Contract.ReplaceL2Output(&_L2OutputOracle.TransactOpts, _l2OutputIndex, _newOutputRoot, _submitter)
}

// ReplaceL2Output is a paid mutator transaction binding the contract method 0xe6646723.
//
// Solidity: function replaceL2Output(uint256 _l2OutputIndex, bytes32 _newOutputRoot, address _submitter) returns()
func (_L2OutputOracle *L2OutputOracleTransactorSession) ReplaceL2Output(_l2OutputIndex *big.Int, _newOutputRoot [32]byte, _submitter common.Address) (*types.Transaction, error) {
	return _L2OutputOracle.Contract.ReplaceL2Output(&_L2OutputOracle.TransactOpts, _l2OutputIndex, _newOutputRoot, _submitter)
}

// SubmitL2Output is a paid mutator transaction binding the contract method 0x80e95ef0.
//
// Solidity: function submitL2Output(bytes32 _outputRoot, uint256 _l2BlockNumber, bytes32 _l1BlockHash, uint256 _l1BlockNumber, uint256 _bondAmount) payable returns()
func (_L2OutputOracle *L2OutputOracleTransactor) SubmitL2Output(opts *bind.TransactOpts, _outputRoot [32]byte, _l2BlockNumber *big.Int, _l1BlockHash [32]byte, _l1BlockNumber *big.Int, _bondAmount *big.Int) (*types.Transaction, error) {
	return _L2OutputOracle.contract.Transact(opts, "submitL2Output", _outputRoot, _l2BlockNumber, _l1BlockHash, _l1BlockNumber, _bondAmount)
}

// SubmitL2Output is a paid mutator transaction binding the contract method 0x80e95ef0.
//
// Solidity: function submitL2Output(bytes32 _outputRoot, uint256 _l2BlockNumber, bytes32 _l1BlockHash, uint256 _l1BlockNumber, uint256 _bondAmount) payable returns()
func (_L2OutputOracle *L2OutputOracleSession) SubmitL2Output(_outputRoot [32]byte, _l2BlockNumber *big.Int, _l1BlockHash [32]byte, _l1BlockNumber *big.Int, _bondAmount *big.Int) (*types.Transaction, error) {
	return _L2OutputOracle.Contract.SubmitL2Output(&_L2OutputOracle.TransactOpts, _outputRoot, _l2BlockNumber, _l1BlockHash, _l1BlockNumber, _bondAmount)
}

// SubmitL2Output is a paid mutator transaction binding the contract method 0x80e95ef0.
//
// Solidity: function submitL2Output(bytes32 _outputRoot, uint256 _l2BlockNumber, bytes32 _l1BlockHash, uint256 _l1BlockNumber, uint256 _bondAmount) payable returns()
func (_L2OutputOracle *L2OutputOracleTransactorSession) SubmitL2Output(_outputRoot [32]byte, _l2BlockNumber *big.Int, _l1BlockHash [32]byte, _l1BlockNumber *big.Int, _bondAmount *big.Int) (*types.Transaction, error) {
	return _L2OutputOracle.Contract.SubmitL2Output(&_L2OutputOracle.TransactOpts, _outputRoot, _l2BlockNumber, _l1BlockHash, _l1BlockNumber, _bondAmount)
}

// L2OutputOracleInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2OutputOracle contract.
type L2OutputOracleInitializedIterator struct {
	Event *L2OutputOracleInitialized // Event containing the contract specifics and raw log

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
func (it *L2OutputOracleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2OutputOracleInitialized)
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
		it.Event = new(L2OutputOracleInitialized)
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
func (it *L2OutputOracleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2OutputOracleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2OutputOracleInitialized represents a Initialized event raised by the L2OutputOracle contract.
type L2OutputOracleInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2OutputOracle *L2OutputOracleFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2OutputOracleInitializedIterator, error) {

	logs, sub, err := _L2OutputOracle.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2OutputOracleInitializedIterator{contract: _L2OutputOracle.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2OutputOracle *L2OutputOracleFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2OutputOracleInitialized) (event.Subscription, error) {

	logs, sub, err := _L2OutputOracle.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2OutputOracleInitialized)
				if err := _L2OutputOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2OutputOracle *L2OutputOracleFilterer) ParseInitialized(log types.Log) (*L2OutputOracleInitialized, error) {
	event := new(L2OutputOracleInitialized)
	if err := _L2OutputOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2OutputOracleOutputReplacedIterator is returned from FilterOutputReplaced and is used to iterate over the raw logs and unpacked data for OutputReplaced events raised by the L2OutputOracle contract.
type L2OutputOracleOutputReplacedIterator struct {
	Event *L2OutputOracleOutputReplaced // Event containing the contract specifics and raw log

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
func (it *L2OutputOracleOutputReplacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2OutputOracleOutputReplaced)
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
		it.Event = new(L2OutputOracleOutputReplaced)
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
func (it *L2OutputOracleOutputReplacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2OutputOracleOutputReplacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2OutputOracleOutputReplaced represents a OutputReplaced event raised by the L2OutputOracle contract.
type L2OutputOracleOutputReplaced struct {
	OutputIndex   *big.Int
	NewOutputRoot [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOutputReplaced is a free log retrieval operation binding the contract event 0xa1b831bb8b6b242db6d0988a6d21f869c610de9f703a5e45e1b7d3dc3137b906.
//
// Solidity: event OutputReplaced(uint256 indexed outputIndex, bytes32 newOutputRoot)
func (_L2OutputOracle *L2OutputOracleFilterer) FilterOutputReplaced(opts *bind.FilterOpts, outputIndex []*big.Int) (*L2OutputOracleOutputReplacedIterator, error) {

	var outputIndexRule []interface{}
	for _, outputIndexItem := range outputIndex {
		outputIndexRule = append(outputIndexRule, outputIndexItem)
	}

	logs, sub, err := _L2OutputOracle.contract.FilterLogs(opts, "OutputReplaced", outputIndexRule)
	if err != nil {
		return nil, err
	}
	return &L2OutputOracleOutputReplacedIterator{contract: _L2OutputOracle.contract, event: "OutputReplaced", logs: logs, sub: sub}, nil
}

// WatchOutputReplaced is a free log subscription operation binding the contract event 0xa1b831bb8b6b242db6d0988a6d21f869c610de9f703a5e45e1b7d3dc3137b906.
//
// Solidity: event OutputReplaced(uint256 indexed outputIndex, bytes32 newOutputRoot)
func (_L2OutputOracle *L2OutputOracleFilterer) WatchOutputReplaced(opts *bind.WatchOpts, sink chan<- *L2OutputOracleOutputReplaced, outputIndex []*big.Int) (event.Subscription, error) {

	var outputIndexRule []interface{}
	for _, outputIndexItem := range outputIndex {
		outputIndexRule = append(outputIndexRule, outputIndexItem)
	}

	logs, sub, err := _L2OutputOracle.contract.WatchLogs(opts, "OutputReplaced", outputIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2OutputOracleOutputReplaced)
				if err := _L2OutputOracle.contract.UnpackLog(event, "OutputReplaced", log); err != nil {
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

// ParseOutputReplaced is a log parse operation binding the contract event 0xa1b831bb8b6b242db6d0988a6d21f869c610de9f703a5e45e1b7d3dc3137b906.
//
// Solidity: event OutputReplaced(uint256 indexed outputIndex, bytes32 newOutputRoot)
func (_L2OutputOracle *L2OutputOracleFilterer) ParseOutputReplaced(log types.Log) (*L2OutputOracleOutputReplaced, error) {
	event := new(L2OutputOracleOutputReplaced)
	if err := _L2OutputOracle.contract.UnpackLog(event, "OutputReplaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2OutputOracleOutputSubmittedIterator is returned from FilterOutputSubmitted and is used to iterate over the raw logs and unpacked data for OutputSubmitted events raised by the L2OutputOracle contract.
type L2OutputOracleOutputSubmittedIterator struct {
	Event *L2OutputOracleOutputSubmitted // Event containing the contract specifics and raw log

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
func (it *L2OutputOracleOutputSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2OutputOracleOutputSubmitted)
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
		it.Event = new(L2OutputOracleOutputSubmitted)
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
func (it *L2OutputOracleOutputSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2OutputOracleOutputSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2OutputOracleOutputSubmitted represents a OutputSubmitted event raised by the L2OutputOracle contract.
type L2OutputOracleOutputSubmitted struct {
	OutputRoot    [32]byte
	L2OutputIndex *big.Int
	L2BlockNumber *big.Int
	L1Timestamp   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOutputSubmitted is a free log retrieval operation binding the contract event 0x457b4388026260019ae0b0b4f16c98235d74fe7359be469bdcba16e6d0d49689.
//
// Solidity: event OutputSubmitted(bytes32 indexed outputRoot, uint256 indexed l2OutputIndex, uint256 indexed l2BlockNumber, uint256 l1Timestamp)
func (_L2OutputOracle *L2OutputOracleFilterer) FilterOutputSubmitted(opts *bind.FilterOpts, outputRoot [][32]byte, l2OutputIndex []*big.Int, l2BlockNumber []*big.Int) (*L2OutputOracleOutputSubmittedIterator, error) {

	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}
	var l2OutputIndexRule []interface{}
	for _, l2OutputIndexItem := range l2OutputIndex {
		l2OutputIndexRule = append(l2OutputIndexRule, l2OutputIndexItem)
	}
	var l2BlockNumberRule []interface{}
	for _, l2BlockNumberItem := range l2BlockNumber {
		l2BlockNumberRule = append(l2BlockNumberRule, l2BlockNumberItem)
	}

	logs, sub, err := _L2OutputOracle.contract.FilterLogs(opts, "OutputSubmitted", outputRootRule, l2OutputIndexRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return &L2OutputOracleOutputSubmittedIterator{contract: _L2OutputOracle.contract, event: "OutputSubmitted", logs: logs, sub: sub}, nil
}

// WatchOutputSubmitted is a free log subscription operation binding the contract event 0x457b4388026260019ae0b0b4f16c98235d74fe7359be469bdcba16e6d0d49689.
//
// Solidity: event OutputSubmitted(bytes32 indexed outputRoot, uint256 indexed l2OutputIndex, uint256 indexed l2BlockNumber, uint256 l1Timestamp)
func (_L2OutputOracle *L2OutputOracleFilterer) WatchOutputSubmitted(opts *bind.WatchOpts, sink chan<- *L2OutputOracleOutputSubmitted, outputRoot [][32]byte, l2OutputIndex []*big.Int, l2BlockNumber []*big.Int) (event.Subscription, error) {

	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}
	var l2OutputIndexRule []interface{}
	for _, l2OutputIndexItem := range l2OutputIndex {
		l2OutputIndexRule = append(l2OutputIndexRule, l2OutputIndexItem)
	}
	var l2BlockNumberRule []interface{}
	for _, l2BlockNumberItem := range l2BlockNumber {
		l2BlockNumberRule = append(l2BlockNumberRule, l2BlockNumberItem)
	}

	logs, sub, err := _L2OutputOracle.contract.WatchLogs(opts, "OutputSubmitted", outputRootRule, l2OutputIndexRule, l2BlockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2OutputOracleOutputSubmitted)
				if err := _L2OutputOracle.contract.UnpackLog(event, "OutputSubmitted", log); err != nil {
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

// ParseOutputSubmitted is a log parse operation binding the contract event 0x457b4388026260019ae0b0b4f16c98235d74fe7359be469bdcba16e6d0d49689.
//
// Solidity: event OutputSubmitted(bytes32 indexed outputRoot, uint256 indexed l2OutputIndex, uint256 indexed l2BlockNumber, uint256 l1Timestamp)
func (_L2OutputOracle *L2OutputOracleFilterer) ParseOutputSubmitted(log types.Log) (*L2OutputOracleOutputSubmitted, error) {
	event := new(L2OutputOracleOutputSubmitted)
	if err := _L2OutputOracle.contract.UnpackLog(event, "OutputSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
