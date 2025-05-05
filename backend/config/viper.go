package config

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/config/runtime"
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

type Flags struct {
	ConfigFile  string
	AutoMigrate bool
}

type RuntimeConfig struct {
	Logging runtime.Logging     `yaml:"logging"`
	RDB     runtime.RDBConfig   `yaml:"rdb"`
	Redis   runtime.RedisConfig `yaml:"redis"`
	Mongo   runtime.MongoConfig `yaml:"mongo"`
	Gin     runtime.GinConfig   `yaml:"gin"`
	Flags   Flags               `yaml:"-"` // As member not embedded
}

func NewRuntimeConfig() *RuntimeConfig {
	viper.AutomaticEnv()

	cfg := &RuntimeConfig{}
	Parse(cfg)

	client := NewConsulClient()
	kvPair, _, err := client.KV().Get(backendConfigKey, nil)
	if err != nil {
		zap.S().Fatalf("failed to load config from consul: %s", err)
	}

	// Put if no key in remote
	if kvPair == nil {
		bytes, err := os.ReadFile(cfg.Flags.ConfigFile)
		if err != nil {
			zap.S().Fatalf("failed to load config file: %s", err)
		}
		_, err = client.KV().Put(&consulapi.KVPair{
			Key:   backendConfigKey,
			Value: bytes,
		}, nil)
		if err != nil {
			zap.S().Fatalf("failed to save config to consul: %s", err)
		}
	}

	// Key is in remote
	// TODO: env-ize endpoint when app is containerized in the future
	err = viper.AddRemoteProvider("consul", "localhost:8500", backendConfigKey)
	if err != nil {
		zap.S().Fatalf("failed to add remote provider: %s", err)
	}
	viper.SetConfigType("yaml")
	err = viper.ReadRemoteConfig()
	if err != nil {
		zap.S().Fatalf("failed to read remote config: %s", err)
	}

	// Unmarshal
	err = viper.Unmarshal(&cfg)
	if err != nil {
		zap.S().Fatalf("failed to unmarshal config: %s", err)
	}

	// Dump
	go Dump(cfg)

	// Watch
	go Watch(cfg)

	return cfg
}

func Parse(cfg *RuntimeConfig) {

	pflag.StringVarP(&cfg.Flags.ConfigFile,
		"config", "f", "./config/settings.yaml", "Configuration file")
	pflag.BoolVarP(&cfg.Flags.AutoMigrate,
		"migrate", "m", false, "Auto migrate models to RDB") // When RDB is ready

	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		zap.S().Fatalf("failed to bind flags: %s", err)
	}
}

func Dump(cfg *RuntimeConfig) {
	byteData, err := yaml.Marshal(cfg)
	if err != nil {
		zap.S().Fatalf("fail to marshal config: %s", err)
	}
	err = os.WriteFile(cfg.Flags.ConfigFile, byteData, 0644)
	if err != nil {
		zap.S().Fatalf("fail to dump config: %s", err)
	}
}

func Watch(cfg *RuntimeConfig) {
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
		err = viper.Unmarshal(cfg)
		if err != nil {
			zap.S().Errorf("fail to unmarshal config: %s", err)
			continue
		}
		Dump(cfg)
		if err != nil {
			zap.S().Errorf("fail to dump config: %s", err)
			continue
		}
	}
}
