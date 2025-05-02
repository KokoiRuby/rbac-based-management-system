package config

import "github.com/KokoiRuby/rbac-based-management-system/backend/config/runtime"

type RuntimeConfig struct {
	Logging runtime.Logging `yaml:"logging"`
}
