package engine

import (
	"bufio"
	"code.google.com/p/gopass"
	"errors"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/shinpei/comstock/model"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var LoginCommand cli.Command = cli.Command{
	Name:  "login",
	Usage: "Login to the cloud",
	Flags: []cli.Flag{
		cli.StringFlag{Name: "username, l", Usage: "username (email address)"},
		cli.StringFlag{Name: "password, p", Usage: "password"},
	},
	Action: LoginAction,
}

func LoginAction(c *cli.Context) {
	if eng.IsLogin() {
		fmt.Printf("Already login as %s\n", eng.userinfo.Mail())
		return
	}
	username := c.String("username")
	password := c.String("password")
	eng.Login(eng.apiServer, username, password)

}

func (e *Engine) Login(loginServer string, u string, p string) {

	// TODO: integrate with isrequireLoginordie
	if e.storager.IsRequireLogin() == false {
		fmt.Println("Your storager doesn't require login")
		return
	}
	if e.userinfo != nil {
		// possibly already login
		e.isLoginPolled = true
		e.isLogin = e.storager.CheckSession(e.userinfo)
		if e.isLogin == true {
			fmt.Println("You already logged in. 'comstock logout' for force logout")
			return
		}
	}

	var mail string
	var registeredNewMail bool
	if u != "" {
		mail = u
	} else if e.config != nil && e.config.User.Mail != "" {
		mail = e.config.User.Mail
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Your registered mail? : ") // username is not defineable
		scanner.Scan()
		mail = scanner.Text()
		registeredNewMail = true
	}

	var password string
	if p != "" {
		password = p
	} else {
		fmt.Printf("Password for %s?:", mail)
		password, _ = gopass.GetPass("")
	}
	e.isLoginPolled = true
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
		err = &model.AuthenticationFailedError{} //ErrAuthenticationFailed
		token = ""
	default:
		err = errors.New("Invalid response.")
		token = ""
		break
	}
	return
}
