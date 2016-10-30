package main

import (
	"fmt"
	"os"

	"github.com/Vilyan01/vowels/version"
)

func main() {
	// get args
	if len(os.Args) <= 1 {
		fmt.Println("Usage: vowels <args>")
		os.Exit(1)
	}
	if ArgsContainsFlag("-v") {
		version.PrintVersion()
	}
}

func ArgsContainsFlag(f string) bool {
	for i := range os.Args {
		if os.Args[i] == f {
			return true
		}
	}
	return false
}
