import sys
import os
from alias import parse_aliases, find_alias, suggest_alias
from debug import Debug

COLOR_WARNING = "\033[0;33m"
COLOR_KEY     = "\033[0;32m"
COLOR_VALUE   = "\033[1;32m"
COLOR_DEBUG   = "\033[0;90m"
COLOR_RESET   = "\033[0m"

DELIMITER = "__oh_my_aliases__DELIMITER"

def main():
    config_expand_alias = os.getenv("ZSH_PLUGINS_OH_MY_ALIASES_EXPAND_ALIAS", "0") == "1"
    config_suggest_alias = os.getenv("ZSH_PLUGINS_OH_MY_ALIASES_SUGGEST_ALIAS", "0") == "1"
    config_enable_debug = os.getenv("ZSH_PLUGINS_OH_MY_ALIASES_DEBUG", "0") == "1"
    config_suggestion_freq_percent = int(os.getenv("ZSH_PLUGINS_OH_MY_ALIASES_SUGGESTION_FREQUENCY_PERCENT", "10"))


    if len(sys.argv) < 3:
        if config_enable_debug:
            print(COLOR_DEBUG + "not enough args" + COLOR_RESET)
        return

    window_width = int(sys.argv[1].split("=").pop())
    inp = sys.argv[2]

    d = Debug(config_enable_debug, window_width, {"ea":config_expand_alias, "sa":config_suggest_alias, "h%": config_suggestion_freq_percent})
    
    alias_rows, history_rows = stdin_rows()
    by_command, by_alias = parse_aliases(alias_rows)

    alias = find_alias(inp, by_command)
    if alias:
        print_key_value("alias found", alias, d)
    else:
        expanded = None

        if config_expand_alias:
            expanded = find_alias(inp, by_alias)
            if expanded:
                print_key_value("run command", expanded, d)

        if config_suggest_alias and (not expanded or not config_expand_alias):
            suggested = suggest_alias(inp, config_suggestion_freq_percent, history_rows)

def print_key_value(key: str, value: str, d: Debug):
    msg = f"{COLOR_KEY}{key}: {COLOR_VALUE}{value}{COLOR_RESET}"
    if d.enabled:
        msg = f"{msg} {COLOR_DEBUG}{d.get_info(len(msg) - 15)}{COLOR_RESET}"

    print(msg)


def stdin_rows():
    groups = ([], [])

    pos = 0

    for row in sys.stdin.readlines():
        if DELIMITER in row:
            pos += 1
            continue

        groups[pos].append(row.strip("\n"))

    return groups 

if __name__ == '__main__':
    main()
