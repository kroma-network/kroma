package derive

const TempValidatorRewardRatio = uint64(2000) // 20%

func CalcValidatorRewardRatio() uint64 {
	// NOTE(pangssu): This is where the ratio for distributing transaction fees as validator rewards is determined.
	// This section requires additional logic to calculate the ratio
	// and is currently set to a temporary constant value of 20%.
	return TempValidatorRewardRatio
}
