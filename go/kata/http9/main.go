package main

import (
	"fmt"
	"net/http"
)

func main() {
	address := "127.0.0.1:8080"

	http.HandleFunc("/", mainPage)
	http.ListenAndServe(address, nil)
}

func mainPage(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "main page")
}