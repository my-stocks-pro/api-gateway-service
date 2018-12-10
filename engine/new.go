package engine

import (
	"github.com/my-stocks-pro/api-gateway-service/handler"
	"github.com/gin-gonic/gin"
)

type Server struct {
	gateway handler.TypeGateway
	Engine  *gin.Engine
}

func New(gateway handler.TypeGateway) Server {
	return Server{
		gateway: gateway,
		Engine:  gin.New(),
	}
}
