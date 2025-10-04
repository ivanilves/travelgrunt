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

func TestNewConfigLinksFlow(t *testing.T) {
	assert := assert.New(t)

	testCases := map[string]bool{
		"travelgrunt.yml.whatever": false,
		"travelgrunt.yml.links":    true,
	}

	for cfgFile, shouldHaveLinks := range testCases {
		configFile = cfgFile

		cfg, err := NewConfig(fixturePath)

		if shouldHaveLinks {
			assert.NotEmpty(cfg.Links)
		} else {
			assert.Empty(cfg.Links)

		}

		assert.Nil(err)
	}
}

func TestNewHomeConfig(t *testing.T) {
	assert := assert.New(t)

	// Save original configFile
	originalConfigFile := configFile
	defer func() { configFile = originalConfigFile }()

	// Test with links fixture
	configFile = "travelgrunt.yml.links"

	cfg, err := NewHomeConfig(fixturePath)

	assert.Nil(err)
	assert.Empty(cfg.Rules, "rules should be empty in home config mode")
	assert.True(cfg.UseLinks, "UseLinks should be true in home config mode")
	assert.False(cfg.IsDefault, "IsDefault should be false")
	assert.NotEmpty(cfg.Links, "links should not be empty")
}

func TestNewHomeConfigNoLinks(t *testing.T) {
	assert := assert.New(t)

	// Save original configFile
	originalConfigFile := configFile
	defer func() { configFile = originalConfigFile }()

	// Test with config that has rules but no links
	configFile = "travelgrunt.yml.terragrunt"

	_, err := NewHomeConfig(fixturePath)
	assert.NotNil(err, "should fail when no links found")
	assert.Contains(err.Error(), "no links found")
}

func TestNewHomeConfigInvalidYAML(t *testing.T) {
	assert := assert.New(t)

	// Save original configFile
	originalConfigFile := configFile
	defer func() { configFile = originalConfigFile }()

	// Test with invalid YAML
	configFile = "travelgrunt.yml.invalid"

	_, err := NewHomeConfig(fixturePath)
	assert.NotNil(err, "should fail with invalid YAML")
	assert.Contains(err.Error(), "failed to parse")
}

func TestNewHomeConfigNonexistent(t *testing.T) {
	assert := assert.New(t)

	// Save original configFile
	originalConfigFile := configFile
	defer func() { configFile = originalConfigFile }()

	// Test with nonexistent file
	configFile = "travelgrunt.yml.doesnotexist"

	_, err := NewHomeConfig(fixturePath)
	assert.NotNil(err, "should fail when file doesn't exist")
	assert.Contains(err.Error(), "failed to read")
}
