package slices

// None returns true if none of the elements match the given predicate.
// If predicate is `nil`, returns true if slice is empty.
func None[S ~[]E, E any](vals S, predicate func(E) bool) bool {
	if predicate == nil {
		return len(vals) == 0
	}

	for _, val := range vals {
		if predicate(val) {
			return false
		}
	}
	return true
}
