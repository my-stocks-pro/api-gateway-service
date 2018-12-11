package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/my-stocks-pro/api-gateway-service/infrastructure"
	"github.com/my-stocks-pro/api-gateway-service/proxy"
)

type Postgres interface {
	Handle(c *gin.Context)
}

type TypePostgres struct {
	Config  infrastructure.Config
	Logger  infrastructure.Logger
	Consul  infrastructure.Consul
	Proxy   proxy.Proxy
}

func NewPostgres(config infrastructure.Config, logger infrastructure.Logger, consul infrastructure.Consul, proxy proxy.Proxy) TypeScheduler {
	return TypeScheduler{
		Config: config,
		Logger: logger,
		Consul: consul,
		Proxy:  proxy,
	}
}


func (v TypePostgres) Handle(c *gin.Context) {

}