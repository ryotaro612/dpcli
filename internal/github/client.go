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

type PullRequest struct{}

func (c Client) ReadPullRequests(ctx context.Context) ([]PullRequest, error) {
	reqs, resp, err := c.client.PullRequests.List(ctx, "alpdr", "data-platform", &github.PullRequestListOptions{
		State: "all",
	})

	if err != nil {
		return nil, err
	}
	return []PullRequest{}, nil
}

func NewClient(l *slog.Logger, gitHubToken string) Client {
	client := github.NewClient(nil).WithAuthToken(gitHubToken)
	//		PullRequestsがserviceのこと
	//.PullRequests.List(nil, "", nil)
	return Client{l, client}
}
