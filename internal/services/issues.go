package services

import (
	"Webhooks/internal/helpers"
	"Webhooks/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-github/github"
	"log"
	"os"
	"strings"
	"time"
)

type GithubIssuesForARepo struct {
	Issues []*github.Issue
	PostBody *models.BaseTypeAndAttachment
	Repo string
}

func NewGithubIssuesForARepo(issues []*github.Issue, repo string, postBody *models.BaseTypeAndAttachment) *GithubIssuesForARepo {
	return &GithubIssuesForARepo{issues,postBody,repo}
}

func (gir *GithubIssuesForARepo) GetPendingIssues()  {
	//Filter issues which are open
	filteredIssues:=gir.FilterIssuesByState("open")

	filteredIssues=gir.FilterLabels(os.Getenv("issuesLabel"))

	//Define pending issues
	var pendingIssues []models.PendingIssue

	//Traverse array and GetDeadlines
	for _,x := range filteredIssues{
		date,err:=gir.GetDeadline(x.GetBody())
		var diff time.Duration
		pendingIssueObject:=models.NewPendingIssue(*x.Title,x.Assignees,*x.HTMLURL)
		if err != nil{
			errorMessage:=fmt.Sprintf("Error in Date format. Eg.%v%v%v",
				os.Getenv("deadlineStartFormat"),
				os.Getenv("dateFormat"),
				os.Getenv("deadlineEndFormat"))
			pendingIssueObject.SetTimeRemaining(errorMessage)
		}else {
			diff = helpers.GetTimeDifference(time.Now(),date)
			timeRemaining:=fmt.Sprintf("%v",diff.Round(time.Minute))
			pendingIssueObject.SetTimeRemaining(timeRemaining)
		}
		pendingIssues = append(pendingIssues, pendingIssueObject)
	}

	for _,x:= range pendingIssues{
		eltContainer:= models.NewElementContainer(x.Title,x.TimeRemaining,x.Assignees,x.IssueUrl)
		gir.PostBody.Attachments[0].Content.Body = append(gir.PostBody.Attachments[0].Content.Body, eltContainer)
	}

	var jsonData []byte
	jsonData, err := json.Marshal(gir.PostBody)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))

	//helpers.NewPostRequest(os.Getenv("hookUrl"),jsonData)
}

//func (gir *GithubIssuesForARepo) PendingReviews()  {
//	filteredIssues:=gir.FilterIssuesByState("open")
//	filteredIssues=gir.FilterLabels(os.Getenv("reviewLabel"))
//
//	PostBody := models.NewPostBody("Pending Reviews")
//
//
//	//Define pending issues
//	var pendingIssues []models.PendingIssue
//
//	//Traverse array and GetDeadlines
//	for _,x := range filteredIssues{
//		date,err:=gir.GetDeadline(x.GetBody())
//		var diff time.Duration
//		pendingIssueObject:=models.NewPendingIssue(*x.Title,x.Assignees,*x.HTMLURL)
//		if err != nil{
//			errorMessage:=fmt.Sprintf("Error in Date format. Eg.%v%v%v",
//				os.Getenv("deadlineStartFormat"),
//				os.Getenv("dateFormat"),
//				os.Getenv("deadlineEndFormat"))
//			pendingIssueObject.SetTimeRemaining(errorMessage)
//		}else {
//			diff = helpers.GetTimeDifference(time.Now(),date)
//			timeRemaining:=fmt.Sprintf("%v",diff.Round(time.Minute))
//			pendingIssueObject.SetTimeRemaining(timeRemaining)
//		}
//		pendingIssues = append(pendingIssues, pendingIssueObject)
//	}
//
//	for _,x:= range pendingIssues{
//		eltContainer:= models.NewElementContainer(x.Title,x.TimeRemaining,x.Assignees,x.IssueUrl)
//		gir.PostBody.Attachments[0].Content.Body = append(gir.PostBody.Attachments[0].Content.Body, eltContainer)
//	}
//
//	var jsonData []byte
//	jsonData, err := json.Marshal(gir.PostBody)
//	if err != nil {
//		log.Println(err)
//	}
//	helpers.NewPostRequest(os.Getenv("hookUrl"),jsonData)
//}

func (gir *GithubIssuesForARepo) FilterIssuesByState(state string)  []*github.Issue{
	return helpers.Filter(gir.Issues, func(i interface{}) bool {
		return i.(*github.Issue).GetState() == state
	}).([]*github.Issue)
}

func (gir *GithubIssuesForARepo) FilterLabels(label string)  []*github.Issue{
	return helpers.Filter(gir.Issues, func(i interface{}) bool {
		return helpers.FindIfLabelExists(i.(*github.Issue),label)
	}).([]*github.Issue)
}

func (gir *GithubIssuesForARepo) GetDeadline(body string) (time.Time,error){
	startFormat:=os.Getenv("deadlineStartFormat")
	endFormat:=os.Getenv("deadlineEndFormat")
	split:=strings.Split(strings.Split(body,endFormat)[0],startFormat)
	if len(split) >1{
		deadline:=split[1]
		return time.Parse(os.Getenv("dateFormat"), deadline)
	}else {
		return time.Now(),errors.New("no date")
	}
}