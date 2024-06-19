package slices_test

import (
	"testing"

	"github.com/dosadczuk/go-slices"
	"github.com/google/go-cmp/cmp"
)

func TestFlatten(t *testing.T) {
	tt := map[string]struct {
		// input
		values [][]int
		// assert
		want []int
	}{
		"empty slice": {
			values: nil, // zero value
			want:   nil, // zero value
		},
		"slice with sub-slice with value": {
			values: [][]int{{1}},
			want:   []int{1},
		},
		"slice with sub-slice with values": {
			values: [][]int{{1, 2, 3, 4, 5}},
			want:   []int{1, 2, 3, 4, 5},
		},
		"slice with sub-slices with value": {
			values: [][]int{{1}, {2}, {3}, {4}, {5}},
			want:   []int{1, 2, 3, 4, 5},
		},
		"slice with sub-slices with values": {
			values: [][]int{{1, 2}, {3, 4}, {5}},
			want:   []int{1, 2, 3, 4, 5},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			have := slices.Flatten(tc.values)

			if !cmp.Equal(tc.want, have) {
				t.Error(cmp.Diff(tc.want, have))
			}
		})
	}
}
