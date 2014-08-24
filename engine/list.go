package engine

import (
	"fmt"
	"github.com/shinpei/comstock/model"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func (e *Engine) List() (err error) {
	if e.storager.IsRequireLogin() == true && e.isLogin == false {
		log.Fatal("You have no valid access token. Please login first.")
	}
	var cmds []model.Command
	if cmds, err = e.storager.List(e.userinfo); err != nil {
		if err == model.ErrSessionExpires {
			e.SetLogout()
		} else if err == model.ErrSessionInvalid {
			e.SetLogout()
		}
	}
	var idx int = 0

	// Modify printing size due to the terminal width, if it's not enough, '...' will be used
	// can be concurrently exec
	sttyCmd := exec.Command("stty", "size")
	sttyCmd.Stdin = os.Stdin

	out, errStty := sttyCmd.Output()
	if errStty != nil {
		log.Fatal(errStty)
	}
	ttyWidthStr := strings.Replace(strings.Split(string(out), " ")[1], "\n", "", 1)
	ttyWidth, _ := strconv.Atoi(ttyWidthStr)

	for _, cmd := range cmds {
		idx++
		if ttyWidth < len(cmd.Cmd) {
			fmt.Printf("%d: %s...\n", idx, cmd.Cmd[:ttyWidth-15])
		} else {
			fmt.Printf("%d: %s\n", idx, cmd.Cmd)
		}
	}
	return
}
