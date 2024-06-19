package slices

// Flatten returns 1D array of all elements from 2D array.
func Flatten[S ~[][]E, R []E, E any](vals S) (out R) {
	var size int
	for _, vals := range vals {
		size += len(vals)
	}

	if size == 0 {
		return
	}

	out = make(R, 0, size)
	for _, vals := range vals {
		out = append(out, vals...)
	}
	return
}
