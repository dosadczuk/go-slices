package slices

// FoldRight returns accumulated value starting with the given initial value
// and applying operation from right to left to current accumulator value.
// Returns initial value, if slice is empty.
func FoldRight[S ~[]E, E any](vals S, initial E, operation func(acc E, val E) E) (acc E) {
	if len(vals) == 0 {
		return initial
	}

	acc = initial
	for i := len(vals) - 1; i >= 0; i-- {
		acc = operation(acc, vals[i])
	}
	return
}
