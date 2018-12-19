package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/my-stocks-pro/api-gateway-service/infrastructure"
	"github.com/my-stocks-pro/api-gateway-service/proxy"
	"net/http"
	"fmt"
	"errors"
)

type Redis interface {
	Handle(c *gin.Context)
}

type TypeRedis struct {
	config infrastructure.Config
	logger infrastructure.Logger
	consul infrastructure.Consul
	proxy  proxy.Proxy
}

func NewRedis(config infrastructure.Config, logger infrastructure.Logger, consul infrastructure.Consul, proxy proxy.Proxy) TypeRedis {
	return TypeRedis{
		config: config,
		logger: logger,
		consul: consul,
		proxy:  proxy,
	}
}

func (r TypeRedis) Handle(c *gin.Context) {
	msg := &CommonMessage{}

	if err := c.BindJSON(msg); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	servicePath, err := r.consul.DiscoveryService(c.Param("service"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(servicePath)

	switch c.Request.Method {
	case http.MethodPost:
		r.proxy.Request(c.Request.Method, servicePath, c.Request.Body)
	case http.MethodGet:
		body, err := r.proxy.Request(c.Request.Method, servicePath, nil)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		r.proxy.Request(c.Request.Method, servicePath, body)
	default:
		c.JSON(http.StatusMethodNotAllowed, errors.New(c.Request.Method))
		return
	}
}
