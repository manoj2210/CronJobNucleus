package config

import (
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
)

func NewGithubClient()  *github.Client{
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("githubToken")},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}
