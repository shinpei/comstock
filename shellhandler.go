package main
import (
	"os"
	"log"
	"bufio"
	"io"
)

type ShellHandler interface {
	readLastHistory (historyfile string) (string, error)
}

type ZshHandler struct {
}

type BashHandler struct{
	
}

func (this *ZshHandler) readLastHistory (filename string) ( ret string, err error) {
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
			// do something
		}
	}
	if err == io.EOF {
		err = nil;
	}
	ret = string(preline[:]);
	return ;
}

func (this *BashHandler) readLastHistory(filename string) (line string, err error) {
	line = "";
	return;
}
