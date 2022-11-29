package include

import (
	"testing"
)

func TestIsTerraformOrTerragrunt(t *testing.T) {
	runSuite(t, IsTerraformOrTerragrunt, "terraform/main.tf", "terragrunt/terragrunt.hcl")
}
