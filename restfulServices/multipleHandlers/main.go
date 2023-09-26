package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.ExpFloat64())
	})

	mux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Intn(1000))
	})

	log.Fatal(http.ListenAndServe(":8000", mux))
}
