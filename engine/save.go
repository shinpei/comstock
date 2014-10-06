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

var ActionCommand cli.Command = cli.Command{
	Name:      "save",
	ShortName: "sv",
	Usage:     "Save previous command",
	Action:    SaveAction,
}

func SaveAction(c *cli.Context) {

	first := c.Args().First()
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
	shellHistoryFilename := e.env.Homepath
	handler, shellHistoryFilename := FetchShellHandler(e, shellHistoryFilename)
	var cmd *model.Command

	//check weather command has given
	if command == "" {
		if isatty.IsTerminal(os.Stdin.Fd()) {
			command, err = handler.ReadLastHistory(shellHistoryFilename)
			if err != nil {
				return
			}
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
