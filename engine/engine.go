package engine

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/shinpei/comstock/model"
	"github.com/shinpei/comstock/storage"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	AppName        string = "comstock"
	AuthFile       string = "authinfo"
	ComVersionFile string = "version"
	SPLITTER       string = "#"
)

// this is TODO.
// How can we pass thunk to cli?
var (
	eng *Engine
)

type Engine struct {
	App       *cli.App
	storager  storage.Storager // storage
	userinfo  *model.AuthInfo
	isLogin   bool
	authInfo  string
	config    *Config
	env       *Env
	apiServer string
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

func NewEngine(version string, apiServer string) *Engine {

	env := NewEnv()
	var config *Config
	configPath := env.Compath + "/" + ConfigFileDefault
	var s storage.Storager
	s = storage.CreateCloudStorager(apiServer)
	config = LoadConfig(configPath)
	switch config.Local.Type {
	case "file":
		s = storage.CreateFileStorager(env.Compath)
	case "mongo":
		s = storage.CreateMongoStorager()
	}

	var isAlreadyLogin bool = false
	authinfo, mail, _ := readAuthInfo(env)
	var userinfo *model.AuthInfo
	if authinfo != "" {
		userinfo = model.CreateUserinfo(authinfo, mail)
		//TODO: maybe we shouldn't check session every time, e.g., --help given
		isAlreadyLogin = s.CheckSession(userinfo)
		config.User.Mail = mail // apply current status
	}

	// TODO: verify comstock version
	versionPath := env.Compath + "/" + ComVersionFile
	if !IsFileExist(versionPath) {
		createVersionFile(versionPath, version)
	} else {
		versionRead := getVersion(versionPath)
		// versioncheck
		if versionRead != version {
			// Version mismatch
			// TODO: warn for environment, double installed?
			//log.Printf("version seems changed, please remove")
		}
	}

	eng = &Engine{
		App:       initApp(version),
		authInfo:  authinfo,
		isLogin:   isAlreadyLogin,
		userinfo:  userinfo,
		storager:  s,
		env:       env,
		config:    config,
		apiServer: apiServer,
	}
	return eng
}
func createVersionFile(path string, version string) {
	versioninfo := []byte(version)
	ioutil.WriteFile(path, versioninfo, 0644)
}

func getVersion(path string) string {
	versioninfo, _ := ioutil.ReadFile(path)
	return string(versioninfo)
}

func readAuthInfo(env *Env) (authinfo string, mail string, err error) {
	authFilePath := env.Compath + "/" + AuthFile
	fi, err := os.Open(authFilePath)
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(fi)
	var lc int = 0
	for scanner.Scan() {
		lc++
		if 1 < lc {
			// error
			log.Fatal("Invalid login info")
		}
		auths := strings.Split(scanner.Text(), SPLITTER)
		// for safety.
		if 1 < len(auths) {
			mail = auths[0]
			authinfo = auths[1]
		} else {
			fmt.Println("comstock version is too old, consider upgrading it.")
			authinfo = auths[0]
		}

	}
	return
}

func (e *Engine) flushAuthInfoOrRemove() {
	// write auth token
	authFilePath := e.env.Compath + "/" + AuthFile
	if e.IsLogin() {
		buffer := bytes.NewBufferString("")
		buffer.WriteString(e.config.User.Mail)
		buffer.WriteString(SPLITTER)
		buffer.WriteString(e.AuthInfo())
		ioutil.WriteFile(authFilePath, []byte(buffer.String()), 0644)
	} else {
		os.Remove(authFilePath)
	}

}
func initApp(version string) *cli.App {
	app := cli.NewApp()
	app.Version = version
	app.Name = AppName
	app.Usage = "save your command to the cloud"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "shell, s", Value: "", Usage: "specify flag"},
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Printf("Command '%v' is unknown.\nRun 'comstock --help' to get help.\n", command)
	}
	app.Commands = []cli.Command{
		ActionCommand,
		StatusCommand,
		ListCommand,
		FetchCommand,
		{
			Name:  "alias",
			Usage: "Make alias for specific command, specified as #number",
			Action: func(c *cli.Context) {
				log.Fatal("Not yet implemented")
			},
		},
		RemoveCommand,
		{
			Name:  "run",
			Usage: "Exec command with #number",
			Action: func(c *cli.Context) {
				log.Fatal("Please execute 'run' from wrapper script")
			},
		},
		LoginCommand,
		ConfigCommand,
		OpenCommand,
		LogoutCommand,
		//		ImportCommand,
		//		RawCommand,
	}

	return app
}

func (e *Engine) Run(args []string) error {
	// initiation

	err := e.App.Run(args)
	e.Close()
	return err
}

func (e *Engine) Close() {
	e.storager.Close()
	e.flushAuthInfoOrRemove()
}

func (e *Engine) Config() {
	e.config.ShowConfig()
}
