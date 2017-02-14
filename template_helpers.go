package main

import (
	"fmt"
	"github.com/h-tko/blog/libraries"
	"html/template"
	"time"
)

var TemplateHelpers = template.FuncMap{
	"raw":      htmlRaw,
	"mbsubstr": mbsubstr,
	"datestr":  datestr,
	"m2h":      m2h,
}

func htmlRaw(html string) template.HTML {
	return template.HTML(html)
}

func mbsubstr(text string, from int, to int) string {

	rntext := []rune(text)

	if len(rntext) <= to {
		return text
	}

	return string(rntext[from:to])
}

func datestr(target time.Time) string {
	return target.Format("2006/01/02")
}

func m2h(text string) string {
	return fmt.Sprintf("%s", libraries.MarkdownToHtml([]byte(text)))
}
