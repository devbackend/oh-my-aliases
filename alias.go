package main

import (
	"fmt"
	"strings"
)

func parseAliases(rows []string) (map[string]string, map[string]string) {
	if len(rows) == 0 {
		return nil, nil
	}

	res := make(map[string]string, len(rows))
	viceVersa := make(map[string]string, len(rows))

	for _, row := range rows {
		parts := strings.Split(row, "=")

		if len(parts) < 2 {
			continue
		}

		alias, command := extract(parts[0]), extract(parts[1])

		res[command] = extract(alias)
		viceVersa[alias] = extract(command)
	}

	return res, viceVersa
}

func find(cmd string, list map[string]string) (value string, found bool) {
	cmd = strings.TrimSpace(cmd)

	if len(cmd) == 0 || len(list) == 0 {
		return
	}

	parts := strings.Split(cmd, "&&")
	if len(parts) > 1 {
		res := make([]string, len(parts))
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])

			a, f := find(parts[i], list)
			if f {
				found = true
				res[i] = a
			} else {
				res[i] = parts[i]
			}
		}
		return strings.Join(res, " && "), found
	}

	value, found = list[cmd]
	if !found {
		var args string

		args, rest := pop(strings.Split(cmd, " "))

		value, found = find(strings.Join(rest, " "), list)
		if !found {
			return
		}

		value = fmt.Sprintf("%s %s", value, args)

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
