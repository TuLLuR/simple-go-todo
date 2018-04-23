package models

import "database/sql"

type Task struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type TaskCollection struct {
	Task []Task `json:"items"`
}

func GetTasks(db *sql.DB) TaskCollection {
	sql := "SELECT * FROM tasks"
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	result := TaskCollection{}

	for rows.Next() {
		task := Task{}
		err2 := rows.Scan(&task.ID, &task.Name)

		if err2 != nil {
			panic(err2)
		}

		result.Task = append(result.Task, task)

	}

	return result
}

func PostTasks(db *sql.DB, name string) (int64, error)  {
	sql := "INSERT INTO tasks (name) VALUES(?)"

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err2 := stmt.Exec(name)

	if err2 != nil {
		panic(err)
	}
	return result.LastInsertId()
}

func DeleteTask(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM tasks WHERE id = ?"

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	result, err2 := stmt.Exec(id)

	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}