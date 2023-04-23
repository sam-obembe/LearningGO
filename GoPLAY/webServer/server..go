package webserver

import (
	"fmt"
	"net/http"
)

func RunWebServer() {
	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprint(writer, "hello from the api")
	})

	http.ListenAndServe(":3000", nil)
}
