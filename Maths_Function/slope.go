package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func slopeFun(p1, p2 Point) float64 {

	// Calculate the change in Y (rise)
	deltaY := p2.Y - p1.Y
	// Calculate the change in X (run)
	deltaX := p2.X - p1.X

	if deltaX == 0 {
		return math.Inf(1)
	}
	return deltaY / deltaX
}

func FindSlope() {

	// Standard Positive Slope m = 1
	point1 := Point{X: 1.0, Y: 2.0}
	point2 := Point{X: 3.0, Y: 4.0}
	slope1 := slopeFun(point1, point2)

	fmt.Printf("Points: (%.1f, %.1f) and (%.1f, %.1f)\n", point1.X, point1.Y, point2.X, point2.Y)
	fmt.Printf("Slope (m) = %.2f\n\n", slope1)

	// Negative Slope m = -2
	point3 := Point{X: 0.0, Y: 5.0}
	point4 := Point{X: 2.0, Y: 1.0}
	slope2 := slopeFun(point3, point4)
	
	fmt.Printf("Points: (%.1f, %.1f) and (%.1f, %.1f)\n", point3.X, point3.Y, point4.X, point4.Y)
	fmt.Printf("Slope (m) = %.2f\n\n", slope2)

	// Zero Slope (Horizontal Line)
	point5 := Point{X: 5.0, Y: 10.0}
	point6 := Point{X: 1.0, Y: 10.0}
	slope3 := slopeFun(point5, point6)
	fmt.Printf("Points: (%.1f, %.1f) and (%.1f, %.1f)\n", point5.X, point5.Y, point6.X, point6.Y)
	fmt.Printf("Slope (m) = %.2f\n\n", slope3)

	// Infinite Slope (Vertical Line)
	point7 := Point{X: -3.0, Y: 5.0}
	point8 := Point{X: -3.0, Y: 10.0}
	slope4 := slopeFun(point7, point8)
	fmt.Printf("Points: (%.1f, %.1f) and (%.1f, %.1f)\n", point7.X, point7.Y, point8.X, point8.Y)
	fmt.Printf("Slope (m) = %v\n", slope4)


}

