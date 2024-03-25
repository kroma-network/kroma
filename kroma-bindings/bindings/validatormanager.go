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

// AssetManagerConstructorParams is an auto generated low-level Go binding around an user-defined struct.
type AssetManagerConstructorParams struct {
	L2OutputOracle         common.Address
	AssetToken             common.Address
	Kgh                    common.Address
	KghManager             common.Address
	SecurityCouncil        common.Address
	MaxOutputFinalizations *big.Int
	BaseReward             *big.Int
	SlashingRateNumerator  *big.Int
	MinSlashingAmount      *big.Int
	MinRegisterAmount      *big.Int
	MinStartAmount         *big.Int
	UndelegationPeriod     *big.Int
}

// ValidatorManagerMetaData contains all meta data concerning the ValidatorManager contract.
var ValidatorManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"contractL2OutputOracle\",\"name\":\"_l2OutputOracle\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_assetToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC721\",\"name\":\"_kgh\",\"type\":\"address\"},{\"internalType\":\"contractIKGHManager\",\"name\":\"_kghManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_securityCouncil\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_maxOutputFinalizations\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_baseReward\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_slashingRateNumerator\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_minSlashingAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_minRegisterAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_minStartAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"_undelegationPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structAssetManager.ConstructorParams\",\"name\":\"_constructorParams\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_trustedValidator\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_commissionRateMinChangeSeconds\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_roundDurationSeconds\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_jailPeriodSeconds\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_jailThreshold\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"ChallengeRewardDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"kroShares\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"kghShares\",\"type\":\"uint128\"}],\"name\":\"KghBatchDelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"kroShares\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"kghShares\",\"type\":\"uint128\"}],\"name\":\"KghBatchUndelegationInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"kroInKgh\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"kroShares\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"kghShares\",\"type\":\"uint128\"}],\"name\":\"KghDelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"KghUndelegationFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"kroShares\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"kghShares\",\"type\":\"uint128\"}],\"name\":\"KghUndelegationInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"shares\",\"type\":\"uint128\"}],\"name\":\"KroDelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"KroUndelegationFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"shares\",\"type\":\"uint128\"}],\"name\":\"KroUndelegationInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"RewardClaimFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"RewardClaimInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"validatorReward\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"baseReward\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"boostedReward\",\"type\":\"uint128\"}],\"name\":\"RewardDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"loser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"oldCommissionRate\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"newCommissionRate\",\"type\":\"uint8\"}],\"name\":\"ValidatorCommissionRateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"expiresAt\",\"type\":\"uint128\"}],\"name\":\"ValidatorJailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"started\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"commissionRate\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"commissionMaxChangeRate\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"assets\",\"type\":\"uint128\"}],\"name\":\"ValidatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startsAt\",\"type\":\"uint256\"}],\"name\":\"ValidatorStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorUnjailed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ASSET_TOKEN\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASE_REWARD\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BOOSTED_REWARD_DENOM\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BOOSTED_REWARD_NUMERATOR\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COMMISSION_RATE_DENOM\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COMMISSION_RATE_MIN_CHANGE_SECONDS\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DECIMAL_OFFSET\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"JAIL_PERIOD_SECONDS\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"JAIL_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"KGH\",\"outputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"KGH_MANAGER\",\"outputs\":[{\"internalType\":\"contractIKGHManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2_ORACLE\",\"outputs\":[{\"internalType\":\"contractL2OutputOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_OUTPUT_FINALIZATIONS\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_REGISTER_AMOUNT\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_SLASHING_AMOUNT\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_START_AMOUNT\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROUND_DURATION_SECONDS\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECURITY_COUNCIL\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASHING_RATE_DENOM\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASHING_RATE_NUMERATOR\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TAX_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TAX_NUMERATOR\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRUSTED_VALIDATOR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNDELEGATION_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VKRO_PER_KGH\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"outputIndex\",\"type\":\"uint256\"}],\"name\":\"afterSubmitL2Output\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newCommissionRate\",\"type\":\"uint8\"}],\"name\":\"changeCommissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"checkSubmissionEligibility\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"assets\",\"type\":\"uint128\"}],\"name\":\"delegate\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"delegateKgh\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"delegateKghBatch\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizeClaimValidatorReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"finalizeUndelegate\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"finalizeUndelegateKgh\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"getCommissionMaxChangeRate\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"getCommissionRate\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getKghTotalBalance\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getKghTotalShareBalance\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"getKroTotalBalance\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"getKroTotalShareBalance\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"enumValidatorManager.ValidatorStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"getWeight\",\"outputs\":[{\"internalType\":\"uint120\",\"name\":\"\",\"type\":\"uint120\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"initClaimValidatorReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"shares\",\"type\":\"uint128\"}],\"name\":\"initUndelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"initUndelegateKgh\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"initUndelegateKghBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"jailExpiresAt\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextValidator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"noSubmissionCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"assets\",\"type\":\"uint128\"}],\"name\":\"previewDelegate\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"previewKghDelegate\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"previewKghUndelegate\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"shares\",\"type\":\"uint128\"}],\"name\":\"previewUndelegate\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"assets\",\"type\":\"uint128\"},{\"internalType\":\"uint8\",\"name\":\"commissionRate\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"commissionMaxChangeRate\",\"type\":\"uint8\"}],\"name\":\"registerValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"loser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"outputIndex\",\"type\":\"uint256\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startedValidatorCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startedValidatorTotalWeight\",\"outputs\":[{\"internalType\":\"uint120\",\"name\":\"\",\"type\":\"uint120\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"totalKghNum\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"totalKroAssets\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tryUnjail\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6102a06040523480156200001257600080fd5b506040516200893038038062008930833981016040819052620000359162000209565b85516001600160a01b0390811660a09081526020880151821660809081526040890151831660c090815260608a0151841660e0908152918a0151909316610100908152918901516001600160801b03908116610120908152938a01518116610140908152918a015181166101605291890151821661018052880151918801518892821691161115620001415760405162461bcd60e51b8152602060048201526044602482018190527f41737365744d616e616765723a206d696e20726567697374657220616d6f756e908201527f742073686f756c64206e6f7420657863656564206d696e20737461727420616d6064820152631bdd5b9d60e21b608482015260a40160405180910390fd5b6101208101516001600160801b039081166101a05261014082015181166101c052610160909101516101e0526001600160a01b039095166102005292841661022052908316610240528216610260521661028052506200037f565b60405161018081016001600160401b0381118282101715620001ce57634e487b7160e01b600052604160045260246000fd5b60405290565b80516001600160a01b0381168114620001ec57600080fd5b919050565b80516001600160801b0381168114620001ec57600080fd5b6000806000806000808688036102208112156200022557600080fd5b610180808212156200023657600080fd5b620002406200019c565b91506200024d89620001d4565b82526200025d60208a01620001d4565b60208301526200027060408a01620001d4565b60408301526200028360608a01620001d4565b60608301526200029660808a01620001d4565b6080830152620002a960a08a01620001f1565b60a0830152620002bc60c08a01620001f1565b60c0830152620002cf60e08a01620001f1565b60e0830152610100620002e4818b01620001f1565b90830152610120620002f88a8201620001f1565b908301526101406200030c8a8201620001f1565b8184015250610160808a015181840152508197506200032d818a01620001d4565b96505050620003406101a08801620001f1565b9350620003516101c08801620001f1565b9250620003626101e08801620001f1565b9150620003736102008801620001f1565b90509295509295509295565b60805160a05160c05160e05161010051610120516101405161016051610180516101a0516101c0516101e051610200516102205161024051610260516102805161838a620005a66000396000818161069b0152615548015260008181610985015261558d0152600081816106c20152611a54015260008181610b440152611fcd0152600081816106740152611aa80152600081816107ae01528181610e360152818161301b015261367401526000818161064d015281816115140152818161294201526146620152600081816104dc015281816110c50152818161149e015281816115cc0152818161266d0152612d450152600081816104b501528181615a690152615aa30152600081816109380152615a3e0152600081816105b1015281816169e101528181616a1a015261783c015260008181610be101526151e201526000818161061e0152615c7c015260008181610ad901528181611170015281816117a001528181612b9d01528181613aac015261496c015260008181610732015281816116d40152818161316001526144920152600081816103ce015281816119bb01528181612335015281816138f8015281816139c601528181613bf201528181614fac0152818161503d015281816150c101528181615255015281816152fc01528181615481015281816156d001526157aa015260008181610b0001528181611014015281816134320152818161387001528181614bfc0152615c5a015261838a6000f3fe608060405234801561001057600080fd5b50600436106103c45760003560e01c8063842d0d3b116101ff578063b7aa324a1161011a578063de313284116100ad578063e0cc26a21161007c578063e0cc26a214610b9a578063e74f823914610bc9578063e7816b7f14610bdc578063eb2ad8cb14610c0357600080fd5b8063de31328414610b22578063de7d4d6a14610b35578063dea1525414610b3f578063dff221b514610b6657600080fd5b8063be119347116100e9578063be11934714610a68578063cf368e8c14610a7b578063d1e288c114610ad4578063d706200514610afb57600080fd5b8063b7aa324a146109fc578063b91b27231461043f578063b9551f8214610a0f578063b9f6131b14610a5557600080fd5b8063970531c111610192578063a85120e411610161578063a85120e41461095a578063a93b7ad41461096d578063abeba44914610980578063ac6c5251146109a757600080fd5b8063970531c1146108ee578063a30c5f3014610920578063a4cf0b2f14610933578063a51c9ace1461043f57600080fd5b80638cdfe8a9116101ce5780638cdfe8a9146108805780638ee4b60214610893578063913f1a9f146108a6578063960a0893146108db57600080fd5b8063842d0d3b1461080157806388576dc914610847578063891aab741461085a5780638c0903501461086d57600080fd5b806331d8e007116102ef57806356576b5b116102825780637533f901116102515780637533f901146107a95780637cd68cd7146107de5780637d2243b4146107f157806382dae3aa146107f957600080fd5b806356576b5b1461072d5780635dd0293b14610754578063631bda01146107675780636b9ffeac1461077757600080fd5b80633ee4d4a3116102be5780633ee4d4a31461066f57806342223ae9146106965780634cca5e6c146106bd57806354fd4d50146106e457600080fd5b806331d8e0071461060657806336086417146106195780633a549046146106405780633bcebcd81461064857600080fd5b80631796e52e11610367578063209a969411610336578063209a96941461059957806322009af6146105ac5780632328bf42146105d357806330ccebb5146105e657600080fd5b80631796e52e146104d75780631a5deb4a146104fe5780631edbc580146105075780631f86f4f11461056657600080fd5b80630bc0b881116103a35780630bc0b88114610437578063110d60691461043f578063150b7a0214610447578063176a86d0146104b057600080fd5b80621c2ff6146103c9578063072df4cb1461040d5780630763fa7e14610417575b600080fd5b6103f07f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b610415610c20565b005b61041f602881565b6040516001600160801b039091168152602001610404565b610415610d11565b61041f606481565b61047f610455366004617c26565b7f150b7a020000000000000000000000000000000000000000000000000000000095945050505050565b6040517fffffffff000000000000000000000000000000000000000000000000000000009091168152602001610404565b61041f7f000000000000000000000000000000000000000000000000000000000000000081565b61041f7f000000000000000000000000000000000000000000000000000000000000000081565b61041f6103e881565b6000805468010000000000000000900463ffffffff1681526001602081905260409091200154600160801b90046effffffffffffffffffffffffffffff165b6040516effffffffffffffffffffffffffffff9091168152602001610404565b610579610574366004617cc5565b611089565b604080516001600160801b03938416815292909116602083015201610404565b61041f6105a7366004617d06565b611278565b61041f7f000000000000000000000000000000000000000000000000000000000000000081565b6104156105e1366004617d3f565b61128d565b6105f96105f4366004617d5c565b61141e565b6040516104049190617da8565b610579610614366004617de9565b611590565b6103f07f000000000000000000000000000000000000000000000000000000000000000081565b6103f06119a3565b61041f7f000000000000000000000000000000000000000000000000000000000000000081565b6103f07f000000000000000000000000000000000000000000000000000000000000000081565b61041f7f000000000000000000000000000000000000000000000000000000000000000081565b61041f7f000000000000000000000000000000000000000000000000000000000000000081565b6107206040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516104049190617e9d565b6103f07f000000000000000000000000000000000000000000000000000000000000000081565b610415610762366004617d06565b611aca565b61041f68056bc75e2d6310000081565b61041f610785366004617d5c565b6001600160a01b03166000908152600360205260409020546001600160801b031690565b6107d07f000000000000000000000000000000000000000000000000000000000000000081565b604051908152602001610404565b6104156107ec366004617cc5565b611c76565b610415611d99565b61041f601481565b61041f61080f366004617eee565b6001600160a01b039182166000908152600360209081526040808320939094168252600b90920190915220546001600160801b031690565b610415610855366004617f32565b611ef6565b610415610868366004617d5c565b61232a565b61041f61087b366004617eee565b61254b565b61041f61088e366004617cc5565b612592565b6104156108a1366004617f4d565b6125dc565b61041f6108b4366004617d5c565b6001600160a01b03166000908152600360205260409020600101546001600160801b031690565b61041f6108e9366004617d06565b6129d9565b61041f6108fc366004617d5c565b6001600160a01b03166000908152600660205260409020546001600160801b031690565b61041561092e366004617de9565b6129e5565b61041f7f000000000000000000000000000000000000000000000000000000000000000081565b61041f610968366004617d06565b612d0a565b61041f61097b366004617d5c565b612ea7565b61041f7f000000000000000000000000000000000000000000000000000000000000000081565b6105466109b5366004617d5c565b6001600160a01b031660009081526002602090815260408083205463ffffffff168352600191829052909120015461010090046effffffffffffffffffffffffffffff1690565b61041f610a0a366004617d5c565b612ebc565b610a43610a1d366004617d5c565b6001600160a01b0316600090815260036020526040902060060154610100900460ff1690565b60405160ff9091168152602001610404565b61041f610a63366004617d5c565b6134b3565b610415610a76366004617f92565b6138ed565b610579610a89366004617fab565b6001600160a01b039283166000908152600360209081526040808320949095168252600c9093018352838120918152915220546001600160801b0380821692600160801b9092041690565b6103f07f000000000000000000000000000000000000000000000000000000000000000081565b6103f07f000000000000000000000000000000000000000000000000000000000000000081565b61041f610b30366004617fab565b613a70565b61041f620f424081565b61041f7f000000000000000000000000000000000000000000000000000000000000000081565b610a43610b74366004617d5c565b6001600160a01b03166000908152600360205260409020600a0154610100900460ff1690565b610a43610ba8366004617d5c565b6001600160a01b031660009081526003602052604090206006015460ff1690565b610415610bd7366004617fab565b613bf0565b61041f7f000000000000000000000000000000000000000000000000000000000000000081565b610c0b613da4565b60405163ffffffff9091168152602001610404565b6004610c2b3361141e565b6006811115610c3c57610c3c617d79565b14610cb45760405162461bcd60e51b815260206004820152603660248201527f56616c696461746f724d616e616765723a2076616c696461746f72207374617260448201527f7420636f6e646974696f6e206973206e6f74206d65740000000000000000000060648201526084015b60405180910390fd5b336000818152600360205260409020610cda9190610cd190613dc7565b60009190613dfa565b60405142815233907fe8e4e936783175825bcf08ad234ab704ad447aeda363141c88312a07a729d0679060200160405180910390a2565b33600090815260036020526040902060078101546005909101906001600160801b0316610da65760405162461bcd60e51b815260206004820152603660248201527f41737365744d616e616765723a206e6f2070656e64696e672076616c6964617460448201527f6f72207265776172647320746f2066696e616c697a65000000000000000000006064820152608401610cab565b600081600301805480602002602001604051908101604052809291908181526020018280548015610df657602002820191906000526020600020905b815481526020019060010190808311610de2575b5050505050905060008060018351610e0e919061801b565b90505b6000838281518110610e2557610e25618032565b60200260200101511115610f3857427f0000000000000000000000000000000000000000000000000000000000000000848381518110610e6757610e67618032565b6020026020010151610e799190618061565b11610f2957836004016000848381518110610e9657610e96618032565b6020026020010151815260200190815260200160002060009054906101000a90046001600160801b031682019150836004016000848381518110610edc57610edc618032565b6020026020010151815260200190815260200160002060006101000a8154906001600160801b030219169055836003018181548110610f1d57610f1d618032565b60009182526020822001555b8015610f385760001901610e11565b506000816001600160801b031611610fb85760405162461bcd60e51b815260206004820152603560248201527f41737365744d616e616765723a206e6f2070656e64696e67207265776172642060448201527f636c61696d20746f2066696e616c697a652079657400000000000000000000006064820152608401610cab565b60028301546001600160801b0380831691161015610fe0575060028201546001600160801b03165b6002830180546fffffffffffffffffffffffffffffffff1981166001600160801b03918216849003821617909155611046907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316903390841661436d565b6040516001600160801b038216815233907f668550d283aec3ba805f3e7c44d0bd95c1f847946fc605a7222d53394ca9e0509060200160405180910390a2505050565b60008083336001600160a01b03821614806110f257506001600160a01b0381166000908152600360205260409020600201546001600160801b037f00000000000000000000000000000000000000000000000000000000000000008116600160801b9092041610155b61113e5760405162461bcd60e51b815260206004820152601f60248201527f41737365744d616e616765723a205661756c7420697320696e616374697665006044820152606401610cab565b6040517fa1d92445000000000000000000000000000000000000000000000000000000008152600481018590526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063a1d9244590602401602060405180830381865afa1580156111bf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111e39190618079565b905060006111f187836129d9565b905060006111fe88612ea7565b905061120d8888858585614439565b604080518881526001600160801b0385811660208301528481168284015283166060820152905133916001600160a01b038b16917f5b545a208eb5f51c3900b6fbf02a83cfdafcfbd2bae035a706827fceef97be8b9181900360800190a39097909650945050505050565b600061128483836145c3565b90505b92915050565b3360009081526003602052604090206005016001600160801b038216158015906112cc575080546001600160801b03600160801b909104811690831611155b61133e5760405162461bcd60e51b815260206004820152602560248201527f41737365744d616e616765723a20496e76616c69642072657761726420746f2060448201527f636c61696d0000000000000000000000000000000000000000000000000000006064820152608401610cab565b80546001600160801b03600160801b808304821685900382160291811691909117825542600081815260048401602090815260408083208054808716890187166fffffffffffffffffffffffffffffffff19918216179091556002870180548088168a019097169690911695909517909455600380860180546001818101835591855283852001949094553380845291529290206113dc929161463e565b6040516001600160801b038316815233907f8ba9d52d538174baf05e0437c09555c022591dc617b7458fc02f53a1de9da4909060200160405180910390a25050565b6001600160a01b0381166000908152600360205260408120600a015460ff1661144957506000919050565b6001600160a01b0382166000908152600660205260409020546001600160801b03161561147857506002919050565b6001600160a01b0382166000908152600360205260409020600201546001600160801b037f00000000000000000000000000000000000000000000000000000000000000008116600160801b9092041610156114d657506001919050565b6001600160a01b0382166000908152600260209081526040808320546003909252909120600181015463ffffffff9092161515916001600160801b037f0000000000000000000000000000000000000000000000000000000000000000811692600160801b909204169061154990613dc7565b6115539190618096565b6001600160801b0316101561157957806115705750600392915050565b50600592915050565b806115875750600492915050565b50600692915050565b60008084336001600160a01b03821614806115f957506001600160a01b0381166000908152600360205260409020600201546001600160801b037f00000000000000000000000000000000000000000000000000000000000000008116600160801b9092041610155b6116455760405162461bcd60e51b815260206004820152601f60248201527f41737365744d616e616765723a205661756c7420697320696e616374697665006044820152606401610cab565b836116b85760405162461bcd60e51b815260206004820152602360248201527f41737365744d616e616765723a2063616e6e6f742064656c656761746520302060448201527f4b474800000000000000000000000000000000000000000000000000000000006064820152608401610cab565b6000806116c488612ea7565b90506000805b8781101561192a577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166342842e0e33308c8c8681811061171557611715618032565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e088901b1681526001600160a01b03958616600482015294909316602485015250602090910201356044820152606401600060405180830381600087803b15801561178457600080fd5b505af1158015611798573d6000803e3d6000fd5b5050505060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663a1d924458b8b858181106117df576117df618032565b905060200201356040518263ffffffff1660e01b815260040161180491815260200190565b602060405180830381865afa158015611821573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118459190618079565b905060006118538c836129d9565b90506040518060400160405280826001600160801b03168152602001866001600160801b0316815250600360008e6001600160a01b03166001600160a01b03168152602001908152602001600020600c016000336001600160a01b03166001600160a01b0316815260200190815260200160002060000160008d8d878181106118de576118de618032565b6020908102929092013583525081810192909252604001600020825192909101516001600160801b03908116600160801b029216919091179055949094019391909101906001016116ca565b5061193587836180be565b915061194489888386866146cc565b336001600160a01b0316896001600160a01b03167f32d1388f6c67f8101e5b61b8d17b70b87d15213fb6e6ebef744aaedfa4a41d1f8a8a878760405161198d94939291906180ed565b60405180910390a3509097909650945050505050565b6005546000906001600160a01b031615611aa55760007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166380446bd26040518163ffffffff1660e01b8152600401602060405180830381865afa158015611a17573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a3b919061815d565b9050804210611a94576000611a50824261801b565b90507f00000000000000000000000000000000000000000000000000000000000000006001600160801b0316811115611a92576001600160a01b039250505090565b505b50506005546001600160a01b031690565b507f000000000000000000000000000000000000000000000000000000000000000090565b6000816001600160801b0316118015611b1657506001600160a01b0382166000908152600360209081526040808320338452600b019091529020546001600160801b0390811690821611155b611b885760405162461bcd60e51b815260206004820152603460248201527f41737365744d616e616765723a20496e76616c696420616d6f756e74206f662060448201527f73686172657320746f20756e64656c65676174650000000000000000000000006064820152608401610cab565b6000611b948383611278565b90506000816001600160801b031611611c155760405162461bcd60e51b815260206004820152602760248201527f41737365744d616e616765723a2063616e6e6f7420756e64656c65676174652060448201527f30206173736574000000000000000000000000000000000000000000000000006064820152608401610cab565b611c2183338385614788565b604080516001600160801b0380841682528416602082015233916001600160a01b038616917fcc7aa86d5cfd922cf42cb1b77d79a7e530bbaa3264669ca2145b564ecc1bf769910160405180910390a3505050565b6001600160a01b0382166000908152600360209081526040808320338452600c0182528083208484529091528120546001600160801b0380821692600160801b9092041690819003611d305760405162461bcd60e51b815260206004820152602d60248201527f41737365744d616e616765723a204e6f2073686172657320666f72207468652060448201527f676976656e20746f6b656e4964000000000000000000000000000000000000006064820152608401610cab565b611d3d843385858561484d565b604080518481526001600160801b038481166020830152831681830152905133916001600160a01b038716917f4d45e37032bac00ed9fe936ed54d9751ff10d249aec97a5dcd23a769a6be31fc9181900360600190a350505050565b6002611da43361141e565b6006811115611db557611db5617d79565b14611e025760405162461bcd60e51b815260206004820152601d60248201527f56616c696461746f724d616e616765723a206e6f7420696e206a61696c0000006044820152606401610cab565b33600090815260066020526040902054426001600160801b039091161115611e925760405162461bcd60e51b815260206004820152602d60248201527f56616c696461746f724d616e616765723a206a61696c20706572696f6420686160448201527f73206e6f7420656c6173706564000000000000000000000000000000000000006064820152608401610cab565b33600081815260066020526040902080546fffffffffffffffffffffffffffffffff19169055611ec190614b17565b6040513381527f9390b453426557da5ebdc31f19a37753ca04addf656d32f35232211bb2af3f199060200160405180910390a1565b6001611f013361141e565b6006811115611f1257611f12617d79565b11611fab5760405162461bcd60e51b815260206004820152604560248201527f56616c696461746f724d616e616765723a2063616e6e6f74206368616e67652060448201527f636f6d6d697373696f6e2072617465206f6620696e6163746976652076616c6960648201527f6461746f72000000000000000000000000000000000000000000000000000000608482015260a401610cab565b33600090815260036020526040902060068101546005909101904290612001907f0000000000000000000000000000000000000000000000000000000000000000906201000090046001600160801b0316618176565b6001600160801b031611156120a45760405162461bcd60e51b815260206004820152604760248201527f56616c696461746f724d616e616765723a206d696e206368616e67652073656360448201527f6f6e6473206f6620636f6d6d697373696f6e207261746520686173206e6f742060648201527f656c617073656400000000000000000000000000000000000000000000000000608482015260a401610cab565b606460ff831611156121455760405162461bcd60e51b8152602060048201526044602482018190527f56616c696461746f724d616e616765723a20746865206d61782076616c756520908201527f6f6620636f6d6d697373696f6e207261746520686173206265656e206578636560648201527f6564656400000000000000000000000000000000000000000000000000000000608482015260a401610cab565b600181015460ff9081169083168190036121c75760405162461bcd60e51b815260206004820152603160248201527f56616c696461746f724d616e616765723a2063616e6e6f74206368616e67652060448201527f746f207468652073616d652076616c75650000000000000000000000000000006064820152608401610cab565b60008160ff168460ff1611156121e8576121e182856181a1565b90506121f5565b6121f284836181a1565b90505b600183015460ff610100909104811690821611156122a15760405162461bcd60e51b815260206004820152604660248201527f56616c696461746f724d616e616765723a206d6178206368616e67652072617460448201527f65206f6620636f6d6d697373696f6e207261746520686173206265656e20657860648201527f6365656465640000000000000000000000000000000000000000000000000000608482015260a401610cab565b60018301805460ff8681167fffffffffffffffffffffffffffff00000000000000000000000000000000ff00909216821762010000426001600160801b031602179092556040805133815292851660208401528201527fc0b29b9b824f7a62d93fde5832bb8307fd62594d0a08d96d533407a0a147ec489060600160405180910390a150505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146123c85760405162461bcd60e51b815260206004820152603c60248201527f56616c696461746f724d616e616765723a204f6e6c79204c324f75747075744f60448201527f7261636c652063616e2063616c6c20746869732066756e6374696f6e000000006064820152608401610cab565b60006123d26119a3565b90506001600160a01b038082161461249257806001600160a01b0316826001600160a01b0316146124925760405162461bcd60e51b8152602060048201526044602482018190527f56616c696461746f724d616e616765723a206f6e6c7920746865206e65787420908201527f73656c65637465642076616c696461746f722063616e207375626d6974206f7560648201527f7470757400000000000000000000000000000000000000000000000000000000608482015260a401610cab565b600661249d8361141e565b60068111156124ae576124ae617d79565b146125475760405162461bcd60e51b815260206004820152604960248201527f56616c696461746f724d616e616765723a2076616c696461746f722073686f7560448201527f6c6420736174697366792074686520636f6e646974696f6e20746f207375626d60648201527f6974206f75747075740000000000000000000000000000000000000000000000608482015260a401610cab565b5050565b6001600160a01b0380831660009081526003602090815260408083209385168352600b9093019052908120546001600160801b031661258a8482611278565b949350505050565b6001600160a01b0382166000908152600360209081526040808320338452600c018252808320848452909152812054611284908490600160801b90046001600160801b0316614b85565b60006125e73361141e565b60068111156125f8576125f8617d79565b1461266b5760405162461bcd60e51b815260206004820152602d60248201527f56616c696461746f724d616e616765723a20616c726561647920696e6974696160448201527f7465642076616c696461746f72000000000000000000000000000000000000006064820152608401610cab565b7f00000000000000000000000000000000000000000000000000000000000000006001600160801b0316836001600160801b0316101561273a5760405162461bcd60e51b8152602060048201526044602482018190527f56616c696461746f724d616e616765723a206e65656420746f20726567697374908201527f65722077697468206174206c65617374206d696e20726567697374657220616d60648201527f6f756e7400000000000000000000000000000000000000000000000000000000608482015260a401610cab565b606460ff831611156127db5760405162461bcd60e51b8152602060048201526044602482018190527f56616c696461746f724d616e616765723a20746865206d61782076616c756520908201527f6f6620636f6d6d697373696f6e207261746520686173206265656e206578636560648201527f6564656400000000000000000000000000000000000000000000000000000000608482015260a401610cab565b606460ff8216111561287b5760405162461bcd60e51b815260206004820152605460248201527f56616c696461746f724d616e616765723a20746865206d61782076616c75652060448201527f6f6620636f6d6d697373696f6e2072617465206d6178206368616e676520726160648201527f746520686173206265656e206578636565646564000000000000000000000000608482015260a401610cab565b336000818152600360205260408120600a810180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556006810180546001600160801b03421662010000027fffffffffffffffffffffffffffff00000000000000000000000000000000ffff60ff888116610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000909416908a16179290921791909116179055916129379181908790614bcd565b506001600160801b037f000000000000000000000000000000000000000000000000000000000000000081169085161080159061297b5761297b33610cd184613dc7565b6040805160ff8087168252851660208201526001600160801b038716918101919091528115159033907f04ba0c4d7cbac9138f7b73ec9fef796e4ad320bf5fb204f080f81fd59c2d48b9906060015b60405180910390a35050505050565b60006112848383614ce3565b80612a585760405162461bcd60e51b815260206004820152602560248201527f41737365744d616e616765723a2063616e6e6f7420756e64656c65676174652060448201527f30204b47480000000000000000000000000000000000000000000000000000006064820152608401610cab565b6001600160a01b0383166000908152600360209081526040808320338452600c01909152812090808080805b86811015612c9f5760008660008a8a85818110612aa357612aa3618032565b602090810292909201358352508101919091526040016000908120546001600160801b0316915087818b8b86818110612ade57612ade618032565b90506020020135815260200190815260200160002060000160109054906101000a90046001600160801b03169050806001600160801b0316600003612b8b5760405162461bcd60e51b815260206004820152602d60248201527f41737365744d616e616765723a204e6f2073686172657320666f72207468652060448201527f676976656e20746f6b656e4964000000000000000000000000000000000000006064820152608401610cab565b95810195948501946001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001663a1d924458b8b86818110612bd457612bd4618032565b905060200201356040518263ffffffff1660e01b8152600401612bf991815260200190565b602060405180830381865afa158015612c16573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c3a9190618079565b85019450612c608b8b8b86818110612c5457612c54618032565b90506020020135612592565b840193508760008b8b86818110612c7957612c79618032565b602090810292909201358352508101919091526040016000908120555050600101612a84565b50612caf88888885888887614d4d565b336001600160a01b0316886001600160a01b03167fcb83bfd06f7180fe00601135ea03ce74be1a309cbdfd2e5a467bc7c42328f6c389898888604051612cf894939291906180ed565b60405180910390a35050505050505050565b600082336001600160a01b0382161480612d7257506001600160a01b0381166000908152600360205260409020600201546001600160801b037f00000000000000000000000000000000000000000000000000000000000000008116600160801b9092041610155b612dbe5760405162461bcd60e51b815260206004820152601f60248201527f41737365744d616e616765723a205661756c7420697320696e616374697665006044820152606401610cab565b6000836001600160801b031611612e3d5760405162461bcd60e51b815260206004820152602560248201527f41737365744d616e616765723a2063616e6e6f742064656c656761746520302060448201527f61737365740000000000000000000000000000000000000000000000000000006064820152608401610cab565b6000612e4c8533866001614bcd565b604080516001600160801b0380881682528316602082015291925033916001600160a01b038816917f334cabe84b7338f2bdd62070c02f24ffbcc7735e46f425fa401db349717e1328910160405180910390a3949350505050565b60006112878268056bc75e2d63100000614edd565b6001600160a01b0381166000908152600360209081526040808320338452600c018252808320600181018054835181860281018601909452808452919385939290830182828015612f2c57602002820191906000526020600020905b815481526020019060010190808311612f18575b505050505090506000815111612faa5760405162461bcd60e51b815260206004820152602c60248201527f41737365744d616e616765723a206e6f20756e64656c65676174696f6e20726560448201527f71756573747320657869737400000000000000000000000000000000000000006064820152608401610cab565b6001600160a01b0384166000908152600360205260408120600481015483519192600160801b9091046001600160801b03161515918190819081908190612ff39060019061801b565b90505b600088828151811061300a5761300a618032565b602002602001015111156132e157427f000000000000000000000000000000000000000000000000000000000000000089838151811061304c5761304c618032565b602002602001015161305e9190618061565b116132d25785156130d15760008960030160008a848151811061308357613083618032565b60209081029190910181015182528181019290925260409081016000208151808301909252546001600160801b03808216808452600160801b90920416919092018190529601959490940193505b60008960020160008a84815181106130eb576130eb618032565b6020026020010151815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561314b57602002820191906000526020600020905b815481526020019060010190808311613137575b5050505050905060005b8151811015613236577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166342842e0e30338585815181106131a1576131a1618032565b60209081029190910101516040517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b1681526001600160a01b0393841660048201529290911660248301526044820152606401600060405180830381600087803b15801561321357600080fd5b505af1158015613227573d6000803e3d6000fd5b50505050806001019050613155565b508960030160008a848151811061324f5761324f618032565b60209081029190910181015182528101919091526040016000908120819055895160028c0191908b908590811061328857613288618032565b6020026020010151815260200190815260200160002060006132aa9190617bdf565b8960010182815481106132bf576132bf618032565b6000918252602082200155506001909101905b80156132e15760001901612ff6565b50600081116133585760405162461bcd60e51b815260206004820152603560248201527f41737365744d616e616765723a206e6f2070656e64696e67204b474820756e6460448201527f656c65676174696f6e20746f2066696e616c697a6500000000000000000000006064820152608401610cab565b841561345d5760038601546004870154600091613384916001600160801b038881169281169116614f24565b600388015460048901549192506000916133b8916001600160801b0388811692600160801b92839004821692900416614f24565b60038901805460048b018054600160801b6001600160801b0380851689900381166fffffffffffffffffffffffffffffffff199586168117839004821688900382168302179095558185168c9003851691909316811783900484168a9003841690920290911790559201935061345d916001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691503390851661436d565b6040516001600160801b038316815233906001600160a01b038c16907fb8e63154a3a976a585c24f01fc46ead9ec9a0038ca84a74e3090e938c61fabe49060200160405180910390a35098975050505050505050565b6001600160a01b038116600090815260036020526040812060048101546001600160801b031661354b5760405162461bcd60e51b815260206004820152602b60248201527f41737365744d616e616765723a204e6f2070656e64696e67207368617265732060448201527f746f2066696e616c697a650000000000000000000000000000000000000000006064820152608401610cab565b6001600160a01b0383166000908152600360209081526040808320338452600b0182528083206001810180548351818602810186019094528084529194939091908301828280156135bb57602002820191906000526020600020905b8154815260200190600101908083116135a7575b5050505050905060008151116136395760405162461bcd60e51b815260206004820152602c60248201527f41737365744d616e616765723a206e6f20756e64656c65676174696f6e20726560448201527f71756573747320657869737400000000000000000000000000000000000000006064820152608401610cab565b60008060006001845161364c919061801b565b90505b600084828151811061366357613663618032565b6020026020010151111561377657427f00000000000000000000000000000000000000000000000000000000000000008583815181106136a5576136a5618032565b60200260200101516136b79190618061565b11613767578460020160008583815181106136d4576136d4618032565b6020026020010151815260200190815260200160002060009054906101000a90046001600160801b03168201915084600201600085838151811061371a5761371a618032565b6020026020010151815260200190815260200160002060006101000a8154906001600160801b03021916905584600101818154811061375b5761375b618032565b60009182526020822001555b8015613776576000190161364f565b506000816001600160801b0316116137f65760405162461bcd60e51b815260206004820152603560248201527f41737365744d616e616765723a206e6f2070656e64696e67204b524f20756e6460448201527f656c65676174696f6e20746f2066696e616c697a6500000000000000000000006064820152608401610cab565b6003850154600486015461381a916001600160801b03848116929181169116614f24565b6003860180546001600160801b0380821684900381166fffffffffffffffffffffffffffffffff19928316179092556004880180548084168690038416921691909117905590925061389a906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016903390851661436d565b6040516001600160801b038316815233906001600160a01b038916907f75cecc4cc21f0ebf6c86948f3a6a9bb934c49e4db83473fb8582ea706d1359149060200160405180910390a35095945050505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461398b5760405162461bcd60e51b815260206004820152603c60248201527f56616c696461746f724d616e616765723a204f6e6c79204c324f75747075744f60448201527f7261636c652063616e2063616c6c20746869732066756e6374696f6e000000006064820152608401610cab565b613993614fa7565b506040517fb0ea09a8000000000000000000000000000000000000000000000000000000008152600481018290526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063b0ea09a890602401602060405180830381865afa158015613a15573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613a3991906181c4565b6005549091506001600160a01b0390811690821603613a6057613a5b81614b17565b613a68565b613a68615517565b61254761568d565b6040517fa1d924450000000000000000000000000000000000000000000000000000000081526004810182905260009081906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a1d9244590602401602060405180830381865afa158015613af3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613b179190618079565b6001600160a01b0386811660009081526003602090815260408083209389168352600c90930181528282208783529052908120549192509068056bc75e2d6310000090613b75908890600160801b90046001600160801b0316612592565b613b7f9190618096565b6001600160a01b038088166000908152600360209081526040808320938a168352600c9093018152828220888352905290812054919250908390613bcd9089906001600160801b0316611278565b613bd79190618096565b9050613be38183618176565b93505050505b9392505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639e45e8f46040518163ffffffff1660e01b8152600401602060405180830381865afa158015613c4e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613c7291906181c4565b6001600160a01b0316336001600160a01b031614613cf85760405162461bcd60e51b815260206004820152603360248201527f41737365744d616e616765723a204f6e6c7920436f6c6f737365756d2063616e60448201527f2063616c6c20746869732066756e6374696f6e000000000000000000000000006064820152608401610cab565b6001600160a01b038316600090815260036020526040812090613d1d82826001615921565b600084815260046020526040902080546fffffffffffffffffffffffffffffffff19166001600160801b0383161790559050613d5b8583600161463e565b6040516001600160801b03821681526001600160a01b0380861691908716907fbfeaf055e3cc2126fdbf006eda97657a7a8f82248db4159264060f31dfa2e2d0906020016129ca565b60008054613dc29063ffffffff6401000000008204811691166181e1565b905090565b600581015481546000916001600160801b03600160801b8204811692613df09282169116618176565b6112879190618176565b60006040518060e00160405280846001600160a01b03168152602001600063ffffffff168152602001600063ffffffff168152602001600063ffffffff168152602001600115158152602001836effffffffffffffffffffffffffffff168152602001836effffffffffffffffffffffffffffff16815250905083600001600081819054906101000a900463ffffffff168092919060010191906101000a81548163ffffffff021916908363ffffffff1602179055505060008460000160009054906101000a900463ffffffff169050818560010160008363ffffffff1663ffffffff16815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160000160146101000a81548163ffffffff021916908363ffffffff16021790555060408201518160000160186101000a81548163ffffffff021916908363ffffffff160217905550606082015181600001601c6101000a81548163ffffffff021916908363ffffffff16021790555060808201518160010160006101000a81548160ff02191690831515021790555060a08201518160010160016101000a8154816effffffffffffffffffffffffffffff02191690836effffffffffffffffffffffffffffff16021790555060c08201518160010160106101000a8154816effffffffffffffffffffffffffffff02191690836effffffffffffffffffffffffffffff16021790555090505080856002016000866001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff1602179055508460000160089054906101000a900463ffffffff1663ffffffff166000036140d557845463ffffffff90911668010000000000000000027fffffffffffffffffffffffffffffffffffffffff00000000ffffffffffffffff90911617909355505050565b845468010000000000000000900463ffffffff165b63ffffffff808216600090815260018089016020526040822090810180546effffffffffffffffffffffffffffff600160801b80830482168b01909116027fff000000000000000000000000000000ffffffffffffffffffffffffffffffff90911617905580549092600160c01b9091041690036142045763ffffffff838116600081815260018a016020526040902080547fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff1674010000000000000000000000000000000000000000938616939093029290921790915581547fffffffff00000000ffffffffffffffffffffffffffffffffffffffffffffffff16600160c01b9091021781556141fb8784615e1f565b50505050505050565b8054600160e01b900463ffffffff166000036142da5763ffffffff838116600081815260018a8101602052604090912080547fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff1674010000000000000000000000000000000000000000948716949094029390931783559190910180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905581547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16600160e01b9091021781556141fb8784615e1f565b805463ffffffff600160e01b8204811660009081526001808b016020526040808320820154600160c01b909504909316825291902001546effffffffffffffffffffffffffffff600160801b9283900481169290910416111561434d578054600160e01b900463ffffffff16915061435f565b8054600160c01b900463ffffffff1691505b506140ea565b505050505050565b6040516001600160a01b0383166024820152604481018290526144349084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261601e565b505050565b6001600160a01b038581166000908152600360205260409081902090517f42842e0e00000000000000000000000000000000000000000000000000000000815233600482015230602482015260448101879052909182917f0000000000000000000000000000000000000000000000000000000000000000909116906342842e0e90606401600060405180830381600087803b1580156144d857600080fd5b505af11580156144ec573d6000803e3d6000fd5b5050825460018085018054600160801b8082046001600160801b039081168d01811682028082169382169390931790940184166fffffffffffffffffffffffffffffffff19928316179092558284168b018316938116841782900483168a018316820290931786556002860180548084168a01841694169390931790925560408051808201825289831681528883166020808301918252336000908152600c8b0182528481208f8252909152929092209051915183169093029082161790915587161591506141fb9050576141fb8783600061463e565b6001600160a01b038216600090815260036020526040812054611284906001600160801b03166145f4906001618176565b6001600160a01b038516600090815260036020526040902054620f424090600160801b90046001600160801b03165b61462d9190618176565b6001600160801b0385169190614f24565b600061464983613dc7565b90508180156146a2575060018301546001600160801b037f000000000000000000000000000000000000000000000000000000000000000081169161469791600160801b9091041683618096565b6001600160801b0316105b156146b8576146b2600085616106565b506146c6565b6146c46000858361629b565b505b50505050565b6001600160a01b03851660009081526003602052604090208054600182018054600160801b8082046001600160801b03908116890181168202808216938216939093178a0181166fffffffffffffffffffffffffffffffff199384161790935582841688018316938216841781900483168701831602909217835560028301805480831686018316931692909217909155841615614365576001600160a01b03861660009081526003602052604081206143659188919061463e565b6001600160a01b03808516600081815260036020908152604080832080546001600160801b03600160801b80830482168a9003821602918116919091178255958916808552600b82019093529220805480861687900386166fffffffffffffffffffffffffffffffff1991821617909155825480861688900390951694169390931781559103614836576002810180546001600160801b03600160801b808304821687900382160291161790555b61484181848461657a565b6146c48582600161463e565b6001600160a01b0385166000908152600360205260408120906148708785611278565b9050600061487e8887612592565b60408051600180825281830190925291925060009182916020808301908036833701905050905087816000815181106148b9576148b9618032565b60209081029190910181019190915285546001600160801b03808216600160801b9283900482168b900382169092029190911787556002870180546fffffffffffffffffffffffffffffffff1981169083168a90039092169190911790556001600160a01b038a81166000908152600c8801835260408082208c835290935282812081905591517fa1d92445000000000000000000000000000000000000000000000000000000008152600481018b90527f00000000000000000000000000000000000000000000000000000000000000009091169063a1d9244590602401602060405180830381865afa1580156149b5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906149d99190618079565b86546001600160801b0380821688900381166fffffffffffffffffffffffffffffffff19928316178955600189018054600160801b80820484168690038416028084169184169190911760001901831690841617905560058901805480831689900368056bc75e2d63100000018316931692909217909155604080518082019091527ffffffffffffffffffffffffffffffffffffffffffffffffa9438a1d29cf0000087019550919250828703916000918190614a99908d16858b614f24565b6001600160801b03168152602001614ac587898d6001600160801b0316614f249092919063ffffffff16565b6001600160801b031690529050614adf888584888561668b565b50505060008285614af09190618176565b6001600160801b03161115614b0b57614b0b8a86600161463e565b50505050505050505050565b6001600160a01b0381166000908152600360205260409020600a0154610100900460ff1615614b82576001600160a01b0381166000908152600360205260409020600a0180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555b50565b6000611284614b93846168fe565b614b9e906001618176565b620f4240614623866001600160a01b03166000908152600360205260409020600201546001600160801b031690565b600080614bda86856129d9565b6001600160a01b038088166000908152600360205260409020919250614c2d907f00000000000000000000000000000000000000000000000000000000000000001687306001600160801b038916616942565b8054600160801b6001600160801b03808316880181166fffffffffffffffffffffffffffffffff19938416811783900482168601821690920290911783556001600160a01b038089166000818152600b8601602052604090208054808516880190941693909416929092179092559088169003614cc7576002810180546001600160801b03600160801b8083048216890182160291161790555b8315614cd957614cd98782600061463e565b5095945050505050565b6001600160a01b03821660009081526003602052604081205461128490620f424090600160801b90046001600160801b0316614d1f9190618176565b6001600160a01b0385166000908152600360205260409020546001600160801b03165b61462d906001618176565b6001600160a01b038716600090815260036020526040812090614d708986611278565b82546002840180546fffffffffffffffffffffffffffffffff198082166001600160801b039283168a9003831617909255600160801b80840482168a90038216810280841694831690831617859003821693909317865560018601805484810483168c90038316909402808416948316908316178c900382169390931790925560058501805491821668056bc75e2d631000008c0288039284168390038416179055604080518082019091529293509188840391600091908190614e37908b168588614f24565b6001600160801b03168152602001614e6385898b6001600160801b0316614f249092919063ffffffff16565b6001600160801b03168152509050614eb3858c8c8080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525087925088915086905061668b565b5060009050614ec28284618176565b6001600160801b03161115614b0b57614b0b8a84600161463e565b6000611284620f4240614f11856001600160a01b03166000908152600360205260409020600201546001600160801b031690565b614f1b9190618176565b614d42856168fe565b6000838302608081901c6001600160801b03841611614f855760405162461bcd60e51b815260206004820152601c60248201527f55696e743132384d6174683a206d756c446976206f766572666c6f77000000006044820152606401610cab565b826001600160801b03168181614f9d57614f9d6181fe565b0495945050505050565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633f98365b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015615008573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061502c919061815d565b615037906001618061565b905060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166369f16eec6040518163ffffffff1660e01b8152600401602060405180830381865afa158015615099573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906150bd919061815d565b90507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663b98debbf6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561511d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061514191906181c4565b6001600160a01b031663ad36d6cc836040518263ffffffff1660e01b815260040161516e91815260200190565b602060405180830381865afa15801561518b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906151af919061822d565b6151bc5760009250505090565b60408051608081018252600080825260208201819052918101829052606081018290525b7f00000000000000000000000000000000000000000000000000000000000000006001600160801b0316826001600160801b03161080156152215750828411155b15615468576040517f33727c4d000000000000000000000000000000000000000000000000000000008152600481018590527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906333727c4d90602401602060405180830381865afa1580156152a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906152c8919061822d565b15615468576040517fa25ae557000000000000000000000000000000000000000000000000000000008152600481018590527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063a25ae55790602401608060405180830381865afa15801561534b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061536f919061824f565b905061537e8160000151616993565b6000848152600460205260409020546001600160801b031680156154325781516001600160a01b031660009081526003602052604081206153c0918390615921565b5060008581526004602090815260409182902080546fffffffffffffffffffffffffffffffff19169055835191516001600160801b03841681526001600160a01b03909216917f568d79fa2b3ed5751db3f4588be94b7eb2127a4696e56c68d8983a04ad0f3f50910160405180910390a25b81516001600160a01b0381166000908152600360205260408120615456929161463e565b846001019450826001019250506151e0565b6001600160801b0382161561550c576001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001663b4c302ff6154b160018761801b565b6040518263ffffffff1660e01b81526004016154cf91815260200190565b600060405180830381600087803b1580156154e957600080fd5b505af11580156154fd573d6000803e3d6000fd5b50505050600194505050505090565b600094505050505090565b6005546001600160a01b03161561568b576005546001600160a01b03166000908152600360205260409020600a01547f00000000000000000000000000000000000000000000000000000000000000006001600160801b031661010090910460ff16106156335760006155b36001600160801b037f00000000000000000000000000000000000000000000000000000000000000001642618061565b600580546001600160a01b0390811660009081526006602090815260409182902080546fffffffffffffffffffffffffffffffff19166001600160801b0387169081179091559354915193845293945016917f95a398f2b6b2ad94f281708c97fe502386fc16adca43daed577a1e992a4cc814910160405180910390a250565b6005546001600160a01b03166000908152600360205260409020600a01805460ff6101008083048216600101909116027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff9091161790555b565b6000805468010000000000000000900463ffffffff16815260016020819052604082200154600160801b90046effffffffffffffffffffffffffffff16905060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633f98365b6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561572c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615750919061815d565b90506000826effffffffffffffffffffffffffffff161180156157735750600081115b156158f5576040517fa25ae557000000000000000000000000000000000000000000000000000000008152600481018290526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063a25ae55790602401608060405180830381865afa1580156157f9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061581d919061824f565b90506000838260200151434144600143615837919061801b565b6040805160208101969096528501939093527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606092831b1691840191909152607483015240609482015260b4016040516020818303038152906040528051906020012060001c6158a891906182f2565b90506158b5600082616b0c565b600580547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055506125479050565b600580547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555050565b6001830154600584015460078501546003860154865460009485946001600160801b03600160801b928390048116958282169590821694928490048216938104821692615972929182169116618176565b61597c9190618176565b6159869190618176565b6159909190618176565b61599a9190618176565b6159a49190618096565b6040805160c081019091526001870154875492935060009282916159db916001600160801b03600160801b90920482169116618096565b6001600160801b0390811682526005890154808216602084015260038a01548083166040850152600160801b908190048316606085015290048116608083015260078901541660a09091015290508315615cb557615a656001600160801b0383167f00000000000000000000000000000000000000000000000000000000000000006103e8614f24565b94507f00000000000000000000000000000000000000000000000000000000000000006001600160801b0316856001600160801b03161015615ac5577f000000000000000000000000000000000000000000000000000000000000000094505b615ae385838360005b60200201516001600160801b03169190614f24565b86546fffffffffffffffffffffffffffffffff1981166001600160801b039182169290920316178655615b198583836001615ace565b6005870180546fffffffffffffffffffffffffffffffff1981166001600160801b039182169390930316919091179055615b568583836002615ace565b600387810180546fffffffffffffffffffffffffffffffff1981166001600160801b03918216949094031692909217909155615b9790869084908490615ace565b6003870180546001600160801b03808216600160801b928390048216949094031602919091179055615bcc8583836004615ace565b600587810180546001600160801b03808216600160801b92839004821695909503160292909217909155615c0590869084908490615ace565b6007870180546001600160801b038082169390930383166fffffffffffffffffffffffffffffffff19909116179055600090615c4690871660146064614f24565b95869003959050615caa6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000167f00000000000000000000000000000000000000000000000000000000000000006001600160801b03841661436d565b859350505050613be9565b615cc28583836000615ace565b86546fffffffffffffffffffffffffffffffff1981166001600160801b039182169290920116178655615cf88583836001615ace565b6005870180546fffffffffffffffffffffffffffffffff1981166001600160801b039182169390930116919091179055615d358583836002615ace565b600387810180546fffffffffffffffffffffffffffffffff1981166001600160801b03918216949094011692909217909155615d7690869084908490615ace565b6003870180546001600160801b03600160801b808304821690940181169093029216919091179055615dab8583836004615ace565b600580880180546001600160801b03600160801b80830482169095018116909402931692909217909155615de490869084908490615ace565b6007870180546fffffffffffffffffffffffffffffffff1981166001600160801b03918216939093011691909117905550839150613be99050565b63ffffffff80821660009081526001840160205260408082208054740100000000000000000000000000000000000000009004909316825290205b815474010000000000000000000000000000000000000000900463ffffffff1615801590615ead5750600180820154908301546effffffffffffffffffffffffffffff6101009283900481169290910416115b156146c657815481547fffffffffffffffffffffffff00000000000000000000000000000000000000008083166001600160a01b03928316178555835416918116919091178255600180840180548483018054610100908190046effffffffffffffffffffffffffffff9081168083027fffffffffffffffffffffffffffffffff000000000000000000000000000000ff80871691909117875584549584900483169384029516949094179092558354929003600160801b808404831691909103909116027fff000000000000000000000000000000ffffffffffffffffffffffffffffffff9091161790558354821660009081526002870160209081526040808320805463ffffffff998a1663ffffffff199182161790915587549654909516835280832080549095167401000000000000000000000000000000000000000096879004891617909455945484900486168082529187019094528184208054939093049094168352909120615e5a565b6000616073826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316616cb49092919063ffffffff16565b9050805160001480616094575080806020019051810190616094919061822d565b6144345760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610cab565b6001600160a01b038116600090815260028301602052604081205463ffffffff16808203616138576000915050611287565b6001600160a01b03831660009081526002850160209081526040808320805463ffffffff1916905563ffffffff8481168452600180890190935292208054910154740100000000000000000000000000000000000000009091049091169061010090046effffffffffffffffffffffffffffff165b63ffffffff8216156162425763ffffffff91821660009081526001808801602052604090912090810180546effffffffffffffffffffffffffffff600160801b8083048216869003909116027fff000000000000000000000000000000ffffffffffffffffffffffffffffffff90911617905554740100000000000000000000000000000000000000009004909116906161ad565b61624c8684616cc3565b50508354600163ffffffff64010000000080840482168301909116027fffffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffff90921691909117855591505092915050565b6001600160a01b038216600090815260028401602052604081205463ffffffff168082036162cd576000915050613be9565b63ffffffff80821660009081526001808801602052604090912090810180546effffffffffffffffffffffffffffff8781166101008181027fffffffffffffffffffffffffffffffff000000000000000000000000000000ff85161790945593549290910416927401000000000000000000000000000000000000000090910416908210156164635763ffffffff83166000908152600188810160205260409091200180547fff000000000000000000000000000000ffffffffffffffffffffffffffffffff8116848803600160801b928390046effffffffffffffffffffffffffffff908116820116909202179091555b63ffffffff8216156164535763ffffffff91821660009081526001808a01602052604090912090810180546effffffffffffffffffffffffffffff600160801b80830482168601909116027fff000000000000000000000000000000ffffffffffffffffffffffffffffffff90911617905554740100000000000000000000000000000000000000009004909116906163bf565b5061645e8784615e1f565b61656d565b63ffffffff83166000908152600188810160205260409091200180547fff000000000000000000000000000000ffffffffffffffffffffffffffffffff8116878503600160801b928390046effffffffffffffffffffffffffffff90811682900316909202179091555b63ffffffff8216156165625763ffffffff91821660009081526001808a01602052604090912090810180546effffffffffffffffffffffffffffff600160801b8083048216869003909116027fff000000000000000000000000000000ffffffffffffffffffffffffffffffff90911617905554740100000000000000000000000000000000000000009004909116906164cd565b5061656d878461741b565b5060019695505050505050565b60038301805483919060009061659a9084906001600160801b0316618176565b92506101000a8154816001600160801b0302191690836001600160801b03160217905550808360000160040160008282829054906101000a90046001600160801b03166165e79190618176565b82546101009290920a6001600160801b03818102199093169183160217909155336000908152600b86016020908152604080832042845260020190915281208054859450909261663991859116618176565b82546001600160801b039182166101009390930a92830291909202199091161790555050336000908152600b909201602090815260408320600190810180549182018155845292204292019190915550565b60005b84518110156166fb57336000908152600c870160209081526040808320428452600201909152902085518690839081106166ca576166ca618032565b60209081029190910181015182546001810184556000938452919092200155806166f381618347565b91505061668e565b5060038501805484919060009061671c9084906001600160801b0316618176565b82546101009290920a6001600160801b0381810219909316918316021790915582516004880180549193509160009161675791859116618176565b92506101000a8154816001600160801b0302191690836001600160801b03160217905550818560000160030160108282829054906101000a90046001600160801b03166167a49190618176565b92506101000a8154816001600160801b0302191690836001600160801b0316021790555080602001518560000160040160108282829054906101000a90046001600160801b03166167f59190618176565b82546101009290920a6001600160801b038181021990931691831602179091558251336000908152600c8901602090815260408083204284526003019091528120805492945092909161684a91859116618176565b82546101009290920a6001600160801b03818102199093169183160217909155602083810151336000908152600c8a0183526040808220428352600301909352919091208054919350916010916168aa918591600160801b900416618176565b82546001600160801b039182166101009390930a92830291909202199091161790555050336000908152600c9094016020908152604085206001908101805491820181558652942042940193909355505050565b6001600160a01b038116600090815260036020526040812060058101546001909101546001600160801b0391821691613df09168056bc75e2d6310000091166180be565b6040516001600160a01b03808516602483015283166044820152606481018290526146c69085907f23b872dd00000000000000000000000000000000000000000000000000000000906084016143b2565b6001600160a01b038116600090815260036020526040812060068101546001820154919260ff909116916169cf906001600160801b031661782c565b9050600080616a0b6001600160801b037f0000000000000000000000000000000000000000000000000000000000000000850116856064614f24565b9050616a456001600160801b037f000000000000000000000000000000000000000000000000000000000000000016606486810390614f24565b9150616a5f6001600160801b038416606486810390614f24565b85546fffffffffffffffffffffffffffffffff198082166001600160801b0392831686018316178855600588018054600160801b9281169084168501841690811783900484168601841690920290911790556040805184831681528583166020820152918316908201529093506001600160a01b038716907f36f11936e926f4c5f13247a0f85bfd1361293f182bc6a64bfff082b39aec64d99060600160405180910390a2505050505050565b815460009068010000000000000000900463ffffffff165b63ffffffff80821660009081526001808701602052604080832054600160c01b9004909316825291902001546effffffffffffffffffffffffffffff808516600160801b909204161115616b985763ffffffff9081166000908152600185016020526040902054600160c01b900416616b24565b63ffffffff8181166000818152600187810160205260408083208054600160c01b900490951683528220810154929091529190910154600160801b9091046effffffffffffffffffffffffffffff90811690940393848116610100909204161115616c245763ffffffff1660009081526001840160205260409020546001600160a01b03169050611287565b63ffffffff818116600090815260018681016020526040808320808301549054600160e01b9004909416835290912001546101009091046effffffffffffffffffffffffffffff90811690940393848116600160801b909204161115616caa5763ffffffff9081166000908152600185016020526040902054600160e01b900416616b24565b6000915050611287565b606061258a8484600085617896565b63ffffffff8082166000908152600184016020526040812080549092600160c01b909104169003616f99578054600160e01b900463ffffffff16600003616e5557805474010000000000000000000000000000000000000000900463ffffffff16600003616d565782547fffffffffffffffffffffffffffffffffffffffff00000000ffffffffffffffff168355616e11565b600181015460ff1615616dbe57805474010000000000000000000000000000000000000000900463ffffffff166000908152600184016020526040902080547fffffffff00000000ffffffffffffffffffffffffffffffffffffffffffffffff169055616e11565b805474010000000000000000000000000000000000000000900463ffffffff166000908152600184016020526040902080547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690555b5063ffffffff1660009081526001918201602052604081209081550180547fff00000000000000000000000000000000000000000000000000000000000000169055565b805463ffffffff600160e01b80830482166000908152600180880160209081526040808420547fffffffffffffffffffffffff00000000000000000000000000000000000000009097166001600160a01b0397881617808955859004861684528084208084018054948a0180546effffffffffffffffffffffffffffff6101009788900481169097027fffffffffffffffffffffffffffffffff000000000000000000000000000000ff8216811783559254600160801b908190049097169096027fff000000000000000000000000000000ffffffffffffffffffffffffffffffff9092167fff000000000000000000000000000000000000000000000000000000000000ff9096169590951717909355915490951682526002880190945292909220805494821663ffffffff199095169490941790935581540490911690617415565b8054600160e01b900463ffffffff166000036170f357805463ffffffff600160c01b80830482166000908152600180880160209081526040808420547fffffffffffffffffffffffff00000000000000000000000000000000000000009097166001600160a01b0397881617808955859004861684528084208084018054948a0180546effffffffffffffffffffffffffffff6101009788900481169097027fffffffffffffffffffffffffffffffff000000000000000000000000000000ff8216811783559254600160801b908190049097169096027fff000000000000000000000000000000ffffffffffffffffffffffffffffffff9092167fff000000000000000000000000000000000000000000000000000000000000ff9096169590951717909355915490951682526002880190945292909220805494821663ffffffff199095169490941790935581540490911690617415565b805463ffffffff600160e01b82048116600090815260018087016020526040808320820154600160c01b909504909316825291902001546effffffffffffffffffffffffffffff610100928390048116929091041611156172b457805463ffffffff600160c01b80830482166000908152600180880160209081526040808420547fffffffffffffffffffffffff00000000000000000000000000000000000000009097166001600160a01b03978816178089558581048716855281852080850180548b870180546effffffffffffffffffffffffffffff6101009384900481169093027fffffffffffffffffffffffffffffffff000000000000000000000000000000ff821681178355600160e01b9096048c168a52868a20909801549254600160801b908190048316938190048316939093019091169091027fff000000000000000000000000000000ffffffffffffffffffffffffffffffff9093167fff000000000000000000000000000000000000000000000000000000000000ff9096169590951791909117909355915490951682526002880190945292909220805494821663ffffffff199095169490941790935581540490911690617415565b805463ffffffff600160e01b80830482166000908152600180880160209081526040808420547fffffffffffffffffffffffff00000000000000000000000000000000000000009097166001600160a01b03978816178089558581048716855281852080850180548b870180546effffffffffffffffffffffffffffff6101009384900481169093027fffffffffffffffffffffffffffffffff000000000000000000000000000000ff8216811783559354600160c01b9096048c168a52868a2090980154600160801b908190048316958190048316959095019091169093027fff000000000000000000000000000000ffffffffffffffffffffffffffffffff9091167fff000000000000000000000000000000000000000000000000000000000000ff90961695909517949094179055915490951682526002880190945292909220805494821663ffffffff1990951694909417909355815404909116905b50616cc3565b5b63ffffffff8082166000908152600180850160205260408083208054600160e01b810486168552828520840154600160c01b90910490951684529220015490916effffffffffffffffffffffffffffff61010091829004811691909204909116111561765757600180820154825463ffffffff600160c01b90910416600090815285830160205260409020909101546effffffffffffffffffffffffffffff6101009283900481169290910416111561443457805463ffffffff600160c01b80830482166000908152600187810160208181526040808520547fffffffffffffffffffffffff0000000000000000000000000000000000000000808a166001600160a01b0392831617808c558890048916875282872080549091169982169990991790985583890180548a548890048916875282872086018054610100908190046effffffffffffffffffffffffffffff9081168083027fffffffffffffffffffffffffffffffff000000000000000000000000000000ff80871691909117909655835494839004821692830294909516939093179091558b548990048a168852838820909601805496909203600160801b808804831691909103909116027fff000000000000000000000000000000ffffffffffffffffffffffffffffffff9095169490941790935587548716845260028a0180825283852080549a881663ffffffff199b8c161790558854869004871680865292825283852054909716845295909552902080549095169092179093559054041661741c565b600180820154825463ffffffff600160e01b90910416600090815285830160205260409020909101546effffffffffffffffffffffffffffff6101009283900481169290910416111561443457805463ffffffff600160e01b80830482166000908152600187810160208181526040808520547fffffffffffffffffffffffff0000000000000000000000000000000000000000808a166001600160a01b0392831617808c558890048916875282872080549091169982169990991790985583890180548a548890048916875282872086018054610100908190046effffffffffffffffffffffffffffff9081168083027fffffffffffffffffffffffffffffffff000000000000000000000000000000ff80871691909117909655835494839004821692830294909516939093179091558b548990048a168852838820909601805496909203600160801b808804831691909103909116027fff000000000000000000000000000000ffffffffffffffffffffffffffffffff9095169490941790935587548716845260028a0180825283852080549a881663ffffffff199b8c161790558854869004871680865292825283852054909716845295909552902080549095169092179093559054041661741c565b6000806178656001600160801b037f00000000000000000000000000000000000000000000000000000000000000001660286064614f24565b9050613be9816001600160801b03166501000000000061788f866001600160801b03166064617988565b9190617a4f565b60608247101561790e5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610cab565b600080866001600160a01b0316858760405161792a9190618361565b60006040518083038185875af1925050503d8060008114617967576040519150601f19603f3d011682016040523d82523d6000602084013e61796c565b606091505b509150915061797d87838387617b41565b979650505050505050565b60008083831080156179a157600181146179b4576179c3565b65010000000000850284900491506179c3565b65010000000000840285900491505b506402ef6c3406818002602890811c808402821c808202831c808302841c808402851c938402851c95909502841c641da06a6e33909502841c6455232d2bb2909202841c640d4ca0c283909302841c643177d95571909102841c64fffe4bcada90960290931c9490940191909101039190910303905081831115611287576501921fb544430392915050565b6000808060001985870985870292508281108382030391505080600003617a8957838281617a7f57617a7f6181fe565b0492505050613be9565b808411617ad85760405162461bcd60e51b815260206004820152601560248201527f4d6174683a206d756c446976206f766572666c6f7700000000000000000000006044820152606401610cab565b60008486880960026001871981018816978890046003810283188082028403028082028403028082028403028082028403028082028403029081029092039091026000889003889004909101858311909403939093029303949094049190911702949350505050565b60608315617bb0578251600003617ba9576001600160a01b0385163b617ba95760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610cab565b508161258a565b61258a8383815115617bc55781518083602001fd5b8060405162461bcd60e51b8152600401610cab9190617e9d565b5080546000825590600052602060002090810190614b8291905b80821115617c0d5760008155600101617bf9565b5090565b6001600160a01b0381168114614b8257600080fd5b600080600080600060808688031215617c3e57600080fd5b8535617c4981617c11565b94506020860135617c5981617c11565b935060408601359250606086013567ffffffffffffffff80821115617c7d57600080fd5b818801915088601f830112617c9157600080fd5b813581811115617ca057600080fd5b896020828501011115617cb257600080fd5b9699959850939650602001949392505050565b60008060408385031215617cd857600080fd5b8235617ce381617c11565b946020939093013593505050565b6001600160801b0381168114614b8257600080fd5b60008060408385031215617d1957600080fd5b8235617d2481617c11565b91506020830135617d3481617cf1565b809150509250929050565b600060208284031215617d5157600080fd5b8135613be981617cf1565b600060208284031215617d6e57600080fd5b8135613be981617c11565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6020810160078310617de3577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b600080600060408486031215617dfe57600080fd5b8335617e0981617c11565b9250602084013567ffffffffffffffff80821115617e2657600080fd5b818601915086601f830112617e3a57600080fd5b813581811115617e4957600080fd5b8760208260051b8501011115617e5e57600080fd5b6020830194508093505050509250925092565b60005b83811015617e8c578181015183820152602001617e74565b838111156146c65750506000910152565b6020815260008251806020840152617ebc816040850160208701617e71565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b60008060408385031215617f0157600080fd5b8235617f0c81617c11565b91506020830135617d3481617c11565b803560ff81168114617f2d57600080fd5b919050565b600060208284031215617f4457600080fd5b61128482617f1c565b600080600060608486031215617f6257600080fd5b8335617f6d81617cf1565b9250617f7b60208501617f1c565b9150617f8960408501617f1c565b90509250925092565b600060208284031215617fa457600080fd5b5035919050565b600080600060608486031215617fc057600080fd5b8335617fcb81617c11565b92506020840135617fdb81617c11565b929592945050506040919091013590565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561802d5761802d617fec565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000821982111561807457618074617fec565b500190565b60006020828403121561808b57600080fd5b8151613be981617cf1565b60006001600160801b03838116908316818110156180b6576180b6617fec565b039392505050565b60006001600160801b03808316818516818304811182151516156180e4576180e4617fec565b02949350505050565b6060815283606082015260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff85111561812657600080fd5b8460051b8087608085013760009083016080019081526001600160801b039485166020840152929093166040909101529392505050565b60006020828403121561816f57600080fd5b5051919050565b60006001600160801b0380831681851680830382111561819857618198617fec565b01949350505050565b600060ff821660ff8416808210156181bb576181bb617fec565b90039392505050565b6000602082840312156181d657600080fd5b8151613be981617c11565b600063ffffffff838116908316818110156180b6576180b6617fec565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006020828403121561823f57600080fd5b81518015158114613be957600080fd5b60006080828403121561826157600080fd5b6040516080810181811067ffffffffffffffff821117156182ab577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405282516182b981617c11565b81526020838101519082015260408301516182d381617cf1565b604082015260608301516182e681617cf1565b60608201529392505050565b60006effffffffffffffffffffffffffffff8084168061833b577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b92169190910692915050565b6000600019820361835a5761835a617fec565b5060010190565b60008251618373818460208701617e71565b919091019291505056fea164736f6c634300080f000a",
}

// ValidatorManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorManagerMetaData.ABI instead.
var ValidatorManagerABI = ValidatorManagerMetaData.ABI

// ValidatorManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValidatorManagerMetaData.Bin instead.
var ValidatorManagerBin = ValidatorManagerMetaData.Bin

// DeployValidatorManager deploys a new Ethereum contract, binding an instance of ValidatorManager to it.
func DeployValidatorManager(auth *bind.TransactOpts, backend bind.ContractBackend, _constructorParams AssetManagerConstructorParams, _trustedValidator common.Address, _commissionRateMinChangeSeconds *big.Int, _roundDurationSeconds *big.Int, _jailPeriodSeconds *big.Int, _jailThreshold *big.Int) (common.Address, *types.Transaction, *ValidatorManager, error) {
	parsed, err := ValidatorManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValidatorManagerBin), backend, _constructorParams, _trustedValidator, _commissionRateMinChangeSeconds, _roundDurationSeconds, _jailPeriodSeconds, _jailThreshold)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidatorManager{ValidatorManagerCaller: ValidatorManagerCaller{contract: contract}, ValidatorManagerTransactor: ValidatorManagerTransactor{contract: contract}, ValidatorManagerFilterer: ValidatorManagerFilterer{contract: contract}}, nil
}

