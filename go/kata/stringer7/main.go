package main

import "fmt"

type User struct {
	name string
	id   string
}

func (user User) String() string {
	return fmt.Sprintf("User - name = %s, id = %s", user.name, user.id)
}

func main() {
	user := User{name: "Ivan", id: "f8ero-3jxmek-sdof"}

	fmt.Println(user)
}
