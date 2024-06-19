package slices

// FindLast returns the last element matching the given predicate.
func FindLast[S ~[]E, E any](vals S, predicate func(E) bool) (out E, found bool) {
	for i := len(vals) - 1; i >= 0; i-- {
		if val := vals[i]; predicate(val) {
			return val, true
		}
	}
	return
}
