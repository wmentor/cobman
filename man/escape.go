package man

import (
	"bytes"
)

var (
	manSpecialRunes = map[rune]string{
		'\\': "\\\\",
		'.':  "\\&.",
		'-':  "\\-",
	}

	manQuoteRunes = map[rune]string{
		'\r': "",
		'\n': " ",
		'"':  "\\\"",
	}
)

func Escape(input string) string {
	result := bytes.NewBuffer(nil)

	for _, r := range string(input) {
		if val, has := manSpecialRunes[r]; has {
			result.WriteString(val)
		} else {
			result.WriteRune(r)
		}
	}

	return result.String()
}

func escapeBytes(input []byte) []byte {
	return []byte(Escape(string(input)))
}

func QuoteEscape(input string) string {
	result := bytes.NewBuffer(nil)

	for _, r := range string(input) {
		if val, has := manQuoteRunes[r]; has {
			result.WriteString(val)
		} else {
			result.WriteRune(r)
		}
	}

	return result.String()
}
