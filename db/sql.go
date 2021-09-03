package db

import "fmt"

const (
	sqlAllTasks = iota
	sqlAllTasksByUserId
	sqlGetTaskById
	sqlAddTask
)

var sqls = []string{
	"SELECT * FROM tasks", // sqlAllTasks
	"SELECT * FROM tasks WHERE user_id=?", // sqlAllTasksByUserId
	"SELECT * FROM tasks WHERE id=?", // sqlGetTaskById
	"INSERT INTO tasks (id, name, user_id, status, due_date, created_at, updated_at) VALUES(?,?,?,?,?,?,?)", // sqlAddTask
}

func sqlError(code int, err error) error {
	return  fmt.Errorf("sql error: %s %v", sqls[code], err)
}