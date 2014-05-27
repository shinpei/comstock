package engine

import (
	"io/ioutil"
	"os"
	"testing"
)

var configStr string = `
[local]
type=file
uri=http://test.com
[remote]
type=heroku
uri= http://comstock.net
[user]
name=test
mail=test@mail.com
`

func TestLoadConfig(t *testing.T) {
	fs, _ := ioutil.TempFile(os.TempDir(), "")
	defer os.Remove(fs.Name())
	ioutil.WriteFile(fs.Name(), []byte(configStr), 0644)
	c := LoadConfig(fs.Name())
	if c != nil {

	} else {
		t.Error("Cannot create config properly")
	}
	if c.Local.Type != "file" {
		t.Error("couldn't read local type properly: " + c.Local.Type)

	}
	if c.Local.URI != "http://test.com" {
		t.Error("Couldn't read local uri properly: " + c.Local.URI)
	}
	if c.Remote.Type != "heroku" {
		t.Error("couldn't read remote type properly: " + c.Remote.Type)
	}
	if c.Remote.URI != "http://comstock.net" {
		t.Error("Couldn't read local uri properly: " + c.Remote.URI)
	}

}
