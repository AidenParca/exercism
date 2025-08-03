package pangram

import "unicode"

// IsPangram determines if a sentence is a pangram.
func IsPangram(input string) bool {
	seen := make(map[rune]bool)

	for _, char := range input {

		if !unicode.IsLetter(char) {
			continue
		}
		seen[unicode.ToLower(char)] = true
	}
	return len(seen) == 26
}