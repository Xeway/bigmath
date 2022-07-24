// This package provides mathematical functions that are not provided by the default math/big package (like logarithm).
package bigmath

import (
	"math"
	"math/big"
	"strconv"
)

// This interface contains a function Sign common to all 3 types of the big/math package (*big.Int, *big.Float, *big.Rat).
// If the type is wrong, an error will occurs.
// Otherwise, a same snippet of code can be used for 3 different types (used in Log10).
type Big interface {
	Sign() int
}

// Log10 returns the decimal logarithm of x.
// x has to be of type *big.Int, *big.Float or *big.Rat.
// For example, if x is of type *big.Float, FloatLog10 will be called.
// Special cases are:
//
//	Log(0) = -Inf
//	Log(x < 0) = NaN
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

// IntLog10 returns the decimal logarithm of x which is type *big.Int.
// The special cases are the same as for Log10.
// Use Log10 for a more generic function that treats all types of the math/big package (*big.Int, *big.Float, *big.Rat).
func IntLog10(x *big.Int) float64 {
	isOverflow := func(bigNum *big.Int) bool {
		return bigNum.String() != strconv.Itoa(int(bigNum.Int64()))
	}

	if !isOverflow(x) {
		return math.Log10(float64(x.Int64()))
	}

	x.Sqrt(x)
	numMul := 2

	for isOverflow(x) {
		x.Sqrt(x)
		numMul = numMul << 1
	}

	bigNumLog := math.Log10(float64(x.Int64())) * float64(numMul)

	return bigNumLog
}

// FloatLog10 returns the decimal logarithm of x which is type *big.Float.
// The special cases are the same as for Log10.
// Use Log10 for a more generic function that treats all types of the math/big package (*big.Int, *big.Float, *big.Rat).
func FloatLog10(x *big.Float) float64 {
	isOverflow := func(bigNum *big.Float) bool {
		_, acc := bigNum.Float64()
		if acc == big.Exact {
			return false
		} else {
			return true
		}
	}

	if !isOverflow(x) {
		v, _ := x.Float64()
		return math.Log10(v)
	}

	x.Sqrt(x)
	numMul := 2

	for isOverflow(x) {
		x.Sqrt(x)
		numMul = numMul << 1
	}

	v, _ := x.Float64()
	bigNumLog := math.Log10(v) * float64(numMul)

	return bigNumLog
}

// RatLog10 returns the decimal logarithm of x which is type *big.Rat.
// The special cases are the same as for Log10.
// Use Log10 for a more generic function that treats all types of the math/big package (*big.Int, *big.Float, *big.Rat).
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

	if !isOverflow(x) {
		v, _ := x.Float64()
		return math.Log10(v)
	}

	thRoot := big.NewFloat(0).Sqrt(big.NewFloat(0).SetRat(x))
	numMul := 2

	for isOverflow(thRoot) {
		thRoot.Sqrt(thRoot)
		numMul = numMul << 1
	}

	v, _ := thRoot.Float64()
	bigNumLog := math.Log10(v) * float64(numMul)

	return bigNumLog
}
