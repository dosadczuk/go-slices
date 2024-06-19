package slices

// Last returns the last element.
func Last[S ~[]E, E any](vals S) (out E, found bool) {
	if len(vals) == 0 {
		return // not found
	}
	return vals[len(vals)-1], true
}
