package model

import (
	"time"
)

type Command struct {
	Cmd       string
	Timestamp time.Time
	Hash      []byte
	Shell     string
}

func CreateCommand(cmd string) *Command {
	return &Command{Cmd: cmd, Timestamp: time.Now()}
}

func (c *Command) SetShell(shell string) {
	// verify?
	c.Shell = shell
}
