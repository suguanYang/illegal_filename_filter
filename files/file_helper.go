package files

import (
	"os"
	"strings"
)

// FileCb deal filepath
type FileCb func(filePath string) string

// IsFileExists is a file exist
func IsFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

// IsHiddenFile is a hidden file or directory
func IsHiddenFile(fileName string) bool {
	return strings.Index(fileName, ".") == 0
}
