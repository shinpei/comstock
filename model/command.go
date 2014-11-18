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
func (nh *NaiveHistory) SetShell(shell string) {
	nh.Shell = shell
}

func (nh *NaiveHistory) Command() string {
	l := len(nh.Cmds)
	symbol := ""
	if l == 1 {
		symbol = nh.Cmds[0]
	} else if l > 1 {
		for _, s := range nh.Cmds {
			symbol += s + " => "
		}
		symbol = symbol[:len(symbol)-4]
	}
	return symbol
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
