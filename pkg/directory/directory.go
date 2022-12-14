package directory

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/ivanilves/travelgrunt/pkg/config"
)

func isHiddenDir(d os.DirEntry) bool {
	return d.IsDir() && string(d.Name()[0]) == "."
}

func isInScope(absPath string, rootPath string) bool {
	return len(absPath) >= len(rootPath)
}

// Collect gets a list of directory path entries containing file "terragrunt.hcl"
func Collect(rootPath string, cfg config.Config) (entries map[string]string, paths []string, err error) {
	entries = make(map[string]string, 0)
	paths = make([]string, 0)

	err = filepath.WalkDir(rootPath,
		func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if isHiddenDir(d) {
				return filepath.SkipDir
			}

			absPath := filepath.Dir(path)

			if isInScope(absPath, rootPath) {
				relPath := strings.TrimPrefix(absPath, rootPath+"/")

				if relPath == absPath {
					relPath = "."
				}

				if cfg.Admit(d, relPath) {
					entries[relPath] = absPath
					paths = append(paths, relPath)
				}
			}

			return nil
		})

	return entries, paths, err
}
