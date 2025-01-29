package main

import "fmt"

type AIO struct {
	name    int
	surname string
}

func (aio AIO) String() string {
	return fmt.Sprintf("AIO - name: %d, surname: %s", aio.name, aio.surname)
}

func main() {

	test := AIO{
		name:    1,
		surname: "test",
	}

	fmt.Println(test)
}
