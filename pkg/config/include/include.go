package include

import (
	"os"
	"strings"
)

func fileOrSymlink(d os.DirEntry) bool {
	return d.Type().IsRegular() || d.Type() == os.ModeSymlink
}

// IsTerragrunt tells us if we operate on Terragrunt config file
func IsTerragrunt(d os.DirEntry) bool {
	return fileOrSymlink(d) && d.Name() == "terragrunt.hcl"
}

// IsTerraform tells us if we operate on Terraform file(s)
func IsTerraform(d os.DirEntry) bool {
	return fileOrSymlink(d) && strings.HasSuffix(d.Name(), ".tf")
}

// IsTerraformOrTerragrunt tells us if we operate on Terraform or Terragrunt file(s)
func IsTerraformOrTerragrunt(d os.DirEntry) bool {
	return IsTerraform(d) || IsTerragrunt(d)
}

// IsDockerfile tells us if we operate on Dockerfile(s) or Dockerfile template(s)
func IsDockerfile(d os.DirEntry) bool {
	return fileOrSymlink(d) && strings.Contains(strings.ToLower(d.Name()), "dockerfile")
}

// IsJenkinsfile tells us we operate on Jenkinsfile(s)
func IsJenkinsfile(d os.DirEntry) bool {
	return fileOrSymlink(d) && strings.Contains(strings.ToLower(d.Name()), "jenkinsfile")
}

// IsGroovy tells us we operate on Groovy file(s)
func IsGroovy(d os.DirEntry) bool {
	return fileOrSymlink(d) && strings.HasSuffix(d.Name(), ".groovy")
}

// IsJenkins tells us we operate on Jenkinsfile(s) or Groovy file(s)
func IsJenkins(d os.DirEntry) bool {
	return IsJenkinsfile(d) || IsGroovy(d)
}
