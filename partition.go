package slices

// Partition splits elements in two partitions, where first contains elements
// for which predicate returns true, while second contains elements for which
// predicate returns false.
func Partition[S ~[]E, T []E, F []E, E any](vals S, predicate func(E) bool) (tout T, fout F) {
	if len(vals) == 0 {
		return
	}

	for _, val := range vals {
		if predicate(val) {
			tout = append(tout, val)
		} else {
			fout = append(fout, val)
		}
	}
	return
}
