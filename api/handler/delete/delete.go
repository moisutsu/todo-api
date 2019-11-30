package delete

import (
	"api/db/delete"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	rTask := regexp.MustCompile("/todo-api/tasks/[0-9]+/?")
	rNum := regexp.MustCompile("[0-9]+")
	if rTask.MatchString(r.RequestURI) {
		num, _ := strconv.Atoi(rNum.FindAllStringSubmatch(r.RequestURI, -1)[0][0])
		delete.DeleteTask(num)
		fmt.Println("Delete task ", num)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
