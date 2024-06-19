package slices_test

import (
	"testing"

	"github.com/dosadczuk/go-slices"
	"github.com/google/go-cmp/cmp"
)

func TestNone(t *testing.T) {
	tt := map[string]struct {
		// input
		values    []int
		predicate func(int) bool
		// assert
		want bool
	}{
		"empty slice": {
			values:    nil, // zero value
			predicate: func(_ int) bool { return true },
			want:      true,
		},
		"empty slice and predicate is nil": {
			values:    nil, // zero value
			predicate: nil,
			want:      true,
		},
		"slice with values and predicate is nil": {
			values:    []int{1, 2, 3, 4, 5},
			predicate: nil,
			want:      false,
		},
		"slice with values matching predicate": {
			values:    []int{1, 2, 3, 4, 5},
			predicate: func(val int) bool { return val%2 == 0 },
			want:      false,
		},
		"slice with values not matching predicate": {
			values:    []int{1, 2, 3, 4, 5},
			predicate: func(val int) bool { return val < 0 },
			want:      true,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			have := slices.None(tc.values, tc.predicate)

			if !cmp.Equal(tc.want, have) {
				t.Error(cmp.Diff(tc.want, have))
			}
		})
	}
}
