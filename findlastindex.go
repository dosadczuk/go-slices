package slices

// FindLastIndex returns index of the last element matching the given predicate.
func FindLastIndex[S ~[]E, E any](vals S, predicate func(E) bool) (idx int, found bool) {
	for i := len(vals) - 1; i >= 0; i-- {
		if val := vals[i]; predicate(val) {
			return i, true
		}
	}
	return
}
