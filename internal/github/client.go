package github

import (
	"context"
	"log/slog"

	"github.com/google/go-github/v63/github"
)

type Client struct {
	logger *slog.Logger
	client *github.Client
}

func (c Client) ReadPullRequests(ctx context.Context, date Date) ([]*github.PullRequest, error) {
	prs, _, err := c.client.PullRequests.List(ctx, "alpdr", "data-platform", &github.PullRequestListOptions{
		State: "all",
	})

	if err != nil {
		return nil, err
	}
	return prs, nil
}

func NewClient(l *slog.Logger, gitHubToken string) Client {
	client := github.NewClient(nil).WithAuthToken(gitHubToken)
	//		PullRequestsがserviceのこと
	//.PullRequests.List(nil, "", nil)
	return Client{l, client}
}
