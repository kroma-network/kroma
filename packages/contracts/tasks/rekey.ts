import * as bip39 from 'bip39'
import { hdkey } from 'ethereumjs-wallet'
import { task } from 'hardhat/config'

task('rekey', 'Generates a new set of keys for a test network').setAction(
  async () => {
    const mnemonic = bip39.generateMnemonic()
    const pathPrefix = "m/44'/60'/0'/0"
    const labels = [
      'deployer',
      'validator',
      'proxyAdminOwner',
      'kromaBaseFeeRecipient',
      'kromaL1FeeRecipient',
      'p2pProposerAddress',
      'batchSenderAddress',
    ]

    const hdwallet = hdkey.fromMasterSeed(await bip39.mnemonicToSeed(mnemonic))

    console.log(`Mnemonic: ${mnemonic}`)
    for (let i = 0; i < labels.length; i++) {
      const label = labels[i]
      const wallet = hdwallet.derivePath(`${pathPrefix}/${i}`).getWallet()
      const addr = '0x' + wallet.getAddress().toString('hex')
      const pk = wallet.getPrivateKey().toString('hex')

      console.log()
      console.log(`${label}: ${addr}`)
      console.log(`Private Key: ${pk}`)
    }
  }
)
