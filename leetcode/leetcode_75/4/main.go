package main

import (
	"fmt"
	"math"
)

// You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in adjacent plots.
//
// Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n, return true if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule and false otherwise.

func main() {
	a := []int{1, 0, 0, 0, 1}
	n := 2

	fmt.Println(canPlaceFlowers(a, n))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	counter := 0
	lastPossiblePlant := math.MinInt

	for i, val := range flowerbed {
		if val == 0 {
			if lastPossiblePlant != i-1 {
				if isNeighbourEmpty(&flowerbed, i-1) && isNeighbourEmpty(&flowerbed, i+1) {
					counter++

					if counter == n {
						return true
					}

					lastPossiblePlant = i
				}
			}
		}
	}

	if n == 0 {
		return true
	}

	return false
}

func isNeighbourEmpty(flowerbed *[]int, index int) bool {
	if index < 0 || index >= len(*flowerbed) {
		return true
	}

	if (*flowerbed)[index] == 1 {
		return false
	}

	return true
}

// func canPlaceFlowers(flowerbed []int, n int) bool {
// 	lastPossiblePlant := -1

// 	if len(flowerbed) > 1 {
// 		if flowerbed[0] == 0 && flowerbed[1] == 0 {
// 			if n == 0 {
// 				return false
// 			}

// 			if n == 1 {
// 				return true
// 			}

// 			n--
// 			lastPossiblePlant = 0
// 		}
// 	}

// 	if n != 0 {
// 		for i := 1; i < len(flowerbed)-1; i++ {
// 			if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 && lastPossiblePlant != i-1 {
// 				if n == 1 {
// 					return true
// 				}

// 				n--
// 				lastPossiblePlant = i
// 			}
// 		}

// 		return false
// 	}

// 	for i := 1; i < len(flowerbed)-1; i++ {
// 		if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 && lastPossiblePlant != i-1 {
// 			return false
// 		}
// 	}

// 	return true
// }
