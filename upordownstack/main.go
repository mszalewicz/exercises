package main

import (
	"fmt"
	"unsafe"
)

func upordown(y *int) bool {
	var x int

	if y == nil {
		return upordown(&x)
	} else {
		return uintptr(unsafe.Pointer(&x)) > uintptr(unsafe.Pointer(y))
	}
}

// Answers question whether stack on the machine grows down or up
func main() {
	isBigger := upordown(nil)

	if isBigger {
		fmt.Println("up")
	} else {
		fmt.Println("down")
	}
}
