package bigmath

import (
	"math"
	"math/big"
	"strconv"
)

func Log10(x Big) float64 {
	if x.Sign() != 1 {
		if x.Sign() == 0 {
			return math.Inf(-1)
		} else {
			return math.NaN()
		}
	}

	switch x := x.(type) {
	case *big.Int:
		return IntLog10(x)
	case *big.Float:
		return FloatLog10(x)
	case *big.Rat:
		return RatLog10(x)
	default:
		return math.NaN()
	}
}

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

func RatLog10(x *big.Rat) float64 {
	isOverflow := func(bigNum interface{}) bool {
		switch bigNum := bigNum.(type) {
		case *big.Float:
			_, acc := bigNum.Float64()
			if acc == big.Exact {
				return false
			} else {
				return true
			}
		case *big.Rat:
			f, _ := bigNum.Float64()
			if strconv.FormatFloat(f, 'E', -1, 64) == "+Inf" {
				return true
			} else {
				return false
			}
		default:
			return false
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
			num = append(num, big.NewFloat(0).Sqrt(big.NewFloat(0).SetRat(x)))
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

type Big interface {
	Sign() int
}
