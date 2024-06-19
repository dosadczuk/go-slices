package slices

// ForEach calls the given action on each element.
func ForEach[S ~[]E, E any](vals S, action func(int, E)) {
	for i, val := range vals {
		action(i, val)
	}
}
