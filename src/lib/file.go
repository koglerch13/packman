package lib

import (
	"errors"
	"os"
)

func FileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	if errors.Is(error, os.ErrNotExist) {
		return false
	}

	if error == nil {
		return true
	}

	return false
}
