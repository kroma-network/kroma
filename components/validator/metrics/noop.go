package metrics

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/kroma-network/kroma/components/node/eth"
	kmetrics "github.com/kroma-network/kroma/utils/service/metrics"
	txmetrics "github.com/kroma-network/kroma/utils/service/txmgr/metrics"
)

type noopMetrics struct {
	kmetrics.NoopRefMetrics
	txmetrics.NoopTxMetrics
}

var NoopMetrics Metricer = new(noopMetrics)

func (*noopMetrics) RecordInfo(version string) {}
func (*noopMetrics) RecordUp()                 {}

func (*noopMetrics) RecordL2OutputSubmitted(l2ref eth.L2BlockRef)   {}
func (*noopMetrics) RecordDepositAmount(amount *big.Int)            {}
func (*noopMetrics) RecordNextValidator(address common.Address)     {}
func (*noopMetrics) RecordChallengeCheckpoint(outputIndex *big.Int) {}
