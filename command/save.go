package command

import (
	"github.com/shinpei/comstock/engine"
	"log"
	"strings"
)

func Save(com *comstock.Comstock, home string, shell string) {
	var shellHistoryFilename string = home
	var handler comstock.Shell = nil
	Save()
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
	//	com.Stock(cmd)
}
