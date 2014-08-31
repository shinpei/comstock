package engine

import (
	"github.com/shinpei/comstock/model"
	"log"
)

func (e *Engine) FetchCommandFromNumber(num int) (cmd *model.Command, err error) {
	if e.storager.IsRequireLogin() == true && e.isLogin == false {
		log.Fatal("You have no valid access token. Please login first.")
	}
	cmd, err = e.storager.FetchCommandFromNumber(e.userinfo, num)
	if err == model.ErrSessionExpires || err == model.ErrSessionInvalid {
		e.SetLogout()
	}
	return
}
