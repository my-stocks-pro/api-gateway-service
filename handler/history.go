package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"errors"
	"fmt"
	"github.com/my-stocks-pro/api-gateway-service/infrastructure"
	"github.com/my-stocks-pro/api-gateway-service/proxy"
)

type History interface {
	Handle(c *gin.Context)
}

type TypeHistory struct {
	config infrastructure.Config
	logger infrastructure.Logger
	consul infrastructure.Consul
	proxy  proxy.Proxy
}

type HistoryMessage struct {
	Service  string
	FromData string
	ToData   string
}

func NewHistory(config infrastructure.Config, logger infrastructure.Logger, consul infrastructure.Consul, proxy proxy.Proxy) TypeHistory {
	return TypeHistory{
		config: config,
		logger: logger,
		consul: consul,
		proxy:  proxy,
	}
}

func (h TypeHistory) Handle(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, errors.New("MethodNotAllowed"))
		return
	}

	msg := &HistoryMessage{}
	if err := c.BindJSON(msg); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	servicePath, err := h.consul.DiscoveryService(c.Param("service"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(servicePath)

	body, err := h.proxy.Request(c.Request.Method, servicePath, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(body)
}
