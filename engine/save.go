package engine

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/mattn/go-isatty"
	"github.com/shinpei/comstock/model"
	"log"
	"os"
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
		//		log.Fatal("Couldn't recognize your shell. Please report your environment through 'comstock sos'")
		log.Fatal("Couldn't recognize your shell. Your env is ", e.env.Shell)
	}
	var cmd *model.Command

	//check weather command has given
	if command == "" {
		if isatty.IsTerminal(os.Stdin.Fd()) {
			command, err = handler.ReadLastHistory(shellHistoryFilename)
		} else {
			// data arrived in stdin
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				// TODO: not only read one line
				command = scanner.Text()
			} else {
				log.Fatal("No command given")
			}
		}
	}
	cmd = model.CreateCommand(command)
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
