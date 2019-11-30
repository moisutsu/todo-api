package delete

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func DeleteTask(index int) {
	db, err := sql.Open("mysql", "root:@/todo")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	var date, body, isFinished string
	err = db.QueryRow("SELECT * FROM tasks LIMIT 1 OFFSET ?;", index).Scan(&date, &body, &isFinished)

	_, err = db.Exec("DELETE FROM tasks WHERE date = ? AND body = ?;", date, body)
	if err != nil {
		fmt.Println(err)
	}

}
