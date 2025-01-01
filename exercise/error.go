package main

import (
	"fmt"
	"math"
)

// Define a custom error type
type ErrNegativeSqrt float64

// Implement the error interface for ErrNegativeSqrt
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt computes the square root of a given number.
// Returns an error if the input is negative.
func SqureRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x) // Return custom error for negative input
	}
	return math.Sqrt(x), nil // Return square root for non-negative input
}

func errorExercise() {
	// Test cases
	fmt.Println(Sqrt(2))  // Should print the square root of 2
	fmt.Println(Sqrt(-2)) // Should print an error message
}
