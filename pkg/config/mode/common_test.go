package mode

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"golang.org/x/exp/slices"
)

const (
	fixturePath = "../../../fixtures/config/mode"
)

func getFnName(fn func(os.DirEntry) bool) string {
	parts := strings.Split(runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(), ".")

	return parts[len(parts)-1] + "()"
}

func runSuite(t *testing.T, fn func(os.DirEntry) bool, paths ...string) {
	assert := assert.New(t)

	for i, p := range paths {
		paths[i] = fixturePath + "/" + p
	}

	for _, p := range paths {
		_, err := os.Stat(p)

		assert.Nilf(err, "file \"%s\" should exist and be accessible", p)
	}

	fnName := getFnName(fn)

	err := filepath.WalkDir(fixturePath,
		func(path string, d os.DirEntry, err error) error {
			assert.Nil(err)

			for _, p := range paths {
				if p == path {
					assert.Truef(fn(d), "func \"%+s\" is expected to include this path: %s", fnName, p)
				} else {
					if !slices.Contains(paths, path) {
						assert.Falsef(fn(d), "func \"%+s\" is expected to exclude this path: %s", fnName, p)
					}
				}
			}

			return nil
		})

	assert.Nil(err)
}
