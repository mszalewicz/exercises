package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	mapping := map[int]string{
		3: "fizz",
		5: "buzz",
		7: "blow",
	}

	numberText := os.Args[1]

	number, err := strconv.Atoi(numberText)

	if err != nil {
		log.Fatal("error: ", err)
	}

	for key, value := range mapping {
		if number%key == 0 {
			fmt.Print(value, " ")
		}
	}

	fmt.Println()
}
