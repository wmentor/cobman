package cobman

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/wmentor/cobman/man"
)

func Md2Man(markdownData string) string {
	parser := parser.NewWithExtensions(parser.CommonExtensions)

	doc := markdown.Parse([]byte(markdownData), parser)
	render := man.NewPlugin()

	result := markdown.Render(doc, render)
	return string(result)
}
