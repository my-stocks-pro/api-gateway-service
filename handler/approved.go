package handler

import (
	"github.com/my-stocks-pro/api-gateway-service/infrastructure"
	"github.com/my-stocks-pro/api-gateway-service/proxy"
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type Approved interface {
	Handle(c *gin.Context)
}

func NewApproved(config infrastructure.Config, logger infrastructure.Logger, consul infrastructure.Consul, proxy proxy.Proxy) TypeApproved {
	return TypeApproved{
		config: config,
		logger: logger,
		consul: consul,
		proxy:  proxy,
	}
}

type TypeApproved struct {
	config infrastructure.Config
	logger infrastructure.Logger
	consul infrastructure.Consul
	proxy  proxy.Proxy
}

func (s TypeApproved) Handle(c *gin.Context) {
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
