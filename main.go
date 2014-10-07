package main

import (
	"github.com/shinpei/comstock/engine"
	"os"
)

//ComstockAPIServer string = "http://localhost:5000"

func main() {
	eng := engine.NewEngine(Version, ComstockAPIServer)
	_ = eng.Run(os.Args)
}
