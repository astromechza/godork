package main

import (
	"bytes"
	"go/doc"
	"html/template"

	"github.com/sourcegraph/syntaxhighlight"
)

func ToText(text string, width int) string {
	buf := new(bytes.Buffer)
	doc.ToText(buf, text, "", "", width)
	return buf.String()
}

func ToHTML(text string) template.HTML {
	buf := new(bytes.Buffer)
	doc.ToHTML(buf, text, map[string]string{})
	return template.HTML(buf.String())
}

func HighlightHTML(code string) template.HTML {
	output, err := syntaxhighlight.AsHTML([]byte(code))
	if err != nil {
		panic(err)
	}
	return template.HTML(output)
}
