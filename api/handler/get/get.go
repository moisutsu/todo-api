package get

import (
	"net/http"
	"regexp"
	"encoding/json"
	"strconv"
	"fmt"
	"api/utils/Type"
)

func Get(w http.ResponseWriter, r *http.Request) {
	rTasks := regexp.MustCompile("/todo-api/tasks/?")
	rTask := regexp.MustCompile("/todo-api/tasks/[0-9]+/?")
	rNum := regexp.MustCompile("[0-9]+")
	if rTask.MatchString(r.RequestURI) {
		num, _ := strconv.Atoi(rNum.FindAllStringSubmatch(r.RequestURI, -1)[0][0])
		json.NewEncoder(w).Encode(getTask(num))
		fmt.Println("Get task ", num)
	} else if rTasks.MatchString(r.RequestURI) {
		json.NewEncoder(w).Encode(getTasks())
		fmt.Println("Get tasks")
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func getTasks() Type.Tasks {
	tasks := generateTasks()
	return tasks
}

func getTask(index int) Type.Task {
	tasks := generateTasks()
	if index < 0 || index > 4 {
		return tasks[0]
	}
	return tasks[index]
}

func generateTasks() Type.Tasks {
	tasks := Type.Tasks{}
	for i := 0; i < 5; i++ {
		body := fmt.Sprintf("Body_%d", i)
		date := fmt.Sprintf("2019-10-1%d", i)
		tasks = append(
			tasks,
			Type.Task{Date: date, Body: body, IsFinished: false},
		)
	}
	return tasks
}
