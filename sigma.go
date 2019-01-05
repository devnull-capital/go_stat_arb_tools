package statsArb

import (
	"errors"
	"math"
)

// CalcSigmaHat ...
func CalcSigmaHat(l1, l2 []float64, beta float64) (float64, error) {
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
		tmp += math.Pow(math.Log(l1[idx]/l2[idx]), 2)
	}
	tmp /= float64(len(l1))

	return tmp - math.Pow(math.Log(beta), 2), nil
}

// CalcEffectiveStandardDeviation ...
func CalcEffectiveStandardDeviation(betaHat, hedgeRatio, sigmaHat float64) float64 {
	return math.Sqrt((math.Pow(betaHat, 2.0) / math.Pow(hedgeRatio, 2.0)) * (sigmaHat + (3.0/2.0)*math.Pow(sigmaHat, 2.0)))
}
