// +build unit

package statsArb

import (
	"fmt"
	"testing"
)

func TestSigmaHat(t *testing.T) {
	// 1. female births
	data, err := fetchFemaleBirthsData()
	if err != nil {
		t.Fatalf("err fetching female births data\n%v", err)
	}

	betaHat, err := CalcBetaHat(data, data)
	if err != nil {
		t.Fatalf("err calcing beta hat\n%v", err)
	}
	sigmaHat, err := CalcSigmaHat(data, data, betaHat)
	if err != nil {
		t.Fatalf("err calcing sigma hat\n%v", err)
	}
	if sigmaHat != 0.0 {
		t.Fatalf("expected 0 received %v", sigmaHat)
	}

	// 2. tech stocks
	data1, err := fetchTechStockData()
	if err != nil {
		t.Fatalf("err fetching tech stock data\n%v", err)
	}

	betaHat, err = CalcBetaHat(data1[0], data1[1])
	if err != nil {
		t.Fatalf("err calcing beta hat\n%v", err)
	}
	sigmaHat, err = CalcSigmaHat(data1[0], data1[1], betaHat)
	if err != nil {
		t.Fatalf("err calcing sigma hat\n%v", err)
	}
	if fmt.Sprintf("%.4f", 0.13820923793111195) != fmt.Sprintf("%.4f", sigmaHat) {
		t.Errorf("expected %s received %s", fmt.Sprintf("%.4f", 0.13820923793111195), fmt.Sprintf("%.4f", sigmaHat))
	}
}

func TestEffectiveStdDev(t *testing.T) {
	tests := [][4]float64{
		[4]float64{1, 1, 1, 1.5811388300842},
		[4]float64{0, 1, 1, 0},
		[4]float64{3, 4, 1, 1.1858541225631},
	}

	for i, tt := range tests {
		res := CalcEffectiveStandardDeviation(tt[0], tt[1], tt[2])

		if fmt.Sprintf("%.4f", tt[3]) != fmt.Sprintf("%.4f", res) {
			t.Errorf("test #%d failed, expected %s, received %s", i+1, fmt.Sprintf("%.4f", tt[3]), fmt.Sprintf("%.4f", res))
		}
	}
}
