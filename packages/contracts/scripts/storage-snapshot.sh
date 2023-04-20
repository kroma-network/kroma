#!/usr/bin/env bash

set -e

if ! command -v forge &> /dev/null
then
    echo "forge could not be found. Please install forge by running:"
    echo "curl -L https://foundry.paradigm.xyz | bash"
    exit
fi

contracts=(
  contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger
  contracts/L1/L1StandardBridge.sol:L1StandardBridge
  contracts/L1/L2OutputOracle.sol:L2OutputOracle
  contracts/L1/KromaPortal.sol:KromaPortal
  contracts/L1/SystemConfig.sol:SystemConfig
  contracts/L2/L1Block.sol:L1Block
  contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger
  contracts/L2/L2StandardBridge.sol:L2StandardBridge
  contracts/L2/L2ToL1MessagePasser.sol:L2ToL1MessagePasser
  contracts/L2/ProposerFeeVault.sol:ProposerFeeVault
  contracts/L2/BaseFeeVault.sol:BaseFeeVault
  contracts/L2/L1FeeVault.sol:L1FeeVault
  contracts/vendor/WETH9.sol:WETH9
  contracts/universal/ProxyAdmin.sol:ProxyAdmin
  contracts/universal/Proxy.sol:Proxy
  contracts/universal/KromaMintableERC20.sol:KromaMintableERC20
  contracts/universal/KromaMintableERC20Factory.sol:KromaMintableERC20Factory
)

dir=$(dirname "$0")

echo "Creating storage layout diagrams.."

echo "=======================" > $dir/../.storage-layout
echo "👁👁 STORAGE LAYOUT snapshot 👁👁" >> $dir/../.storage-layout
echo "=======================" >> $dir/../.storage-layout

for contract in ${contracts[@]}
do
  echo -e "\n=======================" >> $dir/../.storage-layout
  echo "➡ $contract">> $dir/../.storage-layout
  echo -e "=======================\n" >> $dir/../.storage-layout
  forge inspect --pretty $contract storage-layout >> $dir/../.storage-layout
done
echo "Storage layout snapshot stored at $dir/../.storage-layout"
