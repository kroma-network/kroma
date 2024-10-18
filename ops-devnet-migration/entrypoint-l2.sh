#!/bin/sh
set -exu

VERBOSITY=${GETH_VERBOSITY:-3}
GETH_DATA_DIR=/db
GETH_CHAINDATA_DIR="$GETH_DATA_DIR/geth/chaindata"
GENESIS_FILE_PATH="${GENESIS_FILE_PATH:-/genesis.json}"
CHAIN_ID=$CHAIN_ID
RPC_PORT="${RPC_PORT:-8545}"
WS_PORT="${WS_PORT:-8546}"

apk add --no-cache jq

GENESIS_JSON="./genesis.json"
TIMESTAMP_HEX=$(jq -r '.timestamp' "$GENESIS_JSON")
KROMA_MPT_TIME=$((TIMESTAMP_HEX + 60))

if [ ! -d "$GETH_CHAINDATA_DIR" ]; then
  echo "$GETH_CHAINDATA_DIR missing, running init"
  echo "Initializing genesis."
  geth --verbosity="$VERBOSITY" init \
    --datadir="$GETH_DATA_DIR" \
    "$GENESIS_FILE_PATH"
else
  echo "$GETH_CHAINDATA_DIR exists."
fi

# Warning: Archive mode is required, otherwise old trie nodes will be
# pruned within minutes of starting the devnet.

exec geth \
  --datadir="$GETH_DATA_DIR" \
  --verbosity="$VERBOSITY" \
  --http \
  --http.corsdomain="*" \
  --http.vhosts="*" \
  --http.addr=0.0.0.0 \
  --http.port="$RPC_PORT" \
  --http.api="web3,debug,eth,txpool,net,engine,kroma" \
  --ws \
  --ws.addr=0.0.0.0 \
  --ws.port="$WS_PORT" \
  --ws.origins="*" \
  --ws.api="debug,eth,txpool,net,engine,kroma" \
  --syncmode=full \
  --nodiscover \
  --maxpeers=1 \
  --networkid=$CHAIN_ID \
  --authrpc.addr="0.0.0.0" \
  --authrpc.port="8551" \
  --authrpc.vhosts="*" \
  --authrpc.jwtsecret=/config/jwt-secret.txt \
  --gcmode=archive \
  --trace.mptwitness=1 \
  --override.mpt="$KROMA_MPT_TIME" \
  "$@"
