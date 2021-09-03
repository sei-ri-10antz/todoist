package packet

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/sei-ri-10antz/todoist"
	"google.golang.org/protobuf/proto"
)

//go:generate protoc --go_out=. --go_opt paths=source_relative packet.proto

func NewSelfLinks(args ...string) *SelfLinks {
	return &SelfLinks{
		Self: strings.Join(args, "/"),
	}
}

func NewTaskResponse(arg *todoist.Task) *TaskResponse {
	return &TaskResponse{
		Id:            arg.ID,
		Name:          arg.Name,
		DueDate:       arg.DueDate,
		LastUpdatedAt: arg.UpdatedAt.Format(time.RFC3339),
		Links:         NewSelfLinks("/tasks", arg.ID),
	}
}

func NewTasksResponse(args []*todoist.Task) *TasksResponse {
	tasks := make([]*TaskResponse, len(args))
	for i := range args {
		tasks[i] = NewTaskResponse(args[i])
	}
	return &TasksResponse{
		Tasks: tasks,
		Links: NewSelfLinks("/tasks"),
	}
}

func MashalJSON(msg proto.Message) string {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return string(b)
}
