package include

import (
	"testing"
)

func TestIsTerraform(t *testing.T) {
	runSuite(t, IsTerraform, "terraform/main.tf")
}
