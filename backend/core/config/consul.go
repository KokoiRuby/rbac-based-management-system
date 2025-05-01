package config

import consulapi "github.com/hashicorp/consul/api"

func NewConsulClient() *consulapi.Client {
	// TODO: TLS, Cluster
	config := consulapi.DefaultConfig()
	client, err := consulapi.NewClient(config)
	if err != nil {
		panic(err)
	}
	return client
}
