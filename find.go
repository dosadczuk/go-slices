package slices

// Find returns the first element matching the given predicate.
func Find[S ~[]E, E any](vals S, predicate func(E) bool) (out E, found bool) {
	for _, val := range vals {
		if predicate(val) {
			return val, true
		}
	}
	return
}
