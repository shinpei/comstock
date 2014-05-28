package storage

import (
	"encoding/json"
	"fmt"
	"github.com/shinpei/comstock/model"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	//ComstockHost = "http://comstock.herokuapp.com"
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
	command := "/postCommand"
	vals := url.Values{"cmd": {cmd.Cmd}, "authinfo": {user.AuthInfo()}}.Encode()
	requestURI := ComstockHost + command + "?" + vals
	resp, err := http.Get(requestURI)
	if err != nil {
		log.Fatal("error")
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
	case 500: // session expires
		err = model.ErrSessionExpires
		// disable login status
	default:
		log.Fatal("Something went wrong")
		//	body, _ := ioutil.ReadAll(resp.Body)
	}
	return
}

func (hs *HerokuStorager) List(user *model.UserInfo) (err error) {
	command := "/list"
	vals := url.Values{"authinfo": {user.AuthInfo()}}.Encode()
	requestURI := ComstockHost + command + "?" + vals
	resp, err := http.Get(requestURI)
	if err != nil {
		log.Fatal("Couldn't reach server, ", err)
	}
	defer resp.Body.Close()
	var body []byte
	switch resp.StatusCode {
	case 200:
		body, _ = ioutil.ReadAll(resp.Body)
	case 404:
		log.Fatal("Failed to fetch")
		return
	case 403:
		log.Fatal("Login required")
		return
	case 500:
		err = model.ErrSessionExpires
		return
	default:
		fmt.Println("Failed to fetch")
		return
	}
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
	command := "/fetchFromNumber"
	vals := url.Values{"authinfo": {user.AuthInfo()}}.Encode()
	requestURI := ComstockHost + command + "?" + vals
	resp, err := http.Get(requestURI)
	if err != nil {
		log.Fatal("Couldn't reach server, ", err)
	}
	defer resp.Body.Close()
	var body []byte
	switch resp.StatusCode {
	case 200:
		body, _ = ioutil.ReadAll(resp.Body)
	case 404, 403:
		fmt.Println("Number not found")
		return
	default:
		fmt.Println("Failed to fetch")
		return
	}
	var cmds []model.Command
	err = json.Unmarshal(body, &cmds)
	var idx int = 0
	for _, cmd := range cmds {
		idx++
		fmt.Printf("%d: %s\n", idx, cmd.Cmd)
	}
	return
}

func (hs *HerokuStorager) StorageType() string {
	return "HerokuStorager"
}

func (hs *HerokuStorager) Close() (err error) {
	return
}
