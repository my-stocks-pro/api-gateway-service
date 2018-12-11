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
}

func NewConsul() (ConsulType, error) {
	client, err := consul.NewClient(consul.DefaultConfig())
	if err != nil {
		return ConsulType{}, err
	}
	return ConsulType{
		agent: client.Agent(),
	}, nil
}

func (c ConsulType) DiscoveryService(service string) (string, error) {
	//c.agent.Service(service)
	fmt.Println(service)

	return "", nil
}

