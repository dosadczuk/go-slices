package slices

// FindIndex returns index of the first element matching the given predicate.
func FindIndex[S ~[]E, E any](vals S, predicate func(E) bool) (idx int, found bool) {
	for i, val := range vals {
		if predicate(val) {
			return i, true
		}
	}
	return
}
