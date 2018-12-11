package engine

import (
	"github.com/my-stocks-pro/api-gateway-service/handler"
	"github.com/gin-gonic/gin"
	"github.com/my-stocks-pro/api-gateway-service/infrastructure"
	"github.com/my-stocks-pro/api-gateway-service/proxy"
)

type Server struct {
	config  infrastructure.Config
	logger  infrastructure.Logger
	consul  infrastructure.Consul
	proxy   proxy.Proxy
	Engine  *gin.Engine
	Handler map[string]handler.Handler
}

func New(config infrastructure.Config, logger infrastructure.Logger, consul infrastructure.Consul, proxy proxy.Proxy) Server {
	return Server{
		config:  config,
		logger:  logger,
		consul:  consul,
		proxy:   proxy,
		Engine:  gin.New(),
		Handler: map[string]handler.Handler{},
	}
}
