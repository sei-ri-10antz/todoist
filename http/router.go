package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(service Service) http.Handler {
	r := gin.Default()

	r.GET("/ping", service.Ping)

	r.GET("/tasks", service.Tasks)
	r.POST("/tasks", service.CreateTask)
	r.GET("/tasks/:id", service.Task)

	return r
}
