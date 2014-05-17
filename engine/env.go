package engine

import (
	"bufio"
	"fmt"
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
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Compath doesn't exists, will you create dir? [Y/n]: ")
		scanner.Scan()
		answer := scanner.Text()
		switch answer {
		case "Y", "y", "":
			break
		case "N", "n":
			//TODO: make local dir as

			return
		default:
			println("Please enter y or n")
			continue
		}
		break
	}
	err = os.Mkdir(path, 0755)
	if err != nil {
		fmt.Printf("Cannot create compath '%s'\n", path)
		//TODO: setup version file
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
