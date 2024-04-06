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

// UpgradeGovernorMetaData contains all meta data concerning the UpgradeGovernor contract.
var UpgradeGovernorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"BALLOT_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CLOCK_MODE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"COUNTING_MODE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"EXTENDED_BALLOT_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancel\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"calldatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"descriptionHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"castVote\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"support\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"castVoteBySig\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"support\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"castVoteWithReason\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"support\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"castVoteWithReasonAndParams\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"support\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"castVoteWithReasonAndParamsBySig\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"support\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"clock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"calldatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"descriptionHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getVotes\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVotesWithParams\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasVoted\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hashProposal\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"calldatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"descriptionHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_timelock\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"_initialVotingDelay\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_initialVotingPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_initialProposalThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_votesQuorumFraction\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onERC1155BatchReceived\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onERC1155Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onERC721Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proposalDeadline\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalEta\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalProposer\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalSnapshot\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalVotes\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"againstVotes\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"forVotes\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"abstainVotes\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"propose\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"calldatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"queue\",\"inputs\":[{\"name\":\"targets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"calldatas\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"descriptionHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"quorum\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumDenominator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumNumerator\",\"inputs\":[{\"name\":\"timepoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorumNumerator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"relay\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setProposalThreshold\",\"inputs\":[{\"name\":\"newProposalThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setVotingDelay\",\"inputs\":[{\"name\":\"newVotingDelay\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setVotingPeriod\",\"inputs\":[{\"name\":\"newVotingPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"state\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIGovernorUpgradeable.ProposalState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"timelock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC5805Upgradeable\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateQuorumNumerator\",\"inputs\":[{\"name\":\"newQuorumNumerator\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTimelock\",\"inputs\":[{\"name\":\"newTimelock\",\"type\":\"address\",\"internalType\":\"contractTimelockControllerUpgradeable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"votingDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"votingPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProposalCanceled\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProposalCreated\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"proposer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"targets\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"signatures\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"calldatas\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"},{\"name\":\"voteStart\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"voteEnd\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"description\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProposalExecuted\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProposalQueued\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"eta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProposalThresholdSet\",\"inputs\":[{\"name\":\"oldProposalThreshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newProposalThreshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuorumNumeratorUpdated\",\"inputs\":[{\"name\":\"oldQuorumNumerator\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newQuorumNumerator\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TimelockChange\",\"inputs\":[{\"name\":\"oldTimelock\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTimelock\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VoteCast\",\"inputs\":[{\"name\":\"voter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"support\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VoteCastWithParams\",\"inputs\":[{\"name\":\"voter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"proposalId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"support\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"weight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"params\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VotingDelaySet\",\"inputs\":[{\"name\":\"oldVotingDelay\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newVotingDelay\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VotingPeriodSet\",\"inputs\":[{\"name\":\"oldVotingPeriod\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newVotingPeriod\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"Empty\",\"inputs\":[]}]",
	Bin: "0x60806040523480156200001157600080fd5b506200001c62000022565b620000e3565b600054610100900460ff16156200008f5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811614620000e1576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b615bb880620000f36000396000f3fe60806040526004361061030c5760003560e01c80637b3c71d31161019a578063c01f9e37116100e1578063ea0217cf1161008a578063f23a6e6111610064578063f23a6e6114610a9e578063f8ce560a14610ae3578063fc0c546a14610b0357600080fd5b8063ea0217cf14610a3e578063eb9019d414610a5e578063ece40cc114610a7e57600080fd5b8063d33219b4116100bb578063d33219b4146109a5578063dd4e2ba5146109c4578063deaaa7cc14610a0a57600080fd5b8063c01f9e3714610938578063c28bc2fa14610972578063c59057e41461098557600080fd5b80639a802a6d11610143578063ab58fb8e1161011d578063ab58fb8e146108be578063b58131b0146108de578063bc197c81146108f357600080fd5b80639a802a6d14610869578063a7713a7014610889578063a890c9101461089e57600080fd5b806386489ba91161017457806386489ba91461080957806391ddadf41461082957806397c3d3341461085557600080fd5b80637b3c71d3146107a15780637d5e81e2146107c157806384b0196e146107e157600080fd5b80633932abb11161025e578063544ffc9c116102075780635f398a14116101e15780635f398a141461074157806360c4247f1461076157806370b0f6601461078157600080fd5b8063544ffc9c1461068557806354fd4d50146106db578063567813881461072157600080fd5b806343859632116102385780634385963214610605578063452115d6146106505780634bf5d7e91461067057600080fd5b80633932abb1146105a35780633bccf4fd146105b85780633e4f49e6146105d857600080fd5b8063143489d0116102c05780632656227d1161029a5780632656227d146105255780632d63f693146105385780632fe3e2611461056f57600080fd5b8063143489d014610436578063150b7a0214610490578063160cbed71461050557600080fd5b806303420181116102f157806303420181146103d457806306f3f9e6146103f457806306fdde031461041457600080fd5b806301ffc9a71461037c57806302a251a3146103b157600080fd5b36610377573061031a610b24565b6001600160a01b0316146103755760405162461bcd60e51b815260206004820152601f60248201527f476f7665726e6f723a206d7573742073656e6420746f206578656375746f720060448201526064015b60405180910390fd5b005b600080fd5b34801561038857600080fd5b5061039c610397366004614a0c565b610b3e565b60405190151581526020015b60405180910390f35b3480156103bd57600080fd5b506103c6610b4f565b6040519081526020016103a8565b3480156103e057600080fd5b506103c66103ef366004614b8e565b610b5b565b34801561040057600080fd5b5061037561040f366004614c35565b610c53565b34801561042057600080fd5b50610429610d0d565b6040516103a89190614caa565b34801561044257600080fd5b50610478610451366004614c35565b600090815260fe60205260409020546801000000000000000090046001600160a01b031690565b6040516001600160a01b0390911681526020016103a8565b34801561049c57600080fd5b506104d46104ab366004614cd2565b7f150b7a0200000000000000000000000000000000000000000000000000000000949350505050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020016103a8565b34801561051157600080fd5b506103c6610520366004614eac565b610d9f565b6103c6610533366004614eac565b611045565b34801561054457600080fd5b506103c6610553366004614c35565b600090815260fe602052604090205467ffffffffffffffff1690565b34801561057b57600080fd5b506103c67fb3b3f3b703cd84ce352197dcff232b1b5d3cfb2025ce47cf04742d0651f1af8881565b3480156105af57600080fd5b506103c66111ab565b3480156105c457600080fd5b506103c66105d3366004614f3c565b6111b7565b3480156105e457600080fd5b506105f86105f3366004614c35565b61122d565b6040516103a89190614fb9565b34801561061157600080fd5b5061039c610620366004614ffa565b6000828152610161602090815260408083206001600160a01b038516845260030190915290205460ff1692915050565b34801561065c57600080fd5b506103c661066b366004614eac565b611238565b34801561067c57600080fd5b50610429611369565b34801561069157600080fd5b506106c06106a0366004614c35565b600090815261016160205260409020805460018201546002909201549092565b604080519384526020840192909252908201526060016103a8565b3480156106e757600080fd5b5060408051808201909152600581527f312e302e300000000000000000000000000000000000000000000000000000006020820152610429565b34801561072d57600080fd5b506103c661073c36600461502a565b61142f565b34801561074d57600080fd5b506103c661075c366004615056565b611458565b34801561076d57600080fd5b506103c661077c366004614c35565b6114a2565b34801561078d57600080fd5b5061037561079c366004614c35565b611597565b3480156107ad57600080fd5b506103c66107bc3660046150da565b61164e565b3480156107cd57600080fd5b506103c66107dc366004615134565b611696565b3480156107ed57600080fd5b506107f66116ad565b6040516103a89796959493929190615224565b34801561081557600080fd5b506103756108243660046152a0565b61176f565b34801561083557600080fd5b5061083e611951565b60405165ffffffffffff90911681526020016103a8565b34801561086157600080fd5b5060646103c6565b34801561087557600080fd5b506103c66108843660046152f9565b6119de565b34801561089557600080fd5b506103c66119f5565b3480156108aa57600080fd5b506103756108b9366004615352565b611a37565b3480156108ca57600080fd5b506103c66108d9366004614c35565b611aee565b3480156108ea57600080fd5b506103c6611ba3565b3480156108ff57600080fd5b506104d461090e36600461536f565b7fbc197c810000000000000000000000000000000000000000000000000000000095945050505050565b34801561094457600080fd5b506103c6610953366004614c35565b600090815260fe602052604090206001015467ffffffffffffffff1690565b610375610980366004615403565b611baf565b34801561099157600080fd5b506103c66109a0366004614eac565b611ce5565b3480156109b157600080fd5b506101f8546001600160a01b0316610478565b3480156109d057600080fd5b506040805180820190915260208082527f737570706f72743d627261766f2671756f72756d3d666f722c6162737461696e90820152610429565b348015610a1657600080fd5b506103c67f150214d74d59b7d1e90c73fc22ef3d991dd0a76b046543d4d80ab92d2a50328f81565b348015610a4a57600080fd5b50610375610a59366004614c35565b611d1f565b348015610a6a57600080fd5b506103c6610a79366004615447565b611dd6565b348015610a8a57600080fd5b50610375610a99366004614c35565b611df7565b348015610aaa57600080fd5b506104d4610ab9366004615473565b7ff23a6e610000000000000000000000000000000000000000000000000000000095945050505050565b348015610aef57600080fd5b506103c6610afe366004614c35565b611eae565b348015610b0f57600080fd5b5061019354610478906001600160a01b031681565b6000610b396101f8546001600160a01b031690565b905090565b6000610b4982611eb9565b92915050565b6000610b396101305490565b600080610bff610bf77fb3b3f3b703cd84ce352197dcff232b1b5d3cfb2025ce47cf04742d0651f1af888c8c8c8c604051610b979291906154dc565b60405180910390208b80519060200120604051602001610bdc959493929190948552602085019390935260ff9190911660408401526060830152608082015260a00190565b60405160208183030381529060405280519060200120611f0f565b868686611f57565b9050610c458a828b8b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508d9250611f75915050565b9a9950505050505050505050565b610c5b610b24565b6001600160a01b0316336001600160a01b031614610cbb5760405162461bcd60e51b815260206004820152601860248201527f476f7665726e6f723a206f6e6c79476f7665726e616e63650000000000000000604482015260640161036c565b30610cc4610b24565b6001600160a01b031614610d015760008036604051610ce49291906154dc565b604051809103902090505b80610cfa60ff6120e5565b03610cef57505b610d0a816121a2565b50565b606060fd8054610d1c906154ec565b80601f0160208091040260200160405190810160405280929190818152602001828054610d48906154ec565b8015610d955780601f10610d6a57610100808354040283529160200191610d95565b820191906000526020600020905b815481529060010190602001808311610d7857829003601f168201915b5050505050905090565b600080610dae86868686611ce5565b90506004610dbb8261122d565b6007811115610dcc57610dcc614f8a565b14610e3f5760405162461bcd60e51b815260206004820152602160248201527f476f7665726e6f723a2070726f706f73616c206e6f742073756363657373667560448201527f6c00000000000000000000000000000000000000000000000000000000000000606482015260840161036c565b6101f854604080517ff27a0c9200000000000000000000000000000000000000000000000000000000815290516000926001600160a01b03169163f27a0c929160048083019260209291908290030181865afa158015610ea3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ec7919061553f565b6101f8546040517fb1c5f4270000000000000000000000000000000000000000000000000000000081529192506001600160a01b03169063b1c5f42790610f1b908a908a908a906000908b906004016155e6565b602060405180830381865afa158015610f38573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f5c919061553f565b60008381526101f96020526040808220929092556101f85491517f8f2a0bb00000000000000000000000000000000000000000000000000000000081526001600160a01b0390921691638f2a0bb091610fc2918b918b918b91908b908990600401615634565b600060405180830381600087803b158015610fdc57600080fd5b505af1158015610ff0573d6000803e3d6000fd5b505050507f9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda289282824261102291906156bb565b604080519283526020830191909152015b60405180910390a15095945050505050565b60008061105486868686611ce5565b905060006110618261122d565b9050600481600781111561107757611077614f8a565b14806110945750600581600781111561109257611092614f8a565b145b6111065760405162461bcd60e51b815260206004820152602160248201527f476f7665726e6f723a2070726f706f73616c206e6f742073756363657373667560448201527f6c00000000000000000000000000000000000000000000000000000000000000606482015260840161036c565b600082815260fe60205260409081902060020180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055517f712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f906111729084815260200190565b60405180910390a16111878288888888612343565b61119482888888886123e5565b6111a182888888886123f2565b5095945050505050565b6000610b3961012f5490565b604080517f150214d74d59b7d1e90c73fc22ef3d991dd0a76b046543d4d80ab92d2a50328f602082015290810186905260ff85166060820152600090819061120590610bf790608001610bdc565b905061122287828860405180602001604052806000815250612438565b979650505050505050565b6000610b498261245b565b60008061124786868686611ce5565b905060006112548261122d565b600781111561126557611265614f8a565b146112b25760405162461bcd60e51b815260206004820152601c60248201527f476f7665726e6f723a20746f6f206c61746520746f2063616e63656c00000000604482015260640161036c565b600081815260fe60205260409020546801000000000000000090046001600160a01b0316336001600160a01b0316146113535760405162461bcd60e51b815260206004820152602260248201527f476f7665726e6f723a206f6e6c792070726f706f7365722063616e2063616e6360448201527f656c000000000000000000000000000000000000000000000000000000000000606482015260840161036c565b61135f868686866125da565b9695505050505050565b61019354604080517f4bf5d7e900000000000000000000000000000000000000000000000000000000815290516060926001600160a01b031691634bf5d7e99160048083019260009291908290030181865afa9250505080156113ee57506040513d6000823e601f3d908101601f191682016040526113eb91908101906156d3565b60015b61142a575060408051808201909152601d81527f6d6f64653d626c6f636b6e756d6265722666726f6d3d64656661756c74000000602082015290565b919050565b60008033905061145084828560405180602001604052806000815250612438565b949350505050565b60008033905061122287828888888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508a9250611f75915050565b6101c7546000908082036114bb5750506101c654919050565b60006101c76114cb600184615741565b815481106114db576114db615758565b60009182526020918290206040805180820190915291015463ffffffff81168083526401000000009091047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16928201929092529150841061155c57602001517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff169392505050565b611571611568856125e8565b6101c790612668565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16949350505050565b61159f610b24565b6001600160a01b0316336001600160a01b0316146115ff5760405162461bcd60e51b815260206004820152601860248201527f476f7665726e6f723a206f6e6c79476f7665726e616e63650000000000000000604482015260640161036c565b30611608610b24565b6001600160a01b03161461164557600080366040516116289291906154dc565b604051809103902090505b8061163e60ff6120e5565b0361163357505b610d0a81612731565b60008033905061135f86828787878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061243892505050565b60006116a485858585612774565b95945050505050565b6000606080600080600060606065546000801b1480156116cd5750606654155b6117195760405162461bcd60e51b815260206004820152601560248201527f4549503731323a20556e696e697469616c697a65640000000000000000000000604482015260640161036c565b611721612ca9565b611729612cb8565b604080516000808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009b939a50919850469750309650945092509050565b600054610100900460ff161580801561178f5750600054600160ff909116105b806117a95750303b1580156117a9575060005460ff166001145b61181b5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161036c565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561187957600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6118b76040518060400160405280600f81526020017f55706772616465476f7665726e6f720000000000000000000000000000000000815250612cc7565b6118c2858585612d90565b6118ca612e1d565b6118d387612e9c565b6118dc82612f22565b6118e586612fa8565b801561194857600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b61019354604080517f91ddadf400000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916391ddadf49160048083019260209291908290030181865afa9250505080156119d1575060408051601f3d908101601f191682019092526119ce91810190615787565b60015b61142a57610b394361302e565b60006119eb8484846130ac565b90505b9392505050565b6101c75460009015611a2f57611a0c6101c761313c565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16905090565b506101c65490565b611a3f610b24565b6001600160a01b0316336001600160a01b031614611a9f5760405162461bcd60e51b815260206004820152601860248201527f476f7665726e6f723a206f6e6c79476f7665726e616e63650000000000000000604482015260640161036c565b30611aa8610b24565b6001600160a01b031614611ae55760008036604051611ac89291906154dc565b604051809103902090505b80611ade60ff6120e5565b03611ad357505b610d0a81613182565b6101f85460008281526101f960205260408082205490517fd45c44350000000000000000000000000000000000000000000000000000000081526004810191909152909182916001600160a01b039091169063d45c443590602401602060405180830381865afa158015611b66573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b8a919061553f565b905080600114611b9a57806119ee565b60009392505050565b6000610b396101315490565b611bb7610b24565b6001600160a01b0316336001600160a01b031614611c175760405162461bcd60e51b815260206004820152601860248201527f476f7665726e6f723a206f6e6c79476f7665726e616e63650000000000000000604482015260640161036c565b30611c20610b24565b6001600160a01b031614611c5d5760008036604051611c409291906154dc565b604051809103902090505b80611c5660ff6120e5565b03611c4b57505b600080856001600160a01b0316858585604051611c7b9291906154dc565b60006040518083038185875af1925050503d8060008114611cb8576040519150601f19603f3d011682016040523d82523d6000602084013e611cbd565b606091505b50915091506119488282604051806060016040528060288152602001615b8460289139613205565b600084848484604051602001611cfe94939291906157af565b60408051601f19818403018152919052805160209091012095945050505050565b611d27610b24565b6001600160a01b0316336001600160a01b031614611d875760405162461bcd60e51b815260206004820152601860248201527f476f7665726e6f723a206f6e6c79476f7665726e616e63650000000000000000604482015260640161036c565b30611d90610b24565b6001600160a01b031614611dcd5760008036604051611db09291906154dc565b604051809103902090505b80611dc660ff6120e5565b03611dbb57505b610d0a8161321e565b60006119ee8383611df260408051602081019091526000815290565b6130ac565b611dff610b24565b6001600160a01b0316336001600160a01b031614611e5f5760405162461bcd60e51b815260206004820152601860248201527f476f7665726e6f723a206f6e6c79476f7665726e616e63650000000000000000604482015260640161036c565b30611e68610b24565b6001600160a01b031614611ea55760008036604051611e889291906154dc565b604051809103902090505b80611e9e60ff6120e5565b03611e9357505b610d0a816132d7565b6000610b498261331a565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f6e665ced000000000000000000000000000000000000000000000000000000001480610b495750610b49826133c2565b6000610b49611f1c61356c565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b6000806000611f6887878787613576565b915091506111a18161363a565b600085815260fe602052604081206001611f8e8861122d565b6007811115611f9f57611f9f614f8a565b146120125760405162461bcd60e51b815260206004820152602360248201527f476f7665726e6f723a20766f7465206e6f742063757272656e746c792061637460448201527f6976650000000000000000000000000000000000000000000000000000000000606482015260840161036c565b805460009061202d90889067ffffffffffffffff16866130ac565b905061203c888888848861379f565b835160000361209157866001600160a01b03167fb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda48988848960405161208494939291906157fa565b60405180910390a2611222565b866001600160a01b03167fe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb871289888489896040516120d2959493929190615822565b60405180910390a2979650505050505050565b600061210d8254600f81810b700100000000000000000000000000000000909204900b131590565b15612144576040517f3db2a12a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b508054600f0b6000818152600180840160205260408220805492905583547fffffffffffffffffffffffffffffffff000000000000000000000000000000001692016fffffffffffffffffffffffffffffffff169190911790915590565b606481111561223f5760405162461bcd60e51b815260206004820152604360248201527f476f7665726e6f72566f74657351756f72756d4672616374696f6e3a2071756f60448201527f72756d4e756d657261746f72206f7665722071756f72756d44656e6f6d696e6160648201527f746f720000000000000000000000000000000000000000000000000000000000608482015260a40161036c565b60006122496119f5565b9050801580159061225b57506101c754155b156122d65760408051808201909152600081526101c7906020810161227f84613992565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff90811690915282546001810184556000938452602093849020835194909301519091166401000000000263ffffffff909316929092179101555b6123046122f16122e4611951565b65ffffffffffff166125e8565b6122fa84613992565b6101c79190613a26565b505060408051828152602081018490527f0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997910160405180910390a15050565b3061234c610b24565b6001600160a01b0316146123de5760005b84518110156123dc57306001600160a01b031685828151811061238257612382615758565b60200260200101516001600160a01b0316036123cc576123cc8382815181106123ad576123ad615758565b60200260200101518051906020012060ff613a4190919063ffffffff16565b6123d581615868565b905061235d565b505b5050505050565b6123de8585858585613a93565b306123fb610b24565b6001600160a01b0316146123de5760ff54600f81810b700100000000000000000000000000000000909204900b13156123de57600060ff556123de565b60006116a48585858561245660408051602081019091526000815290565b611f75565b60008061246783613b21565b9050600481600781111561247d5761247d614f8a565b146124885792915050565b60008381526101f96020526040902054806124a4575092915050565b6101f8546040517f2ab0f529000000000000000000000000000000000000000000000000000000008152600481018390526001600160a01b0390911690632ab0f52990602401602060405180830381865afa158015612507573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061252b91906158a0565b1561253a575060079392505050565b6101f8546040517f584b153e000000000000000000000000000000000000000000000000000000008152600481018390526001600160a01b039091169063584b153e90602401602060405180830381865afa15801561259d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125c191906158a0565b156125d0575060059392505050565b5060029392505050565b60006116a485858585613c64565b600063ffffffff8211156126645760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203360448201527f3220626974730000000000000000000000000000000000000000000000000000606482015260840161036c565b5090565b8154600090818160058111156126c557600061268384613d33565b61268d9085615741565b60008881526020902090915081015463ffffffff90811690871610156126b5578091506126c3565b6126c08160016156bb565b92505b505b60006126d387878585613e1b565b90508015612724576126f8876126ea600184615741565b600091825260209091200190565b5464010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16611222565b6000979650505050505050565b61012f5460408051918252602082018390527fc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93910160405180910390a161012f55565b6000336127818184613e79565b6127cd5760405162461bcd60e51b815260206004820152601d60248201527f476f7665726e6f723a2070726f706f7365722072657374726963746564000000604482015260640161036c565b60006127d7611951565b65ffffffffffff1690506127e9611ba3565b6127f883610a79600185615741565b101561286c5760405162461bcd60e51b815260206004820152603160248201527f476f7665726e6f723a2070726f706f73657220766f7465732062656c6f77207060448201527f726f706f73616c207468726573686f6c64000000000000000000000000000000606482015260840161036c565b60006128818888888880519060200120611ce5565b905086518851146128fa5760405162461bcd60e51b815260206004820152602160248201527f476f7665726e6f723a20696e76616c69642070726f706f73616c206c656e677460448201527f6800000000000000000000000000000000000000000000000000000000000000606482015260840161036c565b85518851146129715760405162461bcd60e51b815260206004820152602160248201527f476f7665726e6f723a20696e76616c69642070726f706f73616c206c656e677460448201527f6800000000000000000000000000000000000000000000000000000000000000606482015260840161036c565b60008851116129c25760405162461bcd60e51b815260206004820152601860248201527f476f7665726e6f723a20656d7074792070726f706f73616c0000000000000000604482015260640161036c565b600081815260fe602052604090205467ffffffffffffffff1615612a4e5760405162461bcd60e51b815260206004820152602160248201527f476f7665726e6f723a2070726f706f73616c20616c726561647920657869737460448201527f7300000000000000000000000000000000000000000000000000000000000000606482015260840161036c565b6000612a586111ab565b612a6290846156bb565b90506000612a6e610b4f565b612a7890836156bb565b90506040518060e00160405280612a8e84613fc9565b67ffffffffffffffff1681526001600160a01b038716602082015260006040820152606001612abc83613fc9565b67ffffffffffffffff9081168252600060208084018290526040808501839052606094850183905288835260fe8252918290208551815492870151878501519186167fffffffff0000000000000000000000000000000000000000000000000000000090941693909317680100000000000000006001600160a01b039094168402177bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167c010000000000000000000000000000000000000000000000000000000060e09290921c91909102178155938501516080860151908416921c0217600183015560a08301516002909201805460c0909401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00009094169215157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1692909217610100931515939093029290921790558a517f7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e091859188918e918e91811115612c4657612c46614aa8565b604051908082528060200260200182016040528015612c7957816020015b6060815260200190600190039081612c645790505b508d88888f604051612c93999897969594939291906158c2565b60405180910390a1509098975050505050505050565b606060678054610d1c906154ec565b606060688054610d1c906154ec565b600054610100900460ff16612d445760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b612d8781612d8260408051808201909152600581527f312e302e30000000000000000000000000000000000000000000000000000000602082015290565b614049565b610d0a816140ee565b600054610100900460ff16612e0d5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b612e1883838361417b565b505050565b600054610100900460ff16612e9a5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b565b600054610100900460ff16612f195760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b610d0a81614213565b600054610100900460ff16612f9f5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b610d0a816142cb565b600054610100900460ff166130255760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b610d0a81614348565b600065ffffffffffff8211156126645760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203460448201527f3820626974730000000000000000000000000000000000000000000000000000606482015260840161036c565b610193546040517f3a46b1a80000000000000000000000000000000000000000000000000000000081526001600160a01b038581166004830152602482018590526000921690633a46b1a890604401602060405180830381865afa158015613118573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119eb919061553f565b80546000908015611b9a57613156836126ea600184615741565b5464010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff166119ee565b6101f854604080516001600160a01b03928316815291831660208301527f08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401910160405180910390a16101f880547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b606083156132145750816119ee565b6119ee83836143c5565b600081116132945760405162461bcd60e51b815260206004820152602760248201527f476f7665726e6f7253657474696e67733a20766f74696e6720706572696f642060448201527f746f6f206c6f7700000000000000000000000000000000000000000000000000606482015260840161036c565b6101305460408051918252602082018390527f7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828910160405180910390a161013055565b6101315460408051918252602082018390527fccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc05461910160405180910390a161013155565b60006064613327836114a2565b610193546040517f8e539e8c000000000000000000000000000000000000000000000000000000008152600481018690526001600160a01b0390911690638e539e8c90602401602060405180830381865afa15801561338a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133ae919061553f565b6133b8919061599a565b610b499190615a06565b60007f51159c06000000000000000000000000000000000000000000000000000000007fc6fba1f8000000000000000000000000000000000000000000000000000000007fbf26d897000000000000000000000000000000000000000000000000000000007f79dd796f000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000861682148061349c57507fffffffff00000000000000000000000000000000000000000000000000000000868116908216145b806134cb57507fffffffff00000000000000000000000000000000000000000000000000000000868116908516145b8061351757507fffffffff0000000000000000000000000000000000000000000000000000000086167f4e2312e000000000000000000000000000000000000000000000000000000000145b8061135f57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008716149695505050505050565b6000610b396143ef565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156135ad5750600090506003613631565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015613601573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661362a57600060019250925050613631565b9150600090505b94509492505050565b600081600481111561364e5761364e614f8a565b036136565750565b600181600481111561366a5761366a614f8a565b036136b75760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161036c565b60028160048111156136cb576136cb614f8a565b036137185760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161036c565b600381600481111561372c5761372c614f8a565b03610d0a5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f7565000000000000000000000000000000000000000000000000000000000000606482015260840161036c565b6000858152610161602090815260408083206001600160a01b0388168452600381019092529091205460ff161561383e5760405162461bcd60e51b815260206004820152602760248201527f476f7665726e6f72566f74696e6753696d706c653a20766f746520616c72656160448201527f6479206361737400000000000000000000000000000000000000000000000000606482015260840161036c565b6001600160a01b0385166000908152600382016020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905560ff84166138a8578281600001600082825461389d91906156bb565b909155506123dc9050565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60ff8516016138e6578281600101600082825461389d91906156bb565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe60ff851601613924578281600201600082825461389d91906156bb565b60405162461bcd60e51b815260206004820152603560248201527f476f7665726e6f72566f74696e6753696d706c653a20696e76616c696420766160448201527f6c756520666f7220656e756d20566f7465547970650000000000000000000000606482015260840161036c565b60007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8211156126645760405162461bcd60e51b815260206004820152602760248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203260448201527f3234206269747300000000000000000000000000000000000000000000000000606482015260840161036c565b600080613a34858585614463565b915091505b935093915050565b815470010000000000000000000000000000000090819004600f0b6000818152600180860160205260409091209390935583546fffffffffffffffffffffffffffffffff908116939091011602179055565b6101f8546040517fe38335e50000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063e38335e5903490613ae89088908890889060009089906004016155e6565b6000604051808303818588803b158015613b0157600080fd5b505af1158015613b15573d6000803e3d6000fd5b50505050505050505050565b600081815260fe60205260408120600281015460ff1615613b455750600792915050565b6002810154610100900460ff1615613b605750600292915050565b600083815260fe602052604081205467ffffffffffffffff1690819003613bc95760405162461bcd60e51b815260206004820152601d60248201527f476f7665726e6f723a20756e6b6e6f776e2070726f706f73616c206964000000604482015260640161036c565b6000613bd3611951565b65ffffffffffff169050808210613bef57506000949350505050565b600085815260fe602052604090206001015467ffffffffffffffff16818110613c1e5750600195945050505050565b613c278661465a565b8015613c4757506000868152610161602052604090208054600190910154115b15613c585750600495945050505050565b50600395945050505050565b600080613c73868686866146a8565b60008181526101f96020526040902054909150156116a4576101f85460008281526101f96020526040908190205490517fc4d252f50000000000000000000000000000000000000000000000000000000081526001600160a01b039092169163c4d252f591613ce89160040190815260200190565b600060405180830381600087803b158015613d0257600080fd5b505af1158015613d16573d6000803e3d6000fd5b50505060008281526101f960205260408120555095945050505050565b600081600003613d4557506000919050565b60006001613d52846147d1565b901c6001901b90506001818481613d6b57613d6b6159d7565b048201901c90506001818481613d8357613d836159d7565b048201901c90506001818481613d9b57613d9b6159d7565b048201901c90506001818481613db357613db36159d7565b048201901c90506001818481613dcb57613dcb6159d7565b048201901c90506001818481613de357613de36159d7565b048201901c90506001818481613dfb57613dfb6159d7565b048201901c90506119ee81828581613e1557613e156159d7565b04614865565b60005b81831015613e71576000613e32848461487b565b60008781526020902090915063ffffffff86169082015463ffffffff161115613e5d57809250613e6b565b613e688160016156bb565b93505b50613e1e565b509392505050565b80516000906034811015613e91576001915050610b49565b8281017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec01517fffffffffffffffffffffffff000000000000000000000000000000000000000081167f2370726f706f7365723d3078000000000000000000000000000000000000000014613f0b57600192505050610b49565b600080613f19602885615741565b90505b83811015613fa857600080613f68888481518110613f3c57613f3c615758565b01602001517fff0000000000000000000000000000000000000000000000000000000000000016614896565b9150915081613f805760019650505050505050610b49565b8060ff166004856001600160a01b0316901b179350505080613fa190615868565b9050613f1c565b50856001600160a01b0316816001600160a01b031614935050505092915050565b600067ffffffffffffffff8211156126645760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203660448201527f3420626974730000000000000000000000000000000000000000000000000000606482015260840161036c565b600054610100900460ff166140c65760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b60676140d28382615a87565b5060686140df8282615a87565b50506000606581905560665550565b600054610100900460ff1661416b5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b60fd6141778282615a87565b5050565b600054610100900460ff166141f85760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b61420183612731565b61420a8261321e565b612e18816132d7565b600054610100900460ff166142905760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b61019380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b600054610100900460ff16610d015760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b600054610100900460ff16611ae55760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b8151156143d55781518083602001fd5b8060405162461bcd60e51b815260040161036c9190614caa565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f61441a614982565b6144226149db565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b8254600090819080156145eb576000614481876126ea600185615741565b60408051808201909152905463ffffffff8082168084526401000000009092047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16602084015291925090871610156145185760405162461bcd60e51b815260206004820152601b60248201527f436865636b706f696e743a2064656372656173696e67206b6579730000000000604482015260640161036c565b805163ffffffff8088169116036145765784614539886126ea600186615741565b80547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff929092166401000000000263ffffffff9092169190911790556145db565b6040805180820190915263ffffffff80881682527bffffffffffffffffffffffffffffffffffffffffffffffffffffffff80881660208085019182528b54600181018d5560008d81529190912094519151909216640100000000029216919091179101555b602001519250839150613a399050565b50506040805180820190915263ffffffff80851682527bffffffffffffffffffffffffffffffffffffffffffffffffffffffff80851660208085019182528854600181018a5560008a815291822095519251909316640100000000029190931617920191909155905081613a39565b6000818152610161602052604081206002810154600182015461467d91906156bb565b600084815260fe602052604090205461469f9067ffffffffffffffff16611eae565b11159392505050565b6000806146b786868686611ce5565b905060006146c48261122d565b905060028160078111156146da576146da614f8a565b141580156146fa575060068160078111156146f7576146f7614f8a565b14155b80156147185750600781600781111561471557614715614f8a565b14155b6147645760405162461bcd60e51b815260206004820152601d60248201527f476f7665726e6f723a2070726f706f73616c206e6f7420616374697665000000604482015260640161036c565b600082815260fe60205260409081902060020180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16610100179055517f789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c906110339084815260200190565b600080608083901c156147e657608092831c92015b604083901c156147f857604092831c92015b602083901c1561480a57602092831c92015b601083901c1561481c57601092831c92015b600883901c1561482e57600892831c92015b600483901c1561484057600492831c92015b600283901c1561485257600292831c92015b600183901c15610b495760010192915050565b600081831061487457816119ee565b5090919050565b600061488a6002848418615a06565b6119ee908484166156bb565b60008060f883901c602f811180156148b15750603a8160ff16105b156148e4576001947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd09091019350915050565b8060ff1660401080156148fa575060478160ff16105b1561492d576001947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc99091019350915050565b8060ff166060108015614943575060678160ff16105b15614976576001947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa99091019350915050565b50600093849350915050565b60008061498d612ca9565b8051909150156149a4578051602090910120919050565b60655480156149b35792915050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4709250505090565b6000806149e6612cb8565b8051909150156149fd578051602090910120919050565b60665480156149b35792915050565b600060208284031215614a1e57600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146119ee57600080fd5b803560ff8116811461142a57600080fd5b60008083601f840112614a7157600080fd5b50813567ffffffffffffffff811115614a8957600080fd5b602083019150836020828501011115614aa157600080fd5b9250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715614b0057614b00614aa8565b604052919050565b600067ffffffffffffffff821115614b2257614b22614aa8565b50601f01601f191660200190565b6000614b43614b3e84614b08565b614ad7565b9050828152838383011115614b5757600080fd5b828260208301376000602084830101529392505050565b600082601f830112614b7f57600080fd5b6119ee83833560208501614b30565b60008060008060008060008060e0898b031215614baa57600080fd5b88359750614bba60208a01614a4e565b9650604089013567ffffffffffffffff80821115614bd757600080fd5b614be38c838d01614a5f565b909850965060608b0135915080821115614bfc57600080fd5b50614c098b828c01614b6e565b945050614c1860808a01614a4e565b925060a0890135915060c089013590509295985092959890939650565b600060208284031215614c4757600080fd5b5035919050565b60005b83811015614c69578181015183820152602001614c51565b83811115614c78576000848401525b50505050565b60008151808452614c96816020860160208601614c4e565b601f01601f19169290920160200192915050565b6020815260006119ee6020830184614c7e565b6001600160a01b0381168114610d0a57600080fd5b60008060008060808587031215614ce857600080fd5b8435614cf381614cbd565b93506020850135614d0381614cbd565b925060408501359150606085013567ffffffffffffffff811115614d2657600080fd5b614d3287828801614b6e565b91505092959194509250565b600067ffffffffffffffff821115614d5857614d58614aa8565b5060051b60200190565b600082601f830112614d7357600080fd5b81356020614d83614b3e83614d3e565b82815260059290921b84018101918181019086841115614da257600080fd5b8286015b84811015614dc6578035614db981614cbd565b8352918301918301614da6565b509695505050505050565b600082601f830112614de257600080fd5b81356020614df2614b3e83614d3e565b82815260059290921b84018101918181019086841115614e1157600080fd5b8286015b84811015614dc65780358352918301918301614e15565b600082601f830112614e3d57600080fd5b81356020614e4d614b3e83614d3e565b82815260059290921b84018101918181019086841115614e6c57600080fd5b8286015b84811015614dc657803567ffffffffffffffff811115614e905760008081fd5b614e9e8986838b0101614b6e565b845250918301918301614e70565b60008060008060808587031215614ec257600080fd5b843567ffffffffffffffff80821115614eda57600080fd5b614ee688838901614d62565b95506020870135915080821115614efc57600080fd5b614f0888838901614dd1565b94506040870135915080821115614f1e57600080fd5b50614f2b87828801614e2c565b949793965093946060013593505050565b600080600080600060a08688031215614f5457600080fd5b85359450614f6460208701614a4e565b9350614f7260408701614a4e565b94979396509394606081013594506080013592915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6020810160088310614ff4577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b6000806040838503121561500d57600080fd5b82359150602083013561501f81614cbd565b809150509250929050565b6000806040838503121561503d57600080fd5b8235915061504d60208401614a4e565b90509250929050565b60008060008060006080868803121561506e57600080fd5b8535945061507e60208701614a4e565b9350604086013567ffffffffffffffff8082111561509b57600080fd5b6150a789838a01614a5f565b909550935060608801359150808211156150c057600080fd5b506150cd88828901614b6e565b9150509295509295909350565b600080600080606085870312156150f057600080fd5b8435935061510060208601614a4e565b9250604085013567ffffffffffffffff81111561511c57600080fd5b61512887828801614a5f565b95989497509550505050565b6000806000806080858703121561514a57600080fd5b843567ffffffffffffffff8082111561516257600080fd5b61516e88838901614d62565b9550602087013591508082111561518457600080fd5b61519088838901614dd1565b945060408701359150808211156151a657600080fd5b6151b288838901614e2c565b935060608701359150808211156151c857600080fd5b508501601f810187136151da57600080fd5b614d3287823560208401614b30565b600081518084526020808501945080840160005b83811015615219578151875295820195908201906001016151fd565b509495945050505050565b7fff000000000000000000000000000000000000000000000000000000000000008816815260e06020820152600061525f60e0830189614c7e565b82810360408401526152718189614c7e565b90508660608401526001600160a01b03861660808401528460a084015282810360c0840152610c4581856151e9565b60008060008060008060c087890312156152b957600080fd5b86356152c481614cbd565b955060208701356152d481614cbd565b95989597505050506040840135936060810135936080820135935060a0909101359150565b60008060006060848603121561530e57600080fd5b833561531981614cbd565b925060208401359150604084013567ffffffffffffffff81111561533c57600080fd5b61534886828701614b6e565b9150509250925092565b60006020828403121561536457600080fd5b81356119ee81614cbd565b600080600080600060a0868803121561538757600080fd5b853561539281614cbd565b945060208601356153a281614cbd565b9350604086013567ffffffffffffffff808211156153bf57600080fd5b6153cb89838a01614dd1565b945060608801359150808211156153e157600080fd5b6153ed89838a01614dd1565b935060808801359150808211156150c057600080fd5b6000806000806060858703121561541957600080fd5b843561542481614cbd565b935060208501359250604085013567ffffffffffffffff81111561511c57600080fd5b6000806040838503121561545a57600080fd5b823561546581614cbd565b946020939093013593505050565b600080600080600060a0868803121561548b57600080fd5b853561549681614cbd565b945060208601356154a681614cbd565b93506040860135925060608601359150608086013567ffffffffffffffff8111156154d057600080fd5b6150cd88828901614b6e565b8183823760009101908152919050565b600181811c9082168061550057607f821691505b602082108103615539577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b60006020828403121561555157600080fd5b5051919050565b600081518084526020808501945080840160005b838110156152195781516001600160a01b03168752958201959082019060010161556c565b600081518084526020808501808196508360051b8101915082860160005b858110156155d95782840389526155c7848351614c7e565b988501989350908401906001016155af565b5091979650505050505050565b60a0815260006155f960a0830188615558565b828103602084015261560b81886151e9565b9050828103604084015261561f8187615591565b60608401959095525050608001529392505050565b60c08152600061564760c0830189615558565b828103602084015261565981896151e9565b9050828103604084015261566d8188615591565b60608401969096525050608081019290925260a0909101529392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082198211156156ce576156ce61568c565b500190565b6000602082840312156156e557600080fd5b815167ffffffffffffffff8111156156fc57600080fd5b8201601f8101841361570d57600080fd5b805161571b614b3e82614b08565b81815285602083850101111561573057600080fd5b6116a4826020830160208601614c4e565b6000828210156157535761575361568c565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006020828403121561579957600080fd5b815165ffffffffffff811681146119ee57600080fd5b6080815260006157c26080830187615558565b82810360208401526157d481876151e9565b905082810360408401526157e88186615591565b91505082606083015295945050505050565b84815260ff8416602082015282604082015260806060820152600061135f6080830184614c7e565b85815260ff8516602082015283604082015260a06060820152600061584a60a0830185614c7e565b828103608084015261585c8185614c7e565b98975050505050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036158995761589961568c565b5060010190565b6000602082840312156158b257600080fd5b815180151581146119ee57600080fd5b60006101208b835260206001600160a01b038c16818501528160408501526158ec8285018c615558565b91508382036060850152615900828b6151e9565b915083820360808501528189518084528284019150828160051b850101838c0160005b8381101561595157601f1987840301855261593f838351614c7e565b94860194925090850190600101615923565b505086810360a0880152615965818c615591565b9450505050508560c08401528460e084015282810361010084015261598a8185614c7e565b9c9b505050505050505050505050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156159d2576159d261568c565b500290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082615a3c577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b601f821115612e1857600081815260208120601f850160051c81016020861015615a685750805b601f850160051c820191505b818110156123dc57828155600101615a74565b815167ffffffffffffffff811115615aa157615aa1614aa8565b615ab581615aaf84546154ec565b84615a41565b602080601f831160018114615b085760008415615ad25750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b1785556123dc565b600085815260208120601f198616915b82811015615b3757888601518255948401946001909101908401615b18565b5085821015615b7357878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b0190555056fe476f7665726e6f723a2072656c617920726576657274656420776974686f7574206d657373616765a164736f6c634300080f000a",
}

