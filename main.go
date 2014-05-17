package main

import (
	"github.com/shinpei/comstock/engine"
	"os"
)

func main() {
	com = comstock.engine.NewComstock()
	com.Run(os.Args)
}
