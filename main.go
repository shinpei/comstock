/*
Comstock, - Store your command to the cloud, https://comstock.herokuapp.com

Overview

Comstock will store your favorite command to the cloud instead of copy and pasting
it to the text editor. Here're unique features of comstock

1) Set and get what you write in to your CLI
 $ ls -la
 $ comstock save
 Success to save 'ls -la'

 $ comstock list
 1. ls -la

 $ comstock get 1
 ls -la

 $ comstock run 1
 ls -la

 $ comstock rm 1
 Success to delete 'ls -la'

2) Share your commands (coming soon with v0.3)
 - You'll share your commands with LAN, WAN.

*/

package main

import (
	"github.com/shinpei/comstock/engine"
	"os"
)

func main() {
	cli := engine.CreateComstockCli(Version, ComstockAPIServer)
	_ = cli.Run(os.Args)
}
