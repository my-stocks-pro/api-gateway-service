package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/my-stocks-pro/api-gateway-service/handler"
)

const (
	ServiceKey = "service"
	version    = "version"
	health     = "health"
	scheduler  = "scheduler"
	redis      = "redis"
	postgres   = "postgres"
	history    = "history"
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
	case version:
		return handler.NewVersion()
	case health:
		return handler.NewHealth()
	case scheduler:
		return handler.NewScheduler(s.config, s.logger, s.consul, s.proxy)
	case redis:
		return handler.NewRedis(s.config, s.logger, s.consul, s.proxy)
	case postgres:
		return handler.NewPostgres(s.config, s.logger, s.consul, s.proxy)
	case history:
		return handler.NewHistory(s.config, s.logger, s.consul, s.proxy)
	}
	return nil
}
