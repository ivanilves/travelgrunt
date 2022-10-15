package file

import (
	"os"
)

func Write(outFile string, path string) error {
	err := os.WriteFile(outFile, []byte(path), 0644)

	return err
}
