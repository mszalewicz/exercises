package main

import (
	"fmt"
	"math"
)

type Shape interface {
	perimeter() float64
	area() float64
}

func PrintShape(shape Shape) {
	fmt.Println("Perimeter = ", shape.perimeter())
	fmt.Println("Area = ", shape.area())
}

type Circle struct {
	radius float64
}

func (c Circle) String() string {
	return fmt.Sprintf("Cricle with radius %.2f", c.radius)
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) area() float64 {
	return math.Pi * math.Pi * c.radius
}

func isCircle(shape Shape) {
	if circle, ok := shape.(Circle); ok {
		fmt.Println("it is a circle")
		PrintShape(circle)
	} else {
		fmt.Println("it is not a circle")
	}
}

func main() {
	circle := Circle{radius: 6.789}

	isCircle(circle)
	fmt.Println(circle)

}