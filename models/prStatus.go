package models

import (
	"time"
)

type PRStatus struct {
	Username string
	Title string
	Approved bool
	ApprovedStatus string
	PullRequestUrl string
	TimeSinceOpened time.Time
}

type ByStatusAndTime []PRStatus

func (a ByStatusAndTime) Len() int           { return len(a) }
func (a ByStatusAndTime) Less(i, j int) bool {
	if a[i].ApprovedStatus < a[j].ApprovedStatus {
		return false
	}
	if a[i].ApprovedStatus > a[j].ApprovedStatus {
		return true
	}
	return a[j].TimeSinceOpened.Unix() < a[i].TimeSinceOpened.Unix()
}
func (a ByStatusAndTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }