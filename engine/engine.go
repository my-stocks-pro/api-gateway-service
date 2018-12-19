package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/my-stocks-pro/api-gateway-service/handler"
)

const (
	version  = "version"
	health   = "health"
	earnings = "earnings"
	approved = "approved"
	rejected = "rejected"
	redis    = "redis"
	postgres = "postgres"
	history  = "history"
	slack    = "slack"
)

func (s *Server) InitMux() {

	s.Engine.GET("/health", s.GetServiceHandler(health).Handle)
	s.Engine.GET("/version", s.GetServiceHandler(version).Handle)

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

func (s *Server) GetServiceHandler(h string) handler.Handler {
	switch h {
	case version:
		return handler.NewVersion(s.config)
	case health:
		return handler.NewHealth(s.config)
	default:
		return nil
	}
	return nil
}

func (s *Server) HandlerConstruct(serviceType string) handler.Handler {
	switch serviceType {
	case earnings:
		return handler.NewEarnings(s.config, s.logger, s.consul, s.proxy)
	case approved:
		return handler.NewApproved(s.config, s.logger, s.consul, s.proxy)
	case rejected:
		return handler.NewRejected(s.config, s.logger, s.consul, s.proxy)
	case redis:
		return handler.NewRedis(s.config, s.logger, s.consul, s.proxy)
	case postgres:
		return handler.NewPostgres(s.config, s.logger, s.consul, s.proxy)
	case history:
		return handler.NewHistory(s.config, s.logger, s.consul, s.proxy)
	case slack:
		return handler.NewSlack(s.config, s.logger, s.consul, s.proxy)
	}
	return nil
}
