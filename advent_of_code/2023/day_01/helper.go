package main

import (
	"bufio"
	"fmt"
	"os"
)

func open_file(path string) (*os.File, *bufio.Scanner) {

	readFile, err := os.Open(path)

	if err != nil {
		fmt.Println("Could not read file from given path. Error:")
		fmt.Println(err)
		os.Exit(1)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	return readFile, fileScanner
}
