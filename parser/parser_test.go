package parser

import (
	"strings"
	"testing"
)

func assertEqual(t *testing.T, expected interface{}, value interface{}) bool {
	if _, ok := expected.(string); ok {
		// string comparison
		var e string = string(expected.(string))
		var v string = string(value.(string))
		if e != v {
			t.Error("Expected value is not equal to the given value")
			return false
		}
	} else if _, ok := expected.(int); ok {
		var e int = int(expected.(int))
		var v int = int(value.(int))
		if e != v {
			t.Error("Expected value is not equal to the given value")
		}
		return false
	}
	return true
}

func TestParsePlain(t *testing.T) {
	line := "comstock save"
	cmds, err := Parse(line)
	if err != nil {
		t.Error("Cannot parse")
	}
	if b := assertEqual(t, len(cmds), 1); b {
		assertEqual(t, cmds[0], "comstock save")
	}
}

func helperSimplePipe(t *testing.T, line string) {
	answers := strings.Split(line, "|")
	if len(answers) < 2 {
		t.Error("argument for helper func seems wrong, please give line contains at least one '|'")
		return
	}
	cmds, err := Parse(line)
	if err != nil {
		t.Error("Couldn't parse:", line)
	}
	if b := assertEqual(t, len(answers), len(cmds)); b {
		for idx, s := range answers {
			assertEqual(t, s, cmds[idx])
		}
	}
}

func TestParsePipe(t *testing.T) {
	helperSimplePipe(t, "a|b")
	helperSimplePipe(t, "ps aux | grep tmux")
	helperSimplePipe(t, "echo \"   hihihi   \"  | sed -e 's/^[ \t]*//'")
}
