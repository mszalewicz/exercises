package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	mapping := map[int]string{
		3: "fizz",
		5: "buzz",
		7: "clank",
	}

	number, err := strconv.Atoi(os.Args[1])

	if err != nil {
		log.Fatal("error: ", err)
	}

	var builder strings.Builder

	for key, val := range mapping {
		if number%key == 0 {
			builder.WriteString(val+" ")
		}
	}

	fmt.Println(builder.String())
}
