package engine

import (
	"testing"
)

func TestCreateEnvField(t *testing.T) {
	env := NewEnv()
	if env != nil {
		// pass
	} else {
		t.Error("Cannot create env")
	}
}
