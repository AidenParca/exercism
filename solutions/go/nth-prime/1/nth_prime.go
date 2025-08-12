package prime

import (
	"errors"
	"math"
)

// Isprime checks if a number is prime.
func Isprime(n int) bool {
	if n <= 1 {
		return false
	}
	limit := int(math.Sqrt(float64(n)))
	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
// Nth returns the nth prime number. An error is returned if n is zero or negative.
func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("prime: input is not a positive number")
	}
	count := 0
	num := 1
	for count < n {
		num++
		if Isprime(num) {
			count++
		}
	}
	return num, nil
}