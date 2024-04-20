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

	colorYellow = "\033[0;33m"
	colorBlue   = "\033[0;34m"
	colorPurple = "\033[0;95m"

	colorReset = "\033[0m"

	inputDelimiter = "__oh_my_aliases__DELIMITER"
)

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
		printInfo(fmt.Sprintf("alias found: %s", alias))
	} else {
		var expandedCmd string
		var found bool

		if config.expandAlias {
			expandedCmd, found = find(cmd, viceVersa)
			if found {
				printInfo(fmt.Sprintf("run command: %s", expandedCmd))
			}
		}

		if config.suggestNewAliases && (!found || !config.expandAlias) {
			res, has := suggest(cmd, uint8(config.suggestionFrequencyPercent), historyRows)
			if has {
				printInfo(fmt.Sprintf("suggest new alias: '%s'", res))
			}
		}
	}

	return exitSuccess
}

func printInfo(msg string) {
	printMessage(msg, colorPurple)
}

func printWarning(msg string) {
	printMessage(msg, colorYellow)
}

func printMessage(msg string, color string) {
	if config.isDebug {
		msg = fmt.Sprintf("%s %s%s", msg, colorBlue, debug.String())
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
