package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	number, err := strconv.Atoi(os.Args[1])

	if err != nil {
		log.Fatal("error: ", err)
	}

	mapping := map[int]string{
		3: "fizz ",
		5: "buzz ",
		7: "bong",
	}

	for key, value := range mapping {
		if number%key == 0 {
			fmt.Print(value)
		}
	}

	fmt.Println()
}
