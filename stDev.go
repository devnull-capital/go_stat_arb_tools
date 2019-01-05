package statsArb

import (
	"errors"

	"gonum.org/v1/gonum/stat"
)

// StdDev calculates the standard deviation
func StdDev(l1, weights []float64) (float64, error) {
	if l1 == nil || len(l1) == 0 {
		return 0, errors.New("list is required")
	}

	return stat.StdDev(l1, weights), nil
}
