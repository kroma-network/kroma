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

// KromaGovernorMetaData contains all meta data concerning the KromaGovernor contract.
var KromaGovernorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Empty\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"}],\"name\":\"ProposalQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldProposalThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newProposalThreshold\",\"type\":\"uint256\"}],\"name\":\"ProposalThresholdSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldQuorumNumerator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newQuorumNumerator\",\"type\":\"uint256\"}],\"name\":\"QuorumNumeratorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldTimelock\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTimelock\",\"type\":\"address\"}],\"name\":\"TimelockChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"VoteCastWithParams\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVotingDelay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotingDelay\",\"type\":\"uint256\"}],\"name\":\"VotingDelaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVotingPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotingPeriod\",\"type\":\"uint256\"}],\"name\":\"VotingPeriodSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BALLOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLOCK_MODE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COUNTING_MODE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXTENDED_BALLOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"}],\"name\":\"castVote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"castVoteBySig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"castVoteWithReason\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"castVoteWithReasonAndParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"castVoteWithReasonAndParamsBySig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clock\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"}],\"name\":\"getVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"getVotesWithParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"hashProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_timelock\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_initialVotingDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_initialVotingPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_initialProposalThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_votesQuorumFraction\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalDeadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalEta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalProposer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"againstVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"abstainVotes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"queue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"quorum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumDenominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"}],\"name\":\"quorumNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newProposalThreshold\",\"type\":\"uint256\"}],\"name\":\"setProposalThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newVotingDelay\",\"type\":\"uint256\"}],\"name\":\"setVotingDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newVotingPeriod\",\"type\":\"uint256\"}],\"name\":\"setVotingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumIGovernorUpgradeable.ProposalState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timelock\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC5805Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newQuorumNumerator\",\"type\":\"uint256\"}],\"name\":\"updateQuorumNumerator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTimelockControllerUpgradeable\",\"name\":\"newTimelock\",\"type\":\"address\"}],\"name\":\"updateTimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b506200001c62000031565b60006080819052600160a05260c052620000f2565b600054610100900460ff16156200009e5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811614620000f0576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b60805160a05160c051615d5a6200012260003960006114570152600061142e015260006114050152615d5a6000f3fe60806040526004361061030c5760003560e01c80637b3c71d31161019a578063c01f9e37116100e1578063ea0217cf1161008a578063f23a6e6111610064578063f23a6e6114610a6d578063f8ce560a14610ab2578063fc0c546a14610ad257600080fd5b8063ea0217cf14610a0d578063eb9019d414610a2d578063ece40cc114610a4d57600080fd5b8063d33219b4116100bb578063d33219b414610974578063dd4e2ba514610993578063deaaa7cc146109d957600080fd5b8063c01f9e3714610907578063c28bc2fa14610941578063c59057e41461095457600080fd5b80639a802a6d11610143578063ab58fb8e1161011d578063ab58fb8e1461088d578063b58131b0146108ad578063bc197c81146108c257600080fd5b80639a802a6d14610838578063a7713a7014610858578063a890c9101461086d57600080fd5b806386489ba91161017457806386489ba9146107d857806391ddadf4146107f857806397c3d3341461082457600080fd5b80637b3c71d3146107705780637d5e81e21461079057806384b0196e146107b057600080fd5b80633932abb11161025e578063544ffc9c116102075780635f398a14116101e15780635f398a141461071057806360c4247f1461073057806370b0f6601461075057600080fd5b8063544ffc9c1461068557806354fd4d50146106db57806356781388146106f057600080fd5b806343859632116102385780634385963214610605578063452115d6146106505780634bf5d7e91461067057600080fd5b80633932abb1146105a35780633bccf4fd146105b85780633e4f49e6146105d857600080fd5b8063143489d0116102c05780632656227d1161029a5780632656227d146105255780632d63f693146105385780632fe3e2611461056f57600080fd5b8063143489d014610436578063150b7a0214610490578063160cbed71461050557600080fd5b806303420181116102f157806303420181146103d457806306f3f9e6146103f457806306fdde031461041457600080fd5b806301ffc9a71461037c57806302a251a3146103b157600080fd5b36610377573061031a610af3565b6001600160a01b0316146103755760405162461bcd60e51b815260206004820152601f60248201527f476f7665726e6f723a206d7573742073656e6420746f206578656375746f720060448201526064015b60405180910390fd5b005b600080fd5b34801561038857600080fd5b5061039c610397366004614bb0565b610b0d565b60405190151581526020015b60405180910390f35b3480156103bd57600080fd5b506103c6610b1e565b6040519081526020016103a8565b3480156103e057600080fd5b506103c66103ef366004614d32565b610b2a565b34801561040057600080fd5b5061037561040f366004614dd9565b610c22565b34801561042057600080fd5b50610429610cdc565b6040516103a89190614e4e565b34801561044257600080fd5b50610478610451366004614dd9565b600090815260fe60205260409020546801000000000000000090046001600160a01b031690565b6040516001600160a01b0390911681526020016103a8565b34801561049c57600080fd5b506104d46104ab366004614e76565b7f150b7a0200000000000000000000000000000000000000000000000000000000949350505050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020016103a8565b34801561051157600080fd5b506103c6610520366004615050565b610d6e565b6103c6610533366004615050565b611014565b34801561054457600080fd5b506103c6610553366004614dd9565b600090815260fe602052604090205467ffffffffffffffff1690565b34801561057b57600080fd5b506103c67fb3b3f3b703cd84ce352197dcff232b1b5d3cfb2025ce47cf04742d0651f1af8881565b3480156105af57600080fd5b506103c661117a565b3480156105c457600080fd5b506103c66105d33660046150e0565b611186565b3480156105e457600080fd5b506105f86105f3366004614dd9565b6111fc565b6040516103a8919061515d565b34801561061157600080fd5b5061039c61062036600461519e565b6000828152610161602090815260408083206001600160a01b038516845260030190915290205460ff1692915050565b34801561065c57600080fd5b506103c661066b366004615050565b611207565b34801561067c57600080fd5b50610429611338565b34801561069157600080fd5b506106c06106a0366004614dd9565b600090815261016160205260409020805460018201546002909201549092565b604080519384526020840192909252908201526060016103a8565b3480156106e757600080fd5b506104296113fe565b3480156106fc57600080fd5b506103c661070b3660046151ce565b6114a1565b34801561071c57600080fd5b506103c661072b3660046151fa565b6114ca565b34801561073c57600080fd5b506103c661074b366004614dd9565b611514565b34801561075c57600080fd5b5061037561076b366004614dd9565b611609565b34801561077c57600080fd5b506103c661078b36600461527e565b6116c0565b34801561079c57600080fd5b506103c66107ab3660046152d8565b611708565b3480156107bc57600080fd5b506107c561171f565b6040516103a897969594939291906153c8565b3480156107e457600080fd5b506103756107f3366004615444565b6117e1565b34801561080457600080fd5b5061080d6119c3565b60405165ffffffffffff90911681526020016103a8565b34801561083057600080fd5b5060646103c6565b34801561084457600080fd5b506103c661085336600461549d565b611a50565b34801561086457600080fd5b506103c6611a67565b34801561087957600080fd5b506103756108883660046154f6565b611aa9565b34801561089957600080fd5b506103c66108a8366004614dd9565b611b60565b3480156108b957600080fd5b506103c6611c15565b3480156108ce57600080fd5b506104d46108dd366004615513565b7fbc197c810000000000000000000000000000000000000000000000000000000095945050505050565b34801561091357600080fd5b506103c6610922366004614dd9565b600090815260fe602052604090206001015467ffffffffffffffff1690565b61037561094f3660046155a7565b611c21565b34801561096057600080fd5b506103c661096f366004615050565b611d57565b34801561098057600080fd5b506101f8546001600160a01b0316610478565b34801561099f57600080fd5b506040805180820190915260208082527f737570706f72743d627261766f2671756f72756d3d666f722c6162737461696e90820152610429565b3480156109e557600080fd5b506103c67f150214d74d59b7d1e90c73fc22ef3d991dd0a76b046543d4d80ab92d2a50328f81565b348015610a1957600080fd5b50610375610a28366004614dd9565b611d91565b348015610a3957600080fd5b506103c6610a483660046155eb565b611e48565b348015610a5957600080fd5b50610375610a68366004614dd9565b611e69565b348015610a7957600080fd5b506104d4610a88366004615617565b7ff23a6e610000000000000000000000000000000000000000000000000000000095945050505050565b348015610abe57600080fd5b506103c6610acd366004614dd9565b611f20565b348015610ade57600080fd5b5061019354610478906001600160a01b031681565b6000610b086101f8546001600160a01b031690565b905090565b6000610b1882611f2b565b92915050565b6000610b086101305490565b600080610bce610bc67fb3b3f3b703cd84ce352197dcff232b1b5d3cfb2025ce47cf04742d0651f1af888c8c8c8c604051610b66929190615680565b60405180910390208b80519060200120604051602001610bab959493929190948552602085019390935260ff9190911660408401526060830152608082015260a00190565b60405160208183030381529060405280519060200120611f81565b868686611fc9565b9050610c148a828b8b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508d9250611fe7915050565b9a9950505050505050505050565b610c2a610af3565b6001600160a01b0316336001600160a01b031614610c8a5760405162461bcd60e51b815260206004820152601860248201527f476f7665726e6f723a206f6e6c79476f7665726e616e63650000000000000000604482015260640161036c565b30610c93610af3565b6001600160a01b031614610cd05760008036604051610cb3929190615680565b604051809103902090505b80610cc960ff612157565b03610cbe57505b610cd981612214565b50565b606060fd8054610ceb90615690565b80601f0160208091040260200160405190810160405280929190818152602001828054610d1790615690565b8015610d645780601f10610d3957610100808354040283529160200191610d64565b820191906000526020600020905b815481529060010190602001808311610d4757829003601f168201915b5050505050905090565b600080610d7d86868686611d57565b90506004610d8a826111fc565b6007811115610d9b57610d9b61512e565b14610e0e5760405162461bcd60e51b815260206004820152602160248201527f476f7665726e6f723a2070726f706f73616c206e6f742073756363657373667560448201527f6c00000000000000000000000000000000000000000000000000000000000000606482015260840161036c565b6101f854604080517ff27a0c9200000000000000000000000000000000000000000000000000000000815290516000926001600160a01b03169163f27a0c929160048083019260209291908290030181865afa158015610e72573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e9691906156e3565b6101f8546040517fb1c5f4270000000000000000000000000000000000000000000000000000000081529192506001600160a01b03169063b1c5f42790610eea908a908a908a906000908b9060040161578a565b602060405180830381865afa158015610f07573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f2b91906156e3565b60008381526101f96020526040808220929092556101f85491517f8f2a0bb00000000000000000000000000000000000000000000000000000000081526001600160a01b0390921691638f2a0bb091610f91918b918b918b91908b9089906004016157d8565b600060405180830381600087803b158015610fab57600080fd5b505af1158015610fbf573d6000803e3d6000fd5b505050507f9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892828242610ff1919061585f565b604080519283526020830191909152015b60405180910390a15095945050505050565b60008061102386868686611d57565b90506000611030826111fc565b905060048160078111156110465761104661512e565b1480611063575060058160078111156110615761106161512e565b145b6110d55760405162461bcd60e51b815260206004820152602160248201527f476f7665726e6f723a2070726f706f73616c206e6f742073756363657373667560448201527f6c00000000000000000000000000000000000000000000000000000000000000606482015260840161036c565b600082815260fe60205260409081902060020180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055517f712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f906111419084815260200190565b60405180910390a161115682888888886123b5565b6111638288888888612457565b6111708288888888612464565b5095945050505050565b6000610b0861012f5490565b604080517f150214d74d59b7d1e90c73fc22ef3d991dd0a76b046543d4d80ab92d2a50328f602082015290810186905260ff8516606082015260009081906111d490610bc690608001610bab565b90506111f1878288604051806020016040528060008152506124aa565b979650505050505050565b6000610b18826124cd565b60008061121686868686611d57565b90506000611223826111fc565b60078111156112345761123461512e565b146112815760405162461bcd60e51b815260206004820152601c60248201527f476f7665726e6f723a20746f6f206c61746520746f2063616e63656c00000000604482015260640161036c565b600081815260fe60205260409020546801000000000000000090046001600160a01b0316336001600160a01b0316146113225760405162461bcd60e51b815260206004820152602260248201527f476f7665726e6f723a206f6e6c792070726f706f7365722063616e2063616e6360448201527f656c000000000000000000000000000000000000000000000000000000000000606482015260840161036c565b61132e8686868661264c565b9695505050505050565b61019354604080517f4bf5d7e900000000000000000000000000000000000000000000000000000000815290516060926001600160a01b031691634bf5d7e99160048083019260009291908290030181865afa9250505080156113bd57506040513d6000823e601f3d908101601f191682016040526113ba9190810190615877565b60015b6113f9575060408051808201909152601d81527f6d6f64653d626c6f636b6e756d6265722666726f6d3d64656661756c74000000602082015290565b919050565b60606114297f000000000000000000000000000000000000000000000000000000000000000061265a565b6114527f000000000000000000000000000000000000000000000000000000000000000061265a565b61147b7f000000000000000000000000000000000000000000000000000000000000000061265a565b60405160200161148d939291906158e5565b604051602081830303815290604052905090565b6000803390506114c2848285604051806020016040528060008152506124aa565b949350505050565b6000803390506111f187828888888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508a9250611fe7915050565b6101c75460009080820361152d5750506101c654919050565b60006101c761153d60018461595b565b8154811061154d5761154d615972565b60009182526020918290206040805180820190915291015463ffffffff81168083526401000000009091047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1692820192909252915084106115ce57602001517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff169392505050565b6115e36115da856126fa565b6101c79061277a565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16949350505050565b611611610af3565b6001600160a01b0316336001600160a01b0316146116715760405162461bcd60e51b815260206004820152601860248201527f476f7665726e6f723a206f6e6c79476f7665726e616e63650000000000000000604482015260640161036c565b3061167a610af3565b6001600160a01b0316146116b7576000803660405161169a929190615680565b604051809103902090505b806116b060ff612157565b036116a557505b610cd981612843565b60008033905061132e86828787878080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506124aa92505050565b600061171685858585612886565b95945050505050565b6000606080600080600060606065546000801b14801561173f5750606654155b61178b5760405162461bcd60e51b815260206004820152601560248201527f4549503731323a20556e696e697469616c697a65640000000000000000000000604482015260640161036c565b611793612dbb565b61179b612dca565b604080516000808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009b939a50919850469750309650945092509050565b600054610100900460ff16158080156118015750600054600160ff909116105b8061181b5750303b15801561181b575060005460ff166001145b61188d5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161036c565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156118eb57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6119296040518060400160405280600d81526020017f4b726f6d61476f7665726e6f7200000000000000000000000000000000000000815250612dd9565b611934858585612e70565b61193c612efd565b61194587612f7c565b61194e82613002565b61195786613088565b80156119ba57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b61019354604080517f91ddadf400000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916391ddadf49160048083019260209291908290030181865afa925050508015611a43575060408051601f3d908101601f19168201909252611a40918101906159a1565b60015b6113f957610b084361310e565b6000611a5d84848461318c565b90505b9392505050565b6101c75460009015611aa157611a7e6101c761321c565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16905090565b506101c65490565b611ab1610af3565b6001600160a01b0316336001600160a01b031614611b115760405162461bcd60e51b815260206004820152601860248201527f476f7665726e6f723a206f6e6c79476f7665726e616e63650000000000000000604482015260640161036c565b30611b1a610af3565b6001600160a01b031614611b575760008036604051611b3a929190615680565b604051809103902090505b80611b5060ff612157565b03611b4557505b610cd981613262565b6101f85460008281526101f960205260408082205490517fd45c44350000000000000000000000000000000000000000000000000000000081526004810191909152909182916001600160a01b039091169063d45c443590602401602060405180830381865afa158015611bd8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bfc91906156e3565b905080600114611c0c5780611a60565b60009392505050565b6000610b086101315490565b611c29610af3565b6001600160a01b0316336001600160a01b031614611c895760405162461bcd60e51b815260206004820152601860248201527f476f7665726e6f723a206f6e6c79476f7665726e616e63650000000000000000604482015260640161036c565b30611c92610af3565b6001600160a01b031614611ccf5760008036604051611cb2929190615680565b604051809103902090505b80611cc860ff612157565b03611cbd57505b600080856001600160a01b0316858585604051611ced929190615680565b60006040518083038185875af1925050503d8060008114611d2a576040519150601f19603f3d011682016040523d82523d6000602084013e611d2f565b606091505b50915091506119ba8282604051806060016040528060288152602001615d26602891396132e5565b600084848484604051602001611d7094939291906159c9565b60408051601f19818403018152919052805160209091012095945050505050565b611d99610af3565b6001600160a01b0316336001600160a01b031614611df95760405162461bcd60e51b815260206004820152601860248201527f476f7665726e6f723a206f6e6c79476f7665726e616e63650000000000000000604482015260640161036c565b30611e02610af3565b6001600160a01b031614611e3f5760008036604051611e22929190615680565b604051809103902090505b80611e3860ff612157565b03611e2d57505b610cd9816132fe565b6000611a608383611e6460408051602081019091526000815290565b61318c565b611e71610af3565b6001600160a01b0316336001600160a01b031614611ed15760405162461bcd60e51b815260206004820152601860248201527f476f7665726e6f723a206f6e6c79476f7665726e616e63650000000000000000604482015260640161036c565b30611eda610af3565b6001600160a01b031614611f175760008036604051611efa929190615680565b604051809103902090505b80611f1060ff612157565b03611f0557505b610cd9816133b7565b6000610b18826133fa565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f6e665ced000000000000000000000000000000000000000000000000000000001480610b185750610b18826134a2565b6000610b18611f8e61364c565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b6000806000611fda87878787613656565b915091506111708161371a565b600085815260fe602052604081206001612000886111fc565b60078111156120115761201161512e565b146120845760405162461bcd60e51b815260206004820152602360248201527f476f7665726e6f723a20766f7465206e6f742063757272656e746c792061637460448201527f6976650000000000000000000000000000000000000000000000000000000000606482015260840161036c565b805460009061209f90889067ffffffffffffffff168661318c565b90506120ae888888848861387f565b835160000361210357866001600160a01b03167fb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4898884896040516120f69493929190615a14565b60405180910390a26111f1565b866001600160a01b03167fe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb87128988848989604051612144959493929190615a3c565b60405180910390a2979650505050505050565b600061217f8254600f81810b700100000000000000000000000000000000909204900b131590565b156121b6576040517f3db2a12a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b508054600f0b6000818152600180840160205260408220805492905583547fffffffffffffffffffffffffffffffff000000000000000000000000000000001692016fffffffffffffffffffffffffffffffff169190911790915590565b60648111156122b15760405162461bcd60e51b815260206004820152604360248201527f476f7665726e6f72566f74657351756f72756d4672616374696f6e3a2071756f60448201527f72756d4e756d657261746f72206f7665722071756f72756d44656e6f6d696e6160648201527f746f720000000000000000000000000000000000000000000000000000000000608482015260a40161036c565b60006122bb611a67565b905080158015906122cd57506101c754155b156123485760408051808201909152600081526101c790602081016122f184613a54565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff90811690915282546001810184556000938452602093849020835194909301519091166401000000000263ffffffff909316929092179101555b6123766123636123566119c3565b65ffffffffffff166126fa565b61236c84613a54565b6101c79190613ae8565b505060408051828152602081018490527f0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997910160405180910390a15050565b306123be610af3565b6001600160a01b0316146124505760005b845181101561244e57306001600160a01b03168582815181106123f4576123f4615972565b60200260200101516001600160a01b03160361243e5761243e83828151811061241f5761241f615972565b60200260200101518051906020012060ff613b0390919063ffffffff16565b61244781615a82565b90506123cf565b505b5050505050565b6124508585858585613b55565b3061246d610af3565b6001600160a01b0316146124505760ff54600f81810b700100000000000000000000000000000000909204900b131561245057600060ff55612450565b6000611716858585856124c860408051602081019091526000815290565b611fe7565b6000806124d983613be3565b905060048160078111156124ef576124ef61512e565b146124fa5792915050565b60008381526101f9602052604090205480612516575092915050565b6101f8546040517f2ab0f529000000000000000000000000000000000000000000000000000000008152600481018390526001600160a01b0390911690632ab0f52990602401602060405180830381865afa158015612579573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061259d9190615a9c565b156125ac575060079392505050565b6101f8546040517f584b153e000000000000000000000000000000000000000000000000000000008152600481018390526001600160a01b039091169063584b153e90602401602060405180830381865afa15801561260f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126339190615a9c565b15612642575060059392505050565b5060029392505050565b600061171685858585613d26565b6060600061266783613df5565b600101905060008167ffffffffffffffff81111561268757612687614c4c565b6040519080825280601f01601f1916602001820160405280156126b1576020820181803683370190505b5090508181016020015b600019017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a85049450846126bb57509392505050565b600063ffffffff8211156127765760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203360448201527f3220626974730000000000000000000000000000000000000000000000000000606482015260840161036c565b5090565b8154600090818160058111156127d757600061279584613ed7565b61279f908561595b565b60008881526020902090915081015463ffffffff90811690871610156127c7578091506127d5565b6127d281600161585f565b92505b505b60006127e587878585613fbf565b905080156128365761280a876127fc60018461595b565b600091825260209091200190565b5464010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff166111f1565b6000979650505050505050565b61012f5460408051918252602082018390527fc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93910160405180910390a161012f55565b600033612893818461401d565b6128df5760405162461bcd60e51b815260206004820152601d60248201527f476f7665726e6f723a2070726f706f7365722072657374726963746564000000604482015260640161036c565b60006128e96119c3565b65ffffffffffff1690506128fb611c15565b61290a83610a4860018561595b565b101561297e5760405162461bcd60e51b815260206004820152603160248201527f476f7665726e6f723a2070726f706f73657220766f7465732062656c6f77207060448201527f726f706f73616c207468726573686f6c64000000000000000000000000000000606482015260840161036c565b60006129938888888880519060200120611d57565b90508651885114612a0c5760405162461bcd60e51b815260206004820152602160248201527f476f7665726e6f723a20696e76616c69642070726f706f73616c206c656e677460448201527f6800000000000000000000000000000000000000000000000000000000000000606482015260840161036c565b8551885114612a835760405162461bcd60e51b815260206004820152602160248201527f476f7665726e6f723a20696e76616c69642070726f706f73616c206c656e677460448201527f6800000000000000000000000000000000000000000000000000000000000000606482015260840161036c565b6000885111612ad45760405162461bcd60e51b815260206004820152601860248201527f476f7665726e6f723a20656d7074792070726f706f73616c0000000000000000604482015260640161036c565b600081815260fe602052604090205467ffffffffffffffff1615612b605760405162461bcd60e51b815260206004820152602160248201527f476f7665726e6f723a2070726f706f73616c20616c726561647920657869737460448201527f7300000000000000000000000000000000000000000000000000000000000000606482015260840161036c565b6000612b6a61117a565b612b74908461585f565b90506000612b80610b1e565b612b8a908361585f565b90506040518060e00160405280612ba08461416d565b67ffffffffffffffff1681526001600160a01b038716602082015260006040820152606001612bce8361416d565b67ffffffffffffffff9081168252600060208084018290526040808501839052606094850183905288835260fe8252918290208551815492870151878501519186167fffffffff0000000000000000000000000000000000000000000000000000000090941693909317680100000000000000006001600160a01b039094168402177bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167c010000000000000000000000000000000000000000000000000000000060e09290921c91909102178155938501516080860151908416921c0217600183015560a08301516002909201805460c0909401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00009094169215157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1692909217610100931515939093029290921790558a517f7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e091859188918e918e91811115612d5857612d58614c4c565b604051908082528060200260200182016040528015612d8b57816020015b6060815260200190600190039081612d765790505b508d88888f604051612da599989796959493929190615aed565b60405180910390a1509098975050505050505050565b606060678054610ceb90615690565b606060688054610ceb90615690565b600054610100900460ff16612e565760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b612e6781612e626113fe565b6141ed565b610cd981614292565b600054610100900460ff16612eed5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b612ef883838361431f565b505050565b600054610100900460ff16612f7a5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b565b600054610100900460ff16612ff95760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b610cd9816143b7565b600054610100900460ff1661307f5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b610cd98161446f565b600054610100900460ff166131055760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b610cd9816144ec565b600065ffffffffffff8211156127765760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203460448201527f3820626974730000000000000000000000000000000000000000000000000000606482015260840161036c565b610193546040517f3a46b1a80000000000000000000000000000000000000000000000000000000081526001600160a01b038581166004830152602482018590526000921690633a46b1a890604401602060405180830381865afa1580156131f8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a5d91906156e3565b80546000908015611c0c57613236836127fc60018461595b565b5464010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16611a60565b6101f854604080516001600160a01b03928316815291831660208301527f08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401910160405180910390a16101f880547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b606083156132f4575081611a60565b611a608383614569565b600081116133745760405162461bcd60e51b815260206004820152602760248201527f476f7665726e6f7253657474696e67733a20766f74696e6720706572696f642060448201527f746f6f206c6f7700000000000000000000000000000000000000000000000000606482015260840161036c565b6101305460408051918252602082018390527f7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828910160405180910390a161013055565b6101315460408051918252602082018390527fccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc05461910160405180910390a161013155565b6000606461340783611514565b610193546040517f8e539e8c000000000000000000000000000000000000000000000000000000008152600481018690526001600160a01b0390911690638e539e8c90602401602060405180830381865afa15801561346a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061348e91906156e3565b6134989190615bc5565b610b189190615be4565b60007f51159c06000000000000000000000000000000000000000000000000000000007fc6fba1f8000000000000000000000000000000000000000000000000000000007fbf26d897000000000000000000000000000000000000000000000000000000007f79dd796f000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000861682148061357c57507fffffffff00000000000000000000000000000000000000000000000000000000868116908216145b806135ab57507fffffffff00000000000000000000000000000000000000000000000000000000868116908516145b806135f757507fffffffff0000000000000000000000000000000000000000000000000000000086167f4e2312e000000000000000000000000000000000000000000000000000000000145b8061132e57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008716149695505050505050565b6000610b08614593565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561368d5750600090506003613711565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa1580156136e1573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661370a57600060019250925050613711565b9150600090505b94509492505050565b600081600481111561372e5761372e61512e565b036137365750565b600181600481111561374a5761374a61512e565b036137975760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161036c565b60028160048111156137ab576137ab61512e565b036137f85760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161036c565b600381600481111561380c5761380c61512e565b03610cd95760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f7565000000000000000000000000000000000000000000000000000000000000606482015260840161036c565b6000858152610161602090815260408083206001600160a01b0388168452600381019092529091205460ff161561391e5760405162461bcd60e51b815260206004820152602760248201527f476f7665726e6f72566f74696e6753696d706c653a20766f746520616c72656160448201527f6479206361737400000000000000000000000000000000000000000000000000606482015260840161036c565b6001600160a01b0385166000908152600382016020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905560ff8416613988578281600001600082825461397d919061585f565b9091555061244e9050565b60001960ff8516016139a8578281600101600082825461397d919061585f565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe60ff8516016139e6578281600201600082825461397d919061585f565b60405162461bcd60e51b815260206004820152603560248201527f476f7665726e6f72566f74696e6753696d706c653a20696e76616c696420766160448201527f6c756520666f7220656e756d20566f7465547970650000000000000000000000606482015260840161036c565b60007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8211156127765760405162461bcd60e51b815260206004820152602760248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203260448201527f3234206269747300000000000000000000000000000000000000000000000000606482015260840161036c565b600080613af6858585614607565b915091505b935093915050565b815470010000000000000000000000000000000090819004600f0b6000818152600180860160205260409091209390935583546fffffffffffffffffffffffffffffffff908116939091011602179055565b6101f8546040517fe38335e50000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063e38335e5903490613baa90889088908890600090899060040161578a565b6000604051808303818588803b158015613bc357600080fd5b505af1158015613bd7573d6000803e3d6000fd5b50505050505050505050565b600081815260fe60205260408120600281015460ff1615613c075750600792915050565b6002810154610100900460ff1615613c225750600292915050565b600083815260fe602052604081205467ffffffffffffffff1690819003613c8b5760405162461bcd60e51b815260206004820152601d60248201527f476f7665726e6f723a20756e6b6e6f776e2070726f706f73616c206964000000604482015260640161036c565b6000613c956119c3565b65ffffffffffff169050808210613cb157506000949350505050565b600085815260fe602052604090206001015467ffffffffffffffff16818110613ce05750600195945050505050565b613ce9866147fe565b8015613d0957506000868152610161602052604090208054600190910154115b15613d1a5750600495945050505050565b50600395945050505050565b600080613d358686868661484c565b60008181526101f9602052604090205490915015611716576101f85460008281526101f96020526040908190205490517fc4d252f50000000000000000000000000000000000000000000000000000000081526001600160a01b039092169163c4d252f591613daa9160040190815260200190565b600060405180830381600087803b158015613dc457600080fd5b505af1158015613dd8573d6000803e3d6000fd5b50505060008281526101f960205260408120555095945050505050565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310613e3e577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310613e6a576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc100008310613e8857662386f26fc10000830492506010015b6305f5e1008310613ea0576305f5e100830492506008015b6127108310613eb457612710830492506004015b60648310613ec6576064830492506002015b600a8310610b185760010192915050565b600081600003613ee957506000919050565b60006001613ef684614975565b901c6001901b90506001818481613f0f57613f0f615abe565b048201901c90506001818481613f2757613f27615abe565b048201901c90506001818481613f3f57613f3f615abe565b048201901c90506001818481613f5757613f57615abe565b048201901c90506001818481613f6f57613f6f615abe565b048201901c90506001818481613f8757613f87615abe565b048201901c90506001818481613f9f57613f9f615abe565b048201901c9050611a6081828581613fb957613fb9615abe565b04614a09565b60005b81831015614015576000613fd68484614a1f565b60008781526020902090915063ffffffff86169082015463ffffffff1611156140015780925061400f565b61400c81600161585f565b93505b50613fc2565b509392505050565b80516000906034811015614035576001915050610b18565b8281017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec01517fffffffffffffffffffffffff000000000000000000000000000000000000000081167f2370726f706f7365723d30780000000000000000000000000000000000000000146140af57600192505050610b18565b6000806140bd60288561595b565b90505b8381101561414c5760008061410c8884815181106140e0576140e0615972565b01602001517fff0000000000000000000000000000000000000000000000000000000000000016614a3a565b91509150816141245760019650505050505050610b18565b8060ff166004856001600160a01b0316901b17935050508061414590615a82565b90506140c0565b50856001600160a01b0316816001600160a01b031614935050505092915050565b600067ffffffffffffffff8211156127765760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203660448201527f3420626974730000000000000000000000000000000000000000000000000000606482015260840161036c565b600054610100900460ff1661426a5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b60676142768382615c65565b5060686142838282615c65565b50506000606581905560665550565b600054610100900460ff1661430f5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b60fd61431b8282615c65565b5050565b600054610100900460ff1661439c5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b6143a583612843565b6143ae826132fe565b612ef8816133b7565b600054610100900460ff166144345760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b61019380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b600054610100900460ff16610cd05760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b600054610100900460ff16611b575760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b8151156145795781518083602001fd5b8060405162461bcd60e51b815260040161036c9190614e4e565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6145be614b26565b6145c6614b7f565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b82546000908190801561478f576000614625876127fc60018561595b565b60408051808201909152905463ffffffff8082168084526401000000009092047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16602084015291925090871610156146bc5760405162461bcd60e51b815260206004820152601b60248201527f436865636b706f696e743a2064656372656173696e67206b6579730000000000604482015260640161036c565b805163ffffffff80881691160361471a57846146dd886127fc60018661595b565b80547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff929092166401000000000263ffffffff90921691909117905561477f565b6040805180820190915263ffffffff80881682527bffffffffffffffffffffffffffffffffffffffffffffffffffffffff80881660208085019182528b54600181018d5560008d81529190912094519151909216640100000000029216919091179101555b602001519250839150613afb9050565b50506040805180820190915263ffffffff80851682527bffffffffffffffffffffffffffffffffffffffffffffffffffffffff80851660208085019182528854600181018a5560008a815291822095519251909316640100000000029190931617920191909155905081613afb565b60008181526101616020526040812060028101546001820154614821919061585f565b600084815260fe60205260409020546148439067ffffffffffffffff16611f20565b11159392505050565b60008061485b86868686611d57565b90506000614868826111fc565b9050600281600781111561487e5761487e61512e565b1415801561489e5750600681600781111561489b5761489b61512e565b14155b80156148bc575060078160078111156148b9576148b961512e565b14155b6149085760405162461bcd60e51b815260206004820152601d60248201527f476f7665726e6f723a2070726f706f73616c206e6f7420616374697665000000604482015260640161036c565b600082815260fe60205260409081902060020180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16610100179055517f789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c906110029084815260200190565b600080608083901c1561498a57608092831c92015b604083901c1561499c57604092831c92015b602083901c156149ae57602092831c92015b601083901c156149c057601092831c92015b600883901c156149d257600892831c92015b600483901c156149e457600492831c92015b600283901c156149f657600292831c92015b600183901c15610b185760010192915050565b6000818310614a185781611a60565b5090919050565b6000614a2e6002848418615be4565b611a609084841661585f565b60008060f883901c602f81118015614a555750603a8160ff16105b15614a88576001947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd09091019350915050565b8060ff166040108015614a9e575060478160ff16105b15614ad1576001947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc99091019350915050565b8060ff166060108015614ae7575060678160ff16105b15614b1a576001947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa99091019350915050565b50600093849350915050565b600080614b31612dbb565b805190915015614b48578051602090910120919050565b6065548015614b575792915050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4709250505090565b600080614b8a612dca565b805190915015614ba1578051602090910120919050565b6066548015614b575792915050565b600060208284031215614bc257600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114611a6057600080fd5b803560ff811681146113f957600080fd5b60008083601f840112614c1557600080fd5b50813567ffffffffffffffff811115614c2d57600080fd5b602083019150836020828501011115614c4557600080fd5b9250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715614ca457614ca4614c4c565b604052919050565b600067ffffffffffffffff821115614cc657614cc6614c4c565b50601f01601f191660200190565b6000614ce7614ce284614cac565b614c7b565b9050828152838383011115614cfb57600080fd5b828260208301376000602084830101529392505050565b600082601f830112614d2357600080fd5b611a6083833560208501614cd4565b60008060008060008060008060e0898b031215614d4e57600080fd5b88359750614d5e60208a01614bf2565b9650604089013567ffffffffffffffff80821115614d7b57600080fd5b614d878c838d01614c03565b909850965060608b0135915080821115614da057600080fd5b50614dad8b828c01614d12565b945050614dbc60808a01614bf2565b925060a0890135915060c089013590509295985092959890939650565b600060208284031215614deb57600080fd5b5035919050565b60005b83811015614e0d578181015183820152602001614df5565b83811115614e1c576000848401525b50505050565b60008151808452614e3a816020860160208601614df2565b601f01601f19169290920160200192915050565b602081526000611a606020830184614e22565b6001600160a01b0381168114610cd957600080fd5b60008060008060808587031215614e8c57600080fd5b8435614e9781614e61565b93506020850135614ea781614e61565b925060408501359150606085013567ffffffffffffffff811115614eca57600080fd5b614ed687828801614d12565b91505092959194509250565b600067ffffffffffffffff821115614efc57614efc614c4c565b5060051b60200190565b600082601f830112614f1757600080fd5b81356020614f27614ce283614ee2565b82815260059290921b84018101918181019086841115614f4657600080fd5b8286015b84811015614f6a578035614f5d81614e61565b8352918301918301614f4a565b509695505050505050565b600082601f830112614f8657600080fd5b81356020614f96614ce283614ee2565b82815260059290921b84018101918181019086841115614fb557600080fd5b8286015b84811015614f6a5780358352918301918301614fb9565b600082601f830112614fe157600080fd5b81356020614ff1614ce283614ee2565b82815260059290921b8401810191818101908684111561501057600080fd5b8286015b84811015614f6a57803567ffffffffffffffff8111156150345760008081fd5b6150428986838b0101614d12565b845250918301918301615014565b6000806000806080858703121561506657600080fd5b843567ffffffffffffffff8082111561507e57600080fd5b61508a88838901614f06565b955060208701359150808211156150a057600080fd5b6150ac88838901614f75565b945060408701359150808211156150c257600080fd5b506150cf87828801614fd0565b949793965093946060013593505050565b600080600080600060a086880312156150f857600080fd5b8535945061510860208701614bf2565b935061511660408701614bf2565b94979396509394606081013594506080013592915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6020810160088310615198577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b600080604083850312156151b157600080fd5b8235915060208301356151c381614e61565b809150509250929050565b600080604083850312156151e157600080fd5b823591506151f160208401614bf2565b90509250929050565b60008060008060006080868803121561521257600080fd5b8535945061522260208701614bf2565b9350604086013567ffffffffffffffff8082111561523f57600080fd5b61524b89838a01614c03565b9095509350606088013591508082111561526457600080fd5b5061527188828901614d12565b9150509295509295909350565b6000806000806060858703121561529457600080fd5b843593506152a460208601614bf2565b9250604085013567ffffffffffffffff8111156152c057600080fd5b6152cc87828801614c03565b95989497509550505050565b600080600080608085870312156152ee57600080fd5b843567ffffffffffffffff8082111561530657600080fd5b61531288838901614f06565b9550602087013591508082111561532857600080fd5b61533488838901614f75565b9450604087013591508082111561534a57600080fd5b61535688838901614fd0565b9350606087013591508082111561536c57600080fd5b508501601f8101871361537e57600080fd5b614ed687823560208401614cd4565b600081518084526020808501945080840160005b838110156153bd578151875295820195908201906001016153a1565b509495945050505050565b7fff000000000000000000000000000000000000000000000000000000000000008816815260e06020820152600061540360e0830189614e22565b82810360408401526154158189614e22565b90508660608401526001600160a01b03861660808401528460a084015282810360c0840152610c14818561538d565b60008060008060008060c0878903121561545d57600080fd5b863561546881614e61565b9550602087013561547881614e61565b95989597505050506040840135936060810135936080820135935060a0909101359150565b6000806000606084860312156154b257600080fd5b83356154bd81614e61565b925060208401359150604084013567ffffffffffffffff8111156154e057600080fd5b6154ec86828701614d12565b9150509250925092565b60006020828403121561550857600080fd5b8135611a6081614e61565b600080600080600060a0868803121561552b57600080fd5b853561553681614e61565b9450602086013561554681614e61565b9350604086013567ffffffffffffffff8082111561556357600080fd5b61556f89838a01614f75565b9450606088013591508082111561558557600080fd5b61559189838a01614f75565b9350608088013591508082111561526457600080fd5b600080600080606085870312156155bd57600080fd5b84356155c881614e61565b935060208501359250604085013567ffffffffffffffff8111156152c057600080fd5b600080604083850312156155fe57600080fd5b823561560981614e61565b946020939093013593505050565b600080600080600060a0868803121561562f57600080fd5b853561563a81614e61565b9450602086013561564a81614e61565b93506040860135925060608601359150608086013567ffffffffffffffff81111561567457600080fd5b61527188828901614d12565b8183823760009101908152919050565b600181811c908216806156a457607f821691505b6020821081036156dd577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b6000602082840312156156f557600080fd5b5051919050565b600081518084526020808501945080840160005b838110156153bd5781516001600160a01b031687529582019590820190600101615710565b600081518084526020808501808196508360051b8101915082860160005b8581101561577d57828403895261576b848351614e22565b98850198935090840190600101615753565b5091979650505050505050565b60a08152600061579d60a08301886156fc565b82810360208401526157af818861538d565b905082810360408401526157c38187615735565b60608401959095525050608001529392505050565b60c0815260006157eb60c08301896156fc565b82810360208401526157fd818961538d565b905082810360408401526158118188615735565b60608401969096525050608081019290925260a0909101529392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000821982111561587257615872615830565b500190565b60006020828403121561588957600080fd5b815167ffffffffffffffff8111156158a057600080fd5b8201601f810184136158b157600080fd5b80516158bf614ce282614cac565b8181528560208385010111156158d457600080fd5b611716826020830160208601614df2565b600084516158f7818460208901614df2565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551615933816001850160208a01614df2565b6001920191820152835161594e816002840160208801614df2565b0160020195945050505050565b60008282101561596d5761596d615830565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000602082840312156159b357600080fd5b815165ffffffffffff81168114611a6057600080fd5b6080815260006159dc60808301876156fc565b82810360208401526159ee818761538d565b90508281036040840152615a028186615735565b91505082606083015295945050505050565b84815260ff8416602082015282604082015260806060820152600061132e6080830184614e22565b85815260ff8516602082015283604082015260a060608201526000615a6460a0830185614e22565b8281036080840152615a768185614e22565b98975050505050505050565b60006000198203615a9557615a95615830565b5060010190565b600060208284031215615aae57600080fd5b81518015158114611a6057600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006101208b835260206001600160a01b038c1681850152816040850152615b178285018c6156fc565b91508382036060850152615b2b828b61538d565b915083820360808501528189518084528284019150828160051b850101838c0160005b83811015615b7c57601f19878403018552615b6a838351614e22565b94860194925090850190600101615b4e565b505086810360a0880152615b90818c615735565b9450505050508560c08401528460e0840152828103610100840152615bb58185614e22565b9c9b505050505050505050505050565b6000816000190483118215151615615bdf57615bdf615830565b500290565b600082615c1a577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b601f821115612ef857600081815260208120601f850160051c81016020861015615c465750805b601f850160051c820191505b8181101561244e57828155600101615c52565b815167ffffffffffffffff811115615c7f57615c7f614c4c565b615c9381615c8d8454615690565b84615c1f565b602080601f831160018114615cc85760008415615cb05750858301515b600019600386901b1c1916600185901b17855561244e565b600085815260208120601f198616915b82811015615cf757888601518255948401946001909101908401615cd8565b5085821015615d155787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fe476f7665726e6f723a2072656c617920726576657274656420776974686f7574206d657373616765a164736f6c634300080f000a",
}

