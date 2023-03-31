<!-- DOCTOC SKIP -->

<!-- All glossary references in this file. -->

[g-block]: glossary#block
[g-l1]: glossary#l1
[g-l2]: glossary#l2
[g-validator]: glossary.md#validator

# Batch Submitter

The batch submitter, also referred to as the batcher, is the entity submitting the [L2][g-l2] proposer data to
[L1][g-l1], to make it available for [validators][g-validator].

[derivation spec]: derivation.md

The format of the data transactions is defined in the [derivation spec]:
the data is constructed from L2 blocks in the reverse order as it is derived from data into L2 [blocks][g-block].

The timing, operation and transaction signing is implementation-specific: any data can be submitted at any time,
but only the data that matches the [derivation spec] rules will be valid from the validator perspective.

The most minimal batcher implementation can be defined as a loop of the following operations:

1. See if the `unsafe` L2 block number is past the `safe` block number: `unsafe` data needs to be submitted.
2. Iterate over all unsafe L2 blocks, skip any that were previously submitted.
3. Open a channel, buffer all the L2 block data to be submitted,
   while applying the encoding and compression as defined in the [derivation spec].
4. Pull frames from the channel to fill data transactions with, until the channel is empty.
5. Submit the data transactions to L1.

The L2 view of safe/unsafe does not instantly update after data is submitted, nor when it gets confirmed on L1,
so special care may have to be taken to not duplicate data submissions.
