package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"errors"
	"fmt"
	"io/ioutil"
)

type Handler interface {
	Handle(c *gin.Context)
}

const ServiceKey = "service"

type CommonMessage struct {
	Service string
	Data    []byte
}

func (g TypeGateway) CommonHandler(c *gin.Context) {
	var body []byte
	var err error
	msg := &CommonMessage{}

	if err = c.BindJSON(msg); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	servicePath, err := g.consul.DiscoveryService(c.Param(ServiceKey))
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

	body, err = g.proxy.Request(c.Request.Method, servicePath, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
}
