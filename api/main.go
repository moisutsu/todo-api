package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

type Task struct {
	Date       string `json:"date"`
	Body       string `json:"body"`
	IsFinished bool   `json:"isFinished"`
}

type Tasks []Task

func main() {
	http.HandleFunc("/todo-api/", apiHandler)

	http.ListenAndServe(":8080", nil)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		getHandler(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
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

func generateTasks() Tasks {
	tasks := Tasks{}
	for i := 0; i < 5; i++ {
		body := fmt.Sprintf("Body_%d", i)
		date := fmt.Sprintf("2019-10-1%d", i)
		tasks = append(
			tasks,
			Task{Date: date, Body: body, IsFinished: false},
		)
	}
	return tasks
}

func getTasks() Tasks {
	tasks := generateTasks()
	return tasks
}

func getTask(index int) Task {
	tasks := generateTasks()
	if index < 0 || index > 4 {
		return tasks[0]
	}
	return tasks[index]
}
