package main

import (
	"code.google.com/p/gcfg"
)

type Config struct {
	Local struct {
		Type        string
		StoragePath string
	}
}

const (
	CompathDefault    string = ".comstock"
	ConfigFileDefault string = ".comconfig"
)

func LoadConfig() *Config {
	var cfg Config
	_ = gcfg.ReadFileInto(&cfg, ConfigFileDefault)
	println(cfg.Local.Type)
	return &cfg
}