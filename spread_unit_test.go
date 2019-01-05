// +build unit

package statsArb

import (
	"fmt"
	"testing"
)

func TestCalcSpread(t *testing.T) {
	tests := [][4]float64{
		[4]float64{1, 1, 1, 0},
		[4]float64{0, 1, 1, -1},
		[4]float64{1, 0, 1, 1},
		[4]float64{1, 1, 0, 1},
		[4]float64{3, 4, 2, -5},
	}

	for i, tt := range tests {
		res := CalcSpread(tt[0], tt[1], tt[2])
		if fmt.Sprintf("%.4f", tt[3]) != fmt.Sprintf("%.4f", res) {
			t.Errorf("test %d faild, expected %s received %s", i+1, fmt.Sprintf("%.4f", tt[3]), fmt.Sprintf("%.4f", res))
		}
	}
}

func TestCalcSpreadReturn(t *testing.T) {
	tests := [][4]float64{
		[4]float64{1, 1, 1, 0},
		[4]float64{0, 1, 1, -1},
		[4]float64{3, 4, 2, -0.625},
	}

	for i, tt := range tests {
		res := CalcSpreadReturn(tt[0], tt[1], tt[2])
		if fmt.Sprintf("%.4f", tt[3]) != fmt.Sprintf("%.4f", res) {
			t.Errorf("test %d faild, expected %s received %s", i+1, fmt.Sprintf("%.4f", tt[3]), fmt.Sprintf("%.4f", res))
		}
	}
}

func TestCalcDist(t *testing.T) {
	// 1. use birth data
	data, err := fetchFemaleBirthsData()
	if err != nil {
		t.Fatalf("err fetching birth data\n%v", err)
	}
	zData, err := ZNorm(data)
	if err != nil {
		t.Fatalf("err normalizing birth data\n%v", err)
	}

	dist, err := CalcDist(data, zData)
	if err != nil {
		t.Errorf("err calcing dist\n%v", err)
	}

	if fmt.Sprintf("%.4f", 0.0) != fmt.Sprintf("%.4f", dist) {
		t.Errorf("expected %s received %s", fmt.Sprintf("%.4f", 0.0), fmt.Sprintf("%.4f", dist))
	}

	// 2. use tech data
	data1, err := fetchTechStockData()
	if err != nil {
		t.Fatalf("err fetching birth data\n%v", err)
	}
	zData10, err := ZNorm(data1[0])
	if err != nil {
		t.Fatalf("err normalizing first tech data\n%v", err)
	}
	zData11, err := ZNorm(data1[1])
	if err != nil {
		t.Fatalf("err normalizing second tech data\n%v", err)
	}

	dist, err = CalcDist(zData10, zData11)
	if err != nil {
		t.Errorf("err calcing dist\n%v", err)
	}

	if fmt.Sprintf("%.4f", 757.4071870761484) != fmt.Sprintf("%.4f", dist) {
		t.Errorf("expected %s received %s", fmt.Sprintf("%.4f", 757.9458835107689), fmt.Sprintf("%.4f", dist))
	}
}
