package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func (circle Circle) Perimeter() float64 {
	return 2 * math.Pi * circle.radius
}

func isCircle(shape Shape) bool {
	if _, ok := shape.(Circle); ok {
		return true
	}

	return false
}


func main() {
	shape := Circle{radius: 34563.24}

	if isCircle(shape) {
		fmt.Println("it is circle")
	}
}
