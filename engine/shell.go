package engine

import (
	"bufio"
	"io"
	"log"
	"os"
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
