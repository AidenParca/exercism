package wordcount

import (
	"regexp"
	"strings"
)


type Frequency map[string]int
func WordCount(phrase string) Frequency {
	lowerPhrase := strings.ToLower(phrase)
	re := regexp.MustCompile(`[a-z0-9]+(?:'[a-z0-9]+)*`)
	words := re.FindAllString(lowerPhrase, -1)
	counts := make(Frequency)

	for _, word := range words {
		counts[word]++
	}

	return counts
}