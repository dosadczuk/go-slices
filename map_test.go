package slices_test

import (
	"fmt"
	"testing"

	"github.com/dosadczuk/go-slices"
	"github.com/google/go-cmp/cmp"
)

func TestMap(t *testing.T) {
	tt := map[string]struct {
		// input
		values    []int
		transform func(int) string
		// assert
		want []string
	}{
		"empty slice": {
			values:    nil, // zero value
			transform: func(_ int) string { return "" },
			want:      nil, // zero value
		},
		"slice with value": {
			values: []int{1},
			transform: func(val int) string {
				return fmt.Sprintf("%d->%d", val, val*2)
			},
			want: []string{"1->2"},
		},
		"slice with values": {
			values: []int{1, 2, 3, 4, 5},
			transform: func(val int) string {
				return fmt.Sprintf("%d->%d", val, val*2)
			},
			want: []string{
				"1->2",
				"2->4",
				"3->6",
				"4->8",
				"5->10",
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			have := slices.Map(tc.values, tc.transform)

			if !cmp.Equal(tc.want, have) {
				t.Error(cmp.Diff(tc.want, have))
			}
		})
	}
}
