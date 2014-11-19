package engine

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/shinpei/comstock/model"
	"log"
	"strconv"
)

var FetchCommand cli.Command = cli.Command{
	Name:   "get",
	Usage:  "Get command by specifiying number",
	Action: FetchAction,
}

func FetchAction(c *cli.Context) {

	if len(c.Args()) == 0 {
		fmt.Println("'get' requires #number argument, e.g., 'comstock get 1'.")
		return
	}
	num, err := strconv.Atoi(c.Args()[0])
	if err != nil {
		fmt.Println("Invalid argument was given, please retry")
		return
	}
	nh, err := eng.FetchCommandFromNumber(num)
	if err != nil {
		fmt.Println("Command failed: ", err.Error())
	} else {
		fmt.Println(nh.Cmds[0])
	}
}

func (e *Engine) FetchCommandFromNumber(num int) (nh *model.NaiveHistory, err error) {
	if e.storager.IsRequireLogin() == true && e.isLogin == false {
		log.Fatal("You have no valid access token. Please login first.")
	}
	nh, err = e.storager.FetchFromNumber(e.userinfo, num)
	if _, ok := err.(*model.SessionExpiresError); ok {
		e.SetLogout()
	}
	if _, ok := err.(*model.SessionInvalidError); ok {
		e.SetLogout()
	}
	return
}
