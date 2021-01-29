package asset

import (
	"fmt"

	"github.com/cuckflong/gophish-analytics/pkg/events"
)

type Email struct {
	email  string
	events []events.Event
}

func NewEmail(a string) *Email {
	e := &Email{
		email:  a,
		events: []events.Event{},
	}
	return e
}

func (e *Email) GetEmail() string {
	return e.email
}

func (e *Email) AddEvent(t string, m string, p string) {
	newEvent := events.NewEvent(t, m, p)
	e.events = append(e.events, newEvent)
}

func (e *Email) GetPasswds() []string {
	var passwds []string
	for _, event := range e.events {
		if event.GetMessage() == "Submitted Data" {
			passwds = append(passwds, event.GetPasswd())
		}
	}
	fmt.Println(passwds)
	return passwds
}

func (e *Email) GetLogins() []string {
	var logins []string
	for _, event := range e.events {
		if event.GetMessage() == "Submitted Data" {
			logins = append(logins, event.GetLogin())
		}
	}
	fmt.Println(logins)
	return logins
}

func (e Email) IsClicked() bool {
	for _, event := range e.events {
		if event.GetMessage() == "Clicked Link" {
			return true
		}
	}
	return false
}

func (e Email) IsSubmitted() bool {
	for _, event := range e.events {
		if event.GetMessage() == "Submitted Data" {
			return true
		}
	}
	return false
}

func (e Email) GetClickEvents() []events.Event {
	var list []events.Event
	for _, event := range e.events {
		if event.GetMessage() == "Clicked Link" {
			list = append(list, event)
		}
	}
	return list
}

func (e Email) GetSubmitEvents() []events.Event {
	var list []events.Event
	for _, event := range e.events {
		if event.GetMessage() == "Submitted Data" {
			list = append(list, event)
		}
	}
	return list
}
