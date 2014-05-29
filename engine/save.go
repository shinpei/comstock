package engine

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

func (e *Engine) Save(home string, shell string) (err error) {
	if e.isLogin == false {
		err = errors.New("Login required")
		return
	}
	var shellHistoryFilename string = home
	var handler Shell = nil
	if strings.Contains(shell, "zsh") {
		shellHistoryFilename += "/.zsh_history"
		handler = &ZshHandler{}
	} else if strings.Contains(shell, "bash") {
		shellHistoryFilename += "/.bash_history"
		handler = &BashHandler{}
	}
	cmd, err := handler.ReadLastHistory(shellHistoryFilename)
	if err != nil {
		log.Fatal(err)
	}
	cmd.Cmd = strings.TrimSpace(cmd.Cmd)
	// save to the local storage
	// remove whitespaces from cmd
	err = e.storager.Push(e.userinfo, e.env.compath, cmd)
	fmt.Printf("[%s]Saved command '%s'\n", e.storager.StorageType(), cmd.Cmd)
	return
}
