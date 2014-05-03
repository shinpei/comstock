package main

import (
	"bufio"
	"log"
	"os"
)

type Storager interface {
	Open() error
	Close() error
}

type LocalStorager interface {
	Storager
	Push(cmd *Command)
}

type FileStorager struct {
	//localstorage
	Fp *os.File
}

func (fs *FileStorager) Open() (err error) {

	var storageFile string = "comstock.storage"
	fs.Fp, err = os.Open(storageFile)
	if err != nil {
		fs.Fp, err = os.Create(storageFile)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("storage opened:" + storageFile)

	return
}

func (fs *FileStorager) Push(cmd *Command) {

	w := bufio.NewWriter(fs.Fp)
	num, _ := w.WriteString(cmd.Cmd())
	w.Flush()
	log.Println(num)
	fs.Fp.Close()
	return
}

func (fs *FileStorager) Close() (err error) {
	fs.Fp.Close()
	return
}

type RemoteStorager interface{}