// UpgradeGovernorABI is the input ABI used to generate the binding from.
// Deprecated: Use UpgradeGovernorMetaData.ABI instead.
var UpgradeGovernorABI = UpgradeGovernorMetaData.ABI

// UpgradeGovernorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UpgradeGovernorMetaData.Bin instead.
var UpgradeGovernorBin = UpgradeGovernorMetaData.Bin

// DeployUpgradeGovernor deploys a new Ethereum contract, binding an instance of UpgradeGovernor to it.
func DeployUpgradeGovernor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UpgradeGovernor, error) {
	parsed, err := UpgradeGovernorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UpgradeGovernorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UpgradeGovernor{UpgradeGovernorCaller: UpgradeGovernorCaller{contract: contract}, UpgradeGovernorTransactor: UpgradeGovernorTransactor{contract: contract}, UpgradeGovernorFilterer: UpgradeGovernorFilterer{contract: contract}}, nil
}

// UpgradeGovernor is an auto generated Go binding around an Ethereum contract.
type UpgradeGovernor struct {
	UpgradeGovernorCaller     // Read-only binding to the contract
	UpgradeGovernorTransactor // Write-only binding to the contract
	UpgradeGovernorFilterer   // Log filterer for contract events
}

// UpgradeGovernorCaller is an auto generated read-only Go binding around an Ethereum contract.
type UpgradeGovernorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeGovernorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UpgradeGovernorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeGovernorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UpgradeGovernorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeGovernorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UpgradeGovernorSession struct {
	Contract     *UpgradeGovernor  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UpgradeGovernorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UpgradeGovernorCallerSession struct {
	Contract *UpgradeGovernorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// UpgradeGovernorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UpgradeGovernorTransactorSession struct {
	Contract     *UpgradeGovernorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// UpgradeGovernorRaw is an auto generated low-level Go binding around an Ethereum contract.
type UpgradeGovernorRaw struct {
	Contract *UpgradeGovernor // Generic contract binding to access the raw methods on
}

// UpgradeGovernorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UpgradeGovernorCallerRaw struct {
	Contract *UpgradeGovernorCaller // Generic read-only contract binding to access the raw methods on
}

// UpgradeGovernorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UpgradeGovernorTransactorRaw struct {
	Contract *UpgradeGovernorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUpgradeGovernor creates a new instance of UpgradeGovernor, bound to a specific deployed contract.
func NewUpgradeGovernor(address common.Address, backend bind.ContractBackend) (*UpgradeGovernor, error) {
	contract, err := bindUpgradeGovernor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UpgradeGovernor{UpgradeGovernorCaller: UpgradeGovernorCaller{contract: contract}, UpgradeGovernorTransactor: UpgradeGovernorTransactor{contract: contract}, UpgradeGovernorFilterer: UpgradeGovernorFilterer{contract: contract}}, nil
}

// NewUpgradeGovernorCaller creates a new read-only instance of UpgradeGovernor, bound to a specific deployed contract.
func NewUpgradeGovernorCaller(address common.Address, caller bind.ContractCaller) (*UpgradeGovernorCaller, error) {
	contract, err := bindUpgradeGovernor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeGovernorCaller{contract: contract}, nil
}

// NewUpgradeGovernorTransactor creates a new write-only instance of UpgradeGovernor, bound to a specific deployed contract.
func NewUpgradeGovernorTransactor(address common.Address, transactor bind.ContractTransactor) (*UpgradeGovernorTransactor, error) {
	contract, err := bindUpgradeGovernor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeGovernorTransactor{contract: contract}, nil
}

// NewUpgradeGovernorFilterer creates a new log filterer instance of UpgradeGovernor, bound to a specific deployed contract.
func NewUpgradeGovernorFilterer(address common.Address, filterer bind.ContractFilterer) (*UpgradeGovernorFilterer, error) {
	contract, err := bindUpgradeGovernor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UpgradeGovernorFilterer{contract: contract}, nil
}

// bindUpgradeGovernor binds a generic wrapper to an already deployed contract.
func bindUpgradeGovernor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UpgradeGovernorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeGovernor *UpgradeGovernorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpgradeGovernor.Contract.UpgradeGovernorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeGovernor *UpgradeGovernorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.UpgradeGovernorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeGovernor *UpgradeGovernorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.UpgradeGovernorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeGovernor *UpgradeGovernorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpgradeGovernor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeGovernor *UpgradeGovernorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeGovernor *UpgradeGovernorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.contract.Transact(opts, method, params...)
}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_UpgradeGovernor *UpgradeGovernorCaller) BALLOTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UpgradeGovernor.contract.Call(opts, &out, "BALLOT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_UpgradeGovernor *UpgradeGovernorSession) BALLOTTYPEHASH() ([32]byte, error) {
	return _UpgradeGovernor.Contract.BALLOTTYPEHASH(&_UpgradeGovernor.CallOpts)
}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_UpgradeGovernor *UpgradeGovernorCallerSession) BALLOTTYPEHASH() ([32]byte, error) {
	return _UpgradeGovernor.Contract.BALLOTTYPEHASH(&_UpgradeGovernor.CallOpts)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_UpgradeGovernor *UpgradeGovernorCaller) CLOCKMODE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UpgradeGovernor.contract.Call(opts, &out, "CLOCK_MODE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_UpgradeGovernor *UpgradeGovernorSession) CLOCKMODE() (string, error) {
	return _UpgradeGovernor.Contract.CLOCKMODE(&_UpgradeGovernor.CallOpts)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_UpgradeGovernor *UpgradeGovernorCallerSession) CLOCKMODE() (string, error) {
	return _UpgradeGovernor.Contract.CLOCKMODE(&_UpgradeGovernor.CallOpts)
}

