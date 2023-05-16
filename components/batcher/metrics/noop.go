package metrics

import (
	"github.com/kroma-network/kroma/components/node/eth"
	"github.com/kroma-network/kroma/components/node/rollup/derive"
	kmetrics "github.com/kroma-network/kroma/utils/service/metrics"
	txmetrics "github.com/kroma-network/kroma/utils/service/txmgr/metrics"
)

type noopMetrics struct {
	kmetrics.NoopRefMetrics
	txmetrics.NoopTxMetrics
}

var NoopMetrics Metricer = new(noopMetrics)

func (*noopMetrics) Document() []kmetrics.DocumentedMetric { return nil }

func (*noopMetrics) RecordInfo(version string) {}
func (*noopMetrics) RecordUp()                 {}

func (*noopMetrics) RecordLatestL1Block(l1ref eth.L1BlockRef)               {}
func (*noopMetrics) RecordL2BlocksLoaded(eth.L2BlockRef)                    {}
func (*noopMetrics) RecordChannelOpened(derive.ChannelID, int)              {}
func (*noopMetrics) RecordL2BlocksAdded(eth.L2BlockRef, int, int, int, int) {}

func (*noopMetrics) RecordChannelClosed(derive.ChannelID, int, int, int, int, error) {}

func (*noopMetrics) RecordChannelFullySubmitted(derive.ChannelID) {}
func (*noopMetrics) RecordChannelTimedOut(derive.ChannelID)       {}

func (*noopMetrics) RecordBatchTxSubmitted() {}
func (*noopMetrics) RecordBatchTxSuccess()   {}
func (*noopMetrics) RecordBatchTxFailed()    {}
