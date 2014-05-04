package main

import (
	"io/ioutil"
	"log"
	"os"
)

type Storager interface {
	Open() error
	Close() error
	Push(cmd *Command)

	// getter
	StorageType() string
}

type FileStorager struct {
	//localstorage
	Fp *os.File
}

func (fs *FileStorager) Open() (err error) {

	return
}
func (fs *FileStorager) StorageType() string {
	return "FileStorage"
}
func (fs *FileStorager) Push(cmd *Command) {
	filename := "comstock.txt"

	data, _ := ioutil.ReadFile(filename)

	cmdByte := []byte(cmd.Cmd())
	data = append(data, string("\n")...)
	data = append(data, cmdByte...)
	if err := ioutil.WriteFile(filename, data, 0644); err != nil {
		log.Printf("writestring problem")
		log.Fatal(err)
	}
	return
}

func (fs *FileStorager) Close() (err error) {
	return
}

type RemoteStorager interface{}
