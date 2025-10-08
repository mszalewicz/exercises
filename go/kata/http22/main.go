package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	address := "127.0.0.1:4444"

	http.HandleFunc("/", mainpage)

	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatalln("Could not start http server:", err)
	}
}

func mainpage(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "main page")
}