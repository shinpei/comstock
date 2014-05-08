package main

type Config struct {
	compath string
	comroot string
}

const (
	ComstockConfigFilename string = "comstock.yaml"
)

func LoadConfig() *Config {
	return &Config{}
}
