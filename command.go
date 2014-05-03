package main

type Command struct {
	cmd       string
	Timestamp int
}

func (c *Command) Cmd() string {
	return c.cmd
}
