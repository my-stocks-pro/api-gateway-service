package handler

import (
	"github.com/my-stocks-pro/api-gateway-service/infrastructure"
	"github.com/my-stocks-pro/api-gateway-service/proxy"
	"github.com/gin-gonic/gin"
	"net/http"
	"errors"
)

type TypeGateway struct {
	config infrastructure.Config
	logger infrastructure.Logger
	consul infrastructure.Consul
	proxy  proxy.Proxy
}

func New(config infrastructure.Config, logger infrastructure.Logger, consul infrastructure.Consul, proxy proxy.Proxy) TypeGateway {
	return TypeGateway{
		config: config,
		logger: logger,
		consul: consul,
		proxy:  proxy,
	}
}

func (g TypeGateway) HandlerService(c *gin.Context) {
	switch c.Param("service") {
	case "version":
		g.VersionHandle(c)
	case "health":
		g.HealthHandle(c)
	case "scheduler", "redis", "postgres":
		g.CommonHandler(c)
	case "history":
		g.HistoryHandler(c)
	default:
		c.JSON(http.StatusBadRequest, errors.New("StatusBadRequest"))
		return
	}
}
