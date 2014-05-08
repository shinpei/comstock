package main

import (
	"bufio"
	"code.google.com/p/gopass"
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"strings"
)

const (
	Version string = "0.0.1"
	AppName string = "comstock"
)

var (
	com *Comstock
)

type Comstock struct {
	App *cli.App

	storager Storager // storage
	logined  bool
	config   Config
}

func (c *Comstock) Logined() bool {
	return c.logined
}

func NewComstock() *Comstock {
	f := &FileStorager{}
	f.Open()
	return &Comstock{
		App:      initApp(),
		storager: &FileStorager{},
		logined:  false,
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
			Name:      "save",
			ShortName: "sv",
			Usage:     "Save previous command",
			Action: func(c *cli.Context) {
				env := CreateEnv()
				home := env.HomePath()
				var shellHistoryFilename string = home
				var handler Shell = nil
				if strings.Contains(env.Shell(), "zsh") {
					shellHistoryFilename += "/.zsh_history"
					handler = &ZshHandler{}
				} else if strings.Contains(env.Shell(), "bash") {
					shellHistoryFilename += "/.bash_history"
					handler = &BashHandler{}
				}
				cmd, err := handler.ReadLastHistory(shellHistoryFilename)
				if err != nil {
					log.Fatal(err)
				}
				com.Stock(cmd)
			},
			BashComplete: func(c *cli.Context) {
				if len(c.Args()) > 0 {
					return
				}
			},
		},
		{
			Name:        "list",
			ShortName:   "ls",
			Description: "Show the list of stocked commands",
			Usage:       "List stocked command",
			Action: func(c *cli.Context) {
				//				args := c.Args()

				com.List()
			},
		},
		{
			Name:  "push",
			Usage: "Push stocked command to cloud",
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
		{
			Name:  "login",
			Usage: "Login to the cloud",
			Action: func(c *cli.Context) {
				if !com.Logined() {
					scanner := bufio.NewScanner(os.Stdin)
					fmt.Printf("Your registered email address? : ")
					scanner.Scan()
					username := scanner.Text()
					fmt.Printf("And password? :")
					password, _ := gopass.GetPass("")
					com.Login(username, password)
				}
			},
		},
		{
			Name:  "config",
			Usage: "Get and set comstock options",
			Action: func(c *cli.Context) {
				LoadConfig()
				com.ShowConfig()
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
	c.storager.Push(cmd)
	/*
		{
			// push to the internet
			c.PushToRemote()
		} else {
			c.PushToLocal(cmd)
		}
	*/
	fmt.Printf("[%s]Saved command '%s'\n", c.storager.StorageType(), cmd.Cmd())
}

// Push
func (c *Comstock) PushToRemote() {

}

func (c *Comstock) PushToLocal(cmd *Command) {

	c.storager.Push(cmd)
}

func (c *Comstock) Close() {
	c.storager.Close()
}

func (c *Comstock) List() {
	// c.storager.PullCommands()
	if err := c.storager.List(); err != nil {
		log.Fatal(err)
	}
}

func (c *Comstock) Login(username string, password string) string {
	if c.Logined() {
		return "access token"
	} else {
		println("logging in...")
		return "access success"
	}
}

func (c *Comstock) ShowConfig() {
	println("Showing config")
}
