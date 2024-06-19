package slices

// ReduceRight returns accumulated value starting with the last element
// and applying operation from right to left to current accumulator value.
func ReduceRight[S ~[]E, E any](vals S, operation func(acc E, val E) E) (acc E) {
	size := len(vals)
	if size == 0 {
		return
	}

	acc = vals[size-1]
	for i := size - 2; i >= 0; i-- {
		acc = operation(acc, vals[i])
	}
	return
}
