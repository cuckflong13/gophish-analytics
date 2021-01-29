// This part of the code is just pure shit, I got lazy

package events

import (
	"encoding/json"
	"time"
)

type Event struct {
	time      time.Time
	message   string
	rid       string
	address   string
	userAgent string
	login     string
	password  string
}

// Behold the ugliest code ever...
func NewEvent(t string, m string, p string) Event {
	e := Event{}
	e.time, _ = time.Parse(time.RFC3339, t)
	e.message = m

	switch m {
	case "Email Sent":
	case "Clicked Link":
		fallthrough
	case "Email Opened":
		var result map[string]interface{}
		json.Unmarshal([]byte(p), &result)
		e.address = result["browser"].(map[string]interface{})["address"].(string)
		e.userAgent = result["browser"].(map[string]interface{})["user-agent"].(string)
		e.rid = result["payload"].(map[string]interface{})["rid"].([]interface{})[0].(string)
	case "Submitted Data":
		var result map[string]interface{}
		json.Unmarshal([]byte(p), &result)
		e.address = result["browser"].(map[string]interface{})["address"].(string)
		e.userAgent = result["browser"].(map[string]interface{})["user-agent"].(string)
		e.rid = result["payload"].(map[string]interface{})["rid"].([]interface{})[0].(string)
		e.login = result["payload"].(map[string]interface{})["login"].([]interface{})[0].(string)
		e.password = result["payload"].(map[string]interface{})["password"].([]interface{})[0].(string)
	}

	return e
}

func (e Event) GetTime() time.Time {
	return e.time
}

func (e Event) GetMessage() string {
	return e.message
}

func (e Event) GetPasswd() string {
	return e.password
}

func (e Event) GetLogin() string {
	return e.login
}

func (e Event) GetRid() string {
	return e.rid
}

func (e Event) GetAddr() string {
	return e.address
}

func (e Event) GetAgent() string {
	return e.userAgent
}
