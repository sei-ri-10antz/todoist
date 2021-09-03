package mocks

import (
	"context"

	"github.com/sei-ri-10antz/todoist"
)

var _ todoist.TasksStore = &TasksStore{}

// TasksStore implements todoist.TasksStore
type TasksStore struct {
	AddFn func(ctx context.Context, t *todoist.Task) error
	AllFn func(ctx context.Context, q todoist.TaskQuery) ([]*todoist.Task, error)
	GetFn func(ctx context.Context, q todoist.TaskQuery) (*todoist.Task, error)
}

// create a new task
func (s *TasksStore) Add(ctx context.Context, t *todoist.Task) error {
	return s.AddFn(ctx, t)
}

// gets all tasks data
func (s *TasksStore) All(ctx context.Context, q todoist.TaskQuery) ([]*todoist.Task, error) {
	return s.AllFn(ctx, q)
}

// gets a task with query
func (s *TasksStore) Get(ctx context.Context, q todoist.TaskQuery) (*todoist.Task, error) {
	return s.GetFn(ctx, q)
}
