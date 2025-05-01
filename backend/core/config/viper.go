package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var conf = pflag.StringP("config", "f", "./core/config/dev/settings.yaml", "Configuration file")

func Parse() {
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(err)
	}
}

func Load() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(*conf)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func ParseAndLoad() {
	Parse()
	Load()
}
