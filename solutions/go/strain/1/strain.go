package strain

func Keep[T any](items []T, predicate func(T) bool) []T {
	var kept []T
	for _, item := range items {
		if predicate(item) {
			kept = append(kept, item)
		}
	}
	return kept
}
func Discard[T any](items []T, predicate func(T) bool) []T {
	var discarded []T
	for _, item := range items {
		if !predicate(item) {
			discarded = append(discarded, item)
		}
	}
	return discarded
}