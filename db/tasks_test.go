package db_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/sei-ri-10antz/todoist"
)



func Test_tasksStore_Get(t *testing.T) {
	fakeID := "xxx"

	type args struct {
		ctx   context.Context
		query todoist.TaskQuery
	}
	tests := []struct {
		name    string
		args    args
		want    *todoist.Task
		wantErr bool
		addFirst bool
	}{
		{
			name: "not found",
			args: args{
				ctx: context.Background(),
				query: todoist.TaskQuery{
					ID: &fakeID,
				},
			},
			wantErr: true,
		},
		{
			name: "get a user",
			args: args{
				ctx: context.Background(),
				query: todoist.TaskQuery{
					ID: &fakeID,
				},
			},
			want: &todoist.Task{
				ID:        "xxx",
				Name:      "xxx",
				Status:    1,
				UserID:    "xxx",
				DueDate:   "2020-10-22",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			wantErr: false,
			addFirst: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewTestClient(t)
			if err != nil {
				// t.Fatal(err)
			}
			defer client.Close()

			s := client.TasksStore()

			if tt.addFirst {
				s.Add(tt.args.ctx, tt.want)
			}

			got, err := s.Get(tt.args.ctx, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Get() -got/+want %v", diff)
			}
		})
	}
}
