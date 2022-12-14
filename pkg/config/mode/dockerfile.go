package mode

import (
	"os"
	"strings"
)

// IsDockerfile tells us if we operate on Dockerfile(s) or Dockerfile template(s)
func IsDockerfile(d os.DirEntry) bool {
	return FileOrSymlink(d) && strings.Contains(strings.ToLower(d.Name()), "dockerfile")
}
