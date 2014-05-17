package main

import (
	"github.com/shinpei/comstock/engine"
	"os"
)

func main() {
	eng := engine.NewEngine()
	eng.Run(os.Args)
}
