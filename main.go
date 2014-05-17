package main

import (
	"github.com/shinpei/comstock/engine"
	"os"
)

func main() {
	engine = engine.NewEngine()
	engine.Run(os.Args)
}
