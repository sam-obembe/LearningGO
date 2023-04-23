package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe("localhost:3000", nil)
}

func Handler(response http.ResponseWriter, request *http.Request) {
	file, _ := os.Open("./menu.csv")

	io.Copy(response, file)
}
