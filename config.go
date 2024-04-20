package main

import (
	"os"
	"strconv"
)

var config = struct {
	isDebug                    bool
	expandAlias                bool
	suggestNewAliases          bool
	suggestionFrequencyPercent int
}{}

func initConfig() {
	config.expandAlias = envParam("ZSH_PLUGINS_OH_MY_ALIASES_EXPAND_ALIAS", "0") == "1"
	config.suggestNewAliases = envParam("ZSH_PLUGINS_OH_MY_ALIASES_SUGGEST_ALIAS", "0") == "1"
	config.isDebug = envParam("ZSH_PLUGINS_OH_MY_ALIASES_DEBUG", "0") == "1"
	config.suggestionFrequencyPercent = envParamInt("ZSH_PLUGINS_OH_MY_ALIASES_SUGGESTION_FREQUENCY_PERCENT", 10)
}

func envParam(key string, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaultVal
	}

	return val
}

func envParamInt(key string, defaultVal int) int {
	val := envParam(key, "")
	if val == "" {
		return defaultVal
	}

	intVal, _ := strconv.Atoi(val)
	if intVal <= 0 {
		return defaultVal
	}

	return intVal
}
