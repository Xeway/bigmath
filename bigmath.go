package bigmath

import (
	"math"
	"math/big"
	"strconv"
)

func IntLog10(x *big.Int) float64 {
	isOverflow := func(bigNum *big.Int) bool {
		return bigNum.String() != strconv.Itoa(int(bigNum.Int64()))
	}

	isOf := isOverflow(x)

	if !isOf {
		return math.Log10(float64(x.Int64()))
	}

	num := make([]*big.Int, 0)

	firstIteration := true
	for isOf {
		if firstIteration {
			num = append(num, big.NewInt(0).Sqrt(x))
			firstIteration = false
		} else {
			num = append(num, big.NewInt(0).Sqrt(num[len(num)-1]))
		}
		isOf = isOverflow(num[len(num)-1])
	}

	numMul := math.Pow(2, float64(len(num)))

	bigNumLog := math.Log10(float64(num[len(num)-1].Int64())) * numMul

	return bigNumLog
}

func FloatLog10(x *big.Float) float64 {
	isOverflow := func(bigNum *big.Float) bool {
		_, acc := bigNum.Float64()
		if acc == big.Exact {
			return false
		} else {
			return true
		}
	}

	isOf := isOverflow(x)

	if !isOf {
		v, _ := x.Float64()
		return math.Log10(v)
	}

	num := make([]*big.Float, 0)

	firstIteration := true
	for isOf {
		if firstIteration {
			num = append(num, big.NewFloat(0).Sqrt(x))
			firstIteration = false
		} else {
			num = append(num, big.NewFloat(0).Sqrt(num[len(num)-1]))
		}
		isOf = isOverflow(num[len(num)-1])
	}

	numMul := math.Pow(2, float64(len(num)))

	v, _ := num[len(num)-1].Float64()
	bigNumLog := math.Log10(v) * numMul

	return bigNumLog
}
