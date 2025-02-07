package main

import "fmt"

type User struct {
	name string
	ucc  string
}

func (user User) String() string {
	return fmt.Sprintf("name: %s\nucc:  %s", user.name, user.ucc)
}

func main() {
	u := User{name: "Maria", ucc: "283-EU-0asm5768"}
	fmt.Println(u)
}
