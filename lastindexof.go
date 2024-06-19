package slices

// LastIndexOf returns index of the last element matching the target.
// Returns -1 if not in the elements.
func LastIndexOf[S ~[]E, E comparable](vals S, target E) int {
	for i := len(vals) - 1; i >= 0; i-- {
		if vals[i] == target {
			return i
		}
	}
	return -1
}
