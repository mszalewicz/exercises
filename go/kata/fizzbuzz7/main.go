package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func usage() {
	usageText := "\nUsage:\n\n./binary number_you_want_to_check\n"

	fmt.Println(usageText)
}

func main() {
	mapping := map[int]string{
		3: "fizz",
		5: "buzz",
		7: "gong",
	}

	if len(os.Args) != 2 {
		usage()
		os.Exit(0)
	}

	numberString := os.Args[1]
	number, err := strconv.Atoi(numberString)
	if err != nil {
		log.Fatal("error: ", err)
	}

	for key, value := range mapping {
		if number % key == 0 {
			fmt.Print(value, " ")
		}
	}

	fmt.Println()
}
