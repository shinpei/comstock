package engine

import (
	"bufio"
	"code.google.com/p/gopass"
	"errors"
	"fmt"
	"github.com/shinpei/comstock/model"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func (e *Engine) Login(loginServer string) {
	// check login
	// TODO: does storager requires login?

	var mail string
	var registeredNewMail bool
	if e.config != nil && e.config.User.Mail != "" {
		mail = e.config.User.Mail
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Your registered mail? : ") // username is not defineable
		scanner.Scan()
		mail = scanner.Text()
		registeredNewMail = true
	}
	fmt.Printf("Password for %s?:", mail)
	password, _ := gopass.GetPass("")
	token, err := tryLoginWithMail(loginServer, mail, password)
	if err != nil {
		// TODO: register?
		log.Println("Login failed:", err)
	} else {
		// success, write token
		if registeredNewMail == true {
			e.config.User.Mail = mail
			println("New email is registered, you can use config for reserving it")
		}
		e.SetLogin()
		e.SetAuthInfo(token)
	}
}

// this is version dependent.
func tryLoginWithMail(loginServer string, mail string, password string) (token string, err error) {
	requestURI := loginServer + "/loginAs?mail=" + mail + "&password=" + password
	var resp *http.Response
	resp, err = http.Get(requestURI)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	//TODO: control over proxy
	switch resp.StatusCode {
	case http.StatusOK:
		body, _ := ioutil.ReadAll(resp.Body)
		token = string(body) // access token
		// check token has contents
		if token == "" {
			err = errors.New("Server error: fetched token is empty")
		} else {
			fmt.Println("Authentication success.")
		}
	case http.StatusConflict:
		body, _ := ioutil.ReadAll(resp.Body)
		token = string(body)
		fmt.Println("Already logined.")
	case http.StatusNotFound, http.StatusForbidden:
		err = model.ErrAuthenticationFailed
		token = ""
	default:
		err = errors.New("Invalid response.")
		token = ""
		break
	}
	return
}
