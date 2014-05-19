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
	var mail string
	if e.config.User.Mail == "" && e.config.User.Name == "" {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Your registered email? : ") // username is not defineable
		scanner.Scan()
		mail = scanner.Text()
	} else {
		mail = e.config.User.Mail
	}
	fmt.Printf("Password for %s? :", mail)
	password, _ := gopass.GetPass("")
	authInfo := tryLoginWithMail(mail, password)
	if authInfo == "" {
		// TODO: register?
		fmt.Println("Login failed")
		return
	}
	// success, write authinfo
	e.SetLogin()
	e.SetAuthInfo(authInfo)
	fmt.Println("Knock knock ... Success!")
}

func tryLoginWithMail(mail string, password string) string {
	requestURI := ComstockHost + "/loginAs?mail=" + mail + "&password=" + password
	resp, err := http.Get(requestURI)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	//TODO: control over proxy
	body, err := ioutil.ReadAll(resp.Body)
	token := string(body) // access token
	return token
}
