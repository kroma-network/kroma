package validator

const (
	// StatusNone is regarded as a challenge is not in progress.
	// The other status are regarded as a challenge is in progress.
	StatusNone uint8 = iota
	StatusInactive
	StatusActive
	StatusCanStart
	StatusStarted
	StatusCanSubmitOutput
)
