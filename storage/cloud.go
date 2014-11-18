// Cloud storager for heroku
package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	cmodel "github.com/shinpei/comstock/model"
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

func (cs *CloudStorager) Push(user *cmodel.AuthInfo, path string, ns *cmodel.NaiveHistory) (err error) {

	command := "/postHistory"
	objStr, _ := json.Marshal(ns)

	vals := url.Values{"history": {string(objStr)}, "token": {user.Token()}}.Encode()
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
		err = &cmodel.ServerSystemError{} //cmodel.ErrSessionExpires
		// disable login status
	case http.StatusForbidden:
		err = errors.New("Hasn't login, please login first")
	default:
		//	body, _ := ioutil.ReadAll(resp.Body)
	}
	return
}

func (cs *CloudStorager) List(user *cmodel.AuthInfo) (hists []cmodel.NaiveHistory, err error) {

	command := "/list"
	// does it have auto
	vals := url.Values{"token": {user.Token()}}.Encode()
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
		err = &cmodel.SessionInvalidError{}
		return
	case http.StatusInternalServerError:
		err = &cmodel.SessionExpiresError{}
		return
	default:
		fmt.Println("Failed to fetch")
		return
	}
	err = json.Unmarshal(body, &hists)
	return
}

func (cs *CloudStorager) FetchFromNumber(user *cmodel.AuthInfo, index int) (nh *cmodel.NaiveHistory, err error) {

	command := "/fetchCommandFromNumber"
	vals := url.Values{"token": {user.Token()}, "number": {strconv.Itoa(index)}}.Encode()
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
	case http.StatusForbidden:
		err = &cmodel.SessionInvalidError{} //ErrSessionInvalid
		return
	case http.StatusNotFound:
		err = (&cmodel.CommandNotFoundError{}).SetError("No command found for idx=" + strconv.Itoa(index))
		return
	case http.StatusInternalServerError:
		err = &cmodel.ServerSystemError{} //ErrServerSystem
		return
	case http.StatusBadRequest:
		err = (&cmodel.IllegalArgumentError{}).SetError("Invalid argument are given, idx=" + strconv.Itoa(index))
		return
	default:
		err = errors.New("Fetch failed somehow")
		return
	}
	err = json.Unmarshal(body, &nh)
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

func (cs *CloudStorager) CheckSession(user *cmodel.AuthInfo) bool {
	command := "/checkSession"
	vals := url.Values{"token": {user.Token()}}.Encode()
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

func (cs *CloudStorager) RemoveOne(user *cmodel.AuthInfo, index int) (err error) {
	command := "/removeOne"
	vals := url.Values{"token": {user.Token()}, "index": {strconv.Itoa(index)}}.Encode()
	requestURI := cs.StorageHost() + command + "?" + vals
	resp, err := http.Get(requestURI)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusOK:
		// do nothing
	case http.StatusForbidden:
		err = (&cmodel.SessionExpiresError{}).SetError("Session expired, please login again")
	case http.StatusUnauthorized:
		err = (&cmodel.SessionNotFoundError{}).SetError("Session not found") //ErrSessionNotFound
	case http.StatusNotFound:
		err = (&cmodel.CommandNotFoundError{}).SetError("No command found for idx=" + strconv.Itoa(index))
	case http.StatusInternalServerError:
		err = &cmodel.ServerSystemError{} //ErrServerSystem
	default:
		log.Fatal("[SERIOUS] Shouldn't be reache here, please report bug")
	}
	return
}
