package immutables

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/kroma-network/kroma/bindings/bindings"
	"github.com/kroma-network/kroma/bindings/predeploys"
	"github.com/kroma-network/kroma/utils/chain-ops/deployer"
)

// ImmutableValues represents the values to be set in immutable code.
// The key is the name of the variable and the value is the value to set in
// immutable code.
type ImmutableValues map[string]any

// ImmutableConfig represents the immutable configuration for the L2 predeploy
// contracts.
type ImmutableConfig map[string]ImmutableValues

// Check does a sanity check that the specific values that
// Kroma uses are set inside of the ImmutableConfig.
func (i ImmutableConfig) Check() error {
	if _, ok := i["L2CrossDomainMessenger"]["otherMessenger"]; !ok {
		return errors.New("L2CrossDomainMessenger otherMessenger not set")
	}
	if _, ok := i["L2StandardBridge"]["otherBridge"]; !ok {
		return errors.New("L2StandardBridge otherBridge not set")
	}
	if _, ok := i["L2ERC721Bridge"]["messenger"]; !ok {
		return errors.New("L2ERC721Bridge messenger not set")
	}
	if _, ok := i["L2ERC721Bridge"]["otherBridge"]; !ok {
		return errors.New("L2ERC721Bridge otherBridge not set")
	}
	if _, ok := i["KromaMintableERC721Factory"]["bridge"]; !ok {
		return errors.New("KromaMintableERC20Factory bridge not set")
	}
	if _, ok := i["KromaMintableERC721Factory"]["remoteChainId"]; !ok {
		return errors.New("KromaMintableERC20Factory remoteChainId not set")
	}
	if _, ok := i["ValidatorRewardVault"]["recipient"]; !ok {
		return errors.New("ValidatorRewardVault recipient not set")
	}
	if _, ok := i["ProposerRewardVault"]["recipient"]; !ok {
		return errors.New("ProposerRewardVault recipient not set")
	}
	if _, ok := i["ProtocolVault"]["recipient"]; !ok {
		return errors.New("ProtocolVault recipient not set")
	}
	return nil
}

// DeploymentResults represents the output of deploying each of the
// contracts so that the immutables can be set properly in the bytecode.
type DeploymentResults map[string]hexutil.Bytes

// BuildKroma will deploy the L2 predeploys so that their immutables are set
// correctly.
func BuildKroma(immutable ImmutableConfig, zktrie bool) (DeploymentResults, error) {
	if err := immutable.Check(); err != nil {
		return DeploymentResults{}, err
	}

	deployments := []deployer.Constructor{
		{
			Name: "GasPriceOracle",
		},
		{
			Name: "L1Block",
		},
		{
			Name: "L2CrossDomainMessenger",
			Args: []interface{}{
				immutable["L2CrossDomainMessenger"]["otherMessenger"],
			},
		},
		{
			Name: "L2StandardBridge",
			Args: []interface{}{
				immutable["L2StandardBridge"]["otherBridge"],
			},
		},
		{
			Name: "L2ToL1MessagePasser",
		},
		{
			Name: "ValidatorRewardVault",
			Args: []interface{}{
				immutable["ValidatorRewardVault"]["recipient"],
			},
		},
		{
			Name: "ProtocolVault",
			Args: []interface{}{
				immutable["ProtocolVault"]["recipient"],
			},
		},
		{
			Name: "ProposerRewardVault",
			Args: []interface{}{
				immutable["ProposerRewardVault"]["recipient"],
			},
		},
		{
			Name: "KromaMintableERC20Factory",
		},
		{
			Name: "L2ERC721Bridge",
			Args: []interface{}{
				predeploys.L2CrossDomainMessengerAddr,
				immutable["L2ERC721Bridge"]["otherBridge"],
			},
		},
		{
			Name: "KromaMintableERC721Factory",
			Args: []interface{}{
				predeploys.L2ERC721BridgeAddr,
				immutable["KromaMintableERC721Factory"]["remoteChainId"],
			},
		},
	}
	return BuildL2(deployments, zktrie)
}

// BuildL2 will deploy contracts to a simulated backend so that their immutables
// can be properly set. The bytecode returned in the results is suitable to be
// inserted into the state via state surgery.
func BuildL2(constructors []deployer.Constructor, zktrie bool) (DeploymentResults, error) {
	deployments, err := deployer.Deploy(deployer.NewBackend(zktrie), constructors, l2Deployer)
	if err != nil {
		return nil, err
	}
	results := make(DeploymentResults)
	for _, dep := range deployments {
		results[dep.Name] = dep.Bytecode
	}
	return results, nil
}

func l2Deployer(backend *backends.SimulatedBackend, opts *bind.TransactOpts, deployment deployer.Constructor) (*types.Transaction, error) {
	var tx *types.Transaction
	var err error
	switch deployment.Name {
	case "GasPriceOracle":
		_, tx, _, err = bindings.DeployGasPriceOracle(opts, backend)
	case "L1Block":
		// No arguments required for the L1Block contract
		_, tx, _, err = bindings.DeployL1Block(opts, backend)
	case "L2CrossDomainMessenger":
		otherMessenger, ok := deployment.Args[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for otherMessenger")
		}
		_, tx, _, err = bindings.DeployL2CrossDomainMessenger(opts, backend, otherMessenger)
	case "L2StandardBridge":
		otherBridge, ok := deployment.Args[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for otherBridge")
		}
		_, tx, _, err = bindings.DeployL2StandardBridge(opts, backend, otherBridge)
	case "L2ToL1MessagePasser":
		// No arguments required for L2ToL1MessagePasser
		_, tx, _, err = bindings.DeployL2ToL1MessagePasser(opts, backend)
	case "ValidatorRewardVault":
		recipient, ok := deployment.Args[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for recipient")
		}
		_, tx, _, err = bindings.DeployValidatorRewardVault(opts, backend, recipient)
	case "ProtocolVault":
		recipient, ok := deployment.Args[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for recipient")
		}
		_, tx, _, err = bindings.DeployProtocolVault(opts, backend, recipient)
	case "ProposerRewardVault":
		recipient, ok := deployment.Args[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for recipient")
		}
		_, tx, _, err = bindings.DeployProposerRewardVault(opts, backend, recipient)
	case "KromaMintableERC20Factory":
		_, tx, _, err = bindings.DeployKromaMintableERC20Factory(opts, backend, predeploys.L2StandardBridgeAddr)
	case "L2ERC721Bridge":
		// TODO(tynes): messenger should be hardcoded in the contract
		messenger, ok := deployment.Args[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for messenger")
		}
		otherBridge, ok := deployment.Args[1].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for otherBridge")
		}
		_, tx, _, err = bindings.DeployL2ERC721Bridge(opts, backend, messenger, otherBridge)
	case "KromaMintableERC721Factory":
		bridge, ok := deployment.Args[0].(common.Address)
		if !ok {
			return nil, fmt.Errorf("invalid type for bridge")
		}
		remoteChainId, ok := deployment.Args[1].(*big.Int)
		if !ok {
			return nil, fmt.Errorf("invalid type for remoteChainId")
		}
		_, tx, _, err = bindings.DeployKromaMintableERC721Factory(opts, backend, bridge, remoteChainId)
	default:
		return tx, fmt.Errorf("unknown contract: %s", deployment.Name)
	}

	return tx, err
}