// KromaGovernorABI is the input ABI used to generate the binding from.
// Deprecated: Use KromaGovernorMetaData.ABI instead.
var KromaGovernorABI = KromaGovernorMetaData.ABI

// KromaGovernorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use KromaGovernorMetaData.Bin instead.
var KromaGovernorBin = KromaGovernorMetaData.Bin

// DeployKromaGovernor deploys a new Ethereum contract, binding an instance of KromaGovernor to it.
func DeployKromaGovernor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KromaGovernor, error) {
	parsed, err := KromaGovernorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KromaGovernorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KromaGovernor{KromaGovernorCaller: KromaGovernorCaller{contract: contract}, KromaGovernorTransactor: KromaGovernorTransactor{contract: contract}, KromaGovernorFilterer: KromaGovernorFilterer{contract: contract}}, nil
}

// KromaGovernor is an auto generated Go binding around an Ethereum contract.
type KromaGovernor struct {
	KromaGovernorCaller     // Read-only binding to the contract
	KromaGovernorTransactor // Write-only binding to the contract
	KromaGovernorFilterer   // Log filterer for contract events
}

// KromaGovernorCaller is an auto generated read-only Go binding around an Ethereum contract.
type KromaGovernorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KromaGovernorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KromaGovernorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KromaGovernorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KromaGovernorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KromaGovernorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KromaGovernorSession struct {
	Contract     *KromaGovernor    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KromaGovernorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KromaGovernorCallerSession struct {
	Contract *KromaGovernorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// KromaGovernorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KromaGovernorTransactorSession struct {
	Contract     *KromaGovernorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// KromaGovernorRaw is an auto generated low-level Go binding around an Ethereum contract.
type KromaGovernorRaw struct {
	Contract *KromaGovernor // Generic contract binding to access the raw methods on
}

// KromaGovernorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KromaGovernorCallerRaw struct {
	Contract *KromaGovernorCaller // Generic read-only contract binding to access the raw methods on
}

// KromaGovernorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KromaGovernorTransactorRaw struct {
	Contract *KromaGovernorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKromaGovernor creates a new instance of KromaGovernor, bound to a specific deployed contract.
func NewKromaGovernor(address common.Address, backend bind.ContractBackend) (*KromaGovernor, error) {
	contract, err := bindKromaGovernor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KromaGovernor{KromaGovernorCaller: KromaGovernorCaller{contract: contract}, KromaGovernorTransactor: KromaGovernorTransactor{contract: contract}, KromaGovernorFilterer: KromaGovernorFilterer{contract: contract}}, nil
}

// NewKromaGovernorCaller creates a new read-only instance of KromaGovernor, bound to a specific deployed contract.
func NewKromaGovernorCaller(address common.Address, caller bind.ContractCaller) (*KromaGovernorCaller, error) {
	contract, err := bindKromaGovernor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KromaGovernorCaller{contract: contract}, nil
}

// NewKromaGovernorTransactor creates a new write-only instance of KromaGovernor, bound to a specific deployed contract.
func NewKromaGovernorTransactor(address common.Address, transactor bind.ContractTransactor) (*KromaGovernorTransactor, error) {
	contract, err := bindKromaGovernor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KromaGovernorTransactor{contract: contract}, nil
}

// NewKromaGovernorFilterer creates a new log filterer instance of KromaGovernor, bound to a specific deployed contract.
func NewKromaGovernorFilterer(address common.Address, filterer bind.ContractFilterer) (*KromaGovernorFilterer, error) {
	contract, err := bindKromaGovernor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KromaGovernorFilterer{contract: contract}, nil
}

// bindKromaGovernor binds a generic wrapper to an already deployed contract.
func bindKromaGovernor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KromaGovernorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KromaGovernor *KromaGovernorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KromaGovernor.Contract.KromaGovernorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KromaGovernor *KromaGovernorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KromaGovernor.Contract.KromaGovernorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KromaGovernor *KromaGovernorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KromaGovernor.Contract.KromaGovernorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KromaGovernor *KromaGovernorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KromaGovernor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KromaGovernor *KromaGovernorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KromaGovernor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KromaGovernor *KromaGovernorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KromaGovernor.Contract.contract.Transact(opts, method, params...)
}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_KromaGovernor *KromaGovernorCaller) BALLOTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KromaGovernor.contract.Call(opts, &out, "BALLOT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_KromaGovernor *KromaGovernorSession) BALLOTTYPEHASH() ([32]byte, error) {
	return _KromaGovernor.Contract.BALLOTTYPEHASH(&_KromaGovernor.CallOpts)
}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_KromaGovernor *KromaGovernorCallerSession) BALLOTTYPEHASH() ([32]byte, error) {
	return _KromaGovernor.Contract.BALLOTTYPEHASH(&_KromaGovernor.CallOpts)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_KromaGovernor *KromaGovernorCaller) CLOCKMODE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KromaGovernor.contract.Call(opts, &out, "CLOCK_MODE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_KromaGovernor *KromaGovernorSession) CLOCKMODE() (string, error) {
	return _KromaGovernor.Contract.CLOCKMODE(&_KromaGovernor.CallOpts)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_KromaGovernor *KromaGovernorCallerSession) CLOCKMODE() (string, error) {
	return _KromaGovernor.Contract.CLOCKMODE(&_KromaGovernor.CallOpts)
}

