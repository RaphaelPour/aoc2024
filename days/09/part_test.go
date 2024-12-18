package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	// require.Equal(t, 0, part1([]string{})
}

func TestIndex(t *testing.T) {
	for _, tc := range []struct {
		name     string
		in       string
		needle   rune
		index    int
		expected int
	}{
		{name: "valid", in: "1234", needle: '1', index: 0, expected: 0},
		{name: "valid", in: "1234", needle: '1', index: 1, expected: -1},
		{name: "valid", in: "1231", needle: '1', index: 1, expected: 3},
	} {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, Index([]rune(tc.in), tc.needle, tc.index))
		})
	}
}
