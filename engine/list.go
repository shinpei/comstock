package engine

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/shinpei/comstock/model"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var ListCommand cli.Command = cli.Command{
	Name:        "list",
	ShortName:   "ls",
	Description: "List stocked commands",
	Usage:       "List stocked command, default value is 15.",
	Flags: []cli.Flag{
		cli.IntFlag{Name: "number,n", Value: 15, Usage: "Give a number of how many items to list."}},
	Action: ListAction,
}

func ListAction(c *cli.Context) {
	err := eng.List()
	if err != nil {
		fmt.Println("Command failed: ", err.Error())
	}
}

func (e *Engine) List() (err error) {
	e.IsRequireLoginOrDie()
	var hists []model.NaiveHistory
	if hists, err = e.storager.List(e.userinfo); err != nil {
		if _, ok := err.(*model.SessionExpiresError); ok {
			e.SetLogout()
		} else if _, ok := err.(*model.SessionInvalidError); ok {
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

	for _, hist := range hists {
		idx++
		cmd := hist.Command()
		if ttyWidth < len(cmd) {
			fmt.Printf("%d: %s...\n", idx, cmd[:ttyWidth-15])
		} else {
			fmt.Printf("%d: %s\n", idx, cmd)
		}
	}
	return
}
