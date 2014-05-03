package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	println(os.Args[0])

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	dir, err = filepath.Abs(filepath.Dir("hoge.txt"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	com = NewComstock()
	com.Run(os.Args)
}
