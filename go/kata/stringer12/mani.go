package main

import "fmt"

type User struct {
	name string
	id   string
}

func (user User) String() string {
	return fmt.Sprintf("Username:        %s\nIdentification:  %s", user.name, user.id)
}

func main() {
	u := User{name: "Sarah Krita", id: "WYT-489-HUK-19289"}

	fmt.Println(u)
}
