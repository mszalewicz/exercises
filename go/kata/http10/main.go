package main

import (
	"fmt"
	"net/http"
)

func main() {
	address := "127.0.0.1:3000"

	http.HandleFunc("/", home)
	http.ListenAndServe(address, nil)

}

func home(writter http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writter, "this is home page")
}