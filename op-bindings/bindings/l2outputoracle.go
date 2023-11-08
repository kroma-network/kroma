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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractValidatorPool\",\"name\":\"_validatorPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_colosseum\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_submissionInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2BlockTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startingTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_finalizationPeriodSeconds\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"outputIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newOutputRoot\",\"type\":\"bytes32\"}],\"name\":\"OutputReplaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2OutputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2BlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"l1Timestamp\",\"type\":\"uint256\"}],\"name\":\"OutputSubmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"COLOSSEUM\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FINALIZATION_PERIOD_SECONDS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2_BLOCK_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUBMISSION_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_POOL\",\"outputs\":[{\"internalType\":\"contractValidatorPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"}],\"name\":\"computeL2Timestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_outputIndex\",\"type\":\"uint256\"}],\"name\":\"finalizedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2OutputIndex\",\"type\":\"uint256\"}],\"name\":\"getL2Output\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"internalType\":\"structTypes.CheckpointOutput\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"}],\"name\":\"getL2OutputAfter\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"l2BlockNumber\",\"type\":\"uint128\"}],\"internalType\":\"structTypes.CheckpointOutput\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"}],\"name\":\"getL2OutputIndexAfter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_outputIndex\",\"type\":\"uint256\"}],\"name\":\"getSubmitter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startingTimestamp\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_outputIndex\",\"type\":\"uint256\"}],\"name\":\"isFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestOutputIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextOutputIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2OutputIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_newOutputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_submitter\",\"type\":\"address\"}],\"name\":\"replaceL2Output\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startingBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startingTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_l1BlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_l1BlockNumber\",\"type\":\"uint256\"}],\"name\":\"submitL2Output\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101806040523480156200001257600080fd5b5060405162001f8a38038062001f8a833981016040819052620000359162000352565b6001608052600060a081905260c05283620000bd5760405162461bcd60e51b815260206004820152603460248201527f4c324f75747075744f7261636c653a204c3220626c6f636b2074696d65206d7560448201527f73742062652067726561746572207468616e203000000000000000000000000060648201526084015b60405180910390fd5b60008511620001355760405162461bcd60e51b815260206004820152603a60248201527f4c324f75747075744f7261636c653a207375626d697373696f6e20696e74657260448201527f76616c206d7573742062652067726561746572207468616e20300000000000006064820152608401620000b4565b6001600160a01b0380881660e05286166101005261012085905261014084905261016081905262000167838362000174565b50505050505050620003be565b600054610100900460ff1615808015620001955750600054600160ff909116105b80620001c55750620001b2306200032a60201b620016441760201c565b158015620001c5575060005460ff166001145b6200022a5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401620000b4565b6000805460ff1916600117905580156200024e576000805461ff0019166101001790555b42821115620002d45760405162461bcd60e51b8152602060048201526044602482018190527f4c324f75747075744f7261636c653a207374617274696e67204c322074696d65908201527f7374616d70206d757374206265206c657373207468616e2063757272656e742060648201526374696d6560e01b608482015260a401620000b4565b60028290556001839055801562000325576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b6001600160a01b03163b151590565b6001600160a01b03811681146200034f57600080fd5b50565b600080600080600080600060e0888a0312156200036e57600080fd5b87516200037b8162000339565b60208901519097506200038e8162000339565b604089015160608a015160808b015160a08c015160c0909c01519a9d939c50919a90999198509650945092505050565b60805160a05160c05160e05161010051610120516101405161016051611b2762000463600039600081816104d90152818161050001528181610bd901528181610f3901526114ef01526000818161018b015261102601526000818161021701526110770152600081816102f801526112e1015260008181610410015281816106830152610bae015260006106350152600061060c015260006105e30152611b276000f3fe6080604052600436106101745760003560e01c80639e45e8f4116100cb578063cf8e5cf01161007f578063e4a3011611610059578063e4a3011614610487578063e6646723146104a7578063f4daa291146104c757600080fd5b8063cf8e5cf014610432578063d1de856c14610452578063dcec33481461047257600080fd5b8063a48ea6de116100b0578063a48ea6de146103be578063b0ea09a8146103de578063b98debbf146103fe57600080fd5b80639e45e8f4146102e6578063a25ae5571461033f57600080fd5b80635a045f781161012d57806370872aa51161010757806370872aa51461029a5780637f006420146102b057806388786272146102d057600080fd5b80635a045f781461025b57806369f16eec146102705780636abcf5631461028557600080fd5b80634599c7881161015e5780634599c788146101f0578063529933df1461020557806354fd4d501461023957600080fd5b80622134cc1461017957806333727c4d146101c0575b600080fd5b34801561018557600080fd5b506101ad7f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020015b60405180910390f35b3480156101cc57600080fd5b506101e06101db366004611800565b6104fb565b60405190151581526020016101b7565b3480156101fc57600080fd5b506101ad610569565b34801561021157600080fd5b506101ad7f000000000000000000000000000000000000000000000000000000000000000081565b34801561024557600080fd5b5061024e6105dc565b6040516101b79190611849565b61026e61026936600461189a565b61067f565b005b34801561027c57600080fd5b506101ad610c86565b34801561029157600080fd5b506003546101ad565b3480156102a657600080fd5b506101ad60015481565b3480156102bc57600080fd5b506101ad6102cb366004611800565b610c98565b3480156102dc57600080fd5b506101ad60025481565b3480156102f257600080fd5b5061031a7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101b7565b34801561034b57600080fd5b5061035f61035a366004611800565b610e78565b60408051825173ffffffffffffffffffffffffffffffffffffffff16815260208084015190820152828201516fffffffffffffffffffffffffffffffff90811692820192909252606092830151909116918101919091526080016101b7565b3480156103ca57600080fd5b506101ad6103d9366004611800565b610f35565b3480156103ea57600080fd5b5061031a6103f9366004611800565b610fa1565b34801561040a57600080fd5b5061031a7f000000000000000000000000000000000000000000000000000000000000000081565b34801561043e57600080fd5b5061035f61044d366004611800565b610fe3565b34801561045e57600080fd5b506101ad61046d366004611800565b611022565b34801561047e57600080fd5b506101ad61106a565b34801561049357600080fd5b5061026e6104a23660046118cc565b6110b0565b3480156104b357600080fd5b5061026e6104c2366004611913565b6112c9565b3480156104d357600080fd5b506101ad7f000000000000000000000000000000000000000000000000000000000000000081565b6000427f0000000000000000000000000000000000000000000000000000000000000000600384815481106105325761053261194c565b600091825260209091206002600390920201015461056291906fffffffffffffffffffffffffffffffff166119aa565b1092915050565b600354600090156105d35760038054610584906001906119c2565b815481106105945761059461194c565b600091825260209091206003909102016002015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16919050565b6001545b905090565b60606106077f0000000000000000000000000000000000000000000000000000000000000000611660565b6106307f0000000000000000000000000000000000000000000000000000000000000000611660565b6106597f0000000000000000000000000000000000000000000000000000000000000000611660565b60405160200161066b939291906119d9565b604051602081830303815290604052905090565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633a5490466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106ec573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107109190611a4f565b905073ffffffffffffffffffffffffffffffffffffffff808216146107e5573373ffffffffffffffffffffffffffffffffffffffff8216146107e55760405162461bcd60e51b815260206004820152604260248201527f4c324f75747075744f7261636c653a206f6e6c7920746865206e65787420736560448201527f6c65637465642076616c696461746f722063616e207375626d6974206f75747060648201527f7574000000000000000000000000000000000000000000000000000000000000608482015260a4015b60405180910390fd5b6107ed61106a565b84146108875760405162461bcd60e51b815260206004820152604860248201527f4c324f75747075744f7261636c653a20626c6f636b206e756d626572206d757360448201527f7420626520657175616c20746f206e65787420657870656374656420626c6f6360648201527f6b206e756d626572000000000000000000000000000000000000000000000000608482015260a4016107dc565b4261089185611022565b106109045760405162461bcd60e51b815260206004820152603560248201527f4c324f75747075744f7261636c653a2063616e6e6f74207375626d6974204c3260448201527f206f757470757420696e2074686520667574757265000000000000000000000060648201526084016107dc565b846109775760405162461bcd60e51b815260206004820152603c60248201527f4c324f75747075744f7261636c653a204c3220636865636b706f696e74206f7560448201527f747075742063616e6e6f7420626520746865207a65726f20686173680000000060648201526084016107dc565b82158015906109865750814015155b15610a275782824014610a275760405162461bcd60e51b815260206004820152604960248201527f4c324f75747075744f7261636c653a20626c6f636b206861736820646f65732060448201527f6e6f74206d61746368207468652068617368206174207468652065787065637460648201527f6564206865696768740000000000000000000000000000000000000000000000608482015260a4016107dc565b6000610a3260035490565b60408051608081018252338152602081018981526fffffffffffffffffffffffffffffffff428181168486019081528b831660608601908152600380546001810182556000829052965196027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b8101805473ffffffffffffffffffffffffffffffffffffffff989098167fffffffffffffffffffffffff00000000000000000000000000000000000000009098169790971790965593517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85c86015551925182167001000000000000000000000000000000000292909116919091177fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85d9092019190915590519192508691839189917f457b4388026260019ae0b0b4f16c98235d74fe7359be469bdcba16e6d0d4968991610b8f9190815260200190565b60405180910390a473ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001663d38dc7ee82610bfe7f0000000000000000000000000000000000000000000000000000000000000000426119aa565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815260048101929092526fffffffffffffffffffffffffffffffff166024820152604401600060405180830381600087803b158015610c6657600080fd5b505af1158015610c7a573d6000803e3d6000fd5b50505050505050505050565b6003546000906105d7906001906119c2565b6000610ca2610569565b821115610d3d5760405162461bcd60e51b815260206004820152604960248201527f4c324f75747075744f7261636c653a2063616e6e6f7420676574206f7574707560448201527f7420666f72206120626c6f636b207468617420686173206e6f74206265656e2060648201527f7375626d69747465640000000000000000000000000000000000000000000000608482015260a4016107dc565b600354610dd85760405162461bcd60e51b815260206004820152604760248201527f4c324f75747075744f7261636c653a2063616e6e6f7420676574206f7574707560448201527f74206173206e6f206f7574707574732068617665206265656e207375626d697460648201527f7465642079657400000000000000000000000000000000000000000000000000608482015260a4016107dc565b6003546000905b80821015610e715760006002610df583856119aa565b610dff9190611a73565b90508460038281548110610e1557610e1561194c565b600091825260209091206003909102016002015470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff161015610e6757610e608160016119aa565b9250610e6b565b8091505b50610ddf565b5092915050565b60408051608081018252600080825260208201819052918101829052606081019190915260038281548110610eaf57610eaf61194c565b6000918252602091829020604080516080810182526003909302909101805473ffffffffffffffffffffffffffffffffffffffff1683526001810154938301939093526002909201546fffffffffffffffffffffffffffffffff808216938301939093527001000000000000000000000000000000009004909116606082015292915050565b60007f000000000000000000000000000000000000000000000000000000000000000060038381548110610f6b57610f6b61194c565b6000918252602090912060026003909202010154610f9b91906fffffffffffffffffffffffffffffffff166119aa565b92915050565b600060038281548110610fb657610fb661194c565b600091825260209091206003909102015473ffffffffffffffffffffffffffffffffffffffff1692915050565b604080516080810182526000808252602082018190529181018290526060810191909152600361101283610c98565b81548110610eaf57610eaf61194c565b60007f00000000000000000000000000000000000000000000000000000000000000006001548361105391906119c2565b61105d9190611aae565b600254610f9b91906119aa565b600354600090156110a8577f000000000000000000000000000000000000000000000000000000000000000061109e610569565b6105d791906119aa565b6105d7610569565b600054610100900460ff16158080156110d05750600054600160ff909116105b806110ea5750303b1580156110ea575060005460ff166001145b61115c5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016107dc565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156111ba57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b428211156112575760405162461bcd60e51b8152602060048201526044602482018190527f4c324f75747075744f7261636c653a207374617274696e67204c322074696d65908201527f7374616d70206d757374206265206c657373207468616e2063757272656e742060648201527f74696d6500000000000000000000000000000000000000000000000000000000608482015260a4016107dc565b6002829055600183905580156112c457600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461139a5760405162461bcd60e51b815260206004820152604160248201527f4c324f75747075744f7261636c653a206f6e6c792074686520636f6c6f73736560448201527f756d20636f6e74726163742063616e207265706c61636520616e206f7574707560648201527f7400000000000000000000000000000000000000000000000000000000000000608482015260a4016107dc565b73ffffffffffffffffffffffffffffffffffffffff81166114235760405162461bcd60e51b815260206004820152603060248201527f4c324f75747075744f7261636c653a207375626d69747465722061646472657360448201527f732063616e6e6f74206265207a65726f0000000000000000000000000000000060648201526084016107dc565b60035483106114c05760405162461bcd60e51b815260206004820152604660248201527f4c324f75747075744f7261636c653a2063616e6e6f74207265706c616365206160448201527f6e206f757470757420616674657220746865206c6174657374206f757470757460648201527f20696e6465780000000000000000000000000000000000000000000000000000608482015260a4016107dc565b6000600384815481106114d5576114d561194c565b6000918252602090912060039091020160028101549091507f00000000000000000000000000000000000000000000000000000000000000009061152b906fffffffffffffffffffffffffffffffff16426119c2565b106115c45760405162461bcd60e51b815260206004820152604860248201527f4c324f75747075744f7261636c653a2063616e6e6f74207265706c616365206160448201527f6e206f757470757420746861742068617320616c7265616479206265656e206660648201527f696e616c697a6564000000000000000000000000000000000000000000000000608482015260a4016107dc565b6001810183905580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff831617815560405183815284907fa1b831bb8b6b242db6d0988a6d21f869c610de9f703a5e45e1b7d3dc3137b9069060200160405180910390a250505050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b6060600061166d8361171e565b600101905060008167ffffffffffffffff81111561168d5761168d611aeb565b6040519080825280601f01601f1916602001820160405280156116b7576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a85049450846116c157509392505050565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310611767577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310611793576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc1000083106117b157662386f26fc10000830492506010015b6305f5e10083106117c9576305f5e100830492506008015b61271083106117dd57612710830492506004015b606483106117ef576064830492506002015b600a8310610f9b5760010192915050565b60006020828403121561181257600080fd5b5035919050565b60005b8381101561183457818101518382015260200161181c565b83811115611843576000848401525b50505050565b6020815260008251806020840152611868816040850160208701611819565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b600080600080608085870312156118b057600080fd5b5050823594602084013594506040840135936060013592509050565b600080604083850312156118df57600080fd5b50508035926020909101359150565b73ffffffffffffffffffffffffffffffffffffffff8116811461191057600080fd5b50565b60008060006060848603121561192857600080fd5b83359250602084013591506040840135611941816118ee565b809150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082198211156119bd576119bd61197b565b500190565b6000828210156119d4576119d461197b565b500390565b600084516119eb818460208901611819565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551611a27816001850160208a01611819565b60019201918201528351611a42816002840160208801611819565b0160020195945050505050565b600060208284031215611a6157600080fd5b8151611a6c816118ee565b9392505050565b600082611aa9577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611ae657611ae661197b565b500290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea164736f6c634300080f000a",
}

