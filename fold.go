package slices

// Fold returns accumulated value starting with the given initial value
// and applying operation from left to right to current accumulator value.
// Returns initial value, if slice is empty.
func Fold[S ~[]E, E any](vals S, initial E, operation func(acc E, val E) E) (acc E) {
	if len(vals) == 0 {
		return initial
	}

	acc = initial
	for _, val := range vals {
		acc = operation(acc, val)
	}
	return
}
