// Check environment for Comstock
package engine

import (
	"bufio"
	"encoding/json"
	"errors"
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

func CreateEnv() *Env {
	user, err := user.Current()
	var homeDir string
	if err != nil {
		// should warn?
		log.Fatal("Couldn't fetch user")
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

func getShellProcessName(pid int) (procName string, err error) {
	cmd := exec.Command("ps", "-p", strconv.Itoa(pid))
	pipe, err := cmd.StdoutPipe()
	defer pipe.Close()
	err = cmd.Start()
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(pipe)
	for {
		linebuf, _, err := r.ReadLine()
		if err != io.EOF {
			line := string(linebuf)
			fields := strings.Fields(line)
			if 4 < len(fields) {
				procName = fields[4]
			} else {
				procName = fields[3]
			}
		} else {
			break
		}
	}
	if procName == "" {
		err = errors.New("Command Not fonud for pid=" + strconv.Itoa(pid))
	}
	return
}

func getPPID(target int) (ppid int, err error) {
	cmd := exec.Command("ps", "xao", "pid,ppid", strconv.Itoa(target))
	pipe, err := cmd.StdoutPipe()
	defer pipe.Close()
	err = cmd.Start()
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(pipe)
	for {
		linebuf, _, err := r.ReadLine()
		if err != io.EOF {
			line := string(linebuf)
			fields := strings.Fields(line)
			pid, _ := strconv.Atoi(fields[0])
			if pid == target {
				ppid, _ = strconv.Atoi(fields[1])
				break
			}
		} else {
			break
		}
	}
	if ppid == 0 {
		err = errors.New("Command Not fonud for pid=" + strconv.Itoa(target))
	}
	return
}

// Get parent shell from parent process.
// This is critical for comstock
func getShell() (shell string) {

	shell = "unknown"
	ppid := (os.Getppid())
	shell, err := getShellProcessName(ppid)
	if err != nil {
		// Couldn't get parent shell
		log.Fatal("Coulnd't get parent shell")
		return
	}
	if strings.HasSuffix(shell, "comstock") {
		// this case, it's a wrapper
		ppid, err = getPPID(ppid)
		if err != nil {
			log.Fatal(err)
		}
		shell, err = getShellProcessName(ppid)
	} else if strings.HasSuffix(shell, "sh") && !strings.HasSuffix(shell, "bash") {
		// this case, it's also a wrapper
		ppid, err = getPPID(ppid)
		if err != nil {
			log.Fatal(err)
		}
		shell, err = getShellProcessName(ppid)
	}

	if err != nil {
		log.Fatal(err)
	}
	return
}
