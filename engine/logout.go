package engine

import (
	"fmt"
	"github.com/codegangsta/cli"
)

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
