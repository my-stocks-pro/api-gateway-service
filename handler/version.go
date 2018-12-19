package handler

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/my-stocks-pro/api-gateway-service/infrastructure"
)

type Version interface {
	Handle(c *gin.Context)
}

type TypeVersion struct {
	config infrastructure.Config
}

func NewVersion(config infrastructure.Config) TypeVersion {
	return TypeVersion{
		config: config,
	}
}

func (v TypeVersion) Handle(c *gin.Context) {
	c.JSON(200, gin.H{
		"startTime": v.config.StartDate,
		"currDate":  time.Now().Format("2006-01-02 15:04"),
		"version":   "1.0",
		"service":   v.config.Name,
	})
}
