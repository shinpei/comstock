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
	//ComstockHost string = "http://comstock.herokuapp.com"
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
	authInfo, err := tryLoginWithMail(mail, password)
	if err != nil {
		// TODO: register?
		log.Fatal("Login failed:", err)
		return
	}
	// success, write authinfo
	e.SetLogin()
	e.SetAuthInfo(authInfo)
	fmt.Println("Knock knock ... Success!")
}

func tryLoginWithMail(mail string, password string) (token string, err error) {
	requestURI := ComstockHost + "/loginAs?mail=" + mail + "&password=" + password
	var resp *http.Response
	resp, err = http.Get(requestURI)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	//TODO: control over proxy
	switch resp.StatusCode {
	case 200:
		body, _ := ioutil.ReadAll(resp.Body)
		token = string(body) // access token
	case 409:
		err = errors.New("Already loggedin")
		token = ""
	case 404, 403:
		err = errors.New("Wrong username or password")
		token = ""
	default:
		err = errors.New("[Login]Invalid response")
		token = ""
		break
	}
	return

}
