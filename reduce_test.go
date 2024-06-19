package slices_test

import (
	"testing"

	"github.com/dosadczuk/go-slices"
	"github.com/google/go-cmp/cmp"
)

func TestReduce(t *testing.T) {
	tt := map[string]struct {
		// input
		values    []string
		operation func(string, string) string
		// assert
		want string
	}{
		"empty slice": {
			values:    nil, // zero value
			operation: func(_ string, _ string) string { return "" },
			want:      "",
		},
		"slice with value": {
			values:    []string{"1"},
			operation: func(acc string, val string) string { return acc + val },
			want:      "1",
		},
		"slice with values": {
			values:    []string{"1", "2", "3", "4", "5"},
			operation: func(acc string, val string) string { return acc + val },
			want:      "12345",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			have := slices.Reduce(tc.values, tc.operation)

			if !cmp.Equal(tc.want, have) {
				t.Error(cmp.Diff(tc.want, have))
			}
		})
	}
}
