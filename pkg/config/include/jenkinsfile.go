package include

import (
	"os"
	"strings"
)

// IsJenkinsfile tells us we operate on Jenkinsfile(s)
func IsJenkinsfile(d os.DirEntry) bool {
	return fileOrSymlink(d) && strings.Contains(strings.ToLower(d.Name()), "jenkinsfile")
}
