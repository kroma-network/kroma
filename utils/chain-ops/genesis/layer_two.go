package genesis

import (
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/wemixkanvas/kanvas/bindings/predeploys"
	"github.com/wemixkanvas/kanvas/utils/chain-ops/state"
)

// BuildL2DeveloperGenesis will build the developer Kanvas Genesis
// Block. Suitable for devnets.
func BuildL2DeveloperGenesis(config *DeployConfig, l1StartBlock *types.Block, zktrie bool) (*core.Genesis, error) {
	genspec, err := NewL2Genesis(config, l1StartBlock, zktrie)
	if err != nil {
		return nil, err
	}

	db := state.NewMemoryStateDB(genspec)

	if config.FundDevAccounts {
		FundDevAccounts(db)
	}
	SetPrecompileBalances(db)

	storage, err := NewL2StorageConfig(config, l1StartBlock)
	if err != nil {
		return nil, err
	}

	immutable, err := NewL2ImmutableConfig(config, l1StartBlock)
	if err != nil {
		return nil, err
	}

	if err := SetL2Proxies(db); err != nil {
		return nil, err
	}

	if err := SetImplementations(db, storage, immutable, zktrie); err != nil {
		return nil, err
	}

	if err := SetDevOnlyL2Implementations(db, storage, immutable, zktrie); err != nil {
		return nil, err
	}

	return db.Genesis(), nil
}

func L2PredeploysCount(config *DeployConfig) int {
	cnt := PrecompiledCount + int(L2ProxyCount) + len(predeploys.Predeploys)
	if config.FundDevAccounts {
		cnt = cnt + len(DevAccounts)
	}

	return cnt
}
