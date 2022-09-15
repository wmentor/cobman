package man_test

import (
	"testing"

	_ "embed"

	"github.com/stretchr/testify/require"
	"github.com/wmentor/cobman/man"
)

//go:embed list.txt
var listWait string

func TestMdPlugin_Paragraph(t *testing.T) {
	t.Parallel()

	txt := `Package github.com/gomarkdown/markdown is a Go **library** for parsing Markdown text and rendering as HTML.

It's very fast and supports common extensions.
`

	wait := `.PP
Package github\&.com/gomarkdown/markdown is a Go \fBlibrary\fR for parsing Markdown text and rendering as HTML\&.
.PP
It's very fast and supports common extensions\&.`

	res := man.Md2Man(txt)

	require.Equal(t, wait, res)
}

func TestMdPlugin_List(t *testing.T) {
	t.Parallel()

	txt := `1. First item
2. Second item
3. Third item
	1. Indented item
	2. Indented item
4. Fourth item
`
	res := man.Md2Man(txt)

	require.Equal(t, listWait, res)
}