// COUNTINGMODE is a free data retrieval call binding the contract method 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (_KromaGovernor *KromaGovernorCaller) COUNTINGMODE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KromaGovernor.contract.Call(opts, &out, "COUNTING_MODE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// COUNTINGMODE is a free data retrieval call binding the contract method 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (_KromaGovernor *KromaGovernorSession) COUNTINGMODE() (string, error) {
	return _KromaGovernor.Contract.COUNTINGMODE(&_KromaGovernor.CallOpts)
}

// COUNTINGMODE is a free data retrieval call binding the contract method 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (_KromaGovernor *KromaGovernorCallerSession) COUNTINGMODE() (string, error) {
	return _KromaGovernor.Contract.COUNTINGMODE(&_KromaGovernor.CallOpts)
}

// EXTENDEDBALLOTTYPEHASH is a free data retrieval call binding the contract method 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (_KromaGovernor *KromaGovernorCaller) EXTENDEDBALLOTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KromaGovernor.contract.Call(opts, &out, "EXTENDED_BALLOT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EXTENDEDBALLOTTYPEHASH is a free data retrieval call binding the contract method 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (_KromaGovernor *KromaGovernorSession) EXTENDEDBALLOTTYPEHASH() ([32]byte, error) {
	return _KromaGovernor.Contract.EXTENDEDBALLOTTYPEHASH(&_KromaGovernor.CallOpts)
}