// ValidatorManager is an auto generated Go binding around an Ethereum contract.
type ValidatorManager struct {
	ValidatorManagerCaller     // Read-only binding to the contract
	ValidatorManagerTransactor // Write-only binding to the contract
	ValidatorManagerFilterer   // Log filterer for contract events
}

// ValidatorManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorManagerSession struct {
	Contract     *ValidatorManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorManagerCallerSession struct {
	Contract *ValidatorManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ValidatorManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorManagerTransactorSession struct {
	Contract     *ValidatorManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ValidatorManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorManagerRaw struct {
	Contract *ValidatorManager // Generic contract binding to access the raw methods on
}

// ValidatorManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorManagerCallerRaw struct {
	Contract *ValidatorManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorManagerTransactorRaw struct {
	Contract *ValidatorManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorManager creates a new instance of ValidatorManager, bound to a specific deployed contract.
func NewValidatorManager(address common.Address, backend bind.ContractBackend) (*ValidatorManager, error) {
	contract, err := bindValidatorManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorManager{ValidatorManagerCaller: ValidatorManagerCaller{contract: contract}, ValidatorManagerTransactor: ValidatorManagerTransactor{contract: contract}, ValidatorManagerFilterer: ValidatorManagerFilterer{contract: contract}}, nil
}

// NewValidatorManagerCaller creates a new read-only instance of ValidatorManager, bound to a specific deployed contract.
func NewValidatorManagerCaller(address common.Address, caller bind.ContractCaller) (*ValidatorManagerCaller, error) {
	contract, err := bindValidatorManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerCaller{contract: contract}, nil
}

// NewValidatorManagerTransactor creates a new write-only instance of ValidatorManager, bound to a specific deployed contract.
func NewValidatorManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorManagerTransactor, error) {
	contract, err := bindValidatorManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerTransactor{contract: contract}, nil
}

// NewValidatorManagerFilterer creates a new log filterer instance of ValidatorManager, bound to a specific deployed contract.
func NewValidatorManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorManagerFilterer, error) {
	contract, err := bindValidatorManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerFilterer{contract: contract}, nil
}

// bindValidatorManager binds a generic wrapper to an already deployed contract.
func bindValidatorManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValidatorManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorManager *ValidatorManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorManager.Contract.ValidatorManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorManager *ValidatorManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorManager.Contract.ValidatorManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorManager *ValidatorManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorManager.Contract.ValidatorManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorManager *ValidatorManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorManager *ValidatorManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorManager *ValidatorManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorManager.Contract.contract.Transact(opts, method, params...)
}

