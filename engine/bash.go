package engine

type BashHandler struct {
}

// Bash stores it's history in its cache. So we cannot fetch it from history file.
func (b *BashHandler) ReadLastHistory(filename string) (command string, err error) {
	var (
		ret []string
	)
	ret, err = tail(filename, 2)
	command = ret[0]
	return
}

func (b *BashHandler) ReadEveryHistory(filename string) (cmd string, err error) {

	return
}
