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

func manEscape(input []byte) []byte {
	result := bytes.NewBuffer(nil)

	for _, r := range string(input) {
		if val, has := manSpecialRunes[r]; has {
			result.WriteString(val)
		} else {
			result.WriteRune(r)
		}
	}

	return result.Bytes()
}

func manQuoteEscape(input []byte) []byte {
	result := bytes.NewBuffer(nil)

	for _, r := range string(input) {
		if val, has := manQuoteRunes[r]; has {
			result.WriteString(val)
		} else {
			result.WriteRune(r)
		}
	}

	return result.Bytes()
}
