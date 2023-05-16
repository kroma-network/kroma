import assert from 'assert'

import '@kroma-network/hardhat-deploy-config'
import { ethers } from 'ethers'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import { assertContractVariable, deploy } from '../src/deploy-utils'

const uint128Max = ethers.BigNumber.from('0xffffffffffffffffffffffffffffffff')

const deployFn: DeployFunction = async (hre) => {
  const batcherHash = hre.ethers.utils
    .hexZeroPad(hre.deployConfig.batchSenderAddress, 32)
    .toLowerCase()

  await deploy(hre, 'SystemConfig', {
    args: [
      hre.deployConfig.finalSystemOwner,
      hre.deployConfig.gasPriceOracleOverhead,
      hre.deployConfig.gasPriceOracleScalar,
      batcherHash,
      hre.deployConfig.l2GenesisBlockGasLimit,
      hre.deployConfig.p2pProposerAddress,
      {
        maxResourceLimit: 20_000_000,
        elasticityMultiplier: 10,
        baseFeeMaxChangeDenominator: 8,
        systemTxMaxGas: 1_000_000,
        minimumBaseFee: ethers.utils.parseUnits('1', 'gwei'),
        maximumBaseFee: uint128Max,
      },
    ],
    isProxyImpl: true,
    initArgs: [
      hre.deployConfig.finalSystemOwner,
      hre.deployConfig.gasPriceOracleOverhead,
      hre.deployConfig.gasPriceOracleScalar,
      batcherHash,
      hre.deployConfig.l2GenesisBlockGasLimit,
      hre.deployConfig.p2pProposerAddress,
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

      const config = await contract.resourceConfig()
      assert(config.maxResourceLimit === 20_000_000)
      assert(config.elasticityMultiplier === 10)
      assert(config.baseFeeMaxChangeDenominator === 8)
      assert(config.systemTxMaxGas === 1_000_000)
      assert(ethers.utils.parseUnits('1', 'gwei').eq(config.minimumBaseFee))
      assert(config.maximumBaseFee.eq(uint128Max))
    },
  })
}

deployFn.tags = ['SystemConfig', 'setup', 'l1']

export default deployFn
