package main

import (
	"fmt"
	"math"

	"github.com/BuriedInTheGround/unsafe-scaffold/fourier"
)

func main() {
	a := []complex128{-15, -2, 1}
	b := []complex128{-1, 3, -3, 21, 1}
	fmt.Printf("a = %s\n", fourier.PolyString(a))
	fmt.Printf("b = %s\n\n", fourier.PolyString(b))

	y := fourier.FFT(a)
	fmt.Printf("y = FFT(a) = %.1f\n", y)

	aPrime := fourier.InverseFFT(y)
	fmt.Printf("a' = InverseFFT(y) = %s\n\n", fourier.PolyString(aPrime))

	z := fourier.Convolve(a, b)
	fmt.Printf("z = a*b = %s\n\n", fourier.PolyString(z))

	// CARTESIAN_SUM(a, b)
	a = []complex128{1, 0, 1, 1}
	b = []complex128{0, 1, 1, 0}
	z = fourier.Convolve(a, b)
	fmt.Printf("u = (1 0 1 1)*(0 1 1 0) = %.0f\n\n", z)

	x := []complex128{5, 3, 3, 3, 4, 3, 3, 3}
	fmt.Printf("x = %s\n", fourier.PolyString(x))
	y = fourier.FFT(x)
	fmt.Printf("y = FFT(x) = %.1f\n\n", y)

	x = []complex128{5, 3, 0, 0}
	fmt.Printf("x = %s\n", fourier.PolyString(x))
	y = fourier.FFT(x)
	fmt.Printf("y = FFT(x) = %.1f\n\n", y)

	x = []complex128{
		complex(1, 0),
		complex(1, 0),
		complex(1, 0),
		complex(1, 0),
		complex(1, 0),
		complex(1, 0),
		complex(1, 0),
		complex(1, 0),
	}
	fmt.Printf("x = %v\n", x)
	y = fourier.FFT(x)
	fmt.Printf("DFTn(8, 0) = %.1f\n", DFTOfFourierMatrixColumn(8, 0))
	fmt.Printf("y = FFT(x) = %.1f\n\n", y)

	x = []complex128{
		complex(1, 0),
		complex(1/math.Sqrt(2), 1/math.Sqrt(2)),
		complex(0, 1),
		complex(-1/math.Sqrt(2), 1/math.Sqrt(2)),
		complex(-1, 0),
		complex(-1/math.Sqrt(2), -1/math.Sqrt(2)),
		complex(0, -1),
		complex(1/math.Sqrt(2), -1/math.Sqrt(2)),
	}
	fmt.Printf("x = %v\n", x)
	y = fourier.FFT(x)
	fmt.Printf("DFTn(8, 1) = %.1f\n", DFTOfFourierMatrixColumn(8, 1))
	fmt.Printf("y = FFT(x) = %.1f\n\n", y)

	x = []complex128{
		complex(1, 0),
		complex(0, 1),
		complex(-1, 0),
		complex(0, -1),
		complex(1, 0),
		complex(0, 1),
		complex(-1, 0),
		complex(0, -1),
	}
	fmt.Printf("x = %v\n", x)
	y = fourier.FFT(x)
	fmt.Printf("DFTn(8, 2) = %.1f\n", DFTOfFourierMatrixColumn(8, 2))
	fmt.Printf("y = FFT(x) = %.1f\n\n", y)

	x = []complex128{4, 4, 4, 4}
	fmt.Printf("x = %s\n", fourier.PolyString(x))
	y = fourier.FFT(x)
	fmt.Printf("y = FFT(x) = %.1f\n", y)
	z = fourier.InverseFFT(y)
	fmt.Printf("z = InverseFFT(y) = %s\n\n", fourier.PolyString(z))

	x = []complex128{4, 0, 0, 0}
	fmt.Printf("x = %s\n", fourier.PolyString(x))
	y = fourier.FFT(x)
	fmt.Printf("y = FFT(x) = %.1f\n", y)
	z = fourier.InverseFFT(y)
	fmt.Printf("z = InverseFFT(y) = %s\n\n", fourier.PolyString(z))

	x = []complex128{4, 0, 0, 0, 2, 0, 0, 0}
	fmt.Printf("x = %s\n", fourier.PolyString(x))
	y = fourier.FFT(x)
	fmt.Printf("y = FFT(x) = %.1f\n", y)
	z = fourier.InverseFFT(y)
	fmt.Printf("z = InverseFFT(y) = %s\n\n", fourier.PolyString(z))

	g := []complex128{
		1.0,
		0,
		-39.48 / 2.0,
		0,
		1558.55 / 24.0,
		0,
		-61528.91 / 720.0,
		0,
	}
	fmt.Printf("g = %.1f\n", g)
	fmt.Printf("FFT(g) = %.1f\n", fourier.FFT(g))
}

func DFTOfFourierMatrixColumn(order, column int) []complex128 {
	if order < 1 {
		panic("order must be greater or equal to 1")
	}
	if column < 0 || column >= order {
		panic("column out of bounds")
	}
	res := make([]complex128, order)
	index := order - column
	res[index%order] = complex(float64(order), 0)
	return res
}
