package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSuggest(t *testing.T) {
	testCases := map[string]struct {
		frequency uint8
		cmd       string
		expected  string
		history   []string
	}{
		"empty history": {
			cmd:       "go test -v",
			frequency: 10,
			history:   make([]string, 0),
			expected:  "",
		},
		"new command": {
			cmd:       "go vet",
			frequency: 33,
			expected:  "",
			history: []string{
				" 1 go test -v",
				" 2 go run ./...",
				" 3 go build ./...",
			},
		},
		"exists suggestion": {
			cmd:       "go run ./...",
			frequency: 10,
			expected:  "go run ./...",
			history: []string{
				" 1 go test -v",
				" 2 go run ./...",
				" 3 go run ./...",
				" 4 go run ./...",
				" 5 go run ./...",
				" 6 go build ./...",
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual, ok := suggest(tc.cmd, tc.frequency, tc.history)

			if len(tc.expected) == 0 {
				require.False(t, ok)
				return
			}

			require.True(t, ok)
			require.Equal(t, tc.expected, actual)
		})
	}
}
