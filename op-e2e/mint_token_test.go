package op_e2e

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/require"

	"github.com/kroma-network/kroma/kroma-bindings/bindings"
	"github.com/kroma-network/kroma/kroma-bindings/predeploys"
	"github.com/kroma-network/kroma/op-e2e/e2eutils/geth"
	"github.com/kroma-network/kroma/op-e2e/e2eutils/wait"
)

var (
	floorUnit = big.NewInt(1e18 / 1e8)
	zeroTime  = hexutil.Uint64(0)
)

func TestMintToken(t *testing.T) {
	InitParallel(t)

	t.Run("MintAndDistribute", func(t *testing.T) {
		cfg := DefaultSystemConfig(t)

		cfg.DeployConfig.L2GenesisBurgundyTimeOffset = &zeroTime
		cfg.DeployConfig.MintManagerSlidingWindowBlocks = 5

		sys, err := cfg.Start(t)
		require.Nil(t, err, "Error starting up system")
		defer sys.Close()

		l2Client := sys.Clients["sequencer"]

		governanceToken, err := bindings.NewGovernanceTokenCaller(predeploys.GovernanceTokenAddr, l2Client)
		require.NoError(t, err)

		// Ensure tokens are minted and distributed correctly until the epoch 5 ended
		for i := uint64(1); i <= 5; i++ {
			blockNumber := new(big.Int).SetUint64(cfg.DeployConfig.MintManagerSlidingWindowBlocks * i)
			block, err := geth.WaitForBlock(blockNumber, l2Client, 10*time.Second)
			require.NoError(t, err)
			checkBalance(t, cfg, governanceToken, block.Number())
		}
	})

	t.Run("InitMintPerBlock_100", func(t *testing.T) {
		cfg := DefaultSystemConfig(t)

		initMintPerBlock := hexutil.Big(*new(big.Int).Mul(big.NewInt(100), big.NewInt(1e18)))
		cfg.DeployConfig.L2GenesisBurgundyTimeOffset = &zeroTime
		cfg.DeployConfig.MintManagerInitMintPerBlock = &initMintPerBlock
		cfg.DeployConfig.MintManagerSlidingWindowBlocks = 5

		sys, err := cfg.Start(t)
		require.Nil(t, err, "Error starting up system")
		defer sys.Close()

		l2Client := sys.Clients["sequencer"]

		governanceToken, err := bindings.NewGovernanceTokenCaller(predeploys.GovernanceTokenAddr, l2Client)
		require.NoError(t, err)

		prevSupply := big.NewInt(0)

		for i := uint64(1); i <= cfg.DeployConfig.MintManagerSlidingWindowBlocks; i++ {
			blockNumber := new(big.Int).SetUint64(i)
			_, err := geth.WaitForBlock(blockNumber, l2Client, 2*time.Second)
			currentSupply, err := governanceToken.TotalSupply(&bind.CallOpts{BlockNumber: blockNumber})
			require.NoError(t, err)
			mintAmount := new(big.Int).Sub(currentSupply, prevSupply)
			require.Equal(t, mintAmount, cfg.DeployConfig.MintManagerInitMintPerBlock.ToInt())
			prevSupply = currentSupply
		}
	})

	t.Run("Decayed", func(t *testing.T) {
		cfg := DefaultSystemConfig(t)

		cfg.DeployConfig.L2GenesisBurgundyTimeOffset = &zeroTime
		cfg.DeployConfig.MintManagerSlidingWindowBlocks = 1

		sys, err := cfg.Start(t)
		require.Nil(t, err, "Error starting up system")
		defer sys.Close()

		l2Client := sys.Clients["sequencer"]

		governanceToken, err := bindings.NewGovernanceToken(predeploys.GovernanceTokenAddr, l2Client)
		require.NoError(t, err)
		prevMint := new(big.Int).Set(cfg.DeployConfig.MintManagerInitMintPerBlock.ToInt())
		prevSupply, err := governanceToken.TotalSupply(&bind.CallOpts{})
		require.NoError(t, err)

		// Ensure that the number of tokens minted in each block decreases for about 10 blocks.
		for i := int64(1); i <= 10; i++ {
			_, err := geth.WaitForBlock(big.NewInt(i), l2Client, 2*time.Second)
			require.NoError(t, err)
			currentSupply, err := governanceToken.TotalSupply(&bind.CallOpts{})
			require.NoError(t, err)
			require.Equal(t, currentSupply.Cmp(prevSupply), 1)
			currentMint := new(big.Int).Sub(currentSupply, prevSupply)
			require.Equal(t, currentMint.Cmp(prevMint), -1)

			prevSupply = currentSupply
			prevMint = currentMint
		}
	})

	t.Run("Exhausted", func(t *testing.T) {
		cfg := DefaultSystemConfig(t)

		cfg.DeployConfig.L2GenesisBurgundyTimeOffset = &zeroTime
		cfg.DeployConfig.MintManagerSlidingWindowBlocks = 1

		sys, err := cfg.Start(t)
		require.Nil(t, err, "Error starting up system")
		defer sys.Close()

		l2Client := sys.Clients["sequencer"]

		governanceToken, err := bindings.NewGovernanceToken(predeploys.GovernanceTokenAddr, l2Client)
		require.NoError(t, err)
		mintManager, err := bindings.NewMintManagerCaller(predeploys.MintManagerAddr, l2Client)
		require.NoError(t, err)

		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		epochBlocks := new(big.Int).SetUint64(cfg.DeployConfig.MintManagerSlidingWindowBlocks)
		prevSupply := big.NewInt(0)
		blockNum := big.NewInt(1)
		for {
			select {
			case <-ctx.Done():
				require.Fail(t, "token not exhausted", ctx.Err())
				break
			default:
				err := wait.ForNextBlock(ctx, l2Client)
				require.NoError(t, err)
				currentSupply, err := governanceToken.TotalSupply(&bind.CallOpts{})
				require.NoError(t, err)

				if currentSupply.Cmp(prevSupply) == 0 {
					mintAmount, err := mintManager.MintAmountPerBlock(&bind.CallOpts{}, blockNum)
					require.NoError(t, err)
					if mintAmount.Uint64() == 0 {
						return
					}
				}

				prevSupply = currentSupply
				blockNum.Add(blockNum, epochBlocks)
			}
		}
	})
}