// COUNTINGMODE is a free data retrieval call binding the contract method 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (_UpgradeGovernor *UpgradeGovernorCaller) COUNTINGMODE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UpgradeGovernor.contract.Call(opts, &out, "COUNTING_MODE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// COUNTINGMODE is a free data retrieval call binding the contract method 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (_UpgradeGovernor *UpgradeGovernorSession) COUNTINGMODE() (string, error) {
	return _UpgradeGovernor.Contract.COUNTINGMODE(&_UpgradeGovernor.CallOpts)
}

// COUNTINGMODE is a free data retrieval call binding the contract method 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (_UpgradeGovernor *UpgradeGovernorCallerSession) COUNTINGMODE() (string, error) {
	return _UpgradeGovernor.Contract.COUNTINGMODE(&_UpgradeGovernor.CallOpts)
}

// EXTENDEDBALLOTTYPEHASH is a free data retrieval call binding the contract method 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (_UpgradeGovernor *UpgradeGovernorCaller) EXTENDEDBALLOTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UpgradeGovernor.contract.Call(opts, &out, "EXTENDED_BALLOT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EXTENDEDBALLOTTYPEHASH is a free data retrieval call binding the contract method 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (_UpgradeGovernor *UpgradeGovernorSession) EXTENDEDBALLOTTYPEHASH() ([32]byte, error) {
	return _UpgradeGovernor.Contract.EXTENDEDBALLOTTYPEHASH(&_UpgradeGovernor.CallOpts)
}

