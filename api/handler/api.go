package handler

import (
	"api/handler/get"
	"api/handler/post"
	"net/http"
)

func Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		get.Get(w, r)
	case "POST":
		post.Post(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}
