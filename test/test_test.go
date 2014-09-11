package test

import (
	"testing"
)

func TestEqualString(t *testing.T) {
	AssertEqual(t, "Hi", "Hi")
	s := "Test String"
	AssertEqual(t, "Test String", s)
}

func TestNotEqualString(t *testing.T) {
	AssertNotEqual(t, "Hi", "bye")
}

func TestEqualInt(t *testing.T) {
	AssertEqual(t, 12345, 12345)
	AssertNotEqual(t, 12345, 56789)
}