func TestInitialMint(t *testing.T) {
	cfg := DefaultSystemConfig(t)

	burgundyBlock := uint64(5)
	burgundyTimeOffset := hexutil.Uint64(burgundyBlock * cfg.DeployConfig.L2BlockTime)
	cfg.DeployConfig.L2GenesisBurgundyTimeOffset = &burgundyTimeOffset
	cfg.DeployConfig.MintManagerSlidingWindowBlocks = 1

	sys, err := cfg.Start(t)
	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	l2Client := sys.Clients["sequencer"]

	governanceToken, err := bindings.NewGovernanceTokenCaller(predeploys.GovernanceTokenAddr, l2Client)
	require.NoError(t, err)

	targetBlockNum := new(big.Int).SetUint64(1)
	block, err := geth.WaitForBlock(targetBlockNum, l2Client, 20*time.Second)
	require.NoError(t, err)
	supply, err := governanceToken.TotalSupply(&bind.CallOpts{BlockNumber: targetBlockNum})
	require.NoError(t, err)
	require.Equal(t, supply.Uint64(), uint64(0))

	targetBlockNum = new(big.Int).SetUint64(burgundyBlock)
	block, err = geth.WaitForBlock(targetBlockNum, l2Client, 20*time.Second)
	require.NoError(t, err)
	checkBalance(t, cfg, governanceToken, block.Number())
}

func checkBalance(t *testing.T, cfg SystemConfig, token *bindings.GovernanceTokenCaller, blockNumber *big.Int) {
	callOpts := &bind.CallOpts{BlockNumber: blockNumber}

	expected := estimateTotalSupply(cfg, blockNumber)
	totalSupply, err := token.TotalSupply(callOpts)
	require.NoError(t, err)
	require.Equal(t, expected, totalSupply)

	for i, acc := range cfg.DeployConfig.MintManagerRecipients {
		expected := new(big.Int).SetUint64(cfg.DeployConfig.MintManagerShares[i])
		expected.Mul(expected, totalSupply)
		expected.Div(expected, big.NewInt(1e5))
		balance, err := token.BalanceOf(callOpts, acc)
		require.NoError(t, err)
		require.Equal(t, expected, balance)
	}
}

func epochAndOffset(cfg SystemConfig, blockNumber *big.Int) (uint64, uint64) {
	epoch := blockNumber.Uint64()/cfg.DeployConfig.MintManagerSlidingWindowBlocks + 1
	offset := blockNumber.Uint64() % cfg.DeployConfig.MintManagerSlidingWindowBlocks
	if offset == 0 {
		epoch = epoch - 1
		offset = cfg.DeployConfig.MintManagerSlidingWindowBlocks
	}

	return epoch, offset
}

func mintAmountPerBlock(cfg SystemConfig, epoch uint64) *big.Int {
	decayingFactor := new(big.Int).SetUint64(cfg.DeployConfig.MintManagerDecayingFactor)
	decayingDenom := new(big.Int).SetUint64(1e5)

	amount := new(big.Int).Set(cfg.DeployConfig.MintManagerInitMintPerBlock.ToInt())
	for i := uint64(1); i < epoch; i++ {
		amount.Mul(amount, decayingFactor).Div(amount, decayingDenom)
		amount.Div(amount, floorUnit).Mul(amount, floorUnit)
	}

	return amount
}

func estimateTotalSupply(cfg SystemConfig, blockNumber *big.Int) *big.Int {
	lastEpoch, offset := epochAndOffset(cfg, blockNumber)

	totalSupply := big.NewInt(0)
	for epoch := uint64(1); epoch < lastEpoch; epoch++ {
		mintPerBlock := mintAmountPerBlock(cfg, epoch)
		blocks := new(big.Int).SetUint64(cfg.DeployConfig.MintManagerSlidingWindowBlocks)
		amount := new(big.Int).Mul(mintPerBlock, blocks)
		totalSupply.Add(totalSupply, amount)
	}

	if offset > 0 {
		mintPerBlock := mintAmountPerBlock(cfg, lastEpoch)
		blocks := new(big.Int).SetUint64(offset)
		amount := new(big.Int).Mul(mintPerBlock, blocks)
		totalSupply.Add(totalSupply, amount)
	}

	return totalSupply
}
