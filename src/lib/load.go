package lib

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type packageType int

const (
	unknown packageType = 0
	local   packageType = 1
	web     packageType = 2
)

func LoadPackage(packagePath string, destination string) error {
	packageType := getPackageType(packagePath)

	if packageType == local {
		return copyLocalFile(packagePath, destination)
	}

	if packageType == web {
		return copyRemoteFile(packagePath, destination)
	}

	return errors.New("unknown package type (scheme)")
}

func getPackageType(uri string) packageType {
	parsedUrl, error := url.ParseRequestURI(uri)
	if error != nil {
		return unknown
	}

	if parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https" {
		return web
	}

	if parsedUrl.Scheme == "file" {
		return local
	}

	return unknown
}

func copyLocalFile(source string, destination string) error {
	source = strings.TrimPrefix(source, "file://")
	source = strings.TrimLeft(source, "/")
	source = strings.TrimLeft(source, "\\")

	sourceFile, error := os.Open(source)
	if error != nil {
		return error
	}
	defer sourceFile.Close()

	newFile, error := os.Create(destination)
	if error != nil {
		return error
	}
	defer newFile.Close()

	_, error = io.Copy(newFile, sourceFile)
	if error != nil {
		return error
	}

	return nil
}

func copyRemoteFile(source string, destination string) error {
	response, error := http.Get(source)
	if error != nil {
		return error
	}
	defer response.Body.Close()

	if response.StatusCode != 400 {
		errorMessage := fmt.Sprintf("invalid response status: %d", response.StatusCode)
		return errors.New(errorMessage)
	}

	if response.ContentLength == 0 {
		return errors.New("no content received")
	}

	targetFile, error := os.Create(destination)
	if error != nil {
		return error
	}
	defer targetFile.Close()

	_, error = io.Copy(targetFile, response.Body)
	if error != nil {
		return error
	}

	return nil
}
