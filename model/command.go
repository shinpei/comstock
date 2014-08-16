package model

import (
	"time"
)

// time format
const (
	RFC3339 = "2014-01-01T15:04:05Z07:00"
)

type Command struct {
	Cmd       string
	Timestamp string
}

func CreateCommand(cmd string) *Command {

	return &Command{Cmd: cmd, Timestamp: time.Now().Format(RFC3339)}

}
