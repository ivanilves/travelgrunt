package directory

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ivanilves/travelgrunt/pkg/config"
)

const (
	fixturePath = "../../fixtures/directory"
	invalidPath = "/c2RmZ3JlZ2ZncnR3Z3I0aGd3dGd3d3d3d3d3d2RmZGZmZnJlZmdydDRndnQ0M3RodWprbDhpb2s4"
)

func TestGetAbsPath(t *testing.T) {
	assert := assert.New(t)

	const testPath = "/etc/passwd"

	assert.Equal(getAbsPath(testPath, false), "/etc")
	assert.Equal(getAbsPath(testPath, true), "/etc/passwd")
}

func TestCollect(t *testing.T) {
	assert := assert.New(t)

	cfg := config.DefaultConfig()

	testCases := map[string]bool{
		fixturePath: true,
		invalidPath: false,
	}

	for path, expectedSuccess := range testCases {
		entries, paths, err := Collect(path, cfg)

		if expectedSuccess {
			assert.Greater(len(entries), 0)
			assert.Greater(len(paths), 0)
			assert.Nil(err)
		} else {
			assert.Equal(0, len(entries))
			assert.Equal(0, len(paths))
			assert.NotNil(err)
		}
	}
}
