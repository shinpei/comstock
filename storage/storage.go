package storage

import (
	"github.com/shinpei/comstock/model"
)

type Storager interface {
	Open() error
	Close() error
	Push(user *model.AuthInfo, path string, cmd *model.Command) error
	//Push2(user *model.AuthInfo, path string, cmds []model.Command) error
	List(user *model.AuthInfo) (cmds []model.Command, err error)
	FetchCommandFromNumber(user *model.AuthInfo, num int) (cmd *model.Command, err error)
	RemoveOne(user *model.AuthInfo, num int) error
	StorageType() string
	IsRequireLogin() bool
	Status() error
	CheckSession(user *model.AuthInfo) bool
}
