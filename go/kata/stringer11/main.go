package main

import "fmt"

type Package struct {
	id       string
	location string
}

func (p Package) String() string {
	return fmt.Sprintf("Pacakge info - id : %s | location : %s", p.id, p.location)
}

func main() {
	p := Package{id: "503daddl--1313", location: "Paris"}

	fmt.Println(p)
}
