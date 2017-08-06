package main

import (
	"flag"
	"fmt"
	htmlt "html/template"
	"io"
	"io/ioutil"
	"os"
	textt "text/template"

	"github.com/AstromechZA/godork"
)

func OutputModeTemplate(pkg *godork.PackageDoc, w io.Writer) error {

	// parse some options from the command line
	fs := flag.NewFlagSet("t", flag.ExitOnError)
	templateFileFlag := fs.String("template", "", "Comma separated list of paths to Golang style templates for interpreting the structured object")
	ignoreErrFlag := fs.Bool("missingkey-ignore", false, "Ignore missingkey errors in the template")
	htmlModeFlag := fs.Bool("html-mode", false, "Interpret the template in html mode and use escape characters")
	if err := fs.Parse(os.Args[1:]); err != nil {
		return err
	}

	// read list of file paths
	if *templateFileFlag == "" {
		return fmt.Errorf("--template option is required")
	}

	mk := "missingkey=error"
	if *ignoreErrFlag {
		mk = "missingkey=default"
	}

	tcontent, err := ioutil.ReadFile(*templateFileFlag)
	if err != nil {
		return err
	}

	if *htmlModeFlag {
		t := htmlt.New("")
		t = t.Funcs(map[string]interface{}{
			"to_text":        ToText,
			"to_html":        ToHTML,
			"highlight_html": HighlightHTML,
		})
		t, terr := t.Parse(string(tcontent))
		if terr != nil {
			return terr
		}
		return t.Option(mk).Execute(w, pkg)
	}
	t := textt.New("")
	t = t.Funcs(map[string]interface{}{
		"to_text": ToText,
		"to_html": ToHTML,
	})
	t, terr := t.Parse(string(tcontent))
	if terr != nil {
		return terr
	}
	return t.Option(mk).Execute(w, pkg)
}
