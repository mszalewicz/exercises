package main

import "fmt"

func main() {

	sumOfSquares := 0
	squareSum := 0

	for i := 1; i <= 100; i++ {
		sumOfSquares += i * i
		squareSum += i
	}

	squareSum *= squareSum

	fmt.Println(squareSum - sumOfSquares)
}
