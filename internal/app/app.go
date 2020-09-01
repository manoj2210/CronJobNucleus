package app

import (
	"Webhooks/internal/models"
	"Webhooks/internal/services"
	"context"
	"github.com/google/go-github/github"
	"log"
	"os"
)

func StartApplication(githubClient *github.Client)  {
	ctx := context.Background()
	issues, _, err := githubClient.Issues.ListByRepo(ctx,
		os.Getenv("githubOwner"),os.Getenv("githubRepo"),nil)
	if err != nil{
		log.Fatal(err)
	}

	PostBody := models.NewPostBody("Pending Issues")
	GithubIssues := services.NewGithubIssuesForARepo(issues,os.Getenv("githubRepo"),&PostBody)
	GithubIssues.GetPendingIssues()
}
