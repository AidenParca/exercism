package isogram 
import "unicode"

func IsIsogram(word string) bool {
	seen := make(map[rune]bool)
	for _, char := range word {
		if !unicode.IsLetter(char) {
			continue
		}
		lChar := unicode.ToLower(char)
		if seen[lChar] {
			return false
		}
		seen[lChar] = true
	}
	return true
}