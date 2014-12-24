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
	if !eng.storager.IsRequireLogin() {
		fmt.Println("You don't need to login for storager you set:", eng.storager.StorageType())
	}
	if eng.userinfo == nil {
		fmt.Println("You're already logged out.")
		return
	}

	eng.isLoginPolled = true
	eng.Logout(eng.apiServer)
	fmt.Println("Logout, done.")

}

func (e *Engine) Logout(loginServer string) {
	eng.SetLogout()
	return
}
