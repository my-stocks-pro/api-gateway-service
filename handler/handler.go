package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Handle(c *gin.Context)
}

type CommonMessage struct {
	Service string
	Data    []byte
}
