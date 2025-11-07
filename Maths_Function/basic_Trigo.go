package main

import (
	"fmt"
	"math"
)

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func Trigo() {
	// Define an angle in degrees
	angleDegrees := 60.0
	fmt.Printf("--- Calculations for %.0f Degrees ---\n", angleDegrees)

	// Convert the angle to radians
	angleRadians := degreesToRadians(angleDegrees)
	fmt.Printf("Angle in Radians: %.4f\n\n", angleRadians)

	// Calculate Sine (Sin)
	sinValue := math.Sin(angleRadians)
	// For a 60-degree angle, sin(60) should be approximately sqrt(3)/2 ≈ 0.866
	fmt.Printf("Sine (sin(%.0f°)): %.4f\n", angleDegrees, sinValue)

	// Calculate Cosine (Cos)
	cosValue := math.Cos(angleRadians)
	// For a 60-degree angle, cos(60) should be exactly 1/2 = 0.5
	fmt.Printf("Cosine (cos(%.0f°)): %.4f\n", angleDegrees, cosValue)

	// Calculate Tangent (Tan)
	tanValue := math.Tan(angleRadians)
	// For a 60-degree angle, tan(60) should be approximately sqrt(3) ≈ 1.732
	fmt.Printf("Tangent (tan(%.0f°)): %.4f\n\n", angleDegrees, tanValue)

	fmt.Println("--- Inverse Functions (Arcsin) ---")

	restoredAngleRadians := math.Asin(sinValue)
	fmt.Printf("Arcsin(%.4f) in Radians: %.4f\n", sinValue, restoredAngleRadians)
	
	restoredAngleDegrees := restoredAngleRadians * 180 / math.Pi
	fmt.Printf("Arcsin(%.4f) in Degrees: %.0f°\n", sinValue, restoredAngleDegrees)

}