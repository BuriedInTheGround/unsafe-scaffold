package fourier

import (
	"math"
	"math/cmplx"
)

func FFT(coeffs []complex128) []complex128 {
	n := len(coeffs)

	// Round n up to the next power of 2.
	s := n - 1
	s |= s >> 1
	s |= s >> 2
	s |= s >> 4
	s |= s >> 8
	s |= s >> 16
	s |= s >> 32
	s += 1

	extendedCoeffs := make([]complex128, s)
	copy(extendedCoeffs, coeffs)

	return fft(extendedCoeffs, false)
}

func InverseFFT(evals []complex128) []complex128 {
	n := len(evals)

	// Round n up to the next power of 2.
	s := n - 1
	s |= s >> 1
	s |= s >> 2
	s |= s >> 4
	s |= s >> 8
	s |= s >> 16
	s |= s >> 32
	s += 1

	extendedEvals := make([]complex128, s)
	copy(extendedEvals, evals)

	coeffs := fft(extendedEvals, true)

	t := 1 / complex(float64(s), 0)
	for i := range coeffs {
		coeffs[i] *= t
	}

	return coeffs
}

func fft(coeffs []complex128, inverse bool) []complex128 {
	n := len(coeffs)

	// Handle the base case.
	if n == 1 {
		return coeffs
	}

	// Check if the slice of coefficients can be split exactly in a half.
	if n%2 != 0 {
		panic("fourier: length not a power of two")
	}

	// Initialize ωₙ and (ωₙ)⁰.
	var omegaN complex128
	if inverse {
		omegaN = cmplx.Rect(1, -2.0*math.Pi/float64(n))
	} else {
		omegaN = cmplx.Rect(1, 2.0*math.Pi/float64(n))
	}
	kThOmegaN := complex128(1)

	// Split the coefficients slice in two halves, one with the elements in
	// even positions and the other with the odd ones.
	evenCoeffs := make([]complex128, 0, n/2)
	oddCoeffs := make([]complex128, 0, n/2)
	for i := 0; i < n; i += 2 {
		evenCoeffs = append(evenCoeffs, coeffs[i])
		oddCoeffs = append(oddCoeffs, coeffs[i+1])
	}

	// Recursively calculate the FFT.
	evenEvals := fft(evenCoeffs, inverse)
	oddEvals := fft(oddCoeffs, inverse)

	// Reconstruct the evaluations slice using the polynomial identity
	// property.
	evals := make([]complex128, n)
	for k := 0; k < n/2; k++ {
		t := kThOmegaN * oddEvals[k]
		evals[k] = evenEvals[k] + t
		evals[k+n/2] = evenEvals[k] - t
		kThOmegaN = kThOmegaN * omegaN
	}

	return evals
}
