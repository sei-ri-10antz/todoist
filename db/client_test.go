package db_test

import (
	"context"
	"testing"

	"github.com/sei-ri-10antz/todoist/db"

	_ "github.com/mattn/go-sqlite3"
)

func NewTestClient(t testing.TB) (db.Client, error) {
	t.Helper()

	return db.NewClient(context.Background(),
		db.Driver("sqlite3"),
		db.Source(":memory:"),
		db.MaxOpenConns(1),
		db.MaxIdleConns(1),
		db.Migrate(),
	)
}

func TestClient(t *testing.T) {
	if _, err := NewTestClient(t); err != nil {
		t.Error(err)
	}
}

func TestClientWithFile(t *testing.T) {
	_, err := db.NewClient(context.Background(),
		db.Driver("sqlite3"),
		db.Source("./sandbox.db"),
		db.MaxOpenConns(1),
		db.MaxIdleConns(1),
		db.Migrate(),
	)
	if err != nil {
		t.Error(err)
	}
}