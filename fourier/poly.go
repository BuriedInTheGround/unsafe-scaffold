package fourier

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// Convolve calculates the convolution of two polynomials of same degree, given
// the coefficients vectors.
func Convolve(a, b []complex128) []complex128 {
	n := len(a)
	if len(b) > len(a) {
		n = len(b)
	}

	// Round 2n-1 up to the next power of 2.
	s := 2*n - 2
	s |= s >> 1
	s |= s >> 2
	s |= s >> 4
	s |= s >> 8
	s |= s >> 16
	s |= s >> 32
	s += 1

	extendedA := make([]complex128, s)
	extendedB := make([]complex128, s)
	res := make([]complex128, s)

	copy(extendedA, a)
	copy(extendedB, b)

	evalsA := FFT(extendedA)
	evalsB := FFT(extendedB)

	for i := range res {
		res[i] = evalsA[i] * evalsB[i]
	}

	return InverseFFT(res)[:len(b)+n-1]
}

func PolyString(coeffs []complex128) string {
	var res string

	for i := range coeffs {
		coeff, _ := strconv.Atoi(fmt.Sprintf("%.f", real(coeffs[i])))
		if coeff == 0 {
			continue
		}
		if i == 7 {
			fmt.Fprintf(os.Stderr, "%#v\n", coeffs[i])
		}

		if i == 0 {
			res += fmt.Sprintf("%d", coeff)
			continue
		}

		if coeff > 0 {
			res += " +"
		} else {
			res += " -"
		}

		if math.Abs(float64(coeff)) != 1 {
			res += fmt.Sprintf(" %.f", math.Abs(float64(coeff)))
		} else {
			res += " "
		}

		if i > 1 {
			res += fmt.Sprintf("x^%d", i)
		} else {
			res += "x"
		}
	}

	return res
}
