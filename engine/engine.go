package engine

import (
	"bufio"
	"code.google.com/p/gopass"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/shinpei/comstock/model"
	"github.com/shinpei/comstock/storage"
	"log"
	"os"
	"strconv"
)

const (
	Version string = "0.0.1"
	AppName string = "comstock"
)

// this is TODO.
// How can we pass thunk to cli?
var (
	eng *Engine
)

type Engine struct {
	App      *cli.App
	storager storage.Storager // storage
	logined  bool
	config   *Config
	env      *Env
}

func (e *Engine) Logined() bool {
	return e.logined
}

func NewEngine() *Engine {
	env := CreateEnv()
	var config *Config
	configPath := env.compath + "/" + ConfigFileDefault
	if IsFileExist(configPath) {
		config = LoadConfig(configPath)
		fmt.Println("Config loaded")
	}
	eng = &Engine{
		App:      initApp(),
		storager: storage.CreateFileStorager(env.compath),
		logined:  false,
		env:      env,
		config:   config,
	}
	return eng
}

func initApp() *cli.App {
	app := cli.NewApp()
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
			Name:      "get",
			ShortName: "g",
			Usage:     "Get command by specifiying number",
			Action: func(c *cli.Context) {
				if len(c.Args()) == 0 {
					println("'run' requires #number argument, e.g., 'comstock run 1'")
					return
				}
				num, _ := strconv.Atoi(c.Args()[0])
				cmd := eng.FetchCommandFromNumber(num)
				println(cmd.Cmd())
			},
		},
		{
			Name:      "run",
			ShortName: "r",
			Usage:     "Exec command with #number (experiment)",
			Action: func(c *cli.Context) {
				if len(c.Args()) == 0 {
					println("'run' requires #number argument, e.g., 'comstock run 1'")
					return
				}
				num, _ := strconv.Atoi(c.Args()[0])
				cmd := eng.FetchCommandFromNumber(num)
				println(cmd.Cmd())

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

func (e *Engine) Stock(cmd *model.Command) {
	// save to the local storage
	e.storager.Push(e.env.compath, cmd)
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

func (e *Engine) FetchCommandFromNumber(num int) (cmd *model.Command) {
	cmd = e.storager.FetchCommandFromNumber(num)
	return
}
