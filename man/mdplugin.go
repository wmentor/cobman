package man

import (
	"bytes"
	"fmt"
	"io"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
)

var (
	_ markdown.Renderer = (*Plugin)(nil)
)

type Plugin struct {
	listLevel int
	lastText  bool
}

func NewPlugin() *Plugin {
	return &Plugin{
		listLevel: -1,
	}
}

func (plugin *Plugin) RenderFooter(w io.Writer, ast ast.Node) {
}

func (plugin *Plugin) RenderHeader(w io.Writer, ast ast.Node) {
}

func (plugin *Plugin) RenderNode(w io.Writer, node ast.Node, entering bool) ast.WalkStatus {
	switch node := node.(type) {
	case *ast.Text:
		w.Write(escapeBytes(node.Literal))
		plugin.lastText = true
	case *ast.Softbreak:
	case *ast.Hardbreak:
	case *ast.NonBlockingSpace:
	case *ast.Emph:
		if entering {
			w.Write([]byte("\\fI"))
		} else {
			w.Write([]byte("\\fR"))
		}
		plugin.lastText = true
	case *ast.Strong:
		if entering {
			w.Write([]byte("\\fB"))
		} else {
			w.Write([]byte("\\fR"))
		}
		plugin.lastText = true
	case *ast.Del:
	case *ast.BlockQuote:
	case *ast.Aside:
	case *ast.Link:
		w.Write(escapeBytes(node.Literal))
		plugin.lastText = true
	case *ast.CrossReference:
	case *ast.Citation:
	case *ast.Image:
		return ast.SkipChildren
	case *ast.Code:
		w.Write([]byte("\\fB"))
		w.Write(escapeBytes(node.Literal))
		w.Write([]byte("\\fR"))
		plugin.lastText = true
	case *ast.CodeBlock:
		plugin.pushItem(w, ".PP")
		w.Write([]byte("\\fB"))
		w.Write(escapeBytes(bytes.TrimSpace(node.Literal)))
		w.Write([]byte("\\fR"))
		plugin.lastText = true
	case *ast.Caption:
	case *ast.CaptionFigure:
	case *ast.Document:
	case *ast.Paragraph:
		if entering {
			plugin.pushItem(w, ".PP")
		}
	case *ast.HTMLSpan:
	case *ast.HTMLBlock:
	case *ast.Heading:
	case *ast.HorizontalRule:
	case *ast.List:
		if entering {
			plugin.listLevel++
		} else {
			plugin.listLevel--
		}
	case *ast.ListItem:
		str := fmt.Sprintf(".RS %d", plugin.listLevel*4)
		if !entering {
			str = ".RE"
		}
		plugin.pushItem(w, str)
	case *ast.Table:
		return ast.SkipChildren
	case *ast.TableCell:
	case *ast.TableHeader:
	case *ast.TableBody:
	case *ast.TableRow:
	case *ast.TableFooter:
	case *ast.Math:
	case *ast.MathBlock:
	case *ast.DocumentMatter:
	case *ast.Callout:
	case *ast.Index:
	case *ast.Subscript:
		w.Write(escapeBytes(node.Literal))
		plugin.lastText = true
	case *ast.Superscript:
		w.Write(escapeBytes(node.Literal))
		plugin.lastText = true
	case *ast.Footnotes:
	default:
	}

	return ast.GoToNext
}

func (plugin *Plugin) pushItem(w io.Writer, txt string) {
	if plugin.lastText {
		plugin.lastText = false
		w.Write([]byte{'\n'})
	}
	w.Write([]byte(txt))
	w.Write([]byte{'\n'})
}
