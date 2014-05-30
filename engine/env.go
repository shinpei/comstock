package engine

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strconv"
	"strings"
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
	//	shell := os.Getenv("SHELL")
	shell := getShell()
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

// get executing shell from ppid
func getShell() string {

	ppid := strconv.Itoa(os.Getppid())
	cmd := exec.Command("ps", "-p", ppid)
	pipe, err := cmd.StdoutPipe()
	defer pipe.Close()
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(pipe)
	for {
		linebuf, _, err := r.ReadLine()
		if err != io.EOF {
			line := string(linebuf)
			if idx := strings.Index(line, "/bin"); idx != -1 {
				return string(line[idx:])
			}
		} else {
			break
		}

	}

	return "/bin/noshell"
}
