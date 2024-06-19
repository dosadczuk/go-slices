package slices_test

import (
	"testing"

	"github.com/dosadczuk/go-slices"
	"github.com/google/go-cmp/cmp"
)

func TestGroupBy(t *testing.T) {
	tt := map[string]struct {
		// input
		values   []int
		selector func(int) int
		// assert
		want map[int][]int
	}{
		"empty slice": {
			values:   nil, // zero value
			selector: func(val int) int { return 0 },
			want:     nil, // zero value
		},
		"selector returns unique keys": {
			values:   []int{1, 2, 3, 4, 5},
			selector: func(val int) int { return val },
			want: map[int][]int{
				1: {1},
				2: {2},
				3: {3},
				4: {4},
				5: {5},
			},
		},
		"selector returns duplicated keys": {
			values:   []int{1, 2, 3, 4, 5},
			selector: func(val int) int { return val % 2 },
			want: map[int][]int{
				1: {1, 3, 5},
				0: {2, 4},
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			have := slices.GroupBy(tc.values, tc.selector)

			if !cmp.Equal(tc.want, have) {
				t.Error(cmp.Diff(tc.want, have))
			}
		})
	}
}
