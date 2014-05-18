package engine

import (
	"bufio"
	"code.google.com/p/gopass"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	//	ComstockHost string = "http://comstock.herokuapp.com"
	ComstockHost string = "http://localhost:5000"
)

func (e *Engine) Login() {
	// check login
	var username string
	if e.config.User.Mail == "" && e.config.User.Name == "" {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Your registered email or username? : ")
		scanner.Scan()
		username = scanner.Text()
	} else {
		username = e.config.User.Name
		if username == "" {
			username = e.config.User.Mail
		}
	}
	fmt.Printf("Password for %s? :", username)
	password, _ := gopass.GetPass("")
	authInfo := tryLogin(username, password)
	if authInfo != "" {
		// TODO: register?
	}
	// write authinfo
	e.SetLogin()
	e.SetAuthInfo(authInfo)
}

func tryLogin(username string, password string) string {
	requestURI := ComstockHost + "/loginAs?username=" + username + "&password=" + password
	println(requestURI)
	resp, err := http.Get(requestURI)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	//TODO: control over proxy
	body, err := ioutil.ReadAll(resp.Body)
	println(string(body))
	return "access success"
}
