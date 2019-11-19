package get

import (
	"api/db/get"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

func Get(w http.ResponseWriter, r *http.Request) {
	rTasks := regexp.MustCompile("/todo-api/tasks/?")
	rTask := regexp.MustCompile("/todo-api/tasks/[0-9]+/?")
	rNum := regexp.MustCompile("[0-9]+")
	if rTask.MatchString(r.RequestURI) {
		num, _ := strconv.Atoi(rNum.FindAllStringSubmatch(r.RequestURI, -1)[0][0])
		json.NewEncoder(w).Encode(get.GetTask(num))
		fmt.Println("Get task ", num)
	} else if rTasks.MatchString(r.RequestURI) {
		json.NewEncoder(w).Encode(get.GetTasks())
		fmt.Println("Get tasks")
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
