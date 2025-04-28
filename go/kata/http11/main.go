package main

import (
	"fmt"
	"net/http"
)

func main() {
	address := "127.0.0.1:1313"

	http.HandleFunc("/", mainPage)
	http.ListenAndServe(address, nil)
}

func mainPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "main page")
}