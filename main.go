package main

import (
	"github.com/my-stocks-pro/api-gateway-service/infrastructure"
	"github.com/my-stocks-pro/api-gateway-service/engine"
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

	server := engine.New(config, logger, consul, httpClient)

	server.InitMux()

	server.Engine.Run(config.Port)
}
