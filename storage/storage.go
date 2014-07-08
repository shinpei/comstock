package storage

import (
	"github.com/shinpei/comstock/model"
)

type Storager interface {
	Open() error
	Close() error
	Push(user *model.UserInfo, path string, cmd *model.Command) error
	List(user *model.UserInfo) (cmds []model.Command, err error)
	FetchCommandFromNumber(user *model.UserInfo, num int) (cmd *model.Command, err error)
	StorageType() string
	IsRequireLogin() bool
	Status() error
	CheckSession(user *model.UserInfo) bool
}
