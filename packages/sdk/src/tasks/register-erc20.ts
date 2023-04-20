import { getContractDefinition, predeploys } from '@kroma-network/contracts'
import '@nomiclabs/hardhat-ethers'
import { task, types } from 'hardhat/config'
import 'hardhat-deploy'

task('register-erc20', 'Register ERC20 onto L2.')
  .addParam('l1Url', 'L1 provider URL.', 'http://localhost:8545', types.string)
  .addParam(
    'l1Token',
    'Address of token to register onto L2.',
    '',
    types.string
  )
  .setAction(async (args, hre) => {
    const l1Provider = new hre.ethers.providers.StaticJsonRpcProvider(
      args.l1Url
    )

    const Artifact__KromaMintableERC20 = await getContractDefinition(
      'KromaMintableERC20'
    )
    const Artifact__KromaMintableERC20Factory = await getContractDefinition(
      'KromaMintableERC20Factory'
    )

    // get token info
    let l1Token = await hre.ethers.getContractAt(
      Artifact__KromaMintableERC20.abi,
      args.l1Token
    )
    l1Token = l1Token.connect(l1Provider)

    const tokenName = await l1Token.name()
    const tokenSymbol = await l1Token.symbol()
    console.log(`Target token: ${tokenName} (${tokenSymbol})`)

    const factory = await hre.ethers.getContractAt(
      Artifact__KromaMintableERC20Factory.abi,
      predeploys.KromaMintableERC20Factory
    )
    const tx = await factory.createKromaMintableERC20(
      l1Token.address,
      tokenName,
      tokenSymbol
    )
    console.log('Transaction sent: ' + tx.hash)
    const receipt = await tx.wait()
    const createdEvent = receipt.events.find(
      (x) => x.event === 'KromaMintableERC20Created'
    )
    if (createdEvent) {
      console.log(
        `${tokenSymbol} has been registered: ${createdEvent.args.localToken}`
      )
    } else {
      throw new Error(`failed to register ${tokenName}`)
    }
  })
