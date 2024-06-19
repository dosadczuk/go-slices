package slices_test

import (
	"testing"

	"github.com/dosadczuk/go-slices"
	"github.com/google/go-cmp/cmp"
)

func TestFilter(t *testing.T) {
	tt := map[string]struct {
		// input
		values    []int
		predicate func(int) bool
		// assert
		want []int
	}{
		"empty slice": {
			values:    nil, // zero value
			predicate: func(_ int) bool { return true },
			want:      nil, // zero value
		},
		"predicate matching all values": {
			values:    []int{1, 2, 3, 4, 5},
			predicate: func(val int) bool { return val > 0 },
			want:      []int{1, 2, 3, 4, 5},
		},
		"predicate matching few values": {
			values:    []int{1, 2, 3, 4, 5},
			predicate: func(val int) bool { return val%2 == 1 },
			want:      []int{1, 3, 5},
		},
		"predicate matching zero values": {
			values:    []int{1, 2, 3, 4, 5},
			predicate: func(val int) bool { return val < 0 },
			want:      nil, // zero value
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			have := slices.Filter(tc.values, tc.predicate)

			if !cmp.Equal(tc.want, have) {
				t.Error(cmp.Diff(tc.want, have))
			}
		})
	}
}
