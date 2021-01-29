package asset

import (
	"time"
)

func (e Email) TimeSent() time.Time {
	var latestTime time.Time
	for _, event := range e.events {
		if event.GetMessage() == "Email Sent" {
			if latestTime.IsZero() {
				latestTime = event.GetTime()
			} else if latestTime.Before(event.GetTime()) {
				latestTime = event.GetTime()
			}
		}
	}
	if latestTime.IsZero() {
		return time.Time{}
	}
	// fmt.Println(e.GetEmail(), latestTime)
	return latestTime
}

func (e Email) TimeSentString() string {
	return e.TimeSent().Format("2006-01-02 15:04:05")
}

func (e Email) TimeOpened() time.Time {
	var earliestTime time.Time
	for _, event := range e.events {
		if event.GetMessage() == "Email Opened" {
			if earliestTime.IsZero() {
				earliestTime = event.GetTime()
			} else if earliestTime.After(event.GetTime()) {
				earliestTime = event.GetTime()
			}
		}
	}
	if earliestTime.IsZero() {
		return time.Time{}
	}
	// fmt.Println(earliestTime)
	return earliestTime
}

func (e Email) TimeSpentOpen() string {
	if e.TimeOpened().IsZero() {
		return ""
	}
	diff := e.TimeOpened().Sub(e.TimeSent())
	return time.Time{}.Add(diff).Format("15:04:05")
}

func (e Email) TimeClicked() time.Time {
	var earliestTime time.Time
	for _, event := range e.events {
		if event.GetMessage() == "Clicked Link" {
			if earliestTime.IsZero() {
				earliestTime = event.GetTime()
			} else if earliestTime.After(event.GetTime()) {
				earliestTime = event.GetTime()
			}
		}
	}
	if earliestTime.IsZero() {
		return time.Time{}
	}
	// fmt.Println(earliestTime)
	return earliestTime
}

func (e Email) TimeClickedString() string {
	return e.TimeClicked().Format("2006-01-02 15:04:05")
}

func (e Email) TimeSpentClick() string {
	if e.TimeClicked().IsZero() {
		return ""
	}
	diff := e.TimeClicked().Sub(e.TimeSent())
	return time.Time{}.Add(diff).Format("15:04:05")
}

func (e Email) TimeSubmitted() time.Time {
	var earliestTime time.Time
	for _, event := range e.events {
		if event.GetMessage() == "Submitted Data" {
			if earliestTime.IsZero() {
				earliestTime = event.GetTime()
			} else if earliestTime.After(event.GetTime()) {
				earliestTime = event.GetTime()
			}
		}
	}
	if earliestTime.IsZero() {
		return time.Time{}
	}
	// fmt.Println(earliestTime)
	return earliestTime
}

func (e Email) TimeSubmittedString() string {
	return e.TimeSubmitted().Format("2006-01-02 15:04:05")
}

func (e Email) TimeSpentSubmit() string {
	if e.TimeSubmitted().IsZero() {
		return ""
	}
	diff := e.TimeSubmitted().Sub(e.TimeSent())
	return time.Time{}.Add(diff).Format("15:04:05")
}
