package config

import (
	"github.com/ivanilves/travelgrunt/pkg/config/rule"
)

// WithNameEx returns a copy of config with regex-based rule added
func (cfg Config) WithNameEx(nameEx string) Config {
	cfg.Rules = []rule.Rule{rule.Rule{NameEx: nameEx}}

	return cfg
}
