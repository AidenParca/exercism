package etl

import "strings"

// Transform converts a map of score-to-letters into a map of letter-to-score.
func Transform(in map[int][]string) map[string]int {
	out := make(map[string]int)

	//Loop through the input map to get the score and its corresponding letters.
	for score, letters := range in {
		for _, letter := range letters {
			out[strings.ToLower(letter)] = score
		}
	}

	return out
}