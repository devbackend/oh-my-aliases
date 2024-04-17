package main

import (
	"fmt"
	"strings"
)

func parseAliases(rows []string) map[string]string {
	if len(rows) == 0 {
		return nil
	}

	res := make(map[string]string, len(rows))

	for _, row := range rows {
		parts := strings.Split(row, "=")

		if len(parts) < 2 {
			continue
		}

		res[extract(parts[1])] = extract(parts[0])
	}

	return res
}

func findAlias(cmd string, aliases map[string]string) (alias string, found bool) {
	cmd = strings.TrimSpace(cmd)

	if len(cmd) == 0 || len(aliases) == 0 {
		return
	}

	alias, found = aliases[cmd]
	if !found {
		var args string

		args, rest := pop(strings.Split(cmd, " "))

		alias, found = findAlias(strings.Join(rest, " "), aliases)
		if !found {
			return
		}

		alias = fmt.Sprintf("%s %s", alias, args)

		return
	}

	return
}

func extract(str string) string {
	if len(str) == 0 || str[0] != '\'' {
		return str
	}

	return str[1 : len(str)-1]
}

func pop(arr []string) (string, []string) {
	if len(arr) == 0 {
		return "", nil
	}

	lastIx := len(arr)

	return arr[lastIx-1], arr[:lastIx-1]
}
