package engine

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
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

func NewEnv() *Env {
	user, err := user.Current()
	var homeDir string
	if err != nil {
		// should warn?
		homeDir = os.Getenv("HOME")
	} else {
		homeDir = user.HomeDir
	}
	compath := ""
	shell := getShell()

	if compath = os.Getenv("COMSTOCK_PATH"); compath == "" {
		// if it's empty, load default path, ~/.comstock
		compath = homeDir + "/" + CompathDefault
	}

	if !IsFileExist(compath) {
		// we need to init comstock
		CreateComstockPath(compath)
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
	lc := 0
	for {
		linebuf, _, err := r.ReadLine()
		if err != io.EOF {
			lc++
			if lc == 2 {
				line := string(linebuf)
				return strings.Fields(line)[3]
			}
		} else {
			break
		}
	}

	return "/bin/noshell"
}