// ASSETTOKEN is a free data retrieval call binding the contract method 0xd7062005.
//
// Solidity: function ASSET_TOKEN() view returns(address)
func (_ValidatorManager *ValidatorManagerCaller) ASSETTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "ASSET_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ASSETTOKEN is a free data retrieval call binding the contract method 0xd7062005.
//
// Solidity: function ASSET_TOKEN() view returns(address)
func (_ValidatorManager *ValidatorManagerSession) ASSETTOKEN() (common.Address, error) {
	return _ValidatorManager.Contract.ASSETTOKEN(&_ValidatorManager.CallOpts)
}

// ASSETTOKEN is a free data retrieval call binding the contract method 0xd7062005.
//
// Solidity: function ASSET_TOKEN() view returns(address)
func (_ValidatorManager *ValidatorManagerCallerSession) ASSETTOKEN() (common.Address, error) {
	return _ValidatorManager.Contract.ASSETTOKEN(&_ValidatorManager.CallOpts)
}

// BASEREWARD is a free data retrieval call binding the contract method 0x22009af6.
//
// Solidity: function BASE_REWARD() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCaller) BASEREWARD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "BASE_REWARD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASEREWARD is a free data retrieval call binding the contract method 0x22009af6.
//
// Solidity: function BASE_REWARD() view returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) BASEREWARD() (*big.Int, error) {
	return _ValidatorManager.Contract.BASEREWARD(&_ValidatorManager.CallOpts)
}

// BASEREWARD is a free data retrieval call binding the contract method 0x22009af6.
//
// Solidity: function BASE_REWARD() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCallerSession) BASEREWARD() (*big.Int, error) {
	return _ValidatorManager.Contract.BASEREWARD(&_ValidatorManager.CallOpts)
}

// BOOSTEDREWARDDENOM is a free data retrieval call binding the contract method 0x110d6069.
//
// Solidity: function BOOSTED_REWARD_DENOM() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCaller) BOOSTEDREWARDDENOM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "BOOSTED_REWARD_DENOM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BOOSTEDREWARDDENOM is a free data retrieval call binding the contract method 0x110d6069.
//
// Solidity: function BOOSTED_REWARD_DENOM() view returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) BOOSTEDREWARDDENOM() (*big.Int, error) {
	return _ValidatorManager.Contract.BOOSTEDREWARDDENOM(&_ValidatorManager.CallOpts)
}

// BOOSTEDREWARDDENOM is a free data retrieval call binding the contract method 0x110d6069.
//
// Solidity: function BOOSTED_REWARD_DENOM() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCallerSession) BOOSTEDREWARDDENOM() (*big.Int, error) {
	return _ValidatorManager.Contract.BOOSTEDREWARDDENOM(&_ValidatorManager.CallOpts)
}

// BOOSTEDREWARDNUMERATOR is a free data retrieval call binding the contract method 0x0763fa7e.
//
// Solidity: function BOOSTED_REWARD_NUMERATOR() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCaller) BOOSTEDREWARDNUMERATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "BOOSTED_REWARD_NUMERATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BOOSTEDREWARDNUMERATOR is a free data retrieval call binding the contract method 0x0763fa7e.
//
// Solidity: function BOOSTED_REWARD_NUMERATOR() view returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) BOOSTEDREWARDNUMERATOR() (*big.Int, error) {
	return _ValidatorManager.Contract.BOOSTEDREWARDNUMERATOR(&_ValidatorManager.CallOpts)
}

// BOOSTEDREWARDNUMERATOR is a free data retrieval call binding the contract method 0x0763fa7e.
//
// Solidity: function BOOSTED_REWARD_NUMERATOR() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCallerSession) BOOSTEDREWARDNUMERATOR() (*big.Int, error) {
	return _ValidatorManager.Contract.BOOSTEDREWARDNUMERATOR(&_ValidatorManager.CallOpts)
}

// COMMISSIONRATEDENOM is a free data retrieval call binding the contract method 0xb91b2723.
//
// Solidity: function COMMISSION_RATE_DENOM() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCaller) COMMISSIONRATEDENOM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "COMMISSION_RATE_DENOM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COMMISSIONRATEDENOM is a free data retrieval call binding the contract method 0xb91b2723.
//
// Solidity: function COMMISSION_RATE_DENOM() view returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) COMMISSIONRATEDENOM() (*big.Int, error) {
	return _ValidatorManager.Contract.COMMISSIONRATEDENOM(&_ValidatorManager.CallOpts)
}

// COMMISSIONRATEDENOM is a free data retrieval call binding the contract method 0xb91b2723.
//
// Solidity: function COMMISSION_RATE_DENOM() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCallerSession) COMMISSIONRATEDENOM() (*big.Int, error) {
	return _ValidatorManager.Contract.COMMISSIONRATEDENOM(&_ValidatorManager.CallOpts)
}

// COMMISSIONRATEMINCHANGESECONDS is a free data retrieval call binding the contract method 0xdea15254.
//
// Solidity: function COMMISSION_RATE_MIN_CHANGE_SECONDS() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCaller) COMMISSIONRATEMINCHANGESECONDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "COMMISSION_RATE_MIN_CHANGE_SECONDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COMMISSIONRATEMINCHANGESECONDS is a free data retrieval call binding the contract method 0xdea15254.
//
// Solidity: function COMMISSION_RATE_MIN_CHANGE_SECONDS() view returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) COMMISSIONRATEMINCHANGESECONDS() (*big.Int, error) {
	return _ValidatorManager.Contract.COMMISSIONRATEMINCHANGESECONDS(&_ValidatorManager.CallOpts)
}

// COMMISSIONRATEMINCHANGESECONDS is a free data retrieval call binding the contract method 0xdea15254.
//
// Solidity: function COMMISSION_RATE_MIN_CHANGE_SECONDS() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCallerSession) COMMISSIONRATEMINCHANGESECONDS() (*big.Int, error) {
	return _ValidatorManager.Contract.COMMISSIONRATEMINCHANGESECONDS(&_ValidatorManager.CallOpts)
}

// DECIMALOFFSET is a free data retrieval call binding the contract method 0xde7d4d6a.
//
// Solidity: function DECIMAL_OFFSET() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCaller) DECIMALOFFSET(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "DECIMAL_OFFSET")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DECIMALOFFSET is a free data retrieval call binding the contract method 0xde7d4d6a.
//
// Solidity: function DECIMAL_OFFSET() view returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) DECIMALOFFSET() (*big.Int, error) {
	return _ValidatorManager.Contract.DECIMALOFFSET(&_ValidatorManager.CallOpts)
}

// DECIMALOFFSET is a free data retrieval call binding the contract method 0xde7d4d6a.
//
// Solidity: function DECIMAL_OFFSET() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCallerSession) DECIMALOFFSET() (*big.Int, error) {
	return _ValidatorManager.Contract.DECIMALOFFSET(&_ValidatorManager.CallOpts)
}

// JAILPERIODSECONDS is a free data retrieval call binding the contract method 0xabeba449.
//
// Solidity: function JAIL_PERIOD_SECONDS() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCaller) JAILPERIODSECONDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "JAIL_PERIOD_SECONDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// JAILPERIODSECONDS is a free data retrieval call binding the contract method 0xabeba449.
//
// Solidity: function JAIL_PERIOD_SECONDS() view returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) JAILPERIODSECONDS() (*big.Int, error) {
	return _ValidatorManager.Contract.JAILPERIODSECONDS(&_ValidatorManager.CallOpts)
}

// JAILPERIODSECONDS is a free data retrieval call binding the contract method 0xabeba449.
//
// Solidity: function JAIL_PERIOD_SECONDS() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCallerSession) JAILPERIODSECONDS() (*big.Int, error) {
	return _ValidatorManager.Contract.JAILPERIODSECONDS(&_ValidatorManager.CallOpts)
}

// JAILTHRESHOLD is a free data retrieval call binding the contract method 0x42223ae9.
//
// Solidity: function JAIL_THRESHOLD() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCaller) JAILTHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "JAIL_THRESHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// JAILTHRESHOLD is a free data retrieval call binding the contract method 0x42223ae9.
//
// Solidity: function JAIL_THRESHOLD() view returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) JAILTHRESHOLD() (*big.Int, error) {
	return _ValidatorManager.Contract.JAILTHRESHOLD(&_ValidatorManager.CallOpts)
}

// JAILTHRESHOLD is a free data retrieval call binding the contract method 0x42223ae9.
//
// Solidity: function JAIL_THRESHOLD() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCallerSession) JAILTHRESHOLD() (*big.Int, error) {
	return _ValidatorManager.Contract.JAILTHRESHOLD(&_ValidatorManager.CallOpts)
}

// KGH is a free data retrieval call binding the contract method 0x56576b5b.
//
// Solidity: function KGH() view returns(address)
func (_ValidatorManager *ValidatorManagerCaller) KGH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "KGH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KGH is a free data retrieval call binding the contract method 0x56576b5b.
//
// Solidity: function KGH() view returns(address)
func (_ValidatorManager *ValidatorManagerSession) KGH() (common.Address, error) {
	return _ValidatorManager.Contract.KGH(&_ValidatorManager.CallOpts)
}

// KGH is a free data retrieval call binding the contract method 0x56576b5b.
//
// Solidity: function KGH() view returns(address)
func (_ValidatorManager *ValidatorManagerCallerSession) KGH() (common.Address, error) {
	return _ValidatorManager.Contract.KGH(&_ValidatorManager.CallOpts)
}

// KGHMANAGER is a free data retrieval call binding the contract method 0xd1e288c1.
//
// Solidity: function KGH_MANAGER() view returns(address)
func (_ValidatorManager *ValidatorManagerCaller) KGHMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "KGH_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KGHMANAGER is a free data retrieval call binding the contract method 0xd1e288c1.
//
// Solidity: function KGH_MANAGER() view returns(address)
func (_ValidatorManager *ValidatorManagerSession) KGHMANAGER() (common.Address, error) {
	return _ValidatorManager.Contract.KGHMANAGER(&_ValidatorManager.CallOpts)
}

// KGHMANAGER is a free data retrieval call binding the contract method 0xd1e288c1.
//
// Solidity: function KGH_MANAGER() view returns(address)
func (_ValidatorManager *ValidatorManagerCallerSession) KGHMANAGER() (common.Address, error) {
	return _ValidatorManager.Contract.KGHMANAGER(&_ValidatorManager.CallOpts)
}

// L2ORACLE is a free data retrieval call binding the contract method 0x001c2ff6.
//
// Solidity: function L2_ORACLE() view returns(address)
func (_ValidatorManager *ValidatorManagerCaller) L2ORACLE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "L2_ORACLE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2ORACLE is a free data retrieval call binding the contract method 0x001c2ff6.
//
// Solidity: function L2_ORACLE() view returns(address)
func (_ValidatorManager *ValidatorManagerSession) L2ORACLE() (common.Address, error) {
	return _ValidatorManager.Contract.L2ORACLE(&_ValidatorManager.CallOpts)
}

// L2ORACLE is a free data retrieval call binding the contract method 0x001c2ff6.
//
// Solidity: function L2_ORACLE() view returns(address)
func (_ValidatorManager *ValidatorManagerCallerSession) L2ORACLE() (common.Address, error) {
	return _ValidatorManager.Contract.L2ORACLE(&_ValidatorManager.CallOpts)
}

// MAXOUTPUTFINALIZATIONS is a free data retrieval call binding the contract method 0xe7816b7f.
//
// Solidity: function MAX_OUTPUT_FINALIZATIONS() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCaller) MAXOUTPUTFINALIZATIONS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "MAX_OUTPUT_FINALIZATIONS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXOUTPUTFINALIZATIONS is a free data retrieval call binding the contract method 0xe7816b7f.
//
// Solidity: function MAX_OUTPUT_FINALIZATIONS() view returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) MAXOUTPUTFINALIZATIONS() (*big.Int, error) {
	return _ValidatorManager.Contract.MAXOUTPUTFINALIZATIONS(&_ValidatorManager.CallOpts)
}

// MAXOUTPUTFINALIZATIONS is a free data retrieval call binding the contract method 0xe7816b7f.
//
// Solidity: function MAX_OUTPUT_FINALIZATIONS() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCallerSession) MAXOUTPUTFINALIZATIONS() (*big.Int, error) {
	return _ValidatorManager.Contract.MAXOUTPUTFINALIZATIONS(&_ValidatorManager.CallOpts)
}

// MINREGISTERAMOUNT is a free data retrieval call binding the contract method 0x1796e52e.
//
// Solidity: function MIN_REGISTER_AMOUNT() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCaller) MINREGISTERAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "MIN_REGISTER_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINREGISTERAMOUNT is a free data retrieval call binding the contract method 0x1796e52e.
//
// Solidity: function MIN_REGISTER_AMOUNT() view returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) MINREGISTERAMOUNT() (*big.Int, error) {
	return _ValidatorManager.Contract.MINREGISTERAMOUNT(&_ValidatorManager.CallOpts)
}

// MINREGISTERAMOUNT is a free data retrieval call binding the contract method 0x1796e52e.
//
// Solidity: function MIN_REGISTER_AMOUNT() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCallerSession) MINREGISTERAMOUNT() (*big.Int, error) {
	return _ValidatorManager.Contract.MINREGISTERAMOUNT(&_ValidatorManager.CallOpts)
}

// MINSLASHINGAMOUNT is a free data retrieval call binding the contract method 0x176a86d0.
//
// Solidity: function MIN_SLASHING_AMOUNT() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCaller) MINSLASHINGAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "MIN_SLASHING_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINSLASHINGAMOUNT is a free data retrieval call binding the contract method 0x176a86d0.
//
// Solidity: function MIN_SLASHING_AMOUNT() view returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) MINSLASHINGAMOUNT() (*big.Int, error) {
	return _ValidatorManager.Contract.MINSLASHINGAMOUNT(&_ValidatorManager.CallOpts)
}

// MINSLASHINGAMOUNT is a free data retrieval call binding the contract method 0x176a86d0.
//
// Solidity: function MIN_SLASHING_AMOUNT() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCallerSession) MINSLASHINGAMOUNT() (*big.Int, error) {
	return _ValidatorManager.Contract.MINSLASHINGAMOUNT(&_ValidatorManager.CallOpts)
}

// MINSTARTAMOUNT is a free data retrieval call binding the contract method 0x3bcebcd8.
//
// Solidity: function MIN_START_AMOUNT() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCaller) MINSTARTAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "MIN_START_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINSTARTAMOUNT is a free data retrieval call binding the contract method 0x3bcebcd8.
//
// Solidity: function MIN_START_AMOUNT() view returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) MINSTARTAMOUNT() (*big.Int, error) {
	return _ValidatorManager.Contract.MINSTARTAMOUNT(&_ValidatorManager.CallOpts)
}

// MINSTARTAMOUNT is a free data retrieval call binding the contract method 0x3bcebcd8.
//
// Solidity: function MIN_START_AMOUNT() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCallerSession) MINSTARTAMOUNT() (*big.Int, error) {
	return _ValidatorManager.Contract.MINSTARTAMOUNT(&_ValidatorManager.CallOpts)
}

// ROUNDDURATIONSECONDS is a free data retrieval call binding the contract method 0x4cca5e6c.
//
// Solidity: function ROUND_DURATION_SECONDS() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCaller) ROUNDDURATIONSECONDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "ROUND_DURATION_SECONDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ROUNDDURATIONSECONDS is a free data retrieval call binding the contract method 0x4cca5e6c.
//
// Solidity: function ROUND_DURATION_SECONDS() view returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) ROUNDDURATIONSECONDS() (*big.Int, error) {
	return _ValidatorManager.Contract.ROUNDDURATIONSECONDS(&_ValidatorManager.CallOpts)
}

// ROUNDDURATIONSECONDS is a free data retrieval call binding the contract method 0x4cca5e6c.
//
// Solidity: function ROUND_DURATION_SECONDS() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCallerSession) ROUNDDURATIONSECONDS() (*big.Int, error) {
	return _ValidatorManager.Contract.ROUNDDURATIONSECONDS(&_ValidatorManager.CallOpts)
}

// SECURITYCOUNCIL is a free data retrieval call binding the contract method 0x36086417.
//
// Solidity: function SECURITY_COUNCIL() view returns(address)
func (_ValidatorManager *ValidatorManagerCaller) SECURITYCOUNCIL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "SECURITY_COUNCIL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SECURITYCOUNCIL is a free data retrieval call binding the contract method 0x36086417.
//
// Solidity: function SECURITY_COUNCIL() view returns(address)
func (_ValidatorManager *ValidatorManagerSession) SECURITYCOUNCIL() (common.Address, error) {
	return _ValidatorManager.Contract.SECURITYCOUNCIL(&_ValidatorManager.CallOpts)
}

// SECURITYCOUNCIL is a free data retrieval call binding the contract method 0x36086417.
//
// Solidity: function SECURITY_COUNCIL() view returns(address)
func (_ValidatorManager *ValidatorManagerCallerSession) SECURITYCOUNCIL() (common.Address, error) {
	return _ValidatorManager.Contract.SECURITYCOUNCIL(&_ValidatorManager.CallOpts)
}

// SLASHINGRATEDENOM is a free data retrieval call binding the contract method 0x1a5deb4a.
//
// Solidity: function SLASHING_RATE_DENOM() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCaller) SLASHINGRATEDENOM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "SLASHING_RATE_DENOM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SLASHINGRATEDENOM is a free data retrieval call binding the contract method 0x1a5deb4a.
//
// Solidity: function SLASHING_RATE_DENOM() view returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) SLASHINGRATEDENOM() (*big.Int, error) {
	return _ValidatorManager.Contract.SLASHINGRATEDENOM(&_ValidatorManager.CallOpts)
}

// SLASHINGRATEDENOM is a free data retrieval call binding the contract method 0x1a5deb4a.
//
// Solidity: function SLASHING_RATE_DENOM() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCallerSession) SLASHINGRATEDENOM() (*big.Int, error) {
	return _ValidatorManager.Contract.SLASHINGRATEDENOM(&_ValidatorManager.CallOpts)
}

// SLASHINGRATENUMERATOR is a free data retrieval call binding the contract method 0xa4cf0b2f.
//
// Solidity: function SLASHING_RATE_NUMERATOR() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCaller) SLASHINGRATENUMERATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "SLASHING_RATE_NUMERATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SLASHINGRATENUMERATOR is a free data retrieval call binding the contract method 0xa4cf0b2f.
//
// Solidity: function SLASHING_RATE_NUMERATOR() view returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) SLASHINGRATENUMERATOR() (*big.Int, error) {
	return _ValidatorManager.Contract.SLASHINGRATENUMERATOR(&_ValidatorManager.CallOpts)
}

// SLASHINGRATENUMERATOR is a free data retrieval call binding the contract method 0xa4cf0b2f.
//
// Solidity: function SLASHING_RATE_NUMERATOR() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCallerSession) SLASHINGRATENUMERATOR() (*big.Int, error) {
	return _ValidatorManager.Contract.SLASHINGRATENUMERATOR(&_ValidatorManager.CallOpts)
}

// TAXDENOMINATOR is a free data retrieval call binding the contract method 0xa51c9ace.
//
// Solidity: function TAX_DENOMINATOR() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCaller) TAXDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "TAX_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TAXDENOMINATOR is a free data retrieval call binding the contract method 0xa51c9ace.
//
// Solidity: function TAX_DENOMINATOR() view returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) TAXDENOMINATOR() (*big.Int, error) {
	return _ValidatorManager.Contract.TAXDENOMINATOR(&_ValidatorManager.CallOpts)
}

// TAXDENOMINATOR is a free data retrieval call binding the contract method 0xa51c9ace.
//
// Solidity: function TAX_DENOMINATOR() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCallerSession) TAXDENOMINATOR() (*big.Int, error) {
	return _ValidatorManager.Contract.TAXDENOMINATOR(&_ValidatorManager.CallOpts)
}

// TAXNUMERATOR is a free data retrieval call binding the contract method 0x82dae3aa.
//
// Solidity: function TAX_NUMERATOR() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCaller) TAXNUMERATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "TAX_NUMERATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TAXNUMERATOR is a free data retrieval call binding the contract method 0x82dae3aa.
//
// Solidity: function TAX_NUMERATOR() view returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) TAXNUMERATOR() (*big.Int, error) {
	return _ValidatorManager.Contract.TAXNUMERATOR(&_ValidatorManager.CallOpts)
}

