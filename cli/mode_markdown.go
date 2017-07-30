package main

import (
	"io"
	"text/template"

	"github.com/AstromechZA/godork"
)

const markdownTemplate = `
# package ` + "`" + `{{ .PackageName }}` + "`" + `

` + "```" + `golang
import "{{ .ImportPath }}"
` + "```" + `

{{ if .Doc }}{{ .Doc }}{{ end }}

{{ if .Examples }}## Examples
{{ range .Examples }}
### ` + "`" + `{{ .FullName }}` + "`" + `

{{ if .Doc }}{{ .Doc }}{{ end }}
` + "```" + `golang
{{ .Code }}
` + "```" + `
{{ end }}{{ end }}

{{ if .Constants }}## Constants
{{ range .Constants }}
{{ if .Doc }}{{ .Doc }}{{ end }}
` + "```" + `golang
{{ .Code }}
` + "```" + `
{{end }}{{ end }}

{{ if .Variables }}## Variables
{{ range .Variables }}
{{ if .Doc }}{{ .Doc }}{{ end }}
` + "```" + `golang
{{ .Code }}
` + "```" + `
{{end }}{{ end }}

{{ if .Functions }}## Functions
{{ range .Functions }}
### ` + "`" + `{{ .Signature }}` + "`" + `

{{ if .Doc }}{{ .Doc }}{{ end }}
{{ if .Examples }}{{ range .Examples }}
#### ` + "`" + `{{ .FullName }}` + "`" + `

{{ if .Doc }}{{ .Doc }}{{ end }}
` + "```" + `golang
{{ .Code }}
` + "```" + `
{{ end }}{{ end }}

{{ end }}{{ end }}
`

func OutputModeMarkdown(pkg *godork.PackageDoc, w io.Writer) error {
	t, err := template.New("").Parse(markdownTemplate)
	if err != nil {
		return err
	}
	return t.Execute(w, pkg)
}
