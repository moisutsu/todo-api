package handler

import (
	"api/handler/get"
	"fmt"
	"net/http"
)

func Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		get.Get(w, r)
	case "POST":
		fmt.Println("POST")
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}
