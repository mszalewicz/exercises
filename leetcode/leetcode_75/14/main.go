package main

// You are given an integer array nums consisting of n elements, and an integer k.
//
// Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10^(-5) will be accepted.

func findMaxAverage(nums []int, k int) float64 {
	if len(nums) == 0 || k == 0 {
		return 0
	}

	if len(nums) == k {
		var sum float64 = 0
		for _, val := range nums {
			sum += float64(val)
		}

		return sum / float64(k)
	}

	var sum float64 = 0.0
	var maxAverage float64 = 0.0

	for i := 0; i < k; i++ {
		sum += float64(nums[i])
	}

	maxAverage = sum/float64(k)

	for iteration := range len(nums)-k {
		sum = sum - float64(nums[iteration]) + float64(nums[k+iteration])

		newAverage := sum/float64(k)

		if newAverage > maxAverage {
			maxAverage = newAverage
		}
	}

	return maxAverage
}