package handler

import (
	"github.com/my-stocks-pro/api-gateway-service/infrastructure"
	"github.com/my-stocks-pro/api-gateway-service/proxy"
	"github.com/gin-gonic/gin"
	"net/http"
	"errors"
)

type Earnings interface {
	Handle(c *gin.Context)
}

func NewEarnings(config infrastructure.Config, logger infrastructure.Logger, consul infrastructure.Consul, proxy proxy.Proxy) TypeEarnings {
	return TypeEarnings{
		config: config,
		logger: logger,
		consul: consul,
		proxy:  proxy,
	}
}

type TypeEarnings struct {
	config infrastructure.Config
	logger infrastructure.Logger
	consul infrastructure.Consul
	proxy  proxy.Proxy
}

func (s TypeEarnings) Handle(c *gin.Context) {

	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusBadRequest, errors.New("StatusBadRequest"))
		return
	}

	servicePath, err := s.consul.DiscoveryService(c.Param(service))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	_, err = s.proxy.Request(c.Request.Method, servicePath, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
}
