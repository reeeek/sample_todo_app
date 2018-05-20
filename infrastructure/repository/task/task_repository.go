package task

import (
	"database/sql"
	"log"
	"time"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const TableName = "tasks"

type TaskRepository struct {
}

type Task struct {
	ID        int       `json:"id"`
	Status    int       `json:"status"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{}
}

func (r *TaskRepository) FindAll() ([]Task, error) {
	db, err := sql.Open("mysql", "root:@/sample_todo?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM tasks") //
	if err != nil {
		panic(err.Error())
	}
	tasks := make([]Task, 0)
	for rows.Next() {
		t := Task{}
		if err := rows.Scan(&t.ID, &t.Status, &t.Title, &t.Body, &t.CreatedAt, &t.UpdatedAt); err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return tasks, nil
}

func (r *TaskRepository) Store(t *Task) error {
	db, err := sql.Open("mysql", "root:@/sample_todo")
	if err != nil {
		panic(err.Error())
	}
	stmt, err := db.Prepare(fmt.Sprintf("INSERT INTO %s (status, title, body, created_at, updated_at) VALUES (?, ?, ?, ?,?)", TableName))
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(t.Status, t.Title, t.Body, t.CreatedAt, t.UpdatedAt)
	if err != nil {
		panic(err.Error())
	}
	return nil
}

func (r *TaskRepository) Update(id int, t *Task) error {
	db, err := sql.Open("mysql", "root:@/sample_todo")
	if err != nil {
		panic(err.Error())
	}

	stmt, err := db.Prepare(fmt.Sprintf("UPDATE %s SET %s WHERE id=?", TableName, UpdateState))
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(t.Status, t.Title, t.Body, t.CreatedAt, t.UpdatedAt)
	if err != nil {
		panic(err.Error())
	}
	return nil
}
