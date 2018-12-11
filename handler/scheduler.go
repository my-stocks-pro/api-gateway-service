package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/my-stocks-pro/api-gateway-service/infrastructure"
	"github.com/my-stocks-pro/api-gateway-service/proxy"
)

type Scheduler interface {
	Handle(c *gin.Context)
}

func NewScheduler(config infrastructure.Config, logger infrastructure.Logger, consul infrastructure.Consul, proxy proxy.Proxy) TypeScheduler {
	return TypeScheduler{
		Config: config,
		Logger: logger,
		Consul: consul,
		Proxy:  proxy,
	}
}

type TypeScheduler struct {
	Config infrastructure.Config
	Logger infrastructure.Logger
	Consul infrastructure.Consul
	Proxy  proxy.Proxy
}

func (v TypeScheduler) Handle(c *gin.Context) {

}
