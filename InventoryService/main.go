package main

import (
	"fmt"
	"inventoryService/handlers"
	"net/http"
	"time"
)

func middlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before handler; middleware start")
		start := time.Now()
		handler.ServeHTTP(w, r)

		fmt.Printf("middleware finishedd; %s", time.Since(start))
	})
}

func main() {

	fmt.Printf("Listening on port %v", 5000)

	productListHandler := http.HandlerFunc(handlers.ProductsHandler)
	productItemHandler := http.HandlerFunc(handlers.ProductHandler)
	http.Handle("/foo", &handlers.FooHandler{Message: "foo called"})
	http.HandleFunc("/bar", handlers.BarHandler)
	http.Handle("/products", middlewareHandler(productListHandler))
	http.Handle("/product/", middlewareHandler(productItemHandler))
	http.ListenAndServe(":5000", nil)

}
