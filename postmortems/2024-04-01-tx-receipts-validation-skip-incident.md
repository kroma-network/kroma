# 2024-04-01 Chain Halt due to Transaction Receipts Validation Skip Post-Mortem

# Incident Summary

On April 1, 2024, at 02:51:15 UTC, a wrong block was generated at block number [9029744](https://kromascan.com/block/9029744).

Some nodes ignored the batch due to a hash discrepancy in that block compared to the block hash of the sequencer. This 
discrepancy was caused due to that the sequencer omitted a deposited transaction from L1.

The sequencer is recovered by rollback to block number 9029744. Block generation was halted for about 3 hours.

# Background

We operate our own L1 RPC nodes and have duplicated L1 RPC to ensure stable operation of L2 sequencer and RPC nodes. If 
the block head of the active L1 RPC falls behind a certain threshold, it is configured to look towards the fallback RPC 
node. Additionally, all kroma-nodes, including sequencers, have been running with the --l1.trustrpc option, using our L1
RPC as the trust RPC through this option.

# Causes

Block head of our L1 RPC began to fall behind the canonical chain, triggering a switch to the fallback L1 RPC once it 
deviated beyond a certain threshold. During the switch, sequencer attempted to retrieve transaction receipts from the L1
RPC node that fell behind it had not yet synced, resulting in empty transaction receipt responses.

Actually, the block in question did contain [deposited transaction](https://kromascan.com/tx/0xf76e4f34645bc3e172909fb03311cd1770f9c574ce46bd1da3644f9cea82c0e3)
on L1, yet the sequencer proceeded to generate the block without processing this transaction, having received empty 
transaction receipts from the trusted L1 RPC. As a consequence, starting from block 9029744, a different block was 
generated, leading to a generation of wrong block. 
[This](https://kromascan.com/tx/0xf76e4f34645bc3e172909fb03311cd1770f9c574ce46bd1da3644f9cea82c0e3) is the transaction 
not included in the block sequencer generated.

# Recovery

We recovered by resyncing the kroma node from the snapshot data for backup, ensuring the correct processing of deposit 
transactions from block 9029744 to generate the block correctly.

## Timeline (UTC)

- 2024-04-01 0251: wrong block generated at 9029744
- 2024-04-01 0258: notified that block generation was halted by Wemade
- 2024-04-01 0357: stop sequencer and RPC nodes
- 2024-04-01 0358: announced incident occurrence on discord, twitter, and other partners
- 2024-04-01 0453: restarted a sequencer with backup snapshot
- 2024-04-01 0611: recovery of Kroma mainnet (RPC, P2P, validator) completed
- 2024-04-01 0618: announced the recovery of Kroma mainnet on partners
- 2024-04-01 0620: announced the recovery of Kroma mainnet on discord and twitter

# How it is fixed

It is fixed by re-syncing the chain from backup snapshot data and re-generating the block using correct L1 blocks.

# Future works

## Removal of the `--l1.trustrpc`

The --l1.trustrpc option is implemented to skip the validation of transaction receipts which is retrieved from the 
trusted source. It will be turned off to allow all Kroma nodes to validate transaction receipts.

## Enhancement of monitoring system

We are going to strengthen the monitoring system by establishing a system to monitor blocks from trust partners, 
enabling swift detection of incidents and improving our ability to respond to incidents.
