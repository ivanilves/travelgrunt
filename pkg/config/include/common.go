package include

import (
	"os"
)

func fileOrSymlink(d os.DirEntry) bool {
	return d.Type().IsRegular() || d.Type() == os.ModeSymlink
}
