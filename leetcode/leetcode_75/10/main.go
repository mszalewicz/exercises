package main

import "fmt"

// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
//
// Note that you must do this in-place without making a copy of the array.

func main() {
	test := []int{0, 0, 1, 0, 3, 1, 2}
	moveZeroes(test)

	fmt.Println(test)
}

func moveZeroes(nums []int) {
	j := 0

	for i, val := range nums {
		if val != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
