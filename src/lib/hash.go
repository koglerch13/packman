package lib

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

func ComputeHashFromString(value string) string {
	var computedHash = sha256.Sum256([]byte(value))
	var hashString = hex.EncodeToString(computedHash[:])
	return hashString
}

func ComputeHashForFile(filePath string) (string, error) {
	file, error := os.Open(filePath)
	if error != nil {
		return "", error
	}
	defer file.Close()

	var hash = sha256.New()
	_, error = io.Copy(hash, file)
	if error != nil {
		return "", error
	}

	var computedHash = hash.Sum(nil)
	var hashString = hex.EncodeToString(computedHash)
	return hashString, nil
}
