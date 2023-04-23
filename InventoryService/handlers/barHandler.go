package handlers

import "net/http"

func BarHandler(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte("bar called"))
}
