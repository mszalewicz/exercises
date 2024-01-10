package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Node struct {
	name           string
	leftChildText  string
	rightChildText string
	left           *Node
	right          *Node
}

func connectGraph(graph *[]Node) {
	for i := 0; i < len(*graph); i++ {
		for j := 0; j < len(*graph); j++ {
			if (*graph)[i].leftChildText == (*graph)[j].name {
				(*graph)[i].left = &(*graph)[j]
			}
			if (*graph)[i].rightChildText == (*graph)[j].name {
				(*graph)[i].right = &(*graph)[j]
			}
		}
	}
}

func countSteps(graph *[]Node, startingNodeIndex int, instructions string, condition string) int {
	steps := 0
	node := (*graph)[startingNodeIndex]

	for {
		if instructions[steps%len(instructions)] == 'L' {
			node = *node.left
		} else {
			node = *node.right
		}

		steps += 1

		if node.name == "ZZZ" {
			break
		}
	}

	return steps
}

func main() {
	input, err := os.ReadFile("input/day_08")

	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	lines := strings.Split(string(input), "\n")
	instructions := lines[0]
	nodes := lines[2:]

	var graph []Node
	startingNodeIndex := 0

	for index, node := range nodes {
		fields := strings.Fields(node)

		var newNode Node
		newNode.name = fields[0]
		newNode.leftChildText = fields[2][1:4]
		newNode.rightChildText = fields[3][0:3]

		if newNode.name == "AAA" {
			startingNodeIndex = index
		}

		graph = append(graph, newNode)
	}

	connectGraph(&graph)
	steps := countSteps(&graph, startingNodeIndex, instructions, "ZZZ")

	elapsed := time.Since(start)
	fmt.Println(steps)
	fmt.Println("time " + fmt.Sprint(elapsed))
}
