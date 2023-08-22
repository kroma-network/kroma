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
	Bin: "0x60e06040523480156200001157600080fd5b5060006080819052600160a05260c0526200002b62000031565b620000f2565b600054610100900460ff16156200009e5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811614620000f0576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b60805160a05160c051615deb6200012260003960006126a1015260006126780152600061264f0152615deb6000f3fe60806040526004361061030c5760003560e01c80637b3c71d31161019a578063c01f9e37116100e1578063ea0217cf1161008a578063f23a6e6111610064578063f23a6e6114610a6d578063f8ce560a14610ab2578063fc0c546a14610ad257600080fd5b8063ea0217cf14610a0d578063eb9019d414610a2d578063ece40cc114610a4d57600080fd5b8063d33219b4116100bb578063d33219b414610974578063dd4e2ba514610993578063deaaa7cc146109d957600080fd5b8063c01f9e3714610907578063c28bc2fa14610941578063c59057e41461095457600080fd5b80639a802a6d11610143578063ab58fb8e1161011d578063ab58fb8e1461088d578063b58131b0146108ad578063bc197c81146108c257600080fd5b80639a802a6d14610838578063a7713a7014610858578063a890c9101461086d57600080fd5b806386489ba91161017457806386489ba9146107d857806391ddadf4146107f857806397c3d3341461082457600080fd5b80637b3c71d3146107705780637d5e81e21461079057806384b0196e146107b057600080fd5b80633932abb11161025e578063544ffc9c116102075780635f398a14116101e15780635f398a141461071057806360c4247f1461073057806370b0f6601461075057600080fd5b8063544ffc9c1461068557806354fd4d50146106db57806356781388146106f057600080fd5b806343859632116102385780634385963214610605578063452115d6146106505780634bf5d7e91461067057600080fd5b80633932abb1146105a35780633bccf4fd146105b85780633e4f49e6146105d857600080fd5b8063143489d0116102c05780632656227d1161029a5780632656227d146105255780632d63f693146105385780632fe3e2611461056f57600080fd5b8063143489d014610436578063150b7a0214610490578063160cbed71461050557600080fd5b806303420181116102f157806303420181146103d457806306f3f9e6146103f457806306fdde031461041457600080fd5b806301ffc9a71461037c57806302a251a3146103b157600080fd5b36610377573061031a610af3565b6001600160a01b0316146103755760405162461bcd60e51b815260206004820152601f60248201527f476f7665726e6f723a206d7573742073656e6420746f206578656375746f720060448201526064015b60405180910390fd5b005b600080fd5b34801561038857600080fd5b5061039c610397366004614c41565b610b0d565b60405190151581526020015b60405180910390f35b3480156103bd57600080fd5b506103c6610b1e565b6040519081526020016103a8565b3480156103e057600080fd5b506103c66103ef366004614dc3565b610b2a565b34801561040057600080fd5b5061037561040f366004614e6a565b610c22565b34801561042057600080fd5b50610429610cdc565b6040516103a89190614edf565b34801561044257600080fd5b50610478610451366004614e6a565b600090815260fe60205260409020546801000000000000000090046001600160a01b031690565b6040516001600160a01b0390911681526020016103a8565b34801561049c57600080fd5b506104d46104ab366004614f07565b7f150b7a0200000000000000000000000000000000000000000000000000000000949350505050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020016103a8565b34801561051157600080fd5b506103c66105203660046150e1565b610d6e565b6103c66105333660046150e1565b61109b565b34801561054457600080fd5b506103c6610553366004614e6a565b600090815260fe602052604090205467ffffffffffffffff1690565b34801561057b57600080fd5b506103c67fb3b3f3b703cd84ce352197dcff232b1b5d3cfb2025ce47cf04742d0651f1af8881565b3480156105af57600080fd5b506103c6611201565b3480156105c457600080fd5b506103c66105d3366004615171565b61120d565b3480156105e457600080fd5b506105f86105f3366004614e6a565b611283565b6040516103a891906151ee565b34801561061157600080fd5b5061039c61062036600461522f565b6000828152610161602090815260408083206001600160a01b038516845260030190915290205460ff1692915050565b34801561065c57600080fd5b506103c661066b3660046150e1565b61128e565b34801561067c57600080fd5b506104296113bf565b34801561069157600080fd5b506106c06106a0366004614e6a565b600090815261016160205260409020805460018201546002909201549092565b604080519384526020840192909252908201526060016103a8565b3480156106e757600080fd5b50610429611485565b3480156106fc57600080fd5b506103c661070b36600461525f565b61148f565b34801561071c57600080fd5b506103c661072b36600461528b565b6114b8565b34801561073c57600080fd5b506103c661074b366004614e6a565b611502565b34801561075c57600080fd5b5061037561076b366004614e6a565b6115f7565b34801561077c57600080fd5b506103c661078b36600461530f565b6116ae565b34801561079c57600080fd5b506103c66107ab366004615369565b6116f6565b3480156107bc57600080fd5b506107c561170d565b6040516103a89796959493929190615459565b3480156107e457600080fd5b506103756107f33660046154d5565b6117cf565b34801561080457600080fd5b5061080d6119b1565b60405165ffffffffffff90911681526020016103a8565b34801561083057600080fd5b5060646103c6565b34801561084457600080fd5b506103c661085336600461552e565b611a3e565b34801561086457600080fd5b506103c6611a55565b34801561087957600080fd5b50610375610888366004615587565b611a97565b34801561089957600080fd5b506103c66108a8366004614e6a565b611b4e565b3480156108b957600080fd5b506103c6611c03565b3480156108ce57600080fd5b506104d46108dd3660046155a4565b7fbc197c810000000000000000000000000000000000000000000000000000000095945050505050565b34801561091357600080fd5b506103c6610922366004614e6a565b600090815260fe602052604090206001015467ffffffffffffffff1690565b61037561094f366004615638565b611c0f565b34801561096057600080fd5b506103c661096f3660046150e1565b611d45565b34801561098057600080fd5b506101f8546001600160a01b0316610478565b34801561099f57600080fd5b506040805180820190915260208082527f737570706f72743d627261766f2671756f72756d3d666f722c6162737461696e90820152610429565b3480156109e557600080fd5b506103c67f150214d74d59b7d1e90c73fc22ef3d991dd0a76b046543d4d80ab92d2a50328f81565b348015610a1957600080fd5b50610375610a28366004614e6a565b611d7f565b348015610a3957600080fd5b506103c6610a4836600461567c565b611e36565b348015610a5957600080fd5b50610375610a68366004614e6a565b611e57565b348015610a7957600080fd5b506104d4610a883660046156a8565b7ff23a6e610000000000000000000000000000000000000000000000000000000095945050505050565b348015610abe57600080fd5b506103c6610acd366004614e6a565b611f0e565b348015610ade57600080fd5b5061019354610478906001600160a01b031681565b6000610b086101f8546001600160a01b031690565b905090565b6000610b1882611f19565b92915050565b6000610b086101305490565b600080610bce610bc67fb3b3f3b703cd84ce352197dcff232b1b5d3cfb2025ce47cf04742d0651f1af888c8c8c8c604051610b66929190615711565b60405180910390208b80519060200120604051602001610bab959493929190948552602085019390935260ff9190911660408401526060830152608082015260a00190565b60405160208183030381529060405280519060200120611f6f565b868686611fb7565b9050610c148a828b8b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508d9250611fd5915050565b9a9950505050505050505050565b610c2a610af3565b6001600160a01b0316336001600160a01b031614610c8a5760405162461bcd60e51b815260206004820152601860248201527f476f7665726e6f723a206f6e6c79476f7665726e616e63650000000000000000604482015260640161036c565b30610c93610af3565b6001600160a01b031614610cd05760008036604051610cb3929190615711565b604051809103902090505b80610cc960ff612145565b03610cbe57505b610cd981612202565b50565b606060fd8054610ceb90615721565b80601f0160208091040260200160405190810160405280929190818152602001828054610d1790615721565b8015610d645780601f10610d3957610100808354040283529160200191610d64565b820191906000526020600020905b815481529060010190602001808311610d4757829003601f168201915b5050505050905090565b600080610d7d86868686611d45565b90506004610d8a82611283565b6007811115610d9b57610d9b6151bf565b14610e0e5760405162461bcd60e51b815260206004820152602160248201527f476f7665726e6f723a2070726f706f73616c206e6f742073756363657373667560448201527f6c00000000000000000000000000000000000000000000000000000000000000606482015260840161036c565b6101f854604080517ff27a0c9200000000000000000000000000000000000000000000000000000000815290516000926001600160a01b03169163f27a0c929160048083019260209291908290030181865afa158015610e72573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e969190615774565b905060006001600160a01b031687600081518110610eb657610eb661578d565b60200260200101516001600160a01b0316148015610eee575085600081518110610ee257610ee261578d565b60200260200101516000145b8015610f15575084600081518110610f0857610f0861578d565b6020026020010151516000145b15610f1e575060005b6101f8546040517fb1c5f4270000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063b1c5f42790610f71908a908a908a906000908b9060040161584a565b602060405180830381865afa158015610f8e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fb29190615774565b60008381526101f96020526040808220929092556101f85491517f8f2a0bb00000000000000000000000000000000000000000000000000000000081526001600160a01b0390921691638f2a0bb091611018918b918b918b91908b908990600401615898565b600060405180830381600087803b15801561103257600080fd5b505af1158015611046573d6000803e3d6000fd5b505050507f9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892828242611078919061591f565b604080519283526020830191909152015b60405180910390a15095945050505050565b6000806110aa86868686611d45565b905060006110b782611283565b905060048160078111156110cd576110cd6151bf565b14806110ea575060058160078111156110e8576110e86151bf565b145b61115c5760405162461bcd60e51b815260206004820152602160248201527f476f7665726e6f723a2070726f706f73616c206e6f742073756363657373667560448201527f6c00000000000000000000000000000000000000000000000000000000000000606482015260840161036c565b600082815260fe60205260409081902060020180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055517f712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f906111c89084815260200190565b60405180910390a16111dd82888888886123a3565b6111ea8288888888612445565b6111f78288888888612452565b5095945050505050565b6000610b0861012f5490565b604080517f150214d74d59b7d1e90c73fc22ef3d991dd0a76b046543d4d80ab92d2a50328f602082015290810186905260ff85166060820152600090819061125b90610bc690608001610bab565b905061127887828860405180602001604052806000815250612498565b979650505050505050565b6000610b18826124bb565b60008061129d86868686611d45565b905060006112aa82611283565b60078111156112bb576112bb6151bf565b146113085760405162461bcd60e51b815260206004820152601c60248201527f476f7665726e6f723a20746f6f206c61746520746f2063616e63656c00000000604482015260640161036c565b600081815260fe60205260409020546801000000000000000090046001600160a01b0316336001600160a01b0316146113a95760405162461bcd60e51b815260206004820152602260248201527f476f7665726e6f723a206f6e6c792070726f706f7365722063616e2063616e6360448201527f656c000000000000000000000000000000000000000000000000000000000000606482015260840161036c565b6113b58686868661263a565b9695505050505050565b61019354604080517f4bf5d7e900000000000000000000000000000000000000000000000000000000815290516060926001600160a01b031691634bf5d7e99160048083019260009291908290030181865afa92505050801561144457506040513d6000823e601f3d908101601f191682016040526114419190810190615937565b60015b611480575060408051808201909152601d81527f6d6f64653d626c6f636b6e756d6265722666726f6d3d64656661756c74000000602082015290565b919050565b6060610b08612648565b6000803390506114b084828560405180602001604052806000815250612498565b949350505050565b60008033905061127887828888888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508a9250611fd5915050565b6101c75460009080820361151b5750506101c654919050565b60006101c761152b6001846159a5565b8154811061153b5761153b61578d565b60009182526020918290206040805180820190915291015463ffffffff81168083526401000000009091047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1692820192909252915084106115bc57602001517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff169392505050565b6115d16115c8856126eb565b6101c79061276b565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16949350505050565b6115ff610af3565b6001600160a01b0316336001600160a01b03161461165f5760405162461bcd60e51b815260206004820152601860248201527f476f7665726e6f723a206f6e6c79476f7665726e616e63650000000000000000604482015260640161036c565b30611668610af3565b6001600160a01b0316146116a55760008036604051611688929190615711565b604051809103902090505b8061169e60ff612145565b0361169357505b610cd981612834565b6000803390506113b586828787878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061249892505050565b600061170485858585612877565b95945050505050565b6000606080600080600060606065546000801b14801561172d5750606654155b6117795760405162461bcd60e51b815260206004820152601560248201527f4549503731323a20556e696e697469616c697a65640000000000000000000000604482015260640161036c565b611781612dac565b611789612dbb565b604080516000808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009b939a50919850469750309650945092509050565b600054610100900460ff16158080156117ef5750600054600160ff909116105b806118095750303b158015611809575060005460ff166001145b61187b5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161036c565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156118d957600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6119176040518060400160405280600f81526020017f55706772616465476f7665726e6f720000000000000000000000000000000000815250612dca565b611922858585612e61565b61192a612eee565b61193387612f6d565b61193c82612ff3565b61194586613079565b80156119a857600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b61019354604080517f91ddadf400000000000000000000000000000000000000000000000000000000815290516000926001600160a01b0316916391ddadf49160048083019260209291908290030181865afa925050508015611a31575060408051601f3d908101601f19168201909252611a2e918101906159bc565b60015b61148057610b08436130ff565b6000611a4b84848461317d565b90505b9392505050565b6101c75460009015611a8f57611a6c6101c761320d565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16905090565b506101c65490565b611a9f610af3565b6001600160a01b0316336001600160a01b031614611aff5760405162461bcd60e51b815260206004820152601860248201527f476f7665726e6f723a206f6e6c79476f7665726e616e63650000000000000000604482015260640161036c565b30611b08610af3565b6001600160a01b031614611b455760008036604051611b28929190615711565b604051809103902090505b80611b3e60ff612145565b03611b3357505b610cd981613253565b6101f85460008281526101f960205260408082205490517fd45c44350000000000000000000000000000000000000000000000000000000081526004810191909152909182916001600160a01b039091169063d45c443590602401602060405180830381865afa158015611bc6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bea9190615774565b905080600114611bfa5780611a4e565b60009392505050565b6000610b086101315490565b611c17610af3565b6001600160a01b0316336001600160a01b031614611c775760405162461bcd60e51b815260206004820152601860248201527f476f7665726e6f723a206f6e6c79476f7665726e616e63650000000000000000604482015260640161036c565b30611c80610af3565b6001600160a01b031614611cbd5760008036604051611ca0929190615711565b604051809103902090505b80611cb660ff612145565b03611cab57505b600080856001600160a01b0316858585604051611cdb929190615711565b60006040518083038185875af1925050503d8060008114611d18576040519150601f19603f3d011682016040523d82523d6000602084013e611d1d565b606091505b50915091506119a88282604051806060016040528060288152602001615db7602891396132d6565b600084848484604051602001611d5e94939291906159e4565b60408051601f19818403018152919052805160209091012095945050505050565b611d87610af3565b6001600160a01b0316336001600160a01b031614611de75760405162461bcd60e51b815260206004820152601860248201527f476f7665726e6f723a206f6e6c79476f7665726e616e63650000000000000000604482015260640161036c565b30611df0610af3565b6001600160a01b031614611e2d5760008036604051611e10929190615711565b604051809103902090505b80611e2660ff612145565b03611e1b57505b610cd9816132ef565b6000611a4e8383611e5260408051602081019091526000815290565b61317d565b611e5f610af3565b6001600160a01b0316336001600160a01b031614611ebf5760405162461bcd60e51b815260206004820152601860248201527f476f7665726e6f723a206f6e6c79476f7665726e616e63650000000000000000604482015260640161036c565b30611ec8610af3565b6001600160a01b031614611f055760008036604051611ee8929190615711565b604051809103902090505b80611efe60ff612145565b03611ef357505b610cd9816133a8565b6000610b18826133eb565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f6e665ced000000000000000000000000000000000000000000000000000000001480610b185750610b1882613493565b6000610b18611f7c61363d565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b6000806000611fc887878787613647565b915091506111f78161370b565b600085815260fe602052604081206001611fee88611283565b6007811115611fff57611fff6151bf565b146120725760405162461bcd60e51b815260206004820152602360248201527f476f7665726e6f723a20766f7465206e6f742063757272656e746c792061637460448201527f6976650000000000000000000000000000000000000000000000000000000000606482015260840161036c565b805460009061208d90889067ffffffffffffffff168661317d565b905061209c8888888488613870565b83516000036120f157866001600160a01b03167fb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4898884896040516120e49493929190615a2f565b60405180910390a2611278565b866001600160a01b03167fe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb87128988848989604051612132959493929190615a57565b60405180910390a2979650505050505050565b600061216d8254600f81810b700100000000000000000000000000000000909204900b131590565b156121a4576040517f3db2a12a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b508054600f0b6000818152600180840160205260408220805492905583547fffffffffffffffffffffffffffffffff000000000000000000000000000000001692016fffffffffffffffffffffffffffffffff169190911790915590565b606481111561229f5760405162461bcd60e51b815260206004820152604360248201527f476f7665726e6f72566f74657351756f72756d4672616374696f6e3a2071756f60448201527f72756d4e756d657261746f72206f7665722071756f72756d44656e6f6d696e6160648201527f746f720000000000000000000000000000000000000000000000000000000000608482015260a40161036c565b60006122a9611a55565b905080158015906122bb57506101c754155b156123365760408051808201909152600081526101c790602081016122df84613a45565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff90811690915282546001810184556000938452602093849020835194909301519091166401000000000263ffffffff909316929092179101555b6123646123516123446119b1565b65ffffffffffff166126eb565b61235a84613a45565b6101c79190613ad9565b505060408051828152602081018490527f0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997910160405180910390a15050565b306123ac610af3565b6001600160a01b03161461243e5760005b845181101561243c57306001600160a01b03168582815181106123e2576123e261578d565b60200260200101516001600160a01b03160361242c5761242c83828151811061240d5761240d61578d565b60200260200101518051906020012060ff613af490919063ffffffff16565b61243581615a9d565b90506123bd565b505b5050505050565b61243e8585858585613b46565b3061245b610af3565b6001600160a01b03161461243e5760ff54600f81810b700100000000000000000000000000000000909204900b131561243e57600060ff5561243e565b6000611704858585856124b660408051602081019091526000815290565b611fd5565b6000806124c783613bd4565b905060048160078111156124dd576124dd6151bf565b146124e85792915050565b60008381526101f9602052604090205480612504575092915050565b6101f8546040517f2ab0f529000000000000000000000000000000000000000000000000000000008152600481018390526001600160a01b0390911690632ab0f52990602401602060405180830381865afa158015612567573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061258b9190615ab7565b1561259a575060079392505050565b6101f8546040517f584b153e000000000000000000000000000000000000000000000000000000008152600481018390526001600160a01b039091169063584b153e90602401602060405180830381865afa1580156125fd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126219190615ab7565b15612630575060059392505050565b5060029392505050565b600061170485858585613d17565b60606126737f0000000000000000000000000000000000000000000000000000000000000000613de6565b61269c7f0000000000000000000000000000000000000000000000000000000000000000613de6565b6126c57f0000000000000000000000000000000000000000000000000000000000000000613de6565b6040516020016126d793929190615ad9565b604051602081830303815290604052905090565b600063ffffffff8211156127675760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203360448201527f3220626974730000000000000000000000000000000000000000000000000000606482015260840161036c565b5090565b8154600090818160058111156127c857600061278684613e86565b61279090856159a5565b60008881526020902090915081015463ffffffff90811690871610156127b8578091506127c6565b6127c381600161591f565b92505b505b60006127d687878585613f6e565b90508015612827576127fb876127ed6001846159a5565b600091825260209091200190565b5464010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16611278565b6000979650505050505050565b61012f5460408051918252602082018390527fc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93910160405180910390a161012f55565b6000336128848184613fcc565b6128d05760405162461bcd60e51b815260206004820152601d60248201527f476f7665726e6f723a2070726f706f7365722072657374726963746564000000604482015260640161036c565b60006128da6119b1565b65ffffffffffff1690506128ec611c03565b6128fb83610a486001856159a5565b101561296f5760405162461bcd60e51b815260206004820152603160248201527f476f7665726e6f723a2070726f706f73657220766f7465732062656c6f77207060448201527f726f706f73616c207468726573686f6c64000000000000000000000000000000606482015260840161036c565b60006129848888888880519060200120611d45565b905086518851146129fd5760405162461bcd60e51b815260206004820152602160248201527f476f7665726e6f723a20696e76616c69642070726f706f73616c206c656e677460448201527f6800000000000000000000000000000000000000000000000000000000000000606482015260840161036c565b8551885114612a745760405162461bcd60e51b815260206004820152602160248201527f476f7665726e6f723a20696e76616c69642070726f706f73616c206c656e677460448201527f6800000000000000000000000000000000000000000000000000000000000000606482015260840161036c565b6000885111612ac55760405162461bcd60e51b815260206004820152601860248201527f476f7665726e6f723a20656d7074792070726f706f73616c0000000000000000604482015260640161036c565b600081815260fe602052604090205467ffffffffffffffff1615612b515760405162461bcd60e51b815260206004820152602160248201527f476f7665726e6f723a2070726f706f73616c20616c726561647920657869737460448201527f7300000000000000000000000000000000000000000000000000000000000000606482015260840161036c565b6000612b5b611201565b612b65908461591f565b90506000612b71610b1e565b612b7b908361591f565b90506040518060e00160405280612b918461411c565b67ffffffffffffffff1681526001600160a01b038716602082015260006040820152606001612bbf8361411c565b67ffffffffffffffff9081168252600060208084018290526040808501839052606094850183905288835260fe8252918290208551815492870151878501519186167fffffffff0000000000000000000000000000000000000000000000000000000090941693909317680100000000000000006001600160a01b039094168402177bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167c010000000000000000000000000000000000000000000000000000000060e09290921c91909102178155938501516080860151908416921c0217600183015560a08301516002909201805460c0909401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00009094169215157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1692909217610100931515939093029290921790558a517f7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e091859188918e918e91811115612d4957612d49614cdd565b604051908082528060200260200182016040528015612d7c57816020015b6060815260200190600190039081612d675790505b508d88888f604051612d9699989796959493929190615b4f565b60405180910390a1509098975050505050505050565b606060678054610ceb90615721565b606060688054610ceb90615721565b600054610100900460ff16612e475760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b612e5881612e53611485565b61419c565b610cd981614241565b600054610100900460ff16612ede5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b612ee98383836142ce565b505050565b600054610100900460ff16612f6b5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b565b600054610100900460ff16612fea5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b610cd981614366565b600054610100900460ff166130705760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b610cd98161441e565b600054610100900460ff166130f65760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b610cd98161449b565b600065ffffffffffff8211156127675760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203460448201527f3820626974730000000000000000000000000000000000000000000000000000606482015260840161036c565b610193546040517f3a46b1a80000000000000000000000000000000000000000000000000000000081526001600160a01b038581166004830152602482018590526000921690633a46b1a890604401602060405180830381865afa1580156131e9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a4b9190615774565b80546000908015611bfa57613227836127ed6001846159a5565b5464010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16611a4e565b6101f854604080516001600160a01b03928316815291831660208301527f08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401910160405180910390a16101f880547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b606083156132e5575081611a4e565b611a4e8383614518565b600081116133655760405162461bcd60e51b815260206004820152602760248201527f476f7665726e6f7253657474696e67733a20766f74696e6720706572696f642060448201527f746f6f206c6f7700000000000000000000000000000000000000000000000000606482015260840161036c565b6101305460408051918252602082018390527f7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828910160405180910390a161013055565b6101315460408051918252602082018390527fccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc05461910160405180910390a161013155565b600060646133f883611502565b610193546040517f8e539e8c000000000000000000000000000000000000000000000000000000008152600481018690526001600160a01b0390911690638e539e8c90602401602060405180830381865afa15801561345b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061347f9190615774565b6134899190615c27565b610b189190615c75565b60007f51159c06000000000000000000000000000000000000000000000000000000007fc6fba1f8000000000000000000000000000000000000000000000000000000007fbf26d897000000000000000000000000000000000000000000000000000000007f79dd796f000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000861682148061356d57507fffffffff00000000000000000000000000000000000000000000000000000000868116908216145b8061359c57507fffffffff00000000000000000000000000000000000000000000000000000000868116908516145b806135e857507fffffffff0000000000000000000000000000000000000000000000000000000086167f4e2312e000000000000000000000000000000000000000000000000000000000145b806113b557507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008716149695505050505050565b6000610b08614542565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561367e5750600090506003613702565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa1580156136d2573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166136fb57600060019250925050613702565b9150600090505b94509492505050565b600081600481111561371f5761371f6151bf565b036137275750565b600181600481111561373b5761373b6151bf565b036137885760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161036c565b600281600481111561379c5761379c6151bf565b036137e95760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161036c565b60038160048111156137fd576137fd6151bf565b03610cd95760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f7565000000000000000000000000000000000000000000000000000000000000606482015260840161036c565b6000858152610161602090815260408083206001600160a01b0388168452600381019092529091205460ff161561390f5760405162461bcd60e51b815260206004820152602760248201527f476f7665726e6f72566f74696e6753696d706c653a20766f746520616c72656160448201527f6479206361737400000000000000000000000000000000000000000000000000606482015260840161036c565b6001600160a01b0385166000908152600382016020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905560ff8416613979578281600001600082825461396e919061591f565b9091555061243c9050565b60001960ff851601613999578281600101600082825461396e919061591f565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe60ff8516016139d7578281600201600082825461396e919061591f565b60405162461bcd60e51b815260206004820152603560248201527f476f7665726e6f72566f74696e6753696d706c653a20696e76616c696420766160448201527f6c756520666f7220656e756d20566f7465547970650000000000000000000000606482015260840161036c565b60007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8211156127675760405162461bcd60e51b815260206004820152602760248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203260448201527f3234206269747300000000000000000000000000000000000000000000000000606482015260840161036c565b600080613ae78585856145b6565b915091505b935093915050565b815470010000000000000000000000000000000090819004600f0b6000818152600180860160205260409091209390935583546fffffffffffffffffffffffffffffffff908116939091011602179055565b6101f8546040517fe38335e50000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063e38335e5903490613b9b90889088908890600090899060040161584a565b6000604051808303818588803b158015613bb457600080fd5b505af1158015613bc8573d6000803e3d6000fd5b50505050505050505050565b600081815260fe60205260408120600281015460ff1615613bf85750600792915050565b6002810154610100900460ff1615613c135750600292915050565b600083815260fe602052604081205467ffffffffffffffff1690819003613c7c5760405162461bcd60e51b815260206004820152601d60248201527f476f7665726e6f723a20756e6b6e6f776e2070726f706f73616c206964000000604482015260640161036c565b6000613c866119b1565b65ffffffffffff169050808210613ca257506000949350505050565b600085815260fe602052604090206001015467ffffffffffffffff16818110613cd15750600195945050505050565b613cda866147ad565b8015613cfa57506000868152610161602052604090208054600190910154115b15613d0b5750600495945050505050565b50600395945050505050565b600080613d26868686866147fb565b60008181526101f9602052604090205490915015611704576101f85460008281526101f96020526040908190205490517fc4d252f50000000000000000000000000000000000000000000000000000000081526001600160a01b039092169163c4d252f591613d9b9160040190815260200190565b600060405180830381600087803b158015613db557600080fd5b505af1158015613dc9573d6000803e3d6000fd5b50505060008281526101f960205260408120555095945050505050565b60606000613df383614924565b600101905060008167ffffffffffffffff811115613e1357613e13614cdd565b6040519080825280601f01601f191660200182016040528015613e3d576020820181803683370190505b5090508181016020015b600019017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a8504945084613e4757509392505050565b600081600003613e9857506000919050565b60006001613ea584614a06565b901c6001901b90506001818481613ebe57613ebe615c46565b048201901c90506001818481613ed657613ed6615c46565b048201901c90506001818481613eee57613eee615c46565b048201901c90506001818481613f0657613f06615c46565b048201901c90506001818481613f1e57613f1e615c46565b048201901c90506001818481613f3657613f36615c46565b048201901c90506001818481613f4e57613f4e615c46565b048201901c9050611a4e81828581613f6857613f68615c46565b04614a9a565b60005b81831015613fc4576000613f858484614ab0565b60008781526020902090915063ffffffff86169082015463ffffffff161115613fb057809250613fbe565b613fbb81600161591f565b93505b50613f71565b509392505050565b80516000906034811015613fe4576001915050610b18565b8281017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec01517fffffffffffffffffffffffff000000000000000000000000000000000000000081167f2370726f706f7365723d307800000000000000000000000000000000000000001461405e57600192505050610b18565b60008061406c6028856159a5565b90505b838110156140fb576000806140bb88848151811061408f5761408f61578d565b01602001517fff0000000000000000000000000000000000000000000000000000000000000016614acb565b91509150816140d35760019650505050505050610b18565b8060ff166004856001600160a01b0316901b1793505050806140f490615a9d565b905061406f565b50856001600160a01b0316816001600160a01b031614935050505092915050565b600067ffffffffffffffff8211156127675760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203660448201527f3420626974730000000000000000000000000000000000000000000000000000606482015260840161036c565b600054610100900460ff166142195760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b60676142258382615cf6565b5060686142328282615cf6565b50506000606581905560665550565b600054610100900460ff166142be5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b60fd6142ca8282615cf6565b5050565b600054610100900460ff1661434b5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b61435483612834565b61435d826132ef565b612ee9816133a8565b600054610100900460ff166143e35760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b61019380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b600054610100900460ff16610cd05760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b600054610100900460ff16611b455760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161036c565b8151156145285781518083602001fd5b8060405162461bcd60e51b815260040161036c9190614edf565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f61456d614bb7565b614575614c10565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b82546000908190801561473e5760006145d4876127ed6001856159a5565b60408051808201909152905463ffffffff8082168084526401000000009092047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff166020840152919250908716101561466b5760405162461bcd60e51b815260206004820152601b60248201527f436865636b706f696e743a2064656372656173696e67206b6579730000000000604482015260640161036c565b805163ffffffff8088169116036146c9578461468c886127ed6001866159a5565b80547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff929092166401000000000263ffffffff90921691909117905561472e565b6040805180820190915263ffffffff80881682527bffffffffffffffffffffffffffffffffffffffffffffffffffffffff80881660208085019182528b54600181018d5560008d81529190912094519151909216640100000000029216919091179101555b602001519250839150613aec9050565b50506040805180820190915263ffffffff80851682527bffffffffffffffffffffffffffffffffffffffffffffffffffffffff80851660208085019182528854600181018a5560008a815291822095519251909316640100000000029190931617920191909155905081613aec565b600081815261016160205260408120600281015460018201546147d0919061591f565b600084815260fe60205260409020546147f29067ffffffffffffffff16611f0e565b11159392505050565b60008061480a86868686611d45565b9050600061481782611283565b9050600281600781111561482d5761482d6151bf565b1415801561484d5750600681600781111561484a5761484a6151bf565b14155b801561486b57506007816007811115614868576148686151bf565b14155b6148b75760405162461bcd60e51b815260206004820152601d60248201527f476f7665726e6f723a2070726f706f73616c206e6f7420616374697665000000604482015260640161036c565b600082815260fe60205260409081902060020180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16610100179055517f789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c906110899084815260200190565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000831061496d577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310614999576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc1000083106149b757662386f26fc10000830492506010015b6305f5e10083106149cf576305f5e100830492506008015b61271083106149e357612710830492506004015b606483106149f5576064830492506002015b600a8310610b185760010192915050565b600080608083901c15614a1b57608092831c92015b604083901c15614a2d57604092831c92015b602083901c15614a3f57602092831c92015b601083901c15614a5157601092831c92015b600883901c15614a6357600892831c92015b600483901c15614a7557600492831c92015b600283901c15614a8757600292831c92015b600183901c15610b185760010192915050565b6000818310614aa95781611a4e565b5090919050565b6000614abf6002848418615c75565b611a4e9084841661591f565b60008060f883901c602f81118015614ae65750603a8160ff16105b15614b19576001947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd09091019350915050565b8060ff166040108015614b2f575060478160ff16105b15614b62576001947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc99091019350915050565b8060ff166060108015614b78575060678160ff16105b15614bab576001947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa99091019350915050565b50600093849350915050565b600080614bc2612dac565b805190915015614bd9578051602090910120919050565b6065548015614be85792915050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4709250505090565b600080614c1b612dbb565b805190915015614c32578051602090910120919050565b6066548015614be85792915050565b600060208284031215614c5357600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114611a4e57600080fd5b803560ff8116811461148057600080fd5b60008083601f840112614ca657600080fd5b50813567ffffffffffffffff811115614cbe57600080fd5b602083019150836020828501011115614cd657600080fd5b9250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715614d3557614d35614cdd565b604052919050565b600067ffffffffffffffff821115614d5757614d57614cdd565b50601f01601f191660200190565b6000614d78614d7384614d3d565b614d0c565b9050828152838383011115614d8c57600080fd5b828260208301376000602084830101529392505050565b600082601f830112614db457600080fd5b611a4e83833560208501614d65565b60008060008060008060008060e0898b031215614ddf57600080fd5b88359750614def60208a01614c83565b9650604089013567ffffffffffffffff80821115614e0c57600080fd5b614e188c838d01614c94565b909850965060608b0135915080821115614e3157600080fd5b50614e3e8b828c01614da3565b945050614e4d60808a01614c83565b925060a0890135915060c089013590509295985092959890939650565b600060208284031215614e7c57600080fd5b5035919050565b60005b83811015614e9e578181015183820152602001614e86565b83811115614ead576000848401525b50505050565b60008151808452614ecb816020860160208601614e83565b601f01601f19169290920160200192915050565b602081526000611a4e6020830184614eb3565b6001600160a01b0381168114610cd957600080fd5b60008060008060808587031215614f1d57600080fd5b8435614f2881614ef2565b93506020850135614f3881614ef2565b925060408501359150606085013567ffffffffffffffff811115614f5b57600080fd5b614f6787828801614da3565b91505092959194509250565b600067ffffffffffffffff821115614f8d57614f8d614cdd565b5060051b60200190565b600082601f830112614fa857600080fd5b81356020614fb8614d7383614f73565b82815260059290921b84018101918181019086841115614fd757600080fd5b8286015b84811015614ffb578035614fee81614ef2565b8352918301918301614fdb565b509695505050505050565b600082601f83011261501757600080fd5b81356020615027614d7383614f73565b82815260059290921b8401810191818101908684111561504657600080fd5b8286015b84811015614ffb578035835291830191830161504a565b600082601f83011261507257600080fd5b81356020615082614d7383614f73565b82815260059290921b840181019181810190868411156150a157600080fd5b8286015b84811015614ffb57803567ffffffffffffffff8111156150c55760008081fd5b6150d38986838b0101614da3565b8452509183019183016150a5565b600080600080608085870312156150f757600080fd5b843567ffffffffffffffff8082111561510f57600080fd5b61511b88838901614f97565b9550602087013591508082111561513157600080fd5b61513d88838901615006565b9450604087013591508082111561515357600080fd5b5061516087828801615061565b949793965093946060013593505050565b600080600080600060a0868803121561518957600080fd5b8535945061519960208701614c83565b93506151a760408701614c83565b94979396509394606081013594506080013592915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6020810160088310615229577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b6000806040838503121561524257600080fd5b82359150602083013561525481614ef2565b809150509250929050565b6000806040838503121561527257600080fd5b8235915061528260208401614c83565b90509250929050565b6000806000806000608086880312156152a357600080fd5b853594506152b360208701614c83565b9350604086013567ffffffffffffffff808211156152d057600080fd5b6152dc89838a01614c94565b909550935060608801359150808211156152f557600080fd5b5061530288828901614da3565b9150509295509295909350565b6000806000806060858703121561532557600080fd5b8435935061533560208601614c83565b9250604085013567ffffffffffffffff81111561535157600080fd5b61535d87828801614c94565b95989497509550505050565b6000806000806080858703121561537f57600080fd5b843567ffffffffffffffff8082111561539757600080fd5b6153a388838901614f97565b955060208701359150808211156153b957600080fd5b6153c588838901615006565b945060408701359150808211156153db57600080fd5b6153e788838901615061565b935060608701359150808211156153fd57600080fd5b508501601f8101871361540f57600080fd5b614f6787823560208401614d65565b600081518084526020808501945080840160005b8381101561544e57815187529582019590820190600101615432565b509495945050505050565b7fff000000000000000000000000000000000000000000000000000000000000008816815260e06020820152600061549460e0830189614eb3565b82810360408401526154a68189614eb3565b90508660608401526001600160a01b03861660808401528460a084015282810360c0840152610c14818561541e565b60008060008060008060c087890312156154ee57600080fd5b86356154f981614ef2565b9550602087013561550981614ef2565b95989597505050506040840135936060810135936080820135935060a0909101359150565b60008060006060848603121561554357600080fd5b833561554e81614ef2565b925060208401359150604084013567ffffffffffffffff81111561557157600080fd5b61557d86828701614da3565b9150509250925092565b60006020828403121561559957600080fd5b8135611a4e81614ef2565b600080600080600060a086880312156155bc57600080fd5b85356155c781614ef2565b945060208601356155d781614ef2565b9350604086013567ffffffffffffffff808211156155f457600080fd5b61560089838a01615006565b9450606088013591508082111561561657600080fd5b61562289838a01615006565b935060808801359150808211156152f557600080fd5b6000806000806060858703121561564e57600080fd5b843561565981614ef2565b935060208501359250604085013567ffffffffffffffff81111561535157600080fd5b6000806040838503121561568f57600080fd5b823561569a81614ef2565b946020939093013593505050565b600080600080600060a086880312156156c057600080fd5b85356156cb81614ef2565b945060208601356156db81614ef2565b93506040860135925060608601359150608086013567ffffffffffffffff81111561570557600080fd5b61530288828901614da3565b8183823760009101908152919050565b600181811c9082168061573557607f821691505b60208210810361576e577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b60006020828403121561578657600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600081518084526020808501945080840160005b8381101561544e5781516001600160a01b0316875295820195908201906001016157d0565b600081518084526020808501808196508360051b8101915082860160005b8581101561583d57828403895261582b848351614eb3565b98850198935090840190600101615813565b5091979650505050505050565b60a08152600061585d60a08301886157bc565b828103602084015261586f818861541e565b9050828103604084015261588381876157f5565b60608401959095525050608001529392505050565b60c0815260006158ab60c08301896157bc565b82810360208401526158bd818961541e565b905082810360408401526158d181886157f5565b60608401969096525050608081019290925260a0909101529392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008219821115615932576159326158f0565b500190565b60006020828403121561594957600080fd5b815167ffffffffffffffff81111561596057600080fd5b8201601f8101841361597157600080fd5b805161597f614d7382614d3d565b81815285602083850101111561599457600080fd5b611704826020830160208601614e83565b6000828210156159b7576159b76158f0565b500390565b6000602082840312156159ce57600080fd5b815165ffffffffffff81168114611a4e57600080fd5b6080815260006159f760808301876157bc565b8281036020840152615a09818761541e565b90508281036040840152615a1d81866157f5565b91505082606083015295945050505050565b84815260ff841660208201528260408201526080606082015260006113b56080830184614eb3565b85815260ff8516602082015283604082015260a060608201526000615a7f60a0830185614eb3565b8281036080840152615a918185614eb3565b98975050505050505050565b60006000198203615ab057615ab06158f0565b5060010190565b600060208284031215615ac957600080fd5b81518015158114611a4e57600080fd5b60008451615aeb818460208901614e83565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551615b27816001850160208a01614e83565b60019201918201528351615b42816002840160208801614e83565b0160020195945050505050565b60006101208b835260206001600160a01b038c1681850152816040850152615b798285018c6157bc565b91508382036060850152615b8d828b61541e565b915083820360808501528189518084528284019150828160051b850101838c0160005b83811015615bde57601f19878403018552615bcc838351614eb3565b94860194925090850190600101615bb0565b505086810360a0880152615bf2818c6157f5565b9450505050508560c08401528460e0840152828103610100840152615c178185614eb3565b9c9b505050505050505050505050565b6000816000190483118215151615615c4157615c416158f0565b500290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082615cab577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b601f821115612ee957600081815260208120601f850160051c81016020861015615cd75750805b601f850160051c820191505b8181101561243c57828155600101615ce3565b815167ffffffffffffffff811115615d1057615d10614cdd565b615d2481615d1e8454615721565b84615cb0565b602080601f831160018114615d595760008415615d415750858301515b600019600386901b1c1916600185901b17855561243c565b600085815260208120601f198616915b82811015615d8857888601518255948401946001909101908401615d69565b5085821015615da65787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fe476f7665726e6f723a2072656c617920726576657274656420776974686f7574206d657373616765a164736f6c634300080f000a",
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
