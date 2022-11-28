package config

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"reflect"
	"runtime"

	"github.com/ivanilves/travelgrunt/pkg/config/include"
)

const (
	fixturePath = "../../fixtures/config"
)

func TestCornerCases(t *testing.T) {
	assert := assert.New(t)

	testCases := map[string]struct {
		cfg     Config
		success bool
	}{
		"travelgrunt.yml.invalid":     {cfg: Config{Mode: "", IncludeFn: nil, IsDefault: false}, success: false},
		"travelgrunt.yml.illegal":     {cfg: Config{Mode: "bogus", IncludeFn: nil, IsDefault: false}, success: false},
		"travelgrunt.yml.nonexistent": {cfg: Config{Mode: "terragrunt", IncludeFn: include.IsTerragrunt, IsDefault: true}, success: true},
	}

	for cfgFile, expected := range testCases {
		configFile = cfgFile

		cfg, err := NewConfig(fixturePath)

		assert.Equal(expected.cfg.Mode, cfg.Mode)
		assert.Equal(expected.cfg.IsDefault, cfg.IsDefault)

		if expected.success {
			assert.NotNil(cfg.IncludeFn)
			assert.Nil(err)
		} else {
			assert.Nil(cfg.IncludeFn)
			assert.NotNil(err)
		}
	}
}

func getNormalConfig(mode string) Config {
	includeFn, _ := GetIncludeFn(mode)

	return Config{Mode: mode, IncludeFn: includeFn, IsDefault: false}
}

func TestNormalFlow(t *testing.T) {
	assert := assert.New(t)

	testCases := map[string]Config{
		"travelgrunt.yml.terragrunt":              getNormalConfig("terragrunt"),
		"travelgrunt.yml.terraform":               getNormalConfig("terraform"),
		"travelgrunt.yml.terraform_or_terragrunt": getNormalConfig("terraform_or_terragrunt"),
		"travelgrunt.yml.dockerfile":              getNormalConfig("dockerfile"),
		"travelgrunt.yml.jenkins":                 getNormalConfig("jenkins"),
	}

	for cfgFile, expected := range testCases {
		configFile = cfgFile

		cfg, err := NewConfig(fixturePath)

		assert.Equal(expected.Mode, cfg.Mode)
		assert.Equal(expected.IsDefault, false)

		assert.Equalf(
			runtime.FuncForPC(reflect.ValueOf(expected.IncludeFn).Pointer()).Name(),
			runtime.FuncForPC(reflect.ValueOf(cfg.IncludeFn).Pointer()).Name(),
			"got unexpected include function while loading config file: %s", configFile,
		)

		assert.Nil(err)
	}
}
