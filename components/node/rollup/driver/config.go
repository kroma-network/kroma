package driver

type Config struct {
	// SyncerConfDepth is the distance to keep from the L1 head when reading L1 data for L2 derivation.
	SyncerConfDepth uint64 `json:"syncer_conf_depth"`

	// ProposerConfDepth is the distance to keep from the L1 head as origin when proposing new L2 blocks.
	// If this distance is too large, the proposer may:
	// - not adopt a L1 origin within the allowed time (rollup.Config.MaxProposerDrift)
	// - not adopt a L1 origin that can be included on L1 within the allowed range (rollup.Config.ProposerWindowSize)
	// and thus fail to produce a block with anything more than deposits.
	ProposerConfDepth uint64 `json:"proposer_conf_depth"`

	// ProposerEnabled is true when the driver should propose new blocks.
	ProposerEnabled bool `json:"proposer_enabled"`

	// ProposerStopped is false when the driver should propose new blocks.
	ProposerStopped bool `json:"proposer_stopped"`

	// ProposerMaxSafeLag is the maximum number of L2 blocks for restricting the distance between L2 safe and unsafe.
	// Disabled if 0.
	ProposerMaxSafeLag uint64 `json:"proposer_max_safe_lag"`
}
