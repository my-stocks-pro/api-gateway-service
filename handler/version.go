package handler

import (
	"net/http"
	"time"
	"errors"
	"github.com/gin-gonic/gin"
)

type Version interface {
	Handle(c *gin.Context)
}

type TypeVersion struct{}

func NewVersion() TypeVersion {
	return TypeVersion{}
}

func (v TypeVersion) Handle(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, errors.New("MethodNotAllowed"))
		return
	}

	c.JSON(200, gin.H{
		//"startTime": g.config.StartDate,
		"currDate": time.Now().Format("2006-01-02 15:04"),
		"version":  "1.0",
		//"service":   g.config.Name,
	})
}
