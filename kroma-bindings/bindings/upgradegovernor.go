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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Empty\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"}],\"name\":\"ProposalQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldProposalThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newProposalThreshold\",\"type\":\"uint256\"}],\"name\":\"ProposalThresholdSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldQuorumNumerator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newQuorumNumerator\",\"type\":\"uint256\"}],\"name\":\"QuorumNumeratorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldTimelock\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTimelock\",\"type\":\"address\"}],\"name\":\"TimelockChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"VoteCastWithParams\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVotingDelay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotingDelay\",\"type\":\"uint256\"}],\"name\":\"VotingDelaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVotingPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotingPeriod\",\"type\":\"uint256\"}],\"name\":\"VotingPeriodSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BALLOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLOCK_MODE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COUNTING_MODE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXTENDED_BALLOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"}],\"name\":\"castVote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"castVoteBySig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"castVoteWithReason\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"castVoteWithReasonAndParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"castVoteWithReasonAndParamsBySig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clock\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"}],\"name\":\"getVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"getVotesWithParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"hashProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_timelock\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_initialVotingDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_initialVotingPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_initialProposalThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_votesQuorumFraction\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalDeadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalEta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalProposer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"againstVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"abstainVotes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"queue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"quorum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumDenominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"}],\"name\":\"quorumNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newProposalThreshold\",\"type\":\"uint256\"}],\"name\":\"setProposalThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newVotingDelay\",\"type\":\"uint256\"}],\"name\":\"setVotingDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newVotingPeriod\",\"type\":\"uint256\"}],\"name\":\"setVotingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumIGovernorUpgradeable.ProposalState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timelock\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC5805Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newQuorumNumerator\",\"type\":\"uint256\"}],\"name\":\"updateQuorumNumerator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTimelockControllerUpgradeable\",\"name\":\"newTimelock\",\"type\":\"address\"}],\"name\":\"updateTimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b506001608052600060a081905260c0526200002b62000031565b620000f2565b600054610100900460ff16156200009e5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811614620000f0576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b60805160a05160c051615d6462000122600039600061261a015260006125f1015260006125c80152615d646000f3fe60806040526004361061030c5760003560e01c80637b3c71d31161019a578063c01f9e37116100e1578063ea0217cf1161008a578063f23a6e6111610064578063f23a6e6114610a6d578063f8ce560a14610ab2578063fc0c546a14610ad257600080fd5b8063ea0217cf14610a0d578063eb9019d414610a2d578063ece40cc114610a4d57600080fd5b8063d33219b4116100bb578063d33219b414610974578063dd4e2ba514610993578063deaaa7cc146109d957600080fd5b8063c01f9e3714610907578063c28bc2fa14610941578063c59057e41461095457600080fd5b80639a802a6d11610143578063ab58fb8e1161011d578063ab58fb8e1461088d578063b58131b0146108ad578063bc197c81146108c257600080fd5b80639a802a6d14610838578063a7713a7014610858578063a890c9101461086d57600080fd5b806386489ba91161017457806386489ba9146107d857806391ddadf4146107f857806397c3d3341461082457600080fd5b80637b3c71d3146107705780637d5e81e21461079057806384b0196e146107b057600080fd5b80633932abb11161025e578063544ffc9c116102075780635f398a14116101e15780635f398a141461071057806360c4247f1461073057806370b0f6601461075057600080fd5b8063544ffc9c1461068557806354fd4d50146106db57806356781388146106f057600080fd5b806343859632116102385780634385963214610605578063452115d6146106505780634bf5d7e91461067057600080fd5b80633932abb1146105a35780633bccf4fd146105b85780633e4f49e6146105d857600080fd5b8063143489d0116102c05780632656227d1161029a5780632656227d146105255780632d63f693146105385780632fe3e2611461056f57600080fd5b8063143489d014610436578063150b7a0214610490578063160cbed71461050557600080fd5b806303420181116102f157806303420181146103d457806306f3f9e6146103f457806306fdde031461041457600080fd5b806301ffc9a71461037c57806302a251a3146103b157600080fd5b36610377573061031a610af3565b6001600160a01b0316146103755760405162461bcd60e51b815260206004820152601f60248201527f476f7665726e6f723a206d7573742073656e6420746f206578656375746f720060448201526064015b60405180910390fd5b005b600080fd5b34801561038857600080fd5b5061039c610397366004614bba565b610b0d565b60405190151581526020015b60405180910390f35b3480156103bd57600080fd5b506103c6610b1e565b6040519081526020016103a8565b3480156103e057600080fd5b506103c66103ef366004614d3c565b610b2a565b34801561040057600080fd5b5061037561040f366004614de3565b610c22565b34801561042057600080fd5b50610429610cdc565b6040516103a89190614e58565b34801561044257600080fd5b50610478610451366004614de3565b600090815260fe60205260409020546801000000000000000090046001600160a01b031690565b6040516001600160a01b0390911681526020016103a8565b34801561049c57600080fd5b506104d46104ab366004614e80565b7f150b7a0200000000000000000000000000000000000000000000000000000000949350505050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020016103a8565b34801561051157600080fd5b506103c661052036600461505a565b610d6e565b6103c661053336600461505a565b611014565b34801561054457600080fd5b506103c6610553366004614de3565b600090815260fe602052604090205467ffffffffffffffff1690565b34801561057b57600080fd5b506103c67fb3b3f3b703cd84ce352197dcff232b1b5d3cfb2025ce47cf04742d0651f1af8881565b3480156105af57600080fd5b506103c661117a565b3480156105c457600080fd5b506103c66105d33660046150ea565b611186565b3480156105e457600080fd5b506105f86105f3366004614de3565b6111fc565b6040516103a89190615167565b34801561061157600080fd5b5061039c6106203660046151a8565b6000828152610161602090815260408083206001600160a01b038516845260030190915290205460ff1692915050565b34801561065c57600080fd5b506103c661066b36600461505a565b611207565b34801561067c57600080fd5b50610429611338565b34801561069157600080fd5b506106c06106a0366004614de3565b600090815261016160205260409020805460018201546002909201549092565b604080519384526020840192909252908201526060016103a8565b3480156106e757600080fd5b506104296113fe565b3480156106fc57600080fd5b506103c661070b3660046151d8565b611408565b34801561071c57600080fd5b506103c661072b366004615204565b611431565b34801561073c57600080fd5b506103c661074b366004614de3565b61147b565b34801561075c57600080fd5b5061037561076b366004614de3565b611570565b34801561077c57600080fd5b506103c661078b366004615288565b611627565b34801561079c57600080fd5b506103c66107ab3660046152e2565b61166f565b3480156107bc57600080fd5b506107c5611686565b6040516103a897969594939291906153d2565b3480156107e457600080fd5b506103756107f336600461544e565b611748565b34801561080457600080fd5b5061080d61192a565b60405165ffffffffffff90911681526020016103a8565b34801561083057600080fd5b5060646103c6565b34801561084457600080fd5b506103c66108533660046154a7565b6119b7565b34801561086457600080fd5b506103c66119ce565b34801561087957600080fd5b50610375610888366004615500565b611a10565b34801561089957600080fd5b506103c66108a8366004614de3565b611ac7565b3480156108b957600080fd5b506103c6611b7c565b3480156108ce57600080fd5b506104d46108dd36600461551d565b7fbc197c810000000000000000000000000000000000000000000000000000000095945050505050565b34801561091357600080fd5b506103c6610922366004614de3565b600090815260fe602052604090206001015467ffffffffffffffff1690565b61037561094f3660046155b1565b611b88565b34801561096057600080fd5b506103c661096f36600461505a565b611cbe565b34801561098057600080fd5b506101f8546001600160a01b0316610478565b34801561099f57600080fd5b506040805180820190915260208082527f737570706f72743d627261766f2671756f72756d3d666f722c6162737461696e90820152610429565b3480156109e557600080fd5b506103c67f150214d74d59b7d1e90c73fc22ef3d991dd0a76b046543d4d80ab92d2a50328f81565b348015610a1957600080fd5b50610375610a28366004614de3565b611cf8565b348015610a3957600080fd5b506103c6610a483660046155f5565b611daf565b348015610a5957600080fd5b50610375610a68366004614de3565b611dd0565b348015610a7957600080fd5b506104d4610a88366004615621565b7ff23a6e610000000000000000000000000000000000000000000000000000000095945050505050565b348015610abe57600080fd5b506103c6610acd366004614de3565b611e87565b348015610ade57600080fd5b5061019354610478906001600160a01b031681565b6000610b086101f8546001600160a01b031690565b905090565b6000610b1882611e92565b92915050565b6000610b086101305490565b600080610bce610bc67fb3b3f3b703cd84ce352197dcff232b1b5d3cfb2025ce47cf04742d0651f1af888c8c8c8c604051610b6692919061568a565b60405180910390208b80519060200120604051602001610bab959493929190948552602085019390935260ff9190911660408401526060830152608082015260a00190565b60405160208183030381529060405280519060200120611ee8565b868686611f30565b9050610c148a828b8b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508d9250611f4e915050565b9a9950505050505050505050565b610c2a610af3565b6001600160a01b0316336001600160a01b031614610c8a5760405162461bcd60e51b815260206004820152601860248201527f476f7665726e6f723a206f6e6c79476f7665726e616e63650000000000000000604482015260640161036c565b30610c93610af3565b6001600160a01b031614610cd05760008036604051610cb392919061568a565b604051809103902090505b80610cc960ff6120be565b03610cbe57505b610cd98161217b565b50565b606060fd8054610ceb9061569a565b80601f0160208091040260200160405190810160405280929190818152602001828054610d179061569a565b8015610d645780601f10610d3957610100808354040283529160200191610d64565b820191906000526020600020905b815481529060010190602001808311610d4757829003601f168201915b5050505050905090565b600080610d7d86868686611cbe565b90506004610d8a826111fc565b6007811115610d9b57610d9b615138565b14610e0e5760405162461bcd60e51b815260206004820152602160248201527f476f7665726e6f723a2070726f706f73616c206e6f742073756363657373667560448201527f6c00000000000000000000000000000000000000000000000000000000000000606482015260840161036c565b6101f854604080517ff27a0c9200000000000000000000000000000000000000000000000000000000815290516000926001600160a01b03169163f27a0c929160048083019260209291908290030181865afa158015610e72573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e9691906156ed565b6101f8546040517fb1c5f4270000000000000000000000000000000000000000000000000000000081529192506001600160a01b03169063b1c5f42790610eea908a908a908a906000908b90600401615794565b602060405180830381865afa158015610f07573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f2b91906156ed565b60008381526101f96020526040808220929092556101f85491517f8f2a0bb00000000000000000000000000000000000000000000000000000000081526001600160a01b0390921691638f2a0bb091610f91918b918b918b91908b9089906004016157e2565b600060405180830381600087803b158015610fab57600080fd5b505af1158015610fbf573d6000803e3d6000fd5b505050507f9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892828242610ff19190615869565b604080519283526020830191909152015b60405180910390a15095945050505050565b60008061102386868686611cbe565b90506000611030826111fc565b9050600481600781111561104657611046615138565b14806110635750600581600781111561106157611061615138565b145b6110d55760405162461bcd60e51b815260206004820152602160248201527f476f7665726e6f723a2070726f706f73616c206e6f742073756363657373667560448201527f6c00000000000000000000000000000000000000000000000000000000000000606482015260840161036c565b600082815260fe60205260409081902060020180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055517f712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f906111419084815260200190565b60405180910390a1611156828888888861231c565b61116382888888886123be565b61117082888888886123cb565b5095945050505050565b6000610b0861012f5490565b604080517f150214d74d59b7d1e90c73fc22ef3d991dd0a76b046543d4d80ab92d2a50328f602082015290810186905260ff8516606082015260009081906111d490610bc690608001610bab565b90506111f187828860405180602001604052806000815250612411565b979650505050505050565b6000610b1882612434565b60008061121686868686611cbe565b90506000611223826111fc565b600781111561123457611234615138565b146112815760405162461bcd60e51b815260206004820152601c60248201527f476f7665726e6f723a20746f6f206c61746520746f2063616e63656c00000000604482015260640161036c565b600081815260fe60205260409020546801000000000000000090046001600160a01b0316336001600160a01b0316146113225760405162461bcd60e51b815260206004820152602260248201527f476f7665726e6f723a206f6e6c792070726f706f7365722063616e2063616e6360448201527f656c000000000000000000000000000000000000000000000000000000000000606482015260840161036c565b61132e868686866125b3565b9695505050505050565b61019354604080517f4bf5d7e900000000000000000000000000000000000000000000000000000000815290516060926001600160a01b031691634bf5d7e99160048083019260009291908290030181865afa9250505080156113bd57506040513d6000823e601f3d908101601f191682016040526113ba9190810190615881565b60015b6113f9575060408051808201909152601d81527f6d6f64653d626c6f636b6e756d6265722666726f6d3d64656661756c74000000602082015290565b919050565b6060610b086125c1565b60008033905061142984828560405180602001604052806000815250612411565b949350505050565b6000803390506111f187828888888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508a9250611f4e915050565b6101c7546000908082036114945750506101c654919050565b60006101c76114a46001846158ef565b815481106114b4576114b4615906565b60009182526020918290206040805180820190915291015463ffffffff81168083526401000000009091047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16928201929092529150841061153557602001517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff169392505050565b61154a61154185612664565b6101c7906126e4565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16949350505050565b611578610af3565b6001600160a01b0316336001600160a01b0316146115d85760405162461bcd60e51b815260206004820152601860248201527f476f7665726e6f723a206f6e6c79476f7665726e616e63650000000000000000604482015260640161036c565b306115e1610af3565b6001600160a01b03161461161e576000803660405161160192919061568a565b604051809103902090505b8061161760ff6120be565b0361160c57505b610cd9816127ad565b60008033905061132e86828787878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061241192505050565b600061167d858585856127f0565b95945050505050565b6000606080600080600060606065546000801b1480156116a65750606654155b6116f25760405162461bcd60e51b815260206004820152601560248201527f4549503731323a20556e696e697469616c697a65640000000000000000000000604482015260640161036c565b6116fa612d25565b611702612d34565b604080516000808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009b939a50919850469750309650945092509050565b600054610100900460ff16158080156117685750600054600160ff909116105b806117825750303b158015611782575060005460ff166001145b6117f45760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161036c565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561185257600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6118906040518060400160405280600f81526020017f55706772616465476f7665726e6f720000000000000000000000000000000000815250612d43565b61189b858585612dda565b6118a3612e67565b6118ac87612ee6565b6118b582612f6c565b6118be86612ff2565b801561192157600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b61019354604080517f91ddadf400000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916391ddadf49160048083019260209291908290030181865afa9250505080156119aa575060408051601f3d908101601f191682019092526119a791810190615935565b60015b6113f957610b0843613078565b60006119c48484846130f6565b90505b9392505050565b6101c75460009015611a08576119e56101c7613186565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16905090565b506101c65490565b611a18610af3565b6001600160a01b0316336001600160a01b031614611a785760405162461bcd60e51b815260206004820152601860248201527f476f7665726e6f723a206f6e6c79476f7665726e616e63650000000000000000604482015260640161036c565b30611a81610af3565b6001600160a01b031614611abe5760008036604051611aa192919061568a565b604051809103902090505b80611ab760ff6120be565b03611aac57505b610cd9816131cc565b6101f85460008281526101f960205260408082205490517fd45c44350000000000000000000000000000000000000000000000000000000081526004810191909152909182916001600160a01b039091169063d45c443590602401602060405180830381865afa158015611b3f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b6391906156ed565b905080600114611b7357806119c7565b60009392505050565b6000610b086101315490565b611b90610af3565b6001600160a01b0316336001600160a01b031614611bf05760405162461bcd60e51b815260206004820152601860248201527f476f7665726e6f723a206f6e6c79476f7665726e616e63650000000000000000604482015260640161036c565b30611bf9610af3565b6001600160a01b031614611c365760008036604051611c1992919061568a565b604051809103902090505b80611c2f60ff6120be565b03611c2457505b600080856001600160a01b0316858585604051611c5492919061568a565b60006040518083038185875af1925050503d8060008114611c91576040519150601f19603f3d011682016040523d82523d6000602084013e611c96565b606091505b50915091506119218282604051806060016040528060288152602001615d306028913961324f565b600084848484604051602001611cd7949392919061595d565b60408051601f19818403018152919052805160209091012095945050505050565b611d00610af3565b6001600160a01b0316336001600160a01b031614611d605760405162461bcd60e51b815260206004820152601860248201527f476f7665726e6f723a206f6e6c79476f7665726e616e63650000000000000000604482015260640161036c565b30611d69610af3565b6001600160a01b031614611da65760008036604051611d8992919061568a565b604051809103902090505b80611d9f60ff6120be565b03611d9457505b610cd981613268565b60006119c78383611dcb60408051602081019091526000815290565b6130f6565b611dd8610af3565b6001600160a01b0316336001600160a01b031614611e385760405162461bcd60e51b815260206004820152601860248201527f476f7665726e6f723a206f6e6c79476f7665726e616e63650000000000000000604482015260640161036c565b30611e41610af3565b6001600160a01b031614611e7e5760008036604051611e6192919061568a565b604051809103902090505b80611e7760ff6120be565b03611e6c57505b610cd981613321565b6000610b1882613364565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f6e665ced000000000000000000000000000000000000000000000000000000001480610b185750610b188261340c565b6000610b18611ef56135b6565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b6000806000611f41878787876135c0565b9150915061117081613684565b600085815260fe602052604081206001611f67886111fc565b6007811115611f7857611f78615138565b14611feb5760405162461bcd60e51b815260206004820152602360248201527f476f7665726e6f723a20766f7465206e6f742063757272656e746c792061637460448201527f6976650000000000000000000000000000000000000000000000000000000000606482015260840161036c565b805460009061200690889067ffffffffffffffff16866130f6565b905061201588888884886137e9565b835160000361206a57866001600160a01b03167fb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda48988848960405161205d94939291906159a8565b60405180910390a26111f1565b866001600160a01b03167fe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb871289888489896040516120ab9594939291906159d0565b60405180910390a2979650505050505050565b60006120e68254600f81810b700100000000000000000000000000000000909204900b131590565b1561211d576040517f3db2a12a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b508054600f0b6000818152600180840160205260408220805492905583547fffffffffffffffffffffffffffffffff000000000000000000000000000000001692016fffffffffffffffffffffffffffffffff169190911790915590565b60648111156122185760405162461bcd60e51b815260206004820152604360248201527f476f7665726e6f72566f74657351756f72756d4672616374696f6e3a2071756f60448201527f72756d4e756d657261746f72206f7665722071756f72756d44656e6f6d696e6160648201527f746f720000000000000000000000000000000000000000000000000000000000608482015260a40161036c565b60006122226119ce565b9050801580159061223457506101c754155b156122af5760408051808201909152600081526101c79060208101612258846139be565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff90811690915282546001810184556000938452602093849020835194909301519091166401000000000263ffffffff909316929092179101555b6122dd6122ca6122bd61192a565b65ffffffffffff16612664565b6122d3846139be565b6101c79190613a52565b505060408051828152602081018490527f0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997910160405180910390a15050565b30612325610af3565b6001600160a01b0316146123b75760005b84518110156123b557306001600160a01b031685828151811061235b5761235b615906565b60200260200101516001600160a01b0316036123a5576123a583828151811061238657612386615906565b60200260200101518051906020012060ff613a6d90919063ffffffff16565b6123ae81615a16565b9050612336565b505b5050505050565b6123b78585858585613abf565b306123d4610af3565b6001600160a01b0316146123b75760ff54600f81810b700100000000000000000000000000000000909204900b13156123b757600060ff556123b7565b600061167d8585858561242f60408051602081019091526000815290565b611f4e565b60008061244083613b4d565b9050600481600781111561245657612456615138565b146124615792915050565b60008381526101f960205260409020548061247d575092915050565b6101f8546040517f2ab0f529000000000000000000000000000000000000000000000000000000008152600481018390526001600160a01b0390911690632ab0f52990602401602060405180830381865afa1580156124e0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125049190615a30565b15612513575060079392505050565b6101f8546040517f584b153e000000000000000000000000000000000000000000000000000000008152600481018390526001600160a01b039091169063584b153e90602401602060405180830381865afa158015612576573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061259a9190615a30565b156125a9575060059392505050565b5060029392505050565b600061167d85858585613c90565b60606125ec7f0000000000000000000000000000000000000000000000000000000000000000613d5f565b6126157f0000000000000000000000000000000000000000000000000000000000000000613d5f565b61263e7f0000000000000000000000000000000000000000000000000000000000000000613d5f565b60405160200161265093929190615a52565b604051602081830303815290604052905090565b600063ffffffff8211156126e05760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203360448201527f3220626974730000000000000000000000000000000000000000000000000000606482015260840161036c565b5090565b8154600090818160058111156127415760006126ff84613dff565b61270990856158ef565b60008881526020902090915081015463ffffffff90811690871610156127315780915061273f565b61273c816001615869565b92505b505b600061274f87878585613ee7565b905080156127a057612774876127666001846158ef565b600091825260209091200190565b5464010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff166111f1565b6000979650505050505050565b61012f5460408051918252602082018390527fc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93910160405180910390a161012f55565b6000336127fd8184613f45565b6128495760405162461bcd60e51b815260206004820152601d60248201527f476f7665726e6f723a2070726f706f7365722072657374726963746564000000604482015260640161036c565b600061285361192a565b65ffffffffffff169050612865611b7c565b61287483610a486001856158ef565b10156128e85760405162461bcd60e51b815260206004820152603160248201527f476f7665726e6f723a2070726f706f73657220766f7465732062656c6f77207060448201527f726f706f73616c207468726573686f6c64000000000000000000000000000000606482015260840161036c565b60006128fd8888888880519060200120611cbe565b905086518851146129765760405162461bcd60e51b815260206004820152602160248201527f476f7665726e6f723a20696e76616c69642070726f706f73616c206c656e677460448201527f6800000000000000000000000000000000000000000000000000000000000000606482015260840161036c565b85518851146129ed5760405162461bcd60e51b815260206004820152602160248201527f476f7665726e6f723a20696e76616c69642070726f706f73616c206c656e677460448201527f6800000000000000000000000000000000000000000000000000000000000000606482015260840161036c565b6000885111612a3e5760405162461bcd60e51b815260206004820152601860248201527f476f7665726e6f723a20656d7074792070726f706f73616c0000000000000000604482015260640161036c565b600081815260fe602052604090205467ffffffffffffffff1615612aca5760405162461bcd60e51b815260206004820152602160248201527f476f7665726e6f723a2070726f706f73616c20616c726561647920657869737460448201527f7300000000000000000000000000000000000000000000000000000000000000606482015260840161036c565b6000612ad461117a565b612ade9084615869565b90506000612aea610b1e565b612af49083615869565b90506040518060e00160405280612b0a84614095565b67ffffffffffffffff1681526001600160a01b038716602082015260006040820152606001612b3883614095565b67ffffffffffffffff9081168252600060208084018290526040808501839052606094850183905288835260fe8252918290208551815492870151878501519186167fffffffff0000000000000000000000000000000000000000000000000000000090941693909317680100000000000000006001600160a01b039094168402177bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167c010000000000000000000000000000000000000000000000000000000060e09290921c91909102178155938501516080860151908416921c0217600183015560a08301516002909201805460c0909401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00009094169215157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1692909217610100931515939093029290921790558a517f7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e091859188918e918e91811115612cc257612cc2614c56565b604051908082528060200260200182016040528015612cf557816020015b6060815260200190600190039081612ce05790505b508d88888f604051612d0f99989796959493929190615ac8565b60405180910390a1509098975050505050505050565b606060678054610ceb9061569a565b606060688054610ceb9061569a565b600054610100900460ff16612dc05760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b612dd181612dcc6113fe565b614115565b610cd9816141ba565b600054610100900460ff16612e575760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b612e62838383614247565b505050565b600054610100900460ff16612ee45760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b565b600054610100900460ff16612f635760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b610cd9816142df565b600054610100900460ff16612fe95760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b610cd981614397565b600054610100900460ff1661306f5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b610cd981614414565b600065ffffffffffff8211156126e05760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203460448201527f3820626974730000000000000000000000000000000000000000000000000000606482015260840161036c565b610193546040517f3a46b1a80000000000000000000000000000000000000000000000000000000081526001600160a01b038581166004830152602482018590526000921690633a46b1a890604401602060405180830381865afa158015613162573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119c491906156ed565b80546000908015611b73576131a0836127666001846158ef565b5464010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff166119c7565b6101f854604080516001600160a01b03928316815291831660208301527f08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401910160405180910390a16101f880547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6060831561325e5750816119c7565b6119c78383614491565b600081116132de5760405162461bcd60e51b815260206004820152602760248201527f476f7665726e6f7253657474696e67733a20766f74696e6720706572696f642060448201527f746f6f206c6f7700000000000000000000000000000000000000000000000000606482015260840161036c565b6101305460408051918252602082018390527f7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828910160405180910390a161013055565b6101315460408051918252602082018390527fccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc05461910160405180910390a161013155565b600060646133718361147b565b610193546040517f8e539e8c000000000000000000000000000000000000000000000000000000008152600481018690526001600160a01b0390911690638e539e8c90602401602060405180830381865afa1580156133d4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133f891906156ed565b6134029190615ba0565b610b189190615bee565b60007f51159c06000000000000000000000000000000000000000000000000000000007fc6fba1f8000000000000000000000000000000000000000000000000000000007fbf26d897000000000000000000000000000000000000000000000000000000007f79dd796f000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000086168214806134e657507fffffffff00000000000000000000000000000000000000000000000000000000868116908216145b8061351557507fffffffff00000000000000000000000000000000000000000000000000000000868116908516145b8061356157507fffffffff0000000000000000000000000000000000000000000000000000000086167f4e2312e000000000000000000000000000000000000000000000000000000000145b8061132e57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008716149695505050505050565b6000610b086144bb565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156135f7575060009050600361367b565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561364b573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166136745760006001925092505061367b565b9150600090505b94509492505050565b600081600481111561369857613698615138565b036136a05750565b60018160048111156136b4576136b4615138565b036137015760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161036c565b600281600481111561371557613715615138565b036137625760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161036c565b600381600481111561377657613776615138565b03610cd95760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f7565000000000000000000000000000000000000000000000000000000000000606482015260840161036c565b6000858152610161602090815260408083206001600160a01b0388168452600381019092529091205460ff16156138885760405162461bcd60e51b815260206004820152602760248201527f476f7665726e6f72566f74696e6753696d706c653a20766f746520616c72656160448201527f6479206361737400000000000000000000000000000000000000000000000000606482015260840161036c565b6001600160a01b0385166000908152600382016020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905560ff84166138f257828160000160008282546138e79190615869565b909155506123b59050565b60001960ff85160161391257828160010160008282546138e79190615869565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe60ff85160161395057828160020160008282546138e79190615869565b60405162461bcd60e51b815260206004820152603560248201527f476f7665726e6f72566f74696e6753696d706c653a20696e76616c696420766160448201527f6c756520666f7220656e756d20566f7465547970650000000000000000000000606482015260840161036c565b60007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8211156126e05760405162461bcd60e51b815260206004820152602760248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203260448201527f3234206269747300000000000000000000000000000000000000000000000000606482015260840161036c565b600080613a6085858561452f565b915091505b935093915050565b815470010000000000000000000000000000000090819004600f0b6000818152600180860160205260409091209390935583546fffffffffffffffffffffffffffffffff908116939091011602179055565b6101f8546040517fe38335e50000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063e38335e5903490613b14908890889088906000908990600401615794565b6000604051808303818588803b158015613b2d57600080fd5b505af1158015613b41573d6000803e3d6000fd5b50505050505050505050565b600081815260fe60205260408120600281015460ff1615613b715750600792915050565b6002810154610100900460ff1615613b8c5750600292915050565b600083815260fe602052604081205467ffffffffffffffff1690819003613bf55760405162461bcd60e51b815260206004820152601d60248201527f476f7665726e6f723a20756e6b6e6f776e2070726f706f73616c206964000000604482015260640161036c565b6000613bff61192a565b65ffffffffffff169050808210613c1b57506000949350505050565b600085815260fe602052604090206001015467ffffffffffffffff16818110613c4a5750600195945050505050565b613c5386614726565b8015613c7357506000868152610161602052604090208054600190910154115b15613c845750600495945050505050565b50600395945050505050565b600080613c9f86868686614774565b60008181526101f960205260409020549091501561167d576101f85460008281526101f96020526040908190205490517fc4d252f50000000000000000000000000000000000000000000000000000000081526001600160a01b039092169163c4d252f591613d149160040190815260200190565b600060405180830381600087803b158015613d2e57600080fd5b505af1158015613d42573d6000803e3d6000fd5b50505060008281526101f960205260408120555095945050505050565b60606000613d6c8361489d565b600101905060008167ffffffffffffffff811115613d8c57613d8c614c56565b6040519080825280601f01601f191660200182016040528015613db6576020820181803683370190505b5090508181016020015b600019017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a8504945084613dc057509392505050565b600081600003613e1157506000919050565b60006001613e1e8461497f565b901c6001901b90506001818481613e3757613e37615bbf565b048201901c90506001818481613e4f57613e4f615bbf565b048201901c90506001818481613e6757613e67615bbf565b048201901c90506001818481613e7f57613e7f615bbf565b048201901c90506001818481613e9757613e97615bbf565b048201901c90506001818481613eaf57613eaf615bbf565b048201901c90506001818481613ec757613ec7615bbf565b048201901c90506119c781828581613ee157613ee1615bbf565b04614a13565b60005b81831015613f3d576000613efe8484614a29565b60008781526020902090915063ffffffff86169082015463ffffffff161115613f2957809250613f37565b613f34816001615869565b93505b50613eea565b509392505050565b80516000906034811015613f5d576001915050610b18565b8281017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec01517fffffffffffffffffffffffff000000000000000000000000000000000000000081167f2370726f706f7365723d3078000000000000000000000000000000000000000014613fd757600192505050610b18565b600080613fe56028856158ef565b90505b838110156140745760008061403488848151811061400857614008615906565b01602001517fff0000000000000000000000000000000000000000000000000000000000000016614a44565b915091508161404c5760019650505050505050610b18565b8060ff166004856001600160a01b0316901b17935050508061406d90615a16565b9050613fe8565b50856001600160a01b0316816001600160a01b031614935050505092915050565b600067ffffffffffffffff8211156126e05760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203660448201527f3420626974730000000000000000000000000000000000000000000000000000606482015260840161036c565b600054610100900460ff166141925760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b606761419e8382615c6f565b5060686141ab8282615c6f565b50506000606581905560665550565b600054610100900460ff166142375760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b60fd6142438282615c6f565b5050565b600054610100900460ff166142c45760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b6142cd836127ad565b6142d682613268565b612e6281613321565b600054610100900460ff1661435c5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b61019380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b600054610100900460ff16610cd05760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b600054610100900460ff16611abe5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b8151156144a15781518083602001fd5b8060405162461bcd60e51b815260040161036c9190614e58565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6144e6614b30565b6144ee614b89565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b8254600090819080156146b757600061454d876127666001856158ef565b60408051808201909152905463ffffffff8082168084526401000000009092047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16602084015291925090871610156145e45760405162461bcd60e51b815260206004820152601b60248201527f436865636b706f696e743a2064656372656173696e67206b6579730000000000604482015260640161036c565b805163ffffffff8088169116036146425784614605886127666001866158ef565b80547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff929092166401000000000263ffffffff9092169190911790556146a7565b6040805180820190915263ffffffff80881682527bffffffffffffffffffffffffffffffffffffffffffffffffffffffff80881660208085019182528b54600181018d5560008d81529190912094519151909216640100000000029216919091179101555b602001519250839150613a659050565b50506040805180820190915263ffffffff80851682527bffffffffffffffffffffffffffffffffffffffffffffffffffffffff80851660208085019182528854600181018a5560008a815291822095519251909316640100000000029190931617920191909155905081613a65565b600081815261016160205260408120600281015460018201546147499190615869565b600084815260fe602052604090205461476b9067ffffffffffffffff16611e87565b11159392505050565b60008061478386868686611cbe565b90506000614790826111fc565b905060028160078111156147a6576147a6615138565b141580156147c6575060068160078111156147c3576147c3615138565b14155b80156147e4575060078160078111156147e1576147e1615138565b14155b6148305760405162461bcd60e51b815260206004820152601d60248201527f476f7665726e6f723a2070726f706f73616c206e6f7420616374697665000000604482015260640161036c565b600082815260fe60205260409081902060020180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16610100179055517f789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c906110029084815260200190565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f01000000000000000083106148e6577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310614912576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061493057662386f26fc10000830492506010015b6305f5e1008310614948576305f5e100830492506008015b612710831061495c57612710830492506004015b6064831061496e576064830492506002015b600a8310610b185760010192915050565b600080608083901c1561499457608092831c92015b604083901c156149a657604092831c92015b602083901c156149b857602092831c92015b601083901c156149ca57601092831c92015b600883901c156149dc57600892831c92015b600483901c156149ee57600492831c92015b600283901c15614a0057600292831c92015b600183901c15610b185760010192915050565b6000818310614a2257816119c7565b5090919050565b6000614a386002848418615bee565b6119c790848416615869565b60008060f883901c602f81118015614a5f5750603a8160ff16105b15614a92576001947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd09091019350915050565b8060ff166040108015614aa8575060478160ff16105b15614adb576001947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc99091019350915050565b8060ff166060108015614af1575060678160ff16105b15614b24576001947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa99091019350915050565b50600093849350915050565b600080614b3b612d25565b805190915015614b52578051602090910120919050565b6065548015614b615792915050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4709250505090565b600080614b94612d34565b805190915015614bab578051602090910120919050565b6066548015614b615792915050565b600060208284031215614bcc57600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146119c757600080fd5b803560ff811681146113f957600080fd5b60008083601f840112614c1f57600080fd5b50813567ffffffffffffffff811115614c3757600080fd5b602083019150836020828501011115614c4f57600080fd5b9250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715614cae57614cae614c56565b604052919050565b600067ffffffffffffffff821115614cd057614cd0614c56565b50601f01601f191660200190565b6000614cf1614cec84614cb6565b614c85565b9050828152838383011115614d0557600080fd5b828260208301376000602084830101529392505050565b600082601f830112614d2d57600080fd5b6119c783833560208501614cde565b60008060008060008060008060e0898b031215614d5857600080fd5b88359750614d6860208a01614bfc565b9650604089013567ffffffffffffffff80821115614d8557600080fd5b614d918c838d01614c0d565b909850965060608b0135915080821115614daa57600080fd5b50614db78b828c01614d1c565b945050614dc660808a01614bfc565b925060a0890135915060c089013590509295985092959890939650565b600060208284031215614df557600080fd5b5035919050565b60005b83811015614e17578181015183820152602001614dff565b83811115614e26576000848401525b50505050565b60008151808452614e44816020860160208601614dfc565b601f01601f19169290920160200192915050565b6020815260006119c76020830184614e2c565b6001600160a01b0381168114610cd957600080fd5b60008060008060808587031215614e9657600080fd5b8435614ea181614e6b565b93506020850135614eb181614e6b565b925060408501359150606085013567ffffffffffffffff811115614ed457600080fd5b614ee087828801614d1c565b91505092959194509250565b600067ffffffffffffffff821115614f0657614f06614c56565b5060051b60200190565b600082601f830112614f2157600080fd5b81356020614f31614cec83614eec565b82815260059290921b84018101918181019086841115614f5057600080fd5b8286015b84811015614f74578035614f6781614e6b565b8352918301918301614f54565b509695505050505050565b600082601f830112614f9057600080fd5b81356020614fa0614cec83614eec565b82815260059290921b84018101918181019086841115614fbf57600080fd5b8286015b84811015614f745780358352918301918301614fc3565b600082601f830112614feb57600080fd5b81356020614ffb614cec83614eec565b82815260059290921b8401810191818101908684111561501a57600080fd5b8286015b84811015614f7457803567ffffffffffffffff81111561503e5760008081fd5b61504c8986838b0101614d1c565b84525091830191830161501e565b6000806000806080858703121561507057600080fd5b843567ffffffffffffffff8082111561508857600080fd5b61509488838901614f10565b955060208701359150808211156150aa57600080fd5b6150b688838901614f7f565b945060408701359150808211156150cc57600080fd5b506150d987828801614fda565b949793965093946060013593505050565b600080600080600060a0868803121561510257600080fd5b8535945061511260208701614bfc565b935061512060408701614bfc565b94979396509394606081013594506080013592915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60208101600883106151a2577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b600080604083850312156151bb57600080fd5b8235915060208301356151cd81614e6b565b809150509250929050565b600080604083850312156151eb57600080fd5b823591506151fb60208401614bfc565b90509250929050565b60008060008060006080868803121561521c57600080fd5b8535945061522c60208701614bfc565b9350604086013567ffffffffffffffff8082111561524957600080fd5b61525589838a01614c0d565b9095509350606088013591508082111561526e57600080fd5b5061527b88828901614d1c565b9150509295509295909350565b6000806000806060858703121561529e57600080fd5b843593506152ae60208601614bfc565b9250604085013567ffffffffffffffff8111156152ca57600080fd5b6152d687828801614c0d565b95989497509550505050565b600080600080608085870312156152f857600080fd5b843567ffffffffffffffff8082111561531057600080fd5b61531c88838901614f10565b9550602087013591508082111561533257600080fd5b61533e88838901614f7f565b9450604087013591508082111561535457600080fd5b61536088838901614fda565b9350606087013591508082111561537657600080fd5b508501601f8101871361538857600080fd5b614ee087823560208401614cde565b600081518084526020808501945080840160005b838110156153c7578151875295820195908201906001016153ab565b509495945050505050565b7fff000000000000000000000000000000000000000000000000000000000000008816815260e06020820152600061540d60e0830189614e2c565b828103604084015261541f8189614e2c565b90508660608401526001600160a01b03861660808401528460a084015282810360c0840152610c148185615397565b60008060008060008060c0878903121561546757600080fd5b863561547281614e6b565b9550602087013561548281614e6b565b95989597505050506040840135936060810135936080820135935060a0909101359150565b6000806000606084860312156154bc57600080fd5b83356154c781614e6b565b925060208401359150604084013567ffffffffffffffff8111156154ea57600080fd5b6154f686828701614d1c565b9150509250925092565b60006020828403121561551257600080fd5b81356119c781614e6b565b600080600080600060a0868803121561553557600080fd5b853561554081614e6b565b9450602086013561555081614e6b565b9350604086013567ffffffffffffffff8082111561556d57600080fd5b61557989838a01614f7f565b9450606088013591508082111561558f57600080fd5b61559b89838a01614f7f565b9350608088013591508082111561526e57600080fd5b600080600080606085870312156155c757600080fd5b84356155d281614e6b565b935060208501359250604085013567ffffffffffffffff8111156152ca57600080fd5b6000806040838503121561560857600080fd5b823561561381614e6b565b946020939093013593505050565b600080600080600060a0868803121561563957600080fd5b853561564481614e6b565b9450602086013561565481614e6b565b93506040860135925060608601359150608086013567ffffffffffffffff81111561567e57600080fd5b61527b88828901614d1c565b8183823760009101908152919050565b600181811c908216806156ae57607f821691505b6020821081036156e7577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b6000602082840312156156ff57600080fd5b5051919050565b600081518084526020808501945080840160005b838110156153c75781516001600160a01b03168752958201959082019060010161571a565b600081518084526020808501808196508360051b8101915082860160005b85811015615787578284038952615775848351614e2c565b9885019893509084019060010161575d565b5091979650505050505050565b60a0815260006157a760a0830188615706565b82810360208401526157b98188615397565b905082810360408401526157cd818761573f565b60608401959095525050608001529392505050565b60c0815260006157f560c0830189615706565b82810360208401526158078189615397565b9050828103604084015261581b818861573f565b60608401969096525050608081019290925260a0909101529392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000821982111561587c5761587c61583a565b500190565b60006020828403121561589357600080fd5b815167ffffffffffffffff8111156158aa57600080fd5b8201601f810184136158bb57600080fd5b80516158c9614cec82614cb6565b8181528560208385010111156158de57600080fd5b61167d826020830160208601614dfc565b6000828210156159015761590161583a565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006020828403121561594757600080fd5b815165ffffffffffff811681146119c757600080fd5b6080815260006159706080830187615706565b82810360208401526159828187615397565b90508281036040840152615996818661573f565b91505082606083015295945050505050565b84815260ff8416602082015282604082015260806060820152600061132e6080830184614e2c565b85815260ff8516602082015283604082015260a0606082015260006159f860a0830185614e2c565b8281036080840152615a0a8185614e2c565b98975050505050505050565b60006000198203615a2957615a2961583a565b5060010190565b600060208284031215615a4257600080fd5b815180151581146119c757600080fd5b60008451615a64818460208901614dfc565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551615aa0816001850160208a01614dfc565b60019201918201528351615abb816002840160208801614dfc565b0160020195945050505050565b60006101208b835260206001600160a01b038c1681850152816040850152615af28285018c615706565b91508382036060850152615b06828b615397565b915083820360808501528189518084528284019150828160051b850101838c0160005b83811015615b5757601f19878403018552615b45838351614e2c565b94860194925090850190600101615b29565b505086810360a0880152615b6b818c61573f565b9450505050508560c08401528460e0840152828103610100840152615b908185614e2c565b9c9b505050505050505050505050565b6000816000190483118215151615615bba57615bba61583a565b500290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082615c24577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b601f821115612e6257600081815260208120601f850160051c81016020861015615c505750805b601f850160051c820191505b818110156123b557828155600101615c5c565b815167ffffffffffffffff811115615c8957615c89614c56565b615c9d81615c97845461569a565b84615c29565b602080601f831160018114615cd25760008415615cba5750858301515b600019600386901b1c1916600185901b1785556123b5565b600085815260208120601f198616915b82811015615d0157888601518255948401946001909101908401615ce2565b5085821015615d1f5787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fe476f7665726e6f723a2072656c617920726576657274656420776974686f7574206d657373616765a164736f6c634300080f000a",
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
// Solidity: function version() view returns(string)
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
// Solidity: function version() view returns(string)
func (_UpgradeGovernor *UpgradeGovernorSession) Version() (string, error) {
	return _UpgradeGovernor.Contract.Version(&_UpgradeGovernor.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
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