// EXTENDEDBALLOTTYPEHASH is a free data retrieval call binding the contract method 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (_KromaGovernor *KromaGovernorCallerSession) EXTENDEDBALLOTTYPEHASH() ([32]byte, error) {
	return _KromaGovernor.Contract.EXTENDEDBALLOTTYPEHASH(&_KromaGovernor.CallOpts)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_KromaGovernor *KromaGovernorCaller) Clock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KromaGovernor.contract.Call(opts, &out, "clock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_KromaGovernor *KromaGovernorSession) Clock() (*big.Int, error) {
	return _KromaGovernor.Contract.Clock(&_KromaGovernor.CallOpts)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_KromaGovernor *KromaGovernorCallerSession) Clock() (*big.Int, error) {
	return _KromaGovernor.Contract.Clock(&_KromaGovernor.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_KromaGovernor *KromaGovernorCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _KromaGovernor.contract.Call(opts, &out, "eip712Domain")

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
func (_KromaGovernor *KromaGovernorSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _KromaGovernor.Contract.Eip712Domain(&_KromaGovernor.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_KromaGovernor *KromaGovernorCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _KromaGovernor.Contract.Eip712Domain(&_KromaGovernor.CallOpts)
}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (_KromaGovernor *KromaGovernorCaller) GetVotes(opts *bind.CallOpts, account common.Address, timepoint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _KromaGovernor.contract.Call(opts, &out, "getVotes", account, timepoint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (_KromaGovernor *KromaGovernorSession) GetVotes(account common.Address, timepoint *big.Int) (*big.Int, error) {
	return _KromaGovernor.Contract.GetVotes(&_KromaGovernor.CallOpts, account, timepoint)
}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (_KromaGovernor *KromaGovernorCallerSession) GetVotes(account common.Address, timepoint *big.Int) (*big.Int, error) {
	return _KromaGovernor.Contract.GetVotes(&_KromaGovernor.CallOpts, account, timepoint)
}

// GetVotesWithParams is a free data retrieval call binding the contract method 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (_KromaGovernor *KromaGovernorCaller) GetVotesWithParams(opts *bind.CallOpts, account common.Address, timepoint *big.Int, params []byte) (*big.Int, error) {
	var out []interface{}
	err := _KromaGovernor.contract.Call(opts, &out, "getVotesWithParams", account, timepoint, params)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotesWithParams is a free data retrieval call binding the contract method 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (_KromaGovernor *KromaGovernorSession) GetVotesWithParams(account common.Address, timepoint *big.Int, params []byte) (*big.Int, error) {
	return _KromaGovernor.Contract.GetVotesWithParams(&_KromaGovernor.CallOpts, account, timepoint, params)
}

// GetVotesWithParams is a free data retrieval call binding the contract method 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (_KromaGovernor *KromaGovernorCallerSession) GetVotesWithParams(account common.Address, timepoint *big.Int, params []byte) (*big.Int, error) {
	return _KromaGovernor.Contract.GetVotesWithParams(&_KromaGovernor.CallOpts, account, timepoint, params)
}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_KromaGovernor *KromaGovernorCaller) HasVoted(opts *bind.CallOpts, proposalId *big.Int, account common.Address) (bool, error) {
	var out []interface{}
	err := _KromaGovernor.contract.Call(opts, &out, "hasVoted", proposalId, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_KromaGovernor *KromaGovernorSession) HasVoted(proposalId *big.Int, account common.Address) (bool, error) {
	return _KromaGovernor.Contract.HasVoted(&_KromaGovernor.CallOpts, proposalId, account)
}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_KromaGovernor *KromaGovernorCallerSession) HasVoted(proposalId *big.Int, account common.Address) (bool, error) {
	return _KromaGovernor.Contract.HasVoted(&_KromaGovernor.CallOpts, proposalId, account)
}

// HashProposal is a free data retrieval call binding the contract method 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (_KromaGovernor *KromaGovernorCaller) HashProposal(opts *bind.CallOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _KromaGovernor.contract.Call(opts, &out, "hashProposal", targets, values, calldatas, descriptionHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashProposal is a free data retrieval call binding the contract method 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (_KromaGovernor *KromaGovernorSession) HashProposal(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	return _KromaGovernor.Contract.HashProposal(&_KromaGovernor.CallOpts, targets, values, calldatas, descriptionHash)
}

// HashProposal is a free data retrieval call binding the contract method 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (_KromaGovernor *KromaGovernorCallerSession) HashProposal(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	return _KromaGovernor.Contract.HashProposal(&_KromaGovernor.CallOpts, targets, values, calldatas, descriptionHash)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KromaGovernor *KromaGovernorCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KromaGovernor.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KromaGovernor *KromaGovernorSession) Name() (string, error) {
	return _KromaGovernor.Contract.Name(&_KromaGovernor.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KromaGovernor *KromaGovernorCallerSession) Name() (string, error) {
	return _KromaGovernor.Contract.Name(&_KromaGovernor.CallOpts)
}

// ProposalDeadline is a free data retrieval call binding the contract method 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (_KromaGovernor *KromaGovernorCaller) ProposalDeadline(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _KromaGovernor.contract.Call(opts, &out, "proposalDeadline", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalDeadline is a free data retrieval call binding the contract method 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (_KromaGovernor *KromaGovernorSession) ProposalDeadline(proposalId *big.Int) (*big.Int, error) {
	return _KromaGovernor.Contract.ProposalDeadline(&_KromaGovernor.CallOpts, proposalId)
}

// ProposalDeadline is a free data retrieval call binding the contract method 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (_KromaGovernor *KromaGovernorCallerSession) ProposalDeadline(proposalId *big.Int) (*big.Int, error) {
	return _KromaGovernor.Contract.ProposalDeadline(&_KromaGovernor.CallOpts, proposalId)
}

