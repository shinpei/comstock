package engine

import (
	"bufio"
	"code.google.com/p/gopass"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	LoginServer string = "https://comstock.herokuapp.com"
	//LoginServer string = "http://localhost:5000"
)

func (e *Engine) Login() {
	// check login
	var mail string
	var registeredNewEmail bool
	if e.config != nil && e.config.User.Mail != "" {
		mail = e.config.User.Mail
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Your registered email? : ") // username is not defineable
		scanner.Scan()
		mail = scanner.Text()
		registeredNewEmail = true
	}
	fmt.Printf("Password for %s?:", mail)
	password, _ := gopass.GetPass("")
	token, err := tryLoginWithMail(mail, password)
	if err != nil {
		// TODO: register?
		log.Println("Login failed:", err)
	} else {
		// success, write token
		if registeredNewEmail == true {
			e.config.User.Mail = mail
			println("New email is registered, you can use config for reserving it")
		}
		e.SetLogin()
		e.SetAuthInfo(token)
	}
}

// this is version dependent.
func tryLoginWithMail(mail string, password string) (token string, err error) {
	requestURI := LoginServer + "/loginAs?mail=" + mail + "&password=" + password
	var resp *http.Response
	resp, err = http.Get(requestURI)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	//TODO: control over proxy
	switch resp.StatusCode {
	case 200:
		body, _ := ioutil.ReadAll(resp.Body)
		token = string(body) // access token
		println("Authentification success.")
	case 409:
		body, _ := ioutil.ReadAll(resp.Body)
		token = string(body)
		println("Already logined.")
	case 404, 403:
		err = errors.New("Authentification failed.")
		token = ""
	default:
		err = errors.New("Invalid response.")
		token = ""
		break
	}
	return
}
