// +build unit

package statsArb

import (
	"fmt"
	"testing"
)

func TestCalcBetaHat(t *testing.T) {
	// 1. test the female birth data
	data, err := fetchFemaleBirthsData()
	if err != nil {
		t.Fatalf("err fetching data\n%v", err)
	}

	betaHat, err := CalcBetaHat(data, data)
	if err != nil {
		t.Fatalf("err calcing beta hat\n%v", err)
	}

	if betaHat != 1 {
		t.Errorf("expected 1, received %v", betaHat)
	}

	// 2. test the tech stock data
	data2, err := fetchTechStockData()
	if err != nil {
		t.Fatalf("err fetching data\n%v", err)
	}
	if len(data2) != 2 || data2[0] == nil || data2[1] == nil || len(data2[0]) == 0 || len(data2[1]) == 0 {
		t.Fatalf("invalid data\n%v", data)
	}

	betaHat, err = CalcBetaHat(data2[0], data2[1])
	if err != nil {
		t.Fatalf("err calcing beta hat\n%v", err)
	}

	if fmt.Sprintf("%.4f", betaHat) != fmt.Sprintf("%.4f", 0.1479640618306254) {
		t.Errorf("expected %s, received %s", fmt.Sprintf("%.4f", betaHat), fmt.Sprintf("%.4f", 0.1479640618306254))
	}
}

func TestCalcHedgeRatio(t *testing.T) {
	tests := [][3]float64{
		[3]float64{1, 1, 1.5},
		[3]float64{0, 1, 0},
		[3]float64{1, 0, 1},
		[3]float64{3, 4, 9},
	}

	for i, tt := range tests {
		res := CalcHedgeRatio(tt[0], tt[1])
		if tt[2] != res {
			t.Errorf("test %d failed, expected %v received %v", i+1, tt[2], res)
		}
	}
}
