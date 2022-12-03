package directory

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/ivanilves/travelgrunt/pkg/config"
)

func isHidden(d os.DirEntry) bool {
	return d.IsDir() && string(d.Name()[0]) == "."
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

			if isHidden(d) {
				return filepath.SkipDir
			}

			if cfg.IncludeFn(d) {
				abs := filepath.Dir(path)
				rel := strings.TrimPrefix(abs, rootPath+"/")

				entries[rel] = abs
				paths = append(paths, rel)
			}

			return nil
		})

	return entries, paths, err
}
