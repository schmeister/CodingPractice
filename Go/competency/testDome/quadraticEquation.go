package testDome

import "math"

func QuadFindRoots(a, b, c float64) (float64, float64) {

	sqrt := math.Sqrt(math.Pow(b, 2) - (4 * a * c))
	a2 := 2 * a
	x1 := ((-1 * b) + sqrt) / a2
	x2 := ((-1 * b) - sqrt) / a2

	return x1, x2
}