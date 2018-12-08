package main

import (
	consul "github.com/hashicorp/consul/api"
	"log"
	"fmt"
)

type ConsulType struct {
	agent *consul.Agent
}

func NewConsul() ConsulType {
	return ConsulType{}
}


func (c ConsulType) DiscoveryService(service string) string {
	//c.agent.Service(service)
	fmt.Println(service)

	return ""
}

func (c ConsulType) NewClient() *consul.Agent {
	client, err := consul.NewClient(consul.DefaultConfig())
	if err != nil {
		log.Fatal(err)
	}
	return client.Agent()
}