package todoist

import (
	"context"
	"time"
)

type User struct {
	ID        string
	Identity  string
	Passwd    string
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type UserQuery struct {
	ID       *string
	Identity *string
}

type UsersStore interface {
	// create a new user
	Add(context.Context, *User) error
	// gets all users data
	All(context.Context) ([]*User, error)
	// gets a user with query
	Get(context.Context, UserQuery) (*User, error)
}

type Task struct {
	ID        string
	Name      string
	Status    int
	UserID    string `db:"user_id"`
	DueDate   string `db:"due_date"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type TaskQuery struct {
	ID     *string
	UserID *string
}

type TasksStore interface {
	// create a new task
	Add(context.Context, *Task) error
	// gets all tasks data
	All(context.Context, TaskQuery) ([]*Task, error)
	// gets a task with query
	Get(context.Context, TaskQuery) (*Task, error)
}
