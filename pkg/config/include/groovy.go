package include

import (
	"os"
	"strings"
)

// IsGroovy tells us we operate on Groovy file(s)
func IsGroovy(d os.DirEntry) bool {
	return fileOrSymlink(d) && strings.HasSuffix(d.Name(), ".groovy")
}
