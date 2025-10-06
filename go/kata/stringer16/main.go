package main

import "fmt"

type Shape struct {

}

type Circle struct {
	radius float64
}

func(c Circle) String() string {
	return fmt.Sprintf("circle info: radius=%v", c.radius)
}

func main() {
	circle := Circle{radius:5.64}
	fmt.Println(circle)
}