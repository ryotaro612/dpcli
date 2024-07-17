package github

type Client struct {
}

type PullRequest struct{}

func (c Client) ReadPullRequests() ([]PullRequest, error) {
	return []PullRequest{}, nil
}
