package db

import (
	"context"
	"fmt"

	"github.com/sei-ri-10antz/todoist"
)

type tasksStore struct{
	*client
}

func (s *tasksStore) Add(ctx context.Context, task *todoist.Task) error {
	tx := s.db.MustBegin()
	defer tx.Rollback()

	_, err := tx.Exec(sqls[sqlAddTask], task.ID, task.Name, task.UserID, task.Status, task.DueDate, task.CreatedAt, task.UpdatedAt)
	if err != nil {
		return sqlError(sqlAddTask, err)

	}
	return tx.Commit()
}

func (s *tasksStore) All(ctx context.Context, query todoist.TaskQuery) ([]*todoist.Task, error) {
	var tasks []*todoist.Task
	if err := s.db.Select(&tasks, sqls[sqlAllTasksByUserId], query.UserID); err != nil {
		return nil, sqlError(sqlAllTasksByUserId, err)
	}
	return tasks, nil
}

func (s *tasksStore) Get(ctx context.Context, query todoist.TaskQuery) (*todoist.Task, error) {
	if query.ID == nil {
		return nil, fmt.Errorf("invalid ID in TaskQuery")
	}
	var task todoist.Task
	if err := s.db.Get(&task, sqls[sqlGetTaskById], *query.ID); err != nil {
		return nil, sqlError(sqlGetTaskById, err)
	}
	return &task, nil
}
