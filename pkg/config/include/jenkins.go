package include

import (
	"os"
)

// IsJenkins tells us we operate on Jenkinsfile(s) or Groovy file(s)
func IsJenkins(d os.DirEntry) bool {
	return IsJenkinsfile(d) || IsGroovy(d)
}
