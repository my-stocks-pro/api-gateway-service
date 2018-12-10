package main

import (
	"github.com/my-stocks-pro/api-gateway-service/infrastructure"
	"github.com/my-stocks-pro/api-gateway-service/engine"
	"github.com/my-stocks-pro/api-gateway-service/handler"
	"github.com/my-stocks-pro/api-gateway-service/proxy"
)

func main() {

	config := infrastructure.NewConfig()

	logger, err := infrastructure.NewLogger()
	if err != nil {
		panic(err)
	}

	consul, err := infrastructure.NewConsul()
	if err != nil {
		panic(err)
	}

	httpClient := infrastructure.GetHTTPClient()

	prx := proxy.New(httpClient)

	gateway := handler.New(config, logger, consul, prx)

	server := engine.New(gateway)

	server.InitMux()

	server.Engine.Run(config.Port)
}