// ProposalEta is a free data retrieval call binding the contract method 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (_KromaGovernor *KromaGovernorCaller) ProposalEta(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _KromaGovernor.contract.Call(opts, &out, "proposalEta", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalEta is a free data retrieval call binding the contract method 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (_KromaGovernor *KromaGovernorSession) ProposalEta(proposalId *big.Int) (*big.Int, error) {
	return _KromaGovernor.Contract.ProposalEta(&_KromaGovernor.CallOpts, proposalId)
}

// ProposalEta is a free data retrieval call binding the contract method 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (_KromaGovernor *KromaGovernorCallerSession) ProposalEta(proposalId *big.Int) (*big.Int, error) {
	return _KromaGovernor.Contract.ProposalEta(&_KromaGovernor.CallOpts, proposalId)
}

// ProposalProposer is a free data retrieval call binding the contract method 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (_KromaGovernor *KromaGovernorCaller) ProposalProposer(opts *bind.CallOpts, proposalId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KromaGovernor.contract.Call(opts, &out, "proposalProposer", proposalId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposalProposer is a free data retrieval call binding the contract method 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (_KromaGovernor *KromaGovernorSession) ProposalProposer(proposalId *big.Int) (common.Address, error) {
	return _KromaGovernor.Contract.ProposalProposer(&_KromaGovernor.CallOpts, proposalId)
}

// ProposalProposer is a free data retrieval call binding the contract method 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (_KromaGovernor *KromaGovernorCallerSession) ProposalProposer(proposalId *big.Int) (common.Address, error) {
	return _KromaGovernor.Contract.ProposalProposer(&_KromaGovernor.CallOpts, proposalId)
}

// ProposalSnapshot is a free data retrieval call binding the contract method 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (_KromaGovernor *KromaGovernorCaller) ProposalSnapshot(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _KromaGovernor.contract.Call(opts, &out, "proposalSnapshot", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalSnapshot is a free data retrieval call binding the contract method 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (_KromaGovernor *KromaGovernorSession) ProposalSnapshot(proposalId *big.Int) (*big.Int, error) {
	return _KromaGovernor.Contract.ProposalSnapshot(&_KromaGovernor.CallOpts, proposalId)
}

// ProposalSnapshot is a free data retrieval call binding the contract method 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (_KromaGovernor *KromaGovernorCallerSession) ProposalSnapshot(proposalId *big.Int) (*big.Int, error) {
	return _KromaGovernor.Contract.ProposalSnapshot(&_KromaGovernor.CallOpts, proposalId)
}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() view returns(uint256)
func (_KromaGovernor *KromaGovernorCaller) ProposalThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KromaGovernor.contract.Call(opts, &out, "proposalThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() view returns(uint256)
func (_KromaGovernor *KromaGovernorSession) ProposalThreshold() (*big.Int, error) {
	return _KromaGovernor.Contract.ProposalThreshold(&_KromaGovernor.CallOpts)
}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() view returns(uint256)
func (_KromaGovernor *KromaGovernorCallerSession) ProposalThreshold() (*big.Int, error) {
	return _KromaGovernor.Contract.ProposalThreshold(&_KromaGovernor.CallOpts)
}

// ProposalVotes is a free data retrieval call binding the contract method 0x544ffc9c.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (_KromaGovernor *KromaGovernorCaller) ProposalVotes(opts *bind.CallOpts, proposalId *big.Int) (struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}, error) {
	var out []interface{}
	err := _KromaGovernor.contract.Call(opts, &out, "proposalVotes", proposalId)

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
func (_KromaGovernor *KromaGovernorSession) ProposalVotes(proposalId *big.Int) (struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}, error) {
	return _KromaGovernor.Contract.ProposalVotes(&_KromaGovernor.CallOpts, proposalId)
}

// ProposalVotes is a free data retrieval call binding the contract method 0x544ffc9c.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (_KromaGovernor *KromaGovernorCallerSession) ProposalVotes(proposalId *big.Int) (struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}, error) {
	return _KromaGovernor.Contract.ProposalVotes(&_KromaGovernor.CallOpts, proposalId)
}

// Quorum is a free data retrieval call binding the contract method 0xf8ce560a.
//
// Solidity: function quorum(uint256 blockNumber) view returns(uint256)
func (_KromaGovernor *KromaGovernorCaller) Quorum(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _KromaGovernor.contract.Call(opts, &out, "quorum", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quorum is a free data retrieval call binding the contract method 0xf8ce560a.
//
// Solidity: function quorum(uint256 blockNumber) view returns(uint256)
func (_KromaGovernor *KromaGovernorSession) Quorum(blockNumber *big.Int) (*big.Int, error) {
	return _KromaGovernor.Contract.Quorum(&_KromaGovernor.CallOpts, blockNumber)
}

// Quorum is a free data retrieval call binding the contract method 0xf8ce560a.
//
// Solidity: function quorum(uint256 blockNumber) view returns(uint256)
func (_KromaGovernor *KromaGovernorCallerSession) Quorum(blockNumber *big.Int) (*big.Int, error) {
	return _KromaGovernor.Contract.Quorum(&_KromaGovernor.CallOpts, blockNumber)
}

// QuorumDenominator is a free data retrieval call binding the contract method 0x97c3d334.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (_KromaGovernor *KromaGovernorCaller) QuorumDenominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KromaGovernor.contract.Call(opts, &out, "quorumDenominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumDenominator is a free data retrieval call binding the contract method 0x97c3d334.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (_KromaGovernor *KromaGovernorSession) QuorumDenominator() (*big.Int, error) {
	return _KromaGovernor.Contract.QuorumDenominator(&_KromaGovernor.CallOpts)
}

// QuorumDenominator is a free data retrieval call binding the contract method 0x97c3d334.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (_KromaGovernor *KromaGovernorCallerSession) QuorumDenominator() (*big.Int, error) {
	return _KromaGovernor.Contract.QuorumDenominator(&_KromaGovernor.CallOpts)
}

