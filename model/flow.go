package model

type Flow struct {
	Self []byte
	Next []byte
	Hash []byte
}

/*
var NullCommand *Command = &Command{Cmd: "null", Hash: nil}

func NewFlow(cmds []*Command) *Flow {
	var cur *Flow = &Flow{}
	var root *Flow = cur
	for _, cmd := range cmds {
		cur.Self = cmd.Hash
	}
	return root
}
*/
