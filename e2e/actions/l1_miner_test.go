package actions

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"

	"github.com/kroma-network/kroma/components/node/testlog"
	"github.com/kroma-network/kroma/e2e/e2eutils"
)

func TestL1Miner_BuildBlock(gt *testing.T) {
	t := NewDefaultTesting(gt)
	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	sd := e2eutils.Setup(t, dp, defaultAlloc)
	log := testlog.Logger(t, log.LvlDebug)
	miner := NewL1Miner(t, log, sd.L1Cfg)
	t.Cleanup(func() {
		_ = miner.Close()
	})

	cl := miner.EthClient()
	l1Cl := miner.L1Client(t, sd.RollupCfg)
	signer := types.LatestSigner(sd.L1Cfg.Config)

	// send a tx to the miner
	tx := types.MustSignNewTx(dp.Secrets.Alice, signer, &types.DynamicFeeTx{
		ChainID:   sd.L1Cfg.Config.ChainID,
		Nonce:     0,
		GasTipCap: big.NewInt(2 * params.GWei),
		GasFeeCap: new(big.Int).Add(miner.l1Chain.CurrentBlock().BaseFee, big.NewInt(2*params.GWei)),
		Gas:       params.TxGas,
		To:        &dp.Addresses.Bob,
		Value:     e2eutils.Ether(2),
	})
	require.NoError(gt, cl.SendTransaction(t.Ctx(), tx))

	// make an empty block, even though a tx may be waiting
	miner.ActL1StartBlock(10)(t)
	miner.ActL1EndBlock(t)
	bl := miner.l1Chain.CurrentBlock()
	_, txs, err := l1Cl.InfoAndTxsByHash(t.Ctx(), bl.Hash())
	require.NoError(t, err)
	require.Equal(t, uint64(1), bl.Number.Uint64())
	require.Zero(gt, txs.Len())

	// now include the tx when we want it to
	miner.ActL1StartBlock(10)(t)
	miner.ActL1IncludeTx(dp.Addresses.Alice)(t)
	miner.ActL1EndBlock(t)
	bl = miner.l1Chain.CurrentBlock()
	_, txs, err = l1Cl.InfoAndTxsByHash(t.Ctx(), bl.Hash())
	require.NoError(t, err)
	require.Equal(t, uint64(2), bl.Number.Uint64())
	require.Equal(t, 1, txs.Len())
	require.Equal(t, tx.Hash(), txs[0].Hash())

	// now make a replica that syncs these two blocks from the miner
	replica := NewL1Replica(t, log, sd.L1Cfg)
	t.Cleanup(func() {
		_ = replica.Close()
	})
	replica.ActL1Sync(miner.CanonL1Chain())(t)
	replica.ActL1Sync(miner.CanonL1Chain())(t)
	require.Equal(t, replica.l1Chain.CurrentBlock().Hash(), miner.l1Chain.CurrentBlock().Hash())
}
