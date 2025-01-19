package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

func measure(g geometry) {
	fmt.Println("Geometry:", g)
	fmt.Println("Area:", g.area())
	fmt.Println("Perimeter:", g.perimeter())
	fmt.Println()
}

type Circle struct {
	radius float64
}

func (c *Circle) area() float64 {
	return math.Pi * math.Sqrt(c.radius)
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func detectCircle(g geometry) {
	if c, ok := g.(*Circle); ok {
		fmt.Println("Circle with radius", c.radius)
		fmt.Println()
		return
	}

	fmt.Printf("%v is not a circle\n", g)
	fmt.Println()
}

type Rectangle struct {
	height float64
	width  float64
}

func (r *Rectangle) area() float64 {
	return r.width * r.height
}

func (r *Rectangle) perimeter() float64{
	return 2 * (r.width + r.height)
}

func main() {
	circle := Circle{radius: 2.25}
	rectangle := Rectangle{height: 6.78, width: 13.38894}

	measure(&circle)
	measure(&rectangle)

	detectCircle(&circle)
	detectCircle(&rectangle)
}
