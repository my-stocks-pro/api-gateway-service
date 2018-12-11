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
	Config infrastructure.Config
	Logger infrastructure.Logger
	Consul infrastructure.Consul
	Proxy  proxy.Proxy
}

type HistoryMessage struct {
	Service  string
	FromData string
	ToData   string
}

func NewHistory(config infrastructure.Config, logger infrastructure.Logger, consul infrastructure.Consul, proxy proxy.Proxy) TypeScheduler {
	return TypeScheduler{
		Config: config,
		Logger: logger,
		Consul: consul,
		Proxy:  proxy,
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

	servicePath, err := h.Consul.DiscoveryService(c.Param(ServiceKey))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(servicePath)

	body, err := h.Proxy.Request(c.Request.Method, servicePath, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(body)
}
