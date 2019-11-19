package get

import (
	"api/utils/Type"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func GetTask(index int) Type.Task {
	db, err := sql.Open("mysql", "root:@/todo")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	var date, body, isFinished string
	err = db.QueryRow("SELECT * FROM tasks LIMIT 1 OFFSET ?;", index).Scan(&date, &body, &isFinished)

	var task Type.Task
	task.Date = date
	task.Body = body
	task.IsFinished, _ = strconv.ParseBool(isFinished)
	return task

}
