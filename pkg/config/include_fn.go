package config

import (
	"fmt"
	"os"

	"github.com/ivanilves/travelgrunt/pkg/config/include"
)

func getIncludeFn(mode string) (fn func(os.DirEntry) bool, err error) {
	err = nil

	switch mode {
	case "":
		fn = nil
	case "terragrunt":
		fn = include.IsTerragrunt
	case "terraform":
		fn = include.IsTerraform
	case "dockerfile":
		fn = include.IsDockerfile
	case "jenkinsfile":
		fn = include.IsJenkinsfile
	case "groovy":
		fn = include.IsGroovy
	default:
		fn = nil
		err = fmt.Errorf("illegal mode: %s", mode)
	}

	return fn, err
}
