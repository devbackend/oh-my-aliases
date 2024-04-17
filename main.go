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

func main() {
	os.Exit(run())
}

func run() int {
	if len(os.Args) < 2 {
		printWarning("not found command in rows")
		return exitError
	}

	isDebug := envParam("ZSH_PLUGINS_OH_MY_ALIASES_DEBUG", "0") == "1"

	if isDebug {
		defer func(start time.Time) {
			printDebug(fmt.Sprintf("[debug] time: %d mcs", time.Since(start).Microseconds()))
		}(time.Now())
	}

	rows := make([]string, 0)

	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	cmd := os.Args[1]

	alias, ok := findAlias(cmd, parseAliases(rows))
	if ok {
		printInfo(fmt.Sprintf("alias found: %s", alias))
	}

	return exitSuccess
}

func printInfo(msg string) {
	printMessage(msg, colorPurple)
}

func printDebug(msg string) {
	printMessage(msg, colorBlue)
}

func printWarning(msg string) {
	printMessage(msg, colorYellow)
}

func printMessage(msg string, color string) {
	_, _ = fmt.Fprintln(os.Stdout, fmt.Sprintf("%s%s%s", color, msg, colorReset))
}

func envParam(key string, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaultVal
	}

	return val
}
