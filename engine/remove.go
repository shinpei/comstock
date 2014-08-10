package engine

import (
	"github.com/shinpei/comstock/model"
)

func (e *Engine) Remove(index int) (err error) {
	if e.isLogin == false {
		err = errors.New("Login requried")
		return
	}
	return
}
