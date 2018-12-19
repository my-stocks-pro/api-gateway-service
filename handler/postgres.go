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
	config infrastructure.Config
	logger infrastructure.Logger
	consul infrastructure.Consul
	proxy  proxy.Proxy
}

func NewPostgres(config infrastructure.Config, logger infrastructure.Logger, consul infrastructure.Consul, proxy proxy.Proxy) TypePostgres {
	return TypePostgres{
		config: config,
		logger: logger,
		consul: consul,
		proxy:  proxy,
	}
}

func (v TypePostgres) Handle(c *gin.Context) {

}
