package include

import (
	"os"
	"path/filepath"

	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	fixturePath = "../../../fixtures/config/include"
)

func TestIncludeFn(t *testing.T) {
	assert := assert.New(t)

	type testCase struct {
		IsTerragrunt            bool
		IsTerraform             bool
		IsTerraformOrTerragrunt bool
	}

	testCases := map[string]testCase{
		"../../../fixtures/config/include/terragrunt/terragrunt.hcl": testCase{true, false, true},
		"../../../fixtures/config/include/terraform/main.tf":         testCase{false, true, true},
		"../../../fixtures/config/include/nothing/foo.bar":           testCase{false, false, false},
	}

	err := filepath.WalkDir(fixturePath,
		func(path string, d os.DirEntry, err error) error {
			assert.Nil(err)

			for p, expected := range testCases {
				if p == path {
					assert.Equal(expected.IsTerragrunt, IsTerragrunt(d))
					assert.Equal(expected.IsTerraform, IsTerraform(d))
					assert.Equal(expected.IsTerraformOrTerragrunt, IsTerraformOrTerragrunt(d))
				}
			}

			return nil
		})

	assert.Nil(err)
}
