package isbn

import (
	"strconv"
	"strings"
)

// IsValidISBN checks if a given string is a valid ISBN-10 number.
// (d₁*10 + d₂*9 + d₃*8 + d₄*7 + d₅*6 + d₆*5 + d₇*4 + d₈*3 + d₉*2 + d₁₀*1) % 11 == 0
func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
		return false
	}
	total := 0
	j := 0
	for i := 10; i >= 1; i-- {
		var dig int
		char := isbn[j]
		if i == 1 && (char == 'X' || char == 'x') {
			dig = 10
		} else {
			val, err := strconv.Atoi(string(char))
			if err != nil {
				return false
			}
			dig = val
		}
		total += i * dig
		j++ 
	}
	return total % 11 == 0
}
