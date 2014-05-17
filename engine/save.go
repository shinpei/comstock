package engine

import (
	"log"
	"strings"
)

func (e *Engine) Save(home string, shell string) {
	var shellHistoryFilename string = home
	var handler Shell = nil
	if strings.Contains(shell, "zsh") {
		shellHistoryFilename += "/.zsh_history"
		handler = &ZshHandler{}
	} else if strings.Contains(shell, "bash") {
		shellHistoryFilename += "/.bash_history"
		handler = &BashHandler{}
	}
	cmd, err := handler.ReadLastHistory(shellHistoryFilename)
	if err != nil {
		log.Fatal(err)
	}
	eng.Stock(cmd)

}