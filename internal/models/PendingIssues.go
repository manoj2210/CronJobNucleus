package models

import (
	"fmt"
	"github.com/google/go-github/github"
)

type PendingIssue struct {
	Title string
	Assignees []*github.User
	TimeRemaining string
	IssueUrl string
}

func NewPendingIssue(title  string,assignees []*github.User,issueUrl  string) PendingIssue{
	return PendingIssue{
		Title:         title,
		Assignees:     assignees,
		TimeRemaining: "",
		IssueUrl: issueUrl,
	}
}

func (pi *PendingIssue) SetTimeRemaining(time  string){
	pi.TimeRemaining=time
}

func (pi *PendingIssue) Print() string{
	return fmt.Sprintf(
		"Title: %v \nAssignees: %v\nRemaingTime: %v\n,Url: %v\n",
		pi.Title,pi.Assignees,pi.TimeRemaining,pi.IssueUrl)
}