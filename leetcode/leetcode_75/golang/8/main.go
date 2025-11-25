package main

import (
	"math"
)

// Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k]. If no such indices exists, return false.

func main() {
	test := []int{1, 5, 0, 4, 1, 3}
	increasingTriplet(test)
}

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	smallest := math.MaxInt
	smallest2 := math.MaxInt

	for _, val := range nums {
		switch {
			case val <= smallest:
				smallest = val
			case val <= smallest2:
				smallest2 = val
			default:
				return true
		}
	}

	return false
}
