package main

import "fmt"

func main() {

	target := 10
	coins := []int{1, 2, 5, 10, 20, 50, 100, 200}
	table := make([]int, target+1)

	table[0] = 1

	for i := 0; i < len(coins); i++ {

		fmt.Printf("i: %d\n", i)

		for j := coins[i]; j <= target; j++ {

			fmt.Printf("j: %d\n", j)

			table[j] += table[j-coins[i]]

		}
	}

	fmt.Println(table)
	fmt.Println(table[target])
}
