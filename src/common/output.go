package common

import "fmt"

var Verbose = true

func OutputAlways(format string, a ...any) {
	fmt.Printf(format+"\n", a...)
}

func OutputVerbose(format string, a ...any) {
	if Verbose {
		fmt.Printf(format+"\n", a...)
	}
}
