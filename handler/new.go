package handler

import (
	"github.com/my-stocks-pro/api-gateway-service/infrastructure"
	"net/http"
)

type Gateway interface {
}

type GatewayType struct {
	config     infrastructure.Config
	logger     infrastructure.Logger
	consul     infrastructure.Consul
	httpClient *http.Client
}

func New(config infrastructure.Config, logger infrastructure.Logger, consul infrastructure.Consul, httpClient *http.Client) GatewayType {
	return GatewayType{
		config:     config,
		logger:     logger,
		consul:     consul,
		httpClient: httpClient,
	}
}
