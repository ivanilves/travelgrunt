package include

import (
	"testing"
)

func TestIsDockerfile(t *testing.T) {
	runSuite(t, IsDockerfile, "dockerfile/Dockerfile")
}
