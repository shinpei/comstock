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

type RemoteStorager interface{}
