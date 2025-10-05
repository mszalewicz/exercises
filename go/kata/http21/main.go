package main

import (
	"fmt"
	"net/http"
)

func main() {
	address := "127.0.0.1:8080"
	http.HandleFunc("/", response)
	http.ListenAndServe(address, nil)
}

func response(writter http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writter, "test")
}