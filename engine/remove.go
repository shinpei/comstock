package engine

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/shinpei/comstock/model"
	"log"
	"strconv"
)

var RemoveCommand cli.Command = cli.Command{
	Name:      "remove",
	ShortName: "rm",
	Usage:     "Delete stocked command by specifiying command number, e.g. 'comstock rm 16'",
	Action:    RemoveAction,
}

func RemoveAction(c *cli.Context) {
	if len(c.Args()) == 0 {
		fmt.Println("'remove' requires #number argument, e.g., 'comstock rm 1'")
		return
	}
	index, err := strconv.Atoi(c.Args()[0])
	if err != nil {
		fmt.Println("Invalid argument was given, please retry")
		return
	}
	if err := eng.Remove(index); err != nil {
		fmt.Println("Command failed:", err.Error())
		return
	}
	fmt.Println("Successfully remove command #", index)
}

func (e *Engine) Remove(index int) (err error) {
	e.IsRequireLoginOrDie()
	if index < 1 {
		log.Fatal("You cannot specify index as index < 1")
	}
	if err = e.storager.RemoveOne(e.userinfo, index); err != nil {
		if _, ok := err.(*model.SessionExpiresError); ok {
			e.SetLogout()
		} else if _, ok := err.(*model.SessionInvalidError); ok {
			e.SetLogout()
		}
		return
	}
	return
}
