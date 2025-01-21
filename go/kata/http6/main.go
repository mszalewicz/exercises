package main

import (
	"fmt"
	"net/http"
)

func main() {
	address := "127.0.0.1:8080"

	http.HandleFunc("/", mainpage)
	http.ListenAndServe(address, nil)
}

func mainpage(responeWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(responeWriter, "this is main page")
}
