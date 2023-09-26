package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/health", HealthCheck)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func HealthCheck(w http.ResponseWriter, req *http.Request) {
	currentTime := time.Now()
	io.WriteString(w, currentTime.String())
}
