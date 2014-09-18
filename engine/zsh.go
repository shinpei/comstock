package engine

import (
	"bufio"
	"os"
	"regexp"
)

type ZshHandler struct {
}

func (z *ZshHandler) ReadLastHistory(filename string) (cmd string, err error) {
	var (
		preCmd   string
		storeCmd string
	)

	//format
	// ': xxxxxxxxxx:x;cmd\n'
	println(filename)
	fi, err := os.Open(filename)
	scanner := bufio.NewScanner(fi)

	//TODO: fix below algo
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
	cmd = preCmd[15:]

	return
}

func (z *ZshHandler) ReadEveryHistory(filename string) (cmd string, err error) {
	println(filename)
	fi, err := os.Open(filename)
	scanner := bufio.NewScanner(fi)
	var storeCmd string
	println("reading")
	var validLine = regexp.MustCompile("^:")
	for scanner.Scan() {
		line := scanner.Text()
		idx := validLine.FindIndex([]byte(line))
		if idx != nil {
			storeCmd = line[15:]
			println(storeCmd)
		} else {
			storeCmd += line
		}
	}

	cmd = storeCmd
	return
}
