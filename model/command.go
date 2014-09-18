package model

import (
	"time"
)

type Command struct {
	Cmd       string
	Timestamp string
	Hash      []byte
}

func CreateCommand(cmd string) *Command {
	return &Command{Cmd: cmd, Timestamp: time.Now().Format(time.RFC3339)}
}
