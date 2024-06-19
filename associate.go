package slices

// Associate returns a map of key-value pairs created by applying transform
// function to each of the given elements.
func Associate[S ~[]E, M map[K]V, E any, K comparable, V any](vals S, transform func(E) (K, V)) M {
	if len(vals) == 0 {
		return nil
	}

	out := make(M, len(vals))
	for _, val := range vals {
		k, v := transform(val)
		out[k] = v
	}
	return out
}
