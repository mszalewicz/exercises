package main

import "fmt"

type User struct {
	id string
}

func (user User) String() string {
	return fmt.Sprintf("ID: %s", user.id)
}

func main() {
	user := User{id: "kjhjfghdfgsdfs-10-bcvd"}

	fmt.Println(user)
}