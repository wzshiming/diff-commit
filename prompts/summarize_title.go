package prompts

import (
	_ "embed"
	"bytes"
	"text/template"
)

//go:embed summarize_title.tpl
var summarizeTitleTemplate string

var summarizeTitle = template.Must(template.New("summarize_title").Parse(summarizeTitleTemplate))

func SummarizeTitle(summaryPoints string) string {
	buf := bytes.NewBuffer(nil)
	err := summarizeTitle.Execute(buf, map[string]any{
		"summaryPoints": summaryPoints,
	})
	if err != nil {
		panic(err)
	}
	return buf.String()
}
