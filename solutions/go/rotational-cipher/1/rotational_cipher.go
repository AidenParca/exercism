package rotationalcipher

import (
	"strings"
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	abc := "abcdefghijklmnopqrstuvwxyz"
	Cipher := abc[shiftKey:] + abc[:shiftKey]
	var builder strings.Builder
	for _, r := range plain {
		if !unicode.IsLetter(r) {
			builder.WriteRune(r) 
			continue
		}
		isUpper := unicode.IsUpper(r)
		idx := strings.IndexRune(abc, unicode.ToLower(r))
		cipherChar := rune(Cipher[idx])
		if isUpper {
			builder.WriteRune(unicode.ToUpper(cipherChar))
		} else {
			builder.WriteRune(cipherChar)
		}
	}
	return builder.String()
}