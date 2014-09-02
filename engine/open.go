package engine

import (
	"github.com/codegangsta/cli"
	"os/exec"
	"strings"
)

func OpenAction(c *cli.Context) {
	err := eng.Open(eng.apiServer)
	if err != nil {
		fmt.Printf("Seems it's failed to execute browser, visit https://comstock.herokuapp.com directly")
	}
}

// codes from link below
// http://stackoverflow.com/a/14053693/3070610
func (e *Engine) Open(URL string) (err error) {
	browserCommand := ""

	if strings.Contains(e.env.OS, "darwin") {
		browserCommand = "open"
	} else if strings.Contains(e.env.OS, "linux") || strings.Contains(e.env.OS, "bsd") {
		browserCommand = "xdg-open"
	} else if strings.Contains(e.env.OS, "mingw") {
		// couldn't be?
		browserCommand = "start"
	}
	cmd := exec.Command(browserCommand, URL)
	err = cmd.Start()
	return
}
