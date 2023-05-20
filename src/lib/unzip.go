package lib

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Unzip(source string, destination string) error {
	// shamelessly taken from: https://stackoverflow.com/a/24792688

	reader, openReaderError := zip.OpenReader(source)
	if openReaderError != nil {
		return openReaderError
	}
	defer func() {
		if readerCloseError := reader.Close(); readerCloseError != nil {
			panic(readerCloseError)
		}
	}()

	os.MkdirAll(destination, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(zipFile *zip.File) error {
		readCloser, openError := zipFile.Open()
		if openError != nil {
			return openError
		}
		defer func() {
			if err := readCloser.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(destination, zipFile.Name)

		// Check for ZipSlip (Directory traversal)
		if !strings.HasPrefix(path, filepath.Clean(destination)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		if zipFile.FileInfo().IsDir() {
			os.MkdirAll(path, zipFile.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), zipFile.Mode())
			file, openFileError := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, zipFile.Mode())
			if openFileError != nil {
				return openFileError
			}
			defer func() {
				if fileCloseError := file.Close(); fileCloseError != nil {
					panic(fileCloseError)
				}
			}()

			_, openFileError = io.Copy(file, readCloser)
			if openFileError != nil {
				return openFileError
			}
		}
		return nil
	}

	for _, zipFile := range reader.File {
		extractError := extractAndWriteFile(zipFile)
		if extractError != nil {
			return extractError
		}
	}

	return nil
}
