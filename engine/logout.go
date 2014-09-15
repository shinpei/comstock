package engine

import (
	"fmt"
	"github.com/codegangsta/cli"
)

var LogoutCommand cli.Command = cli.Command{
	Name:   "logout",
	Usage:  "Logout from current account",
	Action: LogoutAction,
}

func LogoutAction(c *cli.Context) {
	if eng.IsLogin() == false {
		fmt.Println("Already logout.")
		return
	}
	eng.Logout(eng.apiServer)
	fmt.Println("Logout, done.")

}

func (e *Engine) Logout(loginServer string) {
	eng.SetLogout()
	return
}
