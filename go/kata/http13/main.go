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

func mainPage(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprint(responseWriter, "main page")
}
