package main

import (
	"github.com/gin-gonic/gin"
)

type Message struct {
	Type string
	Data string
}

func NewEngine() *gin.Engine {
	return gin.New()
}

func (g Gateway) Routing() {
	g.server.POST("/scheduler", g.HandleScheduler)

	g.server.POST("/redis", func(c *gin.Context) {})
	g.server.POST("/postres", func(c *gin.Context) {})

	g.server.GET("/redis", func(c *gin.Context) {})
	g.server.GET("/postres", func(c *gin.Context) {})
}
