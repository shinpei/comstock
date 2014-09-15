package engine

import (
	"code.google.com/p/gcfg"
	"fmt"
	"github.com/codegangsta/cli"
	"log"
)

var ConfigCommand cli.Command = cli.Command{
	Name:  "config",
	Usage: "Show comstock configuration",
	Action: func(c *cli.Context) {
		eng.Config()
	},
}

type Config struct {
	verboseMode bool
	Local       struct {
		Type string
		URI  string
	}
	Remote struct {
		Type string
		URI  string
	}
	User struct {
		Name string
		Mail string
	}
	path string
}

const (
	CompathDefault    string = ".comstock"
	ConfigFileDefault string = "config"
)

func LoadConfig(path string) *Config {
	var cfg Config

	if path == "" {
		path = ConfigFileDefault
	}
	if IsFileExist(path) == false {
		return &cfg
	}
	err := gcfg.ReadFileInto(&cfg, path)
	if err != nil {
		log.Println("Invalid configuration file")
		log.Fatal(err)
	}

	// set other values
	cfg.path = path
	cfg.verboseMode = false
	// Set defaults
	if cfg.Local.Type == "" {
		cfg.Local.Type = "file"
	}
	return &cfg
}

func DumpConfig(path string) (err error) {
	return
}

func printConfig(key string, val string) {
	if val != "" {
		fmt.Printf("%s=%s\n", key, val)
	}
}
func (c *Config) ShowConfig() {
	printConfig("local.type", c.Local.Type)
	printConfig("local.uri", c.Local.URI)
	printConfig("remote.type", c.Remote.Type)
	printConfig("remote.uri", c.Remote.URI)
	printConfig("user.name", c.User.Name)
	printConfig("user.mail", c.User.Mail)
}

func (c *Config) Path() string {
	return c.path
}
