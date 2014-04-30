package main
import (
	"os"
	"log"
	"bufio"
	"io"
	"fmt"
)

type ShellHandler interface {
	ReadLastHistory (historyfile string) (string, error)
}

type ZshHandler struct {
}

type BashHandler struct{
	
}

// TODO: support -l option
func tail (filename string) (ret string, err error) {
	var (
		fi *os.File
		line []byte
		preline []byte
		hasMoreInLine bool
	)
	if fi, err = os.Open(filename); err != nil {
		log.Fatal(err);
		return;
	}
	defer fi.Close();
	freader := bufio.NewReader(fi);
	for {
		if line, hasMoreInLine, err = freader.ReadLine(); err != nil {
			// EOF comes here also
			break;
		}
		preline = line;
		if !hasMoreInLine { 
			// do something is required, but don't know yet..
		}
	}
	if err == io.EOF {
		err = nil;
	}
	ret = string(preline[:]);
	return ;
}

func (this *ZshHandler) ReadLastHistory (filename string) ( line string, err error) {
	var (
		ret string
		timestamp int
		linenum int
	)
	ret, err = tail(filename);
	fmt.Sscanf(ret, ": %d:%d;%s\n", &timestamp, &linenum, &line);
	return;
}

func (this *BashHandler) ReadLastHistory(filename string) (line string, err error) {
	line, err = tail(filename);
	return;
}
