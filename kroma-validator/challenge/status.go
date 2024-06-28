package challenge

const (
	// StatusNone is regarded as a challenge that is not in progress.
	// The other status is regarded as a challenge in progress.
	StatusNone uint8 = iota
	StatusChallengerTurn
	StatusAsserterTurn
	StatusChallengerTimeout
	StatusAsserterTimeout
	StatusReadyToProve
)
