package comstock

import (
	"bufio"
	"code.google.com/p/gopass"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/shinpei/comstock/command"
	"log"
	"os"
)

const (
	Version string = "0.0.1"
	AppName string = "comstock"
)

var (
	com *Comstock
)

type Comstock struct {
	App      *cli.App
	storager Storager // storage
	logined  bool
	config   *Config
	env      *Env
}

func (c *Comstock) Logined() bool {
	return c.logined
}

func NewComstock() *Comstock {
	f := &FileStorager{}
	f.Open()
	env := CreateEnv()
	var config *Config
	configPath := env.compath + "/" + ConfigFileDefault
	if IsFileExist(configPath) {
		config = LoadConfig(configPath)
		fmt.Println("Config loaded")
	}
	return &Comstock{
		App:      initApp(),
		storager: &FileStorager{},
		logined:  false,
		env:      env,
		config:   config,
	}
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
				command.Save(com, com.env.HomePath(), com.env.Shell())
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

				com.List()
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
		println("We're preparing this func...")
		return "access success"
	}
}

func (c *Comstock) ShowConfig() {
	println("Showing config")
}
