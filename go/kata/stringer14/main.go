package main

import "fmt"

type User struct {
	name     string
	id       string
	category string
}

func (user User) String() string {
	return fmt.Sprintf("Username: %s | ID: %s | Category: %s", user.name, user.id, user.category)

}

func main() {
	user := User{name: "Bilbo", id: "xcvb098-a1sfd-098xcv0", category: "Client"}

	fmt.Println(user)
}
