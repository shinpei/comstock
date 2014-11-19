package engine

import (
	"fmt"
	"github.com/codegangsta/cli"
)

var PushCommand cli.Command = cli.Command{
	Name:   "push",
	Usage:  "push commands to the cloud",
	Action: PushAction,
}

func PushAction(c *cli.Context) {
	err := eng.Push()
	if err != nil {
		fmt.Println("Command failed: ", err.Error())
	} else {
		fmt.Println("push succsesed for ", eng.AuthInfo())
	}
}

func (e *Engine) Push() (err error) {
	e.IsRequireLoginOrDie()

	return
}
