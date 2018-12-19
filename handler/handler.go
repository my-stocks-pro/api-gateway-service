package handler

import (
	"github.com/gin-gonic/gin"
)

const (
	service = "service"
)

type Handler interface {
	Handle(c *gin.Context)
}

type CommonMessage struct {
	Service string
	Data    []byte
}
