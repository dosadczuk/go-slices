package slices

// Reversed returns the elements in reversed order.
func Reversed[S ~[]E, R []E, E any](vals S) R {
	if len(vals) == 0 {
		return nil
	}

	out := make(R, len(vals))
	for i, j := 0, len(vals)-1; j >= 0; i, j = i+1, j-1 {
		out[i] = vals[j]
	}
	return out
}
