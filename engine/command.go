package engine

// TODO: in future, we make command pluggable
type Command interface {
	Exec() error
	// getters
	ShortName() string
}

func NewCommand() *Command {
	return nil
}
