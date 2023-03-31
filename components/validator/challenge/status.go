package challenge

const (
	StatusNone uint8 = iota
	StatusChallengerTurn
	StatusAsserterTurn
	StatusChallengerTimeout
	StatusAsserterTimeout
	StatusProveReady
)
