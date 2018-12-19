package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/my-stocks-pro/api-gateway-service/infrastructure"
)

type Health interface {
	Handle(c *gin.Context)
}

type TypeHealth struct {
	config infrastructure.Config
}

func NewHealth(config infrastructure.Config) TypeHealth {
	return TypeHealth{
		config: config,
	}
}

func (g TypeHealth) Handle(c *gin.Context) {
	c.JSON(200, gin.H{
		"health": "True",
	})

	//TODO get all services paths from Consul and get Health endpoints
	//TODO return struct of services with current status
}
