package main

import (
	"fmt"
)

// using the Newton-Raphson method
func Sqrt(x float64) float64 {
	if x < 0 {
		return 0 // Return 0 for negative input as sqrt is not defined
	}

	// Initial guess
	z := x / 2.0
	epsilon := 1e-10 // Precision threshold

	// Iteratively improve the guess
	for {
		newZ := z - (z*z-x)/(2*z)

		if absoluteValue(newZ-z) < epsilon {
			break
		}

		z = newZ
	}
	return z
}

// absoluteValue calculates the absolute value of a float64
func absoluteValue(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(Sqrt(2))
}
