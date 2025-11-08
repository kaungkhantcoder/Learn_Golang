package main

import (
	"fmt"
	"math"
)

// Function represents a mathematical function f(x) that takes a float64
// and returns a float64. This is the function we want to integrate.
type Function func(x float64) float64

// DefiniteIntegral approximates the definite integral of f(x) from a to b
// using the Trapezoidal Rule with a specified number of steps (n).
// This function is an implementation of the numerical principle behind
// the Fundamental Theorem of Calculus Part 2 (FTC2).
func DefiniteIntegral(f Function, a, b float64, n int) float64 {
	// Ensure n is a positive integer to avoid division by zero or nonsensical results
	if n <= 0 {
		n = 1000 // Default to 1000 steps if n is invalid
	}

	// Calculate the width of each trapezoid (h or delta x)
	h := (b - a) / float64(n)

	// Initialize the total area with the first and last point contribution
	// The Trapezoidal Rule weights the endpoints by 1/2
	area := 0.5 * (f(a) + f(b))

	// Sum the function values at the interior points (x_1 to x_{n-1})
	for i := 1; i < n; i++ {
		x_i := a + float64(i)*h
		area += f(x_i)
	}

	// Multiply the sum by the width h (delta x) to get the final area approximation
	return area * h
}


// f(x) = x^2
func xSquared(x float64) float64 {
	return x * x
}

// f(x) = sin(x)
func sine(x float64) float64 {
	return math.Sin(x)
}


func FuncDefiniteIntegral() {
	f1 := xSquared
	a1 := 0.0
	b1 := 3.0
	n := 10000
	
	result1 := DefiniteIntegral(f1, a1, b1, n)
	fmt.Printf("1. Integral of x^2 from %.1f to %.1f (n=%d):\n", a1, b1, n)
	fmt.Printf("   Approximation: %.10f\n", result1)
	fmt.Printf("   Actual Value:  9.0\n")
	
	fmt.Println("---")


	f2 := sine
	a2 := 0.0
	b2 := math.Pi
	
	result2 := DefiniteIntegral(f2, a2, b2, n)
	fmt.Printf("2. Integral of sin(x) from %.1f to Pi (n=%d):\n", a2, n)
	fmt.Printf("   Approximation: %.10f\n", result2)
	fmt.Printf("   Actual Value:  2.0\n")
}