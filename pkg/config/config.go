package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/ivanilves/travelgrunt/pkg/config/mode"
	"github.com/ivanilves/travelgrunt/pkg/config/rule"
)

var configFile = ".travelgrunt.yml"

// Config is a travelgrunt repo-level configuration (a sequentially evaluated list of rules)
type Config struct {
	Rules []rule.Rule `yaml:"rules"`

	Links []string `yaml:"links"`

	IsDefault bool
	UseFiles  bool
	UseLinks  bool
}

// DefaultConfig returns default travelgrunt repo-level configuration
func DefaultConfig() Config {
	return Config{
		Rules:     []rule.Rule{{ModeFn: mode.IsTerragrunt}},
		IsDefault: true,
		UseFiles:  false,
		UseLinks:  false,
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

// NewHomeConfig loads config from home directory for non-repo mode
// Only links are loaded, rules are ignored
// If homePath is empty, os.UserHomeDir() is used
func NewHomeConfig(homePath string) (Config, error) {
	if homePath == "" {
		var err error
		homePath, err = os.UserHomeDir()
		if err != nil {
			return Config{}, fmt.Errorf("failed to get home directory: %w", err)
		}
	}

	configPath := filepath.Join(homePath, configFile)
	data, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read %s: %w", configPath, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("failed to parse %s: %w", configPath, err)
	}

	// Clear rules, only use links
	cfg.Rules = []rule.Rule{}
	cfg.UseLinks = true
	cfg.IsDefault = false

	if len(cfg.Links) == 0 {
		return Config{}, fmt.Errorf("no links found in %s", configPath)
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
