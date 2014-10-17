package model

import (
	"github.com/shinpei/comstock/test"
	"test"
)

func TestCreateCommand(t *testing.T) {
	c := CreateCommand("ls -la")
	AssertEqual(t, "ls -la", c.Cmd)
}

func TestCreateUserinfo(t *testing.T) {
	ui := CreateUserinfo("fakeToken", "bob@mail.com")
	AssertEqual(t, "fakeToken", ui.Token())
	AssertEqual(t, "bob@mail.com", ui.Mail())
}
