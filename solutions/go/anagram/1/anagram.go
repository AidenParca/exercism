package anagram

import "strings"

func normalizeAndCount(word string) map[rune]int {
	lowerWord := strings.ToLower(word)
	counts := make(map[rune]int)
	for _, char := range lowerWord {
		counts[char]++
	}
	return counts
}

func Detect(subject string, candidates []string) []string {
	targetCounts := normalizeAndCount(subject)
	normalizedSubject := strings.ToLower(subject)
	targetLength := len(subject)
	
	var anagrams []string
	
	for _, candidate := range candidates {
		
		if len(candidate) != targetLength {
			continue
		}
		
		normalizedCandidate := strings.ToLower(candidate)
		if normalizedCandidate == normalizedSubject {
			continue
		}
		
		candidateCounts := normalizeAndCount(candidate)
		
		if len(targetCounts) != len(candidateCounts) {
			continue
		}
		
		isAnagram := true
		for char, targetCount := range targetCounts {
			if candidateCounts[char] != targetCount {
				isAnagram = false
				break
			}
		}
		
		if isAnagram {
			anagrams = append(anagrams, candidate)
		}
	}
	
	return anagrams
}