package slices_test

import (
	"testing"

	"github.com/dosadczuk/go-slices"
	"github.com/google/go-cmp/cmp"
)

func TestPartition(t *testing.T) {
	tt := map[string]struct {
		// input
		values    []int
		predicate func(int) bool
		// assert
		wantT []int
		wantF []int
	}{
		"empty slice": {
			values:    nil, // zero value
			predicate: func(_ int) bool { return true },
			wantT:     nil, // zero value
			wantF:     nil, // zero value
		},
		"predicate matching first element": {
			values:    []int{1, 2, 3, 4, 5},
			predicate: func(val int) bool { return val == 1 },
			wantT:     []int{1},
			wantF:     []int{2, 3, 4, 5},
		},
		"predicate matching middle element": {
			values:    []int{1, 2, 3, 4, 5},
			predicate: func(val int) bool { return val == 3 },
			wantT:     []int{3},
			wantF:     []int{1, 2, 4, 5},
		},
		"predicate matching last element": {
			values:    []int{1, 2, 3, 4, 5},
			predicate: func(val int) bool { return val == 5 },
			wantT:     []int{5},
			wantF:     []int{1, 2, 3, 4},
		},
		"predicate matching range of elements": {
			values:    []int{1, 2, 3, 4, 5},
			predicate: func(val int) bool { return val < 3 },
			wantT:     []int{1, 2},
			wantF:     []int{3, 4, 5},
		},
		"predicate not matching an element": {
			values:    []int{1, 2, 3, 4, 5},
			predicate: func(val int) bool { return val < 0 },
			wantT:     nil, // zero value
			wantF:     []int{1, 2, 3, 4, 5},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			haveT, haveF := slices.Partition(tc.values, tc.predicate)

			if !cmp.Equal(tc.wantT, haveT) {
				t.Error(cmp.Diff(tc.wantT, haveT))
			}
			if !cmp.Equal(tc.wantF, haveF) {
				t.Error(cmp.Diff(tc.wantF, haveF))
			}
		})
	}
}
