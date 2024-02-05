package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	left  *Node
	right *Node
	value int
}

func maxPathSum(root *Node, maxValue int) int {
	if root == nil {
		return 0
	}

	leftPath := maxPathSum(root.left, maxValue)
	rightPath := maxPathSum(root.right, maxValue)

	maxValue = int(math.Max(float64(maxValue), math.Max(float64(leftPath), float64(rightPath))+float64(root.value)))

	return maxValue
}

func createBTree(numbers [][]int) Node {
	var head Node
	var previousRow []Node
	var currRow []Node

	for y := 0; y < len(numbers); y++ {
		for x := 0; x < len(numbers[y]); x++ {
			newNode := Node{value: numbers[y][x], left: nil, right: nil}

			if x == 0 && y == 0 {
				head = newNode
			} else {
				currRow = append(currRow, newNode)
			}
		}

		if y == 1 {
			head.left = &currRow[0]
			head.right = &currRow[1]
		} else {
			for i := range previousRow {
				previousRow[i].left = &currRow[i]
				previousRow[i].right = &currRow[i+1]
			}
		}

		previousRow = currRow
		currRow = nil
	}
	return head
}

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	var numbers [][]int
	rows := strings.Split(string(input), "\n")

	for _, row := range rows {
		numbersText := strings.Fields(string(row))
		subset := make([]int, len(numbersText))

		for i, text := range numbersText {
			number, err := strconv.Atoi(text)
			if err != nil {
				log.Fatal(err)
			}
			subset[i] = number
		}
		numbers = append(numbers, subset)
	}

	root := createBTree(numbers)
	smallestInt := math.MinInt
	answer := maxPathSum(&root, smallestInt)

	fmt.Println(answer)
}
