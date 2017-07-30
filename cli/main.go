package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/AstromechZA/godork"
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

Usage:

$ godork [MODE] [SRC DIRECTORY] [IMPORT PATH] (options)

Mode should be one of the following:

	json 		Dump the package documentation data as JSON on stdout
	markdown 	Print the documentation as a Markdown file
	html 		Print the documentation as a HTML file
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

	if len(os.Args) != 4 {
		return fmt.Errorf("expected 3 positional arguments: [MODE] [SRC DIRECTORY] [IMPORT PATH] (options)")
	}
	mode := os.Args[1]
	sourceDirectory := os.Args[2]
	importPath := os.Args[3]

	sourceDirectory, err := filepath.Abs(sourceDirectory)
	if err != nil {
		return fmt.Errorf("failed to get absolute path to '%s': %s", sourceDirectory, err)
	}

	fi, err := os.Stat(sourceDirectory)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("source directory '%s' does not exist", sourceDirectory)
		}
		return fmt.Errorf("failed to read source directory '%s': %s", sourceDirectory, err)
	}
	if !fi.IsDir() {
		return fmt.Errorf("source directory must be a directory, not a file")
	}

	if importPath == "" {
		return fmt.Errorf("import path cannot be empty string")
	}
	for _, p := range strings.Split(importPath, "/") {
		p = strings.TrimSpace(p)
		if p == "" {
			return fmt.Errorf("import path is invalid")
		}
	}

	fset, pkg, err := godork.ReadPackageDirectory(sourceDirectory)
	if err != nil {
		return err
	}
	packageDoc, err := godork.BuildPackageDoc(importPath, fset, pkg)
	if err != nil {
		return err
	}

	newArgs := []string{os.Args[0]}
	newArgs = append(newArgs, os.Args[4:]...)
	os.Args = newArgs
	switch mode {
	case "json":
		return OutputModeJSON(packageDoc, os.Stdout)
	case "markdown":
		return OutputModeMarkdown(packageDoc, os.Stdout)
	case "html":
		return fmt.Errorf("output mode 'html' is not implemented")
	case "template":
		return fmt.Errorf("output mode 'template' is not implemented")
	default:
		return fmt.Errorf("unknown output mode '%s' see --help", mode)
	}
}

func main() {
	if err := mainInner(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
}
