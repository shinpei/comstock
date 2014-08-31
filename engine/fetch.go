package engine

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/shinpei/comstock/model"
	"log"
	"strconv"
)

func (e *Engine) FetchAction(c *cli.Context) {
	if len(c.Args()) == 0 {
		fmt.Println("'get' requires #number argument, e.g., 'comstock get 1'.")
		return
	}
	num, err := strconv.Atoi(c.Args()[0])
	if err != nil {
		fmt.Println("Invalid argument was given, please retry")
		return
	}
	cmd, err := e.FetchCommandFromNumber(num)
	if err != nil {
		fmt.Println("Command failed: ", err.Error())
	} else {
		fmt.Println(cmd.Cmd)
	}

}

func (e *Engine) FetchCommandFromNumber(num int) (cmd *model.Command, err error) {
	if e.storager.IsRequireLogin() == true && e.isLogin == false {
		log.Fatal("You have no valid access token. Please login first.")
	}
	cmd, err = e.storager.FetchCommandFromNumber(e.userinfo, num)
	if err == model.ErrSessionExpires || err == model.ErrSessionInvalid {
		e.SetLogout()
	}
	return
}
