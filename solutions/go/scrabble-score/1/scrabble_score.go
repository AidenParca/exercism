package scrabble

import "strings"
import "unicode"

// scoreMap holds the pre-calculated score for each lowercase letter.
var scoreMap = make(map[rune]int)

func init() {
	scoreGroups := map[int]string{
		1:  "AEIOULNRST",
		2:  "DG",
		3:  "BCMP",
		4:  "FHVWY",
		5:  "K",
		8:  "JX",
		10: "QZ",
	}
	for score, letters := range scoreGroups {
		for _, letter := range letters {
			scoreMap[unicode.ToLower(letter)] = score
		}
	}
}
// Score calculates the Scrabble score for a given word.
func Score(word string) int {
	total := 0
	// Loop over each character of the input word.
	for _, char := range strings.ToLower(word) {
		// Look up the character's score in the map and add it to the total.
		total += scoreMap[char]
	}
	return total
}