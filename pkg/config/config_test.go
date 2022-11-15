package config

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"os"
	"reflect"
	"runtime"

	"github.com/ivanilves/travelgrunt/pkg/config/include"
)

const (
	fixturePath = "../../fixtures/config"
)

func getConfig(mode string, isDefault bool) Config {
	return Config{Mode: mode, IsDefault: isDefault}
}

func TestNewConfig(t *testing.T) {
	assert := assert.New(t)

	testCases := map[string]struct {
		cfg       Config
		success   bool
		includeFn func(os.DirEntry) bool
	}{
		"travelgrunt.yml.terragrunt":              {cfg: getConfig("terragrunt", false), success: true, includeFn: include.IsTerragrunt},
		"travelgrunt.yml.terraform":               {cfg: getConfig("terraform", false), success: true, includeFn: include.IsTerraform},
		"travelgrunt.yml.terraform_or_terragrunt": {cfg: getConfig("terraform_or_terragrunt", false), success: true, includeFn: include.IsTerraformOrTerragrunt},
		"travelgrunt.yml.invalid":                 {cfg: getConfig("", false), success: false, includeFn: nil},
		"travelgrunt.yml.illegal":                 {cfg: getConfig("bogus", false), success: false, includeFn: nil},
		"travelgrunt.yml.nonexistent":             {cfg: getConfig("terragrunt", true), success: true, includeFn: include.IsTerragrunt},
	}

	for cfgFile, expected := range testCases {
		configFile = cfgFile

		cfg, err := NewConfig(fixturePath)

		assert.Equal(expected.cfg, cfg)

		assert.Equalf(
			runtime.FuncForPC(reflect.ValueOf(expected.includeFn).Pointer()).Name(),
			runtime.FuncForPC(reflect.ValueOf(cfg.IncludeFn()).Pointer()).Name(),
			"got unexpected include function while loading config file: %s", configFile,
		)

		if expected.success {
			assert.Nil(err)
		} else {
			assert.NotNil(err)
		}
	}
}
