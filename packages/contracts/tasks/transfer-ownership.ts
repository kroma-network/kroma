import assert from 'assert'
import readline from 'readline'

import '@nomiclabs/hardhat-ethers'
import { task } from 'hardhat/config'
import { HardhatRuntimeEnvironment } from 'hardhat/types'

import { getDeploymentAddress } from '../src/deploy-utils'

const isL1Network = (hre: HardhatRuntimeEnvironment): boolean => {
  const networkName = hre.network.name
  const l1Network = ['mainnet', 'sepolia', 'easel', 'devnetL1']
  return l1Network.includes(networkName)
}

const transferProxyAdminOwnership = async (
  hre: HardhatRuntimeEnvironment,
  newOwner: string
) => {
  console.log('transfer ProxyAdmin ownership to TimeLock')
  const proxyAdminAddress = isL1Network(hre)
    ? await getDeploymentAddress(hre, 'ProxyAdmin')
    : await getDeploymentAddress(hre, 'ProxyAdminProxy')
  let proxyAdmin = await hre.ethers.getContractAt(
    'ProxyAdmin',
    proxyAdminAddress
  )

  const currentProxyAdminOwner = await proxyAdmin.owner()
  if (currentProxyAdminOwner === newOwner) {
    console.log(
      'skip the ProxyAdmin owner transfer process because the owner has already been transferred.'
    )
    return
  }

  proxyAdmin = proxyAdmin.connect(
    hre.ethers.provider.getSigner(currentProxyAdminOwner)
  )

  const tx = await proxyAdmin.transferOwnership(newOwner, {
    from: currentProxyAdminOwner,
  })
  await tx.wait()

  assert((await proxyAdmin.owner()) === newOwner)
  console.log('successfully transferred ownership of ProxyAdmin')
}

const transferSecurityCouncilTokenOwnership = async (
  hre: HardhatRuntimeEnvironment,
  newOwner: string
) => {
  console.log(
    'transfer SecurityCouncilToken ownership to SecurityCouncilTokenOwner'
  )
  const scTokenProxyAddress = await getDeploymentAddress(
    hre,
    'SecurityCouncilTokenProxy'
  )
  let scToken = await hre.ethers.getContractAt(
    'SecurityCouncilToken',
    scTokenProxyAddress
  )
  const currentScTokenOwner = await scToken.owner()
  if (
    hre.ethers.utils.getAddress(currentScTokenOwner) ===
    hre.ethers.utils.getAddress(newOwner)
  ) {
    console.log(
      'skip the SecurityCouncilToken owner transfer process because the owner has already been transferred.'
    )
    return
  }

  scToken = scToken.connect(hre.ethers.provider.getSigner(currentScTokenOwner))

  const tx = await scToken.transferOwnership(newOwner, {
    from: currentScTokenOwner,
  })
  await tx.wait()

  assert(
    hre.ethers.utils.getAddress(await scToken.owner()) ===
      hre.ethers.utils.getAddress(newOwner)
  )
  console.log('successfully transferred ownership of SecurityCouncilToken')
}

task('transfer-ownership', 'Transfer ownership').setAction(
  async (args, hre) => {
    const readLineAsync = () => {
      const rl = readline.createInterface({
        input: process.stdin,
      })

      return new Promise((resolve) => {
        rl.prompt()
        rl.on('line', (line) => {
          rl.close()
          resolve(line)
        })
      })
    }

    const run = async () => {
      const networkName = hre.network.name
      const yes = 'yes'
      console.warn(
        '*******************************************************************************************************'
      )
      console.warn(
        '  [WARNING] Do you want to continue with the transferOwnership operation for the ' +
          networkName +
          ' network?'
      )
      console.warn("  Type and enter 'yes' to continue")
      console.warn(
        '*******************************************************************************************************'
      )
      const line = await readLineAsync()
      if (line.toString() !== yes) {
        console.log('The response is invalid. Terminate the task.')
        return
      }

      const timeLockProxyAddress = await getDeploymentAddress(
        hre,
        'TimeLockProxy'
      )
      await transferProxyAdminOwnership(hre, timeLockProxyAddress)
      await transferSecurityCouncilTokenOwnership(hre, timeLockProxyAddress)
    }

    await run()
  }
)
