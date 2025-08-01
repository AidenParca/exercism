package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	// Iterate over each character (rune) in the input string.
	for _, char := range log {
		// Switch on the character to find a match.
		switch char {
		case '❗': // Unicode U+2757
			return "recommendation"
		case '🔍': // Unicode U+1F50D
			return "search"
		case '☀': // Unicode U+2600
			return "weather"
		}
	}
	return "default"
}


// Replace replaces all occurrences of old with new, returning the modified log
// to the caller without using the strings package.
func Replace(log string, oldRune, newRune rune) string {
	// Create a slice of runes to build the new string.
	// Pre-allocating with a capacity can be a small optimization.
	result := make([]rune, 0, len(log))

	for _, char := range log {
		if char == oldRune {
			// If we find the old rune, add the new one to our slice.
			result = append(result, newRune)
		} else {
			// Otherwise, add the original character.
			result = append(result, char)
		}
	}

	// Convert the slice of runes back into a string and return it.
	return string(result)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	Nor := utf8.RuneCountInString(log)
    if Nor <= limit {
        return true
    } else {
        return false
    }
}
