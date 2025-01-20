package main

import (
	"math"
)

// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
//
// Find two lines that together with the x-axis form a container, such that the container contains the most water.
//
// Return the maximum amount of water a container can store.
//
// Notice that you may not slant the container.

func maxArea(height []int) int {
	maxArea := math.MinInt

	start, end := 0, len(height)-1

	for {
		area := min(height[start], height[end]) * (end-start)
		if area > maxArea {
			maxArea = area
		}

		if height[start] < height[end] {
			start++
		} else {
			end--
		}

		if start == end {
			return maxArea
		}
	}
}

func min(a,b int) int {
	if a < b {
		return a
	}

	return b
}
