package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"strings"
)

const (
	Version                string = "0.0.1"
	AppName                string = "comstock"
	ComstockConfigFilename string = "comstock.yaml"
)

var (
	com *Comstock
)

type Comstock struct {
	App               *cli.App
	connectionEnabled bool
	lStorager         LocalStorager /* local storage */
}

func NewComstock() *Comstock {
	f := &FileStorager{}
	f.Open()
	return &Comstock{
		App:               initApp(),
		lStorager:         &FileStorager{},
		connectionEnabled: false,
	}
}

func initApp() *cli.App {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Version = Version
	app.Name = AppName
	app.Usage = "save your command to the cloud"
	app.Action = func(c *cli.Context) {
		println("comstock: error: command is missing. For more details, see 'comstock -h'")
	}
	app.Commands = []cli.Command{
		{
			Name:      "stock",
			ShortName: "s",
			Usage:     "Stock the former command into appropriate storage",
			Action:    nil,
		},
		{
			Name:  "save",
			Usage: "Alias for 'stock'",
			Action: func(c *cli.Context) {
				home := os.Getenv("HOME")
				shell := os.Getenv("SHELL")
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
				com.Stock(cmd)
				fmt.Printf("saved command '%s'\n", cmd.Cmd)
			},
			BashComplete: func(c *cli.Context) {
				if len(c.Args()) > 0 {
					return
				}
			},
		},
		{
			Name:      "list",
			ShortName: "l",
			Usage:     "List stocked command",
			Action: func(c *cli.Context) {
				println("listed")
			},
		},
		{
			Name:      "push",
			ShortName: "p",
			Usage:     "Push stocked command to cloud",
			Action: func(c *cli.Context) {
				println("pushed")
			},
		},
		{
			Name:  "pop",
			Usage: "Pop last stocked command",
			Action: func(c *cli.Context) {
				println("poped")
			},
		},
	}

	return app
}

func (c *Comstock) Run(args []string) {
	c.App.Run(args)
	c.Close()
}

func (c *Comstock) Stock(cmd *Command) {
	// save to the local storage
	if c.connectionEnabled {
		// push to the internet
		c.PushToRemote()
	} else {
		c.PushToLocal(cmd)
	}
	println(cmd.Cmd())
}

// Push
func (c *Comstock) PushToRemote() {

}

func (c *Comstock) PushToLocal(cmd *Command) {

	c.lStorager.Push(cmd)
}

func (c *Comstock) Close() {
	c.lStorager.Close()
}
