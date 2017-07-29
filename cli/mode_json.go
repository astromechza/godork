package main

import (
	"encoding/json"
	"io"

	"github.com/AstromechZA/godork"
)

func OutputModeJSON(pkg *godork.PackageDoc, w io.Writer) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	return enc.Encode(pkg)
}
