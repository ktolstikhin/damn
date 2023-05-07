package util

import "unicode"

func ToSentence(s string) string {
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])

	return string(runes) + "."
}
