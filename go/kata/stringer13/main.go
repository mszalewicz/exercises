package main

import "fmt"

type User struct {
	name string
	id   string
}

func (user User) String() string {
	return fmt.Sprintf("Username: %s, ID: %s", user.name, user.id)
}

func main() {
	user := User{name: "Joanna", id: "0sfkw-w9we0-fvdgfjn"}

	fmt.Println(user)
}
