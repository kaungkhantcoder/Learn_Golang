package main

import (
	"fmt"
	"strings"
)

func printDiamond(size int) {
	if size <= 0 {
		fmt.Println("Error: Size must be a postive integer.")
		return
	}

	fmt.Printf("Generating diamond of size %d:\n", size)

	for i := 1; i <= size; i++ {
		spaces := size - 1

		stars := 2*i - 1

		fmt.Printf("%s%s\n", strings.Repeat(" ", spaces), strings.Repeat("*", stars))
	}

	for i := size - 1; i >= 1; i-- {
		spaces := size - i

		stars := 2*i - 1

		fmt.Printf("%s%s\n", strings.Repeat(" ", spaces), strings.Repeat("*", stars))
	}
}

func Diamond() {
	printDiamond(5)

	fmt.Println("\n" + strings.Repeat("-", 30) + "\n")

	printDiamond(3)

	fmt.Println("\n" + strings.Repeat("-", 30) + "\n")

	printDiamond(1)
}