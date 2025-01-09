package bindings

import (
	"fmt"
	"strings"

	"github.com/ethereum-optimism/superchain-registry/superchain"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

// layouts represents the set of storage layouts. It is populated in an init function.
var layouts = make(map[string]*solc.StorageLayout)

// deployedBytecodes represents the set of deployed bytecodes. It is populated
// in an init function.
var deployedBytecodes = make(map[string]string)

// immutableReferences represents the set of immutable references. It is populated
// in an init function.
var immutableReferences = make(map[string]bool)

// Create2DeployerCodeHash represents the codehash of the Create2Deployer contract.
var Create2DeployerCodeHash = common.HexToHash("0xb0550b5b431e30d38000efb7107aaa0ade03d48a7198a140edda9d27134468b2")

func init() {
	code, err := superchain.LoadContractBytecode(superchain.Hash(Create2DeployerCodeHash))
	if err != nil {
		panic(err)
	}
	deployedBytecodes["Create2Deployer"] = common.Bytes2Hex(code)
}

// GetStorageLayout returns the storage layout of a contract by name.
func GetStorageLayout(name string) (*solc.StorageLayout, error) {
	layout := layouts[name]
	if layout == nil {
		return nil, fmt.Errorf("%s: storage layout not found", name)
	}
	return layout, nil
}

// GetDeployedBytecode returns the deployed bytecode of a contract by name.
func GetDeployedBytecode(name string) ([]byte, error) {
	bc := deployedBytecodes[name]
	if bc == "" {
		return nil, fmt.Errorf("%s: deployed bytecode not found", name)
	}

	if !isHex(bc) {
		return nil, fmt.Errorf("%s: invalid deployed bytecode", name)
	}

	return common.FromHex(bc), nil
}

// HasImmutableReferences returns the immutable references of a contract by name.
func HasImmutableReferences(name string) (bool, error) {
	has, ok := immutableReferences[name]
	if !ok {
		return false, fmt.Errorf("%s: immutable reference not found", name)
	}
	return has, nil
}

// isHexCharacter returns bool of c being a valid hexadecimal.
func isHexCharacter(c byte) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')
}

// isHex validates whether each byte is valid hexadecimal string.
func isHex(str string) bool {
	if len(str)%2 != 0 {
		return false
	}
	str = strings.TrimPrefix(str, "0x")

	for _, c := range []byte(str) {
		if !isHexCharacter(c) {
			return false
		}
	}
	return true
}
