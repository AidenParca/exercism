package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

// Valid determines if a number is valid per the Luhn algorithm.
func Valid(id string) bool {
	cleanID := strings.ReplaceAll(id, " ", "")
	if len(cleanID) <= 1 {
		return false
	}
	sum := 0
	shouldDouble := len(cleanID)%2 == 0
	for _, char := range cleanID {
		if !unicode.IsDigit(char) {
			return false
		}
		digit, _ := strconv.Atoi(string(char))
		if shouldDouble {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		shouldDouble = !shouldDouble
	}
	return sum%10 == 0
}