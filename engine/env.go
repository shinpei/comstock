package engine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"runtime"
)

type Env struct {
	Compath  string
	Homepath string
	OS       string
	Arch     string
	Shell    string
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
		Compath:  compath,
		Homepath: homeDir,
		OS:       runtime.GOOS,
		Arch:     runtime.GOARCH,
		Shell:    shell,
	}
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

func createVersionFile(path string) {
	versioninfo := []byte(Version)
	ioutil.WriteFile(path, versioninfo, 0644)
}

func getVersion(path string) string {
	versioninfo, _ := ioutil.ReadFile(path)
	return string(versioninfo)
}

func (e *Env) Dump() string {
	dumpStr, err := json.Marshal(e)
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(dumpStr)
}

func (e *Env) AsMap() map[string]string {
	m := map[string]string{
		"Compath":  e.Compath,
		"Homepath": e.Homepath,
		"Arch":     e.Arch,
		"OS":       e.OS,
		"Shell":    e.Shell,
	}
	return m
}