// EXTENDEDBALLOTTYPEHASH is a free data retrieval call binding the contract method 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (_UpgradeGovernor *UpgradeGovernorCallerSession) EXTENDEDBALLOTTYPEHASH() ([32]byte, error) {
	return _UpgradeGovernor.Contract.EXTENDEDBALLOTTYPEHASH(&_UpgradeGovernor.CallOpts)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_UpgradeGovernor *UpgradeGovernorCaller) Clock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeGovernor.contract.Call(opts, &out, "clock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_UpgradeGovernor *UpgradeGovernorSession) Clock() (*big.Int, error) {
	return _UpgradeGovernor.Contract.Clock(&_UpgradeGovernor.CallOpts)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_UpgradeGovernor *UpgradeGovernorCallerSession) Clock() (*big.Int, error) {
	return _UpgradeGovernor.Contract.Clock(&_UpgradeGovernor.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_UpgradeGovernor *UpgradeGovernorCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _UpgradeGovernor.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_UpgradeGovernor *UpgradeGovernorSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _UpgradeGovernor.Contract.Eip712Domain(&_UpgradeGovernor.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_UpgradeGovernor *UpgradeGovernorCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _UpgradeGovernor.Contract.Eip712Domain(&_UpgradeGovernor.CallOpts)
}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorCaller) GetVotes(opts *bind.CallOpts, account common.Address, timepoint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeGovernor.contract.Call(opts, &out, "getVotes", account, timepoint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorSession) GetVotes(account common.Address, timepoint *big.Int) (*big.Int, error) {
	return _UpgradeGovernor.Contract.GetVotes(&_UpgradeGovernor.CallOpts, account, timepoint)
}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorCallerSession) GetVotes(account common.Address, timepoint *big.Int) (*big.Int, error) {
	return _UpgradeGovernor.Contract.GetVotes(&_UpgradeGovernor.CallOpts, account, timepoint)
}

// GetVotesWithParams is a free data retrieval call binding the contract method 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorCaller) GetVotesWithParams(opts *bind.CallOpts, account common.Address, timepoint *big.Int, params []byte) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeGovernor.contract.Call(opts, &out, "getVotesWithParams", account, timepoint, params)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotesWithParams is a free data retrieval call binding the contract method 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorSession) GetVotesWithParams(account common.Address, timepoint *big.Int, params []byte) (*big.Int, error) {
	return _UpgradeGovernor.Contract.GetVotesWithParams(&_UpgradeGovernor.CallOpts, account, timepoint, params)
}

// GetVotesWithParams is a free data retrieval call binding the contract method 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorCallerSession) GetVotesWithParams(account common.Address, timepoint *big.Int, params []byte) (*big.Int, error) {
	return _UpgradeGovernor.Contract.GetVotesWithParams(&_UpgradeGovernor.CallOpts, account, timepoint, params)
}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_UpgradeGovernor *UpgradeGovernorCaller) HasVoted(opts *bind.CallOpts, proposalId *big.Int, account common.Address) (bool, error) {
	var out []interface{}
	err := _UpgradeGovernor.contract.Call(opts, &out, "hasVoted", proposalId, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_UpgradeGovernor *UpgradeGovernorSession) HasVoted(proposalId *big.Int, account common.Address) (bool, error) {
	return _UpgradeGovernor.Contract.HasVoted(&_UpgradeGovernor.CallOpts, proposalId, account)
}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_UpgradeGovernor *UpgradeGovernorCallerSession) HasVoted(proposalId *big.Int, account common.Address) (bool, error) {
	return _UpgradeGovernor.Contract.HasVoted(&_UpgradeGovernor.CallOpts, proposalId, account)
}

// HashProposal is a free data retrieval call binding the contract method 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorCaller) HashProposal(opts *bind.CallOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeGovernor.contract.Call(opts, &out, "hashProposal", targets, values, calldatas, descriptionHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashProposal is a free data retrieval call binding the contract method 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorSession) HashProposal(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	return _UpgradeGovernor.Contract.HashProposal(&_UpgradeGovernor.CallOpts, targets, values, calldatas, descriptionHash)
}

// HashProposal is a free data retrieval call binding the contract method 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorCallerSession) HashProposal(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	return _UpgradeGovernor.Contract.HashProposal(&_UpgradeGovernor.CallOpts, targets, values, calldatas, descriptionHash)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UpgradeGovernor *UpgradeGovernorCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UpgradeGovernor.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UpgradeGovernor *UpgradeGovernorSession) Name() (string, error) {
	return _UpgradeGovernor.Contract.Name(&_UpgradeGovernor.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UpgradeGovernor *UpgradeGovernorCallerSession) Name() (string, error) {
	return _UpgradeGovernor.Contract.Name(&_UpgradeGovernor.CallOpts)
}

// ProposalDeadline is a free data retrieval call binding the contract method 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorCaller) ProposalDeadline(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeGovernor.contract.Call(opts, &out, "proposalDeadline", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalDeadline is a free data retrieval call binding the contract method 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorSession) ProposalDeadline(proposalId *big.Int) (*big.Int, error) {
	return _UpgradeGovernor.Contract.ProposalDeadline(&_UpgradeGovernor.CallOpts, proposalId)
}

