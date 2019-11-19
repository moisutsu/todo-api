package main

import (
	"api/db/get"
	"api/handler"
	"net/http"
)

func main() {
	get.GetTasks()
	http.HandleFunc("/todo-api/", handler.Api)

	http.ListenAndServe(":8080", nil)
}
