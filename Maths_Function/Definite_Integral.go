package main

import (
	"fmt"
	"math"
)

type Function func(x float64) float64

// DefiniteIntegral approximates the definite integral of f(x) from a to b
// using the Trapezoidal Rule with a specified number of steps (n).
// This function is an implementation of the numerical principle behind
// the Fundamental Theorem of Calculus Part 2 (FTC2).
func DefiniteIntegral(f Function, a, b float64, n int) float64 {

	if n <= 0 {
		n = 1000 
	}

	h := (b - a) / float64(n)

	area := 0.5 * (f(a) + f(b))

	for i := 1; i < n; i++ {
		x_i := a + float64(i)*h
		area += f(x_i)
	}

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