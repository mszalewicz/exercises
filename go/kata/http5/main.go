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

func mainpage(reponseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprint(reponseWriter, "main page")
}