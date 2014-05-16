package main

import (
	"log"
	"strings"
)

func Save() {
	home := com.env.HomePath()
	var shellHistoryFilename string = home
	var handler Shell = nil
	Save()
	if strings.Contains(com.env.Shell(), "zsh") {
		shellHistoryFilename += "/.zsh_history"
		handler = &ZshHandler{}
	} else if strings.Contains(com.env.Shell(), "bash") {
		shellHistoryFilename += "/.bash_history"
		handler = &BashHandler{}
	}
	cmd, err := handler.ReadLastHistory(shellHistoryFilename)
	if err != nil {
		log.Fatal(err)
	}
	com.Stock(cmd)
}
