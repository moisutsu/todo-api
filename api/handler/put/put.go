package put

import (
	"api/db/put"
	"api/utils/Type"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

func Put(w http.ResponseWriter, r *http.Request) {
	rTask := regexp.MustCompile("/todo-api/tasks/[0-9]+/?")
	rNum := regexp.MustCompile("[0-9]+")
	if rTask.MatchString(r.RequestURI) {
		if (r.Header.Get("Content-Type") != "application/json;charset=UTF-8") && (r.Header.Get("Content-Type") != "application/json") {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		body, _ := ioutil.ReadAll(r.Body)
		var task Type.Task
		json.Unmarshal(body, &task)

		num, _ := strconv.Atoi(rNum.FindAllStringSubmatch(r.RequestURI, -1)[0][0])

		put.PutTask(num, task)
		fmt.Println("Put index ", num, " task ", task)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
