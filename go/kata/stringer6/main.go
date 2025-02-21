package main

import "fmt"

type User struct {
	name string
	id   string
}

func (user User) String() string {
	return fmt.Sprintf("User info: username=%s, id=%s", user.name, user.id)
}

func main() {
	user := User{name: "John", id: "ID-2344-sf9988"}

	fmt.Println(user)
}
