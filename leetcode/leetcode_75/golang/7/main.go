package main

import "slices"

// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
//
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
//
// You must write an algorithm that runs in O(n) time and without using the division operation.

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	inbetween := 1

	for index, number := range nums {
		result[index] = inbetween
		inbetween *= number
	}

	inbetween = 1

	for index, number := range slices.Backward(nums) {
		result[index] *= inbetween
		inbetween *= number
	}

	return result
}