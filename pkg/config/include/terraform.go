package include

import (
	"os"
	"strings"
)

// IsTerraform tells us if we operate on Terraform file(s)
func IsTerraform(d os.DirEntry) bool {
	return fileOrSymlink(d) && strings.HasSuffix(d.Name(), ".tf")
}
