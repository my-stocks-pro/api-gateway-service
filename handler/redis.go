package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/my-stocks-pro/api-gateway-service/infrastructure"
	"github.com/my-stocks-pro/api-gateway-service/proxy"
	"net/http"
	"fmt"
	"io/ioutil"
	"errors"
)

type Redis interface {
	Handle(c *gin.Context)
}

type TypeRedis struct {
	Config infrastructure.Config
	Logger infrastructure.Logger
	Consul infrastructure.Consul
	Proxy  proxy.Proxy
}

func NewRedis(config infrastructure.Config, logger infrastructure.Logger, consul infrastructure.Consul, proxy proxy.Proxy) TypeScheduler {
	return TypeScheduler{
		Config: config,
		Logger: logger,
		Consul: consul,
		Proxy:  proxy,
	}
}

func (r TypeRedis) Handle(c *gin.Context) {
	var body []byte
	var err error
	msg := &CommonMessage{}

	if err = c.BindJSON(msg); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	servicePath, err := r.Consul.DiscoveryService(c.Param(ServiceKey))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(servicePath)

	switch c.Request.Method {
	case http.MethodPost:
		body, err = ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	case http.MethodGet:
		break
	default:
		c.JSON(http.StatusMethodNotAllowed, errors.New(c.Request.Method))
		return
	}

	fmt.Println(c.Request.Method, c.Param(ServiceKey))
	fmt.Println(body)

	body, err = r.Proxy.Request(c.Request.Method, servicePath, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
}
