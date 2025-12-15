package main

import (
	"fmt"
)

func twoSumHashMap(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if complementIndex, ok := seen[complement]; ok {
			return []int{complementIndex, i}
		}
		seen[num] = i
	}

	// No solution found
	return nil
}

func TwoSum() {
	nums := []int{3,2,4}
	target := 6
	result := twoSumHashMap(nums, target)
	fmt.Printf("Hash Mash Solution:\n")
	if result != nil {
		fmt.Printf("Input: %v, Target: %d\n", nums, target)
		fmt.Printf("Indices: %v (Values: %d + %d = %d)\n", result, nums[result[0]], nums[result[1]], target)
	} else {
		fmt.Println("No solution found.")
	}
}