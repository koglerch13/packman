package main

import (
	"fmt"
	"packman/common"
	"packman/lib"
	"path/filepath"
	"strings"
)

func Load(configFilePath string) {
	cacheDirectory := GetCacheDir()

	packages := getPackagesToLoad(configFilePath)

	checkCacheDirectory(cacheDirectory)

	for _, currentPackage := range packages {
		var loadedPackage = loadPackage(currentPackage, cacheDirectory)
		extractPackage(loadedPackage, currentPackage.Destination, currentPackage.ClearDestination)
	}
}

func getPackagesToLoad(configFilePath string) []lib.Package {
	common.OutputAlways("Looking for config file: '%s'.", configFilePath)
	config, error := lib.LoadConfig(configFilePath)
	common.ExitOnError("Cannot load config file.", error)
	return config.Packages
}

func checkCacheDirectory(cacheDirectory string) {
	common.OutputVerbose("Check cache directory '%s'", cacheDirectory)
	error := lib.CheckDirectory(cacheDirectory)
	common.ExitOnError("Error checking cache directory.", error)
}

func loadPackage(packageToLoad lib.Package, cacheDirectory string) string {
	common.OutputAlways("Loading package '%s'.", packageToLoad.Uri)
	var cachedFileName = lib.ComputeHashFromString(packageToLoad.Uri)
	var cachedFilePath = filepath.Join(cacheDirectory, cachedFileName)

	common.OutputVerbose("Checking cache.")
	if doesValidCachedFileExist(cachedFilePath, packageToLoad.Hash) {
		// package was found in cache. skip the loading.
		return cachedFilePath
	}

	common.OutputVerbose("Loading package directly from source.")
	error := lib.LoadPackage(packageToLoad.Uri, cachedFilePath)
	common.ExitOnError("Cannot load package.", error)
	common.OutputVerbose("Loaded package.")

	if packageToLoad.Hash != "" {
		common.OutputVerbose("Verifying hash.")
		computedHash, error := lib.ComputeHashForFile(cachedFilePath)
		common.ExitOnError("Cannot compute hash.", error)

		if !strings.EqualFold(packageToLoad.Hash, computedHash) {
			var errorMessage = fmt.Sprintf("Hashes do not match. (Expected: '%s', Actual: '%s')", packageToLoad.Hash, computedHash)
			common.PrintErrorAndExit(errorMessage)
		}
	}

	return cachedFilePath
}

func extractPackage(loadedPackagePath string, destinationDirectoryPath string, clearDestination bool) {
	common.OutputVerbose("Extract package to '%s'.", destinationDirectoryPath)

	error := lib.CheckDirectory(destinationDirectoryPath)
	common.ExitOnError("Cannot check destionation directory.", error)

	if clearDestination {
		common.OutputAlways("Clearing directory '%s'.", destinationDirectoryPath)
		error := lib.ClearDirectory(destinationDirectoryPath)
		common.ExitOnError("Cannot clear directory.", error)
	}

	common.OutputVerbose("Extract package to '%s'.", destinationDirectoryPath)
	error = lib.Unzip(loadedPackagePath, destinationDirectoryPath)
	common.ExitOnError("Error when extracting.", error)
	common.OutputAlways("Extracted package to '%s'.", destinationDirectoryPath)
}

func doesValidCachedFileExist(cachedFilePath string, expectedHash string) bool {
	if !lib.FileExists(cachedFilePath) {
		return false
	}

	if expectedHash == "" {
		return true
	}

	computedHash, error := lib.ComputeHashForFile(cachedFilePath)
	if error != nil {
		common.OutputVerbose("Found package in cache, but it appears to be invalid.")
		return false
	}

	if strings.EqualFold(expectedHash, computedHash) {
		common.OutputVerbose("Found package in cache.")
		return true
	}

	common.OutputVerbose("Found package in cache, but it appears to be invalid.")
	return false
}
