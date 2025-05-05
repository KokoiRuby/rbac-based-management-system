package runtime

import (
	"fmt"
	"time"
)

type GinConfig struct {
	IP      string        `yaml:"ip"`
	Port    string        `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func (g GinConfig) Addr() string {
	return fmt.Sprintf("%s:%s", g.IP, g.Port)
}
