package mode

import (
	"os"
)

// IsTerragrunt tells us if we operate on Terragrunt config file
func IsTerragrunt(d os.DirEntry) bool {
	return FileOrSymlink(d) && d.Name() == "terragrunt.hcl"
}
