package http_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/sei-ri-10antz/todoist"
	todohttp "github.com/sei-ri-10antz/todoist/http"
	"github.com/sei-ri-10antz/todoist/http/packet"
	"github.com/sei-ri-10antz/todoist/mocks"
)

var (
	TaskTestData = []*todoist.Task{
		{
			ID:      "xx",
			Name:    "xx",
			Status:  0,
			UserID:  "xxx",
			DueDate: "2020-10-01",
		},
		{
			ID:      "xy",
			Name:    "xy",
			Status:  1,
			UserID:  "xxx",
			DueDate: "2020-10-01",
		},
	}
)

func TestService_Task(t *testing.T) {
	type fields struct {
		Store todohttp.DataStore
	}
	type args struct {
		c  TestContext
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   TestResponse
	}{
		{
			name: "ok",
			fields: fields{
				Store: &todohttp.Store{
					TasksStore: &mocks.TasksStore{
						GetFn: func(ctx context.Context, q todoist.TaskQuery) (*todoist.Task, error) {
							for i := range TaskTestData {
								if TaskTestData[i].ID == *q.ID {
									return TaskTestData[i], nil
								}
							}
							return nil, fmt.Errorf("not found")
						},
					},
				},
			},
			args: args{
				c: TestContext{
					w: httptest.NewRecorder(),
					r: httptest.NewRequest("GET", "http://any.url", nil),
				},
				id: TaskTestData[0].ID,
			},
			want: TestResponse{
				StatusCode: http.StatusOK,
				Body:       packet.MashalJSON(packet.NewTaskResponse(TaskTestData[0])),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &todohttp.Service{
				Store: tt.fields.Store,
			}

			s.Task(tt.args.c.Context(gin.Param{Key: "id", Value: tt.args.id}))

			resp := tt.args.c.w.Result()
			body, _ := ioutil.ReadAll(tt.args.c.w.Body)

			if diff := cmp.Diff(resp.StatusCode, tt.want.StatusCode); diff != "" {
				t.Errorf("StatusCode: -got/+want\n%s", diff)
			}
			if diff := cmp.Diff(string(body), tt.want.Body); diff != "" {
				t.Errorf("Body: -got/+want\n%s", diff)
			}
		})
	}
}
