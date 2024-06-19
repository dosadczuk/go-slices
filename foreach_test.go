package slices_test

import (
	"testing"

	"github.com/dosadczuk/go-slices"
	"github.com/google/go-cmp/cmp"
)

func TestForEach(t *testing.T) {
	tt := map[string]struct {
		// input
		values []int
	}{
		"empty slice": {
			values: nil, // zero value
		},
		"slice with value": {
			values: []int{1},
		},
		"slice with values": {
			values: []int{1, 2, 3, 4, 5},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			var have []int
			slices.ForEach(tc.values, func(_ int, val int) {
				have = append(have, val)
			})

			if !cmp.Equal(tc.values, have) {
				t.Error(cmp.Diff(tc.values, have))
			}
		})
	}
}
