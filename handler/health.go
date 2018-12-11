package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"errors"
)

type Health interface {
	Handle(c *gin.Context)
}

type TypeHealth struct{}

func NewHealth() TypeHealth {
	return TypeHealth{}
}

func (g TypeHealth) Handle(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, errors.New("MethodNotAllowed"))
		return
	}

	c.JSON(200, gin.H{
		"health":  "True",
	})

	//TODO get all services paths from Consul and get Health endpoints
	//TODO return struct of services with current status
}
