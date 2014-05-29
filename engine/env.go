package engine

import (
	"fmt"
	"io/ioutil"
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

const (
	ComVersionFile string = "version"
)

func NewEnv() *Env {
	user, _ := user.Current()
	shell := os.Getenv("SHELL")
	homeDir := user.HomeDir
	compath := ""

	if compath = os.Getenv("COMSTOCK_PATH"); compath == "" {
		// if it's empty, load default path, ~/.comstock
		compath = homeDir + "/" + CompathDefault
	}

	if !IsFileExist(compath) {
		// we need to init comstock
		CreateComstockPath(compath)
	}

	// TODO: verify comstock version
	versionPath := compath + "/" + ComVersionFile
	if IsFileExist(versionPath) {
		createVersionFile(versionPath)
	} else {
		version := getVersion(versionPath)
		// versioncheck
		if version != Version {
			// Version mismatch
		}
	}
	return &Env{
		compath:  compath,
		homepath: homeDir,
		os:       runtime.GOOS,
		arch:     runtime.GOARCH,
		shell:    shell,
	}
}

func (e *Env) ComPath() string {
	return e.compath
}

func CreateComstockPath(path string) (err error) {
	err = os.Mkdir(path, 0755)
	if err != nil {
		fmt.Printf("Cannot create compath '%s'\n", path)
	} else {
		fmt.Printf("Create compath as '%s'\n", path)
	}
	return
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

func createVersionFile(path string) {
	versioninfo := []byte(Version)
	ioutil.WriteFile(path, versioninfo, 0644)
}

func getVersion(path string) string {
	versioninfo, _ := ioutil.ReadFile(path)
	return string(versioninfo)
}
