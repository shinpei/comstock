package test

import (
	"testing"
)

func AssertEqual(t *testing.T, expected interface{}, value interface{}) bool {
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
