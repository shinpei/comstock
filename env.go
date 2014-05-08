package main

import (
	"os"
	"os/user"
	"runtime"
)

type Env struct {
	homepath string
	os       string
	arch     string
	shell    string
}

func CreateEnv() *Env {
	user, _ := user.Current()
	shell := os.Getenv("SHELL")
	return &Env{homepath: user.HomeDir, os: runtime.GOOS, arch: runtime.GOARCH, shell: shell}
}

func (e *Env) Shell() string {
	return e.shell
}

func (e *Env) HomePath() string {
	return e.homepath
}
