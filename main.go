package main

import (
	"github.com/codegangsta/cli"
	"os"
	"strings"
	"log"
)

func main() {
	app := cli.NewApp();
	app.EnableBashCompletion = true;
	app.Version = "0.0.1";
	app.Name = "comstock"
	app.Usage = "save your command"
	app.Action = func (c *cli.Context) {
		println("comstock: error: command is missing");
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
				var handler ShellHandler = nil;
				if strings.Contains(shell, "zsh") {
					shellHistoryFilename += "/.zsh_history";
					handler = &ZshHandler{};
				} else if strings.Contains(shell, "bash") {
					shellHistoryFilename += "/.bash_history";
					handler = &BashHandler{};
				}
				line, err := handler.ReadLastHistory(shellHistoryFilename);
				if err != nil {
					log.Fatal(err);
				}
				//comstock.Stock(line)
				println(line)
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
	};

	app.Run(os.Args);
}

