package handlers

import "net/http"

type FooHandler struct {
	Message string
}

func (handle *FooHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte(handle.Message))
}
