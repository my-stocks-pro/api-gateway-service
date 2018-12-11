package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"errors"
	"fmt"
)

type HistoryMessage struct {
	Service  string
	FromData string
	ToData   string
}

func (g TypeGateway) HistoryHandler(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, errors.New("MethodNotAllowed"))
		return
	}

	msg := &HistoryMessage{}
	if err := c.BindJSON(msg); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	servicePath, err := g.consul.DiscoveryService(c.Param(ServiceKey))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(servicePath)

	body, err := g.proxy.Request(c.Request.Method, servicePath, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(body)
}
