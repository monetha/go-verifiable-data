package cmd

import (
	"flag"
	"fmt"
)

var (
	// Version (set by compiler) is the version of program
	Version = "undefined"
	// BuildTime (set by compiler) is the program build time in '+%Y-%m-%dT%H:%M:%SZ' format
	BuildTime = "undefined"
	// GitHash (set by compiler) is the git commit hash of source tree
	GitHash = "undefined"

	printVersion bool
)

func init() {
	flag.BoolVar(&printVersion, "v", false, "outputs the binary version")
}

// PrintVersion outputs the binary version when command-line flag "-v" is set.
// It returns true if flag is set.
func PrintVersion() bool {
	if !flag.Parsed() {
		flag.Parse()
	}

	fmt.Printf("Version: %s\nBuildTime: %v\nGitHash: %s\n", Version, BuildTime, GitHash)

	return printVersion
}
