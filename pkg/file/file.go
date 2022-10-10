package file

import (
	"os"
)

func WriteAndExit(outFile string, path string) error {
	err := os.WriteFile(outFile, []byte(path), 0644)

	if err == nil {
		os.Exit(0)
	}

	return err
}
