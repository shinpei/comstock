package engine

import (
	"bufio"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/shinpei/comstock/model"
	"github.com/shinpei/comstock/storage"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

const (
	Version  string = "0.1.0"
	AppName  string = "comstock"
	AuthFile string = "authinfo"
)

// this is TODO.
// How can we pass thunk to cli?
var (
	eng *Engine
)

type Engine struct {
	App      *cli.App
	storager storage.Storager // storage
	isLogin  bool
	authInfo string
	config   *Config
	env      *Env
}

func (e *Engine) IsLogin() bool {
	return e.isLogin
}
func (e *Engine) SetLogin() {
	e.isLogin = true
}
func (e *Engine) SetLogout() {
	e.isLogin = false
}
func (e *Engine) AuthInfo() string {
	return e.authInfo
}
func (e *Engine) SetAuthInfo(auth string) {
	e.authInfo = auth
}

func NewEngine() *Engine {
	env := CreateEnv()
	var config *Config
	configPath := env.compath + "/" + ConfigFileDefault
	if IsFileExist(configPath) {
		config = LoadConfig(configPath)
	}
	var s storage.Storager
	switch config.Local.Type {
	case "file":
		s = storage.CreateFileStorager(env.compath)
	case "mongo":
		s = storage.CreateMongoStorager()
	default:
		s = storage.CreateFileStorager(env.compath)
	}
	var isAlreadyLogin bool = false
	authinfo := readAuthInfo(eng.env)
	if authinfo != "" {
		isAlreadyLogin = true
	}

	eng = &Engine{
		App:      initApp(),
		authInfo: authinfo,
		isLogin:  isAlreadyLogin,
		storager: s,
		env:      env,
		config:   config,
	}

	return eng
}

func readAuthInfo(env *Env) string {
	authFilePath := env.compath + "/" + AuthFile
	fi, _ := os.Open(authFilePath)
	scanner := bufio.NewScanner(fi)
	var lc int = 0
	var authinfo string = ""
	for scanner.Scan() {
		lc++
		if 1 < lc {
			// error
			log.Fatal("Invalid login info")
		}
		authinfo = scanner.Text()
	}
	return authinfo
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
			Name:      "status",
			ShortName: "st",
			Usage:     "Show comstock status",
			Action: func(c *cli.Context) {
				//TODO
				log.Fatal("Not yet implemented")
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
				println(cmd.Cmd)
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
				// TODO fix with parseint
				num, err := strconv.Atoi(c.Args()[0])
				if err != nil {
					log.Fatal(err)
				}
				cmd := eng.FetchCommandFromNumber(num)
				println(cmd.Cmd)

			},
		},
		{
			Name:  "login",
			Usage: "Login to the cloud",
			Action: func(c *cli.Context) {
				if eng.IsLogin() {
					fmt.Printf("Already login as %s\n", eng.config.User.Mail)
					return
				}
				eng.Login()
			},
		},
		{
			Name:  "config",
			Usage: "Show comstock configuration",
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
	// remove whitespaces from cmd
	e.storager.Push(e.env.compath, cmd)
	/*
		{
			// push to the internet
			e.PushToRemote()
		} else {
			e.PushToLocal(cmd)
		}
	*/
	fmt.Printf("[%s]Saved command '%s'\n", e.storager.StorageType(), cmd.Cmd)
}

func (e *Engine) Close() {
	e.storager.Close()

	// write needed info
	if e.IsLogin() {
		authFilePath := e.env.compath + "/" + AuthFile
		authinfo := []byte(e.AuthInfo())
		ioutil.WriteFile(authFilePath, authinfo, 0644)
		return
	}

}

func (e *Engine) List() {
	// e.storager.PullCommands()
	if err := e.storager.List(); err != nil {
		log.Fatal(err)
	}
}

func (e *Engine) ShowConfig() {
	e.config.ShowConfig()
}

func (e *Engine) FetchCommandFromNumber(num int) (cmd *model.Command) {

	cmd = e.storager.FetchCommandFromNumber(num)
	return
}