// TAXNUMERATOR is a free data retrieval call binding the contract method 0x82dae3aa.
//
// Solidity: function TAX_NUMERATOR() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCallerSession) TAXNUMERATOR() (*big.Int, error) {
	return _ValidatorManager.Contract.TAXNUMERATOR(&_ValidatorManager.CallOpts)
}

// TRUSTEDVALIDATOR is a free data retrieval call binding the contract method 0x3ee4d4a3.
//
// Solidity: function TRUSTED_VALIDATOR() view returns(address)
func (_ValidatorManager *ValidatorManagerCaller) TRUSTEDVALIDATOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "TRUSTED_VALIDATOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TRUSTEDVALIDATOR is a free data retrieval call binding the contract method 0x3ee4d4a3.
//
// Solidity: function TRUSTED_VALIDATOR() view returns(address)
func (_ValidatorManager *ValidatorManagerSession) TRUSTEDVALIDATOR() (common.Address, error) {
	return _ValidatorManager.Contract.TRUSTEDVALIDATOR(&_ValidatorManager.CallOpts)
}

// TRUSTEDVALIDATOR is a free data retrieval call binding the contract method 0x3ee4d4a3.
//
// Solidity: function TRUSTED_VALIDATOR() view returns(address)
func (_ValidatorManager *ValidatorManagerCallerSession) TRUSTEDVALIDATOR() (common.Address, error) {
	return _ValidatorManager.Contract.TRUSTEDVALIDATOR(&_ValidatorManager.CallOpts)
}

// UNDELEGATIONPERIOD is a free data retrieval call binding the contract method 0x7533f901.
//
// Solidity: function UNDELEGATION_PERIOD() view returns(uint256)
func (_ValidatorManager *ValidatorManagerCaller) UNDELEGATIONPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "UNDELEGATION_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UNDELEGATIONPERIOD is a free data retrieval call binding the contract method 0x7533f901.
//
// Solidity: function UNDELEGATION_PERIOD() view returns(uint256)
func (_ValidatorManager *ValidatorManagerSession) UNDELEGATIONPERIOD() (*big.Int, error) {
	return _ValidatorManager.Contract.UNDELEGATIONPERIOD(&_ValidatorManager.CallOpts)
}

// UNDELEGATIONPERIOD is a free data retrieval call binding the contract method 0x7533f901.
//
// Solidity: function UNDELEGATION_PERIOD() view returns(uint256)
func (_ValidatorManager *ValidatorManagerCallerSession) UNDELEGATIONPERIOD() (*big.Int, error) {
	return _ValidatorManager.Contract.UNDELEGATIONPERIOD(&_ValidatorManager.CallOpts)
}

// VKROPERKGH is a free data retrieval call binding the contract method 0x631bda01.
//
// Solidity: function VKRO_PER_KGH() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCaller) VKROPERKGH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "VKRO_PER_KGH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VKROPERKGH is a free data retrieval call binding the contract method 0x631bda01.
//
// Solidity: function VKRO_PER_KGH() view returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) VKROPERKGH() (*big.Int, error) {
	return _ValidatorManager.Contract.VKROPERKGH(&_ValidatorManager.CallOpts)
}

// VKROPERKGH is a free data retrieval call binding the contract method 0x631bda01.
//
// Solidity: function VKRO_PER_KGH() view returns(uint128)
func (_ValidatorManager *ValidatorManagerCallerSession) VKROPERKGH() (*big.Int, error) {
	return _ValidatorManager.Contract.VKROPERKGH(&_ValidatorManager.CallOpts)
}

// CheckSubmissionEligibility is a free data retrieval call binding the contract method 0x891aab74.
//
// Solidity: function checkSubmissionEligibility(address validator) view returns()
func (_ValidatorManager *ValidatorManagerCaller) CheckSubmissionEligibility(opts *bind.CallOpts, validator common.Address) error {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "checkSubmissionEligibility", validator)

	if err != nil {
		return err
	}

	return err

}

// CheckSubmissionEligibility is a free data retrieval call binding the contract method 0x891aab74.
//
// Solidity: function checkSubmissionEligibility(address validator) view returns()
func (_ValidatorManager *ValidatorManagerSession) CheckSubmissionEligibility(validator common.Address) error {
	return _ValidatorManager.Contract.CheckSubmissionEligibility(&_ValidatorManager.CallOpts, validator)
}

// CheckSubmissionEligibility is a free data retrieval call binding the contract method 0x891aab74.
//
// Solidity: function checkSubmissionEligibility(address validator) view returns()
func (_ValidatorManager *ValidatorManagerCallerSession) CheckSubmissionEligibility(validator common.Address) error {
	return _ValidatorManager.Contract.CheckSubmissionEligibility(&_ValidatorManager.CallOpts, validator)
}

// GetCommissionMaxChangeRate is a free data retrieval call binding the contract method 0xb9551f82.
//
// Solidity: function getCommissionMaxChangeRate(address validator) view returns(uint8)
func (_ValidatorManager *ValidatorManagerCaller) GetCommissionMaxChangeRate(opts *bind.CallOpts, validator common.Address) (uint8, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "getCommissionMaxChangeRate", validator)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetCommissionMaxChangeRate is a free data retrieval call binding the contract method 0xb9551f82.
//
// Solidity: function getCommissionMaxChangeRate(address validator) view returns(uint8)
func (_ValidatorManager *ValidatorManagerSession) GetCommissionMaxChangeRate(validator common.Address) (uint8, error) {
	return _ValidatorManager.Contract.GetCommissionMaxChangeRate(&_ValidatorManager.CallOpts, validator)
}

// GetCommissionMaxChangeRate is a free data retrieval call binding the contract method 0xb9551f82.
//
// Solidity: function getCommissionMaxChangeRate(address validator) view returns(uint8)
func (_ValidatorManager *ValidatorManagerCallerSession) GetCommissionMaxChangeRate(validator common.Address) (uint8, error) {
	return _ValidatorManager.Contract.GetCommissionMaxChangeRate(&_ValidatorManager.CallOpts, validator)
}

// GetCommissionRate is a free data retrieval call binding the contract method 0xe0cc26a2.
//
// Solidity: function getCommissionRate(address validator) view returns(uint8)
func (_ValidatorManager *ValidatorManagerCaller) GetCommissionRate(opts *bind.CallOpts, validator common.Address) (uint8, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "getCommissionRate", validator)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetCommissionRate is a free data retrieval call binding the contract method 0xe0cc26a2.
//
// Solidity: function getCommissionRate(address validator) view returns(uint8)
func (_ValidatorManager *ValidatorManagerSession) GetCommissionRate(validator common.Address) (uint8, error) {
	return _ValidatorManager.Contract.GetCommissionRate(&_ValidatorManager.CallOpts, validator)
}

// GetCommissionRate is a free data retrieval call binding the contract method 0xe0cc26a2.
//
// Solidity: function getCommissionRate(address validator) view returns(uint8)
func (_ValidatorManager *ValidatorManagerCallerSession) GetCommissionRate(validator common.Address) (uint8, error) {
	return _ValidatorManager.Contract.GetCommissionRate(&_ValidatorManager.CallOpts, validator)
}

// GetKghTotalBalance is a free data retrieval call binding the contract method 0xde313284.
//
// Solidity: function getKghTotalBalance(address validator, address delegator, uint256 tokenId) view returns(uint128)
func (_ValidatorManager *ValidatorManagerCaller) GetKghTotalBalance(opts *bind.CallOpts, validator common.Address, delegator common.Address, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "getKghTotalBalance", validator, delegator, tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetKghTotalBalance is a free data retrieval call binding the contract method 0xde313284.
//
// Solidity: function getKghTotalBalance(address validator, address delegator, uint256 tokenId) view returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) GetKghTotalBalance(validator common.Address, delegator common.Address, tokenId *big.Int) (*big.Int, error) {
	return _ValidatorManager.Contract.GetKghTotalBalance(&_ValidatorManager.CallOpts, validator, delegator, tokenId)
}

// GetKghTotalBalance is a free data retrieval call binding the contract method 0xde313284.
//
// Solidity: function getKghTotalBalance(address validator, address delegator, uint256 tokenId) view returns(uint128)
func (_ValidatorManager *ValidatorManagerCallerSession) GetKghTotalBalance(validator common.Address, delegator common.Address, tokenId *big.Int) (*big.Int, error) {
	return _ValidatorManager.Contract.GetKghTotalBalance(&_ValidatorManager.CallOpts, validator, delegator, tokenId)
}

// GetKghTotalShareBalance is a free data retrieval call binding the contract method 0xcf368e8c.
//
// Solidity: function getKghTotalShareBalance(address validator, address delegator, uint256 tokenId) view returns(uint128, uint128)
func (_ValidatorManager *ValidatorManagerCaller) GetKghTotalShareBalance(opts *bind.CallOpts, validator common.Address, delegator common.Address, tokenId *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "getKghTotalShareBalance", validator, delegator, tokenId)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetKghTotalShareBalance is a free data retrieval call binding the contract method 0xcf368e8c.
//
// Solidity: function getKghTotalShareBalance(address validator, address delegator, uint256 tokenId) view returns(uint128, uint128)
func (_ValidatorManager *ValidatorManagerSession) GetKghTotalShareBalance(validator common.Address, delegator common.Address, tokenId *big.Int) (*big.Int, *big.Int, error) {
	return _ValidatorManager.Contract.GetKghTotalShareBalance(&_ValidatorManager.CallOpts, validator, delegator, tokenId)
}

// GetKghTotalShareBalance is a free data retrieval call binding the contract method 0xcf368e8c.
//
// Solidity: function getKghTotalShareBalance(address validator, address delegator, uint256 tokenId) view returns(uint128, uint128)
func (_ValidatorManager *ValidatorManagerCallerSession) GetKghTotalShareBalance(validator common.Address, delegator common.Address, tokenId *big.Int) (*big.Int, *big.Int, error) {
	return _ValidatorManager.Contract.GetKghTotalShareBalance(&_ValidatorManager.CallOpts, validator, delegator, tokenId)
}

// GetKroTotalBalance is a free data retrieval call binding the contract method 0x8c090350.
//
// Solidity: function getKroTotalBalance(address validator, address delegator) view returns(uint128)
func (_ValidatorManager *ValidatorManagerCaller) GetKroTotalBalance(opts *bind.CallOpts, validator common.Address, delegator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "getKroTotalBalance", validator, delegator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetKroTotalBalance is a free data retrieval call binding the contract method 0x8c090350.
//
// Solidity: function getKroTotalBalance(address validator, address delegator) view returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) GetKroTotalBalance(validator common.Address, delegator common.Address) (*big.Int, error) {
	return _ValidatorManager.Contract.GetKroTotalBalance(&_ValidatorManager.CallOpts, validator, delegator)
}

// GetKroTotalBalance is a free data retrieval call binding the contract method 0x8c090350.
//
// Solidity: function getKroTotalBalance(address validator, address delegator) view returns(uint128)
func (_ValidatorManager *ValidatorManagerCallerSession) GetKroTotalBalance(validator common.Address, delegator common.Address) (*big.Int, error) {
	return _ValidatorManager.Contract.GetKroTotalBalance(&_ValidatorManager.CallOpts, validator, delegator)
}

// GetKroTotalShareBalance is a free data retrieval call binding the contract method 0x842d0d3b.
//
// Solidity: function getKroTotalShareBalance(address validator, address delegator) view returns(uint128)
func (_ValidatorManager *ValidatorManagerCaller) GetKroTotalShareBalance(opts *bind.CallOpts, validator common.Address, delegator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "getKroTotalShareBalance", validator, delegator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetKroTotalShareBalance is a free data retrieval call binding the contract method 0x842d0d3b.
//
// Solidity: function getKroTotalShareBalance(address validator, address delegator) view returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) GetKroTotalShareBalance(validator common.Address, delegator common.Address) (*big.Int, error) {
	return _ValidatorManager.Contract.GetKroTotalShareBalance(&_ValidatorManager.CallOpts, validator, delegator)
}

// GetKroTotalShareBalance is a free data retrieval call binding the contract method 0x842d0d3b.
//
// Solidity: function getKroTotalShareBalance(address validator, address delegator) view returns(uint128)
func (_ValidatorManager *ValidatorManagerCallerSession) GetKroTotalShareBalance(validator common.Address, delegator common.Address) (*big.Int, error) {
	return _ValidatorManager.Contract.GetKroTotalShareBalance(&_ValidatorManager.CallOpts, validator, delegator)
}

// GetStatus is a free data retrieval call binding the contract method 0x30ccebb5.
//
// Solidity: function getStatus(address validator) view returns(uint8)
func (_ValidatorManager *ValidatorManagerCaller) GetStatus(opts *bind.CallOpts, validator common.Address) (uint8, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "getStatus", validator)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetStatus is a free data retrieval call binding the contract method 0x30ccebb5.
//
// Solidity: function getStatus(address validator) view returns(uint8)
func (_ValidatorManager *ValidatorManagerSession) GetStatus(validator common.Address) (uint8, error) {
	return _ValidatorManager.Contract.GetStatus(&_ValidatorManager.CallOpts, validator)
}

// GetStatus is a free data retrieval call binding the contract method 0x30ccebb5.
//
// Solidity: function getStatus(address validator) view returns(uint8)
func (_ValidatorManager *ValidatorManagerCallerSession) GetStatus(validator common.Address) (uint8, error) {
	return _ValidatorManager.Contract.GetStatus(&_ValidatorManager.CallOpts, validator)
}

// GetWeight is a free data retrieval call binding the contract method 0xac6c5251.
//
// Solidity: function getWeight(address validator) view returns(uint120)
func (_ValidatorManager *ValidatorManagerCaller) GetWeight(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "getWeight", validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWeight is a free data retrieval call binding the contract method 0xac6c5251.
//
// Solidity: function getWeight(address validator) view returns(uint120)
func (_ValidatorManager *ValidatorManagerSession) GetWeight(validator common.Address) (*big.Int, error) {
	return _ValidatorManager.Contract.GetWeight(&_ValidatorManager.CallOpts, validator)
}

// GetWeight is a free data retrieval call binding the contract method 0xac6c5251.
//
// Solidity: function getWeight(address validator) view returns(uint120)
func (_ValidatorManager *ValidatorManagerCallerSession) GetWeight(validator common.Address) (*big.Int, error) {
	return _ValidatorManager.Contract.GetWeight(&_ValidatorManager.CallOpts, validator)
}

// JailExpiresAt is a free data retrieval call binding the contract method 0x970531c1.
//
// Solidity: function jailExpiresAt(address validator) view returns(uint128)
func (_ValidatorManager *ValidatorManagerCaller) JailExpiresAt(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "jailExpiresAt", validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// JailExpiresAt is a free data retrieval call binding the contract method 0x970531c1.
//
// Solidity: function jailExpiresAt(address validator) view returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) JailExpiresAt(validator common.Address) (*big.Int, error) {
	return _ValidatorManager.Contract.JailExpiresAt(&_ValidatorManager.CallOpts, validator)
}

// JailExpiresAt is a free data retrieval call binding the contract method 0x970531c1.
//
// Solidity: function jailExpiresAt(address validator) view returns(uint128)
func (_ValidatorManager *ValidatorManagerCallerSession) JailExpiresAt(validator common.Address) (*big.Int, error) {
	return _ValidatorManager.Contract.JailExpiresAt(&_ValidatorManager.CallOpts, validator)
}

// NextValidator is a free data retrieval call binding the contract method 0x3a549046.
//
// Solidity: function nextValidator() view returns(address)
func (_ValidatorManager *ValidatorManagerCaller) NextValidator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "nextValidator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NextValidator is a free data retrieval call binding the contract method 0x3a549046.
//
// Solidity: function nextValidator() view returns(address)
func (_ValidatorManager *ValidatorManagerSession) NextValidator() (common.Address, error) {
	return _ValidatorManager.Contract.NextValidator(&_ValidatorManager.CallOpts)
}

// NextValidator is a free data retrieval call binding the contract method 0x3a549046.
//
// Solidity: function nextValidator() view returns(address)
func (_ValidatorManager *ValidatorManagerCallerSession) NextValidator() (common.Address, error) {
	return _ValidatorManager.Contract.NextValidator(&_ValidatorManager.CallOpts)
}

// NoSubmissionCount is a free data retrieval call binding the contract method 0xdff221b5.
//
// Solidity: function noSubmissionCount(address validator) view returns(uint8)
func (_ValidatorManager *ValidatorManagerCaller) NoSubmissionCount(opts *bind.CallOpts, validator common.Address) (uint8, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "noSubmissionCount", validator)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// NoSubmissionCount is a free data retrieval call binding the contract method 0xdff221b5.
//
// Solidity: function noSubmissionCount(address validator) view returns(uint8)
func (_ValidatorManager *ValidatorManagerSession) NoSubmissionCount(validator common.Address) (uint8, error) {
	return _ValidatorManager.Contract.NoSubmissionCount(&_ValidatorManager.CallOpts, validator)
}

// NoSubmissionCount is a free data retrieval call binding the contract method 0xdff221b5.
//
// Solidity: function noSubmissionCount(address validator) view returns(uint8)
func (_ValidatorManager *ValidatorManagerCallerSession) NoSubmissionCount(validator common.Address) (uint8, error) {
	return _ValidatorManager.Contract.NoSubmissionCount(&_ValidatorManager.CallOpts, validator)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_ValidatorManager *ValidatorManagerCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_ValidatorManager *ValidatorManagerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _ValidatorManager.Contract.OnERC721Received(&_ValidatorManager.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_ValidatorManager *ValidatorManagerCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _ValidatorManager.Contract.OnERC721Received(&_ValidatorManager.CallOpts, arg0, arg1, arg2, arg3)
}

// PreviewDelegate is a free data retrieval call binding the contract method 0x960a0893.
//
// Solidity: function previewDelegate(address validator, uint128 assets) view returns(uint128)
func (_ValidatorManager *ValidatorManagerCaller) PreviewDelegate(opts *bind.CallOpts, validator common.Address, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "previewDelegate", validator, assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDelegate is a free data retrieval call binding the contract method 0x960a0893.
//
// Solidity: function previewDelegate(address validator, uint128 assets) view returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) PreviewDelegate(validator common.Address, assets *big.Int) (*big.Int, error) {
	return _ValidatorManager.Contract.PreviewDelegate(&_ValidatorManager.CallOpts, validator, assets)
}

// PreviewDelegate is a free data retrieval call binding the contract method 0x960a0893.
//
// Solidity: function previewDelegate(address validator, uint128 assets) view returns(uint128)
func (_ValidatorManager *ValidatorManagerCallerSession) PreviewDelegate(validator common.Address, assets *big.Int) (*big.Int, error) {
	return _ValidatorManager.Contract.PreviewDelegate(&_ValidatorManager.CallOpts, validator, assets)
}

// PreviewKghDelegate is a free data retrieval call binding the contract method 0xa93b7ad4.
//
// Solidity: function previewKghDelegate(address validator) view returns(uint128)
func (_ValidatorManager *ValidatorManagerCaller) PreviewKghDelegate(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "previewKghDelegate", validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewKghDelegate is a free data retrieval call binding the contract method 0xa93b7ad4.
//
// Solidity: function previewKghDelegate(address validator) view returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) PreviewKghDelegate(validator common.Address) (*big.Int, error) {
	return _ValidatorManager.Contract.PreviewKghDelegate(&_ValidatorManager.CallOpts, validator)
}

// PreviewKghDelegate is a free data retrieval call binding the contract method 0xa93b7ad4.
//
// Solidity: function previewKghDelegate(address validator) view returns(uint128)
func (_ValidatorManager *ValidatorManagerCallerSession) PreviewKghDelegate(validator common.Address) (*big.Int, error) {
	return _ValidatorManager.Contract.PreviewKghDelegate(&_ValidatorManager.CallOpts, validator)
}

// PreviewKghUndelegate is a free data retrieval call binding the contract method 0x8cdfe8a9.
//
// Solidity: function previewKghUndelegate(address validator, uint256 tokenId) view returns(uint128)
func (_ValidatorManager *ValidatorManagerCaller) PreviewKghUndelegate(opts *bind.CallOpts, validator common.Address, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "previewKghUndelegate", validator, tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewKghUndelegate is a free data retrieval call binding the contract method 0x8cdfe8a9.
//
// Solidity: function previewKghUndelegate(address validator, uint256 tokenId) view returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) PreviewKghUndelegate(validator common.Address, tokenId *big.Int) (*big.Int, error) {
	return _ValidatorManager.Contract.PreviewKghUndelegate(&_ValidatorManager.CallOpts, validator, tokenId)
}

// PreviewKghUndelegate is a free data retrieval call binding the contract method 0x8cdfe8a9.
//
// Solidity: function previewKghUndelegate(address validator, uint256 tokenId) view returns(uint128)
func (_ValidatorManager *ValidatorManagerCallerSession) PreviewKghUndelegate(validator common.Address, tokenId *big.Int) (*big.Int, error) {
	return _ValidatorManager.Contract.PreviewKghUndelegate(&_ValidatorManager.CallOpts, validator, tokenId)
}

// PreviewUndelegate is a free data retrieval call binding the contract method 0x209a9694.
//
// Solidity: function previewUndelegate(address validator, uint128 shares) view returns(uint128)
func (_ValidatorManager *ValidatorManagerCaller) PreviewUndelegate(opts *bind.CallOpts, validator common.Address, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "previewUndelegate", validator, shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewUndelegate is a free data retrieval call binding the contract method 0x209a9694.
//
// Solidity: function previewUndelegate(address validator, uint128 shares) view returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) PreviewUndelegate(validator common.Address, shares *big.Int) (*big.Int, error) {
	return _ValidatorManager.Contract.PreviewUndelegate(&_ValidatorManager.CallOpts, validator, shares)
}

// PreviewUndelegate is a free data retrieval call binding the contract method 0x209a9694.
//
// Solidity: function previewUndelegate(address validator, uint128 shares) view returns(uint128)
func (_ValidatorManager *ValidatorManagerCallerSession) PreviewUndelegate(validator common.Address, shares *big.Int) (*big.Int, error) {
	return _ValidatorManager.Contract.PreviewUndelegate(&_ValidatorManager.CallOpts, validator, shares)
}

// StartedValidatorCount is a free data retrieval call binding the contract method 0xeb2ad8cb.
//
// Solidity: function startedValidatorCount() view returns(uint32)
func (_ValidatorManager *ValidatorManagerCaller) StartedValidatorCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "startedValidatorCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// StartedValidatorCount is a free data retrieval call binding the contract method 0xeb2ad8cb.
//
// Solidity: function startedValidatorCount() view returns(uint32)
func (_ValidatorManager *ValidatorManagerSession) StartedValidatorCount() (uint32, error) {
	return _ValidatorManager.Contract.StartedValidatorCount(&_ValidatorManager.CallOpts)
}

// StartedValidatorCount is a free data retrieval call binding the contract method 0xeb2ad8cb.
//
// Solidity: function startedValidatorCount() view returns(uint32)
func (_ValidatorManager *ValidatorManagerCallerSession) StartedValidatorCount() (uint32, error) {
	return _ValidatorManager.Contract.StartedValidatorCount(&_ValidatorManager.CallOpts)
}

// StartedValidatorTotalWeight is a free data retrieval call binding the contract method 0x1edbc580.
//
// Solidity: function startedValidatorTotalWeight() view returns(uint120)
func (_ValidatorManager *ValidatorManagerCaller) StartedValidatorTotalWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "startedValidatorTotalWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartedValidatorTotalWeight is a free data retrieval call binding the contract method 0x1edbc580.
//
// Solidity: function startedValidatorTotalWeight() view returns(uint120)
func (_ValidatorManager *ValidatorManagerSession) StartedValidatorTotalWeight() (*big.Int, error) {
	return _ValidatorManager.Contract.StartedValidatorTotalWeight(&_ValidatorManager.CallOpts)
}

// StartedValidatorTotalWeight is a free data retrieval call binding the contract method 0x1edbc580.
//
// Solidity: function startedValidatorTotalWeight() view returns(uint120)
func (_ValidatorManager *ValidatorManagerCallerSession) StartedValidatorTotalWeight() (*big.Int, error) {
	return _ValidatorManager.Contract.StartedValidatorTotalWeight(&_ValidatorManager.CallOpts)
}

// TotalKghNum is a free data retrieval call binding the contract method 0x913f1a9f.
//
// Solidity: function totalKghNum(address validator) view returns(uint128)
func (_ValidatorManager *ValidatorManagerCaller) TotalKghNum(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "totalKghNum", validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalKghNum is a free data retrieval call binding the contract method 0x913f1a9f.
//
// Solidity: function totalKghNum(address validator) view returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) TotalKghNum(validator common.Address) (*big.Int, error) {
	return _ValidatorManager.Contract.TotalKghNum(&_ValidatorManager.CallOpts, validator)
}

// TotalKghNum is a free data retrieval call binding the contract method 0x913f1a9f.
//
// Solidity: function totalKghNum(address validator) view returns(uint128)
func (_ValidatorManager *ValidatorManagerCallerSession) TotalKghNum(validator common.Address) (*big.Int, error) {
	return _ValidatorManager.Contract.TotalKghNum(&_ValidatorManager.CallOpts, validator)
}

// TotalKroAssets is a free data retrieval call binding the contract method 0x6b9ffeac.
//
// Solidity: function totalKroAssets(address validator) view returns(uint128)
func (_ValidatorManager *ValidatorManagerCaller) TotalKroAssets(opts *bind.CallOpts, validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "totalKroAssets", validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalKroAssets is a free data retrieval call binding the contract method 0x6b9ffeac.
//
// Solidity: function totalKroAssets(address validator) view returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) TotalKroAssets(validator common.Address) (*big.Int, error) {
	return _ValidatorManager.Contract.TotalKroAssets(&_ValidatorManager.CallOpts, validator)
}

// TotalKroAssets is a free data retrieval call binding the contract method 0x6b9ffeac.
//
// Solidity: function totalKroAssets(address validator) view returns(uint128)
func (_ValidatorManager *ValidatorManagerCallerSession) TotalKroAssets(validator common.Address) (*big.Int, error) {
	return _ValidatorManager.Contract.TotalKroAssets(&_ValidatorManager.CallOpts, validator)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ValidatorManager *ValidatorManagerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ValidatorManager.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ValidatorManager *ValidatorManagerSession) Version() (string, error) {
	return _ValidatorManager.Contract.Version(&_ValidatorManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ValidatorManager *ValidatorManagerCallerSession) Version() (string, error) {
	return _ValidatorManager.Contract.Version(&_ValidatorManager.CallOpts)
}

// AfterSubmitL2Output is a paid mutator transaction binding the contract method 0xbe119347.
//
// Solidity: function afterSubmitL2Output(uint256 outputIndex) returns()
func (_ValidatorManager *ValidatorManagerTransactor) AfterSubmitL2Output(opts *bind.TransactOpts, outputIndex *big.Int) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "afterSubmitL2Output", outputIndex)
}

// AfterSubmitL2Output is a paid mutator transaction binding the contract method 0xbe119347.
//
// Solidity: function afterSubmitL2Output(uint256 outputIndex) returns()
func (_ValidatorManager *ValidatorManagerSession) AfterSubmitL2Output(outputIndex *big.Int) (*types.Transaction, error) {
	return _ValidatorManager.Contract.AfterSubmitL2Output(&_ValidatorManager.TransactOpts, outputIndex)
}

// AfterSubmitL2Output is a paid mutator transaction binding the contract method 0xbe119347.
//
// Solidity: function afterSubmitL2Output(uint256 outputIndex) returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) AfterSubmitL2Output(outputIndex *big.Int) (*types.Transaction, error) {
	return _ValidatorManager.Contract.AfterSubmitL2Output(&_ValidatorManager.TransactOpts, outputIndex)
}

