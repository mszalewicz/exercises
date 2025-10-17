package main

import (
	"fmt"
	"net/http"
)

func main() {
	address := "127.0.0.1:6060"
	
	http.HandleFunc("/", MainPage)
	http.ListenAndServe(address, nil)
}

func MainPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "main page")
}
