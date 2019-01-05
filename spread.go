package statsArb

import (
	"errors"
	"math"
)

func CalcSpread(p, q, hedgeRatio float64) float64 {
	return p - hedgeRatio*q
}

func CalcSpreadReturn(p, q, hedgeRatio float64) float64 {
	return (p - hedgeRatio*q) / (hedgeRatio * q)
}

func CalcDist(l1, l2 []float64) (float64, error) {
	if l1 == nil || len(l1) == 0 {
		return 0, errors.New("l1 is required")
	}
	if l2 == nil || len(l2) == 0 {
		return 0, errors.New("l2 is required")
	}
	if len(l1) != len(l2) {
		return 0, errors.New("lists are not the same length")
	}

	zL1, err := ZNorm(l1)
	if err != nil {
		return 0.0, err
	}
	zL2, err := ZNorm(l2)
	if err != nil {
		return 0.0, err
	}

	dist := 0.0
	for idx := range zL1 {
		dist += math.Pow(zL1[idx]-zL2[idx], 2)
	}

	return dist, nil
}
