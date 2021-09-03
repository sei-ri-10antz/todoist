package http

import "github.com/sei-ri-10antz/todoist"

type DataStore interface {
	Users() todoist.UsersStore
	Tasks() todoist.TasksStore
}

type Store struct {
	UsersStore todoist.UsersStore
	TasksStore todoist.TasksStore
}

func (s *Store) Users() todoist.UsersStore {
	return s.UsersStore
}

func (s *Store) Tasks() todoist.TasksStore {
	return s.TasksStore
}