// ProposalDeadline is a free data retrieval call binding the contract method 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorCallerSession) ProposalDeadline(proposalId *big.Int) (*big.Int, error) {
	return _UpgradeGovernor.Contract.ProposalDeadline(&_UpgradeGovernor.CallOpts, proposalId)
}

// ProposalEta is a free data retrieval call binding the contract method 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorCaller) ProposalEta(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeGovernor.contract.Call(opts, &out, "proposalEta", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalEta is a free data retrieval call binding the contract method 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorSession) ProposalEta(proposalId *big.Int) (*big.Int, error) {
	return _UpgradeGovernor.Contract.ProposalEta(&_UpgradeGovernor.CallOpts, proposalId)
}

// ProposalEta is a free data retrieval call binding the contract method 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorCallerSession) ProposalEta(proposalId *big.Int) (*big.Int, error) {
	return _UpgradeGovernor.Contract.ProposalEta(&_UpgradeGovernor.CallOpts, proposalId)
}

// ProposalProposer is a free data retrieval call binding the contract method 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (_UpgradeGovernor *UpgradeGovernorCaller) ProposalProposer(opts *bind.CallOpts, proposalId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _UpgradeGovernor.contract.Call(opts, &out, "proposalProposer", proposalId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposalProposer is a free data retrieval call binding the contract method 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (_UpgradeGovernor *UpgradeGovernorSession) ProposalProposer(proposalId *big.Int) (common.Address, error) {
	return _UpgradeGovernor.Contract.ProposalProposer(&_UpgradeGovernor.CallOpts, proposalId)
}

// ProposalProposer is a free data retrieval call binding the contract method 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (_UpgradeGovernor *UpgradeGovernorCallerSession) ProposalProposer(proposalId *big.Int) (common.Address, error) {
	return _UpgradeGovernor.Contract.ProposalProposer(&_UpgradeGovernor.CallOpts, proposalId)
}

// ProposalSnapshot is a free data retrieval call binding the contract method 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorCaller) ProposalSnapshot(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeGovernor.contract.Call(opts, &out, "proposalSnapshot", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalSnapshot is a free data retrieval call binding the contract method 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorSession) ProposalSnapshot(proposalId *big.Int) (*big.Int, error) {
	return _UpgradeGovernor.Contract.ProposalSnapshot(&_UpgradeGovernor.CallOpts, proposalId)
}

// ProposalSnapshot is a free data retrieval call binding the contract method 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorCallerSession) ProposalSnapshot(proposalId *big.Int) (*big.Int, error) {
	return _UpgradeGovernor.Contract.ProposalSnapshot(&_UpgradeGovernor.CallOpts, proposalId)
}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorCaller) ProposalThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeGovernor.contract.Call(opts, &out, "proposalThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorSession) ProposalThreshold() (*big.Int, error) {
	return _UpgradeGovernor.Contract.ProposalThreshold(&_UpgradeGovernor.CallOpts)
}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorCallerSession) ProposalThreshold() (*big.Int, error) {
	return _UpgradeGovernor.Contract.ProposalThreshold(&_UpgradeGovernor.CallOpts)
}

// ProposalVotes is a free data retrieval call binding the contract method 0x544ffc9c.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (_UpgradeGovernor *UpgradeGovernorCaller) ProposalVotes(opts *bind.CallOpts, proposalId *big.Int) (struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}, error) {
	var out []interface{}
	err := _UpgradeGovernor.contract.Call(opts, &out, "proposalVotes", proposalId)

	outstruct := new(struct {
		AgainstVotes *big.Int
		ForVotes     *big.Int
		AbstainVotes *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AgainstVotes = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ForVotes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AbstainVotes = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProposalVotes is a free data retrieval call binding the contract method 0x544ffc9c.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (_UpgradeGovernor *UpgradeGovernorSession) ProposalVotes(proposalId *big.Int) (struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}, error) {
	return _UpgradeGovernor.Contract.ProposalVotes(&_UpgradeGovernor.CallOpts, proposalId)
}

// ProposalVotes is a free data retrieval call binding the contract method 0x544ffc9c.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (_UpgradeGovernor *UpgradeGovernorCallerSession) ProposalVotes(proposalId *big.Int) (struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}, error) {
	return _UpgradeGovernor.Contract.ProposalVotes(&_UpgradeGovernor.CallOpts, proposalId)
}

// Quorum is a free data retrieval call binding the contract method 0xf8ce560a.
//
// Solidity: function quorum(uint256 blockNumber) view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorCaller) Quorum(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeGovernor.contract.Call(opts, &out, "quorum", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quorum is a free data retrieval call binding the contract method 0xf8ce560a.
//
// Solidity: function quorum(uint256 blockNumber) view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorSession) Quorum(blockNumber *big.Int) (*big.Int, error) {
	return _UpgradeGovernor.Contract.Quorum(&_UpgradeGovernor.CallOpts, blockNumber)
}

// Quorum is a free data retrieval call binding the contract method 0xf8ce560a.
//
// Solidity: function quorum(uint256 blockNumber) view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorCallerSession) Quorum(blockNumber *big.Int) (*big.Int, error) {
	return _UpgradeGovernor.Contract.Quorum(&_UpgradeGovernor.CallOpts, blockNumber)
}

// QuorumDenominator is a free data retrieval call binding the contract method 0x97c3d334.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorCaller) QuorumDenominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeGovernor.contract.Call(opts, &out, "quorumDenominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumDenominator is a free data retrieval call binding the contract method 0x97c3d334.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorSession) QuorumDenominator() (*big.Int, error) {
	return _UpgradeGovernor.Contract.QuorumDenominator(&_UpgradeGovernor.CallOpts)
}

// QuorumDenominator is a free data retrieval call binding the contract method 0x97c3d334.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorCallerSession) QuorumDenominator() (*big.Int, error) {
	return _UpgradeGovernor.Contract.QuorumDenominator(&_UpgradeGovernor.CallOpts)
}

// QuorumNumerator is a free data retrieval call binding the contract method 0x60c4247f.
//
// Solidity: function quorumNumerator(uint256 timepoint) view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorCaller) QuorumNumerator(opts *bind.CallOpts, timepoint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeGovernor.contract.Call(opts, &out, "quorumNumerator", timepoint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumNumerator is a free data retrieval call binding the contract method 0x60c4247f.
//
// Solidity: function quorumNumerator(uint256 timepoint) view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorSession) QuorumNumerator(timepoint *big.Int) (*big.Int, error) {
	return _UpgradeGovernor.Contract.QuorumNumerator(&_UpgradeGovernor.CallOpts, timepoint)
}

// QuorumNumerator is a free data retrieval call binding the contract method 0x60c4247f.
//
// Solidity: function quorumNumerator(uint256 timepoint) view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorCallerSession) QuorumNumerator(timepoint *big.Int) (*big.Int, error) {
	return _UpgradeGovernor.Contract.QuorumNumerator(&_UpgradeGovernor.CallOpts, timepoint)
}

// QuorumNumerator0 is a free data retrieval call binding the contract method 0xa7713a70.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorCaller) QuorumNumerator0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeGovernor.contract.Call(opts, &out, "quorumNumerator0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumNumerator0 is a free data retrieval call binding the contract method 0xa7713a70.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorSession) QuorumNumerator0() (*big.Int, error) {
	return _UpgradeGovernor.Contract.QuorumNumerator0(&_UpgradeGovernor.CallOpts)
}

// QuorumNumerator0 is a free data retrieval call binding the contract method 0xa7713a70.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorCallerSession) QuorumNumerator0() (*big.Int, error) {
	return _UpgradeGovernor.Contract.QuorumNumerator0(&_UpgradeGovernor.CallOpts)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_UpgradeGovernor *UpgradeGovernorCaller) State(opts *bind.CallOpts, proposalId *big.Int) (uint8, error) {
	var out []interface{}
	err := _UpgradeGovernor.contract.Call(opts, &out, "state", proposalId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_UpgradeGovernor *UpgradeGovernorSession) State(proposalId *big.Int) (uint8, error) {
	return _UpgradeGovernor.Contract.State(&_UpgradeGovernor.CallOpts, proposalId)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_UpgradeGovernor *UpgradeGovernorCallerSession) State(proposalId *big.Int) (uint8, error) {
	return _UpgradeGovernor.Contract.State(&_UpgradeGovernor.CallOpts, proposalId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UpgradeGovernor *UpgradeGovernorCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _UpgradeGovernor.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UpgradeGovernor *UpgradeGovernorSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _UpgradeGovernor.Contract.SupportsInterface(&_UpgradeGovernor.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UpgradeGovernor *UpgradeGovernorCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _UpgradeGovernor.Contract.SupportsInterface(&_UpgradeGovernor.CallOpts, interfaceId)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (_UpgradeGovernor *UpgradeGovernorCaller) Timelock(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UpgradeGovernor.contract.Call(opts, &out, "timelock")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (_UpgradeGovernor *UpgradeGovernorSession) Timelock() (common.Address, error) {
	return _UpgradeGovernor.Contract.Timelock(&_UpgradeGovernor.CallOpts)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (_UpgradeGovernor *UpgradeGovernorCallerSession) Timelock() (common.Address, error) {
	return _UpgradeGovernor.Contract.Timelock(&_UpgradeGovernor.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_UpgradeGovernor *UpgradeGovernorCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UpgradeGovernor.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_UpgradeGovernor *UpgradeGovernorSession) Token() (common.Address, error) {
	return _UpgradeGovernor.Contract.Token(&_UpgradeGovernor.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_UpgradeGovernor *UpgradeGovernorCallerSession) Token() (common.Address, error) {
	return _UpgradeGovernor.Contract.Token(&_UpgradeGovernor.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_UpgradeGovernor *UpgradeGovernorCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UpgradeGovernor.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_UpgradeGovernor *UpgradeGovernorSession) Version() (string, error) {
	return _UpgradeGovernor.Contract.Version(&_UpgradeGovernor.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_UpgradeGovernor *UpgradeGovernorCallerSession) Version() (string, error) {
	return _UpgradeGovernor.Contract.Version(&_UpgradeGovernor.CallOpts)
}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorCaller) VotingDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeGovernor.contract.Call(opts, &out, "votingDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorSession) VotingDelay() (*big.Int, error) {
	return _UpgradeGovernor.Contract.VotingDelay(&_UpgradeGovernor.CallOpts)
}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorCallerSession) VotingDelay() (*big.Int, error) {
	return _UpgradeGovernor.Contract.VotingDelay(&_UpgradeGovernor.CallOpts)
}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorCaller) VotingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeGovernor.contract.Call(opts, &out, "votingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorSession) VotingPeriod() (*big.Int, error) {
	return _UpgradeGovernor.Contract.VotingPeriod(&_UpgradeGovernor.CallOpts)
}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorCallerSession) VotingPeriod() (*big.Int, error) {
	return _UpgradeGovernor.Contract.VotingPeriod(&_UpgradeGovernor.CallOpts)
}

// Cancel is a paid mutator transaction binding the contract method 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorTransactor) Cancel(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _UpgradeGovernor.contract.Transact(opts, "cancel", targets, values, calldatas, descriptionHash)
}

// Cancel is a paid mutator transaction binding the contract method 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorSession) Cancel(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.Cancel(&_UpgradeGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Cancel is a paid mutator transaction binding the contract method 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorTransactorSession) Cancel(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.Cancel(&_UpgradeGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorTransactor) CastVote(opts *bind.TransactOpts, proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _UpgradeGovernor.contract.Transact(opts, "castVote", proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorSession) CastVote(proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.CastVote(&_UpgradeGovernor.TransactOpts, proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorTransactorSession) CastVote(proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.CastVote(&_UpgradeGovernor.TransactOpts, proposalId, support)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x3bccf4fd.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, uint8 v, bytes32 r, bytes32 s) returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorTransactor) CastVoteBySig(opts *bind.TransactOpts, proposalId *big.Int, support uint8, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UpgradeGovernor.contract.Transact(opts, "castVoteBySig", proposalId, support, v, r, s)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x3bccf4fd.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, uint8 v, bytes32 r, bytes32 s) returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorSession) CastVoteBySig(proposalId *big.Int, support uint8, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.CastVoteBySig(&_UpgradeGovernor.TransactOpts, proposalId, support, v, r, s)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x3bccf4fd.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, uint8 v, bytes32 r, bytes32 s) returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorTransactorSession) CastVoteBySig(proposalId *big.Int, support uint8, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.CastVoteBySig(&_UpgradeGovernor.TransactOpts, proposalId, support, v, r, s)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorTransactor) CastVoteWithReason(opts *bind.TransactOpts, proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _UpgradeGovernor.contract.Transact(opts, "castVoteWithReason", proposalId, support, reason)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorSession) CastVoteWithReason(proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.CastVoteWithReason(&_UpgradeGovernor.TransactOpts, proposalId, support, reason)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorTransactorSession) CastVoteWithReason(proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.CastVoteWithReason(&_UpgradeGovernor.TransactOpts, proposalId, support, reason)
}

// CastVoteWithReasonAndParams is a paid mutator transaction binding the contract method 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorTransactor) CastVoteWithReasonAndParams(opts *bind.TransactOpts, proposalId *big.Int, support uint8, reason string, params []byte) (*types.Transaction, error) {
	return _UpgradeGovernor.contract.Transact(opts, "castVoteWithReasonAndParams", proposalId, support, reason, params)
}

// CastVoteWithReasonAndParams is a paid mutator transaction binding the contract method 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorSession) CastVoteWithReasonAndParams(proposalId *big.Int, support uint8, reason string, params []byte) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.CastVoteWithReasonAndParams(&_UpgradeGovernor.TransactOpts, proposalId, support, reason, params)
}

// CastVoteWithReasonAndParams is a paid mutator transaction binding the contract method 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorTransactorSession) CastVoteWithReasonAndParams(proposalId *big.Int, support uint8, reason string, params []byte) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.CastVoteWithReasonAndParams(&_UpgradeGovernor.TransactOpts, proposalId, support, reason, params)
}

// CastVoteWithReasonAndParamsBySig is a paid mutator transaction binding the contract method 0x03420181.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, string reason, bytes params, uint8 v, bytes32 r, bytes32 s) returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorTransactor) CastVoteWithReasonAndParamsBySig(opts *bind.TransactOpts, proposalId *big.Int, support uint8, reason string, params []byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UpgradeGovernor.contract.Transact(opts, "castVoteWithReasonAndParamsBySig", proposalId, support, reason, params, v, r, s)
}

