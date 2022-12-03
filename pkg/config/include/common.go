package include

import (
	"os"
)

// FileOrSymlink tells us if dir entry in question is a regular file or a symlink
func FileOrSymlink(d os.DirEntry) bool {
	return d.Type().IsRegular() || d.Type() == os.ModeSymlink
}
