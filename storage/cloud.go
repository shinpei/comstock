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

type CloudStorager struct {
	storageHost string
}

func (cs *CloudStorager) Open() (err error) {
	return
}

func (cs *CloudStorager) StorageHost() string {
	return cs.storageHost
}

func CreateCloudStorager(host string) (h *CloudStorager) {
	return &CloudStorager{
		storageHost: host,
	}
}

func (cs *CloudStorager) Push(user *model.AuthInfo, path string, cmd *model.Command) (err error) {
	command := "/postCommand"
	vals := url.Values{"cmd": {cmd.Cmd}, "authinfo": {user.Token()}}.Encode()
	requestURI := cs.StorageHost() + command + "?" + vals
	resp, err := http.Get(requestURI)
	if err != nil {
		log.Fatal("Couldn't reach server", err)
		return
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusOK:
		// do nothing.
	case http.StatusInternalServerError: // session expires
		err = model.ErrSessionExpires
		// disable login status
	case http.StatusForbidden:
		err = errors.New("Hasn't login, please login first")
	default:
		//	body, _ := ioutil.ReadAll(resp.Body)
	}
	return
}

func (cs *CloudStorager) List(user *model.AuthInfo) (cmds []model.Command, err error) {

	command := "/list"
	// does it have auto
	vals := url.Values{"authinfo": {user.Token()}}.Encode()
	requestURI := cs.StorageHost() + command + "?" + vals
	resp, err := http.Get(requestURI)
	if err != nil {
		log.Fatal("Couldn't reach server, ", err)
		return
	}
	defer resp.Body.Close()

	var body []byte
	switch resp.StatusCode {
	case http.StatusOK:
		body, _ = ioutil.ReadAll(resp.Body)
	case http.StatusNotFound:
		err = errors.New("Not found")
		return
	case http.StatusForbidden:
		err = model.ErrSessionInvalid
		return
	case http.StatusInternalServerError:
		err = model.ErrSessionExpires
		return
	default:
		fmt.Println("Failed to fetch")
		return
	}
	err = json.Unmarshal(body, &cmds)
	return
}

func (cs *CloudStorager) FetchCommandFromNumber(user *model.AuthInfo, num int) (cmd *model.Command, err error) {
	command := "/fetchCommandFromNumber"
	vals := url.Values{"authinfo": {user.Token()}, "number": {strconv.Itoa(num)}}.Encode()
	requestURI := cs.StorageHost() + command + "?" + vals
	resp, err := http.Get(requestURI)
	if err != nil {
		log.Fatal("Couldn't reach Host: ", err)
	}
	defer resp.Body.Close()
	var body []byte
	switch resp.StatusCode {
	case http.StatusOK:
		body, _ = ioutil.ReadAll(resp.Body)
	case http.StatusForbidden, http.StatusNotFound:
		err = errors.New("Number not found")
		return
	case http.StatusInternalServerError:
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

func (cs *CloudStorager) StorageType() string {
	return "CloudStorager"
}

func (cs *CloudStorager) Close() (err error) {
	return
}
func (cs *CloudStorager) IsRequireLogin() bool {
	return true
}

func (cs *CloudStorager) Status() (err error) {
	var m map[string]string = make(map[string]string)
	m["StoragerType"] = cs.StorageType()
	m["StorageURL"] = cs.StorageHost()
	for k, v := range m {
		fmt.Println(k, ":", v)
	}
	return
}

func (cs *CloudStorager) Search() (err error) {
	return
}

func (cs *CloudStorager) CheckSession(user *model.AuthInfo) bool {
	command := "/checkSession"
	vals := url.Values{"authinfo": {user.Token()}}.Encode()
	requestURI := cs.StorageHost() + command + "?" + vals
	resp, err := http.Get(requestURI)
	if err != nil {
		log.Fatal("Couldn't reach server")
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusOK:
		return true
	case http.StatusBadRequest:
		return false
	}
	return false
}

func (cs *CloudStorager) RemoveOne(user *model.AuthInfo, num int) bool {
	//	command := "/removeOne"
	return false
}
