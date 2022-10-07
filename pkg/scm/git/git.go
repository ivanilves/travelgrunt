package git

import (
	"os"

	git "github.com/go-git/go-git/v5"
)

// Probe tells you if you are in the Git repository or not
func Probe() bool {
	_, err := RootPath()

	return err == nil
}

// RootPath gets top level path of the Git repository
func RootPath() (string, error) {
	var err error

	cwd := "."

	for {
		_, err = os.Open(cwd)

		if err != nil {
			return "", err
		}

		_, err = git.PlainOpen(cwd)

		if err == nil {
			err = os.Chdir(cwd)

			if err != nil {
				return "", err
			}

			return os.Getwd()
		}

		cwd = cwd + "/.."
	}
}
