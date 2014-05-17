package main

import (
	"github.com/shinpei/comstock"
	"os"
)

func main() {
	com = comstock.NewComstock()
	com.Run(os.Args)
}
