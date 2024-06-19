package slices

// Min returns the smallest value from the elements.
func Min[S ~[]E, E orderable](vals S) (min E) {
	if len(vals) == 0 {
		return
	}

	min = vals[0]
	for i := 1; i < len(vals); i++ {
		if val := vals[i]; val < min {
			min = val
		}
	}
	return
}

// MinFunc returns the smallest value from the elements.
//
// Function compares two arguments for order and returns:
//
//	 0 - if the arguments are equal,
//	-1 - if the first argument is less than the second,
//	+1 - if the second argument is less than the first.
func MinFunc[S ~[]E, E any](vals S, cmp func(E, E) int) (min E) {
	if len(vals) == 0 {
		return
	}

	min = vals[0]
	for i := 1; i < len(vals); i++ {
		if val := vals[i]; cmp(val, min) < 0 {
			min = val
		}
	}
	return
}
