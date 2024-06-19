package slices_test

import (
	"testing"

	"github.com/dosadczuk/go-slices"
	"github.com/google/go-cmp/cmp"
)

func TestContains(t *testing.T) {
	tt := map[string]struct {
		// input
		values []int
		target int
		// assert
		want bool
	}{
		"empty slice": {
			values: nil, // zero value
			target: 0,
			want:   false,
		},
		"slice with matching value": {
			values: []int{1, 2, 3, 4, 5},
			target: 4,
			want:   true,
		},
		"slice with not matching value": {
			values: []int{1, 2, 3, 4, 5},
			target: 0,
			want:   false,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			have := slices.Contains(tc.values, tc.target)

			if !cmp.Equal(tc.want, have) {
				t.Error(cmp.Diff(tc.want, have))
			}
		})
	}
}
