package engine

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
	eng *Engine
)

type Engine struct {
	App      *cli.App
	storager Storager // storage
	logined  bool
	config   *Config
	env      *Env
}

func (e *Engine) Logined() bool {
	return e.logined
}

func NewEngine() *Engine {
	f := &FileStorager{}
	f.Open()
	env := CreateEnv()
	var config *Config
	configPath := env.compath + "/" + ConfigFileDefault
	if IsFileExist(configPath) {
		config = LoadConfig(configPath)
		fmt.Println("Config loaded")
	}
	eng = &Engine{
		App:      initApp(),
		storager: &FileStorager{},
		logined:  false,
		env:      env,
		config:   config,
	}
	return eng
}

func initApp() *cli.App {
	app := cli.NewApp()
	//app.EnableBashCompletion = true
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
				eng.Save(eng.env.HomePath(), eng.env.Shell())
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

				eng.List()
			},
		},
		{
			Name:      "run",
			ShortName: "r",
			Usage:     "Exec command with #number",
			Action: func(c *cli.Context) {
				if len(c.Args()) == 0 {
					println("'run' requires #number argument, e.g., 'comstock run 1'")
					return
				}
				println("Hello", c.Args()[0])
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
				if !eng.Logined() {
					scanner := bufio.NewScanner(os.Stdin)
					fmt.Printf("Your registered email address? : ")
					scanner.Scan()
					username := scanner.Text()
					fmt.Printf("And password? :")
					password, _ := gopass.GetPass("")
					eng.Login(username, password)
				}
			},
		},
		{
			Name:  "config",
			Usage: "Get and set comstock options",
			Action: func(c *cli.Context) {
				eng.ShowConfig()
			},
		},
	}

	return app
}

func (e *Engine) Run(args []string) {
	e.App.Run(args)
	e.Close()
}

func (e *Engine) Stock(cmd *Command) {
	// save to the local storage
	e.storager.Push(cmd)
	/*
		{
			// push to the internet
			e.PushToRemote()
		} else {
			e.PushToLocal(cmd)
		}
	*/
	fmt.Printf("[%s]Saved command '%s'\n", e.storager.StorageType(), cmd.Cmd())
}

// Push
func (e *Engine) PushToRemote() {

}

func (e *Engine) PushToLocal(cmd *Command) {

	e.storager.Push(cmd)
}

func (e *Engine) Close() {
	e.storager.Close()
}

func (e *Engine) List() {
	// e.storager.PullCommands()
	if err := e.storager.List(); err != nil {
		log.Fatal(err)
	}
}

func (e *Engine) Login(username string, password string) string {
	if e.Logined() {

		return "access token"
	} else {
		println("We're preparing this func...")
		return "access success"
	}
}

func (e *Engine) ShowConfig() {
	println("Showing config")
}

func (e *Engine) Save(home string, shell string) {
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
	eng.Stock(cmd)
}