// CastVoteWithReasonAndParamsBySig is a paid mutator transaction binding the contract method 0x03420181.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, string reason, bytes params, uint8 v, bytes32 r, bytes32 s) returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorSession) CastVoteWithReasonAndParamsBySig(proposalId *big.Int, support uint8, reason string, params []byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.CastVoteWithReasonAndParamsBySig(&_UpgradeGovernor.TransactOpts, proposalId, support, reason, params, v, r, s)
}

// CastVoteWithReasonAndParamsBySig is a paid mutator transaction binding the contract method 0x03420181.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, string reason, bytes params, uint8 v, bytes32 r, bytes32 s) returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorTransactorSession) CastVoteWithReasonAndParamsBySig(proposalId *big.Int, support uint8, reason string, params []byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.CastVoteWithReasonAndParamsBySig(&_UpgradeGovernor.TransactOpts, proposalId, support, reason, params, v, r, s)
}

// Execute is a paid mutator transaction binding the contract method 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorTransactor) Execute(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _UpgradeGovernor.contract.Transact(opts, "execute", targets, values, calldatas, descriptionHash)
}

// Execute is a paid mutator transaction binding the contract method 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorSession) Execute(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.Execute(&_UpgradeGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Execute is a paid mutator transaction binding the contract method 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorTransactorSession) Execute(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.Execute(&_UpgradeGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Initialize is a paid mutator transaction binding the contract method 0x86489ba9.
//
// Solidity: function initialize(address _token, address _timelock, uint256 _initialVotingDelay, uint256 _initialVotingPeriod, uint256 _initialProposalThreshold, uint256 _votesQuorumFraction) returns()
func (_UpgradeGovernor *UpgradeGovernorTransactor) Initialize(opts *bind.TransactOpts, _token common.Address, _timelock common.Address, _initialVotingDelay *big.Int, _initialVotingPeriod *big.Int, _initialProposalThreshold *big.Int, _votesQuorumFraction *big.Int) (*types.Transaction, error) {
	return _UpgradeGovernor.contract.Transact(opts, "initialize", _token, _timelock, _initialVotingDelay, _initialVotingPeriod, _initialProposalThreshold, _votesQuorumFraction)
}

// Initialize is a paid mutator transaction binding the contract method 0x86489ba9.
//
// Solidity: function initialize(address _token, address _timelock, uint256 _initialVotingDelay, uint256 _initialVotingPeriod, uint256 _initialProposalThreshold, uint256 _votesQuorumFraction) returns()
func (_UpgradeGovernor *UpgradeGovernorSession) Initialize(_token common.Address, _timelock common.Address, _initialVotingDelay *big.Int, _initialVotingPeriod *big.Int, _initialProposalThreshold *big.Int, _votesQuorumFraction *big.Int) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.Initialize(&_UpgradeGovernor.TransactOpts, _token, _timelock, _initialVotingDelay, _initialVotingPeriod, _initialProposalThreshold, _votesQuorumFraction)
}

// Initialize is a paid mutator transaction binding the contract method 0x86489ba9.
//
// Solidity: function initialize(address _token, address _timelock, uint256 _initialVotingDelay, uint256 _initialVotingPeriod, uint256 _initialProposalThreshold, uint256 _votesQuorumFraction) returns()
func (_UpgradeGovernor *UpgradeGovernorTransactorSession) Initialize(_token common.Address, _timelock common.Address, _initialVotingDelay *big.Int, _initialVotingPeriod *big.Int, _initialProposalThreshold *big.Int, _votesQuorumFraction *big.Int) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.Initialize(&_UpgradeGovernor.TransactOpts, _token, _timelock, _initialVotingDelay, _initialVotingPeriod, _initialProposalThreshold, _votesQuorumFraction)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_UpgradeGovernor *UpgradeGovernorTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _UpgradeGovernor.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_UpgradeGovernor *UpgradeGovernorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.OnERC1155BatchReceived(&_UpgradeGovernor.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_UpgradeGovernor *UpgradeGovernorTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.OnERC1155BatchReceived(&_UpgradeGovernor.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_UpgradeGovernor *UpgradeGovernorTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _UpgradeGovernor.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_UpgradeGovernor *UpgradeGovernorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.OnERC1155Received(&_UpgradeGovernor.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_UpgradeGovernor *UpgradeGovernorTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.OnERC1155Received(&_UpgradeGovernor.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_UpgradeGovernor *UpgradeGovernorTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _UpgradeGovernor.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_UpgradeGovernor *UpgradeGovernorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.OnERC721Received(&_UpgradeGovernor.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_UpgradeGovernor *UpgradeGovernorTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.OnERC721Received(&_UpgradeGovernor.TransactOpts, arg0, arg1, arg2, arg3)
}

// Propose is a paid mutator transaction binding the contract method 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorTransactor) Propose(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _UpgradeGovernor.contract.Transact(opts, "propose", targets, values, calldatas, description)
}

// Propose is a paid mutator transaction binding the contract method 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorSession) Propose(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.Propose(&_UpgradeGovernor.TransactOpts, targets, values, calldatas, description)
}

// Propose is a paid mutator transaction binding the contract method 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorTransactorSession) Propose(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.Propose(&_UpgradeGovernor.TransactOpts, targets, values, calldatas, description)
}

// Queue is a paid mutator transaction binding the contract method 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorTransactor) Queue(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _UpgradeGovernor.contract.Transact(opts, "queue", targets, values, calldatas, descriptionHash)
}

// Queue is a paid mutator transaction binding the contract method 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorSession) Queue(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.Queue(&_UpgradeGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Queue is a paid mutator transaction binding the contract method 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_UpgradeGovernor *UpgradeGovernorTransactorSession) Queue(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.Queue(&_UpgradeGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Relay is a paid mutator transaction binding the contract method 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (_UpgradeGovernor *UpgradeGovernorTransactor) Relay(opts *bind.TransactOpts, target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _UpgradeGovernor.contract.Transact(opts, "relay", target, value, data)
}

// Relay is a paid mutator transaction binding the contract method 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (_UpgradeGovernor *UpgradeGovernorSession) Relay(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.Relay(&_UpgradeGovernor.TransactOpts, target, value, data)
}

// Relay is a paid mutator transaction binding the contract method 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (_UpgradeGovernor *UpgradeGovernorTransactorSession) Relay(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.Relay(&_UpgradeGovernor.TransactOpts, target, value, data)
}

// SetProposalThreshold is a paid mutator transaction binding the contract method 0xece40cc1.
//
// Solidity: function setProposalThreshold(uint256 newProposalThreshold) returns()
func (_UpgradeGovernor *UpgradeGovernorTransactor) SetProposalThreshold(opts *bind.TransactOpts, newProposalThreshold *big.Int) (*types.Transaction, error) {
	return _UpgradeGovernor.contract.Transact(opts, "setProposalThreshold", newProposalThreshold)
}

// SetProposalThreshold is a paid mutator transaction binding the contract method 0xece40cc1.
//
// Solidity: function setProposalThreshold(uint256 newProposalThreshold) returns()
func (_UpgradeGovernor *UpgradeGovernorSession) SetProposalThreshold(newProposalThreshold *big.Int) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.SetProposalThreshold(&_UpgradeGovernor.TransactOpts, newProposalThreshold)
}

// SetProposalThreshold is a paid mutator transaction binding the contract method 0xece40cc1.
//
// Solidity: function setProposalThreshold(uint256 newProposalThreshold) returns()
func (_UpgradeGovernor *UpgradeGovernorTransactorSession) SetProposalThreshold(newProposalThreshold *big.Int) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.SetProposalThreshold(&_UpgradeGovernor.TransactOpts, newProposalThreshold)
}

// SetVotingDelay is a paid mutator transaction binding the contract method 0x70b0f660.
//
// Solidity: function setVotingDelay(uint256 newVotingDelay) returns()
func (_UpgradeGovernor *UpgradeGovernorTransactor) SetVotingDelay(opts *bind.TransactOpts, newVotingDelay *big.Int) (*types.Transaction, error) {
	return _UpgradeGovernor.contract.Transact(opts, "setVotingDelay", newVotingDelay)
}

// SetVotingDelay is a paid mutator transaction binding the contract method 0x70b0f660.
//
// Solidity: function setVotingDelay(uint256 newVotingDelay) returns()
func (_UpgradeGovernor *UpgradeGovernorSession) SetVotingDelay(newVotingDelay *big.Int) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.SetVotingDelay(&_UpgradeGovernor.TransactOpts, newVotingDelay)
}

// SetVotingDelay is a paid mutator transaction binding the contract method 0x70b0f660.
//
// Solidity: function setVotingDelay(uint256 newVotingDelay) returns()
func (_UpgradeGovernor *UpgradeGovernorTransactorSession) SetVotingDelay(newVotingDelay *big.Int) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.SetVotingDelay(&_UpgradeGovernor.TransactOpts, newVotingDelay)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xea0217cf.
//
// Solidity: function setVotingPeriod(uint256 newVotingPeriod) returns()
func (_UpgradeGovernor *UpgradeGovernorTransactor) SetVotingPeriod(opts *bind.TransactOpts, newVotingPeriod *big.Int) (*types.Transaction, error) {
	return _UpgradeGovernor.contract.Transact(opts, "setVotingPeriod", newVotingPeriod)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xea0217cf.
//
// Solidity: function setVotingPeriod(uint256 newVotingPeriod) returns()
func (_UpgradeGovernor *UpgradeGovernorSession) SetVotingPeriod(newVotingPeriod *big.Int) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.SetVotingPeriod(&_UpgradeGovernor.TransactOpts, newVotingPeriod)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xea0217cf.
//
// Solidity: function setVotingPeriod(uint256 newVotingPeriod) returns()
func (_UpgradeGovernor *UpgradeGovernorTransactorSession) SetVotingPeriod(newVotingPeriod *big.Int) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.SetVotingPeriod(&_UpgradeGovernor.TransactOpts, newVotingPeriod)
}

// UpdateQuorumNumerator is a paid mutator transaction binding the contract method 0x06f3f9e6.
//
// Solidity: function updateQuorumNumerator(uint256 newQuorumNumerator) returns()
func (_UpgradeGovernor *UpgradeGovernorTransactor) UpdateQuorumNumerator(opts *bind.TransactOpts, newQuorumNumerator *big.Int) (*types.Transaction, error) {
	return _UpgradeGovernor.contract.Transact(opts, "updateQuorumNumerator", newQuorumNumerator)
}

// UpdateQuorumNumerator is a paid mutator transaction binding the contract method 0x06f3f9e6.
//
// Solidity: function updateQuorumNumerator(uint256 newQuorumNumerator) returns()
func (_UpgradeGovernor *UpgradeGovernorSession) UpdateQuorumNumerator(newQuorumNumerator *big.Int) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.UpdateQuorumNumerator(&_UpgradeGovernor.TransactOpts, newQuorumNumerator)
}

// UpdateQuorumNumerator is a paid mutator transaction binding the contract method 0x06f3f9e6.
//
// Solidity: function updateQuorumNumerator(uint256 newQuorumNumerator) returns()
func (_UpgradeGovernor *UpgradeGovernorTransactorSession) UpdateQuorumNumerator(newQuorumNumerator *big.Int) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.UpdateQuorumNumerator(&_UpgradeGovernor.TransactOpts, newQuorumNumerator)
}

// UpdateTimelock is a paid mutator transaction binding the contract method 0xa890c910.
//
// Solidity: function updateTimelock(address newTimelock) returns()
func (_UpgradeGovernor *UpgradeGovernorTransactor) UpdateTimelock(opts *bind.TransactOpts, newTimelock common.Address) (*types.Transaction, error) {
	return _UpgradeGovernor.contract.Transact(opts, "updateTimelock", newTimelock)
}

// UpdateTimelock is a paid mutator transaction binding the contract method 0xa890c910.
//
// Solidity: function updateTimelock(address newTimelock) returns()
func (_UpgradeGovernor *UpgradeGovernorSession) UpdateTimelock(newTimelock common.Address) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.UpdateTimelock(&_UpgradeGovernor.TransactOpts, newTimelock)
}

