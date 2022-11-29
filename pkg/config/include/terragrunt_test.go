package include

import (
	"testing"
)

func TestIsTerragrunt(t *testing.T) {
	runSuite(t, IsTerragrunt, "terragrunt/terragrunt.hcl")
}
