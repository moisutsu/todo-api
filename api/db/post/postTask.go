package post

import (
	"api/utils/Type"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func PostTask(task Type.Task) {
	db, err := sql.Open("mysql", "root:@/todo")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO tasks (date, body, is_finished) values (?, ?, ?);", task.Date, task.Body, task.IsFinished)
	if err != nil {
		fmt.Println(err)
	}
}
