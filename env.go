package main

import (
	"os"
	"os/user"
	"runtime"
)

type Env struct {
	compath  string
	homepath string
	os       string
	arch     string
	shell    string
}

func CreateEnv() *Env {
	user, _ := user.Current()
	shell := os.Getenv("SHELL")
	return &Env{
		compath:  CompathDefault,
		homepath: user.HomeDir,
		os:       runtime.GOOS,
		arch:     runtime.GOARCH,
		shell:    shell,
	}
}

func (e *Env) ComPath() string {
	return e.compath
}

func (e *Env) HomePath() string {
	return e.homepath
}

func (e *Env) Arch() string {
	return e.arch
}

func (e *Env) OS() string {
	return e.os
}

func (e *Env) Shell() string {
	return e.shell
}
