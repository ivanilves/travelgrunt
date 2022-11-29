package include

import (
	"os"
)

// IsTerragrunt tells us if we operate on Terragrunt config file
func IsTerragrunt(d os.DirEntry) bool {
	return fileOrSymlink(d) && d.Name() == "terragrunt.hcl"
}
