package lib

import (
	"errors"
	"os"
	"path/filepath"
)

func CheckDirectory(path string) error {
	fileInfo, error := os.Stat(path)
	if os.IsNotExist(error) {
		error = os.MkdirAll(path, os.ModePerm)
		if error != nil {
			return error
		}
	} else if !fileInfo.IsDir() {
		return errors.New("not a directory")
	}

	return nil
}

func ClearDirectory(path string) error {
	dir, error := os.ReadDir(path)
	if error != nil {
		return error
	}

	for _, d := range dir {
		error = os.RemoveAll(filepath.Join([]string{path, d.Name()}...))
		if error != nil {
			return error
		}
	}

	return nil
}
