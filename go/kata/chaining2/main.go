package main

import "fmt"

type Operation struct {
	currency        string
	currency_amount float64
	product         string
	product_amount  int
}

func NewOperation() *Operation {
	return &Operation{}
}

func (operation *Operation) WithCurrency(currency string) *Operation {
	operation.currency = currency
	return operation
}

func (operation *Operation) CurrencyAmount(amount float64) *Operation {
	operation.currency_amount = amount
	return operation
}

func (operation *Operation) WithProduct(product string) *Operation {
	operation.product = product
	return operation
}

func (operation *Operation) ProductAmount(amount int) *Operation {
	operation.product_amount = amount
	return operation
}

func main() {

	operation := NewOperation().
		WithCurrency("EUR").
		CurrencyAmount(3645.23).
		WithProduct("Subscription").
		ProductAmount(1)

	fmt.Println(operation)
}
