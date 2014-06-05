package engine

import (
	"os/exec"
	"strings"
)

// codes from link below
// http://stackoverflow.com/a/14053693/3070610
func (e *Engine) Open(URL string) (err error) {
	var browserCommand string = ""

	if strings.Contains(e.env.OS, "darwin") {
		browserCommand = "open"
	} else if strings.Contains(e.env.OS, "linux") || strings.Contains(e.env.OS, "bsd") {
		browserCommand = "xdg-open"
	} else if strings.Contains(e.env.OS, "mingw") {
		browserCommand = "start"
	}
	cmd := exec.Command(browserCommand, URL)
	err = cmd.Start()

	return
}
