package slices

// GroupBy groups elements by the key returned by the given selector function
// applied to each element.
func GroupBy[S ~[]E, M map[K][]E, K comparable, E any](vals S, selector func(E) K) M {
	if len(vals) == 0 {
		return nil
	}

	out := make(M)
	for _, val := range vals {
		k := selector(val)
		out[k] = append(out[k], val)
	}
	return out
}
