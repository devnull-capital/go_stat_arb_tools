package statsArb

import (
	"errors"

	"gonum.org/v1/gonum/stat"
)

// ZNorm calculates the ZNormalized list defined as
// Z = ( x - mean ) / std
func ZNorm(l1 []float64) ([]float64, error) {
	if l1 == nil || len(l1) == 0 {
		return nil, errors.New("l1 is required")
	}

	std := stat.StdDev(l1, nil)
	if std == 0 {
		return nil, errors.New("std dev is zero")
	}

	mean := stat.Mean(l1, nil)

	for idx := range l1 {
		l1[idx] = (l1[idx] - mean) / std
	}

	return l1, nil
}
