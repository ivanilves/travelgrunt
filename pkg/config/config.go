package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/ivanilves/travelgrunt/pkg/config/include"
)

var configFile = ".travelgrunt.yml"

// Config is a travelgrunt repo-level configuration
type Config struct {
	Mode string `yaml:"mode"`

	IsDefault bool
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

	return cfg, validate(cfg)
}

func validate(cfg Config) error {
	allowedModes := []string{"terragrunt", "terraform", "terraform_or_terragrunt"}

	for _, mode := range allowedModes {
		if cfg.Mode == mode {
			return nil
		}
	}

	return fmt.Errorf("illegal mode: %s", cfg.Mode)
}

// DefaultConfig returns default travelgrunt repo-level configuration
func DefaultConfig() Config {
	return Config{Mode: "terragrunt", IsDefault: true}
}

// IncludeFn returns the "include" function used to select relevant directories
func (cfg Config) IncludeFn() (fn func(os.DirEntry) bool) {
	switch cfg.Mode {
	case "terragrunt":
		fn = include.IsTerragrunt
	case "terraform":
		fn = include.IsTerraform
	case "terraform_or_terragrunt":
		fn = include.IsTerraformOrTerragrunt
	default:
		fn = nil
	}

	return fn
}
