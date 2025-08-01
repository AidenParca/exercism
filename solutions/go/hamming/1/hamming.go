package hamming

import "errors"

func Distance(a, b string) (int, error) {
	// First, check if the lengths are compatible.
	// Converting to runes ensures you're counting characters, not bytes.
	runesA := []rune(a)
	runesB := []rune(b)

	if len(runesA) != len(runesB) {
		return 0, errors.New("inputs must have the same length")
	}

	// Loop through the strings and count the differences.
	distance := 0
	for i := range runesA {
		if runesA[i] != runesB[i] {
			distance++
		}
	}

	return distance, nil
}