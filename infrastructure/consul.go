package infrastructure

import (
	consul "github.com/hashicorp/consul/api"
	"fmt"
)

type Consul interface {
	DiscoveryService(service string) (string, error)
}

type ConsulType struct {
	agent *consul.Agent
	service map[string]string
}

func NewConsul() (ConsulType, error) {
	//client, err := consul.NewClient(consul.DefaultConfig())
	//if err != nil {
	//	return ConsulType{}, err
	//}
	return ConsulType{
		//agent: client.Agent(),
		service: map[string]string{
			"earnings": "http://127.0.0.1:9002",
			"approved": "http://127.0.0.1:9003",
			"rejected": "http://127.0.0.1:9004",
		},
	}, nil

}

func (c ConsulType) DiscoveryService(service string) (string, error) {
	//c.agent.Service(service)
	fmt.Println(service)

	return c.service[service], nil
}

