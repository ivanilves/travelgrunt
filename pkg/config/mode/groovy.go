package mode

import (
	"os"
	"strings"
)

// IsGroovy tells us we operate on Groovy file(s)
func IsGroovy(d os.DirEntry) bool {
	return FileOrSymlink(d) && strings.HasSuffix(d.Name(), ".groovy")
}
