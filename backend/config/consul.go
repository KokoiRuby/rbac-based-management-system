package config

import (
	consulapi "github.com/hashicorp/consul/api"
	"os"
)

func NewConsulClient() *consulapi.Client {
	// TODO: TLS, Cluster
	consulAddr := os.Getenv("CONSUL_HTTP_ADDR")
	if consulAddr == "" {
		consulAddr = "http://consul:8500"
	}

	config := consulapi.DefaultConfig()
	config.Address = consulAddr

	client, err := consulapi.NewClient(config)
	if err != nil {
		panic(err)
	}
	return client
}
