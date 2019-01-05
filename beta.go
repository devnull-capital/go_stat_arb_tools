package statsArb

import (
	"errors"
	"math"
)

// CalcBetaHat estimates beta hat
func CalcBetaHat(l1, l2 []float64) (float64, error) {
	if l1 == nil || len(l1) == 0 {
		return 0, errors.New("l1 is required")
	}
	if l2 == nil || len(l2) == 0 {
		return 0, errors.New("l2 is required")
	}
	if len(l1) != len(l2) {
		return 0, errors.New("lists are not the same length")
	}

	var tmp float64
	for idx := range l1 {
		tmp += math.Log(l1[idx] / l2[idx])
	}

	return math.Exp(tmp / float64(len(l1))), nil
}

// CalcHedgeRatio ...
func CalcHedgeRatio(betaHat, sigmaHat float64) float64 {
	return betaHat * (1.0 + 0.5*sigmaHat)
}
