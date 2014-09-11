package model

import (
	"crypto/md5"
	"time"
)

type Command struct {
	Cmd       string
	Timestamp string
	Hash      []byte
}

func CreateCommand(cmd string) *Command {
	h := md5.New()
	return &Command{Cmd: cmd, Timestamp: time.Now().Format(time.RFC3339), Hash: h.Sum([]byte(cmd))}
}
