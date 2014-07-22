package engine

import (
	"errors"
	"fmt"
	"github.com/shinpei/comstock/model"
	"log"
	"strings"
)

func (e *Engine) Save(command string) (err error) {
	if e.isLogin == false {
		err = errors.New("Login required")
		return
	}
	var shellHistoryFilename string = e.env.Homepath
	var handler Shell = nil
	if strings.Contains(e.env.Shell, "zsh") {
		shellHistoryFilename += "/.zsh_history"
		handler = &ZshHandler{}
	} else if strings.Contains(e.env.Shell, "bash") {
		shellHistoryFilename += "/.bash_history"
		handler = &BashHandler{}
	} else {
		log.Fatal("Couldn't recognize your shell. Please report your environment through 'comstock sos'")
	}
	var cmd *model.Command
	if command != "" {
		cmd = model.CreateCommand(command)
	}
	fmt.Println(handler)
	//	cmd, err = handler.ReadLastHistory(shellHistoryFilename)
	if err != nil {
		log.Fatal(err)
	}
	cmd.Cmd = strings.TrimSpace(cmd.Cmd)
	// save to the local storage
	// remove whitespaces from cmd

	err = e.storager.Push(e.userinfo, e.env.Compath, cmd)

	if e.config.verboseMode {
		fmt.Printf("[%s]Saved command '%s'\n", e.storager.StorageType(), cmd.Cmd)
	} else {
		fmt.Printf("Saved command '%s'\n", cmd.Cmd)
	}
	return
}
