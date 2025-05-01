package config

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"os"
	"time"
)

const backendConfigKey = "backend/config"

var conf = pflag.StringP("config", "f", "./core/config/dev/settings.yaml", "Configuration file")

func Parse() {
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(err)
	}
}

// Load
// Application should  (?!) always fetch configuration from a configuration registry and load to memory, which means
// the configurations should be injected to the configuration registry before application starts.
// At the dev stage, directly update in consul gui once key is persisted.
// TODO: Best practice in production?!
// TODO: struct-ize configuration?!
func Load() {

	client := NewConsulClient()

	//$ consul kv get -detailed backend/config
	//CreateIndex      112
	//Flags            0
	//Key              backend/config
	//LockIndex        0
	//ModifyIndex      112
	//Session          -
	//Value            version: "0.0.1"

	kvPair, _, err := client.KV().Get(backendConfigKey, nil)
	if err != nil {
		panic(err)
	}

	// Put if no key in Consul
	if kvPair == nil {
		bytes, err := os.ReadFile(*conf)
		if err != nil {
			panic(err)
		}
		_, err = client.KV().Put(&consulapi.KVPair{
			Key:   backendConfigKey,
			Value: bytes,
		}, nil)
		if err != nil {
			panic(err)
		}
	}

	// TODO: env-ize endpoint when app is containerized in the future
	err = viper.AddRemoteProvider("consul", "localhost:8500", backendConfigKey)
	if err != nil {
		panic(err)
	}
	viper.SetConfigType("yaml")
	err = viper.ReadRemoteConfig()
	if err != nil {
		panic(err)
	}

	// Open a goroutine to watch remote changes forever
	go func() {
		for {
			// TODO: env-ize sleep time
			time.Sleep(time.Second * 5)

			err := viper.WatchRemoteConfig()
			if err != nil {
				continue
			}
			err = viper.ReadRemoteConfig()
			if err != nil {
				panic(err)
			}
			fmt.Println(viper.Get("version"))
		}
	}()

}

func ParseAndLoad() {
	Parse()
	Load()
}
