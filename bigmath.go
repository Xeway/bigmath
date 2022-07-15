package bigmath

import (
	"math"
	"math/big"
	"strconv"
	"sync"
)

func IntLog10(x *big.Int) float64 {
	isOverflow := func(bigNum *big.Int) bool {
		return bigNum.String() != strconv.Itoa(int(bigNum.Int64()))
	}

	isOf := isOverflow(x)

	if isOf {
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

		var bigNumLog float64

		var wg sync.WaitGroup
		wg.Add(int(numMul))

		for k := 0; k < int(numMul); k++ {
			go func() {
				bigNumLog += math.Log10(float64(num[len(num)-1].Int64()))

				wg.Done()
			}()
		}
		wg.Wait()

		return bigNumLog
	} else {
		return math.Log10(float64(x.Int64()))
	}
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

	if isOf {
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

		var bigNumLog float64

		var wg sync.WaitGroup
		wg.Add(int(numMul))

		for k := 0; k < int(numMul); k++ {
			go func() {
				v, _ := num[len(num)-1].Float64()
				bigNumLog += math.Log10(v)

				wg.Done()
			}()
		}
		wg.Wait()

		return bigNumLog
	} else {
		v, _ := x.Float64()
		return math.Log10(v)
	}
}
