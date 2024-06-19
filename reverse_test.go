package slices_test

import (
	"testing"

	"github.com/dosadczuk/go-slices"
	"github.com/google/go-cmp/cmp"
)

func TestReverse(t *testing.T) {
	tt := map[string]struct {
		// input
		values []int
		// assert
		want []int
	}{
		"empty slice": {
			values: nil, // zero value
			want:   nil, // zero value
		},
		"slice with value": {
			values: []int{1},
			want:   []int{1},
		},
		"slice with values": {
			values: []int{1, 2, 3, 4, 5},
			want:   []int{5, 4, 3, 2, 1},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			have := tc.values
			slices.Reverse(have)

			if !cmp.Equal(tc.want, have) {
				t.Error(cmp.Diff(tc.want, have))
			}
		})
	}
}
