package config

import (
	"os"

	"gopkg.in/yaml.v3"

	"github.com/ivanilves/travelgrunt/pkg/config/mode"
	"github.com/ivanilves/travelgrunt/pkg/config/rule"
)

var configFile = ".travelgrunt.yml"

// Config is a travelgrunt repo-level configuration (a sequentially evaluated list of rules)
type Config struct {
	Rules []rule.Rule `yaml:"rules"`

	IsDefault bool
}

// DefaultConfig returns default travelgrunt repo-level configuration
func DefaultConfig() Config {
	return Config{
		Rules:     []rule.Rule{{ModeFn: mode.IsTerragrunt}},
		IsDefault: true,
	}
}

// NewConfig creates new travelgrunt repo-level configuration
func NewConfig(path string) (cfg Config, err error) {
	var data []byte

	data, err = os.ReadFile(path + "/" + configFile)
	// We don't care about config file not being read,
	// it's a common case and we just return default config.
	if err != nil {
		return DefaultConfig(), nil
	}
	// If we have a file, it should be a correct one though!
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return cfg, err
	}

	for idx := range cfg.Rules {
		cfg.Rules[idx].ModeFn, err = getModeFn(cfg.Rules[idx].Mode)

		if err != nil {
			cfg.Rules = nil

			return cfg, err
		}
	}

	return cfg, nil
}

// Admit is a global "decider" function that includes/excludes the path given
func (cfg Config) Admit(d os.DirEntry, relPath string) bool {
	for idx := range cfg.Rules {
		if cfg.Rules[idx].Admit(d, relPath) {
			return !cfg.Rules[idx].Negate
		}
	}

	return false
}
