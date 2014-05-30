package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shinpei/comstock/model"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const (
	ComstockHost = "http://comstock.herokuapp.com"
	//ComstockHost = "http://localhost:5000"
)

type CloudStorager struct {
}

func (hs *CloudStorager) Open() (err error) {
	return
}

func CreateCloudStorager() (h *CloudStorager) {
	return &CloudStorager{}
}

func (hs *CloudStorager) Push(user *model.UserInfo, path string, cmd *model.Command) (err error) {
	command := "/postCommand"
	vals := url.Values{"cmd": {cmd.Cmd}, "authinfo": {user.AuthInfo()}}.Encode()
	requestURI := ComstockHost + command + "?" + vals
	resp, err := http.Get(requestURI)
	if err != nil {
		log.Fatal("Couldn't reach server", err)
		return
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		// do nothing.
	case 500: // session expires
		err = model.ErrSessionExpires
		// disable login status
	case 403:
		err = errors.New("Hasn't login, please login first")
	default:
		//	body, _ := ioutil.ReadAll(resp.Body)
	}
	return
}

func (hs *CloudStorager) List(user *model.UserInfo) (err error) {

	command := "/list"
	// does it have auto
	vals := url.Values{"authinfo": {user.AuthInfo()}}.Encode()
	requestURI := ComstockHost + command + "?" + vals
	resp, err := http.Get(requestURI)
	if err != nil {
		log.Fatal("Couldn't reach server, ", err)
		return
	}
	defer resp.Body.Close()

	var body []byte
	switch resp.StatusCode {
	case 200:
		body, _ = ioutil.ReadAll(resp.Body)
	case 404:
		err = errors.New("Not found")
		return
	case 403:
		err = model.ErrSessionInvalid
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

func (hs *CloudStorager) FetchCommandFromNumber(user *model.UserInfo, num int) (cmd *model.Command, err error) {
	command := "/fetchCommandFromNumber"
	vals := url.Values{"authinfo": {user.AuthInfo()}, "number": {strconv.Itoa(num)}}.Encode()
	requestURI := ComstockHost + command + "?" + vals
	resp, err := http.Get(requestURI)
	if err != nil {
		log.Fatal("Couldn't reach Host: ", err)
	}
	defer resp.Body.Close()
	var body []byte
	switch resp.StatusCode {
	case 200:
		body, _ = ioutil.ReadAll(resp.Body)
	case 404, 403:
		err = errors.New("Number not found")
		return
	case 500:
		err = model.ErrSessionExpires
		return
	default:
		err = errors.New("Fetch failed somehow")
		return
	}
	var cmds []model.Command
	err = json.Unmarshal(body, &cmds)
	cmd = &cmds[0]
	return
}

func (hs *CloudStorager) StorageType() string {
	return "CloudStorager"
}

func (hs *CloudStorager) Close() (err error) {
	return
}
func (hs *CloudStorager) IsRequireLogin() bool {
	return true
}
