package db

import "github.com/sei-ri-10antz/todoist"

type Client interface {
	UsersStore() todoist.UsersStore
	TasksStore() todoist.TasksStore
	Close() error
}
