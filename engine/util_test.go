package engine

import (
	"testing"
)

func TestIsFileExist(t *testing.T) {
	if IsFileExist("./") {

	} else {
		t.Error("Couldn't stat current directory?")
	}
}
