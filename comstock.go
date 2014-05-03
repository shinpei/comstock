package main

import (
	"github.com/codegangsta/cli"
	"os"
	"strings"
	"log"
	"fmt"

)

const (
	Version string = "0.0.1"
	AppName string = "comstock"
	ComstockConfigFilename string = "comstock.yaml"
)

type Comstock struct {
	App *cli.App
	connectionEnabled bool
}

func NewComstock() *Comstock {
	app := cli.NewApp();
	app.EnableBashCompletion = true;
	app.Version = Version;
	app.Name = AppName;
	app.Usage = "save your command to the cloud"
	app.Action = func (c *cli.Context) {
		println("comstock: error: command is missing. For more details, see 'comstock -h'");
	}

	app.Commands = []cli.Command {
		{
			Name: "save",
			ShortName: "s",
			Usage: "save former command ",
			Action: func (c *cli.Context) {
				home := os.Getenv("HOME");
				shell := os.Getenv("SHELL");
				var shellHistoryFilename string = home;
				var handler Shell = nil;
				if strings.Contains(shell, "zsh") {
					shellHistoryFilename += "/.zsh_history";
					handler = &ZshHandler{};
				} else if strings.Contains(shell, "bash") {
					shellHistoryFilename += "/.bash_history";
					handler = &BashHandler{};
				}
				cmd, err := handler.ReadLastHistory(shellHistoryFilename);
				if err != nil {
					log.Fatal(err);
				}
				//comstock.Stock(line)
				fmt.Printf("saved command '%s'\n", cmd.Cmd);
			},
			BashComplete: func(c *cli.Context) {
				if len(c.Args()) > 0 {
					return
				}
		    },
		},
		{
			Name: "list",
			ShortName: "l",
			Usage: "List stocked command",
			Action:func (c *cli.Context) {
				println("listed");
			},
		},
		{
			Name: "push",
			ShortName: "p",
			Usage: "Push stocked command to cloud",
			Action: func(c *cli.Context) {
				println("pushed");
			},
		},
		{
			Name: "pop",
			Usage: "Pop last stocked command",
			Action:func (c *cli.Context) {
				println("poped");
			},
		},
	};

	return &Comstock{App: app};
}

func (this *Comstock) Run (args []string) {
	this.App.Run(args);
}

func (this *Comstock) Stock(command string) {
	// save to the local storage
	if this.connectionEnabled {
		// push to the internet
		this.PushToRemote();
	} else {
		this.PushToLocal();
	}
	println(command);
}
func (this *Comstock) PushToRemote() {

}

func (this *Comstock) PushToLocal() {
	

}