// ChangeCommissionRate is a paid mutator transaction binding the contract method 0x88576dc9.
//
// Solidity: function changeCommissionRate(uint8 newCommissionRate) returns()
func (_ValidatorManager *ValidatorManagerTransactor) ChangeCommissionRate(opts *bind.TransactOpts, newCommissionRate uint8) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "changeCommissionRate", newCommissionRate)
}

// ChangeCommissionRate is a paid mutator transaction binding the contract method 0x88576dc9.
//
// Solidity: function changeCommissionRate(uint8 newCommissionRate) returns()
func (_ValidatorManager *ValidatorManagerSession) ChangeCommissionRate(newCommissionRate uint8) (*types.Transaction, error) {
	return _ValidatorManager.Contract.ChangeCommissionRate(&_ValidatorManager.TransactOpts, newCommissionRate)
}

// ChangeCommissionRate is a paid mutator transaction binding the contract method 0x88576dc9.
//
// Solidity: function changeCommissionRate(uint8 newCommissionRate) returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) ChangeCommissionRate(newCommissionRate uint8) (*types.Transaction, error) {
	return _ValidatorManager.Contract.ChangeCommissionRate(&_ValidatorManager.TransactOpts, newCommissionRate)
}

// Delegate is a paid mutator transaction binding the contract method 0xa85120e4.
//
// Solidity: function delegate(address validator, uint128 assets) returns(uint128)
func (_ValidatorManager *ValidatorManagerTransactor) Delegate(opts *bind.TransactOpts, validator common.Address, assets *big.Int) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "delegate", validator, assets)
}

// Delegate is a paid mutator transaction binding the contract method 0xa85120e4.
//
// Solidity: function delegate(address validator, uint128 assets) returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) Delegate(validator common.Address, assets *big.Int) (*types.Transaction, error) {
	return _ValidatorManager.Contract.Delegate(&_ValidatorManager.TransactOpts, validator, assets)
}

// Delegate is a paid mutator transaction binding the contract method 0xa85120e4.
//
// Solidity: function delegate(address validator, uint128 assets) returns(uint128)
func (_ValidatorManager *ValidatorManagerTransactorSession) Delegate(validator common.Address, assets *big.Int) (*types.Transaction, error) {
	return _ValidatorManager.Contract.Delegate(&_ValidatorManager.TransactOpts, validator, assets)
}

// DelegateKgh is a paid mutator transaction binding the contract method 0x1f86f4f1.
//
// Solidity: function delegateKgh(address validator, uint256 tokenId) returns(uint128, uint128)
func (_ValidatorManager *ValidatorManagerTransactor) DelegateKgh(opts *bind.TransactOpts, validator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "delegateKgh", validator, tokenId)
}

// DelegateKgh is a paid mutator transaction binding the contract method 0x1f86f4f1.
//
// Solidity: function delegateKgh(address validator, uint256 tokenId) returns(uint128, uint128)
func (_ValidatorManager *ValidatorManagerSession) DelegateKgh(validator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ValidatorManager.Contract.DelegateKgh(&_ValidatorManager.TransactOpts, validator, tokenId)
}

// DelegateKgh is a paid mutator transaction binding the contract method 0x1f86f4f1.
//
// Solidity: function delegateKgh(address validator, uint256 tokenId) returns(uint128, uint128)
func (_ValidatorManager *ValidatorManagerTransactorSession) DelegateKgh(validator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ValidatorManager.Contract.DelegateKgh(&_ValidatorManager.TransactOpts, validator, tokenId)
}

// DelegateKghBatch is a paid mutator transaction binding the contract method 0x31d8e007.
//
// Solidity: function delegateKghBatch(address validator, uint256[] tokenIds) returns(uint128, uint128)
func (_ValidatorManager *ValidatorManagerTransactor) DelegateKghBatch(opts *bind.TransactOpts, validator common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "delegateKghBatch", validator, tokenIds)
}

// DelegateKghBatch is a paid mutator transaction binding the contract method 0x31d8e007.
//
// Solidity: function delegateKghBatch(address validator, uint256[] tokenIds) returns(uint128, uint128)
func (_ValidatorManager *ValidatorManagerSession) DelegateKghBatch(validator common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _ValidatorManager.Contract.DelegateKghBatch(&_ValidatorManager.TransactOpts, validator, tokenIds)
}

// DelegateKghBatch is a paid mutator transaction binding the contract method 0x31d8e007.
//
// Solidity: function delegateKghBatch(address validator, uint256[] tokenIds) returns(uint128, uint128)
func (_ValidatorManager *ValidatorManagerTransactorSession) DelegateKghBatch(validator common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _ValidatorManager.Contract.DelegateKghBatch(&_ValidatorManager.TransactOpts, validator, tokenIds)
}

// FinalizeClaimValidatorReward is a paid mutator transaction binding the contract method 0x0bc0b881.
//
// Solidity: function finalizeClaimValidatorReward() returns()
func (_ValidatorManager *ValidatorManagerTransactor) FinalizeClaimValidatorReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "finalizeClaimValidatorReward")
}

// FinalizeClaimValidatorReward is a paid mutator transaction binding the contract method 0x0bc0b881.
//
// Solidity: function finalizeClaimValidatorReward() returns()
func (_ValidatorManager *ValidatorManagerSession) FinalizeClaimValidatorReward() (*types.Transaction, error) {
	return _ValidatorManager.Contract.FinalizeClaimValidatorReward(&_ValidatorManager.TransactOpts)
}

// FinalizeClaimValidatorReward is a paid mutator transaction binding the contract method 0x0bc0b881.
//
// Solidity: function finalizeClaimValidatorReward() returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) FinalizeClaimValidatorReward() (*types.Transaction, error) {
	return _ValidatorManager.Contract.FinalizeClaimValidatorReward(&_ValidatorManager.TransactOpts)
}

// FinalizeUndelegate is a paid mutator transaction binding the contract method 0xb9f6131b.
//
// Solidity: function finalizeUndelegate(address validator) returns(uint128)
func (_ValidatorManager *ValidatorManagerTransactor) FinalizeUndelegate(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "finalizeUndelegate", validator)
}

// FinalizeUndelegate is a paid mutator transaction binding the contract method 0xb9f6131b.
//
// Solidity: function finalizeUndelegate(address validator) returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) FinalizeUndelegate(validator common.Address) (*types.Transaction, error) {
	return _ValidatorManager.Contract.FinalizeUndelegate(&_ValidatorManager.TransactOpts, validator)
}

// FinalizeUndelegate is a paid mutator transaction binding the contract method 0xb9f6131b.
//
// Solidity: function finalizeUndelegate(address validator) returns(uint128)
func (_ValidatorManager *ValidatorManagerTransactorSession) FinalizeUndelegate(validator common.Address) (*types.Transaction, error) {
	return _ValidatorManager.Contract.FinalizeUndelegate(&_ValidatorManager.TransactOpts, validator)
}

// FinalizeUndelegateKgh is a paid mutator transaction binding the contract method 0xb7aa324a.
//
// Solidity: function finalizeUndelegateKgh(address validator) returns(uint128)
func (_ValidatorManager *ValidatorManagerTransactor) FinalizeUndelegateKgh(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "finalizeUndelegateKgh", validator)
}

// FinalizeUndelegateKgh is a paid mutator transaction binding the contract method 0xb7aa324a.
//
// Solidity: function finalizeUndelegateKgh(address validator) returns(uint128)
func (_ValidatorManager *ValidatorManagerSession) FinalizeUndelegateKgh(validator common.Address) (*types.Transaction, error) {
	return _ValidatorManager.Contract.FinalizeUndelegateKgh(&_ValidatorManager.TransactOpts, validator)
}

// FinalizeUndelegateKgh is a paid mutator transaction binding the contract method 0xb7aa324a.
//
// Solidity: function finalizeUndelegateKgh(address validator) returns(uint128)
func (_ValidatorManager *ValidatorManagerTransactorSession) FinalizeUndelegateKgh(validator common.Address) (*types.Transaction, error) {
	return _ValidatorManager.Contract.FinalizeUndelegateKgh(&_ValidatorManager.TransactOpts, validator)
}

// InitClaimValidatorReward is a paid mutator transaction binding the contract method 0x2328bf42.
//
// Solidity: function initClaimValidatorReward(uint128 amount) returns()
func (_ValidatorManager *ValidatorManagerTransactor) InitClaimValidatorReward(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "initClaimValidatorReward", amount)
}

// InitClaimValidatorReward is a paid mutator transaction binding the contract method 0x2328bf42.
//
// Solidity: function initClaimValidatorReward(uint128 amount) returns()
func (_ValidatorManager *ValidatorManagerSession) InitClaimValidatorReward(amount *big.Int) (*types.Transaction, error) {
	return _ValidatorManager.Contract.InitClaimValidatorReward(&_ValidatorManager.TransactOpts, amount)
}

// InitClaimValidatorReward is a paid mutator transaction binding the contract method 0x2328bf42.
//
// Solidity: function initClaimValidatorReward(uint128 amount) returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) InitClaimValidatorReward(amount *big.Int) (*types.Transaction, error) {
	return _ValidatorManager.Contract.InitClaimValidatorReward(&_ValidatorManager.TransactOpts, amount)
}

// InitUndelegate is a paid mutator transaction binding the contract method 0x5dd0293b.
//
// Solidity: function initUndelegate(address validator, uint128 shares) returns()
func (_ValidatorManager *ValidatorManagerTransactor) InitUndelegate(opts *bind.TransactOpts, validator common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "initUndelegate", validator, shares)
}

// InitUndelegate is a paid mutator transaction binding the contract method 0x5dd0293b.
//
// Solidity: function initUndelegate(address validator, uint128 shares) returns()
func (_ValidatorManager *ValidatorManagerSession) InitUndelegate(validator common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ValidatorManager.Contract.InitUndelegate(&_ValidatorManager.TransactOpts, validator, shares)
}

// InitUndelegate is a paid mutator transaction binding the contract method 0x5dd0293b.
//
// Solidity: function initUndelegate(address validator, uint128 shares) returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) InitUndelegate(validator common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ValidatorManager.Contract.InitUndelegate(&_ValidatorManager.TransactOpts, validator, shares)
}

// InitUndelegateKgh is a paid mutator transaction binding the contract method 0x7cd68cd7.
//
// Solidity: function initUndelegateKgh(address validator, uint256 tokenId) returns()
func (_ValidatorManager *ValidatorManagerTransactor) InitUndelegateKgh(opts *bind.TransactOpts, validator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "initUndelegateKgh", validator, tokenId)
}

// InitUndelegateKgh is a paid mutator transaction binding the contract method 0x7cd68cd7.
//
// Solidity: function initUndelegateKgh(address validator, uint256 tokenId) returns()
func (_ValidatorManager *ValidatorManagerSession) InitUndelegateKgh(validator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ValidatorManager.Contract.InitUndelegateKgh(&_ValidatorManager.TransactOpts, validator, tokenId)
}

// InitUndelegateKgh is a paid mutator transaction binding the contract method 0x7cd68cd7.
//
// Solidity: function initUndelegateKgh(address validator, uint256 tokenId) returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) InitUndelegateKgh(validator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ValidatorManager.Contract.InitUndelegateKgh(&_ValidatorManager.TransactOpts, validator, tokenId)
}

// InitUndelegateKghBatch is a paid mutator transaction binding the contract method 0xa30c5f30.
//
// Solidity: function initUndelegateKghBatch(address validator, uint256[] tokenIds) returns()
func (_ValidatorManager *ValidatorManagerTransactor) InitUndelegateKghBatch(opts *bind.TransactOpts, validator common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "initUndelegateKghBatch", validator, tokenIds)
}

// InitUndelegateKghBatch is a paid mutator transaction binding the contract method 0xa30c5f30.
//
// Solidity: function initUndelegateKghBatch(address validator, uint256[] tokenIds) returns()
func (_ValidatorManager *ValidatorManagerSession) InitUndelegateKghBatch(validator common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _ValidatorManager.Contract.InitUndelegateKghBatch(&_ValidatorManager.TransactOpts, validator, tokenIds)
}

// InitUndelegateKghBatch is a paid mutator transaction binding the contract method 0xa30c5f30.
//
// Solidity: function initUndelegateKghBatch(address validator, uint256[] tokenIds) returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) InitUndelegateKghBatch(validator common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _ValidatorManager.Contract.InitUndelegateKghBatch(&_ValidatorManager.TransactOpts, validator, tokenIds)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x8ee4b602.
//
// Solidity: function registerValidator(uint128 assets, uint8 commissionRate, uint8 commissionMaxChangeRate) returns()
func (_ValidatorManager *ValidatorManagerTransactor) RegisterValidator(opts *bind.TransactOpts, assets *big.Int, commissionRate uint8, commissionMaxChangeRate uint8) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "registerValidator", assets, commissionRate, commissionMaxChangeRate)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x8ee4b602.
//
// Solidity: function registerValidator(uint128 assets, uint8 commissionRate, uint8 commissionMaxChangeRate) returns()
func (_ValidatorManager *ValidatorManagerSession) RegisterValidator(assets *big.Int, commissionRate uint8, commissionMaxChangeRate uint8) (*types.Transaction, error) {
	return _ValidatorManager.Contract.RegisterValidator(&_ValidatorManager.TransactOpts, assets, commissionRate, commissionMaxChangeRate)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x8ee4b602.
//
// Solidity: function registerValidator(uint128 assets, uint8 commissionRate, uint8 commissionMaxChangeRate) returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) RegisterValidator(assets *big.Int, commissionRate uint8, commissionMaxChangeRate uint8) (*types.Transaction, error) {
	return _ValidatorManager.Contract.RegisterValidator(&_ValidatorManager.TransactOpts, assets, commissionRate, commissionMaxChangeRate)
}

// Slash is a paid mutator transaction binding the contract method 0xe74f8239.
//
// Solidity: function slash(address loser, address winner, uint256 outputIndex) returns()
func (_ValidatorManager *ValidatorManagerTransactor) Slash(opts *bind.TransactOpts, loser common.Address, winner common.Address, outputIndex *big.Int) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "slash", loser, winner, outputIndex)
}

// Slash is a paid mutator transaction binding the contract method 0xe74f8239.
//
// Solidity: function slash(address loser, address winner, uint256 outputIndex) returns()
func (_ValidatorManager *ValidatorManagerSession) Slash(loser common.Address, winner common.Address, outputIndex *big.Int) (*types.Transaction, error) {
	return _ValidatorManager.Contract.Slash(&_ValidatorManager.TransactOpts, loser, winner, outputIndex)
}

// Slash is a paid mutator transaction binding the contract method 0xe74f8239.
//
// Solidity: function slash(address loser, address winner, uint256 outputIndex) returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) Slash(loser common.Address, winner common.Address, outputIndex *big.Int) (*types.Transaction, error) {
	return _ValidatorManager.Contract.Slash(&_ValidatorManager.TransactOpts, loser, winner, outputIndex)
}

// StartValidator is a paid mutator transaction binding the contract method 0x072df4cb.
//
// Solidity: function startValidator() returns()
func (_ValidatorManager *ValidatorManagerTransactor) StartValidator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "startValidator")
}

// StartValidator is a paid mutator transaction binding the contract method 0x072df4cb.
//
// Solidity: function startValidator() returns()
func (_ValidatorManager *ValidatorManagerSession) StartValidator() (*types.Transaction, error) {
	return _ValidatorManager.Contract.StartValidator(&_ValidatorManager.TransactOpts)
}

// StartValidator is a paid mutator transaction binding the contract method 0x072df4cb.
//
// Solidity: function startValidator() returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) StartValidator() (*types.Transaction, error) {
	return _ValidatorManager.Contract.StartValidator(&_ValidatorManager.TransactOpts)
}

// TryUnjail is a paid mutator transaction binding the contract method 0x7d2243b4.
//
// Solidity: function tryUnjail() returns()
func (_ValidatorManager *ValidatorManagerTransactor) TryUnjail(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorManager.contract.Transact(opts, "tryUnjail")
}

// TryUnjail is a paid mutator transaction binding the contract method 0x7d2243b4.
//
// Solidity: function tryUnjail() returns()
func (_ValidatorManager *ValidatorManagerSession) TryUnjail() (*types.Transaction, error) {
	return _ValidatorManager.Contract.TryUnjail(&_ValidatorManager.TransactOpts)
}

// TryUnjail is a paid mutator transaction binding the contract method 0x7d2243b4.
//
// Solidity: function tryUnjail() returns()
func (_ValidatorManager *ValidatorManagerTransactorSession) TryUnjail() (*types.Transaction, error) {
	return _ValidatorManager.Contract.TryUnjail(&_ValidatorManager.TransactOpts)
}

