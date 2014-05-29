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
	//LoginServer string = "http://comstock.herokuapp.com"
	LoginServer string = "http://localhost:5000"
)

func (e *Engine) Login() {
	// check login
	var mail string
	if e.config.User.Mail == "" {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Your registered email? : ") // username is not defineable
		scanner.Scan()
		mail = scanner.Text()
	} else {
		mail = e.config.User.Mail
	}
	fmt.Printf("Password for %s?:", mail)
	password, _ := gopass.GetPass("")
	authInfo, err := tryLoginWithMail(mail, password)
	if err != nil {
		// TODO: register?
		log.Println("Login failed:", err)
	} else {
		// success, write authinfo
		e.SetLogin()
		e.SetAuthInfo(authInfo)
	}
}

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
