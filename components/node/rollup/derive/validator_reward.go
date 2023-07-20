package derive

const TempValidatorRewardScalar = uint64(2000) // Denominator is 10000, so the ratio is 20%

func CalcValidatorRewardScalar() uint64 {
	// NOTE(pangssu): This is where the scalar for distributing transaction fees as validator rewards is determined.
	// This section requires additional logic to calculate the scalar
	// and is currently set to a temporary constant value of 2000.
	return TempValidatorRewardScalar
}
