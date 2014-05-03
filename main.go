package main

import (
	"os"
)

func main() {
	comstock := NewComstock()
	comstock.Run(os.Args)
}
