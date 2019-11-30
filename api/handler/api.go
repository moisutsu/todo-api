package handler

import (
	"api/handler/delete"
	"api/handler/get"
	"api/handler/post"
	"net/http"
)

func Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		get.Get(w, r)
	case http.MethodPost:
		post.Post(w, r)
	case http.MethodDelete:
		delete.Delete(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}
