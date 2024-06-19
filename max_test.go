package slices_test

import (
	"testing"

	"github.com/dosadczuk/go-slices"
	"github.com/google/go-cmp/cmp"
)

func TestMax(t *testing.T) {
	tt := map[string]struct {
		// input
		values []int
		// assert
		want int
	}{
		"empty slice": {
			values: nil, // zero value
			want:   0,   // zero value
		},
		"slice with value": {
			values: []int{1},
			want:   1,
		},
		"slice with unique values": {
			values: []int{1, 4, 2, 5, 3},
			want:   5,
		},
		"slice with duplicate valued": {
			values: []int{1, 1, 5, 5, 5, 5, 3, 4, 2, 6, 6, 3},
			want:   6,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			have := slices.Max(tc.values)

			if !cmp.Equal(tc.want, have) {
				t.Error(cmp.Diff(tc.want, have))
			}
		})
	}
}

func TestMaxFunc(t *testing.T) {
	tt := map[string]struct {
		// input
		values []int
		// assert
		want int
	}{
		"empty slice": {
			values: nil, // zero value
			want:   0,   // zero value
		},
		"slice with value": {
			values: []int{1},
			want:   1,
		},
		"slice with unique values": {
			values: []int{1, 4, 2, 5, 3},
			want:   5,
		},
		"slice with duplicate valued": {
			values: []int{1, 1, 5, 5, 5, 5, 3, 4, 2, 6, 6, 3},
			want:   6,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			have := slices.MaxFunc(tc.values, func(a, b int) int { return a - b })

			if !cmp.Equal(tc.want, have) {
				t.Error(cmp.Diff(tc.want, have))
			}
		})
	}
}
