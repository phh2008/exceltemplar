package richtext

import (
	"testing"

	excel "github.com/nikitaxru/exceltemplar"
)

func TestRichtext(t *testing.T) {
	// Path to an .xlsx file containing template directives
	templatePath := "./files/template1.xlsx"
	outputPath := "./files/output1.xlsx"

	// One or more JSON strings providing data for rendering
	data := []string{
		`{"tasks":[{"name":"aaa","priority":"high"},{"name":"bbb","priority":"low"},{"name":"ccc","priority":"high"}],"project":{"name":"项目1","version":"v1.0.0","count":112}}`,
	}

	if err := excel.WriteResultsWithTemplate(templatePath, outputPath, data); err != nil {
		t.Fatalf("render failed: %v", err)
	}
}
