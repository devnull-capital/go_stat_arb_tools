package statsArb

// note: these lists are not normalized!
// note: this is not working...
//func areNonNormalizedListsStationary(l1, l2 []float64) (bool, float64, error) {
//if l1 == nil || len(l1) == 0 {
//return false, 0.0, errors.New("l1 is required")
//}
//if l2 == nil || len(l2) == 0 {
//return false, 0.0, errors.New("l2 is required")
//}
//if len(l1) != len(l2) {
//return false, 0.0, errors.New("lists are not of equal length")
//}

//zL1, err := ZNorm(l1)
//if err != nil {
//return false, 0.0, err
//}

//zL2, err := ZNorm(l2)
//if err != nil {
//return false, 0.0, err
//}

//beta, err := CalcBetaHat(zL1, zL2)
//if err != nil {
//return false, 0.0, err
//}

//var ut []float64
//for idx := range zL1 {
//ut = append(ut, math.Log(zL1[idx])-math.Log(zL2[idx])-math.Log(beta))
//}

//return augmentedDickeyFullerTest(ut, adf.DefaultPValue, -1)
//}

// note: this is not working.
//func augmentedDickeyFullerTest(series []float64, pVal float64, lag int) (bool, float64, error) {
//if series == nil || len(series) == 0 {
//return false, 0.0, errors.New("series is required")
//}

//t := adf.New(series, pVal, lag)
//t.Run()
//return t.IsStationary(), t.Statistic, nil
//}
