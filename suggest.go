package main

import (
	"strings"
)

func suggest(cmd string, freq uint8, history []string) (string, bool) {
	freq = max(freq, 1)   // MIN: 1
	freq = min(freq, 100) // MAX: 100

	if len(history) == 0 {
		return "", false
	}

	count := 1

	for _, row := range history {
		row = strings.TrimSpace(strings.Join(strings.Split(strings.TrimSpace(row), " ")[1:], " "))

		if row == cmd {
			count++
		}
	}

	percent := uint8(count * 100 / (len(history) + 1))
	if percent >= freq {
		return cmd, true
	}

	return "", false
}
