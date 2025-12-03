// package main

// import (
// 	"fmt"
// 	"math"
// )

// func Circle() {
// 	const radius = 10.0
// 	const tolerance = 0.7

// 	centerX := radius
// 	centerY := radius

// 	for y := 0.0; y <= 2*radius; y++ {

// 		for x := 0.0; x <= 2*radius; x++ {

// 			dx := x - centerX
// 			dy := y - centerY

// 			distance := math.Sqrt(dx*dx + dy*dy)


// 			if math.Abs(distance - radius) < tolerance {
// 				fmt.Print("* ")
// 			} else {
// 				fmt.Print("  ")
// 			}
// 		}
// 		fmt.Println()
// 	}
// }