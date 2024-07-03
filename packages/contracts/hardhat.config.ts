import dotenv from 'dotenv'
import { ethers } from 'ethers'
import { HardhatUserConfig } from 'hardhat/config'

// Hardhat plugins
import '@foundry-rs/hardhat-forge'
import '@kroma/hardhat-deploy-config'
import '@nomiclabs/hardhat-ethers'
import 'hardhat-deploy'

// Hardhat tasks
import './tasks'

// Deploy configuration
import { deployConfigSpec } from './src/deploy-config'

// Load environment variables
dotenv.config()

// Private key of deployer(0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266)
const PRIVATE_KEY_DEPLOYER_DEVNET =
  'ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80'

const config: HardhatUserConfig = {
  networks: {
    hardhat: {
      live: false,
    },
    mainnet: {
      url: process.env.L1_RPC_MAINNET || '',
      accounts: [
        process.env.PRIVATE_KEY_DEPLOYER_MAINNET || ethers.constants.HashZero,
      ],
      companionNetworks: {
        l2: 'kroma',
      },
      deploy: ['./deploy/L1'],
    },
    kroma: {
      url: process.env.L2_RPC_KROMA_MAINNET || '',
      accounts: [
        process.env.PRIVATE_KEY_DEPLOYER_MAINNET || ethers.constants.HashZero,
      ],
      companionNetworks: {
        l1: 'mainnet',
      },
      deploy: ['./deploy/L2'],
    },
    holesky: {
      chainId: 17000,
      url: process.env.L1_RPC_HOLESKY || '',
      accounts: [
        process.env.PRIVATE_KEY_DEPLOYER_HOLESKY || ethers.constants.HashZero,
      ],
      companionNetworks: {
        l2: 'kromaHolesky',
      },
      deploy: ['./deploy/L1'],
    },
    kromaHolesky: {
      url: process.env.L2_RPC_KROMA_HOLESKY || '',
      accounts: [
        process.env.PRIVATE_KEY_DEPLOYER_HOLESKY || ethers.constants.HashZero,
      ],
      companionNetworks: {
        l1: 'holesky',
      },
      deploy: ['./deploy/L2'],
    },
    sepolia: {
      chainId: 11155111,
      url: process.env.L1_RPC_SEPOLIA || '',
      accounts: [
        process.env.PRIVATE_KEY_DEPLOYER_SEPOLIA || ethers.constants.HashZero,
      ],
      companionNetworks: {
        l2: 'kromaSepolia',
      },
      deploy: ['./deploy/L1'],
    },
    kromaSepolia: {
      chainId: 2358,
      url: process.env.L2_RPC_KROMA_SEPOLIA || '',
      accounts: [
        process.env.PRIVATE_KEY_DEPLOYER_SEPOLIA || ethers.constants.HashZero,
      ],
      companionNetworks: {
        l1: 'sepolia',
      },
      deploy: ['./deploy/L2'],
    },
    devnetL1: {
      live: false,
      url: 'http://localhost:8545',
      accounts: [PRIVATE_KEY_DEPLOYER_DEVNET],
      saveDeployments: true,
      deploy: ['./deploy/L1'],
    },
    devnetL2: {
      live: false,
      url: process.env.RPC_URL || 'http://localhost:9545',
      accounts: [PRIVATE_KEY_DEPLOYER_DEVNET],
      saveDeployments: true,
      companionNetworks: {
        l1: 'devnetL1',
      },
      deploy: ['./deploy/L2'],
    },
  },
  foundry: {
    buildInfo: true,
  },
  paths: {
    deploy: './deploy',
    deployments: './deployments',
    deployConfig: './deploy-config',
  },
  namedAccounts: {
    deployer: {
      default: 0,
    },
  },
  deployConfigSpec,
  external: {
    contracts: [
      {
        artifacts: '../contracts/artifacts',
      },
    ],
    deployments: {
      mainnet: ['../contracts/deployments/mainnet'],
      sepolia: ['../contracts/deployments/sepolia'],
      holesky: ['../contracts/deployments/holesky'],
    },
  },
  solidity: {
    compilers: [
      {
        version: '0.8.15',
        settings: {
          optimizer: {
            enabled: true,
            runs: 10_000,
          },
        },
      },
      {
        version: '0.5.17', // Required for WETH9
        settings: {
          optimizer: {
            enabled: true,
            runs: 10_000,
          },
        },
      },
    ],
    settings: {
      metadata: {
        bytecodeHash:
          process.env.FOUNDRY_PROFILE === 'echidna' ? 'ipfs' : 'none',
      },
      outputSelection: {
        '*': {
          '*': ['metadata', 'storageLayout'],
        },
      },
    },
  },
}

export default config
