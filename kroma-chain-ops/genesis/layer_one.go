package genesis

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum-optimism/optimism/op-chain-ops/state"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	gstate "github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"

	"github.com/kroma-network/kroma/kroma-bindings/bindings"
)

var (
	// uint128Max is type(uint128).max and is set in the init function.
	uint128Max = new(big.Int)
	// The default values for the ResourceConfig, used as part of
	// an EIP-1559 curve for deposit gas.
	DefaultResourceConfig = bindings.ResourceMeteringResourceConfig{
		MaxResourceLimit:            20_000_000,
		ElasticityMultiplier:        10,
		BaseFeeMaxChangeDenominator: 8,
		MinimumBaseFee:              params.GWei,
		SystemTxMaxGas:              1_000_000,
	}
)

func init() {
	var ok bool
	uint128Max, ok = new(big.Int).SetString("ffffffffffffffffffffffffffffffff", 16)
	if !ok {
		panic("bad uint128Max")
	}
	// Set the maximum base fee on the default config.
	DefaultResourceConfig.MaximumBaseFee = uint128Max
}

// BuildL1DeveloperGenesis will create a L1 genesis block after creating
// all of the state required for an Optimism network to function.
// It is expected that the dump contains all of the required state to bootstrap
// the L1 chain.
func BuildL1DeveloperGenesis(config *DeployConfig, dump *gstate.Dump, l1Deployments *L1Deployments, postProcess bool) (*core.Genesis, error) {
	log.Info("Building developer L1 genesis block")
	genesis, err := NewL1Genesis(config)
	if err != nil {
		return nil, fmt.Errorf("cannot create L1 developer genesis: %w", err)
	}

	memDB := state.NewMemoryStateDB(genesis)
	FundDevAccounts(memDB)
	SetPrecompileBalances(memDB)

	if dump != nil {
		for address, account := range dump.Accounts {
			name := "<unknown>"
			if l1Deployments != nil {
				if n := l1Deployments.GetName(address); n != "" {
					name = n
				}
			}
			log.Info("Setting account", "name", name, "address", address.Hex())
			memDB.CreateAccount(address)
			memDB.SetNonce(address, account.Nonce)

			balance, ok := new(big.Int).SetString(account.Balance, 10)
			if !ok {
				return nil, fmt.Errorf("failed to parse balance for %s", address)
			}
			memDB.AddBalance(address, balance)
			memDB.SetCode(address, account.Code)
			for key, value := range account.Storage {
				log.Info("Setting storage", "name", name, "key", key.Hex(), "value", value)
				memDB.SetState(address, key, common.HexToHash(value))
			}
		}

		// This should only be used if we are expecting Optimism specific state to be set
		if postProcess {
			if err := PostProcessL1DeveloperGenesis(memDB, l1Deployments); err != nil {
				return nil, fmt.Errorf("failed to post process L1 developer genesis: %w", err)
			}
		}
	}

	return memDB.Genesis(), nil
}

