package engine

import (
	"errors"
	"fmt"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var RawCommand cli.Command = cli.Command{
	Name:   "raw",
	Usage:  "Tap raw command",
	Action: RawAction,
}

// 'raw' is an action which sends raw commands to the server
func RawAction(c *cli.Context) {
	first := c.Args().First()
	if first == "" {
		fmt.Println("'raw' requires one argument")
		return
	}
	err := eng.Raw(eng.apiServer, first)
	if err != nil {
		fmt.Println("Failed: " + err.Error())
	}
}

func (e *Engine) Raw(host string, cmd string) (err error) {

	if e.isLogin == false {
		err = errors.New("Login required")
		return
	}
	if strings.HasPrefix(cmd, "/") == false {
		cmd = "/" + cmd
	}
	vals := url.Values{"authinfo": {e.userinfo.Token()}}.Encode()
	requestURI := host + cmd + "?" + vals
	resp, err := http.Get(requestURI)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusOK:
		// do nothing
	case http.StatusNotFound:
		err = errors.New("Command '" + cmd + "' seems not found in server-side")

	default:
		body, _ := ioutil.ReadAll(resp.Body)
		log.Fatal("Failed: " + string(body))
	}
	return
}
