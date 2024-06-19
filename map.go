package slices

// Map returns a slice of values created by applying transform function
// to each of the given elements.
func Map[S ~[]E1, R []E2, E1, E2 any](vals S, transform func(E1) E2) R {
	if len(vals) == 0 {
		return nil
	}

	out := make(R, 0, len(vals))
	for _, val := range vals {
		out = append(out, transform(val))
	}
	return out
}
