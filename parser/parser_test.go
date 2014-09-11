package parser

import (
	. "github.com/shinpei/comstock/test"
	"strings"
	"testing"
)

func TestParsePlain(t *testing.T) {
	line := "comstock save"
	cmds, err := Parse(line)
	if err != nil {
		t.Error("Cannot parse")
	}
	if b := AssertEqual(t, len(cmds), 1); b {
		AssertEqual(t, cmds[0], "comstock save")
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
	if b := AssertEqual(t, len(answers), len(cmds)); b {
		for idx, s := range answers {
			AssertEqual(t, s, cmds[idx])
		}
	}
}

func TestParsePipe(t *testing.T) {
	helperSimplePipe(t, "a|b")
	helperSimplePipe(t, "ps aux | grep tmux")
	helperSimplePipe(t, "echo \"   hihihi   \"  | sed -e 's/^[ \t]*//'")
}
