package runtime

import (
	"fmt"
)

type GinConfig struct {
	IP      string `yaml:"ip"`
	Port    string `yaml:"port"`
	Timeout uint   `yaml:"timeout"`
	TLS     bool   `yaml:"tls"`
	MTLS    bool   `yaml:"mtls"`
}

func (g GinConfig) Addr() string {
	return fmt.Sprintf("%s:%s", g.IP, g.Port)
}
