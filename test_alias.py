from unittest import TestCase, main
from alias import parse_aliases, find_alias, suggest_alias

class TestAliases(TestCase):
    def test_parse_aliases(self):
        rows = [
            "rd=rmdir",
            "l='ls -lah'",
            "'md'=mkdir -p",
            "'gc!'='git commit --verbose --amend'",
            "gotest='gotestsum --hide-summary=skipped --format=testname'"
        ]

        expected_by_cmd = {
            "rmdir":                                               "rd",
            "ls -lah":                                             "l",
            "mkdir -p":                                            "md",
            "git commit --verbose --amend":                        "gc!",
            "gotestsum --hide-summary=skipped --format=testname": "gotest",
        }

        expected_by_alias = {
            "rd":     "rmdir",
            "l":      "ls -lah",
            "md":     "mkdir -p",
            "gc!":    "git commit --verbose --amend",
            "gotest": "gotestsum --hide-summary=skipped --format=testname",
        }

        actual_by_cmd, actual_by_alias = parse_aliases(rows)

        self.assertEqual(actual_by_cmd, expected_by_cmd)
        self.assertEqual(actual_by_alias, expected_by_alias)

    def test_find_alias(self):
        aliases = {
            "go run":        "gor",
            "go test":       "got",
            "go test -race": "gotr",
            "go mod":        "gom",
            "git":           "g",
        }

        self.assertEqual(find_alias("", aliases), None)
        self.assertEqual(find_alias("go test", {}), None)
        self.assertEqual(find_alias("go build", aliases), None)
        self.assertEqual(find_alias("go test", aliases), "got")
        self.assertEqual(find_alias("go test -v -count=1 ./...", aliases), "got -v -count=1 ./...")
        self.assertEqual(find_alias("go test -race ./...", aliases), "gotr ./...")
        self.assertEqual(find_alias("go mod", aliases), "gom")

    def test_suggest_alias(self):
        self.assertEqual(suggest_alias("go test -v", 10, []), None)
        self.assertEqual(suggest_alias("go vet", 33, [" 1 go test -v", " 2 go run ./...", " 3 go build ./..."]), None)
        self.assertEqual(suggest_alias("go run ./...", 10, [" 1 go test -v", " 2 go run ./...", " 3 go run ./...", " 4 go run ./...", " 5 go run ./...", " 6 go build ./..."]), "go run ./...")

if __name__ == '__main__':
    main()
