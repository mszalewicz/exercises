package main

import "fmt"

func main() {

	// Solution works by finding consecutive corner values.
	//
	// Smallest spiral:
	//
	//     7 8 9
	//     6 1 2
	//     5 4 3
	//
	// Consecutive corners: 3, 5, 7, 9
	// We can see that all corners differ by 2.
	//
	// Corners of 1 level bigger spiral are: 13, 17, 21, 25
	// Corners differ by 4.
	//
	// Rule holds for next levels of spiral, where each consecutive level adds 2 to the corner difference.
	// Implementation iterates and adds corners of all levels up to spiral with side size = 1001

	const maxSpiralSideSize = 1001

	diagonalsSum := 1
	increase := 2
	current := 1
	step := 1

	for {
		if step%5 == 0 {
			increase += 2
			step = 1

			if increase > maxSpiralSideSize {
				break
			}
		}

		diagonalsSum += current + increase
		current = current + increase
		step++
	}

	fmt.Println(diagonalsSum)
}
