package engine

import (
	"errors"
	"fmt"
	"github.com/codegangsta/cli"
)

var ImportCommand cli.Command = cli.Command{
	Name:  "import",
	Usage: "Import from zshell files",
	Flags: []cli.Flag{
		cli.StringFlag{Name: "shell, s", Value: "", Usage: "Forcly change the shell handler. for example, bash compatible shell can use 'bash'"},
	},
	Action: ImportAction,
}

func ImportAction(c *cli.Context) {
	shellstr := c.String("shell")
	if shellstr != "" {
		eng.env.Shell = shellstr
	}
	err := eng.Import()
	if err != nil {
		fmt.Println("Command fialed: " + err.Error())
	}
}

func (e *Engine) Import() (err error) {

	if e.isLogin == false {
		err = errors.New("Login required")
		return
	}

	histFile := e.env.Homepath
	handler, histFile := FetchShellHandler(e, histFile)
	cmds, err := handler.ReadEveryHistory(histFile)
	for idx, cmd := range cmds {
		eng.Save(cmd)
		if idx > 2 {
			break
		}
		println(idx, cmd)
	}
	return
}
