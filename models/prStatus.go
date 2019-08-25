package models

type PRStatus struct {
	Username string
	Title string
	ApprovedStatus string
	PullRequestUrl string
}

type ByStatus []PRStatus

func (a ByStatus) Len() int           { return len(a) }
func (a ByStatus) Less(i, j int) bool { return a[i].ApprovedStatus < a[j].ApprovedStatus }
func (a ByStatus) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }