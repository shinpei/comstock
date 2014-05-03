package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Shell interface {
	ReadLastHistory(historyfile string) (*Command, error)
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

func (this *ZshHandler) ReadLastHistory(filename string) (cmd *Command, err error) {
	var (
		ret       []string
		timestamp int
		linenum   int
	)
	ret, err = tail(filename, 2)
	//format
	// ': xxxxxxxxxx:x;cmd\n'
	var ignore string
	fmt.Sscanf(ret[0], ": %d:%d;%s\n", &timestamp, &linenum, &ignore)
	cmd = &Command{Cmd: ret[0][15:], Timestamp: timestamp} // FIXME: slice number should be more smart
	return
}

func (this *BashHandler) ReadLastHistory(filename string) (cmd *Command, err error) {
	var (
		ret []string
	)
	ret, err = tail(filename, 2)
	cmd = &Command{Cmd: ret[0]}
	return
}
