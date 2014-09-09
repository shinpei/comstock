package parser

import (
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

func TestParsePipe(t *testing.T) {
	line := "a|b"
	cmds, err := Parse(line)
	if err != nil {
		t.Error("Cannot parse")
	}
	if b := assertEqual(t, len(cmds), 2); b {
		assertEqual(t, cmds[0], "a")
		assertEqual(t, cmds[1], "b")
	}

}