// QuorumNumerator is a free data retrieval call binding the contract method 0x60c4247f.
//
// Solidity: function quorumNumerator(uint256 timepoint) view returns(uint256)
func (_KromaGovernor *KromaGovernorCaller) QuorumNumerator(opts *bind.CallOpts, timepoint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _KromaGovernor.contract.Call(opts, &out, "quorumNumerator", timepoint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumNumerator is a free data retrieval call binding the contract method 0x60c4247f.
//
// Solidity: function quorumNumerator(uint256 timepoint) view returns(uint256)
func (_KromaGovernor *KromaGovernorSession) QuorumNumerator(timepoint *big.Int) (*big.Int, error) {
	return _KromaGovernor.Contract.QuorumNumerator(&_KromaGovernor.CallOpts, timepoint)
}

// QuorumNumerator is a free data retrieval call binding the contract method 0x60c4247f.
//
// Solidity: function quorumNumerator(uint256 timepoint) view returns(uint256)
func (_KromaGovernor *KromaGovernorCallerSession) QuorumNumerator(timepoint *big.Int) (*big.Int, error) {
	return _KromaGovernor.Contract.QuorumNumerator(&_KromaGovernor.CallOpts, timepoint)
}

// QuorumNumerator0 is a free data retrieval call binding the contract method 0xa7713a70.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (_KromaGovernor *KromaGovernorCaller) QuorumNumerator0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KromaGovernor.contract.Call(opts, &out, "quorumNumerator0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumNumerator0 is a free data retrieval call binding the contract method 0xa7713a70.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (_KromaGovernor *KromaGovernorSession) QuorumNumerator0() (*big.Int, error) {
	return _KromaGovernor.Contract.QuorumNumerator0(&_KromaGovernor.CallOpts)
}

// QuorumNumerator0 is a free data retrieval call binding the contract method 0xa7713a70.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (_KromaGovernor *KromaGovernorCallerSession) QuorumNumerator0() (*big.Int, error) {
	return _KromaGovernor.Contract.QuorumNumerator0(&_KromaGovernor.CallOpts)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_KromaGovernor *KromaGovernorCaller) State(opts *bind.CallOpts, proposalId *big.Int) (uint8, error) {
	var out []interface{}
	err := _KromaGovernor.contract.Call(opts, &out, "state", proposalId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_KromaGovernor *KromaGovernorSession) State(proposalId *big.Int) (uint8, error) {
	return _KromaGovernor.Contract.State(&_KromaGovernor.CallOpts, proposalId)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_KromaGovernor *KromaGovernorCallerSession) State(proposalId *big.Int) (uint8, error) {
	return _KromaGovernor.Contract.State(&_KromaGovernor.CallOpts, proposalId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KromaGovernor *KromaGovernorCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _KromaGovernor.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KromaGovernor *KromaGovernorSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KromaGovernor.Contract.SupportsInterface(&_KromaGovernor.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KromaGovernor *KromaGovernorCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KromaGovernor.Contract.SupportsInterface(&_KromaGovernor.CallOpts, interfaceId)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (_KromaGovernor *KromaGovernorCaller) Timelock(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KromaGovernor.contract.Call(opts, &out, "timelock")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (_KromaGovernor *KromaGovernorSession) Timelock() (common.Address, error) {
	return _KromaGovernor.Contract.Timelock(&_KromaGovernor.CallOpts)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (_KromaGovernor *KromaGovernorCallerSession) Timelock() (common.Address, error) {
	return _KromaGovernor.Contract.Timelock(&_KromaGovernor.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_KromaGovernor *KromaGovernorCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KromaGovernor.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_KromaGovernor *KromaGovernorSession) Token() (common.Address, error) {
	return _KromaGovernor.Contract.Token(&_KromaGovernor.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_KromaGovernor *KromaGovernorCallerSession) Token() (common.Address, error) {
	return _KromaGovernor.Contract.Token(&_KromaGovernor.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_KromaGovernor *KromaGovernorCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KromaGovernor.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_KromaGovernor *KromaGovernorSession) Version() (string, error) {
	return _KromaGovernor.Contract.Version(&_KromaGovernor.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_KromaGovernor *KromaGovernorCallerSession) Version() (string, error) {
	return _KromaGovernor.Contract.Version(&_KromaGovernor.CallOpts)
}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() view returns(uint256)
func (_KromaGovernor *KromaGovernorCaller) VotingDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KromaGovernor.contract.Call(opts, &out, "votingDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() view returns(uint256)
func (_KromaGovernor *KromaGovernorSession) VotingDelay() (*big.Int, error) {
	return _KromaGovernor.Contract.VotingDelay(&_KromaGovernor.CallOpts)
}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() view returns(uint256)
func (_KromaGovernor *KromaGovernorCallerSession) VotingDelay() (*big.Int, error) {
	return _KromaGovernor.Contract.VotingDelay(&_KromaGovernor.CallOpts)
}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
func (_KromaGovernor *KromaGovernorCaller) VotingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KromaGovernor.contract.Call(opts, &out, "votingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
func (_KromaGovernor *KromaGovernorSession) VotingPeriod() (*big.Int, error) {
	return _KromaGovernor.Contract.VotingPeriod(&_KromaGovernor.CallOpts)
}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
func (_KromaGovernor *KromaGovernorCallerSession) VotingPeriod() (*big.Int, error) {
	return _KromaGovernor.Contract.VotingPeriod(&_KromaGovernor.CallOpts)
}

// Cancel is a paid mutator transaction binding the contract method 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_KromaGovernor *KromaGovernorTransactor) Cancel(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _KromaGovernor.contract.Transact(opts, "cancel", targets, values, calldatas, descriptionHash)
}

// Cancel is a paid mutator transaction binding the contract method 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_KromaGovernor *KromaGovernorSession) Cancel(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _KromaGovernor.Contract.Cancel(&_KromaGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Cancel is a paid mutator transaction binding the contract method 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_KromaGovernor *KromaGovernorTransactorSession) Cancel(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _KromaGovernor.Contract.Cancel(&_KromaGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (_KromaGovernor *KromaGovernorTransactor) CastVote(opts *bind.TransactOpts, proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _KromaGovernor.contract.Transact(opts, "castVote", proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (_KromaGovernor *KromaGovernorSession) CastVote(proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _KromaGovernor.Contract.CastVote(&_KromaGovernor.TransactOpts, proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (_KromaGovernor *KromaGovernorTransactorSession) CastVote(proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _KromaGovernor.Contract.CastVote(&_KromaGovernor.TransactOpts, proposalId, support)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x3bccf4fd.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, uint8 v, bytes32 r, bytes32 s) returns(uint256)
func (_KromaGovernor *KromaGovernorTransactor) CastVoteBySig(opts *bind.TransactOpts, proposalId *big.Int, support uint8, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _KromaGovernor.contract.Transact(opts, "castVoteBySig", proposalId, support, v, r, s)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x3bccf4fd.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, uint8 v, bytes32 r, bytes32 s) returns(uint256)
func (_KromaGovernor *KromaGovernorSession) CastVoteBySig(proposalId *big.Int, support uint8, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _KromaGovernor.Contract.CastVoteBySig(&_KromaGovernor.TransactOpts, proposalId, support, v, r, s)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x3bccf4fd.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, uint8 v, bytes32 r, bytes32 s) returns(uint256)
func (_KromaGovernor *KromaGovernorTransactorSession) CastVoteBySig(proposalId *big.Int, support uint8, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _KromaGovernor.Contract.CastVoteBySig(&_KromaGovernor.TransactOpts, proposalId, support, v, r, s)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (_KromaGovernor *KromaGovernorTransactor) CastVoteWithReason(opts *bind.TransactOpts, proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _KromaGovernor.contract.Transact(opts, "castVoteWithReason", proposalId, support, reason)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (_KromaGovernor *KromaGovernorSession) CastVoteWithReason(proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _KromaGovernor.Contract.CastVoteWithReason(&_KromaGovernor.TransactOpts, proposalId, support, reason)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (_KromaGovernor *KromaGovernorTransactorSession) CastVoteWithReason(proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _KromaGovernor.Contract.CastVoteWithReason(&_KromaGovernor.TransactOpts, proposalId, support, reason)
}

// CastVoteWithReasonAndParams is a paid mutator transaction binding the contract method 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (_KromaGovernor *KromaGovernorTransactor) CastVoteWithReasonAndParams(opts *bind.TransactOpts, proposalId *big.Int, support uint8, reason string, params []byte) (*types.Transaction, error) {
	return _KromaGovernor.contract.Transact(opts, "castVoteWithReasonAndParams", proposalId, support, reason, params)
}

// CastVoteWithReasonAndParams is a paid mutator transaction binding the contract method 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (_KromaGovernor *KromaGovernorSession) CastVoteWithReasonAndParams(proposalId *big.Int, support uint8, reason string, params []byte) (*types.Transaction, error) {
	return _KromaGovernor.Contract.CastVoteWithReasonAndParams(&_KromaGovernor.TransactOpts, proposalId, support, reason, params)
}

// CastVoteWithReasonAndParams is a paid mutator transaction binding the contract method 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (_KromaGovernor *KromaGovernorTransactorSession) CastVoteWithReasonAndParams(proposalId *big.Int, support uint8, reason string, params []byte) (*types.Transaction, error) {
	return _KromaGovernor.Contract.CastVoteWithReasonAndParams(&_KromaGovernor.TransactOpts, proposalId, support, reason, params)
}

// CastVoteWithReasonAndParamsBySig is a paid mutator transaction binding the contract method 0x03420181.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, string reason, bytes params, uint8 v, bytes32 r, bytes32 s) returns(uint256)
func (_KromaGovernor *KromaGovernorTransactor) CastVoteWithReasonAndParamsBySig(opts *bind.TransactOpts, proposalId *big.Int, support uint8, reason string, params []byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _KromaGovernor.contract.Transact(opts, "castVoteWithReasonAndParamsBySig", proposalId, support, reason, params, v, r, s)
}

// CastVoteWithReasonAndParamsBySig is a paid mutator transaction binding the contract method 0x03420181.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, string reason, bytes params, uint8 v, bytes32 r, bytes32 s) returns(uint256)
func (_KromaGovernor *KromaGovernorSession) CastVoteWithReasonAndParamsBySig(proposalId *big.Int, support uint8, reason string, params []byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _KromaGovernor.Contract.CastVoteWithReasonAndParamsBySig(&_KromaGovernor.TransactOpts, proposalId, support, reason, params, v, r, s)
}

// CastVoteWithReasonAndParamsBySig is a paid mutator transaction binding the contract method 0x03420181.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, string reason, bytes params, uint8 v, bytes32 r, bytes32 s) returns(uint256)
func (_KromaGovernor *KromaGovernorTransactorSession) CastVoteWithReasonAndParamsBySig(proposalId *big.Int, support uint8, reason string, params []byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _KromaGovernor.Contract.CastVoteWithReasonAndParamsBySig(&_KromaGovernor.TransactOpts, proposalId, support, reason, params, v, r, s)
}

// Execute is a paid mutator transaction binding the contract method 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (_KromaGovernor *KromaGovernorTransactor) Execute(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _KromaGovernor.contract.Transact(opts, "execute", targets, values, calldatas, descriptionHash)
}

// Execute is a paid mutator transaction binding the contract method 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (_KromaGovernor *KromaGovernorSession) Execute(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _KromaGovernor.Contract.Execute(&_KromaGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Execute is a paid mutator transaction binding the contract method 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (_KromaGovernor *KromaGovernorTransactorSession) Execute(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _KromaGovernor.Contract.Execute(&_KromaGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Initialize is a paid mutator transaction binding the contract method 0x86489ba9.
//
// Solidity: function initialize(address _token, address _timelock, uint256 _initialVotingDelay, uint256 _initialVotingPeriod, uint256 _initialProposalThreshold, uint256 _votesQuorumFraction) returns()
func (_KromaGovernor *KromaGovernorTransactor) Initialize(opts *bind.TransactOpts, _token common.Address, _timelock common.Address, _initialVotingDelay *big.Int, _initialVotingPeriod *big.Int, _initialProposalThreshold *big.Int, _votesQuorumFraction *big.Int) (*types.Transaction, error) {
	return _KromaGovernor.contract.Transact(opts, "initialize", _token, _timelock, _initialVotingDelay, _initialVotingPeriod, _initialProposalThreshold, _votesQuorumFraction)
}

// Initialize is a paid mutator transaction binding the contract method 0x86489ba9.
//
// Solidity: function initialize(address _token, address _timelock, uint256 _initialVotingDelay, uint256 _initialVotingPeriod, uint256 _initialProposalThreshold, uint256 _votesQuorumFraction) returns()
func (_KromaGovernor *KromaGovernorSession) Initialize(_token common.Address, _timelock common.Address, _initialVotingDelay *big.Int, _initialVotingPeriod *big.Int, _initialProposalThreshold *big.Int, _votesQuorumFraction *big.Int) (*types.Transaction, error) {
	return _KromaGovernor.Contract.Initialize(&_KromaGovernor.TransactOpts, _token, _timelock, _initialVotingDelay, _initialVotingPeriod, _initialProposalThreshold, _votesQuorumFraction)
}

// Initialize is a paid mutator transaction binding the contract method 0x86489ba9.
//
// Solidity: function initialize(address _token, address _timelock, uint256 _initialVotingDelay, uint256 _initialVotingPeriod, uint256 _initialProposalThreshold, uint256 _votesQuorumFraction) returns()
func (_KromaGovernor *KromaGovernorTransactorSession) Initialize(_token common.Address, _timelock common.Address, _initialVotingDelay *big.Int, _initialVotingPeriod *big.Int, _initialProposalThreshold *big.Int, _votesQuorumFraction *big.Int) (*types.Transaction, error) {
	return _KromaGovernor.Contract.Initialize(&_KromaGovernor.TransactOpts, _token, _timelock, _initialVotingDelay, _initialVotingPeriod, _initialProposalThreshold, _votesQuorumFraction)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_KromaGovernor *KromaGovernorTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _KromaGovernor.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_KromaGovernor *KromaGovernorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _KromaGovernor.Contract.OnERC1155BatchReceived(&_KromaGovernor.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_KromaGovernor *KromaGovernorTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _KromaGovernor.Contract.OnERC1155BatchReceived(&_KromaGovernor.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_KromaGovernor *KromaGovernorTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _KromaGovernor.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_KromaGovernor *KromaGovernorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _KromaGovernor.Contract.OnERC1155Received(&_KromaGovernor.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_KromaGovernor *KromaGovernorTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _KromaGovernor.Contract.OnERC1155Received(&_KromaGovernor.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_KromaGovernor *KromaGovernorTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _KromaGovernor.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_KromaGovernor *KromaGovernorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _KromaGovernor.Contract.OnERC721Received(&_KromaGovernor.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_KromaGovernor *KromaGovernorTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _KromaGovernor.Contract.OnERC721Received(&_KromaGovernor.TransactOpts, arg0, arg1, arg2, arg3)
}

// Propose is a paid mutator transaction binding the contract method 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (_KromaGovernor *KromaGovernorTransactor) Propose(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _KromaGovernor.contract.Transact(opts, "propose", targets, values, calldatas, description)
}

// Propose is a paid mutator transaction binding the contract method 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (_KromaGovernor *KromaGovernorSession) Propose(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _KromaGovernor.Contract.Propose(&_KromaGovernor.TransactOpts, targets, values, calldatas, description)
}

// Propose is a paid mutator transaction binding the contract method 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (_KromaGovernor *KromaGovernorTransactorSession) Propose(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _KromaGovernor.Contract.Propose(&_KromaGovernor.TransactOpts, targets, values, calldatas, description)
}

// Queue is a paid mutator transaction binding the contract method 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_KromaGovernor *KromaGovernorTransactor) Queue(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _KromaGovernor.contract.Transact(opts, "queue", targets, values, calldatas, descriptionHash)
}

// Queue is a paid mutator transaction binding the contract method 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_KromaGovernor *KromaGovernorSession) Queue(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _KromaGovernor.Contract.Queue(&_KromaGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Queue is a paid mutator transaction binding the contract method 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_KromaGovernor *KromaGovernorTransactorSession) Queue(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _KromaGovernor.Contract.Queue(&_KromaGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Relay is a paid mutator transaction binding the contract method 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (_KromaGovernor *KromaGovernorTransactor) Relay(opts *bind.TransactOpts, target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _KromaGovernor.contract.Transact(opts, "relay", target, value, data)
}

// Relay is a paid mutator transaction binding the contract method 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (_KromaGovernor *KromaGovernorSession) Relay(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _KromaGovernor.Contract.Relay(&_KromaGovernor.TransactOpts, target, value, data)
}

// Relay is a paid mutator transaction binding the contract method 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (_KromaGovernor *KromaGovernorTransactorSession) Relay(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _KromaGovernor.Contract.Relay(&_KromaGovernor.TransactOpts, target, value, data)
}

// SetProposalThreshold is a paid mutator transaction binding the contract method 0xece40cc1.
//
// Solidity: function setProposalThreshold(uint256 newProposalThreshold) returns()
func (_KromaGovernor *KromaGovernorTransactor) SetProposalThreshold(opts *bind.TransactOpts, newProposalThreshold *big.Int) (*types.Transaction, error) {
	return _KromaGovernor.contract.Transact(opts, "setProposalThreshold", newProposalThreshold)
}

// SetProposalThreshold is a paid mutator transaction binding the contract method 0xece40cc1.
//
// Solidity: function setProposalThreshold(uint256 newProposalThreshold) returns()
func (_KromaGovernor *KromaGovernorSession) SetProposalThreshold(newProposalThreshold *big.Int) (*types.Transaction, error) {
	return _KromaGovernor.Contract.SetProposalThreshold(&_KromaGovernor.TransactOpts, newProposalThreshold)
}

// SetProposalThreshold is a paid mutator transaction binding the contract method 0xece40cc1.
//
// Solidity: function setProposalThreshold(uint256 newProposalThreshold) returns()
func (_KromaGovernor *KromaGovernorTransactorSession) SetProposalThreshold(newProposalThreshold *big.Int) (*types.Transaction, error) {
	return _KromaGovernor.Contract.SetProposalThreshold(&_KromaGovernor.TransactOpts, newProposalThreshold)
}

// SetVotingDelay is a paid mutator transaction binding the contract method 0x70b0f660.
//
// Solidity: function setVotingDelay(uint256 newVotingDelay) returns()
func (_KromaGovernor *KromaGovernorTransactor) SetVotingDelay(opts *bind.TransactOpts, newVotingDelay *big.Int) (*types.Transaction, error) {
	return _KromaGovernor.contract.Transact(opts, "setVotingDelay", newVotingDelay)
}

// SetVotingDelay is a paid mutator transaction binding the contract method 0x70b0f660.
//
// Solidity: function setVotingDelay(uint256 newVotingDelay) returns()
func (_KromaGovernor *KromaGovernorSession) SetVotingDelay(newVotingDelay *big.Int) (*types.Transaction, error) {
	return _KromaGovernor.Contract.SetVotingDelay(&_KromaGovernor.TransactOpts, newVotingDelay)
}

// SetVotingDelay is a paid mutator transaction binding the contract method 0x70b0f660.
//
// Solidity: function setVotingDelay(uint256 newVotingDelay) returns()
func (_KromaGovernor *KromaGovernorTransactorSession) SetVotingDelay(newVotingDelay *big.Int) (*types.Transaction, error) {
	return _KromaGovernor.Contract.SetVotingDelay(&_KromaGovernor.TransactOpts, newVotingDelay)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xea0217cf.
//
// Solidity: function setVotingPeriod(uint256 newVotingPeriod) returns()
func (_KromaGovernor *KromaGovernorTransactor) SetVotingPeriod(opts *bind.TransactOpts, newVotingPeriod *big.Int) (*types.Transaction, error) {
	return _KromaGovernor.contract.Transact(opts, "setVotingPeriod", newVotingPeriod)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xea0217cf.
//
// Solidity: function setVotingPeriod(uint256 newVotingPeriod) returns()
func (_KromaGovernor *KromaGovernorSession) SetVotingPeriod(newVotingPeriod *big.Int) (*types.Transaction, error) {
	return _KromaGovernor.Contract.SetVotingPeriod(&_KromaGovernor.TransactOpts, newVotingPeriod)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xea0217cf.
//
// Solidity: function setVotingPeriod(uint256 newVotingPeriod) returns()
func (_KromaGovernor *KromaGovernorTransactorSession) SetVotingPeriod(newVotingPeriod *big.Int) (*types.Transaction, error) {
	return _KromaGovernor.Contract.SetVotingPeriod(&_KromaGovernor.TransactOpts, newVotingPeriod)
}

// UpdateQuorumNumerator is a paid mutator transaction binding the contract method 0x06f3f9e6.
//
// Solidity: function updateQuorumNumerator(uint256 newQuorumNumerator) returns()
func (_KromaGovernor *KromaGovernorTransactor) UpdateQuorumNumerator(opts *bind.TransactOpts, newQuorumNumerator *big.Int) (*types.Transaction, error) {
	return _KromaGovernor.contract.Transact(opts, "updateQuorumNumerator", newQuorumNumerator)
}

// UpdateQuorumNumerator is a paid mutator transaction binding the contract method 0x06f3f9e6.
//
// Solidity: function updateQuorumNumerator(uint256 newQuorumNumerator) returns()
func (_KromaGovernor *KromaGovernorSession) UpdateQuorumNumerator(newQuorumNumerator *big.Int) (*types.Transaction, error) {
	return _KromaGovernor.Contract.UpdateQuorumNumerator(&_KromaGovernor.TransactOpts, newQuorumNumerator)
}

// UpdateQuorumNumerator is a paid mutator transaction binding the contract method 0x06f3f9e6.
//
// Solidity: function updateQuorumNumerator(uint256 newQuorumNumerator) returns()
func (_KromaGovernor *KromaGovernorTransactorSession) UpdateQuorumNumerator(newQuorumNumerator *big.Int) (*types.Transaction, error) {
	return _KromaGovernor.Contract.UpdateQuorumNumerator(&_KromaGovernor.TransactOpts, newQuorumNumerator)
}

// UpdateTimelock is a paid mutator transaction binding the contract method 0xa890c910.
//
// Solidity: function updateTimelock(address newTimelock) returns()
func (_KromaGovernor *KromaGovernorTransactor) UpdateTimelock(opts *bind.TransactOpts, newTimelock common.Address) (*types.Transaction, error) {
	return _KromaGovernor.contract.Transact(opts, "updateTimelock", newTimelock)
}

// UpdateTimelock is a paid mutator transaction binding the contract method 0xa890c910.
//
// Solidity: function updateTimelock(address newTimelock) returns()
func (_KromaGovernor *KromaGovernorSession) UpdateTimelock(newTimelock common.Address) (*types.Transaction, error) {
	return _KromaGovernor.Contract.UpdateTimelock(&_KromaGovernor.TransactOpts, newTimelock)
}

// UpdateTimelock is a paid mutator transaction binding the contract method 0xa890c910.
//
// Solidity: function updateTimelock(address newTimelock) returns()
func (_KromaGovernor *KromaGovernorTransactorSession) UpdateTimelock(newTimelock common.Address) (*types.Transaction, error) {
	return _KromaGovernor.Contract.UpdateTimelock(&_KromaGovernor.TransactOpts, newTimelock)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_KromaGovernor *KromaGovernorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KromaGovernor.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_KromaGovernor *KromaGovernorSession) Receive() (*types.Transaction, error) {
	return _KromaGovernor.Contract.Receive(&_KromaGovernor.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_KromaGovernor *KromaGovernorTransactorSession) Receive() (*types.Transaction, error) {
	return _KromaGovernor.Contract.Receive(&_KromaGovernor.TransactOpts)
}

// KromaGovernorEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the KromaGovernor contract.
type KromaGovernorEIP712DomainChangedIterator struct {
	Event *KromaGovernorEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *KromaGovernorEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KromaGovernorEIP712DomainChanged)
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
		it.Event = new(KromaGovernorEIP712DomainChanged)
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
func (it *KromaGovernorEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KromaGovernorEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KromaGovernorEIP712DomainChanged represents a EIP712DomainChanged event raised by the KromaGovernor contract.
type KromaGovernorEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_KromaGovernor *KromaGovernorFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*KromaGovernorEIP712DomainChangedIterator, error) {

	logs, sub, err := _KromaGovernor.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &KromaGovernorEIP712DomainChangedIterator{contract: _KromaGovernor.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_KromaGovernor *KromaGovernorFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *KromaGovernorEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _KromaGovernor.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KromaGovernorEIP712DomainChanged)
				if err := _KromaGovernor.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_KromaGovernor *KromaGovernorFilterer) ParseEIP712DomainChanged(log types.Log) (*KromaGovernorEIP712DomainChanged, error) {
	event := new(KromaGovernorEIP712DomainChanged)
	if err := _KromaGovernor.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KromaGovernorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the KromaGovernor contract.
type KromaGovernorInitializedIterator struct {
	Event *KromaGovernorInitialized // Event containing the contract specifics and raw log

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
func (it *KromaGovernorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KromaGovernorInitialized)
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
		it.Event = new(KromaGovernorInitialized)
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
func (it *KromaGovernorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KromaGovernorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KromaGovernorInitialized represents a Initialized event raised by the KromaGovernor contract.
type KromaGovernorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_KromaGovernor *KromaGovernorFilterer) FilterInitialized(opts *bind.FilterOpts) (*KromaGovernorInitializedIterator, error) {

	logs, sub, err := _KromaGovernor.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &KromaGovernorInitializedIterator{contract: _KromaGovernor.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_KromaGovernor *KromaGovernorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *KromaGovernorInitialized) (event.Subscription, error) {

	logs, sub, err := _KromaGovernor.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KromaGovernorInitialized)
				if err := _KromaGovernor.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_KromaGovernor *KromaGovernorFilterer) ParseInitialized(log types.Log) (*KromaGovernorInitialized, error) {
	event := new(KromaGovernorInitialized)
	if err := _KromaGovernor.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KromaGovernorProposalCanceledIterator is returned from FilterProposalCanceled and is used to iterate over the raw logs and unpacked data for ProposalCanceled events raised by the KromaGovernor contract.
type KromaGovernorProposalCanceledIterator struct {
	Event *KromaGovernorProposalCanceled // Event containing the contract specifics and raw log

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
func (it *KromaGovernorProposalCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KromaGovernorProposalCanceled)
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
		it.Event = new(KromaGovernorProposalCanceled)
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
func (it *KromaGovernorProposalCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KromaGovernorProposalCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KromaGovernorProposalCanceled represents a ProposalCanceled event raised by the KromaGovernor contract.
type KromaGovernorProposalCanceled struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalCanceled is a free log retrieval operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_KromaGovernor *KromaGovernorFilterer) FilterProposalCanceled(opts *bind.FilterOpts) (*KromaGovernorProposalCanceledIterator, error) {

	logs, sub, err := _KromaGovernor.contract.FilterLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return &KromaGovernorProposalCanceledIterator{contract: _KromaGovernor.contract, event: "ProposalCanceled", logs: logs, sub: sub}, nil
}

// WatchProposalCanceled is a free log subscription operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_KromaGovernor *KromaGovernorFilterer) WatchProposalCanceled(opts *bind.WatchOpts, sink chan<- *KromaGovernorProposalCanceled) (event.Subscription, error) {

	logs, sub, err := _KromaGovernor.contract.WatchLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KromaGovernorProposalCanceled)
				if err := _KromaGovernor.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
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
func (_KromaGovernor *KromaGovernorFilterer) ParseProposalCanceled(log types.Log) (*KromaGovernorProposalCanceled, error) {
	event := new(KromaGovernorProposalCanceled)
	if err := _KromaGovernor.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KromaGovernorProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the KromaGovernor contract.
type KromaGovernorProposalCreatedIterator struct {
	Event *KromaGovernorProposalCreated // Event containing the contract specifics and raw log

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
func (it *KromaGovernorProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KromaGovernorProposalCreated)
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
		it.Event = new(KromaGovernorProposalCreated)
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
func (it *KromaGovernorProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KromaGovernorProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KromaGovernorProposalCreated represents a ProposalCreated event raised by the KromaGovernor contract.
type KromaGovernorProposalCreated struct {
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
func (_KromaGovernor *KromaGovernorFilterer) FilterProposalCreated(opts *bind.FilterOpts) (*KromaGovernorProposalCreatedIterator, error) {

	logs, sub, err := _KromaGovernor.contract.FilterLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return &KromaGovernorProposalCreatedIterator{contract: _KromaGovernor.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 voteStart, uint256 voteEnd, string description)
func (_KromaGovernor *KromaGovernorFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *KromaGovernorProposalCreated) (event.Subscription, error) {

	logs, sub, err := _KromaGovernor.contract.WatchLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KromaGovernorProposalCreated)
				if err := _KromaGovernor.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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
func (_KromaGovernor *KromaGovernorFilterer) ParseProposalCreated(log types.Log) (*KromaGovernorProposalCreated, error) {
	event := new(KromaGovernorProposalCreated)
	if err := _KromaGovernor.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KromaGovernorProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the KromaGovernor contract.
type KromaGovernorProposalExecutedIterator struct {
	Event *KromaGovernorProposalExecuted // Event containing the contract specifics and raw log

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
func (it *KromaGovernorProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KromaGovernorProposalExecuted)
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
		it.Event = new(KromaGovernorProposalExecuted)
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
func (it *KromaGovernorProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KromaGovernorProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KromaGovernorProposalExecuted represents a ProposalExecuted event raised by the KromaGovernor contract.
type KromaGovernorProposalExecuted struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_KromaGovernor *KromaGovernorFilterer) FilterProposalExecuted(opts *bind.FilterOpts) (*KromaGovernorProposalExecutedIterator, error) {

	logs, sub, err := _KromaGovernor.contract.FilterLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return &KromaGovernorProposalExecutedIterator{contract: _KromaGovernor.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_KromaGovernor *KromaGovernorFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *KromaGovernorProposalExecuted) (event.Subscription, error) {

	logs, sub, err := _KromaGovernor.contract.WatchLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KromaGovernorProposalExecuted)
				if err := _KromaGovernor.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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
func (_KromaGovernor *KromaGovernorFilterer) ParseProposalExecuted(log types.Log) (*KromaGovernorProposalExecuted, error) {
	event := new(KromaGovernorProposalExecuted)
	if err := _KromaGovernor.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KromaGovernorProposalQueuedIterator is returned from FilterProposalQueued and is used to iterate over the raw logs and unpacked data for ProposalQueued events raised by the KromaGovernor contract.
type KromaGovernorProposalQueuedIterator struct {
	Event *KromaGovernorProposalQueued // Event containing the contract specifics and raw log

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
func (it *KromaGovernorProposalQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KromaGovernorProposalQueued)
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
		it.Event = new(KromaGovernorProposalQueued)
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
func (it *KromaGovernorProposalQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KromaGovernorProposalQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KromaGovernorProposalQueued represents a ProposalQueued event raised by the KromaGovernor contract.
type KromaGovernorProposalQueued struct {
	ProposalId *big.Int
	Eta        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalQueued is a free log retrieval operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 eta)
func (_KromaGovernor *KromaGovernorFilterer) FilterProposalQueued(opts *bind.FilterOpts) (*KromaGovernorProposalQueuedIterator, error) {

	logs, sub, err := _KromaGovernor.contract.FilterLogs(opts, "ProposalQueued")
	if err != nil {
		return nil, err
	}
	return &KromaGovernorProposalQueuedIterator{contract: _KromaGovernor.contract, event: "ProposalQueued", logs: logs, sub: sub}, nil
}

// WatchProposalQueued is a free log subscription operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 eta)
func (_KromaGovernor *KromaGovernorFilterer) WatchProposalQueued(opts *bind.WatchOpts, sink chan<- *KromaGovernorProposalQueued) (event.Subscription, error) {

	logs, sub, err := _KromaGovernor.contract.WatchLogs(opts, "ProposalQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KromaGovernorProposalQueued)
				if err := _KromaGovernor.contract.UnpackLog(event, "ProposalQueued", log); err != nil {
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
func (_KromaGovernor *KromaGovernorFilterer) ParseProposalQueued(log types.Log) (*KromaGovernorProposalQueued, error) {
	event := new(KromaGovernorProposalQueued)
	if err := _KromaGovernor.contract.UnpackLog(event, "ProposalQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KromaGovernorProposalThresholdSetIterator is returned from FilterProposalThresholdSet and is used to iterate over the raw logs and unpacked data for ProposalThresholdSet events raised by the KromaGovernor contract.
type KromaGovernorProposalThresholdSetIterator struct {
	Event *KromaGovernorProposalThresholdSet // Event containing the contract specifics and raw log

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
func (it *KromaGovernorProposalThresholdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KromaGovernorProposalThresholdSet)
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
		it.Event = new(KromaGovernorProposalThresholdSet)
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
func (it *KromaGovernorProposalThresholdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KromaGovernorProposalThresholdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KromaGovernorProposalThresholdSet represents a ProposalThresholdSet event raised by the KromaGovernor contract.
type KromaGovernorProposalThresholdSet struct {
	OldProposalThreshold *big.Int
	NewProposalThreshold *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterProposalThresholdSet is a free log retrieval operation binding the contract event 0xccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc05461.
//
// Solidity: event ProposalThresholdSet(uint256 oldProposalThreshold, uint256 newProposalThreshold)
func (_KromaGovernor *KromaGovernorFilterer) FilterProposalThresholdSet(opts *bind.FilterOpts) (*KromaGovernorProposalThresholdSetIterator, error) {

	logs, sub, err := _KromaGovernor.contract.FilterLogs(opts, "ProposalThresholdSet")
	if err != nil {
		return nil, err
	}
	return &KromaGovernorProposalThresholdSetIterator{contract: _KromaGovernor.contract, event: "ProposalThresholdSet", logs: logs, sub: sub}, nil
}

// WatchProposalThresholdSet is a free log subscription operation binding the contract event 0xccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc05461.
//
// Solidity: event ProposalThresholdSet(uint256 oldProposalThreshold, uint256 newProposalThreshold)
func (_KromaGovernor *KromaGovernorFilterer) WatchProposalThresholdSet(opts *bind.WatchOpts, sink chan<- *KromaGovernorProposalThresholdSet) (event.Subscription, error) {

	logs, sub, err := _KromaGovernor.contract.WatchLogs(opts, "ProposalThresholdSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KromaGovernorProposalThresholdSet)
				if err := _KromaGovernor.contract.UnpackLog(event, "ProposalThresholdSet", log); err != nil {
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
func (_KromaGovernor *KromaGovernorFilterer) ParseProposalThresholdSet(log types.Log) (*KromaGovernorProposalThresholdSet, error) {
	event := new(KromaGovernorProposalThresholdSet)
	if err := _KromaGovernor.contract.UnpackLog(event, "ProposalThresholdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KromaGovernorQuorumNumeratorUpdatedIterator is returned from FilterQuorumNumeratorUpdated and is used to iterate over the raw logs and unpacked data for QuorumNumeratorUpdated events raised by the KromaGovernor contract.
type KromaGovernorQuorumNumeratorUpdatedIterator struct {
	Event *KromaGovernorQuorumNumeratorUpdated // Event containing the contract specifics and raw log

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
func (it *KromaGovernorQuorumNumeratorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KromaGovernorQuorumNumeratorUpdated)
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
		it.Event = new(KromaGovernorQuorumNumeratorUpdated)
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
func (it *KromaGovernorQuorumNumeratorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KromaGovernorQuorumNumeratorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KromaGovernorQuorumNumeratorUpdated represents a QuorumNumeratorUpdated event raised by the KromaGovernor contract.
type KromaGovernorQuorumNumeratorUpdated struct {
	OldQuorumNumerator *big.Int
	NewQuorumNumerator *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterQuorumNumeratorUpdated is a free log retrieval operation binding the contract event 0x0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997.
//
// Solidity: event QuorumNumeratorUpdated(uint256 oldQuorumNumerator, uint256 newQuorumNumerator)
func (_KromaGovernor *KromaGovernorFilterer) FilterQuorumNumeratorUpdated(opts *bind.FilterOpts) (*KromaGovernorQuorumNumeratorUpdatedIterator, error) {

	logs, sub, err := _KromaGovernor.contract.FilterLogs(opts, "QuorumNumeratorUpdated")
	if err != nil {
		return nil, err
	}
	return &KromaGovernorQuorumNumeratorUpdatedIterator{contract: _KromaGovernor.contract, event: "QuorumNumeratorUpdated", logs: logs, sub: sub}, nil
}

// WatchQuorumNumeratorUpdated is a free log subscription operation binding the contract event 0x0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997.
//
// Solidity: event QuorumNumeratorUpdated(uint256 oldQuorumNumerator, uint256 newQuorumNumerator)
func (_KromaGovernor *KromaGovernorFilterer) WatchQuorumNumeratorUpdated(opts *bind.WatchOpts, sink chan<- *KromaGovernorQuorumNumeratorUpdated) (event.Subscription, error) {

	logs, sub, err := _KromaGovernor.contract.WatchLogs(opts, "QuorumNumeratorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KromaGovernorQuorumNumeratorUpdated)
				if err := _KromaGovernor.contract.UnpackLog(event, "QuorumNumeratorUpdated", log); err != nil {
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
func (_KromaGovernor *KromaGovernorFilterer) ParseQuorumNumeratorUpdated(log types.Log) (*KromaGovernorQuorumNumeratorUpdated, error) {
	event := new(KromaGovernorQuorumNumeratorUpdated)
	if err := _KromaGovernor.contract.UnpackLog(event, "QuorumNumeratorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KromaGovernorTimelockChangeIterator is returned from FilterTimelockChange and is used to iterate over the raw logs and unpacked data for TimelockChange events raised by the KromaGovernor contract.
type KromaGovernorTimelockChangeIterator struct {
	Event *KromaGovernorTimelockChange // Event containing the contract specifics and raw log

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
func (it *KromaGovernorTimelockChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KromaGovernorTimelockChange)
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
		it.Event = new(KromaGovernorTimelockChange)
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
func (it *KromaGovernorTimelockChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KromaGovernorTimelockChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KromaGovernorTimelockChange represents a TimelockChange event raised by the KromaGovernor contract.
type KromaGovernorTimelockChange struct {
	OldTimelock common.Address
	NewTimelock common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTimelockChange is a free log retrieval operation binding the contract event 0x08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401.
//
// Solidity: event TimelockChange(address oldTimelock, address newTimelock)
func (_KromaGovernor *KromaGovernorFilterer) FilterTimelockChange(opts *bind.FilterOpts) (*KromaGovernorTimelockChangeIterator, error) {

	logs, sub, err := _KromaGovernor.contract.FilterLogs(opts, "TimelockChange")
	if err != nil {
		return nil, err
	}
	return &KromaGovernorTimelockChangeIterator{contract: _KromaGovernor.contract, event: "TimelockChange", logs: logs, sub: sub}, nil
}

// WatchTimelockChange is a free log subscription operation binding the contract event 0x08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401.
//
// Solidity: event TimelockChange(address oldTimelock, address newTimelock)
func (_KromaGovernor *KromaGovernorFilterer) WatchTimelockChange(opts *bind.WatchOpts, sink chan<- *KromaGovernorTimelockChange) (event.Subscription, error) {

	logs, sub, err := _KromaGovernor.contract.WatchLogs(opts, "TimelockChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KromaGovernorTimelockChange)
				if err := _KromaGovernor.contract.UnpackLog(event, "TimelockChange", log); err != nil {
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
func (_KromaGovernor *KromaGovernorFilterer) ParseTimelockChange(log types.Log) (*KromaGovernorTimelockChange, error) {
	event := new(KromaGovernorTimelockChange)
	if err := _KromaGovernor.contract.UnpackLog(event, "TimelockChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KromaGovernorVoteCastIterator is returned from FilterVoteCast and is used to iterate over the raw logs and unpacked data for VoteCast events raised by the KromaGovernor contract.
type KromaGovernorVoteCastIterator struct {
	Event *KromaGovernorVoteCast // Event containing the contract specifics and raw log

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
func (it *KromaGovernorVoteCastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KromaGovernorVoteCast)
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
		it.Event = new(KromaGovernorVoteCast)
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
func (it *KromaGovernorVoteCastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KromaGovernorVoteCastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KromaGovernorVoteCast represents a VoteCast event raised by the KromaGovernor contract.
type KromaGovernorVoteCast struct {
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
func (_KromaGovernor *KromaGovernorFilterer) FilterVoteCast(opts *bind.FilterOpts, voter []common.Address) (*KromaGovernorVoteCastIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _KromaGovernor.contract.FilterLogs(opts, "VoteCast", voterRule)
	if err != nil {
		return nil, err
	}
	return &KromaGovernorVoteCastIterator{contract: _KromaGovernor.contract, event: "VoteCast", logs: logs, sub: sub}, nil
}

// WatchVoteCast is a free log subscription operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (_KromaGovernor *KromaGovernorFilterer) WatchVoteCast(opts *bind.WatchOpts, sink chan<- *KromaGovernorVoteCast, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _KromaGovernor.contract.WatchLogs(opts, "VoteCast", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KromaGovernorVoteCast)
				if err := _KromaGovernor.contract.UnpackLog(event, "VoteCast", log); err != nil {
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
func (_KromaGovernor *KromaGovernorFilterer) ParseVoteCast(log types.Log) (*KromaGovernorVoteCast, error) {
	event := new(KromaGovernorVoteCast)
	if err := _KromaGovernor.contract.UnpackLog(event, "VoteCast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KromaGovernorVoteCastWithParamsIterator is returned from FilterVoteCastWithParams and is used to iterate over the raw logs and unpacked data for VoteCastWithParams events raised by the KromaGovernor contract.
type KromaGovernorVoteCastWithParamsIterator struct {
	Event *KromaGovernorVoteCastWithParams // Event containing the contract specifics and raw log

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
func (it *KromaGovernorVoteCastWithParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KromaGovernorVoteCastWithParams)
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
		it.Event = new(KromaGovernorVoteCastWithParams)
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
func (it *KromaGovernorVoteCastWithParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KromaGovernorVoteCastWithParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KromaGovernorVoteCastWithParams represents a VoteCastWithParams event raised by the KromaGovernor contract.
type KromaGovernorVoteCastWithParams struct {
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
func (_KromaGovernor *KromaGovernorFilterer) FilterVoteCastWithParams(opts *bind.FilterOpts, voter []common.Address) (*KromaGovernorVoteCastWithParamsIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _KromaGovernor.contract.FilterLogs(opts, "VoteCastWithParams", voterRule)
	if err != nil {
		return nil, err
	}
	return &KromaGovernorVoteCastWithParamsIterator{contract: _KromaGovernor.contract, event: "VoteCastWithParams", logs: logs, sub: sub}, nil
}

// WatchVoteCastWithParams is a free log subscription operation binding the contract event 0xe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (_KromaGovernor *KromaGovernorFilterer) WatchVoteCastWithParams(opts *bind.WatchOpts, sink chan<- *KromaGovernorVoteCastWithParams, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _KromaGovernor.contract.WatchLogs(opts, "VoteCastWithParams", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KromaGovernorVoteCastWithParams)
				if err := _KromaGovernor.contract.UnpackLog(event, "VoteCastWithParams", log); err != nil {
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
func (_KromaGovernor *KromaGovernorFilterer) ParseVoteCastWithParams(log types.Log) (*KromaGovernorVoteCastWithParams, error) {
	event := new(KromaGovernorVoteCastWithParams)
	if err := _KromaGovernor.contract.UnpackLog(event, "VoteCastWithParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KromaGovernorVotingDelaySetIterator is returned from FilterVotingDelaySet and is used to iterate over the raw logs and unpacked data for VotingDelaySet events raised by the KromaGovernor contract.
type KromaGovernorVotingDelaySetIterator struct {
	Event *KromaGovernorVotingDelaySet // Event containing the contract specifics and raw log

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
func (it *KromaGovernorVotingDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KromaGovernorVotingDelaySet)
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
		it.Event = new(KromaGovernorVotingDelaySet)
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
func (it *KromaGovernorVotingDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KromaGovernorVotingDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KromaGovernorVotingDelaySet represents a VotingDelaySet event raised by the KromaGovernor contract.
type KromaGovernorVotingDelaySet struct {
	OldVotingDelay *big.Int
	NewVotingDelay *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterVotingDelaySet is a free log retrieval operation binding the contract event 0xc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93.
//
// Solidity: event VotingDelaySet(uint256 oldVotingDelay, uint256 newVotingDelay)
func (_KromaGovernor *KromaGovernorFilterer) FilterVotingDelaySet(opts *bind.FilterOpts) (*KromaGovernorVotingDelaySetIterator, error) {

	logs, sub, err := _KromaGovernor.contract.FilterLogs(opts, "VotingDelaySet")
	if err != nil {
		return nil, err
	}
	return &KromaGovernorVotingDelaySetIterator{contract: _KromaGovernor.contract, event: "VotingDelaySet", logs: logs, sub: sub}, nil
}

// WatchVotingDelaySet is a free log subscription operation binding the contract event 0xc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93.
//
// Solidity: event VotingDelaySet(uint256 oldVotingDelay, uint256 newVotingDelay)
func (_KromaGovernor *KromaGovernorFilterer) WatchVotingDelaySet(opts *bind.WatchOpts, sink chan<- *KromaGovernorVotingDelaySet) (event.Subscription, error) {

	logs, sub, err := _KromaGovernor.contract.WatchLogs(opts, "VotingDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KromaGovernorVotingDelaySet)
				if err := _KromaGovernor.contract.UnpackLog(event, "VotingDelaySet", log); err != nil {
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
func (_KromaGovernor *KromaGovernorFilterer) ParseVotingDelaySet(log types.Log) (*KromaGovernorVotingDelaySet, error) {
	event := new(KromaGovernorVotingDelaySet)
	if err := _KromaGovernor.contract.UnpackLog(event, "VotingDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KromaGovernorVotingPeriodSetIterator is returned from FilterVotingPeriodSet and is used to iterate over the raw logs and unpacked data for VotingPeriodSet events raised by the KromaGovernor contract.
type KromaGovernorVotingPeriodSetIterator struct {
	Event *KromaGovernorVotingPeriodSet // Event containing the contract specifics and raw log

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
func (it *KromaGovernorVotingPeriodSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KromaGovernorVotingPeriodSet)
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
		it.Event = new(KromaGovernorVotingPeriodSet)
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
func (it *KromaGovernorVotingPeriodSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KromaGovernorVotingPeriodSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KromaGovernorVotingPeriodSet represents a VotingPeriodSet event raised by the KromaGovernor contract.
type KromaGovernorVotingPeriodSet struct {
	OldVotingPeriod *big.Int
	NewVotingPeriod *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVotingPeriodSet is a free log retrieval operation binding the contract event 0x7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828.
//
// Solidity: event VotingPeriodSet(uint256 oldVotingPeriod, uint256 newVotingPeriod)
func (_KromaGovernor *KromaGovernorFilterer) FilterVotingPeriodSet(opts *bind.FilterOpts) (*KromaGovernorVotingPeriodSetIterator, error) {

	logs, sub, err := _KromaGovernor.contract.FilterLogs(opts, "VotingPeriodSet")
	if err != nil {
		return nil, err
	}
	return &KromaGovernorVotingPeriodSetIterator{contract: _KromaGovernor.contract, event: "VotingPeriodSet", logs: logs, sub: sub}, nil
}

// WatchVotingPeriodSet is a free log subscription operation binding the contract event 0x7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828.
//
// Solidity: event VotingPeriodSet(uint256 oldVotingPeriod, uint256 newVotingPeriod)
func (_KromaGovernor *KromaGovernorFilterer) WatchVotingPeriodSet(opts *bind.WatchOpts, sink chan<- *KromaGovernorVotingPeriodSet) (event.Subscription, error) {

	logs, sub, err := _KromaGovernor.contract.WatchLogs(opts, "VotingPeriodSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KromaGovernorVotingPeriodSet)
				if err := _KromaGovernor.contract.UnpackLog(event, "VotingPeriodSet", log); err != nil {
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
func (_KromaGovernor *KromaGovernorFilterer) ParseVotingPeriodSet(log types.Log) (*KromaGovernorVotingPeriodSet, error) {
	event := new(KromaGovernorVotingPeriodSet)
	if err := _KromaGovernor.contract.UnpackLog(event, "VotingPeriodSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
