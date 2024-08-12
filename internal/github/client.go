package github

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/go-github/v63/github"
)

type Client struct {
	logger *slog.Logger
	client *github.Client
}

func (c *Client) ReadPullRequests(ctx context.Context, offset time.Time) ([]*github.PullRequest, error) {
	page := 0
	var prs []*github.PullRequest

	done := false
	fetched := map[int64]bool{}

	for !done {
		next, res, err := c.readPullRequestsPage(ctx, page)
		if err != nil {
			return nil, err
		}
		if res.StatusCode/100 != 2 {
			return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
		}
		if len(next) == 0 {
			done = true
		}
		for _, pr := range next {
			if pr.CreatedAt.Before(offset) {
				done = true
				break
			}
			if !fetched[*pr.ID] {
				prs = append(prs, pr)
				fetched[*pr.ID] = true
			}
		}
		page++
	}

	return prs, nil
}

// Instantiate a new client with the given logger and GitHub token.
func NewClient(l *slog.Logger, gitHubToken string) Client {
	client := github.NewClient(nil).WithAuthToken(gitHubToken)
	//		PullRequestsがserviceのこと
	//.PullRequests.List(nil, "", nil)
	return Client{l, client}
}

func (c *Client) readPullRequestsPage(ctx context.Context, page int) ([]*github.PullRequest, *github.Response, error) {
	prs, res, err := c.client.PullRequests.List(ctx, "alpdr", "data-platform", &github.PullRequestListOptions{
		State:       "all",
		ListOptions: github.ListOptions{Page: page},
	})
	return prs, res, err
}
