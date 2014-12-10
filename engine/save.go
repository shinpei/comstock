package engine

import (
	"bufio"
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
	Name:  "save",
	Usage: "Save last executed command",
	Flags: []cli.Flag{
		cli.StringFlag{Name: "shell, s", Value: "", Usage: "Forcly change the shell handler. for example, bash compatible shell can use 'bash'"},
	},
	Action: SaveAction,
}

func SaveAction(c *cli.Context) {

	first := c.Args().First()
	shellstr := c.String("shell")
	if shellstr != "" {
		eng.env.Shell = shellstr
	}
	err := eng.Save(first)
	if err != nil {
		fmt.Println("Command failed: ", err)
	}
}

func (e *Engine) Save(command string) (err error) {

	e.IsRequireLoginOrDie()
	shellHistoryFilename := e.env.Homepath
	handler, shellHistoryFilename := FetchShellHandler(e, shellHistoryFilename)

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
	cmds := []string{}
	for _, cmdStr := range commands {
		cmds = append(cmds, strings.TrimSpace(cmdStr))
	}
	nh := model.CreateNaiveHistory(cmds, "")
	nh.SetShell(e.env.Shell)
	// save to the local storage
	// remove whitespaces from cmd

	err = e.storager.Push(e.userinfo, e.env.Compath, nh)
	if e.config.verboseMode {
		fmt.Printf("[%s]Saved command '%s'\n", e.storager.StorageType(), nh.Command())
	} else {
		fmt.Printf("Saved command '%s'\n", nh.Command())
	}

	return
}
