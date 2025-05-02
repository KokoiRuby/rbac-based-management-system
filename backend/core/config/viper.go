package config

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/config"
	"github.com/KokoiRuby/rbac-based-management-system/backend/global"
	"github.com/fsnotify/fsnotify"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

const backendConfigKey = "backend/config"

var conf = pflag.StringP("config", "f", "./config/settings.yaml", "Configuration file")

func Parse() error {
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		return err
	}
	return nil
}

// Load
// Application should  (?!) always fetch configuration from a configuration registry and load to memory, which means
// the configurations should be injected to the configuration registry before application starts.
// At the dev stage, directly update in consul gui once key is persisted.
// TODO: Best practice in production?!
// TODO: struct-ize configuration?!
// TODO: It's necessary when we dump in-mem config for persistence in case remote failure.
func Load() error {

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
		return err
	}

	// Put if no key in remote
	if kvPair == nil {
		bytes, err := os.ReadFile(*conf)
		if err != nil {
			return err
		}
		_, err = client.KV().Put(&consulapi.KVPair{
			Key:   backendConfigKey,
			Value: bytes,
		}, nil)
		if err != nil {
			return err
		}
	}

	// TODO: env-ize endpoint when app is containerized in the future
	err = viper.AddRemoteProvider("consul", "localhost:8500", backendConfigKey)
	if err != nil {
		return err
	}
	viper.SetConfigType("yaml")
	err = viper.ReadRemoteConfig()
	if err != nil {
		return err
	}

	// Save to global
	cfg := config.RuntimeConfig{}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return err
	}
	global.RuntimeConfig = &cfg

	// Dump
	go func() {
		err = Dump()
		if err != nil {
			panic(err)
		}
	}()

	// Open a goroutine to watch remote changes forever
	go func() {

		viper.OnConfigChange(func(e fsnotify.Event) {
			// TODO: Detect diff by runtime config struct and reflection (?!)
			zap.S().Infof("Config changed: %v, operation: %v", e.Name, e.Op)
		})

		for {
			// TODO: env-ize sleep time
			time.Sleep(time.Second * 5)

			err := viper.WatchRemoteConfig()
			if err != nil {
				zap.S().Errorf("unable to watch remote config: %v", err)
				continue
			}

			// TODO: dump on update only
			// unmarshal new config into our runtime config struct. you can also use channel
			// to implement a signal to notify the system of the changes
			err = viper.Unmarshal(global.RuntimeConfig)
			if err != nil {
				zap.S().Errorf("fail to unmarshal config: %s", err)
				continue
			}
			err = Dump()
			if err != nil {
				zap.S().Errorf("fail to dump config: %s", err)
				continue
			}
		}
	}()

	return nil
}

func ParseAndLoad() error {
	err := Parse()
	if err != nil {
		return err
	}
	err = Load()
	if err != nil {
		return err
	}
	return nil
}

func Dump() error {
	byteData, err := yaml.Marshal(global.RuntimeConfig)
	if err != nil {
		zap.S().Errorf("fail to dump config: %s", err)
		return err
	}
	err = os.WriteFile(*conf, byteData, 0644)
	if err != nil {
		zap.S().Errorf("fail to dump config: %s", err)
		return err
	}
	//zap.S().Debugf("Dump config to %v successfully", *conf)
	return nil
}
