package model

import (
	"time"
)

type NaiveHistory struct {
	Date        time.Time
	Description string
	Cmds        []string
	Shell       string
}

func CreateNaiveHistory(cmds []string, desc string) *NaiveHistory {
	return &NaiveHistory{Date: time.Now(), Cmds: cmds, Description: desc}
}
func (c *NaiveHistory) SetShell(shell string) {
	c.Shell = shell
}

type Command struct {
	Cmd       string
	Timestamp string
	Hash      []byte
	Shell     string
}

func _CreateCommand(cmd string) *Command {
	return &Command{Cmd: cmd, Timestamp: time.Now().Format(time.RFC3339)}
}

func (c *Command) SetShell(shell string) {
	// verify?
	c.Shell = shell
}
