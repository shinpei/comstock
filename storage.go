package main

import (
	"os"
)

type Storage interface {
	Open() error
	GetReader()
	GetWriter()
}

type LocalStorage interface {
	Storage
	Push(cmd string)
}

type FileStorage struct {
	//localstorage
	Fp *os.File
}

func (fs *FileStorage) Open() (err error) {

	var storageFile string = "comstock.storage"
	fs.Fp, err = os.Open(storageFile)
	return
}
func (fs *FileStorage) Push(cmd string) {
	return
}

type RemoteStorage interface{}
