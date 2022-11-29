package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/ivanilves/travelgrunt/pkg/config/include"
)

var configFile = ".travelgrunt.yml"

// Config contains travelgrunt repo-level configuration
type Config struct {
	Mode      string `yaml:"mode"`
	IncludeFn func(os.DirEntry) bool

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

	cfg.IncludeFn, err = GetIncludeFn(cfg.Mode)

	return cfg, err
}

// DefaultConfig returns default travelgrunt repo-level configuration
func DefaultConfig() Config {
	return Config{
		Mode:      "terragrunt",
		IncludeFn: include.IsTerragrunt,
		IsDefault: true,
	}
}

// GetIncludeFn gets an "include" func for the given mode (if mode is unknown, it returns a nil func and a non-nil error)
func GetIncludeFn(mode string) (fn func(os.DirEntry) bool, err error) {
	err = nil

	switch mode {
	case "terragrunt":
		fn = include.IsTerragrunt
	case "terraform":
		fn = include.IsTerraform
	case "terraform_or_terragrunt":
		fn = include.IsTerraformOrTerragrunt
	case "dockerfile":
		fn = include.IsDockerfile
	case "jenkins":
		fn = include.IsJenkins
	default:
		fn = nil
		err = fmt.Errorf("illegal mode: %s", mode)
	}

	return fn, err
}
