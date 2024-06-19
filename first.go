package slices

// First returns the first element.
func First[S ~[]E, E any](vals S) (out E, found bool) {
	if len(vals) == 0 {
		return // not found
	}
	return vals[0], true
}
