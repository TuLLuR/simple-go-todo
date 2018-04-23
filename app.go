package main

import (
	"SinglePageApp/handlers"
	"database/sql"

	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := initDB("storage.db")
	migrate(db)
	e := echo.New()

	e.Static("/", "public")

	e.GET("/tasks", handlers.GetTasks(db))
	e.POST("/tasks", handlers.PostTask(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))
	e.Logger.Fatal(e.Start(":1012"))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db = nil")
	}
	return db
}

func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL
);`

	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}

}