// ValidatorManagerChallengeRewardDistributedIterator is returned from FilterChallengeRewardDistributed and is used to iterate over the raw logs and unpacked data for ChallengeRewardDistributed events raised by the ValidatorManager contract.
type ValidatorManagerChallengeRewardDistributedIterator struct {
	Event *ValidatorManagerChallengeRewardDistributed // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerChallengeRewardDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerChallengeRewardDistributed)
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
		it.Event = new(ValidatorManagerChallengeRewardDistributed)
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
func (it *ValidatorManagerChallengeRewardDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerChallengeRewardDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerChallengeRewardDistributed represents a ChallengeRewardDistributed event raised by the ValidatorManager contract.
type ValidatorManagerChallengeRewardDistributed struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChallengeRewardDistributed is a free log retrieval operation binding the contract event 0x568d79fa2b3ed5751db3f4588be94b7eb2127a4696e56c68d8983a04ad0f3f50.
//
// Solidity: event ChallengeRewardDistributed(address indexed recipient, uint128 amount)
func (_ValidatorManager *ValidatorManagerFilterer) FilterChallengeRewardDistributed(opts *bind.FilterOpts, recipient []common.Address) (*ValidatorManagerChallengeRewardDistributedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "ChallengeRewardDistributed", recipientRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerChallengeRewardDistributedIterator{contract: _ValidatorManager.contract, event: "ChallengeRewardDistributed", logs: logs, sub: sub}, nil
}

// WatchChallengeRewardDistributed is a free log subscription operation binding the contract event 0x568d79fa2b3ed5751db3f4588be94b7eb2127a4696e56c68d8983a04ad0f3f50.
//
// Solidity: event ChallengeRewardDistributed(address indexed recipient, uint128 amount)
func (_ValidatorManager *ValidatorManagerFilterer) WatchChallengeRewardDistributed(opts *bind.WatchOpts, sink chan<- *ValidatorManagerChallengeRewardDistributed, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "ChallengeRewardDistributed", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerChallengeRewardDistributed)
				if err := _ValidatorManager.contract.UnpackLog(event, "ChallengeRewardDistributed", log); err != nil {
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

// ParseChallengeRewardDistributed is a log parse operation binding the contract event 0x568d79fa2b3ed5751db3f4588be94b7eb2127a4696e56c68d8983a04ad0f3f50.
//
// Solidity: event ChallengeRewardDistributed(address indexed recipient, uint128 amount)
func (_ValidatorManager *ValidatorManagerFilterer) ParseChallengeRewardDistributed(log types.Log) (*ValidatorManagerChallengeRewardDistributed, error) {
	event := new(ValidatorManagerChallengeRewardDistributed)
	if err := _ValidatorManager.contract.UnpackLog(event, "ChallengeRewardDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorManagerKghBatchDelegatedIterator is returned from FilterKghBatchDelegated and is used to iterate over the raw logs and unpacked data for KghBatchDelegated events raised by the ValidatorManager contract.
type ValidatorManagerKghBatchDelegatedIterator struct {
	Event *ValidatorManagerKghBatchDelegated // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerKghBatchDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerKghBatchDelegated)
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
		it.Event = new(ValidatorManagerKghBatchDelegated)
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
func (it *ValidatorManagerKghBatchDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerKghBatchDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerKghBatchDelegated represents a KghBatchDelegated event raised by the ValidatorManager contract.
type ValidatorManagerKghBatchDelegated struct {
	Validator common.Address
	Delegator common.Address
	TokenIds  []*big.Int
	KroShares *big.Int
	KghShares *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterKghBatchDelegated is a free log retrieval operation binding the contract event 0x32d1388f6c67f8101e5b61b8d17b70b87d15213fb6e6ebef744aaedfa4a41d1f.
//
// Solidity: event KghBatchDelegated(address indexed validator, address indexed delegator, uint256[] tokenIds, uint128 kroShares, uint128 kghShares)
func (_ValidatorManager *ValidatorManagerFilterer) FilterKghBatchDelegated(opts *bind.FilterOpts, validator []common.Address, delegator []common.Address) (*ValidatorManagerKghBatchDelegatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "KghBatchDelegated", validatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerKghBatchDelegatedIterator{contract: _ValidatorManager.contract, event: "KghBatchDelegated", logs: logs, sub: sub}, nil
}

// WatchKghBatchDelegated is a free log subscription operation binding the contract event 0x32d1388f6c67f8101e5b61b8d17b70b87d15213fb6e6ebef744aaedfa4a41d1f.
//
// Solidity: event KghBatchDelegated(address indexed validator, address indexed delegator, uint256[] tokenIds, uint128 kroShares, uint128 kghShares)
func (_ValidatorManager *ValidatorManagerFilterer) WatchKghBatchDelegated(opts *bind.WatchOpts, sink chan<- *ValidatorManagerKghBatchDelegated, validator []common.Address, delegator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "KghBatchDelegated", validatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerKghBatchDelegated)
				if err := _ValidatorManager.contract.UnpackLog(event, "KghBatchDelegated", log); err != nil {
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

// ParseKghBatchDelegated is a log parse operation binding the contract event 0x32d1388f6c67f8101e5b61b8d17b70b87d15213fb6e6ebef744aaedfa4a41d1f.
//
// Solidity: event KghBatchDelegated(address indexed validator, address indexed delegator, uint256[] tokenIds, uint128 kroShares, uint128 kghShares)
func (_ValidatorManager *ValidatorManagerFilterer) ParseKghBatchDelegated(log types.Log) (*ValidatorManagerKghBatchDelegated, error) {
	event := new(ValidatorManagerKghBatchDelegated)
	if err := _ValidatorManager.contract.UnpackLog(event, "KghBatchDelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorManagerKghBatchUndelegationInitiatedIterator is returned from FilterKghBatchUndelegationInitiated and is used to iterate over the raw logs and unpacked data for KghBatchUndelegationInitiated events raised by the ValidatorManager contract.
type ValidatorManagerKghBatchUndelegationInitiatedIterator struct {
	Event *ValidatorManagerKghBatchUndelegationInitiated // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerKghBatchUndelegationInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerKghBatchUndelegationInitiated)
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
		it.Event = new(ValidatorManagerKghBatchUndelegationInitiated)
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
func (it *ValidatorManagerKghBatchUndelegationInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerKghBatchUndelegationInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerKghBatchUndelegationInitiated represents a KghBatchUndelegationInitiated event raised by the ValidatorManager contract.
type ValidatorManagerKghBatchUndelegationInitiated struct {
	Validator common.Address
	Delegator common.Address
	TokenIds  []*big.Int
	KroShares *big.Int
	KghShares *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterKghBatchUndelegationInitiated is a free log retrieval operation binding the contract event 0xcb83bfd06f7180fe00601135ea03ce74be1a309cbdfd2e5a467bc7c42328f6c3.
//
// Solidity: event KghBatchUndelegationInitiated(address indexed validator, address indexed delegator, uint256[] tokenIds, uint128 kroShares, uint128 kghShares)
func (_ValidatorManager *ValidatorManagerFilterer) FilterKghBatchUndelegationInitiated(opts *bind.FilterOpts, validator []common.Address, delegator []common.Address) (*ValidatorManagerKghBatchUndelegationInitiatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "KghBatchUndelegationInitiated", validatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerKghBatchUndelegationInitiatedIterator{contract: _ValidatorManager.contract, event: "KghBatchUndelegationInitiated", logs: logs, sub: sub}, nil
}

// WatchKghBatchUndelegationInitiated is a free log subscription operation binding the contract event 0xcb83bfd06f7180fe00601135ea03ce74be1a309cbdfd2e5a467bc7c42328f6c3.
//
// Solidity: event KghBatchUndelegationInitiated(address indexed validator, address indexed delegator, uint256[] tokenIds, uint128 kroShares, uint128 kghShares)
func (_ValidatorManager *ValidatorManagerFilterer) WatchKghBatchUndelegationInitiated(opts *bind.WatchOpts, sink chan<- *ValidatorManagerKghBatchUndelegationInitiated, validator []common.Address, delegator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "KghBatchUndelegationInitiated", validatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerKghBatchUndelegationInitiated)
				if err := _ValidatorManager.contract.UnpackLog(event, "KghBatchUndelegationInitiated", log); err != nil {
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

// ParseKghBatchUndelegationInitiated is a log parse operation binding the contract event 0xcb83bfd06f7180fe00601135ea03ce74be1a309cbdfd2e5a467bc7c42328f6c3.
//
// Solidity: event KghBatchUndelegationInitiated(address indexed validator, address indexed delegator, uint256[] tokenIds, uint128 kroShares, uint128 kghShares)
func (_ValidatorManager *ValidatorManagerFilterer) ParseKghBatchUndelegationInitiated(log types.Log) (*ValidatorManagerKghBatchUndelegationInitiated, error) {
	event := new(ValidatorManagerKghBatchUndelegationInitiated)
	if err := _ValidatorManager.contract.UnpackLog(event, "KghBatchUndelegationInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorManagerKghDelegatedIterator is returned from FilterKghDelegated and is used to iterate over the raw logs and unpacked data for KghDelegated events raised by the ValidatorManager contract.
type ValidatorManagerKghDelegatedIterator struct {
	Event *ValidatorManagerKghDelegated // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerKghDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerKghDelegated)
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
		it.Event = new(ValidatorManagerKghDelegated)
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
func (it *ValidatorManagerKghDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerKghDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerKghDelegated represents a KghDelegated event raised by the ValidatorManager contract.
type ValidatorManagerKghDelegated struct {
	Validator common.Address
	Delegator common.Address
	TokenId   *big.Int
	KroInKgh  *big.Int
	KroShares *big.Int
	KghShares *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterKghDelegated is a free log retrieval operation binding the contract event 0x5b545a208eb5f51c3900b6fbf02a83cfdafcfbd2bae035a706827fceef97be8b.
//
// Solidity: event KghDelegated(address indexed validator, address indexed delegator, uint256 tokenId, uint128 kroInKgh, uint128 kroShares, uint128 kghShares)
func (_ValidatorManager *ValidatorManagerFilterer) FilterKghDelegated(opts *bind.FilterOpts, validator []common.Address, delegator []common.Address) (*ValidatorManagerKghDelegatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "KghDelegated", validatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerKghDelegatedIterator{contract: _ValidatorManager.contract, event: "KghDelegated", logs: logs, sub: sub}, nil
}

// WatchKghDelegated is a free log subscription operation binding the contract event 0x5b545a208eb5f51c3900b6fbf02a83cfdafcfbd2bae035a706827fceef97be8b.
//
// Solidity: event KghDelegated(address indexed validator, address indexed delegator, uint256 tokenId, uint128 kroInKgh, uint128 kroShares, uint128 kghShares)
func (_ValidatorManager *ValidatorManagerFilterer) WatchKghDelegated(opts *bind.WatchOpts, sink chan<- *ValidatorManagerKghDelegated, validator []common.Address, delegator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "KghDelegated", validatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerKghDelegated)
				if err := _ValidatorManager.contract.UnpackLog(event, "KghDelegated", log); err != nil {
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

// ParseKghDelegated is a log parse operation binding the contract event 0x5b545a208eb5f51c3900b6fbf02a83cfdafcfbd2bae035a706827fceef97be8b.
//
// Solidity: event KghDelegated(address indexed validator, address indexed delegator, uint256 tokenId, uint128 kroInKgh, uint128 kroShares, uint128 kghShares)
func (_ValidatorManager *ValidatorManagerFilterer) ParseKghDelegated(log types.Log) (*ValidatorManagerKghDelegated, error) {
	event := new(ValidatorManagerKghDelegated)
	if err := _ValidatorManager.contract.UnpackLog(event, "KghDelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorManagerKghUndelegationFinalizedIterator is returned from FilterKghUndelegationFinalized and is used to iterate over the raw logs and unpacked data for KghUndelegationFinalized events raised by the ValidatorManager contract.
type ValidatorManagerKghUndelegationFinalizedIterator struct {
	Event *ValidatorManagerKghUndelegationFinalized // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerKghUndelegationFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerKghUndelegationFinalized)
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
		it.Event = new(ValidatorManagerKghUndelegationFinalized)
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
func (it *ValidatorManagerKghUndelegationFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerKghUndelegationFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerKghUndelegationFinalized represents a KghUndelegationFinalized event raised by the ValidatorManager contract.
type ValidatorManagerKghUndelegationFinalized struct {
	Validator common.Address
	Delegator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterKghUndelegationFinalized is a free log retrieval operation binding the contract event 0xb8e63154a3a976a585c24f01fc46ead9ec9a0038ca84a74e3090e938c61fabe4.
//
// Solidity: event KghUndelegationFinalized(address indexed validator, address indexed delegator, uint128 amount)
func (_ValidatorManager *ValidatorManagerFilterer) FilterKghUndelegationFinalized(opts *bind.FilterOpts, validator []common.Address, delegator []common.Address) (*ValidatorManagerKghUndelegationFinalizedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "KghUndelegationFinalized", validatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerKghUndelegationFinalizedIterator{contract: _ValidatorManager.contract, event: "KghUndelegationFinalized", logs: logs, sub: sub}, nil
}

// WatchKghUndelegationFinalized is a free log subscription operation binding the contract event 0xb8e63154a3a976a585c24f01fc46ead9ec9a0038ca84a74e3090e938c61fabe4.
//
// Solidity: event KghUndelegationFinalized(address indexed validator, address indexed delegator, uint128 amount)
func (_ValidatorManager *ValidatorManagerFilterer) WatchKghUndelegationFinalized(opts *bind.WatchOpts, sink chan<- *ValidatorManagerKghUndelegationFinalized, validator []common.Address, delegator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "KghUndelegationFinalized", validatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerKghUndelegationFinalized)
				if err := _ValidatorManager.contract.UnpackLog(event, "KghUndelegationFinalized", log); err != nil {
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

// ParseKghUndelegationFinalized is a log parse operation binding the contract event 0xb8e63154a3a976a585c24f01fc46ead9ec9a0038ca84a74e3090e938c61fabe4.
//
// Solidity: event KghUndelegationFinalized(address indexed validator, address indexed delegator, uint128 amount)
func (_ValidatorManager *ValidatorManagerFilterer) ParseKghUndelegationFinalized(log types.Log) (*ValidatorManagerKghUndelegationFinalized, error) {
	event := new(ValidatorManagerKghUndelegationFinalized)
	if err := _ValidatorManager.contract.UnpackLog(event, "KghUndelegationFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorManagerKghUndelegationInitiatedIterator is returned from FilterKghUndelegationInitiated and is used to iterate over the raw logs and unpacked data for KghUndelegationInitiated events raised by the ValidatorManager contract.
type ValidatorManagerKghUndelegationInitiatedIterator struct {
	Event *ValidatorManagerKghUndelegationInitiated // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerKghUndelegationInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerKghUndelegationInitiated)
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
		it.Event = new(ValidatorManagerKghUndelegationInitiated)
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
func (it *ValidatorManagerKghUndelegationInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerKghUndelegationInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerKghUndelegationInitiated represents a KghUndelegationInitiated event raised by the ValidatorManager contract.
type ValidatorManagerKghUndelegationInitiated struct {
	Validator common.Address
	Delegator common.Address
	TokenId   *big.Int
	KroShares *big.Int
	KghShares *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterKghUndelegationInitiated is a free log retrieval operation binding the contract event 0x4d45e37032bac00ed9fe936ed54d9751ff10d249aec97a5dcd23a769a6be31fc.
//
// Solidity: event KghUndelegationInitiated(address indexed validator, address indexed delegator, uint256 tokenId, uint128 kroShares, uint128 kghShares)
func (_ValidatorManager *ValidatorManagerFilterer) FilterKghUndelegationInitiated(opts *bind.FilterOpts, validator []common.Address, delegator []common.Address) (*ValidatorManagerKghUndelegationInitiatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "KghUndelegationInitiated", validatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerKghUndelegationInitiatedIterator{contract: _ValidatorManager.contract, event: "KghUndelegationInitiated", logs: logs, sub: sub}, nil
}

// WatchKghUndelegationInitiated is a free log subscription operation binding the contract event 0x4d45e37032bac00ed9fe936ed54d9751ff10d249aec97a5dcd23a769a6be31fc.
//
// Solidity: event KghUndelegationInitiated(address indexed validator, address indexed delegator, uint256 tokenId, uint128 kroShares, uint128 kghShares)
func (_ValidatorManager *ValidatorManagerFilterer) WatchKghUndelegationInitiated(opts *bind.WatchOpts, sink chan<- *ValidatorManagerKghUndelegationInitiated, validator []common.Address, delegator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "KghUndelegationInitiated", validatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerKghUndelegationInitiated)
				if err := _ValidatorManager.contract.UnpackLog(event, "KghUndelegationInitiated", log); err != nil {
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

// ParseKghUndelegationInitiated is a log parse operation binding the contract event 0x4d45e37032bac00ed9fe936ed54d9751ff10d249aec97a5dcd23a769a6be31fc.
//
// Solidity: event KghUndelegationInitiated(address indexed validator, address indexed delegator, uint256 tokenId, uint128 kroShares, uint128 kghShares)
func (_ValidatorManager *ValidatorManagerFilterer) ParseKghUndelegationInitiated(log types.Log) (*ValidatorManagerKghUndelegationInitiated, error) {
	event := new(ValidatorManagerKghUndelegationInitiated)
	if err := _ValidatorManager.contract.UnpackLog(event, "KghUndelegationInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorManagerKroDelegatedIterator is returned from FilterKroDelegated and is used to iterate over the raw logs and unpacked data for KroDelegated events raised by the ValidatorManager contract.
type ValidatorManagerKroDelegatedIterator struct {
	Event *ValidatorManagerKroDelegated // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerKroDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerKroDelegated)
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
		it.Event = new(ValidatorManagerKroDelegated)
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
func (it *ValidatorManagerKroDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerKroDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerKroDelegated represents a KroDelegated event raised by the ValidatorManager contract.
type ValidatorManagerKroDelegated struct {
	Validator common.Address
	Delegator common.Address
	Amount    *big.Int
	Shares    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterKroDelegated is a free log retrieval operation binding the contract event 0x334cabe84b7338f2bdd62070c02f24ffbcc7735e46f425fa401db349717e1328.
//
// Solidity: event KroDelegated(address indexed validator, address indexed delegator, uint128 amount, uint128 shares)
func (_ValidatorManager *ValidatorManagerFilterer) FilterKroDelegated(opts *bind.FilterOpts, validator []common.Address, delegator []common.Address) (*ValidatorManagerKroDelegatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "KroDelegated", validatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerKroDelegatedIterator{contract: _ValidatorManager.contract, event: "KroDelegated", logs: logs, sub: sub}, nil
}

// WatchKroDelegated is a free log subscription operation binding the contract event 0x334cabe84b7338f2bdd62070c02f24ffbcc7735e46f425fa401db349717e1328.
//
// Solidity: event KroDelegated(address indexed validator, address indexed delegator, uint128 amount, uint128 shares)
func (_ValidatorManager *ValidatorManagerFilterer) WatchKroDelegated(opts *bind.WatchOpts, sink chan<- *ValidatorManagerKroDelegated, validator []common.Address, delegator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "KroDelegated", validatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerKroDelegated)
				if err := _ValidatorManager.contract.UnpackLog(event, "KroDelegated", log); err != nil {
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

// ParseKroDelegated is a log parse operation binding the contract event 0x334cabe84b7338f2bdd62070c02f24ffbcc7735e46f425fa401db349717e1328.
//
// Solidity: event KroDelegated(address indexed validator, address indexed delegator, uint128 amount, uint128 shares)
func (_ValidatorManager *ValidatorManagerFilterer) ParseKroDelegated(log types.Log) (*ValidatorManagerKroDelegated, error) {
	event := new(ValidatorManagerKroDelegated)
	if err := _ValidatorManager.contract.UnpackLog(event, "KroDelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorManagerKroUndelegationFinalizedIterator is returned from FilterKroUndelegationFinalized and is used to iterate over the raw logs and unpacked data for KroUndelegationFinalized events raised by the ValidatorManager contract.
type ValidatorManagerKroUndelegationFinalizedIterator struct {
	Event *ValidatorManagerKroUndelegationFinalized // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerKroUndelegationFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerKroUndelegationFinalized)
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
		it.Event = new(ValidatorManagerKroUndelegationFinalized)
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
func (it *ValidatorManagerKroUndelegationFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerKroUndelegationFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerKroUndelegationFinalized represents a KroUndelegationFinalized event raised by the ValidatorManager contract.
type ValidatorManagerKroUndelegationFinalized struct {
	Validator common.Address
	Delegator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterKroUndelegationFinalized is a free log retrieval operation binding the contract event 0x75cecc4cc21f0ebf6c86948f3a6a9bb934c49e4db83473fb8582ea706d135914.
//
// Solidity: event KroUndelegationFinalized(address indexed validator, address indexed delegator, uint128 amount)
func (_ValidatorManager *ValidatorManagerFilterer) FilterKroUndelegationFinalized(opts *bind.FilterOpts, validator []common.Address, delegator []common.Address) (*ValidatorManagerKroUndelegationFinalizedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "KroUndelegationFinalized", validatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerKroUndelegationFinalizedIterator{contract: _ValidatorManager.contract, event: "KroUndelegationFinalized", logs: logs, sub: sub}, nil
}

// WatchKroUndelegationFinalized is a free log subscription operation binding the contract event 0x75cecc4cc21f0ebf6c86948f3a6a9bb934c49e4db83473fb8582ea706d135914.
//
// Solidity: event KroUndelegationFinalized(address indexed validator, address indexed delegator, uint128 amount)
func (_ValidatorManager *ValidatorManagerFilterer) WatchKroUndelegationFinalized(opts *bind.WatchOpts, sink chan<- *ValidatorManagerKroUndelegationFinalized, validator []common.Address, delegator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "KroUndelegationFinalized", validatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerKroUndelegationFinalized)
				if err := _ValidatorManager.contract.UnpackLog(event, "KroUndelegationFinalized", log); err != nil {
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

// ParseKroUndelegationFinalized is a log parse operation binding the contract event 0x75cecc4cc21f0ebf6c86948f3a6a9bb934c49e4db83473fb8582ea706d135914.
//
// Solidity: event KroUndelegationFinalized(address indexed validator, address indexed delegator, uint128 amount)
func (_ValidatorManager *ValidatorManagerFilterer) ParseKroUndelegationFinalized(log types.Log) (*ValidatorManagerKroUndelegationFinalized, error) {
	event := new(ValidatorManagerKroUndelegationFinalized)
	if err := _ValidatorManager.contract.UnpackLog(event, "KroUndelegationFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorManagerKroUndelegationInitiatedIterator is returned from FilterKroUndelegationInitiated and is used to iterate over the raw logs and unpacked data for KroUndelegationInitiated events raised by the ValidatorManager contract.
type ValidatorManagerKroUndelegationInitiatedIterator struct {
	Event *ValidatorManagerKroUndelegationInitiated // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerKroUndelegationInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerKroUndelegationInitiated)
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
		it.Event = new(ValidatorManagerKroUndelegationInitiated)
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
func (it *ValidatorManagerKroUndelegationInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerKroUndelegationInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerKroUndelegationInitiated represents a KroUndelegationInitiated event raised by the ValidatorManager contract.
type ValidatorManagerKroUndelegationInitiated struct {
	Validator common.Address
	Delegator common.Address
	Amount    *big.Int
	Shares    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterKroUndelegationInitiated is a free log retrieval operation binding the contract event 0xcc7aa86d5cfd922cf42cb1b77d79a7e530bbaa3264669ca2145b564ecc1bf769.
//
// Solidity: event KroUndelegationInitiated(address indexed validator, address indexed delegator, uint128 amount, uint128 shares)
func (_ValidatorManager *ValidatorManagerFilterer) FilterKroUndelegationInitiated(opts *bind.FilterOpts, validator []common.Address, delegator []common.Address) (*ValidatorManagerKroUndelegationInitiatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "KroUndelegationInitiated", validatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerKroUndelegationInitiatedIterator{contract: _ValidatorManager.contract, event: "KroUndelegationInitiated", logs: logs, sub: sub}, nil
}

// WatchKroUndelegationInitiated is a free log subscription operation binding the contract event 0xcc7aa86d5cfd922cf42cb1b77d79a7e530bbaa3264669ca2145b564ecc1bf769.
//
// Solidity: event KroUndelegationInitiated(address indexed validator, address indexed delegator, uint128 amount, uint128 shares)
func (_ValidatorManager *ValidatorManagerFilterer) WatchKroUndelegationInitiated(opts *bind.WatchOpts, sink chan<- *ValidatorManagerKroUndelegationInitiated, validator []common.Address, delegator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "KroUndelegationInitiated", validatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerKroUndelegationInitiated)
				if err := _ValidatorManager.contract.UnpackLog(event, "KroUndelegationInitiated", log); err != nil {
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

// ParseKroUndelegationInitiated is a log parse operation binding the contract event 0xcc7aa86d5cfd922cf42cb1b77d79a7e530bbaa3264669ca2145b564ecc1bf769.
//
// Solidity: event KroUndelegationInitiated(address indexed validator, address indexed delegator, uint128 amount, uint128 shares)
func (_ValidatorManager *ValidatorManagerFilterer) ParseKroUndelegationInitiated(log types.Log) (*ValidatorManagerKroUndelegationInitiated, error) {
	event := new(ValidatorManagerKroUndelegationInitiated)
	if err := _ValidatorManager.contract.UnpackLog(event, "KroUndelegationInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorManagerRewardClaimFinalizedIterator is returned from FilterRewardClaimFinalized and is used to iterate over the raw logs and unpacked data for RewardClaimFinalized events raised by the ValidatorManager contract.
type ValidatorManagerRewardClaimFinalizedIterator struct {
	Event *ValidatorManagerRewardClaimFinalized // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerRewardClaimFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerRewardClaimFinalized)
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
		it.Event = new(ValidatorManagerRewardClaimFinalized)
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
func (it *ValidatorManagerRewardClaimFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerRewardClaimFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerRewardClaimFinalized represents a RewardClaimFinalized event raised by the ValidatorManager contract.
type ValidatorManagerRewardClaimFinalized struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimFinalized is a free log retrieval operation binding the contract event 0x668550d283aec3ba805f3e7c44d0bd95c1f847946fc605a7222d53394ca9e050.
//
// Solidity: event RewardClaimFinalized(address indexed validator, uint128 amount)
func (_ValidatorManager *ValidatorManagerFilterer) FilterRewardClaimFinalized(opts *bind.FilterOpts, validator []common.Address) (*ValidatorManagerRewardClaimFinalizedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "RewardClaimFinalized", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerRewardClaimFinalizedIterator{contract: _ValidatorManager.contract, event: "RewardClaimFinalized", logs: logs, sub: sub}, nil
}

// WatchRewardClaimFinalized is a free log subscription operation binding the contract event 0x668550d283aec3ba805f3e7c44d0bd95c1f847946fc605a7222d53394ca9e050.
//
// Solidity: event RewardClaimFinalized(address indexed validator, uint128 amount)
func (_ValidatorManager *ValidatorManagerFilterer) WatchRewardClaimFinalized(opts *bind.WatchOpts, sink chan<- *ValidatorManagerRewardClaimFinalized, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "RewardClaimFinalized", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerRewardClaimFinalized)
				if err := _ValidatorManager.contract.UnpackLog(event, "RewardClaimFinalized", log); err != nil {
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

// ParseRewardClaimFinalized is a log parse operation binding the contract event 0x668550d283aec3ba805f3e7c44d0bd95c1f847946fc605a7222d53394ca9e050.
//
// Solidity: event RewardClaimFinalized(address indexed validator, uint128 amount)
func (_ValidatorManager *ValidatorManagerFilterer) ParseRewardClaimFinalized(log types.Log) (*ValidatorManagerRewardClaimFinalized, error) {
	event := new(ValidatorManagerRewardClaimFinalized)
	if err := _ValidatorManager.contract.UnpackLog(event, "RewardClaimFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorManagerRewardClaimInitiatedIterator is returned from FilterRewardClaimInitiated and is used to iterate over the raw logs and unpacked data for RewardClaimInitiated events raised by the ValidatorManager contract.
type ValidatorManagerRewardClaimInitiatedIterator struct {
	Event *ValidatorManagerRewardClaimInitiated // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerRewardClaimInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerRewardClaimInitiated)
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
		it.Event = new(ValidatorManagerRewardClaimInitiated)
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
func (it *ValidatorManagerRewardClaimInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerRewardClaimInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerRewardClaimInitiated represents a RewardClaimInitiated event raised by the ValidatorManager contract.
type ValidatorManagerRewardClaimInitiated struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimInitiated is a free log retrieval operation binding the contract event 0x8ba9d52d538174baf05e0437c09555c022591dc617b7458fc02f53a1de9da490.
//
// Solidity: event RewardClaimInitiated(address indexed validator, uint128 amount)
func (_ValidatorManager *ValidatorManagerFilterer) FilterRewardClaimInitiated(opts *bind.FilterOpts, validator []common.Address) (*ValidatorManagerRewardClaimInitiatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "RewardClaimInitiated", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerRewardClaimInitiatedIterator{contract: _ValidatorManager.contract, event: "RewardClaimInitiated", logs: logs, sub: sub}, nil
}

// WatchRewardClaimInitiated is a free log subscription operation binding the contract event 0x8ba9d52d538174baf05e0437c09555c022591dc617b7458fc02f53a1de9da490.
//
// Solidity: event RewardClaimInitiated(address indexed validator, uint128 amount)
func (_ValidatorManager *ValidatorManagerFilterer) WatchRewardClaimInitiated(opts *bind.WatchOpts, sink chan<- *ValidatorManagerRewardClaimInitiated, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "RewardClaimInitiated", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerRewardClaimInitiated)
				if err := _ValidatorManager.contract.UnpackLog(event, "RewardClaimInitiated", log); err != nil {
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

// ParseRewardClaimInitiated is a log parse operation binding the contract event 0x8ba9d52d538174baf05e0437c09555c022591dc617b7458fc02f53a1de9da490.
//
// Solidity: event RewardClaimInitiated(address indexed validator, uint128 amount)
func (_ValidatorManager *ValidatorManagerFilterer) ParseRewardClaimInitiated(log types.Log) (*ValidatorManagerRewardClaimInitiated, error) {
	event := new(ValidatorManagerRewardClaimInitiated)
	if err := _ValidatorManager.contract.UnpackLog(event, "RewardClaimInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorManagerRewardDistributedIterator is returned from FilterRewardDistributed and is used to iterate over the raw logs and unpacked data for RewardDistributed events raised by the ValidatorManager contract.
type ValidatorManagerRewardDistributedIterator struct {
	Event *ValidatorManagerRewardDistributed // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerRewardDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerRewardDistributed)
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
		it.Event = new(ValidatorManagerRewardDistributed)
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
func (it *ValidatorManagerRewardDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerRewardDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerRewardDistributed represents a RewardDistributed event raised by the ValidatorManager contract.
type ValidatorManagerRewardDistributed struct {
	Validator       common.Address
	ValidatorReward *big.Int
	BaseReward      *big.Int
	BoostedReward   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRewardDistributed is a free log retrieval operation binding the contract event 0x36f11936e926f4c5f13247a0f85bfd1361293f182bc6a64bfff082b39aec64d9.
//
// Solidity: event RewardDistributed(address indexed validator, uint128 validatorReward, uint128 baseReward, uint128 boostedReward)
func (_ValidatorManager *ValidatorManagerFilterer) FilterRewardDistributed(opts *bind.FilterOpts, validator []common.Address) (*ValidatorManagerRewardDistributedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "RewardDistributed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerRewardDistributedIterator{contract: _ValidatorManager.contract, event: "RewardDistributed", logs: logs, sub: sub}, nil
}

// WatchRewardDistributed is a free log subscription operation binding the contract event 0x36f11936e926f4c5f13247a0f85bfd1361293f182bc6a64bfff082b39aec64d9.
//
// Solidity: event RewardDistributed(address indexed validator, uint128 validatorReward, uint128 baseReward, uint128 boostedReward)
func (_ValidatorManager *ValidatorManagerFilterer) WatchRewardDistributed(opts *bind.WatchOpts, sink chan<- *ValidatorManagerRewardDistributed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "RewardDistributed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerRewardDistributed)
				if err := _ValidatorManager.contract.UnpackLog(event, "RewardDistributed", log); err != nil {
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

// ParseRewardDistributed is a log parse operation binding the contract event 0x36f11936e926f4c5f13247a0f85bfd1361293f182bc6a64bfff082b39aec64d9.
//
// Solidity: event RewardDistributed(address indexed validator, uint128 validatorReward, uint128 baseReward, uint128 boostedReward)
func (_ValidatorManager *ValidatorManagerFilterer) ParseRewardDistributed(log types.Log) (*ValidatorManagerRewardDistributed, error) {
	event := new(ValidatorManagerRewardDistributed)
	if err := _ValidatorManager.contract.UnpackLog(event, "RewardDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorManagerSlashedIterator is returned from FilterSlashed and is used to iterate over the raw logs and unpacked data for Slashed events raised by the ValidatorManager contract.
type ValidatorManagerSlashedIterator struct {
	Event *ValidatorManagerSlashed // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerSlashed)
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
		it.Event = new(ValidatorManagerSlashed)
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
func (it *ValidatorManagerSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerSlashed represents a Slashed event raised by the ValidatorManager contract.
type ValidatorManagerSlashed struct {
	Loser  common.Address
	Winner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSlashed is a free log retrieval operation binding the contract event 0xbfeaf055e3cc2126fdbf006eda97657a7a8f82248db4159264060f31dfa2e2d0.
//
// Solidity: event Slashed(address indexed loser, address indexed winner, uint128 amount)
func (_ValidatorManager *ValidatorManagerFilterer) FilterSlashed(opts *bind.FilterOpts, loser []common.Address, winner []common.Address) (*ValidatorManagerSlashedIterator, error) {

	var loserRule []interface{}
	for _, loserItem := range loser {
		loserRule = append(loserRule, loserItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "Slashed", loserRule, winnerRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerSlashedIterator{contract: _ValidatorManager.contract, event: "Slashed", logs: logs, sub: sub}, nil
}

// WatchSlashed is a free log subscription operation binding the contract event 0xbfeaf055e3cc2126fdbf006eda97657a7a8f82248db4159264060f31dfa2e2d0.
//
// Solidity: event Slashed(address indexed loser, address indexed winner, uint128 amount)
func (_ValidatorManager *ValidatorManagerFilterer) WatchSlashed(opts *bind.WatchOpts, sink chan<- *ValidatorManagerSlashed, loser []common.Address, winner []common.Address) (event.Subscription, error) {

	var loserRule []interface{}
	for _, loserItem := range loser {
		loserRule = append(loserRule, loserItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "Slashed", loserRule, winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerSlashed)
				if err := _ValidatorManager.contract.UnpackLog(event, "Slashed", log); err != nil {
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

// ParseSlashed is a log parse operation binding the contract event 0xbfeaf055e3cc2126fdbf006eda97657a7a8f82248db4159264060f31dfa2e2d0.
//
// Solidity: event Slashed(address indexed loser, address indexed winner, uint128 amount)
func (_ValidatorManager *ValidatorManagerFilterer) ParseSlashed(log types.Log) (*ValidatorManagerSlashed, error) {
	event := new(ValidatorManagerSlashed)
	if err := _ValidatorManager.contract.UnpackLog(event, "Slashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorManagerValidatorCommissionRateChangedIterator is returned from FilterValidatorCommissionRateChanged and is used to iterate over the raw logs and unpacked data for ValidatorCommissionRateChanged events raised by the ValidatorManager contract.
type ValidatorManagerValidatorCommissionRateChangedIterator struct {
	Event *ValidatorManagerValidatorCommissionRateChanged // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerValidatorCommissionRateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerValidatorCommissionRateChanged)
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
		it.Event = new(ValidatorManagerValidatorCommissionRateChanged)
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
func (it *ValidatorManagerValidatorCommissionRateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerValidatorCommissionRateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerValidatorCommissionRateChanged represents a ValidatorCommissionRateChanged event raised by the ValidatorManager contract.
type ValidatorManagerValidatorCommissionRateChanged struct {
	Validator         common.Address
	OldCommissionRate uint8
	NewCommissionRate uint8
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterValidatorCommissionRateChanged is a free log retrieval operation binding the contract event 0xc0b29b9b824f7a62d93fde5832bb8307fd62594d0a08d96d533407a0a147ec48.
//
// Solidity: event ValidatorCommissionRateChanged(address validator, uint8 oldCommissionRate, uint8 newCommissionRate)
func (_ValidatorManager *ValidatorManagerFilterer) FilterValidatorCommissionRateChanged(opts *bind.FilterOpts) (*ValidatorManagerValidatorCommissionRateChangedIterator, error) {

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "ValidatorCommissionRateChanged")
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerValidatorCommissionRateChangedIterator{contract: _ValidatorManager.contract, event: "ValidatorCommissionRateChanged", logs: logs, sub: sub}, nil
}

// WatchValidatorCommissionRateChanged is a free log subscription operation binding the contract event 0xc0b29b9b824f7a62d93fde5832bb8307fd62594d0a08d96d533407a0a147ec48.
//
// Solidity: event ValidatorCommissionRateChanged(address validator, uint8 oldCommissionRate, uint8 newCommissionRate)
func (_ValidatorManager *ValidatorManagerFilterer) WatchValidatorCommissionRateChanged(opts *bind.WatchOpts, sink chan<- *ValidatorManagerValidatorCommissionRateChanged) (event.Subscription, error) {

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "ValidatorCommissionRateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerValidatorCommissionRateChanged)
				if err := _ValidatorManager.contract.UnpackLog(event, "ValidatorCommissionRateChanged", log); err != nil {
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

// ParseValidatorCommissionRateChanged is a log parse operation binding the contract event 0xc0b29b9b824f7a62d93fde5832bb8307fd62594d0a08d96d533407a0a147ec48.
//
// Solidity: event ValidatorCommissionRateChanged(address validator, uint8 oldCommissionRate, uint8 newCommissionRate)
func (_ValidatorManager *ValidatorManagerFilterer) ParseValidatorCommissionRateChanged(log types.Log) (*ValidatorManagerValidatorCommissionRateChanged, error) {
	event := new(ValidatorManagerValidatorCommissionRateChanged)
	if err := _ValidatorManager.contract.UnpackLog(event, "ValidatorCommissionRateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorManagerValidatorJailedIterator is returned from FilterValidatorJailed and is used to iterate over the raw logs and unpacked data for ValidatorJailed events raised by the ValidatorManager contract.
type ValidatorManagerValidatorJailedIterator struct {
	Event *ValidatorManagerValidatorJailed // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerValidatorJailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerValidatorJailed)
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
		it.Event = new(ValidatorManagerValidatorJailed)
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
func (it *ValidatorManagerValidatorJailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerValidatorJailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerValidatorJailed represents a ValidatorJailed event raised by the ValidatorManager contract.
type ValidatorManagerValidatorJailed struct {
	Validator common.Address
	ExpiresAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorJailed is a free log retrieval operation binding the contract event 0x95a398f2b6b2ad94f281708c97fe502386fc16adca43daed577a1e992a4cc814.
//
// Solidity: event ValidatorJailed(address indexed validator, uint128 expiresAt)
func (_ValidatorManager *ValidatorManagerFilterer) FilterValidatorJailed(opts *bind.FilterOpts, validator []common.Address) (*ValidatorManagerValidatorJailedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "ValidatorJailed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerValidatorJailedIterator{contract: _ValidatorManager.contract, event: "ValidatorJailed", logs: logs, sub: sub}, nil
}

// WatchValidatorJailed is a free log subscription operation binding the contract event 0x95a398f2b6b2ad94f281708c97fe502386fc16adca43daed577a1e992a4cc814.
//
// Solidity: event ValidatorJailed(address indexed validator, uint128 expiresAt)
func (_ValidatorManager *ValidatorManagerFilterer) WatchValidatorJailed(opts *bind.WatchOpts, sink chan<- *ValidatorManagerValidatorJailed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "ValidatorJailed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerValidatorJailed)
				if err := _ValidatorManager.contract.UnpackLog(event, "ValidatorJailed", log); err != nil {
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

// ParseValidatorJailed is a log parse operation binding the contract event 0x95a398f2b6b2ad94f281708c97fe502386fc16adca43daed577a1e992a4cc814.
//
// Solidity: event ValidatorJailed(address indexed validator, uint128 expiresAt)
func (_ValidatorManager *ValidatorManagerFilterer) ParseValidatorJailed(log types.Log) (*ValidatorManagerValidatorJailed, error) {
	event := new(ValidatorManagerValidatorJailed)
	if err := _ValidatorManager.contract.UnpackLog(event, "ValidatorJailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorManagerValidatorRegisteredIterator is returned from FilterValidatorRegistered and is used to iterate over the raw logs and unpacked data for ValidatorRegistered events raised by the ValidatorManager contract.
type ValidatorManagerValidatorRegisteredIterator struct {
	Event *ValidatorManagerValidatorRegistered // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerValidatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerValidatorRegistered)
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
		it.Event = new(ValidatorManagerValidatorRegistered)
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
func (it *ValidatorManagerValidatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerValidatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerValidatorRegistered represents a ValidatorRegistered event raised by the ValidatorManager contract.
type ValidatorManagerValidatorRegistered struct {
	Validator               common.Address
	Started                 bool
	CommissionRate          uint8
	CommissionMaxChangeRate uint8
	Assets                  *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegistered is a free log retrieval operation binding the contract event 0x04ba0c4d7cbac9138f7b73ec9fef796e4ad320bf5fb204f080f81fd59c2d48b9.
//
// Solidity: event ValidatorRegistered(address indexed validator, bool indexed started, uint8 commissionRate, uint8 commissionMaxChangeRate, uint128 assets)
func (_ValidatorManager *ValidatorManagerFilterer) FilterValidatorRegistered(opts *bind.FilterOpts, validator []common.Address, started []bool) (*ValidatorManagerValidatorRegisteredIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var startedRule []interface{}
	for _, startedItem := range started {
		startedRule = append(startedRule, startedItem)
	}

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "ValidatorRegistered", validatorRule, startedRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerValidatorRegisteredIterator{contract: _ValidatorManager.contract, event: "ValidatorRegistered", logs: logs, sub: sub}, nil
}

// WatchValidatorRegistered is a free log subscription operation binding the contract event 0x04ba0c4d7cbac9138f7b73ec9fef796e4ad320bf5fb204f080f81fd59c2d48b9.
//
// Solidity: event ValidatorRegistered(address indexed validator, bool indexed started, uint8 commissionRate, uint8 commissionMaxChangeRate, uint128 assets)
func (_ValidatorManager *ValidatorManagerFilterer) WatchValidatorRegistered(opts *bind.WatchOpts, sink chan<- *ValidatorManagerValidatorRegistered, validator []common.Address, started []bool) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var startedRule []interface{}
	for _, startedItem := range started {
		startedRule = append(startedRule, startedItem)
	}

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "ValidatorRegistered", validatorRule, startedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerValidatorRegistered)
				if err := _ValidatorManager.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
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

// ParseValidatorRegistered is a log parse operation binding the contract event 0x04ba0c4d7cbac9138f7b73ec9fef796e4ad320bf5fb204f080f81fd59c2d48b9.
//
// Solidity: event ValidatorRegistered(address indexed validator, bool indexed started, uint8 commissionRate, uint8 commissionMaxChangeRate, uint128 assets)
func (_ValidatorManager *ValidatorManagerFilterer) ParseValidatorRegistered(log types.Log) (*ValidatorManagerValidatorRegistered, error) {
	event := new(ValidatorManagerValidatorRegistered)
	if err := _ValidatorManager.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorManagerValidatorStartedIterator is returned from FilterValidatorStarted and is used to iterate over the raw logs and unpacked data for ValidatorStarted events raised by the ValidatorManager contract.
type ValidatorManagerValidatorStartedIterator struct {
	Event *ValidatorManagerValidatorStarted // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerValidatorStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerValidatorStarted)
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
		it.Event = new(ValidatorManagerValidatorStarted)
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
func (it *ValidatorManagerValidatorStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerValidatorStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerValidatorStarted represents a ValidatorStarted event raised by the ValidatorManager contract.
type ValidatorManagerValidatorStarted struct {
	Validator common.Address
	StartsAt  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorStarted is a free log retrieval operation binding the contract event 0xe8e4e936783175825bcf08ad234ab704ad447aeda363141c88312a07a729d067.
//
// Solidity: event ValidatorStarted(address indexed validator, uint256 startsAt)
func (_ValidatorManager *ValidatorManagerFilterer) FilterValidatorStarted(opts *bind.FilterOpts, validator []common.Address) (*ValidatorManagerValidatorStartedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "ValidatorStarted", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerValidatorStartedIterator{contract: _ValidatorManager.contract, event: "ValidatorStarted", logs: logs, sub: sub}, nil
}

// WatchValidatorStarted is a free log subscription operation binding the contract event 0xe8e4e936783175825bcf08ad234ab704ad447aeda363141c88312a07a729d067.
//
// Solidity: event ValidatorStarted(address indexed validator, uint256 startsAt)
func (_ValidatorManager *ValidatorManagerFilterer) WatchValidatorStarted(opts *bind.WatchOpts, sink chan<- *ValidatorManagerValidatorStarted, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "ValidatorStarted", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerValidatorStarted)
				if err := _ValidatorManager.contract.UnpackLog(event, "ValidatorStarted", log); err != nil {
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

// ParseValidatorStarted is a log parse operation binding the contract event 0xe8e4e936783175825bcf08ad234ab704ad447aeda363141c88312a07a729d067.
//
// Solidity: event ValidatorStarted(address indexed validator, uint256 startsAt)
func (_ValidatorManager *ValidatorManagerFilterer) ParseValidatorStarted(log types.Log) (*ValidatorManagerValidatorStarted, error) {
	event := new(ValidatorManagerValidatorStarted)
	if err := _ValidatorManager.contract.UnpackLog(event, "ValidatorStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorManagerValidatorUnjailedIterator is returned from FilterValidatorUnjailed and is used to iterate over the raw logs and unpacked data for ValidatorUnjailed events raised by the ValidatorManager contract.
type ValidatorManagerValidatorUnjailedIterator struct {
	Event *ValidatorManagerValidatorUnjailed // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerValidatorUnjailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerValidatorUnjailed)
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
		it.Event = new(ValidatorManagerValidatorUnjailed)
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
func (it *ValidatorManagerValidatorUnjailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerValidatorUnjailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerValidatorUnjailed represents a ValidatorUnjailed event raised by the ValidatorManager contract.
type ValidatorManagerValidatorUnjailed struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorUnjailed is a free log retrieval operation binding the contract event 0x9390b453426557da5ebdc31f19a37753ca04addf656d32f35232211bb2af3f19.
//
// Solidity: event ValidatorUnjailed(address validator)
func (_ValidatorManager *ValidatorManagerFilterer) FilterValidatorUnjailed(opts *bind.FilterOpts) (*ValidatorManagerValidatorUnjailedIterator, error) {

	logs, sub, err := _ValidatorManager.contract.FilterLogs(opts, "ValidatorUnjailed")
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerValidatorUnjailedIterator{contract: _ValidatorManager.contract, event: "ValidatorUnjailed", logs: logs, sub: sub}, nil
}

// WatchValidatorUnjailed is a free log subscription operation binding the contract event 0x9390b453426557da5ebdc31f19a37753ca04addf656d32f35232211bb2af3f19.
//
// Solidity: event ValidatorUnjailed(address validator)
func (_ValidatorManager *ValidatorManagerFilterer) WatchValidatorUnjailed(opts *bind.WatchOpts, sink chan<- *ValidatorManagerValidatorUnjailed) (event.Subscription, error) {

	logs, sub, err := _ValidatorManager.contract.WatchLogs(opts, "ValidatorUnjailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerValidatorUnjailed)
				if err := _ValidatorManager.contract.UnpackLog(event, "ValidatorUnjailed", log); err != nil {
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

// ParseValidatorUnjailed is a log parse operation binding the contract event 0x9390b453426557da5ebdc31f19a37753ca04addf656d32f35232211bb2af3f19.
//
// Solidity: event ValidatorUnjailed(address validator)
func (_ValidatorManager *ValidatorManagerFilterer) ParseValidatorUnjailed(log types.Log) (*ValidatorManagerValidatorUnjailed, error) {
	event := new(ValidatorManagerValidatorUnjailed)
	if err := _ValidatorManager.contract.UnpackLog(event, "ValidatorUnjailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
