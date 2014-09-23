package engine

import (
	"fmt"
	"github.com/codegangsta/cli"
)

var RawCommand cli.Command = cli.Command{
	Name:   "raw",
	Usage:  "Tap raw command",
	Action: RawAction,
}

// 'raw' is an action which sends raw commands to the server
func RawAction(c *cli.Context) {
	first := c.Args().First()

	err := eng.Raw(eng.apiServer, first)
	if err != nil {
		fmt.Println("Seems, it's failed")
	}
}

func (e *Engine) Raw(URL string, cmd string) (err error) {
	browserCommand := cmd
	println(browserCommand)
	return
}
