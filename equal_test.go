package slices_test

import (
	"testing"

	"github.com/dosadczuk/go-slices"
	"github.com/google/go-cmp/cmp"
)

func TestEqual(t *testing.T) {
	tt := map[string]struct {
		// input
		values1 []int
		values2 []int
		// assert
		want bool
	}{
		"empty slices": {
			values1: nil, // zero value
			values2: nil, // zero value
			want:    true,
		},
		"slices with different size": {
			values1: []int{1, 2},
			values2: []int{3, 4, 5},
			want:    false,
		},
		"slices of values in the same order": {
			values1: []int{1, 2, 3, 4, 5},
			values2: []int{1, 2, 3, 4, 5},
			want:    true,
		},
		"slices of values in different order": {
			values1: []int{1, 2, 3, 4, 5},
			values2: []int{1, 5, 4, 2, 3},
			want:    false,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			have := slices.Equal(tc.values1, tc.values2)

			if !cmp.Equal(tc.want, have) {
				t.Error(cmp.Diff(tc.want, have))
			}
		})
	}
}

func TestEqualFunc(t *testing.T) {
	tt := map[string]struct {
		// input
		values1 []int
		values2 []int
		// assert
		want bool
	}{
		"empty slices": {
			values1: nil, // zero value
			values2: nil, // zero value
			want:    true,
		},
		"slices with different size": {
			values1: []int{1, 2},
			values2: []int{3, 4, 5},
			want:    false,
		},
		"slices of values in the same order": {
			values1: []int{1, 2, 3, 4, 5},
			values2: []int{1, 2, 3, 4, 5},
			want:    true,
		},
		"slices of values in different order": {
			values1: []int{1, 2, 3, 4, 5},
			values2: []int{1, 5, 4, 2, 3},
			want:    false,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			have := slices.EqualFunc(tc.values1, tc.values2, func(v1, v2 int) bool { return v1 == v2 })

			if !cmp.Equal(tc.want, have) {
				t.Error(cmp.Diff(tc.want, have))
			}
		})
	}
}
