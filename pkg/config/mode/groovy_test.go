package mode

import (
	"testing"
)

func TestIsGroovy(t *testing.T) {
	runSuite(t, IsGroovy, "groovy/script.groovy")
}
