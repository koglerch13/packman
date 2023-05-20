package common

import (
	"fmt"
	"os"
)

func ExitOnError(message string, e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "%s\n%v\n", message, e)
		os.Exit(1)
	}
}

func PrintErrorAndExit(message string) {
	fmt.Fprint(os.Stderr, message)
	os.Exit(1)
}
