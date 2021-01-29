package runner

import (
	"log"
	"regexp"

	"github.com/cuckflong/gophish-analytics/pkg/asset"
)

func (r *Runner) addRecord(e string, time string, message string, payload string) {
	if !r.emailExists(e) {
		n := asset.NewEmail(e)
		r.emails = append(r.emails, n)
	}
	email := r.getEmailByAddr(e)
	if email == nil {
		log.Fatalln("email doesn't exist")
	}
	email.AddEvent(time, message, payload)
}

func (r *Runner) emailExists(e string) bool {
	for _, email := range r.emails {
		if email.GetEmail() == e {
			return true
		}
	}
	return false
}

func isValid(e string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}

func (r *Runner) getEmailByAddr(a string) *asset.Email {
	for _, email := range r.emails {
		if email.GetEmail() == a {
			return email
		}
	}
	return nil
}
