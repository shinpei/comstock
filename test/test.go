package test

import (
	"testing"
)

func assertEqual(t *testing.T, expected interface{}, value interface{}) bool {
	if _, ok := expected.(string); ok {
		// string comparison
		var e string = string(expected.(string))
		var v string = string(value.(string))
		if e != v {
			return false
		}
	} else if _, ok := expected.(int); ok {
		var e int = int(expected.(int))
		var v int = int(value.(int))
		if e != v {
			return false
		}
	} else if _, ok := expected.(float64); ok {

	}
	return true

}

func AssertEqual(t *testing.T, expected interface{}, value interface{}) bool {

	if ok := assertEqual(t, expected, value); !ok {
		t.Error("Expected value is not equal to the given value")
		return false
	}
	return true
}

// Maybe not needed...
func AssertNotEqual(t *testing.T, expectedNot interface{}, value interface{}) bool {

	if ok := assertEqual(t, expectedNot, value); ok {
		t.Error("Unexpectedly the given value is matched with expectedNot value")
		return false

	}
	return true
}
