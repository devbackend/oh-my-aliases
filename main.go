package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	exitSuccess = 0
	exitError   = 1

	colorWarning = "\033[0;33m"
	colorKey     = "\033[0;32m"
	colorValue   = "\033[1;32m"
	colorDebug   = "\033[0;90m"

	colorReset = "\033[0m"

	inputDelimiter = "__oh_my_aliases__DELIMITER"
)

type keyValMsg struct {
	key string
	val string
}

func (kv keyValMsg) String() string {
	return fmt.Sprintf("%s%s: %s%s%s", colorKey, kv.key, colorValue, kv.val, colorReset)
}

func main() {
	initConfig()

	os.Exit(run())
}

func run() int {
	debug = debugInfo{start: time.Now()}

	if len(os.Args) < 2 {
		printWarning("not found command in aliasRows")
		return exitError
	}

	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))

	aliasRows := readRows(scanner)
	historyRows := readRows(scanner)

	debug.amountAliases = len(aliasRows)
	debug.amountHistory = len(historyRows)

	cmd := os.Args[1]

	aliases, viceVersa := parseAliases(aliasRows)

	alias, ok := find(cmd, aliases)
	if ok {
		printKeyValue("alias found", alias)
	} else {
		var expandedCmd string
		var found bool

		if config.expandAlias {
			expandedCmd, found = find(cmd, viceVersa)
			if found {
				printKeyValue("run command", expandedCmd)
			}
		}

		if config.suggestNewAliases && (!found || !config.expandAlias) {
			res, has := suggest(cmd, uint8(config.suggestionFrequencyPercent), historyRows)
			if has {
				printKeyValue("suggest new alias", res)
			}
		}
	}

	return exitSuccess
}

func printWarning(msg string) {
	printMessage(msg, colorWarning)
}

func printKeyValue(k, v string) {
	msg := keyValMsg{key: k, val: v}.String()
	if config.isDebug {
		msg = fmt.Sprintf("%s %s%s%s", msg, colorDebug, debug.String(), colorReset)
	}

	_, _ = fmt.Fprintln(os.Stdout, msg)
}

func printMessage(msg string, color string) {
	if config.isDebug {
		msg = fmt.Sprintf("%s %s%s", msg, colorDebug, debug.String())
	}

	_, _ = fmt.Fprintln(os.Stdout, fmt.Sprintf("%s%s%s", color, msg, colorReset))
}

func readRows(s *bufio.Scanner) []string {
	res := make([]string, 0)

	for s.Scan() {
		row := s.Text()
		if row == inputDelimiter {
			break
		}

		res = append(res, row)
	}

	return res
}
