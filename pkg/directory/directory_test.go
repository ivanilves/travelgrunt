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

func TestCollect(t *testing.T) {
	assert := assert.New(t)

	cfg := config.DefaultConfig()

	testCases := map[string]bool{
		fixturePath: true,
		invalidPath: false,
	}

	for path, expectedSuccess := range testCases {
		entries, paths, err := Collect(path, cfg.IncludeFn)

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
