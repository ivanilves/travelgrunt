package include

import (
	"testing"
)

func TestIsJenkinsfile(t *testing.T) {
	runSuite(t, IsJenkinsfile, "jenkinsfile/Jenkinsfile")
}
