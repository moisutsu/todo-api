package put

import (
	"api/db/get"
	"api/utils/Type"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func PutTask(index int, task Type.Task) {
	db, err := sql.Open("mysql", "root:@/todo")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	getTask := get.GetTask(index)

	_, err = db.Exec("UPDATE tasks SET date=?, body=?, is_finished=? WHERE date=? and body=? and is_finished=?;", task.Date, task.Body, task.IsFinished, getTask.Date, getTask.Body, getTask.IsFinished)
	if err != nil {
		fmt.Println(err)
	}
}
