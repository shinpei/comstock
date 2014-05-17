package storage

import (
	"bufio"
	"fmt"
	"github.com/shinpei/comstock/model"
	"io/ioutil"
	"log"
	"os"
)

const (
	storagerFile string = "comstock.txt"
)

type Storager interface {
	Open() error
	Close() error
	Push(path string, cmd *model.Command) error
	Pull() error
	List() error
	FetchCommandFromNumber(num int) (cmd *model.Command)
	// getter
	StorageType() string
}

/* FileStorager is a storager which store commands for the local file.
 * file to use can be configurable, and should be devidable
 */

type FileStorager struct {
	//localstorage
	Fp       *os.File
	filepath string
}

// Common interface for storager,
// Although FileStorager doesn't have one.
func (fs *FileStorager) Open() (err error) {

	return
}

func CreateFileStorager(basepath string) *FileStorager {

	return &FileStorager{
		filepath: basepath + "/" + storagerFile,
	}
}

// Return static string which represents storager type
func (fs *FileStorager) StorageType() string {
	return "FileStorage"
}

// Store the command
func (fs *FileStorager) Push(path string, cmd *model.Command) (err error) {

	data, _ := ioutil.ReadFile(fs.filepath)
	cmdByte := []byte(cmd.Cmd())
	cmdByte = append(cmdByte, string("\n")...)
	data = append(data, cmdByte...)
	err = ioutil.WriteFile(fs.filepath, data, 0644)
	return
}

// Close the storge connection
func (fs *FileStorager) Close() (err error) {
	return
}

// List all commands
func (fs *FileStorager) List() (err error) {
	var fi *os.File
	fi, err = os.Open(fs.filepath)
	scanner := bufio.NewScanner(fi)
	var idx int = 0
	for scanner.Scan() {
		idx++
		fmt.Printf("%d: %s\n", idx, scanner.Text())
	}
	return
}

func (fs *FileStorager) Pull() (err error) {
	return
}

func (fs *FileStorager) FetchCommandFromNumber(num int) (cmd *model.Command) {
	var fi *os.File
	// TODO
	fi, _ = os.Open(fs.filepath)
	scanner := bufio.NewScanner(fi)
	var idx int = 0
	for scanner.Scan() {
		idx++
		if idx == num {
			cmd = model.CreateCommand(scanner.Text())
			return
		}
	}
	log.Fatal("Not exist for your specified number=" + string(num))
	return
}

type RemoteStorager interface{}
