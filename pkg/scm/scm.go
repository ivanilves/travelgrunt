package scm

import (
	"fmt"

	"github.com/ivanilves/travelgrunt/pkg/scm/git"
)

func detectSCM() string {
	if git.Probe() {
		return "git"
	}

	return "unknown"
}

// RootPath gets top level path of the SCM repository of choice
func RootPath() (string, error) {
	switch scm := detectSCM(); scm {
	case "git":
		return git.RootPath()
	default:
		return "", fmt.Errorf("unknown SCM: %s", scm)
	}
}
