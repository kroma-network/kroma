import assert from 'assert'

import '@kroma-network/hardhat-deploy-config'
import { BigNumber } from 'ethers'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import { defaultResourceConfig } from '../src'
import { assertContractVariable, deploy } from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const batcherHash = hre.ethers.utils
    .hexZeroPad(hre.deployConfig.batchSenderAddress, 32)
    .toLowerCase()

  const l2GenesisBlockGasLimit = BigNumber.from(
    hre.deployConfig.l2GenesisBlockGasLimit
  )
  const l2GasLimitLowerBound = BigNumber.from(
    defaultResourceConfig.systemTxMaxGas +
      defaultResourceConfig.maxResourceLimit
  )
  if (l2GenesisBlockGasLimit.lt(l2GasLimitLowerBound)) {
    throw new Error(
      `L2 genesis block gas limit must be at least ${l2GasLimitLowerBound}`
    )
  }

  await deploy(hre, 'SystemConfig', {
    args: [
      hre.deployConfig.finalSystemOwner,
      hre.deployConfig.gasPriceOracleOverhead,
      hre.deployConfig.gasPriceOracleScalar,
      batcherHash,
      hre.deployConfig.l2GenesisBlockGasLimit,
      hre.deployConfig.p2pProposerAddress,
      defaultResourceConfig,
      hre.deployConfig.validatorRewardScalar,
    ],
    isProxyImpl: true,
    initArgs: [
      hre.deployConfig.finalSystemOwner,
      hre.deployConfig.gasPriceOracleOverhead,
      hre.deployConfig.gasPriceOracleScalar,
      batcherHash,
      hre.deployConfig.l2GenesisBlockGasLimit,
      hre.deployConfig.p2pProposerAddress,
      defaultResourceConfig,
      hre.deployConfig.validatorRewardScalar,
    ],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'owner',
        hre.deployConfig.finalSystemOwner
      )
      await assertContractVariable(
        contract,
        'overhead',
        hre.deployConfig.gasPriceOracleOverhead
      )
      await assertContractVariable(
        contract,
        'scalar',
        hre.deployConfig.gasPriceOracleScalar
      )
      await assertContractVariable(contract, 'batcherHash', batcherHash)
      await assertContractVariable(
        contract,
        'unsafeBlockSigner',
        hre.deployConfig.p2pProposerAddress
      )
      await assertContractVariable(
        contract,
        'validatorRewardScalar',
        hre.deployConfig.validatorRewardScalar
      )

      const config = await contract.resourceConfig()
      assert(config.maxResourceLimit === defaultResourceConfig.maxResourceLimit)
      assert(
        config.elasticityMultiplier ===
          defaultResourceConfig.elasticityMultiplier
      )
      assert(
        config.baseFeeMaxChangeDenominator ===
          defaultResourceConfig.baseFeeMaxChangeDenominator
      )
      assert(
        BigNumber.from(config.systemTxMaxGas).eq(
          defaultResourceConfig.systemTxMaxGas
        )
      )
      assert(
        BigNumber.from(config.minimumBaseFee).eq(
          defaultResourceConfig.minimumBaseFee
        )
      )
      assert(
        BigNumber.from(config.maximumBaseFee).eq(
          defaultResourceConfig.maximumBaseFee
        )
      )
    },
  })
}

deployFn.tags = ['SystemConfig', 'setup', 'l1']

export default deployFn