// PostProcessL1DeveloperGenesis will apply post processing to the L1 genesis
// state. This is required to handle edge cases in the genesis generation.
// `block.number` is used during deployment and without specifically setting
// the value to 0, it will cause underflow reverts for deposits in testing.
func PostProcessL1DeveloperGenesis(stateDB *state.MemoryStateDB, deployments *L1Deployments) error {
	log.Info("Post processing state")

	if stateDB == nil {
		return errors.New("cannot post process nil stateDB")
	}
	if deployments == nil {
		return errors.New("cannot post process dump with nil deployments")
	}

	if !stateDB.Exist(deployments.KromaPortalProxy) {
		return fmt.Errorf("portal proxy doesn't exist at %s", deployments.KromaPortalProxy)
	}

	// [Kroma: START]
	slot, err := getStorageSlot("KromaPortal", "params")
	if err != nil {
		return err
	}

	stateDB.SetState(deployments.KromaPortalProxy, slot, common.Hash{})
	log.Info("Post process update", "name", "KromaPortal", "address", deployments.KromaPortalProxy, "slot", slot.Hex(), "value", common.Hash{}.Hex())

	// Transfer ownership of SystemConfig to ProxyAdminOwner for test
	if !stateDB.Exist(deployments.SystemConfigProxy) {
		return fmt.Errorf("sysCfg proxy doesn't exist at %s", deployments.SystemConfigProxy)
	}

	slot, err = getStorageSlot("SystemConfig", "_owner")
	if err != nil {
		return err
	}

	val := stateDB.GetState(deployments.ProxyAdmin, common.BigToHash(common.Big0))
	stateDB.SetState(deployments.SystemConfigProxy, slot, val)
	log.Info("Post process update", "name", "SystemConfig", "address", deployments.SystemConfigProxy, "slot", slot.Hex(), "value", val.Hex())

	// Change the key of _quorumNumeratorHistory in UpgradeGovernor to 1 which means that quorumNumerator has been set at L1 block number 1 for guardian test
	if !stateDB.Exist(deployments.UpgradeGovernorProxy) {
		return fmt.Errorf("upgardeGovernor proxy doesn't exist at %s", deployments.UpgradeGovernorProxy)
	}

	slot, err = getStorageSlot("UpgradeGovernor", "_quorumNumeratorHistory")
	if err != nil {
		return err
	}
	slot = crypto.Keccak256Hash(slot.Bytes())

	beforeVal := stateDB.GetState(deployments.UpgradeGovernorProxy, slot)
	checkpointVal := make([]byte, 28)
	copy(checkpointVal, beforeVal[:28])
	checkpointKey := [4]byte{}
	checkpointKey[3] = 0x01
	val = common.BytesToHash(append(checkpointVal, checkpointKey[:]...))

	stateDB.SetState(deployments.UpgradeGovernorProxy, slot, val)
	log.Info("Post process update", "name", "UpgradeGovernor", "address", deployments.UpgradeGovernorProxy, "slot", slot.Hex(), "beforeVal", beforeVal.Hex(), "afterVal", val.Hex())

	// Change the keys of _totalCheckpoints in SecurityCouncilToken to 1 which means that tokens have been minted at L1 block number 1 for guardian test
	if !stateDB.Exist(deployments.SecurityCouncilTokenProxy) {
		return fmt.Errorf("securityCouncilToken proxy doesn't exist at %s", deployments.SecurityCouncilTokenProxy)
	}

	slot, err = getStorageSlot("SecurityCouncilToken", "_totalCheckpoints")
	if err != nil {
		return err
	}
	startSlot := new(big.Int).SetBytes(crypto.Keccak256(slot.Bytes()))

	mintedNum := stateDB.GetState(deployments.SecurityCouncilTokenProxy, slot).Big().Uint64()
	for i := 0; uint64(i) < mintedNum; i++ {
		slot = common.BigToHash(new(big.Int).Add(startSlot, big.NewInt(int64(i))))

		beforeVal = stateDB.GetState(deployments.SecurityCouncilTokenProxy, slot)
		checkpointVal = make([]byte, 28)
		copy(checkpointVal, beforeVal[:28])
		checkpointKey = [4]byte{}
		checkpointKey[3] = 0x01
		val = common.BytesToHash(append(checkpointVal, checkpointKey[:]...))

		stateDB.SetState(deployments.SecurityCouncilTokenProxy, slot, val)
		log.Info("Post process update", "name", "SecurityCouncilToken", "address", deployments.SecurityCouncilTokenProxy, "slot", slot.Hex(), "beforeVal", beforeVal.Hex(), "afterVal", val.Hex())
	}
	return nil
}

func getStorageSlot(contractName, entryName string) (common.Hash, error) {
	layout, err := bindings.GetStorageLayout(contractName)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get storage layout for %s", contractName)
	}

	entry, err := layout.GetStorageLayoutEntry(entryName)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get storage layout entry for %s.%s", contractName, entryName)
	}

	return common.BigToHash(big.NewInt(int64(entry.Slot))), nil
}

// [Kroma: END]
