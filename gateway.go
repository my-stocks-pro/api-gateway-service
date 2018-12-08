package main

import (
	"github.com/gin-gonic/gin"
)

type Gateway struct {
	server *gin.Engine
	proxy  Proxy
	consul ConsulType
}

func NewGateway() Gateway {
	return Gateway{
		server: NewEngine(),
		proxy:  NewProxy(),
		consul: NewConsul(),
	}
}