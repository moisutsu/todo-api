package Type

type Task struct {
	Date       string `json:"date"`
	Body       string `json:"body"`
	IsFinished bool   `json:"is_finished"`
}

type Tasks []Task
