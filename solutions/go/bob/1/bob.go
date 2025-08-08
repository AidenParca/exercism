package bob

import (
	"strings"
)

// Hey responds to a remark based on a set of rules.
func Hey(remark string) string {
	// 1. Clean the input string by removing all leading/trailing whitespace.
	remark = strings.TrimSpace(remark)

	// 2. Check for silence. This must be the first check.
	if remark == "" {
		return "Fine. Be that way!"
	}

	// 3. Determine the properties of the remark.
	isQuestion := strings.HasSuffix(remark, "?")
	hasLetters := strings.ContainsAny(remark, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	isYelling := hasLetters && (strings.ToUpper(remark) == remark)

	// 4. Apply the rules in the correct order of priority.
	if isYelling && isQuestion {
		return "Calm down, I know what I'm doing!"
	}

	if isYelling {
		return "Whoa, chill out!"
	}

	if isQuestion {
		return "Sure."
	}

	// 5. If no other rule matches, return the default response.
	return "Whatever."
}