/*
Comstock, - Store your command to the cloud, https://comstock.herokuapp.com

Basics

Comstock will store your favorite command, such as hard to remember,
less used but convenient command.

*/
package main

import (
	"github.com/shinpei/comstock/engine"
	"os"
)

func main() {
	eng := engine.NewEngine(Version, ComstockAPIServer)
	_ = eng.Run(os.Args)
}
