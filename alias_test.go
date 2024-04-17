package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindAlias(t *testing.T) {
	testCases := map[string]struct {
		cmd           string
		aliases       map[string]string
		expectedAlias string
	}{
		"nil aliases": {
			cmd:           "go test",
			aliases:       nil,
			expectedAlias: "",
		},
		"empty aliases": {
			cmd:           "go test",
			aliases:       map[string]string{},
			expectedAlias: "",
		},
		"exact": {
			cmd:           "go test",
			expectedAlias: "got",
			aliases: map[string]string{
				"go run":  "gor",
				"go test": "got",
			},
		},
		"with args": {
			cmd:           "go test -v",
			expectedAlias: "got -v",
			aliases: map[string]string{
				"go run":  "gor",
				"go test": "got",
			},
		},
		"with deeper args": {
			cmd:           "go test -v -count=1 ./...",
			expectedAlias: "got -v -count=1 ./...",
			aliases: map[string]string{
				"go run":  "gor",
				"go test": "got",
			},
		},
		"unknown command": {
			cmd:           "go test -v -count=1 ./...",
			expectedAlias: "",
			aliases: map[string]string{
				"go run": "gor",
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual, ok := find(tc.cmd, tc.aliases)
			if len(tc.expectedAlias) == 0 {
				require.False(t, ok)
				return
			}

			require.True(t, ok)
			require.Equal(t, tc.expectedAlias, actual)
		})
	}
}

func TestParseAlias(t *testing.T) {
	testCases := map[string]struct {
		rows              []string
		expected          map[string]string
		expectedViceVersa map[string]string
	}{
		"empty rows": {
			rows:              nil,
			expected:          nil,
			expectedViceVersa: nil,
		},
		"list": {
			rows: []string{
				"rd=rmdir",
				"l='ls -lah'",
				"'md'=mkdir -p",
				"'gc!'='git commit --verbose --amend'",
			},
			expected: map[string]string{
				"rmdir":                        "rd",
				"ls -lah":                      "l",
				"mkdir -p":                     "md",
				"git commit --verbose --amend": "gc!",
			},
			expectedViceVersa: map[string]string{
				"rd":  "rmdir",
				"l":   "ls -lah",
				"md":  "mkdir -p",
				"gc!": "git commit --verbose --amend",
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual, viceVersa := parseAliases(tc.rows)
			require.Equal(t, tc.expected, actual)
			require.Equal(t, tc.expectedViceVersa, viceVersa)
		})
	}
}
