package engine

import (
	"bufio"
	"os"
	"regexp"
)

type ZshHandler struct {
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

func (z *ZshHandler) ReadEveryHistory(filename string) (cmd string, err error) {

	return
}
