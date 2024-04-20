package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	freqPercent = 10
)

type debugInfo struct {
	start         time.Time
	amountAliases int
	amountHistory int
}

func (d debugInfo) String() string {
	return fmt.Sprintf(
		"(took: %d μs; aliases: %d; history: %d)",
		time.Since(d.start).Microseconds(),
		d.amountAliases,
		d.amountHistory,
	)
}

var debug debugInfo

func main() {
	os.Exit(run())
}

func run() int {
	debug = debugInfo{start: time.Now()}

	if len(os.Args) < 2 {
		printWarning("not found command in aliasRows")
		return exitError
	}

	expandAlias := envParam("ZSH_PLUGINS_OH_MY_ALIASES_EXPAND_ALIAS", "0") == "1"
	suggestAlias := envParam("ZSH_PLUGINS_OH_MY_ALIASES_SUGGEST_ALIAS", "0") == "1"

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

		if expandAlias {
			expandedCmd, found = find(cmd, viceVersa)
			if found {
				printInfo(fmt.Sprintf("run command: %s", expandedCmd))
			}
		}

		if suggestAlias && (!found || !expandAlias) {
			res, has := suggest(cmd, freqPercent, historyRows)
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
	isDebug := envParam("ZSH_PLUGINS_OH_MY_ALIASES_DEBUG", "0") == "1"
	if isDebug {
		msg = fmt.Sprintf("%s %s%s", msg, colorBlue, debug.String())
	}

	_, _ = fmt.Fprintln(os.Stdout, fmt.Sprintf("%s%s%s", color, msg, colorReset))
}

func envParam(key string, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaultVal
	}

	return val
}

func envParamInt(key string, defaultVal int64) int64 {
	val := envParam(key, "")
	if val == "" {
		return defaultVal
	}

	intVal, _ := strconv.ParseInt(val, 10, 64)
	if intVal <= 0 {
		return defaultVal
	}

	return intVal
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
