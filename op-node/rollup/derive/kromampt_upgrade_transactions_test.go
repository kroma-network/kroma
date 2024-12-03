package derive

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ethereum-optimism/optimism/op-bindings/predeploys"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func TestMPTSourcesMatchSpec(t *testing.T) {
	for _, test := range []struct {
		source       UpgradeDepositSource
		expectedHash string
	}{
		{
			source:       deployBaseFeeVaultSource,
			expectedHash: "0xdd882729c9833d55eb4ad39f08f578c0b92e171234e1b4ddc8f91d7878cd8686",
		},
		{
			source:       updateBaseFeeVaultProxySource,
			expectedHash: "0x747e2dc0ea5ed790a8746601b59c7588bfdce93e5d4dc0f6fd92a4d32c07596d",
		},
		{
			source:       deployL1FeeVaultSource,
			expectedHash: "0xb88c42fa962a1aa3d495761f85b93fd496a9607b57ffda85a1e3840b0d854794",
		},
		{
			source:       updateL1FeeVaultProxySource,
			expectedHash: "0xd4b39da88babd8b9c3d96bfc88b2c4be2a977c3b2629d0f2498fc290e1c523be",
		},
		{
			source:       deploySequencerFeeVaultSource,
			expectedHash: "0x4618df4b1692fcaee1da6b0532641b361f5dbfc3bb5714ae22802f989dc59c6f",
		},
		{
			source:       updateSequencerFeeVaultProxySource,
			expectedHash: "0xaf7e0d5c4b92ad5918fe238414031d8a82b87841ed70f1ba247d50121f9c04eb",
		},
		{
			source:       deployL1BlockMPTSource,
			expectedHash: "0x3c459db0fc6553f1b48a3b7f1e20de2e80d7f28976085817070699d1502dbe74",
		},
		{
			source:       updateL1BlockMPTProxySource,
			expectedHash: "0xb3a45c33c31a2322c4c4db99788dcae68606b98e4c90ea9fb3d1bc93ddfbf3bf",
		},
	} {
		require.Equal(t, common.HexToHash(test.expectedHash), test.source.SourceHash())
	}
}

