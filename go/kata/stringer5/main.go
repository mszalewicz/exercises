package main

import "fmt"

type Product struct {
	name    string
	iisd    string
	quality int
}

func (product Product) String() string {
	return fmt.Sprintf("name: %s, iisd: %s, quality: %d", product.name, product.iisd, product.quality)
}

func main() {
	product := Product{name: "vacum", iisd: "19810-TUF-1393", quality: 8}

	fmt.Println(product)
}
