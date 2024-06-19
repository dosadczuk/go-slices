package slices

// All returns true if all elements match the given predicate.
func All[S ~[]E, E any](vals S, predicate func(E) bool) bool {
	for _, val := range vals {
		if !predicate(val) {
			return false
		}
	}
	return true
}
