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
)

var start time.Time

func main() {
	os.Exit(run())
}

func run() int {
	start = time.Now()

	if len(os.Args) < 2 {
		printWarning("not found command in rows")
		return exitError
	}

	expandAlias := envParam("ZSH_PLUGINS_OH_MY_ALIASES_EXPAND_ALIAS", "0") == "1"

	rows := make([]string, 0)

	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	cmd := os.Args[1]

	aliases, viceVersa := parseAliases(rows)

	alias, ok := find(cmd, aliases)
	if ok {
		printInfo(fmt.Sprintf("alias found: %s", alias))
	} else {
		if expandAlias {
			cmd, ok := find(cmd, viceVersa)
			if ok {
				printInfo(fmt.Sprintf("run command: %s", cmd))
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
		msg = fmt.Sprintf("%s %s(took: %d μs)", msg, colorBlue, time.Since(start).Microseconds())
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
