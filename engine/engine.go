package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/my-stocks-pro/api-gateway-service/handler"
)

func (s *Server) InitMux() {

	s.Engine.GET("/gateway/:service", s.GetHandler)
	s.Engine.POST("/gateway/:service", s.GetHandler)
}


func (s *Server) GetHandler(c *gin.Context) {
	serviceType := c.Param("service")
	_, ok := s.Handler[serviceType]
	if !ok {
		s.Handler[serviceType] = s.HandlerConstruct(serviceType)
	}
	s.Handler[serviceType].Handle(c)
}

func (s *Server) HandlerConstruct(serviceType string) handler.Handler {
	switch serviceType {
	case "version":
		return handler.TypeVersion{}
	case "health":
		return handler.TypeVersion{}
	case "scheduler":
		return handler.TypeVersion{}
	case "redis":
		return handler.TypeVersion{}
	case "postgres":
		return handler.TypeVersion{}
	}
	return nil
}