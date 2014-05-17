package comstock

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	storagerFile string = "comstock.txt"
)

type Storager interface {
	Open() error
	Close() error
	Push(cmd *Command) error
	Pull() error
	List() error
	// getter
	StorageType() string
}

/* FileStorager is a storager which store commands for the local file.
 * file to use can be configurable, and should be devidable
 */

type FileStorager struct {
	//localstorage
	Fp *os.File
}

// Common interface for storager,
// Although FileStorager doesn't have one.
func (fs *FileStorager) Open() (err error) {

	return
}

// Return static string which represents storager type
func (fs *FileStorager) StorageType() string {
	return "FileStorage"
}

// Store the command
func (fs *FileStorager) Push(cmd *Command) (err error) {

	data, _ := ioutil.ReadFile(storagerFile)
	cmdByte := []byte(cmd.Cmd())
	cmdByte = append(cmdByte, string("\n")...)
	data = append(data, cmdByte...)
	err = ioutil.WriteFile(storagerFile, data, 0644)
	return
}

// Close the storge connection
func (fs *FileStorager) Close() (err error) {
	return
}

// List all commands
func (fs *FileStorager) List() (err error) {
	var fi *os.File
	fi, err = os.Open(storagerFile)
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

type RemoteStorager interface{}
