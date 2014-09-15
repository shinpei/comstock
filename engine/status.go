package engine

import (
	"fmt"
	"github.com/codegangsta/cli"
)

var StatusCommand cli.Command = cli.Command{
	Name:      "status",
	ShortName: "st",
	Usage:     "Show comstock status",
	Action: func(c *cli.Context) {
		//TODO
		eng.Status()
	},
}

func (e *Engine) Status() {
	fmt.Println("[Comstock envieonment]")
	m := e.env.AsMap()
	for k, v := range m {
		fmt.Println(k, ":", v)
	}
	fmt.Println("")
	fmt.Println("[Storager environment]")
	_ = e.storager.Status()
}
