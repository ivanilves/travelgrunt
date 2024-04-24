package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithNameEx(t *testing.T) {
	assert := assert.New(t)

	cfg := DefaultConfig()

	cfg = cfg.WithNameEx("*.tf")

	assert.Len(cfg.Rules, 1)
}
