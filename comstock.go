package main

import (
	"github.com/codegangsta/cli"
	"os"
	"strings"
	"log"
)

var tasks = []string{"cook", "clean"}
func main() {
	app := cli.NewApp();
	app.EnableBashCompletion = true;

	app.Name = "comstock"
	app.Usage = "save your command"
	app.Action = func (c *cli.Context) {
		println("comstock: error: command is missing");
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
		println(shellHistoryFilename);
		line, err := handler.ReadLastHistory(shellHistoryFilename);
		if err != nil {
			log.Fatal(err);
		}
		println(line);
	}

	app.Commands = []cli.Command {
		{
			Name: "save",
			ShortName: "s",
			Usage: "save former command ",
			Action: func (c *cli.Context) {
				println("save command to comstock cloud")
			},
			BashComplete: func(c *cli.Context) {
				if len(c.Args()) > 0 {
					return
				}
			    for _, t := range tasks {
					println(t)
				}
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
