package engine

import (
	"os/exec"
	"strings"
)

func (e *Engine) Open(URL string) (err error) {
	if strings.Contains(e.env.OS, "darwin") {
		cmd := exec.Command("open", URL)
		err = cmd.Start()
	}

	return
}
