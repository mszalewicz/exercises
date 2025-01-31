package main

import (
	"fmt"
)

type Book struct {
	isbn  string
	title string
	price float64
}

func (book Book) String() string {
	return fmt.Sprintf("isbn: %s || title: %s || price: %f", book.isbn, book.title, book.price)
}

func main() {
	book := Book{isbn: "ISBN-2025-123098-2342", title: "U2, BUT not me", price: 345.57}

	fmt.Println(book)
}
