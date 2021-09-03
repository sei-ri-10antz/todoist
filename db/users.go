package db

import (
	"context"

	"github.com/sei-ri-10antz/todoist"
)

type usersStore struct{
	*client
}

func (s *usersStore) Add(ctx context.Context, user *todoist.User) error {
	panic("implement me")
}

func (s *usersStore) All(ctx context.Context) ([]*todoist.User, error) {
	panic("implement me")
}

func (s *usersStore) Get(ctx context.Context, query todoist.UserQuery) (*todoist.User, error) {
	panic("implement me")
}
