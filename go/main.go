package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/todo-api", ApiHandler)

	http.ListenAndServe(":8080", nil)
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("task")
}
