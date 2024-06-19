package slices_test

import (
	"testing"

	"github.com/dosadczuk/go-slices"
	"github.com/google/go-cmp/cmp"
)

func TestFirst(t *testing.T) {
	tt := map[string]struct {
		// input
		values []int
		// assert
		wantValue int
		wantFound bool
	}{
		"empty slice": {
			values:    nil,   // zero value
			wantValue: 0,     // zero value
			wantFound: false, // zero value
		},
		"slice with value": {
			values:    []int{1},
			wantValue: 1,
			wantFound: true,
		},
		"slice with values": {
			values:    []int{1, 2, 3, 4, 5},
			wantValue: 1,
			wantFound: true,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			haveValue, haveFound := slices.First(tc.values)

			if !cmp.Equal(tc.wantValue, haveValue) {
				t.Error(cmp.Diff(tc.wantValue, haveValue))
			}
			if !cmp.Equal(tc.wantFound, haveFound) {
				t.Error(cmp.Diff(tc.wantFound, haveFound))
			}
		})
	}
}
