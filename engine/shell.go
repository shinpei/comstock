package engine

import (
	"bufio"
	"github.com/shinpei/comstock/model"
	"io"
	"log"
	"os"
	"regexp"
)

type Shell interface {
	ReadLastHistory(historyfile string) (*model.Command, error)
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

func (z *ZshHandler) ReadLastHistory(filename string) (cmd *model.Command, err error) {
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
	cmd = model.CreateCommand(preCmd[15:])
	//cmd = model.CreateCommand(ret[0][15:])
	return
}

func (b *BashHandler) ReadLastHistory(filename string) (cmd *model.Command, err error) {
	var (
		ret []string
	)
	ret, err = tail(filename, 2)
	cmd = model.CreateCommand(ret[0])
	return
}
