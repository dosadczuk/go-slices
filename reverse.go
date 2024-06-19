package slices

// Reverse reverses order of the elements in-place.
func Reverse[S ~[]E, E any](vals S) {
	for i, j := 0, len(vals)-1; i < j; i, j = i+1, j-1 {
		vals[i], vals[j] = vals[j], vals[i]
	}
}
