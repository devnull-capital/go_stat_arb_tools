package statsArb

// OptimalPercentPortfolio ...
func OptimalPercentPortfolio(chanceOfWin, avgWin, avgLoss float64) float64 {
	return chanceOfWin*(1.0/avgLoss) - (1.0-chanceOfWin)*(1.0/avgWin)
}
