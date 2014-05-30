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
	Version  string = "0.1.1"
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
	userinfo *model.UserInfo
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
	env := NewEnv()
	var config *Config
	configPath := env.compath + "/" + ConfigFileDefault
	var s storage.Storager
	s = storage.CreateCloudStorager()
	config = LoadConfig(configPath)
	switch config.Local.Type {
	case "file":
		s = storage.CreateFileStorager(env.compath)
	case "mongo":
		s = storage.CreateMongoStorager()
	}

	var isAlreadyLogin bool = false
	authinfo := readAuthInfo(env)
	var userinfo *model.UserInfo
	if authinfo != "" {
		userinfo = model.CreateUserinfo(authinfo)
		isAlreadyLogin = true
	}

	eng = &Engine{
		App:      initApp(),
		authInfo: authinfo,
		isLogin:  isAlreadyLogin,
		userinfo: userinfo,
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
				err := eng.Save(eng.env.HomePath(), eng.env.Shell())
				if err != nil {
					fmt.Println("Command failed: ", err)
				}
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
				err := eng.List()
				if err != nil {
					fmt.Println("Command failed: ", err)
				}
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
				cmd, err := eng.FetchCommandFromNumber(num)
				if err != nil {
					fmt.Println("Command failed: ", err)
				} else {
					println(cmd.Cmd)
				}
			},
		},
		{
			Name:      "run",
			ShortName: "r",
			Usage:     "Exec command with #number (experiment)",
			Action: func(c *cli.Context) {
				log.Fatal("Not yet implemented")
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

func (e *Engine) Close() {
	e.storager.Close()
	// write auth token
	authFilePath := e.env.compath + "/" + AuthFile
	if e.IsLogin() {
		token := []byte(e.AuthInfo())
		//		username := e.config.User.Mail
		ioutil.WriteFile(authFilePath, token, 0644)
	} else {
		os.Remove(authFilePath)
	}
}

func (e *Engine) List() (err error) {
	if e.storager.IsRequireLogin() == true && e.isLogin == false {
		log.Fatal("You have no valid access token. Please login first.")
	}
	if err = e.storager.List(e.userinfo); err != nil {
		if err == model.ErrSessionExpires {
			e.SetLogout()
		} else if err == model.ErrSessionInvalid {
			e.SetLogout()
		}
	}
	return
}

func (e *Engine) ShowConfig() {
	e.config.ShowConfig()
}

func (e *Engine) FetchCommandFromNumber(num int) (cmd *model.Command, err error) {

	cmd, err = e.storager.FetchCommandFromNumber(e.userinfo, num)
	if err == model.ErrSessionExpires {
		e.SetLogout()
	}
	return
}
