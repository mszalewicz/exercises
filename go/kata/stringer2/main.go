package main

import "fmt"

type User struct {
	name    string
	surname string
	age     int
}

func (user User) String() string {
	return fmt.Sprintf("-----------------------------------\n\n  User\n\n    name:     %s\n    surname:  %s\n    age:      %d\n\n-----------------------------------", user.name, user.surname, user.age)
}

func main() {
	newUser := User{name: "John", surname: "Sirenwood", age: 45}
	newUser2 := User{name: "Jaina", surname: "Treeage", age: 100}

	fmt.Println(newUser)
	fmt.Println(newUser2)
}
