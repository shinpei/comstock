package comstock

import (
	"code.google.com/p/gcfg"
	"log"
)

type Config struct {
	Local struct {
		Type        string
		StoragePath string
	}
	Remote struct {
		Type        string
		StoragePath string
	}
	Alias struct {
	}
}

const (
	CompathDefault    string = ".comstock"
	ConfigFileDefault string = ".comconfig"
)

func LoadConfig(path string) *Config {
	var cfg Config
	if path == "" {
		path = ConfigFileDefault
	}

	err := gcfg.ReadFileInto(&cfg, path)
	if err != nil {
		log.Fatal(err)
	}
	return &cfg
}
