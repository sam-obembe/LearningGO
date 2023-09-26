package main

import (
	"fmt"
	"net/http"
)

func main() {
	originalHandler := http.HandlerFunc(handle)

	http.Handle("/", middleware(originalHandler))
	http.ListenAndServe(":8000", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing main handler")
	w.Write([]byte("Ok"))
}

func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware before req")
		handler.ServeHTTP(w, r)
		fmt.Println("Executing middleware after req")
	})
}
