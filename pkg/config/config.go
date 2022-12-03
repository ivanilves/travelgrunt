package config

import (
	"os"

	"gopkg.in/yaml.v3"

	"github.com/ivanilves/travelgrunt/pkg/config/include"
)

var configFile = ".travelgrunt.yml"

// Config is a travelgrunt repo-level configuration (a sequentially evaluated list of rules)
type Config struct {
	Rules []Rule `yaml:"rules"`

	IsDefault bool
}

// DefaultConfig returns default travelgrunt repo-level configuration
func DefaultConfig() Config {
	return Config{
		Rules:     []Rule{{IncludeFn: include.IsTerragrunt}},
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
		cfg.Rules[idx].IncludeFn, err = getIncludeFn(cfg.Rules[idx].Mode)

		if err != nil {
			cfg.Rules = nil

			return cfg, err
		}
	}

	return cfg, nil
}

// Include is a global "decider" function that includes/excludes the path given according to rules
func (cfg Config) Include(d os.DirEntry, rel string) bool {
	for idx := range cfg.Rules {
		if cfg.Rules[idx].Include(d, rel) {
			return !cfg.Rules[idx].Exclude
		}
	}

	return false
}
