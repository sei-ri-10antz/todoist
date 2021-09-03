package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sei-ri-10antz/todoist"
	"github.com/sei-ri-10antz/todoist/http/packet"
)

func (s *Service) Task(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	task, err := s.Store.Tasks().Get(ctx, todoist.TaskQuery{ID: &id})
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, packet.NewTaskResponse(task))
}

func (s *Service) Tasks(c *gin.Context) {
	ctx := c.Request.Context()
	uid := c.Query("uid")
	tasks, err := s.Store.Tasks().All(ctx, todoist.TaskQuery{UserID: &uid})
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, packet.NewTasksResponse(tasks))
}

func (s *Service) CreateTask(c *gin.Context) {
	ctx := c.Request.Context()
	var payload packet.CreateTask
	if err := c.BindJSON(&payload); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// TODO: get user_id from access_token
	newTask := &todoist.Task{
		Name:    payload.Name,
		UserID:  "xxx",
		DueDate: payload.EndDate,
	}
	if err := s.Store.Tasks().Add(ctx, newTask); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, packet.NewSelfLinks("/users", newTask.ID))
}
