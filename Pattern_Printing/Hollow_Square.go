package main

import (
	"fmt"
)

func printHollowSquare(N int) {
	
	if N <= 0 {
		fmt.Println("Side length must be a positive integer.")
		return
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ { 

			if i == 1 || i == N || j == 1 || j == N {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func HollowSquare() {
	sideLength := 5 

	fmt.Printf("Hollow Square Pattern (N=%d):\n", sideLength)
	printHollowSquare(sideLength)

	fmt.Println("\n---")
	
	sideLength = 8 
	fmt.Printf("Hollow Square Pattern (N=%d):\n", sideLength)
	printHollowSquare(sideLength)
}