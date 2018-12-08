package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func (g Gateway) HandleScheduler(c *gin.Context) {
	service := &Message{}
	if err := c.BindJSON(service); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	servicePath := g.consul.DiscoveryService(service.Type)
	fmt.Println(servicePath)

	if err := g.proxy.Do(servicePath); err != nil {
		fmt.Println(err)
	}

}
