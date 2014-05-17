package comstock

import (
	"time"
)

// time format
const (
	RFC3339 = "2006-01-02T15:04:05Z07:00"
)

type Command struct {
	cmd       string
	timestamp string
}

func (c *Command) Cmd() string {
	return c.cmd
}

func CreateCommand(cmd string) *Command {

	return &Command{cmd: cmd, timestamp: time.Now().Format(RFC3339)}

}
