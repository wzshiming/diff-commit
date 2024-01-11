package prompts

import (
	_ "embed"
	"text/template"
	"bytes"
)

//go:embed conventional_commit.tpl
var conventionalCommitTemplate string

var conventionalCommit = template.Must(template.New("conventional_commit").Parse(conventionalCommitTemplate))

func ConventionalCommit(summaryPoints string) string {
	buf := bytes.NewBuffer(nil)
	err := conventionalCommit.Execute(buf, map[string]any{
		"summaryPoints": summaryPoints,
	})
	if err != nil {
		panic(err)
	}
	return buf.String()
}
