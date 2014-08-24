package engine

import (
	"github.com/shinpei/comstock/model"
	"log"
)

func (e *Engine) Remove(index int) (err error) {
	if e.storager.IsRequireLogin() == true && e.isLogin == false {
		log.Fatal("You have no valid access token. Please login first.")
	}
	if err = e.storager.RemoveOne(e.userinfo, index); err != nil {
		if err == model.ErrSessionExpires {
			e.SetLogout()
		} else if err == model.ErrSessionInvalid {
			e.SetLogout()
		}
		log.Println("Couldn't delete: ", err.Error())
		return
	}
	log.Println("Successfully remove command #", index)
	return
}
