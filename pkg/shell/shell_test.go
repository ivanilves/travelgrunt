package shell

import (
	"os"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectShell(t *testing.T) {
	assert := assert.New(t)

	cases := map[string]string{
		"/bin/bash": "/bin/bash",
		"/bin/zsh":  "/bin/zsh",
		"/bin/yolo": defaultShell,
		"":          defaultShell,
	}

	for input, expected := range cases {
		os.Setenv("SHELL", input)

		assert.Equal(expected, detectShell())
	}
}

func TestDetectLines(t *testing.T) {
	assert := assert.New(t)

	cases := map[string]int{
		"50":  50,
		"xyz": defaultLines,
		"":    defaultLines,
	}

	for input, expected := range cases {
		os.Setenv("LINES", input)

		assert.Equal(expected, detectLines())
	}
}

func TestIsMocked(t *testing.T) {
	assert := assert.New(t)

	cases := map[string]bool{
		"/bin/bash":  false,
		"/bin/zsh":   false,
		"/bin/true":  true,
		"/bin/false": true,
	}

	for input, expected := range cases {
		os.Setenv("SHELL", input)

		assert.Equal(expected, isMocked())
	}
}

func TestName(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(os.Args[0], Name())
}

func TestIsRunningInside(t *testing.T) {
	assert := assert.New(t)

	cases := map[string]bool{
		"true":     true,
		"false":    false,
		"whatever": false,
	}

	for input, expected := range cases {
		os.Setenv("TTG", input)

		assert.Equal(expected, IsRunningInside())
	}

	os.Setenv("TTG", "")

	assert.Equal(false, IsRunningInside())
}

func TestGetppid(t *testing.T) {
	assert := assert.New(t)

	cases := map[string]int{
		"42424242": 42424242,
		"xyz":      0,
	}

	for input, expected := range cases {
		os.Setenv("TTG_PID", input)

		assert.Equal(expected, Getppid())
	}

	os.Setenv("TTG_PID", "")

	assert.Equal(0, Getppid())
}

func TestSpawn(t *testing.T) {
	assert := assert.New(t)

	cases := map[string]bool{
		"/bin/true":  true,
		"/bin/false": false,
	}

	for input, expected := range cases {
		os.Setenv("SHELL", input)

		assert.Equal(expected, Spawn("/") == nil)
	}
}
