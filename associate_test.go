package slices_test

import (
	"testing"

	"github.com/dosadczuk/go-slices"
	"github.com/google/go-cmp/cmp"
)

func TestAssociate(t *testing.T) {
	tt := map[string]struct {
		// input
		values    []int
		transform func(int) (int, int)
		// assert
		want map[int]int
	}{
		"empty slice": {
			values:    nil, // zero value
			transform: func(val int) (int, int) { return val, 0 },
			want:      nil, // zero value
		},
		"transform returns unique keys": {
			values: []int{1, 2, 3, 4, 5},
			transform: func(val int) (int, int) {
				return val, val % 2
			},
			want: map[int]int{
				1: 1,
				2: 0,
				3: 1,
				4: 0,
				5: 1,
			},
		},
		"transform returns duplicated keys": {
			values: []int{1, 2, 3, 4, 5},
			transform: func(val int) (int, int) {
				return val % 2, val
			},
			want: map[int]int{
				0: 4,
				1: 5,
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			have := slices.Associate(tc.values, tc.transform)

			if !cmp.Equal(tc.want, have) {
				t.Error(cmp.Diff(tc.want, have))
			}
		})
	}
}
