def parse_aliases(rows):
    by_cmd, by_alias = {}, {}

    for row in rows:
        parts = row.split("=")
        if len(parts) != 2:
            continue

        alias, command = extract(parts[0]), extract(parts[1])

        by_cmd[command] = alias
        by_alias[alias] = command

    return by_cmd, by_alias

def find_alias(cmd, items):
    if len(cmd) == 0 or len(items) == 0:
        return None

    if cmd in items:
        return items[cmd]

    sorted_items = dict(sorted(items.items(), reverse=True))

    for full, alias in sorted_items.items():
        if cmd.find(full) == 0:
            return cmd.replace(full, alias)

    return None

def suggest_alias(cmd, freq, history):
    freq = max(freq, 1)
    freq = min(freq, 100)

    if len(cmd) == 0 or len(history) == 0:
        return None

    count = 1
    for row in history:
        row = " ".join(" ".split(row.strip())[1:]).strip()
        if row == cmd:
            count += 1

    call_percent = count * 100 / (len(history) + 1)
    if call_percent >= freq:
        return cmd

    return None

def extract(str):
    if len(str) == 0 or str[0] != "'":
        return str

    return str[1:len(str) - 1]
