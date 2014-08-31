package main

import (
	"github.com/shinpei/comstock/engine"
	"os"
)

func main() {
	eng := engine.NewEngine(Version)
	_ = eng.Run(os.Args)

}
