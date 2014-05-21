package storage

import (
	"encoding/json"
	"fmt"
	"github.com/shinpei/comstock/model"
	"io/ioutil"
	"net/http"
)

const (
	// ComstockHost = "http://comstock.herokuapp.com"
	ComstockHost = "http://localhost:5000"
)

type HerokuStorager struct {
}

func (hs *HerokuStorager) Open() (err error) {
	return
}

func CreateHerokuStorager() (h *HerokuStorager) {
	return &HerokuStorager{}
}

func (hs *HerokuStorager) Push(user *model.UserInfo, path string, cmd *model.Command) (err error) {
	return
}

func (hs *HerokuStorager) List(user *model.UserInfo) (err error) {
	command := "/list?authinfo=" + user.AuthInfo()
	requestURI := ComstockHost + command
	resp, _ := http.Get(requestURI)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var cmds []model.Command
	err = json.Unmarshal(body, &cmds)
	var idx int = 0
	for _, cmd := range cmds {
		idx++
		fmt.Printf("%d: %s\n", idx, cmd.Cmd)
	}
	return
}

func (hs *HerokuStorager) FetchCommandFromNumber(user *model.UserInfo, num int) (cmd *model.Command) {
	return
}

func (hs *HerokuStorager) StorageType() string {
	return "HerokuStorager"
}

func (hs *HerokuStorager) Close() (err error) {
	return
}
