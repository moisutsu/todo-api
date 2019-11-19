package get

import (
	"api/utils/Type"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func GetTasks() Type.Tasks {
	db, err := sql.Open("mysql", "root:@/todo")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tasks;")
	if err != nil {
		fmt.Println(err)
	}

	columns, err := rows.Columns() // カラム名を取得
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	var tasks Type.Tasks
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		var task Type.Task
		for i, col := range values {
			value := string(col)
			switch columns[i] {
			case "date":
				task.Date = value
			case "body":
				task.Body = value
			case "is_finished":
				task.IsFinished, _ = strconv.ParseBool(value)
			}
		}
		tasks = append(tasks, task)
	}
	return tasks
}
