package summultiples


// take the limit, finds every multiples of divisors up to the limit and sum them
func SumMultiples(limit int, divisors ...int) int {
    multiples := make(map[int]bool)
	for _, divisor := range divisors {
		if divisor == 0 {
			continue
		}
		for i := divisor; i < limit; i += divisor {
			multiples[i] = true
		}
	}
	sum := 0
	for num := range multiples {
		sum += num
	}
	return sum
}