func TestMPTNetworkTransactions(t *testing.T) {
	upgradeTxns, err := KromaMPTNetworkUpgradeTransactions()
	require.NoError(t, err)
	require.Len(t, upgradeTxns, 8)

	deployL1BlockSender, deployL1Block := toDepositTxn(t, upgradeTxns[0])
	require.Equal(t, deployL1BlockSender, L1BlockDeployerAddress)
	require.Equal(t, deployL1BlockMPTSource.SourceHash(), deployL1Block.SourceHash())
	require.Nil(t, deployL1Block.To())
	require.Equal(t, uint64(375_000), deployL1Block.Gas())
	require.Equal(t, hexutil.Bytes(l1BlockMPTDeploymentBytecode).String(), hexutil.Bytes(deployL1Block.Data()).String())

	deployBaseFeeVaultSender, deployBaseFeeVault := toDepositTxn(t, upgradeTxns[1])
	require.Equal(t, deployBaseFeeVaultSender, KromaMPTDeployerAddress)
	require.Equal(t, deployBaseFeeVaultSource.SourceHash(), deployBaseFeeVault.SourceHash())
	require.Nil(t, deployBaseFeeVault.To())
	require.Equal(t, uint64(1_000_000), deployBaseFeeVault.Gas())
	require.Equal(t, hexutil.Bytes(feeVaultDeploymentBytecode).String(), hexutil.Bytes(deployBaseFeeVault.Data()).String())

	deployL1FeeVaultSender, deployL1FeeVault := toDepositTxn(t, upgradeTxns[2])
	require.Equal(t, deployL1FeeVaultSender, KromaMPTDeployerAddress)
	require.Equal(t, deployL1FeeVaultSource.SourceHash(), deployL1FeeVault.SourceHash())
	require.Nil(t, deployL1FeeVault.To())
	require.Equal(t, uint64(1_000_000), deployL1FeeVault.Gas())
	require.Equal(t, hexutil.Bytes(feeVaultDeploymentBytecode).String(), hexutil.Bytes(deployL1FeeVault.Data()).String())

	deployL1FeeVaultSender, deploySequencerFeeVault := toDepositTxn(t, upgradeTxns[3])
	require.Equal(t, deployL1FeeVaultSender, KromaMPTDeployerAddress)
	require.Equal(t, deploySequencerFeeVaultSource.SourceHash(), deploySequencerFeeVault.SourceHash())
	require.Nil(t, deploySequencerFeeVault.To())
	require.Equal(t, uint64(1_000_000), deploySequencerFeeVault.Gas())
	require.Equal(t, hexutil.Bytes(feeVaultDeploymentBytecode).String(), hexutil.Bytes(deploySequencerFeeVault.Data()).String())

	updateL1BlockProxySender, updateL1BlockProxy := toDepositTxn(t, upgradeTxns[4])
	require.Equal(t, updateL1BlockProxySender, common.Address{})
	require.Equal(t, updateL1BlockMPTProxySource.SourceHash(), updateL1BlockProxy.SourceHash())
	require.NotNil(t, updateL1BlockProxy.To())
	require.Equal(t, *updateL1BlockProxy.To(), predeploys.L1BlockAddr)
	require.Equal(t, uint64(50_000), updateL1BlockProxy.Gas())
	require.Equal(t, common.FromHex("0x3659cfe6000000000000000000000000cd7467a8926d13f8b41ea035ff3761326c822bad"), updateL1BlockProxy.Data())

	updateBaseFeeVaultProxySender, updateBaseFeeVaultProxy := toDepositTxn(t, upgradeTxns[5])
	require.Equal(t, updateBaseFeeVaultProxySender, common.Address{})
	require.Equal(t, updateBaseFeeVaultProxySource.SourceHash(), updateBaseFeeVaultProxy.SourceHash())
	require.NotNil(t, updateBaseFeeVaultProxy.To())
	require.Equal(t, *updateBaseFeeVaultProxy.To(), predeploys.BaseFeeVaultAddr)
	require.Equal(t, uint64(50_000), updateBaseFeeVaultProxy.Gas())
	require.Equal(t, common.FromHex("0x3659cfe60000000000000000000000008ab182eaa3d481ee66d13727c82c37a220032372"), updateBaseFeeVaultProxy.Data())

	updateL1FeeVaultProxySender, updateL1FeeVaultProxy := toDepositTxn(t, upgradeTxns[6])
	require.Equal(t, updateL1FeeVaultProxySender, common.Address{})
	require.Equal(t, updateL1FeeVaultProxySource.SourceHash(), updateL1FeeVaultProxy.SourceHash())
	require.NotNil(t, updateL1FeeVaultProxy.To())
	require.Equal(t, *updateL1FeeVaultProxy.To(), predeploys.L1FeeVaultAddr)
	require.Equal(t, uint64(50_000), updateL1FeeVaultProxy.Gas())
	require.Equal(t, common.FromHex("0x3659cfe60000000000000000000000008974d3621be88d87f4ff264e1215428fd60134bb"), updateL1FeeVaultProxy.Data())

	updateSequencerFeeVaultProxySender, updateSequencerFeeVaultProxy := toDepositTxn(t, upgradeTxns[7])
	require.Equal(t, updateSequencerFeeVaultProxySender, common.Address{})
	require.Equal(t, updateSequencerFeeVaultProxySource.SourceHash(), updateSequencerFeeVaultProxy.SourceHash())
	require.NotNil(t, updateSequencerFeeVaultProxy.To())
	require.Equal(t, *updateSequencerFeeVaultProxy.To(), predeploys.SequencerFeeVaultAddr)
	require.Equal(t, uint64(50_000), updateSequencerFeeVaultProxy.Gas())
	require.Equal(t, common.FromHex("0x3659cfe6000000000000000000000000ab31a20195e56a4858628b7a1c80178fa281d010"), updateSequencerFeeVaultProxy.Data())
}
