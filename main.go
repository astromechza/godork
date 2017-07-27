package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Version variable embedded by build tool
var Version = "unknown"

// GitCommitSHA variable embedded by build tool
var GitCommitSHA = "unknown"

// GitCommitTime variable embedded by build tool
var GitCommitTime = "1970-01-01T00:00:00Z"

const usageString = `
godork generates friendly html Golang package documentation along the lines of godoc but aimed at static interlinked
html hosted in your own web server.
`

func mainInner() error {
	versionFlag := flag.Bool("version", false, "print version information")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, strings.TrimSpace(usageString)+"\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *versionFlag {
		fmt.Fprintf(os.Stderr, "Github:     https://github.com/AstromechZA/godork\n")
		fmt.Fprintf(os.Stderr, "Version:    %s\n", Version)
		fmt.Fprintf(os.Stderr, "Git Commit: %s\n", GitCommitSHA)
		fmt.Fprintf(os.Stderr, "Git Date:   %s\n", GitCommitTime)
		os.Exit(1)
	}

	return nil
}

func main() {
	if err := mainInner(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
}
