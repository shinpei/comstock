package storage

import (
	"github.com/shinpei/comstock/model"
)

type Storager interface {
	Open() error
	Close() error
	Push(path string, cmd *model.Command) error
	List() error
	FetchCommandFromNumber(num int) (cmd *model.Command)
	// getter
	StorageType() string
}

type RemoteStorager interface {
	Open() error
	Close() error
	Push(user *model.UserInfo, path string, cmd *model.Command) error
	List(user *model.UserInfo)
	FetchCommandFromNumber(user *model.UserInfo, num int) (cmd *model.Command)
	StorageType() string
}
