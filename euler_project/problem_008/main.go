package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./input")

	if err != nil {
		log.Fatal(err)
	}

	digitText := strings.Replace(string(input), "\n", "", -1)
	result := 0

	for i := 0; i < len(digitText)-13; i++ {
		tmp := 1
		for j := 0; j < 13; j++ {
			number, err := strconv.Atoi(string(digitText[i+j]))

			if err != nil {
				log.Fatal(err)
			}

			tmp *= number
		}
		if tmp > result {
			result = tmp
		}
	}
	fmt.Println(result)
}
