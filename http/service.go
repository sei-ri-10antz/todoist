package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Service struct {
	Store DataStore
}

func (s *Service) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}
