package links

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	invalidPath = "/c2RmZ3JlZ2ZncnR3Z3I0aGd3dGd3d3d3d3d3d2RmZGZmZnJlZmdydDRndnQ0M3RodWprbDhpb2s4"
)

func TestExpandPath(t *testing.T) {
	assert := assert.New(t)

	testCases := map[string]bool{
		"/tmp":           true,  // absolute, no expansion
		"relative/path":  true,  // relative, no expansion
		"~":              true,  // home directory
		"~/projects":     true,  // home with path
		"~user/projects": false, // ~username not supported
	}

	for path, expectedSuccess := range testCases {
		result, err := ExpandPath(path)

		if expectedSuccess {
			assert.NotEmpty(result)
			assert.Nil(err)
		} else {
			assert.Empty(result)
			assert.NotNil(err)
		}
	}
}

func TestGetAbsPath(t *testing.T) {
	assert := assert.New(t)

	tmpDir := t.TempDir()

	testCases := map[string]struct {
		path            string
		rootPath        string
		expectedSuccess bool
	}{
		"absolute path": {
			path:            "/tmp",
			rootPath:        tmpDir,
			expectedSuccess: true,
		},
		"relative path": {
			path:            "build",
			rootPath:        tmpDir,
			expectedSuccess: true,
		},
		"empty path": {
			path:            "",
			rootPath:        tmpDir,
			expectedSuccess: false,
		},
		"tilde path": {
			path:            "~/projects",
			rootPath:        tmpDir,
			expectedSuccess: true,
		},
		"parent traversal": {
			path:            "../../etc",
			rootPath:        tmpDir,
			expectedSuccess: false, // Should be blocked
		},
	}

	for name, tc := range testCases {
		result, err := GetAbsPath(tc.path, tc.rootPath)

		if tc.expectedSuccess {
			assert.NotEmpty(result, "test case: %s", name)
			assert.Nil(err, "test case: %s", name)
			assert.True(filepath.IsAbs(result), "test case: %s should return absolute path", name)
		} else {
			assert.Empty(result, "test case: %s", name)
			assert.NotNil(err, "test case: %s", name)
		}
	}
}

func TestValidatePath(t *testing.T) {
	assert := assert.New(t)

	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "testfile.txt")

	err := os.WriteFile(tmpFile, []byte("test"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	testCases := map[string]bool{
		tmpDir:      true,  // valid directory
		tmpFile:     false, // file, not directory
		invalidPath: false, // non-existent
	}

	for path, expectedSuccess := range testCases {
		err := ValidatePath(path)

		if expectedSuccess {
			assert.Nil(err)
		} else {
			assert.NotNil(err)
		}
	}
}

func TestCollect(t *testing.T) {
	assert := assert.New(t)

	tmpDir := t.TempDir()
	subDir := filepath.Join(tmpDir, "subdir")

	err := os.Mkdir(subDir, 0755)
	if err != nil {
		t.Fatal(err)
	}

	testCases := map[string]struct {
		rootPath        string
		linkPaths       []string
		expectedSuccess bool
	}{
		"valid absolute and relative": {
			rootPath:        tmpDir,
			linkPaths:       []string{tmpDir, "subdir"},
			expectedSuccess: true,
		},
		"empty links": {
			rootPath:        tmpDir,
			linkPaths:       []string{},
			expectedSuccess: false,
		},
		"invalid path": {
			rootPath:        tmpDir,
			linkPaths:       []string{invalidPath},
			expectedSuccess: false,
		},
		"mixed valid and empty": {
			rootPath:        tmpDir,
			linkPaths:       []string{tmpDir, "", "subdir"},
			expectedSuccess: true,
		},
		"only empty strings": {
			rootPath:        tmpDir,
			linkPaths:       []string{"", ""},
			expectedSuccess: false,
		},
	}

	for name, tc := range testCases {
		entries, paths, err := Collect(tc.rootPath, tc.linkPaths)

		if tc.expectedSuccess {
			assert.Greater(len(entries), 0, "test case: %s", name)
			assert.Greater(len(paths), 0, "test case: %s", name)
			assert.Nil(err, "test case: %s", name)
		} else {
			assert.Equal(0, len(entries), "test case: %s", name)
			assert.Equal(0, len(paths), "test case: %s", name)
			assert.NotNil(err, "test case: %s", name)
		}
	}
}
