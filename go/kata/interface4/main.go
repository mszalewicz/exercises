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
	return math.Pi * circle.radius * circle.radius
}

func (circle Circle) perimeter() float64 {
	return 2 * math.Pi * circle.radius
}

func isCircle(shape Shape) {
	if circle, ok := shape.(Circle); ok {
		fmt.Println("It is circle.")
		fmt.Println("Area:", circle.area())
		fmt.Println("Perimeter:", circle.perimeter())
	} else {
		fmt.Println("It is not a circle.")
	}
}

func main() {
	circle := Circle{radius: 4.65}

	isCircle(circle)
}