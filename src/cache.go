package main

import (
	"os"
	"packman/common"
	"packman/lib"
	"path/filepath"
)

const cacheFolder = ".packman-cache"

func GetCacheDir() string {
	cacheDir, error := os.UserCacheDir()

	appCacheDir := filepath.Join(cacheDir, cacheFolder)
	lib.CheckDirectory(appCacheDir)

	common.ExitOnError("Cannot determine cache directory", error)
	return cacheDir
}
