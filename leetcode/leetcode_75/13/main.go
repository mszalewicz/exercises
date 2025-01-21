package main

import (
	"fmt"
	"slices"
)

func main() {
	test := []int{1, 1, 3, 4, 4, 4}
	operations := maxOperations(test, 5)
	fmt.Println(operations)
}

func maxOperations(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	slices.Sort(nums)
	operations := 0

	start := 0
	end := len(nums) - 1

	for start < end {
		switch {
			case nums[start]+nums[end] == k:
				operations++
				start++
				end--
			case nums[start]+nums[end] < k:
				start++
			case nums[start]+nums[end] > k:
				end--
		}
	}

	return operations
}

