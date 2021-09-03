package db

import (
	"context"
	"embed"
	"fmt"
	"io/fs"

	"github.com/jmoiron/sqlx"
	"github.com/sei-ri-10antz/todoist"
)

var _ Client = &client{}

//go:embed schema
var schemaFS embed.FS

type Options struct {
	Driver         string
	Source         string
	MaxIdleConns   int
	MaxOpenConns   int
	MigrateEnabled bool
}

type Option func(o *Options)

func Driver(s string) Option {
	return func(o *Options) {
		o.Driver = s
	}
}

func Source(s string) Option {
	return func(o *Options) {
		o.Source = s
	}
}

func MaxIdleConns(n int) Option {
	return func(o *Options) {
		o.MaxIdleConns = n
	}
}

func MaxOpenConns(n int) Option {
	return func(o *Options) {
		o.MaxOpenConns = n
	}
}

func Migrate() Option {
	return func(o *Options) {
		o.MigrateEnabled = true
	}
}

type client struct {
	opts Options
	db   *sqlx.DB
}

func NewClient(ctx context.Context, opts ...Option) (Client, error) {
	c := &client{
		opts: Options{
			Driver:       "mysql",
			Source:       "root:sa@tcp(127.0.0.1:3306)/sandbox?parseTime=true",
			MaxIdleConns: 10,
			MaxOpenConns: 20,
		},
	}

	// apply options
	for i := range opts {
		opts[i](&c.opts)
	}

	return c, c.open(ctx)
}

func (c *client) open(ctx context.Context) error {
	db, err := sqlx.ConnectContext(ctx, c.opts.Driver, c.opts.Source)
	if err != nil {
		return err
	}

	db.SetMaxIdleConns(c.opts.MaxIdleConns)
	db.SetMaxOpenConns(c.opts.MaxOpenConns)
	// ....
	c.db = db

	if !c.opts.MigrateEnabled {
		return nil
	}
	return c.migrate()
}

func (c *client) migrate() error {

	names, err := fs.Glob(schemaFS, fmt.Sprintf("schema/%s/*.sql", c.opts.Driver))
	if err != nil {
		return err
	}

	tx := c.db.MustBegin()
	defer tx.Rollback()

	for _, name := range names {
		buf, err := fs.ReadFile(schemaFS, name)
		if err != nil {
			return err
		}
		tx.Exec(string(buf))
	}

	return tx.Commit()
}

func (c *client) UsersStore() todoist.UsersStore {
	return &usersStore{c}
}

func (c *client) TasksStore() todoist.TasksStore {
	return &tasksStore{c}
}

func (c *client) Close() error {
	return c.db.Close()
}
