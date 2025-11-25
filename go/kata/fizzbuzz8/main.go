package main

import "fmt"

func main() {
	mapping := map[int]string{
		3: "fizzz",
		5: "buzz",
		8: "nice",
	}

	numbers := []int{3, 5, 8, 120, 8765432}

	for _, number := range numbers {
		fmt.Print(number, "  ->   ")
		for key, value := range mapping {
			if number%key == 0 {
				fmt.Print(value, " ")
			}
		}
		fmt.Println()
	}
}
