package file

import (
	"os"
)

// Write writes a file with the content specified
func Write(outFile string, content string) error {
	err := os.WriteFile(outFile, []byte(content), 0644)

	return err
}
