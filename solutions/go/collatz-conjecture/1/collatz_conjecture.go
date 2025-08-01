package collatzconjecture
import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		// Create and return a new error for invalid input.
		return 0, errors.New("input must be a positive number")
	}

	steps := 0
	for n != 1 {
		if n % 2 == 0 {
			n = n / 2
		} else {
			n = (3 * n) + 1
		}
		steps++
	}

	// On success, return the step count and 'nil' for the error.
	return steps, nil
}
