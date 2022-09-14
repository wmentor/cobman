package man

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

func Md2Man(markdownData string) string {
	parser := parser.NewWithExtensions(parser.CommonExtensions)

	doc := markdown.Parse([]byte(markdownData), parser)
	render := NewPlugin()

	result := markdown.Render(doc, render)
	return string(result)
}
