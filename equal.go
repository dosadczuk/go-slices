package slices

// Equal returns true if two slices have the same values in the same order.
func Equal[S ~[]E, E comparable](s1, s2 S) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

// EqualFunc returns true if two slices have the same values in the same order.
func EqualFunc[S1 ~[]E1, S2 ~[]E2, E1, E2 any](s1 S1, s2 S2, eq func(E1, E2) bool) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if !eq(s1[i], s2[i]) {
			return false
		}
	}
	return true
}
