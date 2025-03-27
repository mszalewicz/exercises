package main

import "fmt"

type User struct {
	name string
	id   string
}

func (user User) String() string {
	return fmt.Sprintf("name: %s, id: %s", user.name, user.id)
}

func main() {
	user := User{name: "test", id: "2342340-sdfs-234"}

	fmt.Println(user)
}
