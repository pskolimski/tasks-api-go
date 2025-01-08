package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDb() {
	cfg := mysql.Config{
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		User:                 "root",
		Passwd:               "admin",
		DBName:               "todolist",
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

func getTasksFromDB() ([]Task, error) {
	result, err := db.Query("SELECT * FROM tasks")

	if err != nil {
		return nil, err
	}

	defer func(result *sql.Rows) {
		err := result.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(result)

	var tasks []Task

	for result.Next() {
		var task Task
		if err := result.Scan(&task.Id, &task.Title, &task.IsDone); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func addTaskToDB(task *Task) error {
	_, err := db.Exec("INSERT INTO tasks VALUES (?, ?, ?)", task.Id, task.Title, task.IsDone)

	if err != nil {
		return err
	}

	return nil
}

func updateTaskInDB(task *Task) error {
	_, err := db.Exec("UPDATE tasks SET title = ?, isDone = ? WHERE uuid = ?", task.Title, task.IsDone, task.Id)

	if err != nil {
		return err
	}

	return nil
}

func getTaskFromDB(id string) (*Task, error) {
	row := db.QueryRow("SELECT * FROM tasks WHERE uuid = ?", id)

	var task Task
	if err := row.Scan(&task.Id, &task.Title, &task.IsDone); err != nil {
		return nil, err
	}

	return &task, nil
}
