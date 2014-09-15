package engine

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

type ShellHandler interface {
	ReadLastHistory(historyfile string) (string, error)
	ReadEveryHistory(historyfile string) (string, error)
}

func FetchShellHandler(e *Engine) (handler ShellHandler) {
	var shellHistoryFilename string
	if strings.Contains(e.env.Shell, "zsh") {
		shellHistoryFilename += "/.zsh_history"
		handler = &ZshHandler{}
	} else if strings.Contains(e.env.Shell, "bash") {
		shellHistoryFilename += "/.bash_history"
		handler = &BashHandler{}
	} else {
		log.Fatal("Couldn't recognize your shell. Your env is ", e.env.Shell)
	}
	return
}

type ZshHandler struct {
}

type BashHandler struct {
}

func tail(filename string, numberLines int) (ret []string, err error) {
	var (
		fi            *os.File
		line          []byte
		hasMoreInLine bool
	)

	if fi, err = os.Open(filename); err != nil {
		log.Fatal(err)
		return
	}
	defer fi.Close()
	freader := bufio.NewReader(fi)
	for {
		if line, hasMoreInLine, err = freader.ReadLine(); err != nil {
			// EOF comes here also
			break
		}
		if len(ret) == numberLines {
			ret = ret[1:len(ret)]
		}
		ret = append(ret, string(line[:]))
		if !hasMoreInLine {
			// do something is required, but don't know yet..
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}

func (z *ZshHandler) ReadLastHistory(filename string) (command string, err error) {
	var (
		preCmd   string
		storeCmd string
	)

	//format
	// ': xxxxxxxxxx:x;cmd\n'
	fi, _ := os.Open(filename)
	scanner := bufio.NewScanner(fi)

	var validLine = regexp.MustCompile("^:")
	for scanner.Scan() {
		line := scanner.Text()
		idx := validLine.FindIndex([]byte(line))
		if idx != nil {
			preCmd = storeCmd
			storeCmd = line
		} else {
			storeCmd += line
		}
	}
	//fmt.Sscanf(preCmd, ": %d:%d;%s", &timestamp, &linenum, &ignore)
	command = preCmd[15:]

	return
}

// Bash stores it's history in its cache. So we cannot fetch it from history file.
func (b *BashHandler) ReadLastHistory(filename string) (command string, err error) {
	var (
		ret []string
	)
	ret, err = tail(filename, 2)
	command = ret[0]
	return
}

func (z *ZshHandler) ReadEveryHistory(filename string) (cmd string, err error) {
	return
}
func (b *BashHandler) ReadEveryHistory(filename string) (cmd string, err error) {
	return
}
