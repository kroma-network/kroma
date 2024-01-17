package state

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/log"

	opstate "github.com/ethereum-optimism/optimism/op-chain-ops/state"

	"github.com/kroma-network/kroma/kroma-bindings/bindings"
)

// SetStorage will set the storage values in a db given a contract name,
// address and the storage values
func SetStorage(name string, address common.Address, values opstate.StorageValues, db vm.StateDB) error {
	layout, err := bindings.GetStorageLayout(name)
	if err != nil {
		return fmt.Errorf("cannot set storage: %w", err)
	}
	slots, err := opstate.ComputeStorageSlots(layout, values)
	if err != nil {
		return fmt.Errorf("%s: %w", name, err)
	}
	for _, slot := range slots {
		db.SetState(address, slot.Key, slot.Value)
		log.Trace("setting storage", "address", address.Hex(), "key", slot.Key.Hex(), "value", slot.Value.Hex())
	}
	return nil
}
