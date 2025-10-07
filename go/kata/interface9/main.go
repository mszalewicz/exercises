package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	radius float64
}

func (circle Circle) area() float64 {
	return math.Pi*circle.radius*circle.radius
}

func (circle Circle) perimeter() float64 {
	return 2 * math.Pi * circle.radius
}

func (circle Circle) String() string {
	return fmt.Sprintf("Circle: radius=%v", circle.radius)
}

func isCircle(shape Shape) bool {
	_, ok := shape.(Circle)

	return ok
}

func main() {
	shape := Circle{radius: 346.234}

	if isCircle(shape) {
		fmt.Println(shape)
	}
}