package slices_test

import (
	"testing"

	"github.com/dosadczuk/go-slices"
	"github.com/google/go-cmp/cmp"
)

func TestFindLast(t *testing.T) {
	tt := map[string]struct {
		// input
		values    []int
		predicate func(int) bool
		// assert
		wantValue int
		wantFound bool
	}{
		"empty slice": {
			values:    nil, // zero value
			predicate: func(_ int) bool { return true },
			wantValue: 0,     // zero value
			wantFound: false, // zero value
		},
		"slice with value matching predicate": {
			values:    []int{1, 2, 3, 4, 5},
			predicate: func(val int) bool { return val%2 == 0 },
			wantValue: 4,
			wantFound: true,
		},
		"slice with value not matching predicate": {
			values:    []int{1, 2, 3, 4, 5},
			predicate: func(val int) bool { return val < 0 },
			wantValue: 0,     // zero value
			wantFound: false, // zero value
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			haveValue, haveFound := slices.FindLast(tc.values, tc.predicate)

			if !cmp.Equal(tc.wantValue, haveValue) {
				t.Error(cmp.Diff(tc.wantValue, haveValue))
			}
			if !cmp.Equal(tc.wantFound, haveFound) {
				t.Error(cmp.Diff(tc.wantFound, haveFound))
			}
		})
	}
}
