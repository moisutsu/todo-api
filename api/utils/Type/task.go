package Type

type Task struct {
	Date       string `json:"date"`
	Body       string `json:"body"`
	IsFinished bool   `json:"isFinished"`
}

type Tasks []Task
