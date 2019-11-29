package post

import (
	"api/db/post"
	"api/utils/Type"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	body, _ := ioutil.ReadAll(r.Body)
	var task Type.Task
	json.Unmarshal(body, &task)
	post.PostTask(task)
	fmt.Println("POST ", task)
}
