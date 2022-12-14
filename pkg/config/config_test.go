package config

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"reflect"
	"runtime"

	"github.com/ivanilves/travelgrunt/pkg/config/mode"
	"github.com/ivanilves/travelgrunt/pkg/config/rule"
)

const (
	fixturePath = "../../fixtures/config"
)

func TestNewConfigCornerCases(t *testing.T) {
	assert := assert.New(t)

	testCases := map[string]struct {
		cfg     Config
		success bool
	}{
		"travelgrunt.yml.invalid":     {cfg: Config{Rules: nil, IsDefault: false}, success: false},
		"travelgrunt.yml.illegal":     {cfg: Config{Rules: nil, IsDefault: false}, success: false},
		"travelgrunt.yml.nonexistent": {cfg: Config{Rules: []rule.Rule{{Mode: "terragrunt", ModeFn: mode.IsTerragrunt}}, IsDefault: true}, success: true},
	}

	for cfgFile, expected := range testCases {
		configFile = cfgFile

		cfg, err := NewConfig(fixturePath)

		assert.Equal(expected.cfg.IsDefault, cfg.IsDefault)

		if expected.success {
			assert.NotNil(cfg.Rules)
			assert.Nil(err)
		} else {
			assert.Nil(cfg.Rules)
			assert.NotNil(err)
		}
	}
}

func getNormalConfig(mode string) Config {
	fn, _ := getModeFn(mode)

	return Config{Rules: []rule.Rule{{Mode: mode, ModeFn: fn}}, IsDefault: false}
}

func TestNewConfigNormalFlow(t *testing.T) {
	assert := assert.New(t)

	testCases := map[string]Config{
		"travelgrunt.yml.terragrunt":  getNormalConfig("terragrunt"),
		"travelgrunt.yml.terraform":   getNormalConfig("terraform"),
		"travelgrunt.yml.dockerfile":  getNormalConfig("dockerfile"),
		"travelgrunt.yml.jenkinsfile": getNormalConfig("jenkinsfile"),
		"travelgrunt.yml.groovy":      getNormalConfig("groovy"),
	}

	for cfgFile, expected := range testCases {
		configFile = cfgFile

		cfg, err := NewConfig(fixturePath)

		assert.NotNil(expected.Rules, cfg.Rules)
		assert.Equal(expected.IsDefault, false)

		assert.Equalf(
			runtime.FuncForPC(reflect.ValueOf(expected.Rules[0].ModeFn).Pointer()).Name(),
			runtime.FuncForPC(reflect.ValueOf(cfg.Rules[0].ModeFn).Pointer()).Name(),
			"got unexpected mode function while loading config file: %s", configFile,
		)

		assert.Nil(err)
	}
}
