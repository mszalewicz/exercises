package main

import "fmt"

type User struct {
	name string
}

func (user User) String() string{
	return fmt.Sprint("Imię: ", user.name)
}

func main() {
	user := User{name: "Paweł"}
	fmt.Println(user)
}