package statsArb

// OptimalStopLoss ...
func OptimalStopLoss(M, gain, loss, probCorrect float64) float64 {
	return 0.5 * (M + (1.0-probCorrect)*loss - probCorrect*gain)
}

// OptimalVarStopLoss ...
// return is long / short
func OptimalVarStopLoss(b, variance, entry float64) (float64, float64) {
	return (1.0 - b*variance) * entry, (1.0 + b*variance) * entry
}
