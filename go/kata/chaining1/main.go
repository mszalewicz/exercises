package main

import (
	"errors"
	"fmt"
	"log"
)

type User struct {
	err      error
	username string
	address  string
	price    float64
	total    float64
	amount   int
}

func NewUser() *User {
	return &User{}
}

func (user *User) WithUsername(username string) *User {
	if len(username) < 1 || len(username) > 25 {
		user.err = errors.New("Username must be between 1 and 25 characters long")
	} else {
		user.username = username
	}

	return user
}

func (user *User) WithAddress(address string) *User {
	if user.err != nil {
		return user
	}

	if len(address) < 1 || len(address) > 25 {
		user.err = errors.New("Username must be between 1 and 25 characters long")
	} else {
		user.address = address
	}

	return user
}

func (user *User) WithAmount(amount int) *User {
	if user.err != nil {
		return user
	}

	if amount < 1 {
		user.err = errors.New("Amount can't be lower than 1")
	} else {
		user.amount = amount
	}

	return user
}

func (user *User) WithPrice(price float64) *User {
	if user.err != nil {
		return user
	}

	if price <= 0 {
		user.err = errors.New("Price can't be zero or lower")
	} else {
		user.price = price
	}

	return user
}

func (user *User) Calculate() (*User, error) {
	if user.err != nil {
		return nil, user.err
	}

	user.total = float64(user.amount) * user.price
	return user, nil
}

func main() {
	user, err := NewUser().
		WithUsername("John").
		WithAddress("Nowhere Strasse").
		WithAmount(45).
		WithPrice(756.213).
		Calculate()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user.total)

	user2, err := NewUser().
		WithUsername("JoanOfArcJoanOfArcJoanOfArcJoanOfArc").
		WithAddress("Nowhere Strasse").
		WithAmount(45).
		WithPrice(756.213).
		Calculate()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user2.total)
}
