package main

import (
	"fmt"
	"net/http"
)

func main() {
	address := "127.0.0.1:7070"
	http.HandleFunc("/", MainPage)
	http.ListenAndServe(address, nil)
}

func MainPage(response http.ResponseWriter,reqest *http.Request,) {
	fmt.Fprintf(response, "main page")
}
