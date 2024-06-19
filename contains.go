package slices

// Contains returns true if target is found in the elements.
func Contains[S ~[]E, E comparable](vals S, target E) bool {
	for _, val := range vals {
		if val == target {
			return true
		}
	}
	return false
}
