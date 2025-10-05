package main

import "fmt"

type Circle struct {
	radius float64
}

func (c *Circle) String() string {
	return fmt.Sprintf("Circle info: radius=%v", c.radius)
}

func main() {
	circle := Circle{radius: 43.23}
	fmt.Println(&circle)
}
