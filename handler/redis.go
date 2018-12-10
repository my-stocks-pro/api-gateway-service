package handler

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (g TypeGateway) HandleRedis(c *gin.Context) {

	msg := Message{}
	if err := c.BindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	servicePath := g.consul.DiscoveryService(msg.Service)
	fmt.Println(servicePath)

	switch c.Request.Method {
	case http.MethodGet:
		fmt.Println("GET")
	case http.MethodPost:
		fmt.Println("POST")
	default:
		fmt.Println("default")
	}
}
