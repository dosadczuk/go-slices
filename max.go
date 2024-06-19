package slices

// Max returns the largest value from the elements.
func Max[S ~[]E, E orderable](vals S) (max E) {
	if len(vals) == 0 {
		return
	}

	max = vals[0]
	for i := 1; i < len(vals); i++ {
		if val := vals[i]; val > max {
			max = val
		}
	}
	return
}

// MaxFunc returns the largest value from the elements.
//
// Function compares two arguments for order and returns:
//
//	 0 - if the arguments are equal,
//	-1 - if the first argument is less than the second,
//	+1 - if the second argument is less than the first.
func MaxFunc[S ~[]E, E any](vals S, cmp func(E, E) int) (max E) {
	if len(vals) == 0 {
		return
	}

	max = vals[0]
	for i := 1; i < len(vals); i++ {
		if val := vals[i]; cmp(val, max) > 0 {
			max = val
		}
	}
	return
}
