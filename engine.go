package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Message struct {
	Type string
	Data string
}

func NewEngine() *gin.Engine {

	router := gin.New()

	//in Body type: earnings approved rejected
	router.POST("/scheduler", func(c *gin.Context) {
		service := &Message{}
		if err := c.BindJSON(service); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		Proxy1(service.Type)
	})

	router.POST("/redis", func(c *gin.Context) {})
	router.POST("/postres", func(c *gin.Context) {})

	router.GET("/redis", func(c *gin.Context) {})
	router.GET("/postres", func(c *gin.Context) {})

	return router
}
