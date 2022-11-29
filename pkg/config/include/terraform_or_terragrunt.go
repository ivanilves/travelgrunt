package include

import (
	"os"
)

// IsTerraformOrTerragrunt tells us if we operate on Terraform or Terragrunt file(s)
func IsTerraformOrTerragrunt(d os.DirEntry) bool {
	return IsTerraform(d) || IsTerragrunt(d)
}
