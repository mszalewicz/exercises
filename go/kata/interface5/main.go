package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

func Information(shape Shape) string {
	return fmt.Sprintf("Shape - area = %f, perimeter = %f", shape.area(), shape.perimeter())
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

func (circle Circle) String() string {
	return fmt.Sprint("Circle - radius = ", circle.radius)
}

func isCircle(shape Shape) {
	if circle, ok := shape.(Circle); ok {
		fmt.Println(circle)
		fmt.Println(Information(circle))
	} else {
		fmt.Println("Not a circle.")
	}
}

func main() {
	circle := Circle{radius: 2.489}

	isCircle(circle)
}
