package main

import (
	"api/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/todo-api/", handler.Api)
	http.ListenAndServe(":8080", nil)
}
