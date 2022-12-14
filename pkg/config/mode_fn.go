package config

import (
	"fmt"
	"os"

	"github.com/ivanilves/travelgrunt/pkg/config/mode"
)

func getModeFn(m string) (fn func(os.DirEntry) bool, err error) {
	err = nil

	switch m {
	case "":
		fn = nil
	case "terragrunt":
		fn = mode.IsTerragrunt
	case "terraform":
		fn = mode.IsTerraform
	case "dockerfile":
		fn = mode.IsDockerfile
	case "jenkinsfile":
		fn = mode.IsJenkinsfile
	case "groovy":
		fn = mode.IsGroovy
	default:
		fn = nil
		err = fmt.Errorf("illegal mode: %s", m)
	}

	return fn, err
}
