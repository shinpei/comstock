package engine

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/mattn/go-isatty"
	"github.com/shinpei/comstock/model"
	"github.com/shinpei/comstock/parser"
	"log"
	"os"
	"strings"
)

func SaveAction(c *cli.Context) {
	first := c.Args().First()
	//fmt.Println("first: ", first)
	shellstr := c.GlobalString("shell")
	if shellstr != "" {
		eng.env.Shell = shellstr
	}
	err := eng.Save(first)
	if err != nil {
		fmt.Println("Command failed: ", err)
	}
}

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
	// split with parser
	commands, _ := parser.Parse(command)
	for _, cmdStr := range commands {
		cmd = model.CreateCommand(cmdStr)
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

	}
	return
}
