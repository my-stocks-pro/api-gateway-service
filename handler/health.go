package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"errors"
)

func (g TypeGateway) HealthHandle(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, errors.New("MethodNotAllowed"))
		return
	}
	//TODO get all services paths from Consul and get Health endpoints
	//TODO return struct of services with current status
}
