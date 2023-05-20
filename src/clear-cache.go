package main

import (
	"packman/common"
	"packman/lib"
)

func ClearCache() {
	cacheDirectory := GetCacheDir()
	common.OutputAlways("Clear cache: '%s'.", cacheDirectory)
	error := lib.ClearDirectory(cacheDirectory)
	common.ExitOnError("Cannot clear cache", error)
}
