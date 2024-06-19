package slices_test

import (
	"testing"

	"github.com/dosadczuk/go-slices"
	"github.com/google/go-cmp/cmp"
)

func TestFoldRight(t *testing.T) {
	tt := map[string]struct {
		// input
		values    []string
		initial   string
		operation func(string, string) string
		// assert
		want string
	}{
		"empty slice": {
			values:    nil, // zero value
			initial:   "",  // zero value
			operation: func(_ string, _ string) string { return "" },
			want:      "",
		},
		"empty slice and initial value": {
			values:    nil, // zero value
			initial:   "order:",
			operation: func(_ string, _ string) string { return "" },
			want:      "order:",
		},
		"slice with value": {
			values:    []string{"1"},
			initial:   "",
			operation: func(acc string, val string) string { return acc + val },
			want:      "1",
		},
		"slice with value and initial value": {
			values:    []string{"1"},
			initial:   "order:",
			operation: func(acc string, val string) string { return acc + val },
			want:      "order:1",
		},
		"slice with values": {
			values:    []string{"1", "2", "3", "4", "5"},
			initial:   "",
			operation: func(acc string, val string) string { return acc + val },
			want:      "54321",
		},
		"slice with values and initial value": {
			values:    []string{"1", "2", "3", "4", "5"},
			initial:   "order:",
			operation: func(acc string, val string) string { return acc + val },
			want:      "order:54321",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			have := slices.FoldRight(tc.values, tc.initial, tc.operation)

			if !cmp.Equal(tc.want, have) {
				t.Error(cmp.Diff(tc.want, have))
			}
		})
	}
}
