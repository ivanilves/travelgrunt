package include

import (
	"testing"
)

func TestIsJenkins(t *testing.T) {
	runSuite(t, IsJenkins, "jenkinsfile/Jenkinsfile", "groovy/script.groovy")
}