// UpdateTimelock is a paid mutator transaction binding the contract method 0xa890c910.
//
// Solidity: function updateTimelock(address newTimelock) returns()
func (_UpgradeGovernor *UpgradeGovernorTransactorSession) UpdateTimelock(newTimelock common.Address) (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.UpdateTimelock(&_UpgradeGovernor.TransactOpts, newTimelock)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UpgradeGovernor *UpgradeGovernorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeGovernor.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UpgradeGovernor *UpgradeGovernorSession) Receive() (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.Receive(&_UpgradeGovernor.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UpgradeGovernor *UpgradeGovernorTransactorSession) Receive() (*types.Transaction, error) {
	return _UpgradeGovernor.Contract.Receive(&_UpgradeGovernor.TransactOpts)
}

// UpgradeGovernorEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the UpgradeGovernor contract.
type UpgradeGovernorEIP712DomainChangedIterator struct {
	Event *UpgradeGovernorEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *UpgradeGovernorEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeGovernorEIP712DomainChanged)
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
		it.Event = new(UpgradeGovernorEIP712DomainChanged)
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
func (it *UpgradeGovernorEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeGovernorEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeGovernorEIP712DomainChanged represents a EIP712DomainChanged event raised by the UpgradeGovernor contract.
type UpgradeGovernorEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_UpgradeGovernor *UpgradeGovernorFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*UpgradeGovernorEIP712DomainChangedIterator, error) {

	logs, sub, err := _UpgradeGovernor.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &UpgradeGovernorEIP712DomainChangedIterator{contract: _UpgradeGovernor.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_UpgradeGovernor *UpgradeGovernorFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *UpgradeGovernorEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _UpgradeGovernor.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeGovernorEIP712DomainChanged)
				if err := _UpgradeGovernor.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_UpgradeGovernor *UpgradeGovernorFilterer) ParseEIP712DomainChanged(log types.Log) (*UpgradeGovernorEIP712DomainChanged, error) {
	event := new(UpgradeGovernorEIP712DomainChanged)
	if err := _UpgradeGovernor.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeGovernorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the UpgradeGovernor contract.
type UpgradeGovernorInitializedIterator struct {
	Event *UpgradeGovernorInitialized // Event containing the contract specifics and raw log

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
func (it *UpgradeGovernorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeGovernorInitialized)
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
		it.Event = new(UpgradeGovernorInitialized)
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
func (it *UpgradeGovernorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeGovernorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeGovernorInitialized represents a Initialized event raised by the UpgradeGovernor contract.
type UpgradeGovernorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UpgradeGovernor *UpgradeGovernorFilterer) FilterInitialized(opts *bind.FilterOpts) (*UpgradeGovernorInitializedIterator, error) {

	logs, sub, err := _UpgradeGovernor.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &UpgradeGovernorInitializedIterator{contract: _UpgradeGovernor.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UpgradeGovernor *UpgradeGovernorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *UpgradeGovernorInitialized) (event.Subscription, error) {

	logs, sub, err := _UpgradeGovernor.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeGovernorInitialized)
				if err := _UpgradeGovernor.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_UpgradeGovernor *UpgradeGovernorFilterer) ParseInitialized(log types.Log) (*UpgradeGovernorInitialized, error) {
	event := new(UpgradeGovernorInitialized)
	if err := _UpgradeGovernor.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeGovernorProposalCanceledIterator is returned from FilterProposalCanceled and is used to iterate over the raw logs and unpacked data for ProposalCanceled events raised by the UpgradeGovernor contract.
type UpgradeGovernorProposalCanceledIterator struct {
	Event *UpgradeGovernorProposalCanceled // Event containing the contract specifics and raw log

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
func (it *UpgradeGovernorProposalCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeGovernorProposalCanceled)
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
		it.Event = new(UpgradeGovernorProposalCanceled)
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
func (it *UpgradeGovernorProposalCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeGovernorProposalCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeGovernorProposalCanceled represents a ProposalCanceled event raised by the UpgradeGovernor contract.
type UpgradeGovernorProposalCanceled struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalCanceled is a free log retrieval operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_UpgradeGovernor *UpgradeGovernorFilterer) FilterProposalCanceled(opts *bind.FilterOpts) (*UpgradeGovernorProposalCanceledIterator, error) {

	logs, sub, err := _UpgradeGovernor.contract.FilterLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return &UpgradeGovernorProposalCanceledIterator{contract: _UpgradeGovernor.contract, event: "ProposalCanceled", logs: logs, sub: sub}, nil
}

// WatchProposalCanceled is a free log subscription operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_UpgradeGovernor *UpgradeGovernorFilterer) WatchProposalCanceled(opts *bind.WatchOpts, sink chan<- *UpgradeGovernorProposalCanceled) (event.Subscription, error) {

	logs, sub, err := _UpgradeGovernor.contract.WatchLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeGovernorProposalCanceled)
				if err := _UpgradeGovernor.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
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

// ParseProposalCanceled is a log parse operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_UpgradeGovernor *UpgradeGovernorFilterer) ParseProposalCanceled(log types.Log) (*UpgradeGovernorProposalCanceled, error) {
	event := new(UpgradeGovernorProposalCanceled)
	if err := _UpgradeGovernor.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeGovernorProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the UpgradeGovernor contract.
type UpgradeGovernorProposalCreatedIterator struct {
	Event *UpgradeGovernorProposalCreated // Event containing the contract specifics and raw log

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
func (it *UpgradeGovernorProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeGovernorProposalCreated)
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
		it.Event = new(UpgradeGovernorProposalCreated)
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
func (it *UpgradeGovernorProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeGovernorProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeGovernorProposalCreated represents a ProposalCreated event raised by the UpgradeGovernor contract.
type UpgradeGovernorProposalCreated struct {
	ProposalId  *big.Int
	Proposer    common.Address
	Targets     []common.Address
	Values      []*big.Int
	Signatures  []string
	Calldatas   [][]byte
	VoteStart   *big.Int
	VoteEnd     *big.Int
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 voteStart, uint256 voteEnd, string description)
func (_UpgradeGovernor *UpgradeGovernorFilterer) FilterProposalCreated(opts *bind.FilterOpts) (*UpgradeGovernorProposalCreatedIterator, error) {

	logs, sub, err := _UpgradeGovernor.contract.FilterLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return &UpgradeGovernorProposalCreatedIterator{contract: _UpgradeGovernor.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 voteStart, uint256 voteEnd, string description)
func (_UpgradeGovernor *UpgradeGovernorFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *UpgradeGovernorProposalCreated) (event.Subscription, error) {

	logs, sub, err := _UpgradeGovernor.contract.WatchLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeGovernorProposalCreated)
				if err := _UpgradeGovernor.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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

// ParseProposalCreated is a log parse operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 voteStart, uint256 voteEnd, string description)
func (_UpgradeGovernor *UpgradeGovernorFilterer) ParseProposalCreated(log types.Log) (*UpgradeGovernorProposalCreated, error) {
	event := new(UpgradeGovernorProposalCreated)
	if err := _UpgradeGovernor.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeGovernorProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the UpgradeGovernor contract.
type UpgradeGovernorProposalExecutedIterator struct {
	Event *UpgradeGovernorProposalExecuted // Event containing the contract specifics and raw log

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
func (it *UpgradeGovernorProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeGovernorProposalExecuted)
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
		it.Event = new(UpgradeGovernorProposalExecuted)
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
func (it *UpgradeGovernorProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeGovernorProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeGovernorProposalExecuted represents a ProposalExecuted event raised by the UpgradeGovernor contract.
type UpgradeGovernorProposalExecuted struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_UpgradeGovernor *UpgradeGovernorFilterer) FilterProposalExecuted(opts *bind.FilterOpts) (*UpgradeGovernorProposalExecutedIterator, error) {

	logs, sub, err := _UpgradeGovernor.contract.FilterLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return &UpgradeGovernorProposalExecutedIterator{contract: _UpgradeGovernor.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_UpgradeGovernor *UpgradeGovernorFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *UpgradeGovernorProposalExecuted) (event.Subscription, error) {

	logs, sub, err := _UpgradeGovernor.contract.WatchLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeGovernorProposalExecuted)
				if err := _UpgradeGovernor.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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

// ParseProposalExecuted is a log parse operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_UpgradeGovernor *UpgradeGovernorFilterer) ParseProposalExecuted(log types.Log) (*UpgradeGovernorProposalExecuted, error) {
	event := new(UpgradeGovernorProposalExecuted)
	if err := _UpgradeGovernor.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeGovernorProposalQueuedIterator is returned from FilterProposalQueued and is used to iterate over the raw logs and unpacked data for ProposalQueued events raised by the UpgradeGovernor contract.
type UpgradeGovernorProposalQueuedIterator struct {
	Event *UpgradeGovernorProposalQueued // Event containing the contract specifics and raw log

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
func (it *UpgradeGovernorProposalQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeGovernorProposalQueued)
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
		it.Event = new(UpgradeGovernorProposalQueued)
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
func (it *UpgradeGovernorProposalQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeGovernorProposalQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeGovernorProposalQueued represents a ProposalQueued event raised by the UpgradeGovernor contract.
type UpgradeGovernorProposalQueued struct {
	ProposalId *big.Int
	Eta        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalQueued is a free log retrieval operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 eta)
func (_UpgradeGovernor *UpgradeGovernorFilterer) FilterProposalQueued(opts *bind.FilterOpts) (*UpgradeGovernorProposalQueuedIterator, error) {

	logs, sub, err := _UpgradeGovernor.contract.FilterLogs(opts, "ProposalQueued")
	if err != nil {
		return nil, err
	}
	return &UpgradeGovernorProposalQueuedIterator{contract: _UpgradeGovernor.contract, event: "ProposalQueued", logs: logs, sub: sub}, nil
}

// WatchProposalQueued is a free log subscription operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 eta)
func (_UpgradeGovernor *UpgradeGovernorFilterer) WatchProposalQueued(opts *bind.WatchOpts, sink chan<- *UpgradeGovernorProposalQueued) (event.Subscription, error) {

	logs, sub, err := _UpgradeGovernor.contract.WatchLogs(opts, "ProposalQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeGovernorProposalQueued)
				if err := _UpgradeGovernor.contract.UnpackLog(event, "ProposalQueued", log); err != nil {
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

// ParseProposalQueued is a log parse operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 eta)
func (_UpgradeGovernor *UpgradeGovernorFilterer) ParseProposalQueued(log types.Log) (*UpgradeGovernorProposalQueued, error) {
	event := new(UpgradeGovernorProposalQueued)
	if err := _UpgradeGovernor.contract.UnpackLog(event, "ProposalQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeGovernorProposalThresholdSetIterator is returned from FilterProposalThresholdSet and is used to iterate over the raw logs and unpacked data for ProposalThresholdSet events raised by the UpgradeGovernor contract.
type UpgradeGovernorProposalThresholdSetIterator struct {
	Event *UpgradeGovernorProposalThresholdSet // Event containing the contract specifics and raw log

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
func (it *UpgradeGovernorProposalThresholdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeGovernorProposalThresholdSet)
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
		it.Event = new(UpgradeGovernorProposalThresholdSet)
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
func (it *UpgradeGovernorProposalThresholdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeGovernorProposalThresholdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeGovernorProposalThresholdSet represents a ProposalThresholdSet event raised by the UpgradeGovernor contract.
type UpgradeGovernorProposalThresholdSet struct {
	OldProposalThreshold *big.Int
	NewProposalThreshold *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterProposalThresholdSet is a free log retrieval operation binding the contract event 0xccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc05461.
//
// Solidity: event ProposalThresholdSet(uint256 oldProposalThreshold, uint256 newProposalThreshold)
func (_UpgradeGovernor *UpgradeGovernorFilterer) FilterProposalThresholdSet(opts *bind.FilterOpts) (*UpgradeGovernorProposalThresholdSetIterator, error) {

	logs, sub, err := _UpgradeGovernor.contract.FilterLogs(opts, "ProposalThresholdSet")
	if err != nil {
		return nil, err
	}
	return &UpgradeGovernorProposalThresholdSetIterator{contract: _UpgradeGovernor.contract, event: "ProposalThresholdSet", logs: logs, sub: sub}, nil
}

// WatchProposalThresholdSet is a free log subscription operation binding the contract event 0xccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc05461.
//
// Solidity: event ProposalThresholdSet(uint256 oldProposalThreshold, uint256 newProposalThreshold)
func (_UpgradeGovernor *UpgradeGovernorFilterer) WatchProposalThresholdSet(opts *bind.WatchOpts, sink chan<- *UpgradeGovernorProposalThresholdSet) (event.Subscription, error) {

	logs, sub, err := _UpgradeGovernor.contract.WatchLogs(opts, "ProposalThresholdSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeGovernorProposalThresholdSet)
				if err := _UpgradeGovernor.contract.UnpackLog(event, "ProposalThresholdSet", log); err != nil {
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

// ParseProposalThresholdSet is a log parse operation binding the contract event 0xccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc05461.
//
// Solidity: event ProposalThresholdSet(uint256 oldProposalThreshold, uint256 newProposalThreshold)
func (_UpgradeGovernor *UpgradeGovernorFilterer) ParseProposalThresholdSet(log types.Log) (*UpgradeGovernorProposalThresholdSet, error) {
	event := new(UpgradeGovernorProposalThresholdSet)
	if err := _UpgradeGovernor.contract.UnpackLog(event, "ProposalThresholdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeGovernorQuorumNumeratorUpdatedIterator is returned from FilterQuorumNumeratorUpdated and is used to iterate over the raw logs and unpacked data for QuorumNumeratorUpdated events raised by the UpgradeGovernor contract.
type UpgradeGovernorQuorumNumeratorUpdatedIterator struct {
	Event *UpgradeGovernorQuorumNumeratorUpdated // Event containing the contract specifics and raw log

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
func (it *UpgradeGovernorQuorumNumeratorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeGovernorQuorumNumeratorUpdated)
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
		it.Event = new(UpgradeGovernorQuorumNumeratorUpdated)
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
func (it *UpgradeGovernorQuorumNumeratorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeGovernorQuorumNumeratorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeGovernorQuorumNumeratorUpdated represents a QuorumNumeratorUpdated event raised by the UpgradeGovernor contract.
type UpgradeGovernorQuorumNumeratorUpdated struct {
	OldQuorumNumerator *big.Int
	NewQuorumNumerator *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterQuorumNumeratorUpdated is a free log retrieval operation binding the contract event 0x0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997.
//
// Solidity: event QuorumNumeratorUpdated(uint256 oldQuorumNumerator, uint256 newQuorumNumerator)
func (_UpgradeGovernor *UpgradeGovernorFilterer) FilterQuorumNumeratorUpdated(opts *bind.FilterOpts) (*UpgradeGovernorQuorumNumeratorUpdatedIterator, error) {

	logs, sub, err := _UpgradeGovernor.contract.FilterLogs(opts, "QuorumNumeratorUpdated")
	if err != nil {
		return nil, err
	}
	return &UpgradeGovernorQuorumNumeratorUpdatedIterator{contract: _UpgradeGovernor.contract, event: "QuorumNumeratorUpdated", logs: logs, sub: sub}, nil
}

// WatchQuorumNumeratorUpdated is a free log subscription operation binding the contract event 0x0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997.
//
// Solidity: event QuorumNumeratorUpdated(uint256 oldQuorumNumerator, uint256 newQuorumNumerator)
func (_UpgradeGovernor *UpgradeGovernorFilterer) WatchQuorumNumeratorUpdated(opts *bind.WatchOpts, sink chan<- *UpgradeGovernorQuorumNumeratorUpdated) (event.Subscription, error) {

	logs, sub, err := _UpgradeGovernor.contract.WatchLogs(opts, "QuorumNumeratorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeGovernorQuorumNumeratorUpdated)
				if err := _UpgradeGovernor.contract.UnpackLog(event, "QuorumNumeratorUpdated", log); err != nil {
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

// ParseQuorumNumeratorUpdated is a log parse operation binding the contract event 0x0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997.
//
// Solidity: event QuorumNumeratorUpdated(uint256 oldQuorumNumerator, uint256 newQuorumNumerator)
func (_UpgradeGovernor *UpgradeGovernorFilterer) ParseQuorumNumeratorUpdated(log types.Log) (*UpgradeGovernorQuorumNumeratorUpdated, error) {
	event := new(UpgradeGovernorQuorumNumeratorUpdated)
	if err := _UpgradeGovernor.contract.UnpackLog(event, "QuorumNumeratorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeGovernorTimelockChangeIterator is returned from FilterTimelockChange and is used to iterate over the raw logs and unpacked data for TimelockChange events raised by the UpgradeGovernor contract.
type UpgradeGovernorTimelockChangeIterator struct {
	Event *UpgradeGovernorTimelockChange // Event containing the contract specifics and raw log

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
func (it *UpgradeGovernorTimelockChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeGovernorTimelockChange)
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
		it.Event = new(UpgradeGovernorTimelockChange)
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
func (it *UpgradeGovernorTimelockChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeGovernorTimelockChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeGovernorTimelockChange represents a TimelockChange event raised by the UpgradeGovernor contract.
type UpgradeGovernorTimelockChange struct {
	OldTimelock common.Address
	NewTimelock common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTimelockChange is a free log retrieval operation binding the contract event 0x08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401.
//
// Solidity: event TimelockChange(address oldTimelock, address newTimelock)
func (_UpgradeGovernor *UpgradeGovernorFilterer) FilterTimelockChange(opts *bind.FilterOpts) (*UpgradeGovernorTimelockChangeIterator, error) {

	logs, sub, err := _UpgradeGovernor.contract.FilterLogs(opts, "TimelockChange")
	if err != nil {
		return nil, err
	}
	return &UpgradeGovernorTimelockChangeIterator{contract: _UpgradeGovernor.contract, event: "TimelockChange", logs: logs, sub: sub}, nil
}

// WatchTimelockChange is a free log subscription operation binding the contract event 0x08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401.
//
// Solidity: event TimelockChange(address oldTimelock, address newTimelock)
func (_UpgradeGovernor *UpgradeGovernorFilterer) WatchTimelockChange(opts *bind.WatchOpts, sink chan<- *UpgradeGovernorTimelockChange) (event.Subscription, error) {

	logs, sub, err := _UpgradeGovernor.contract.WatchLogs(opts, "TimelockChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeGovernorTimelockChange)
				if err := _UpgradeGovernor.contract.UnpackLog(event, "TimelockChange", log); err != nil {
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

// ParseTimelockChange is a log parse operation binding the contract event 0x08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401.
//
// Solidity: event TimelockChange(address oldTimelock, address newTimelock)
func (_UpgradeGovernor *UpgradeGovernorFilterer) ParseTimelockChange(log types.Log) (*UpgradeGovernorTimelockChange, error) {
	event := new(UpgradeGovernorTimelockChange)
	if err := _UpgradeGovernor.contract.UnpackLog(event, "TimelockChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeGovernorVoteCastIterator is returned from FilterVoteCast and is used to iterate over the raw logs and unpacked data for VoteCast events raised by the UpgradeGovernor contract.
type UpgradeGovernorVoteCastIterator struct {
	Event *UpgradeGovernorVoteCast // Event containing the contract specifics and raw log

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
func (it *UpgradeGovernorVoteCastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeGovernorVoteCast)
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
		it.Event = new(UpgradeGovernorVoteCast)
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
func (it *UpgradeGovernorVoteCastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeGovernorVoteCastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeGovernorVoteCast represents a VoteCast event raised by the UpgradeGovernor contract.
type UpgradeGovernorVoteCast struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    uint8
	Weight     *big.Int
	Reason     string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteCast is a free log retrieval operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (_UpgradeGovernor *UpgradeGovernorFilterer) FilterVoteCast(opts *bind.FilterOpts, voter []common.Address) (*UpgradeGovernorVoteCastIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _UpgradeGovernor.contract.FilterLogs(opts, "VoteCast", voterRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeGovernorVoteCastIterator{contract: _UpgradeGovernor.contract, event: "VoteCast", logs: logs, sub: sub}, nil
}

// WatchVoteCast is a free log subscription operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (_UpgradeGovernor *UpgradeGovernorFilterer) WatchVoteCast(opts *bind.WatchOpts, sink chan<- *UpgradeGovernorVoteCast, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _UpgradeGovernor.contract.WatchLogs(opts, "VoteCast", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeGovernorVoteCast)
				if err := _UpgradeGovernor.contract.UnpackLog(event, "VoteCast", log); err != nil {
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

// ParseVoteCast is a log parse operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (_UpgradeGovernor *UpgradeGovernorFilterer) ParseVoteCast(log types.Log) (*UpgradeGovernorVoteCast, error) {
	event := new(UpgradeGovernorVoteCast)
	if err := _UpgradeGovernor.contract.UnpackLog(event, "VoteCast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeGovernorVoteCastWithParamsIterator is returned from FilterVoteCastWithParams and is used to iterate over the raw logs and unpacked data for VoteCastWithParams events raised by the UpgradeGovernor contract.
type UpgradeGovernorVoteCastWithParamsIterator struct {
	Event *UpgradeGovernorVoteCastWithParams // Event containing the contract specifics and raw log

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
func (it *UpgradeGovernorVoteCastWithParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeGovernorVoteCastWithParams)
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
		it.Event = new(UpgradeGovernorVoteCastWithParams)
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
func (it *UpgradeGovernorVoteCastWithParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeGovernorVoteCastWithParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeGovernorVoteCastWithParams represents a VoteCastWithParams event raised by the UpgradeGovernor contract.
type UpgradeGovernorVoteCastWithParams struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    uint8
	Weight     *big.Int
	Reason     string
	Params     []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteCastWithParams is a free log retrieval operation binding the contract event 0xe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (_UpgradeGovernor *UpgradeGovernorFilterer) FilterVoteCastWithParams(opts *bind.FilterOpts, voter []common.Address) (*UpgradeGovernorVoteCastWithParamsIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _UpgradeGovernor.contract.FilterLogs(opts, "VoteCastWithParams", voterRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeGovernorVoteCastWithParamsIterator{contract: _UpgradeGovernor.contract, event: "VoteCastWithParams", logs: logs, sub: sub}, nil
}

// WatchVoteCastWithParams is a free log subscription operation binding the contract event 0xe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (_UpgradeGovernor *UpgradeGovernorFilterer) WatchVoteCastWithParams(opts *bind.WatchOpts, sink chan<- *UpgradeGovernorVoteCastWithParams, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _UpgradeGovernor.contract.WatchLogs(opts, "VoteCastWithParams", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeGovernorVoteCastWithParams)
				if err := _UpgradeGovernor.contract.UnpackLog(event, "VoteCastWithParams", log); err != nil {
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

// ParseVoteCastWithParams is a log parse operation binding the contract event 0xe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (_UpgradeGovernor *UpgradeGovernorFilterer) ParseVoteCastWithParams(log types.Log) (*UpgradeGovernorVoteCastWithParams, error) {
	event := new(UpgradeGovernorVoteCastWithParams)
	if err := _UpgradeGovernor.contract.UnpackLog(event, "VoteCastWithParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeGovernorVotingDelaySetIterator is returned from FilterVotingDelaySet and is used to iterate over the raw logs and unpacked data for VotingDelaySet events raised by the UpgradeGovernor contract.
type UpgradeGovernorVotingDelaySetIterator struct {
	Event *UpgradeGovernorVotingDelaySet // Event containing the contract specifics and raw log

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
func (it *UpgradeGovernorVotingDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeGovernorVotingDelaySet)
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
		it.Event = new(UpgradeGovernorVotingDelaySet)
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
func (it *UpgradeGovernorVotingDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeGovernorVotingDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeGovernorVotingDelaySet represents a VotingDelaySet event raised by the UpgradeGovernor contract.
type UpgradeGovernorVotingDelaySet struct {
	OldVotingDelay *big.Int
	NewVotingDelay *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterVotingDelaySet is a free log retrieval operation binding the contract event 0xc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93.
//
// Solidity: event VotingDelaySet(uint256 oldVotingDelay, uint256 newVotingDelay)
func (_UpgradeGovernor *UpgradeGovernorFilterer) FilterVotingDelaySet(opts *bind.FilterOpts) (*UpgradeGovernorVotingDelaySetIterator, error) {

	logs, sub, err := _UpgradeGovernor.contract.FilterLogs(opts, "VotingDelaySet")
	if err != nil {
		return nil, err
	}
	return &UpgradeGovernorVotingDelaySetIterator{contract: _UpgradeGovernor.contract, event: "VotingDelaySet", logs: logs, sub: sub}, nil
}

// WatchVotingDelaySet is a free log subscription operation binding the contract event 0xc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93.
//
// Solidity: event VotingDelaySet(uint256 oldVotingDelay, uint256 newVotingDelay)
func (_UpgradeGovernor *UpgradeGovernorFilterer) WatchVotingDelaySet(opts *bind.WatchOpts, sink chan<- *UpgradeGovernorVotingDelaySet) (event.Subscription, error) {

	logs, sub, err := _UpgradeGovernor.contract.WatchLogs(opts, "VotingDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeGovernorVotingDelaySet)
				if err := _UpgradeGovernor.contract.UnpackLog(event, "VotingDelaySet", log); err != nil {
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

// ParseVotingDelaySet is a log parse operation binding the contract event 0xc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93.
//
// Solidity: event VotingDelaySet(uint256 oldVotingDelay, uint256 newVotingDelay)
func (_UpgradeGovernor *UpgradeGovernorFilterer) ParseVotingDelaySet(log types.Log) (*UpgradeGovernorVotingDelaySet, error) {
	event := new(UpgradeGovernorVotingDelaySet)
	if err := _UpgradeGovernor.contract.UnpackLog(event, "VotingDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeGovernorVotingPeriodSetIterator is returned from FilterVotingPeriodSet and is used to iterate over the raw logs and unpacked data for VotingPeriodSet events raised by the UpgradeGovernor contract.
type UpgradeGovernorVotingPeriodSetIterator struct {
	Event *UpgradeGovernorVotingPeriodSet // Event containing the contract specifics and raw log

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
func (it *UpgradeGovernorVotingPeriodSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeGovernorVotingPeriodSet)
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
		it.Event = new(UpgradeGovernorVotingPeriodSet)
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
func (it *UpgradeGovernorVotingPeriodSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeGovernorVotingPeriodSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeGovernorVotingPeriodSet represents a VotingPeriodSet event raised by the UpgradeGovernor contract.
type UpgradeGovernorVotingPeriodSet struct {
	OldVotingPeriod *big.Int
	NewVotingPeriod *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVotingPeriodSet is a free log retrieval operation binding the contract event 0x7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828.
//
// Solidity: event VotingPeriodSet(uint256 oldVotingPeriod, uint256 newVotingPeriod)
func (_UpgradeGovernor *UpgradeGovernorFilterer) FilterVotingPeriodSet(opts *bind.FilterOpts) (*UpgradeGovernorVotingPeriodSetIterator, error) {

	logs, sub, err := _UpgradeGovernor.contract.FilterLogs(opts, "VotingPeriodSet")
	if err != nil {
		return nil, err
	}
	return &UpgradeGovernorVotingPeriodSetIterator{contract: _UpgradeGovernor.contract, event: "VotingPeriodSet", logs: logs, sub: sub}, nil
}

// WatchVotingPeriodSet is a free log subscription operation binding the contract event 0x7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828.
//
// Solidity: event VotingPeriodSet(uint256 oldVotingPeriod, uint256 newVotingPeriod)
func (_UpgradeGovernor *UpgradeGovernorFilterer) WatchVotingPeriodSet(opts *bind.WatchOpts, sink chan<- *UpgradeGovernorVotingPeriodSet) (event.Subscription, error) {

	logs, sub, err := _UpgradeGovernor.contract.WatchLogs(opts, "VotingPeriodSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeGovernorVotingPeriodSet)
				if err := _UpgradeGovernor.contract.UnpackLog(event, "VotingPeriodSet", log); err != nil {
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

// ParseVotingPeriodSet is a log parse operation binding the contract event 0x7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828.
//
// Solidity: event VotingPeriodSet(uint256 oldVotingPeriod, uint256 newVotingPeriod)
func (_UpgradeGovernor *UpgradeGovernorFilterer) ParseVotingPeriodSet(log types.Log) (*UpgradeGovernorVotingPeriodSet, error) {
	event := new(UpgradeGovernorVotingPeriodSet)
	if err := _UpgradeGovernor.contract.UnpackLog(event, "VotingPeriodSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