// L2OutputOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use L2OutputOracleMetaData.ABI instead.
var L2OutputOracleABI = L2OutputOracleMetaData.ABI

// L2OutputOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2OutputOracleMetaData.Bin instead.
var L2OutputOracleBin = L2OutputOracleMetaData.Bin

// DeployL2OutputOracle deploys a new Ethereum contract, binding an instance of L2OutputOracle to it.
func DeployL2OutputOracle(auth *bind.TransactOpts, backend bind.ContractBackend, _validatorPool common.Address, _colosseum common.Address, _submissionInterval *big.Int, _l2BlockTime *big.Int, _startingBlockNumber *big.Int, _startingTimestamp *big.Int, _finalizationPeriodSeconds *big.Int) (common.Address, *types.Transaction, *L2OutputOracle, error) {
	parsed, err := L2OutputOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2OutputOracleBin), backend, _validatorPool, _colosseum, _submissionInterval, _l2BlockTime, _startingBlockNumber, _startingTimestamp, _finalizationPeriodSeconds)
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

// FinalizedAt is a free data retrieval call binding the contract method 0xa48ea6de.
//
// Solidity: function finalizedAt(uint256 _outputIndex) view returns(uint256)
func (_L2OutputOracle *L2OutputOracleCaller) FinalizedAt(opts *bind.CallOpts, _outputIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _L2OutputOracle.contract.Call(opts, &out, "finalizedAt", _outputIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FinalizedAt is a free data retrieval call binding the contract method 0xa48ea6de.
//
// Solidity: function finalizedAt(uint256 _outputIndex) view returns(uint256)
func (_L2OutputOracle *L2OutputOracleSession) FinalizedAt(_outputIndex *big.Int) (*big.Int, error) {
	return _L2OutputOracle.Contract.FinalizedAt(&_L2OutputOracle.CallOpts, _outputIndex)
}

// FinalizedAt is a free data retrieval call binding the contract method 0xa48ea6de.
//
// Solidity: function finalizedAt(uint256 _outputIndex) view returns(uint256)
func (_L2OutputOracle *L2OutputOracleCallerSession) FinalizedAt(_outputIndex *big.Int) (*big.Int, error) {
	return _L2OutputOracle.Contract.FinalizedAt(&_L2OutputOracle.CallOpts, _outputIndex)
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

// IsFinalized is a free data retrieval call binding the contract method 0x33727c4d.
//
// Solidity: function isFinalized(uint256 _outputIndex) view returns(bool)
func (_L2OutputOracle *L2OutputOracleCaller) IsFinalized(opts *bind.CallOpts, _outputIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _L2OutputOracle.contract.Call(opts, &out, "isFinalized", _outputIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFinalized is a free data retrieval call binding the contract method 0x33727c4d.
//
// Solidity: function isFinalized(uint256 _outputIndex) view returns(bool)
func (_L2OutputOracle *L2OutputOracleSession) IsFinalized(_outputIndex *big.Int) (bool, error) {
	return _L2OutputOracle.Contract.IsFinalized(&_L2OutputOracle.CallOpts, _outputIndex)
}

// IsFinalized is a free data retrieval call binding the contract method 0x33727c4d.
//
// Solidity: function isFinalized(uint256 _outputIndex) view returns(bool)
func (_L2OutputOracle *L2OutputOracleCallerSession) IsFinalized(_outputIndex *big.Int) (bool, error) {
	return _L2OutputOracle.Contract.IsFinalized(&_L2OutputOracle.CallOpts, _outputIndex)
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

// SubmitL2Output is a paid mutator transaction binding the contract method 0x5a045f78.
//
// Solidity: function submitL2Output(bytes32 _outputRoot, uint256 _l2BlockNumber, bytes32 _l1BlockHash, uint256 _l1BlockNumber) payable returns()
func (_L2OutputOracle *L2OutputOracleTransactor) SubmitL2Output(opts *bind.TransactOpts, _outputRoot [32]byte, _l2BlockNumber *big.Int, _l1BlockHash [32]byte, _l1BlockNumber *big.Int) (*types.Transaction, error) {
	return _L2OutputOracle.contract.Transact(opts, "submitL2Output", _outputRoot, _l2BlockNumber, _l1BlockHash, _l1BlockNumber)
}

// SubmitL2Output is a paid mutator transaction binding the contract method 0x5a045f78.
//
// Solidity: function submitL2Output(bytes32 _outputRoot, uint256 _l2BlockNumber, bytes32 _l1BlockHash, uint256 _l1BlockNumber) payable returns()
func (_L2OutputOracle *L2OutputOracleSession) SubmitL2Output(_outputRoot [32]byte, _l2BlockNumber *big.Int, _l1BlockHash [32]byte, _l1BlockNumber *big.Int) (*types.Transaction, error) {
	return _L2OutputOracle.Contract.SubmitL2Output(&_L2OutputOracle.TransactOpts, _outputRoot, _l2BlockNumber, _l1BlockHash, _l1BlockNumber)
}

// SubmitL2Output is a paid mutator transaction binding the contract method 0x5a045f78.
//
// Solidity: function submitL2Output(bytes32 _outputRoot, uint256 _l2BlockNumber, bytes32 _l1BlockHash, uint256 _l1BlockNumber) payable returns()
func (_L2OutputOracle *L2OutputOracleTransactorSession) SubmitL2Output(_outputRoot [32]byte, _l2BlockNumber *big.Int, _l1BlockHash [32]byte, _l1BlockNumber *big.Int) (*types.Transaction, error) {
	return _L2OutputOracle.Contract.SubmitL2Output(&_L2OutputOracle.TransactOpts, _outputRoot, _l2BlockNumber, _l1BlockHash, _l1BlockNumber)
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
