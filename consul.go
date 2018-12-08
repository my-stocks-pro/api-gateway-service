package main

import (
	consul "github.com/hashicorp/consul/api"
	"log"
)

type ConsulType struct {
	agent *consul.Agent
}

func NewConsul() ConsulType {
	return ConsulType{}
}


func (c ConsulType) DiscoveryService(service string) string {

	return ""
}

func (c ConsulType) NewClient() *consul.Agent {
	client, err := consul.NewClient(consul.DefaultConfig())
	if err != nil {
		log.Fatal(err)
	}
	return client.Agent()
}